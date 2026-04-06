package cli

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/cache"
	"github.com/dl-alexandre/gdrv/internal/logging"
	"github.com/dl-alexandre/gdrv/internal/resolver"
	"github.com/dl-alexandre/gdrv/internal/types"
	"github.com/dl-alexandre/gdrv/pkg/version"
)

// ============================================================
// KONG FOUNDATION - new CLI architecture
// ============================================================

// Globals holds all persistent flags inherited by every command.
// AfterApply runs before any command Run method.
type Globals struct {
	Profile             string         `help:"Authentication profile to use" default:"default" name:"profile"`
	DriveID             string         `help:"Shared Drive ID to operate in" name:"drive-id"`
	Output              string         `help:"Output format (json, table)" default:"json" name:"output"`
	Quiet               bool           `help:"Suppress non-essential output" short:"q" name:"quiet"`
	Verbose             bool           `help:"Enable verbose logging" short:"v" name:"verbose"`
	Debug               bool           `help:"Enable debug output" name:"debug"`
	Strict              bool           `help:"Convert warnings to errors" name:"strict"`
	NoCache             bool           `help:"Bypass path resolution cache" name:"no-cache"`
	CacheTTL            int            `help:"Path cache TTL in seconds" default:"300" name:"cache-ttl"`
	IncludeSharedWithMe bool           `help:"Include shared-with-me items" name:"include-shared-with-me"`
	Config              string         `help:"Path to configuration file" name:"config"`
	LogFile             string         `help:"Path to log file" name:"log-file"`
	DryRun              bool           `help:"Show what would be done without making changes" name:"dry-run"`
	Force               bool           `help:"Force operation without confirmation" short:"f" name:"force"`
	Yes                 bool           `help:"Answer yes to all prompts" short:"y" name:"yes"`
	JSON                bool           `help:"Output in JSON format (alias for --output json)" name:"json"`
	Logger              logging.Logger `kong:"-"`
	Cache               cache.Cache    `kong:"-"`
}

// AfterApply replaces cobra PersistentPreRunE for kong commands.
func (g *Globals) AfterApply() error {
	if g.JSON {
		g.Output = "json"
	}

	if g.Output != string(types.OutputFormatJSON) && g.Output != string(types.OutputFormatTable) {
		return fmt.Errorf("invalid output format: %s (must be json or table)", g.Output)
	}

	logConfig := logging.LogConfig{
		Level:           logging.INFO,
		OutputFile:      g.LogFile,
		EnableConsole:   !g.Quiet,
		EnableDebug:     g.Debug,
		RedactSensitive: true,
		EnableColor:     true,
		EnableTimestamp: true,
	}
	if g.Verbose {
		logConfig.Level = logging.DEBUG
	}
	if g.Output == string(types.OutputFormatJSON) && !g.Verbose && !g.Debug {
		logConfig.EnableConsole = false
	}

	var err error
	g.Logger, err = logging.NewLogger(logConfig)
	if err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}

	globalFlags = g.ToGlobalFlags()
	logger = g.Logger

	// Initialize cache for update checking
	if g.Cache == nil {
		g.Cache = cache.NewMemoryCache(cache.MemoryCacheOptions{
			DefaultTTL: time.Duration(g.CacheTTL) * time.Second,
		})
	}

	return nil
}

// ToGlobalFlags converts kong globals to legacy manager-compatible flags.
func (g *Globals) ToGlobalFlags() types.GlobalFlags {
	outputFormat := types.OutputFormatJSON
	if g.Output == string(types.OutputFormatTable) {
		outputFormat = types.OutputFormatTable
	}

	return types.GlobalFlags{
		Profile:             g.Profile,
		DriveID:             g.DriveID,
		OutputFormat:        outputFormat,
		Quiet:               g.Quiet,
		Verbose:             g.Verbose,
		Debug:               g.Debug,
		Strict:              g.Strict,
		NoCache:             g.NoCache,
		CacheTTL:            g.CacheTTL,
		IncludeSharedWithMe: g.IncludeSharedWithMe,
		Config:              g.Config,
		LogFile:             g.LogFile,
		DryRun:              g.DryRun,
		Force:               g.Force,
		Yes:                 g.Yes,
		JSON:                g.JSON,
	}
}

// CLI is the kong root command tree.
type CLI struct {
	Globals

	Version      VersionCmd      `cmd:"" help:"Print the version number"`
	About        AboutCmd        `cmd:"" help:"Display Drive account information and API capabilities"`
	Files        FilesCmd        `cmd:"" help:"File operations"`
	Folders      FoldersCmd      `cmd:"" help:"Folder operations"`
	Auth         AuthCmd         `cmd:"" help:"Authentication commands"`
	Permissions  PermissionsCmd  `cmd:"" aliases:"perm" help:"Permission operations"`
	Drives       DrivesCmd       `cmd:"" help:"Manage Shared Drives"`
	Sheets       SheetsCmd       `cmd:"" help:"Google Sheets operations"`
	Docs         DocsCmd         `cmd:"" help:"Google Docs operations"`
	Slides       SlidesCmd       `cmd:"" help:"Google Slides operations"`
	Admin        AdminCmd        `cmd:"" help:"Google Workspace Admin SDK operations"`
	Groups       GroupsCmd       `cmd:"" help:"Cloud Identity Groups operations (distinct from 'admin groups' which uses Admin SDK)"`
	Changes      ChangesCmd      `cmd:"" help:"Drive Changes API operations"`
	Labels       LabelsCmd       `cmd:"" help:"Drive Labels API operations"`
	Activity     ActivityCmd     `cmd:"" help:"Drive Activity API operations"`
	Forms        FormsCmd        `cmd:"" help:"Google Forms operations"`
	People       PeopleCmd       `cmd:"" help:"Google People/Contacts operations"`
	Chat         ChatCmd         `cmd:"" help:"Google Chat operations"`
	Gmail        GmailCmd        `cmd:"" help:"Gmail operations"`
	Calendar     CalendarCmd     `cmd:"" help:"Google Calendar operations"`
	Tasks        TasksCmd        `cmd:"" help:"Google Tasks operations"`
	AppScript    AppScriptCmd    `cmd:"" help:"Google Apps Script operations"`
	Sync         SyncCmd         `cmd:"" help:"Sync local folders with Drive"`
	Config       ConfigCmd       `cmd:"" help:"Configuration management"`
	Cache        CacheCmd        `cmd:"" help:"Cache management"`
	Completion   CompletionCmd   `cmd:"" help:"Generate shell completion scripts" `
	AI           AICmd           `cmd:"" help:"Google Generative AI (Gemini) operations"`
	Meet         MeetCmd         `cmd:"" help:"Google Meet operations"`
	Logging      CloudLoggingCmd `cmd:"" help:"Cloud Logging operations"`
	Monitoring   MonitoringCmd   `cmd:"" help:"Cloud Monitoring operations"`
	IAMAdmin     IAMAdminCmd     `cmd:"" aliases:"iam-admin" help:"IAM Admin operations"`
	CheckUpdates UpdateCheckCmd  `cmd:"" aliases:"check-update" help:"Check for available updates"`
	TUI          TUICmd          `cmd:"" help:"Interactive TUI file browser"`
}

// VersionCmd prints the version.
type VersionCmd struct{}

func (v *VersionCmd) Run() error {
	fmt.Println(version.Version)
	return nil
}

// globalFlags and logger are set by Globals.AfterApply() during kong initialization
var (
	globalFlags types.GlobalFlags
	logger      logging.Logger
)

// GetGlobalFlags returns the current global flags.
func GetGlobalFlags() types.GlobalFlags {
	return globalFlags
}

// GetLogger returns the current logger.
func GetLogger() logging.Logger {
	return logger
}

// ResolveFileID resolves a file ID from either a direct ID or a path
// If the input starts with "/" or contains "/", it's treated as a path
// Otherwise, it's treated as a direct file ID
func ResolveFileID(ctx context.Context, client *api.Client, flags types.GlobalFlags, fileIDOrPath string) (string, error) {
	// Check if this looks like a path (contains "/" or starts with a path-like name)
	if !isPath(fileIDOrPath) {
		// Treat as direct file ID
		return fileIDOrPath, nil
	}

	// Create path resolver
	cacheTTL := time.Duration(flags.CacheTTL) * time.Second
	if flags.NoCache {
		cacheTTL = 0
	}
	pathResolver := resolver.NewPathResolver(client, cacheTTL)

	// Create request context
	reqCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypeListOrSearch)

	// Resolve path
	result, err := pathResolver.Resolve(ctx, reqCtx, fileIDOrPath, resolver.ResolveOptions{
		DriveID:             flags.DriveID,
		IncludeSharedWithMe: flags.IncludeSharedWithMe,
		UseCache:            !flags.NoCache,
		StrictMode:          flags.Strict,
	})
	if err != nil {
		return "", err
	}

	return result.FileID, nil
}

// isPath determines if the input looks like a path rather than a file ID
func isPath(input string) bool {
	// If it contains "/", it's definitely a path
	if strings.Contains(input, "/") {
		return true
	}
	// Google Drive file IDs are typically long alphanumeric strings
	// Paths typically contain common characters like spaces, dots, etc.
	// If it contains spaces, dots (except at start), or is short, treat as path
	if strings.Contains(input, " ") || strings.Contains(input, ".") {
		return true
	}
	return false
}

// GetPathResolver creates a path resolver with the current flags
func GetPathResolver(client *api.Client, flags types.GlobalFlags) *resolver.PathResolver {
	cacheTTL := time.Duration(flags.CacheTTL) * time.Second
	if flags.NoCache {
		cacheTTL = 0
	}
	return resolver.NewPathResolver(client, cacheTTL)
}

// GetResolveOptions creates resolve options from global flags
func GetResolveOptions(flags types.GlobalFlags) resolver.ResolveOptions {
	return resolver.ResolveOptions{
		DriveID:             flags.DriveID,
		IncludeSharedWithMe: flags.IncludeSharedWithMe,
		UseCache:            !flags.NoCache,
		StrictMode:          flags.Strict,
	}
}
