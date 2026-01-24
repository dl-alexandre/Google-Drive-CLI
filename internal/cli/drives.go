package cli

import (
	"context"

	"github.com/dl-alexandre/gdrive/internal/api"
	"github.com/dl-alexandre/gdrive/internal/auth"
	"github.com/dl-alexandre/gdrive/internal/drives"
	"github.com/dl-alexandre/gdrive/internal/types"
	"github.com/dl-alexandre/gdrive/internal/utils"
	"github.com/spf13/cobra"
)

var drivesCmd = &cobra.Command{
	Use:   "drives",
	Short: "Manage Shared Drives",
	Long:  `Commands for listing and managing Shared Drives (formerly Team Drives).`,
}

var drivesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Shared Drives",
	Long:  `List all Shared Drives accessible by the authenticated user.`,
	RunE:  runDrivesList,
}

var drivesGetCmd = &cobra.Command{
	Use:   "get <drive-id>",
	Short: "Get Shared Drive details",
	Long:  `Get details of a specific Shared Drive by its ID.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runDrivesGet,
}

var (
	drivesListPageSize  int
	drivesListPageToken string
	drivesListPaginate  bool
)

func init() {
	rootCmd.AddCommand(drivesCmd)
	drivesCmd.AddCommand(drivesListCmd)
	drivesCmd.AddCommand(drivesGetCmd)

	drivesListCmd.Flags().IntVar(&drivesListPageSize, "page-size", 100, "Maximum number of drives to return per page")
	drivesListCmd.Flags().StringVar(&drivesListPageToken, "page-token", "", "Page token for pagination")
	drivesListCmd.Flags().BoolVar(&drivesListPaginate, "paginate", false, "Automatically fetch all pages")
}

func runDrivesList(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	flags := GetGlobalFlags()

	writer := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	// Get authenticated client
	client, err := getAPIClient(ctx, flags.Profile)
	if err != nil {
		return handleError(writer, "drives list", err)
	}

	// Create drives manager
	manager := drives.NewManager(client)

	// Create request context
	reqCtx := api.NewRequestContext(flags.Profile, "", types.RequestTypeListOrSearch)

	// If --paginate flag is set, fetch all pages
	if drivesListPaginate {
		var allDrives []*drives.SharedDrive
		pageToken := drivesListPageToken
		for {
			result, err := manager.List(ctx, reqCtx, drivesListPageSize, pageToken)
			if err != nil {
				return handleError(writer, "drives list", err)
			}
			allDrives = append(allDrives, result.Drives...)
			if result.NextPageToken == "" {
				break
			}
			pageToken = result.NextPageToken
		}
		return writer.WriteSuccess("drives list", map[string]interface{}{
			"drives": allDrives,
		})
	}

	// List drives
	result, err := manager.List(ctx, reqCtx, drivesListPageSize, drivesListPageToken)
	if err != nil {
		return handleError(writer, "drives list", err)
	}

	// Output result
	return writer.WriteSuccess("drives list", result)
}

func runDrivesGet(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	flags := GetGlobalFlags()
	driveID := args[0]

	writer := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	// Get authenticated client
	client, err := getAPIClient(ctx, flags.Profile)
	if err != nil {
		return handleError(writer, "drives get", err)
	}

	// Create drives manager
	manager := drives.NewManager(client)

	// Create request context
	reqCtx := api.NewRequestContext(flags.Profile, driveID, types.RequestTypeGetByID)

	// Get drive
	result, err := manager.Get(ctx, reqCtx, driveID, "")
	if err != nil {
		return handleError(writer, "drives get", err)
	}

	// Output result
	return writer.WriteSuccess("drives get", result)
}

// getAPIClient creates an API client for the given profile
func getAPIClient(ctx context.Context, profile string) (*api.Client, error) {
	// Get config directory
	configDir := getConfigDir()

	// Get auth manager
	authMgr := auth.NewManager(configDir)

	// Get valid credentials for profile
	creds, err := authMgr.GetValidCredentials(ctx, profile)
	if err != nil {
		return nil, err
	}

	// Create Drive service
	driveService, err := authMgr.GetDriveService(ctx, creds)
	if err != nil {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeAuthRequired,
			"Failed to create Drive service: "+err.Error()).Build())
	}

	// Create API client
	return api.NewClient(driveService, utils.DefaultMaxRetries, utils.DefaultRetryDelayMs, GetLogger()), nil
}

// handleError converts errors to CLI output
func handleError(writer *OutputWriter, command string, err error) error {
	if appErr, ok := err.(*utils.AppError); ok {
		return writer.WriteError(command, appErr.CLIError)
	}
	return writer.WriteError(command, utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
}
