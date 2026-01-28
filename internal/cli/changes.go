package cli

import (
	"context"
	"fmt"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/auth"
	"github.com/dl-alexandre/gdrv/internal/changes"
	"github.com/dl-alexandre/gdrv/internal/types"
	"github.com/dl-alexandre/gdrv/internal/utils"
	"github.com/spf13/cobra"
)

var changesCmd = &cobra.Command{
	Use:   "changes",
	Short: "Drive Changes API operations",
	Long:  "Track changes to files and folders for real-time synchronization and automation",
}

var changesStartPageTokenCmd = &cobra.Command{
	Use:   "start-page-token",
	Short: "Get the starting page token",
	Long: `Get the starting page token for listing changes.

This token represents the current state of the Drive and can be used
to track all future changes.

Examples:
  # Get the starting page token
  gdrv changes start-page-token --json

  # Get the starting page token for a Shared Drive
  gdrv changes start-page-token --drive-id <drive-id> --json`,
	RunE: runChangesStartPageToken,
}

var changesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List changes since a page token",
	Long: `List changes to files and folders since a given page token.

Use the page token from start-page-token or from a previous list response
to track incremental changes.

Examples:
  # List changes since a page token
  gdrv changes list --page-token "12345" --json

  # List changes with auto-pagination
  gdrv changes list --page-token "12345" --paginate --json

  # List changes for a specific Shared Drive
  gdrv changes list --page-token "12345" --drive-id <drive-id> --json

  # List changes including removed files
  gdrv changes list --page-token "12345" --include-removed --json

  # List changes with specific fields
  gdrv changes list --page-token "12345" --fields "nextPageToken,newStartPageToken,changes(fileId,time,removed,file(name,mimeType))" --json`,
	RunE: runChangesList,
}

var changesWatchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch for changes via webhook",
	Long: `Set up a webhook to receive notifications when changes occur.

The webhook URL must be accessible from the internet and must use HTTPS.
Google will send POST requests to this URL when changes occur.

Examples:
  # Watch for changes
  gdrv changes watch --page-token "12345" --webhook-url "https://example.com/webhook" --json

  # Watch with custom expiration (Unix timestamp in milliseconds)
  gdrv changes watch --page-token "12345" --webhook-url "https://example.com/webhook" --expiration 1706745600000 --json

  # Watch with a token for verification
  gdrv changes watch --page-token "12345" --webhook-url "https://example.com/webhook" --token "my-secret-token" --json`,
	RunE: runChangesWatch,
}

var changesStopCmd = &cobra.Command{
	Use:   "stop <channel-id> <resource-id>",
	Short: "Stop watching for changes",
	Long: `Stop a notification channel created by the watch command.

The channel-id and resource-id are returned by the watch command.

Examples:
  # Stop watching for changes
  gdrv changes stop <channel-id> <resource-id>`,
	Args: cobra.ExactArgs(2),
	RunE: runChangesStop,
}

var (
	changesPageToken                 string
	changesIncludeCorpusRemovals     bool
	changesIncludeItemsFromAllDrives bool
	changesIncludePermissionsForView string
	changesIncludeRemoved            bool
	changesRestrictToMyDrive         bool
	changesSupportsAllDrives         bool
	changesLimit                     int
	changesFields                    string
	changesSpaces                    string
	changesPaginate                  bool
	changesWebhookURL                string
	changesExpiration                int64
	changesToken                     string
)

func init() {
	changesStartPageTokenCmd.Flags().StringVar(&globalFlags.DriveID, "drive-id", "", "Shared Drive ID")

	changesListCmd.Flags().StringVar(&changesPageToken, "page-token", "", "Page token to list changes from (required)")
	changesListCmd.Flags().StringVar(&globalFlags.DriveID, "drive-id", "", "Shared Drive ID")
	changesListCmd.Flags().BoolVar(&changesIncludeCorpusRemovals, "include-corpus-removals", false, "Include changes outside target corpus")
	changesListCmd.Flags().BoolVar(&changesIncludeItemsFromAllDrives, "include-items-from-all-drives", false, "Include items from all drives")
	changesListCmd.Flags().StringVar(&changesIncludePermissionsForView, "include-permissions-for-view", "", "Include permissions with published view")
	changesListCmd.Flags().BoolVar(&changesIncludeRemoved, "include-removed", false, "Include removed items")
	changesListCmd.Flags().BoolVar(&changesRestrictToMyDrive, "restrict-to-my-drive", false, "Restrict to My Drive only")
	changesListCmd.Flags().BoolVar(&changesSupportsAllDrives, "supports-all-drives", true, "Support all drives")
	changesListCmd.Flags().IntVar(&changesLimit, "limit", 100, "Maximum results per page")
	changesListCmd.Flags().StringVar(&changesFields, "fields", "", "Fields to return")
	changesListCmd.Flags().StringVar(&changesSpaces, "spaces", "", "Comma-separated list of spaces (drive, appDataFolder, photos)")
	changesListCmd.Flags().BoolVar(&changesPaginate, "paginate", false, "Auto-paginate through all changes")
	changesListCmd.MarkFlagRequired("page-token")

	changesWatchCmd.Flags().StringVar(&changesPageToken, "page-token", "", "Page token to watch from (required)")
	changesWatchCmd.Flags().StringVar(&changesWebhookURL, "webhook-url", "", "Webhook URL for notifications (required)")
	changesWatchCmd.Flags().StringVar(&globalFlags.DriveID, "drive-id", "", "Shared Drive ID")
	changesWatchCmd.Flags().BoolVar(&changesIncludeCorpusRemovals, "include-corpus-removals", false, "Include changes outside target corpus")
	changesWatchCmd.Flags().BoolVar(&changesIncludeItemsFromAllDrives, "include-items-from-all-drives", false, "Include items from all drives")
	changesWatchCmd.Flags().StringVar(&changesIncludePermissionsForView, "include-permissions-for-view", "", "Include permissions with published view")
	changesWatchCmd.Flags().BoolVar(&changesIncludeRemoved, "include-removed", false, "Include removed items")
	changesWatchCmd.Flags().BoolVar(&changesRestrictToMyDrive, "restrict-to-my-drive", false, "Restrict to My Drive only")
	changesWatchCmd.Flags().BoolVar(&changesSupportsAllDrives, "supports-all-drives", true, "Support all drives")
	changesWatchCmd.Flags().StringVar(&changesSpaces, "spaces", "", "Comma-separated list of spaces (drive, appDataFolder, photos)")
	changesWatchCmd.Flags().Int64Var(&changesExpiration, "expiration", 0, "Webhook expiration time (Unix timestamp in milliseconds)")
	changesWatchCmd.Flags().StringVar(&changesToken, "token", "", "Arbitrary token for webhook verification")
	changesWatchCmd.MarkFlagRequired("page-token")
	changesWatchCmd.MarkFlagRequired("webhook-url")

	changesCmd.AddCommand(changesStartPageTokenCmd)
	changesCmd.AddCommand(changesListCmd)
	changesCmd.AddCommand(changesWatchCmd)
	changesCmd.AddCommand(changesStopCmd)
	rootCmd.AddCommand(changesCmd)
}

func getChangesManager(ctx context.Context, flags types.GlobalFlags) (*changes.Manager, *api.Client, *types.RequestContext, *OutputWriter, error) {
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	configDir := getConfigDir()
	authMgr := auth.NewManager(configDir)

	creds, err := authMgr.GetValidCredentials(ctx, flags.Profile)
	if err != nil {
		return nil, nil, nil, out, err
	}

	service, err := authMgr.GetDriveService(ctx, creds)
	if err != nil {
		return nil, nil, nil, out, err
	}

	client := api.NewClient(service, utils.DefaultMaxRetries, utils.DefaultRetryDelayMs, GetLogger())
	mgr := changes.NewManager(client)
	reqCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypeListOrSearch)

	return mgr, client, reqCtx, out, nil
}

func runChangesStartPageToken(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	mgr, _, reqCtx, out, err := getChangesManager(ctx, globalFlags)
	if err != nil {
		return err
	}

	token, err := mgr.GetStartPageToken(ctx, reqCtx, globalFlags.DriveID)
	if err != nil {
		return err
	}

	if globalFlags.OutputFormat == types.OutputFormatJSON {
		return out.WriteSuccess("changes start-page-token", map[string]string{"startPageToken": token})
	}

	fmt.Println(token)
	return nil
}

func runChangesList(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	mgr, _, reqCtx, out, err := getChangesManager(ctx, globalFlags)
	if err != nil {
		return err
	}

	opts := types.ListOptions{
		PageToken:                 changesPageToken,
		DriveID:                   globalFlags.DriveID,
		IncludeCorpusRemovals:     changesIncludeCorpusRemovals,
		IncludeItemsFromAllDrives: changesIncludeItemsFromAllDrives,
		IncludePermissionsForView: changesIncludePermissionsForView,
		IncludeRemoved:            changesIncludeRemoved,
		RestrictToMyDrive:         changesRestrictToMyDrive,
		SupportsAllDrives:         changesSupportsAllDrives,
		Limit:                     changesLimit,
		Fields:                    changesFields,
		Spaces:                    changesSpaces,
	}

	if changesPaginate {
		return paginateChanges(ctx, mgr, reqCtx, out, opts)
	}

	result, err := mgr.List(ctx, reqCtx, opts)
	if err != nil {
		return err
	}

	if globalFlags.OutputFormat == types.OutputFormatJSON {
		return out.WriteSuccess("changes list", result)
	}

	if len(result.Changes) == 0 {
		fmt.Println("No changes found")
		return nil
	}

	fmt.Printf("Found %d change(s)\n", len(result.Changes))
	for _, change := range result.Changes {
		if change.Removed {
			fmt.Printf("  [REMOVED] %s (File ID: %s) at %s\n", change.ChangeType, change.FileID, change.Time.Format("2006-01-02 15:04:05"))
		} else if change.File != nil {
			fmt.Printf("  [%s] %s (ID: %s) at %s\n", change.ChangeType, change.File.Name, change.FileID, change.Time.Format("2006-01-02 15:04:05"))
		} else if change.Drive != nil {
			fmt.Printf("  [%s] %s (ID: %s) at %s\n", change.ChangeType, change.Drive.Name, change.DriveID, change.Time.Format("2006-01-02 15:04:05"))
		}
	}

	if result.NextPageToken != "" {
		fmt.Printf("\nNext page token: %s\n", result.NextPageToken)
	}

	if result.NewStartPageToken != "" {
		fmt.Printf("New start page token: %s\n", result.NewStartPageToken)
	}

	return nil
}

func paginateChanges(ctx context.Context, mgr *changes.Manager, reqCtx *types.RequestContext, out *OutputWriter, opts types.ListOptions) error {
	allChanges := []types.Change{}
	pageToken := opts.PageToken
	var newStartPageToken string

	for {
		opts.PageToken = pageToken
		result, err := mgr.List(ctx, reqCtx, opts)
		if err != nil {
			return err
		}

		allChanges = append(allChanges, result.Changes...)

		if result.NewStartPageToken != "" {
			newStartPageToken = result.NewStartPageToken
		}

		if result.NextPageToken == "" {
			break
		}

		pageToken = result.NextPageToken
	}

	if globalFlags.OutputFormat == types.OutputFormatJSON {
		return out.WriteSuccess("changes list", map[string]interface{}{
			"changes":           allChanges,
			"newStartPageToken": newStartPageToken,
			"totalChanges":      len(allChanges),
		})
	}

	fmt.Printf("Found %d total change(s)\n", len(allChanges))
	for _, change := range allChanges {
		if change.Removed {
			fmt.Printf("  [REMOVED] %s (File ID: %s) at %s\n", change.ChangeType, change.FileID, change.Time.Format("2006-01-02 15:04:05"))
		} else if change.File != nil {
			fmt.Printf("  [%s] %s (ID: %s) at %s\n", change.ChangeType, change.File.Name, change.FileID, change.Time.Format("2006-01-02 15:04:05"))
		} else if change.Drive != nil {
			fmt.Printf("  [%s] %s (ID: %s) at %s\n", change.ChangeType, change.Drive.Name, change.DriveID, change.Time.Format("2006-01-02 15:04:05"))
		}
	}

	if newStartPageToken != "" {
		fmt.Printf("\nNew start page token: %s\n", newStartPageToken)
	}

	return nil
}

func runChangesWatch(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	mgr, _, reqCtx, out, err := getChangesManager(ctx, globalFlags)
	if err != nil {
		return err
	}

	opts := types.WatchOptions{
		PageToken:                 changesPageToken,
		DriveID:                   globalFlags.DriveID,
		IncludeCorpusRemovals:     changesIncludeCorpusRemovals,
		IncludeItemsFromAllDrives: changesIncludeItemsFromAllDrives,
		IncludePermissionsForView: changesIncludePermissionsForView,
		IncludeRemoved:            changesIncludeRemoved,
		RestrictToMyDrive:         changesRestrictToMyDrive,
		SupportsAllDrives:         changesSupportsAllDrives,
		Spaces:                    changesSpaces,
		WebhookURL:                changesWebhookURL,
		Expiration:                changesExpiration,
		Token:                     changesToken,
	}

	channel, err := mgr.Watch(ctx, reqCtx, changesPageToken, changesWebhookURL, opts)
	if err != nil {
		return err
	}

	if globalFlags.OutputFormat == types.OutputFormatJSON {
		return out.WriteSuccess("changes watch", channel)
	}

	fmt.Printf("Watching for changes\n")
	fmt.Printf("Channel ID: %s\n", channel.ID)
	fmt.Printf("Resource ID: %s\n", channel.ResourceID)
	if channel.Expiration > 0 {
		fmt.Printf("Expiration: %d\n", channel.Expiration)
	}
	fmt.Printf("\nTo stop watching, run:\n")
	fmt.Printf("  gdrv changes stop %s %s\n", channel.ID, channel.ResourceID)

	return nil
}

func runChangesStop(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	mgr, _, reqCtx, out, err := getChangesManager(ctx, globalFlags)
	if err != nil {
		return err
	}

	channelID := args[0]
	resourceID := args[1]

	err = mgr.Stop(ctx, reqCtx, channelID, resourceID)
	if err != nil {
		return err
	}

	if globalFlags.OutputFormat == types.OutputFormatJSON {
		return out.WriteSuccess("changes stop", map[string]string{
			"status":     "stopped",
			"channelId":  channelID,
			"resourceId": resourceID,
		})
	}

	fmt.Printf("Stopped watching channel %s\n", channelID)
	return nil
}
