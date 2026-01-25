package cli

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/auth"
	"github.com/dl-alexandre/gdrv/internal/sync/conflict"
	syncengine "github.com/dl-alexandre/gdrv/internal/sync"
	"github.com/dl-alexandre/gdrv/internal/sync/diff"
	"github.com/dl-alexandre/gdrv/internal/sync/index"
	"github.com/dl-alexandre/gdrv/internal/types"
	"github.com/dl-alexandre/gdrv/internal/utils"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync <config-id>",
	Short: "Sync local folders with Drive",
	Args:  cobra.ExactArgs(1),
	RunE:  runSyncBidirectional,
}

var syncInitCmd = &cobra.Command{
	Use:   "init <local-path> <remote-folder>",
	Short: "Initialize a sync configuration",
	Args:  cobra.ExactArgs(2),
	RunE:  runSyncInit,
}

var syncPushCmd = &cobra.Command{
	Use:   "push <config-id>",
	Short: "Push local changes to Drive",
	Args:  cobra.ExactArgs(1),
	RunE:  runSyncPush,
}

var syncPullCmd = &cobra.Command{
	Use:   "pull <config-id>",
	Short: "Pull remote changes to local",
	Args:  cobra.ExactArgs(1),
	RunE:  runSyncPull,
}

var syncStatusCmd = &cobra.Command{
	Use:   "status <config-id>",
	Short: "Show pending sync changes",
	Args:  cobra.ExactArgs(1),
	RunE:  runSyncStatus,
}

var syncListCmd = &cobra.Command{
	Use:   "list",
	Short: "List sync configurations",
	RunE:  runSyncList,
}

var syncRemoveCmd = &cobra.Command{
	Use:   "remove <config-id>",
	Short: "Remove a sync configuration",
	Args:  cobra.ExactArgs(1),
	RunE:  runSyncRemove,
}

var (
	syncExclude     string
	syncConflict    string
	syncDirection   string
	syncConfigID    string
	syncDelete      bool
	syncConcurrency int
	syncUseChanges  bool
)

func init() {
	syncInitCmd.Flags().StringVar(&syncExclude, "exclude", "", "Comma-separated exclude patterns")
	syncInitCmd.Flags().StringVar(&syncConflict, "conflict", "rename-both", "Conflict policy (local-wins, remote-wins, rename-both)")
	syncInitCmd.Flags().StringVar(&syncDirection, "direction", "bidirectional", "Sync direction (push, pull, bidirectional)")
	syncInitCmd.Flags().StringVar(&syncConfigID, "id", "", "Optional sync configuration ID")

	syncCmd.Flags().BoolVar(&syncDelete, "delete", false, "Propagate deletions")
	syncCmd.Flags().StringVar(&syncConflict, "conflict", "", "Override conflict policy")
	syncCmd.Flags().IntVar(&syncConcurrency, "concurrency", 5, "Concurrent transfers")
	syncCmd.Flags().BoolVar(&syncUseChanges, "use-changes", true, "Use Drive Changes API when available")

	syncPushCmd.Flags().BoolVar(&syncDelete, "delete", false, "Propagate deletions")
	syncPushCmd.Flags().StringVar(&syncConflict, "conflict", "", "Override conflict policy")
	syncPushCmd.Flags().IntVar(&syncConcurrency, "concurrency", 5, "Concurrent transfers")
	syncPushCmd.Flags().BoolVar(&syncUseChanges, "use-changes", true, "Use Drive Changes API when available")

	syncPullCmd.Flags().BoolVar(&syncDelete, "delete", false, "Propagate deletions")
	syncPullCmd.Flags().StringVar(&syncConflict, "conflict", "", "Override conflict policy")
	syncPullCmd.Flags().IntVar(&syncConcurrency, "concurrency", 5, "Concurrent transfers")
	syncPullCmd.Flags().BoolVar(&syncUseChanges, "use-changes", true, "Use Drive Changes API when available")

	syncStatusCmd.Flags().BoolVar(&syncDelete, "delete", false, "Include deletions in status")
	syncStatusCmd.Flags().StringVar(&syncConflict, "conflict", "", "Override conflict policy")
	syncStatusCmd.Flags().BoolVar(&syncUseChanges, "use-changes", true, "Use Drive Changes API when available")

	syncCmd.AddCommand(syncInitCmd)
	syncCmd.AddCommand(syncPushCmd)
	syncCmd.AddCommand(syncPullCmd)
	syncCmd.AddCommand(syncStatusCmd)
	syncCmd.AddCommand(syncListCmd)
	syncCmd.AddCommand(syncRemoveCmd)
	rootCmd.AddCommand(syncCmd)
}

func runSyncInit(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	localPath := args[0]
	remotePath := args[1]

	stat, err := os.Stat(localPath)
	if err != nil || !stat.IsDir() {
		return out.WriteError("sync.init", utils.NewCLIError(utils.ErrCodeInvalidArgument, "Local path must be a directory").Build())
	}

	absLocal, err := filepath.Abs(localPath)
	if err != nil {
		return out.WriteError("sync.init", utils.NewCLIError(utils.ErrCodeInvalidArgument, err.Error()).Build())
	}

	_, client, _, _, err := getFileManager(ctx, flags)
	if err != nil {
		return out.WriteError("sync.init", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	remoteID, err := ResolveFileID(ctx, client, flags, remotePath)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("sync.init", appErr.CLIError)
		}
		return out.WriteError("sync.init", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	db, err := openSyncDB()
	if err != nil {
		return out.WriteError("sync.init", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}
	defer db.Close()

	configID := syncConfigID
	if configID == "" {
		configID = uuid.New().String()
	}

	excludes := []string{}
	if syncExclude != "" {
		for _, part := range strings.Split(syncExclude, ",") {
			part = strings.TrimSpace(part)
			if part != "" {
				excludes = append(excludes, part)
			}
		}
	}

	cfg := index.SyncConfig{
		ID:              configID,
		LocalRoot:       absLocal,
		RemoteRootID:    remoteID,
		ExcludePatterns: excludes,
		ConflictPolicy:  syncConflict,
		Direction:       syncDirection,
	}
	if err := syncengine.EnsureConfig(&cfg); err != nil {
		return out.WriteError("sync.init", utils.NewCLIError(utils.ErrCodeInvalidArgument, err.Error()).Build())
	}

	if err := db.UpsertConfig(ctx, cfg); err != nil {
		return out.WriteError("sync.init", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("sync.init", cfg)
}

func runSyncBidirectional(cmd *cobra.Command, args []string) error {
	return runSyncWithMode(args[0], diff.ModeBidirectional, "sync", false)
}

func runSyncPush(cmd *cobra.Command, args []string) error {
	return runSyncWithMode(args[0], diff.ModePush, "sync.push", false)
}

func runSyncPull(cmd *cobra.Command, args []string) error {
	return runSyncWithMode(args[0], diff.ModePull, "sync.pull", false)
}

func runSyncStatus(cmd *cobra.Command, args []string) error {
	return runSyncWithMode(args[0], diff.ModeBidirectional, "sync.status", true)
}

func runSyncWithMode(configID string, mode diff.Mode, command string, planOnly bool) error {
	flags := GetGlobalFlags()
	ctx := context.Background()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	engine, reqCtx, cfg, err := loadSyncEngine(ctx, flags, configID)
	if err != nil {
		return out.WriteError(command, utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}
	defer engine.Close()

	opts := syncengine.Options{
		Mode:        mode,
		Delete:      syncDelete,
		DryRun:      flags.DryRun || planOnly,
		Force:       flags.Force,
		Yes:         flags.Yes,
		Concurrency: syncConcurrency,
		UseChanges:  syncUseChanges,
	}
	if syncConflict != "" {
		opts.ConflictPolicy = conflict.Policy(syncConflict)
	}

	plan, err := engine.Plan(ctx, cfg, opts, reqCtx)
	if err != nil {
		return out.WriteError(command, utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	if planOnly {
		return out.WriteSuccess(command, map[string]interface{}{
			"configId":  cfg.ID,
			"actions":   plan.Actions,
			"conflicts": plan.Conflicts,
		})
	}

	if len(plan.Conflicts) > 0 {
		return out.WriteError(command, utils.NewCLIError(utils.ErrCodeUnknown, "Conflicts detected").Build())
	}

	applyCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypeMutation)
	result, err := engine.Apply(ctx, cfg, plan, opts, applyCtx)
	if err != nil {
		return out.WriteError(command, utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess(command, map[string]interface{}{
		"configId": cfg.ID,
		"summary":  result.Summary,
	})
}

func runSyncList(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	db, err := openSyncDB()
	if err != nil {
		return out.WriteError("sync.list", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}
	defer db.Close()

	configs, err := db.ListConfigs(context.Background())
	if err != nil {
		return out.WriteError("sync.list", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("sync.list", map[string]interface{}{
		"configs": configs,
	})
}

func runSyncRemove(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	db, err := openSyncDB()
	if err != nil {
		return out.WriteError("sync.remove", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}
	defer db.Close()

	if err := db.DeleteEntries(context.Background(), args[0]); err != nil {
		return out.WriteError("sync.remove", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}
	if err := db.DeleteConfig(context.Background(), args[0]); err != nil {
		return out.WriteError("sync.remove", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("sync.remove", map[string]interface{}{
		"configId": args[0],
	})
}

func openSyncDB() (*index.DB, error) {
	configDir := getConfigDir()
	dbPath := filepath.Join(configDir, "sync", "index.db")
	return index.Open(dbPath)
}

func loadSyncEngine(ctx context.Context, flags types.GlobalFlags, configID string) (*syncengine.Engine, *types.RequestContext, index.SyncConfig, error) {
	configDir := getConfigDir()
	authMgr := auth.NewManager(configDir)

	creds, err := authMgr.GetValidCredentials(ctx, flags.Profile)
	if err != nil {
		return nil, nil, index.SyncConfig{}, err
	}

	service, err := authMgr.GetDriveService(ctx, creds)
	if err != nil {
		return nil, nil, index.SyncConfig{}, err
	}

	client := api.NewClient(service, utils.DefaultMaxRetries, utils.DefaultRetryDelayMs, GetLogger())
	reqCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypeListOrSearch)

	db, err := openSyncDB()
	if err != nil {
		return nil, nil, index.SyncConfig{}, err
	}

	cfg, err := db.GetConfig(ctx, configID)
	if err != nil {
		_ = db.Close()
		return nil, nil, index.SyncConfig{}, err
	}

	engine := syncengine.NewEngine(client, db)
	return engine, reqCtx, *cfg, nil
}
