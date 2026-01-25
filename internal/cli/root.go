package cli

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/logging"
	"github.com/dl-alexandre/gdrv/internal/resolver"
	"github.com/dl-alexandre/gdrv/internal/types"
	"github.com/dl-alexandre/gdrv/internal/utils"
	"github.com/dl-alexandre/gdrv/pkg/version"
	"github.com/spf13/cobra"
)

var (
	globalFlags types.GlobalFlags
	logger      logging.Logger
)

var rootCmd = &cobra.Command{
	Use:   "gdrv",
	Short: "Google Drive CLI - Command line interface for Google Drive",
	Long: `gdrv is a command-line tool for interacting with Google Drive.
It supports file operations, folder management, permissions, and more.

All commands support JSON output for automation and scripting.`,
	Version: version.Version,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if err := validateGlobalFlags(); err != nil {
			return err
		}

		// Initialize logging
		logConfig := logging.LogConfig{
			Level:           logging.INFO,
			OutputFile:      globalFlags.LogFile,
			EnableConsole:   !globalFlags.Quiet,
			EnableDebug:     globalFlags.Debug,
			RedactSensitive: true,
			EnableColor:     true,
			EnableTimestamp: true,
		}
		if globalFlags.Verbose {
			logConfig.Level = logging.DEBUG
		}
		if globalFlags.OutputFormat == types.OutputFormatJSON && !globalFlags.Verbose && !globalFlags.Debug {
			logConfig.EnableConsole = false
		}

		var err error
		logger, err = logging.NewLogger(logConfig)
		if err != nil {
			return fmt.Errorf("failed to initialize logger: %w", err)
		}

		return nil
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  "Print the version number of gdrv",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Version)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&globalFlags.Profile, "profile", "default", "Authentication profile to use")
	rootCmd.PersistentFlags().StringVar(&globalFlags.DriveID, "drive-id", "", "Shared Drive ID to operate in")
	rootCmd.PersistentFlags().StringVar((*string)(&globalFlags.OutputFormat), "output", "json", "Output format (json, table)")
	rootCmd.PersistentFlags().BoolVarP(&globalFlags.Quiet, "quiet", "q", false, "Suppress non-essential output")
	rootCmd.PersistentFlags().BoolVarP(&globalFlags.Verbose, "verbose", "v", false, "Enable verbose logging")
	rootCmd.PersistentFlags().BoolVar(&globalFlags.Debug, "debug", false, "Enable debug output")
	rootCmd.PersistentFlags().BoolVar(&globalFlags.Strict, "strict", false, "Convert warnings to errors")
	rootCmd.PersistentFlags().BoolVar(&globalFlags.NoCache, "no-cache", false, "Bypass path resolution cache")
	rootCmd.PersistentFlags().IntVar(&globalFlags.CacheTTL, "cache-ttl", 300, "Path cache TTL in seconds")
	rootCmd.PersistentFlags().BoolVar(&globalFlags.IncludeSharedWithMe, "include-shared-with-me", false, "Include shared-with-me items")
	rootCmd.PersistentFlags().StringVar(&globalFlags.Config, "config", "", "Path to configuration file")
	rootCmd.PersistentFlags().StringVar(&globalFlags.LogFile, "log-file", "", "Path to log file")
	rootCmd.PersistentFlags().BoolVar(&globalFlags.DryRun, "dry-run", false, "Show what would be done without making changes")
	rootCmd.PersistentFlags().BoolVarP(&globalFlags.Force, "force", "f", false, "Force operation without confirmation")
	rootCmd.PersistentFlags().BoolVarP(&globalFlags.Yes, "yes", "y", false, "Answer yes to all prompts")
	rootCmd.PersistentFlags().BoolVar(&globalFlags.JSON, "json", false, "Output in JSON format (alias for --output json)")

	// Add subcommands
	rootCmd.AddCommand(versionCmd)
}

func validateGlobalFlags() error {
	// Handle --json flag as alias for --output json
	if globalFlags.JSON {
		globalFlags.OutputFormat = types.OutputFormatJSON
	}

	if globalFlags.OutputFormat != types.OutputFormatJSON && globalFlags.OutputFormat != types.OutputFormatTable {
		return fmt.Errorf("invalid output format: %s", globalFlags.OutputFormat)
	}
	return nil
}

// Execute runs the root command
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(utils.ExitUnknown)
	}
	return nil
}

// GetGlobalFlags returns the global flags
func GetGlobalFlags() types.GlobalFlags {
	return globalFlags
}

// GetLogger returns the global logger
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
