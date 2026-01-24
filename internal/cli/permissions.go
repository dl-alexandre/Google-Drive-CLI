package cli

import (
	"context"
	"os"

	"github.com/dl-alexandre/gdrive/internal/api"
	"github.com/dl-alexandre/gdrive/internal/auth"
	"github.com/dl-alexandre/gdrive/internal/permissions"
	"github.com/dl-alexandre/gdrive/internal/types"
	"github.com/dl-alexandre/gdrive/internal/utils"
	"github.com/spf13/cobra"
)

var permissionsCmd = &cobra.Command{
	Use:     "permissions",
	Aliases: []string{"perm"},
	Short:   "Permission operations",
	Long:    "Commands for managing file and folder permissions in Google Drive",
}

var permListCmd = &cobra.Command{
	Use:   "list <file-id>",
	Short: "List permissions",
	Long:  "List all permissions for a file or folder",
	Args:  cobra.ExactArgs(1),
	RunE:  runPermList,
}

var permCreateCmd = &cobra.Command{
	Use:   "create <file-id>",
	Short: "Create a permission",
	Long:  "Create a new permission to a file or folder",
	Args:  cobra.ExactArgs(1),
	RunE:  runPermCreate,
}

var permUpdateCmd = &cobra.Command{
	Use:   "update <file-id> <permission-id>",
	Short: "Update a permission",
	Long:  "Update an existing permission's role",
	Args:  cobra.ExactArgs(2),
	RunE:  runPermUpdate,
}

var permRemoveCmd = &cobra.Command{
	Use:   "remove <file-id> <permission-id>",
	Short: "Remove a permission",
	Long:  "Remove a permission from a file or folder",
	Args:  cobra.ExactArgs(2),
	RunE:  runPermRemove,
}

var permCreateLinkCmd = &cobra.Command{
	Use:   "create-link <file-id>",
	Short: "Create a public link",
	Long:  "Create a public sharing link for a file or folder",
	Args:  cobra.ExactArgs(1),
	RunE:  runPermCreateLink,
}

// Flags
var (
	permType               string
	permRole               string
	permEmail              string
	permDomain             string
	permSendNotification   bool
	permEmailMessage       string
	permTransferOwnership  bool
	permAllowFileDiscovery bool
)

func init() {
	rootCmd.AddCommand(permissionsCmd)

	permissionsCmd.AddCommand(permListCmd)
	permissionsCmd.AddCommand(permCreateCmd)
	permissionsCmd.AddCommand(permUpdateCmd)
	permissionsCmd.AddCommand(permRemoveCmd)
	permissionsCmd.AddCommand(permCreateLinkCmd)

	// Create flags
	permCreateCmd.Flags().StringVar(&permType, "type", "", "Permission type (user, group, domain, anyone)")
	permCreateCmd.Flags().StringVar(&permRole, "role", "", "Permission role (reader, commenter, writer, organizer)")
	permCreateCmd.Flags().StringVar(&permEmail, "email", "", "Email address (for user/group type)")
	permCreateCmd.Flags().StringVar(&permDomain, "domain", "", "Domain (for domain type)")
	permCreateCmd.Flags().BoolVar(&permSendNotification, "send-notification", true, "Send email notification")
	permCreateCmd.Flags().StringVar(&permEmailMessage, "message", "", "Custom email message")
	permCreateCmd.Flags().BoolVar(&permTransferOwnership, "transfer-ownership", false, "Transfer ownership (requires owner role)")
	permCreateCmd.Flags().BoolVar(&permAllowFileDiscovery, "allow-discovery", false, "Allow file discovery (for anyone type)")
	_ = permCreateCmd.MarkFlagRequired("type")
	_ = permCreateCmd.MarkFlagRequired("role")

	// Update flags
	permUpdateCmd.Flags().StringVar(&permRole, "role", "", "New permission role")
	_ = permUpdateCmd.MarkFlagRequired("role")

	// Create link flags
	permCreateLinkCmd.Flags().StringVar(&permRole, "role", "reader", "Permission role (reader, commenter, writer)")
	permCreateLinkCmd.Flags().BoolVar(&permAllowFileDiscovery, "allow-discovery", false, "Allow file discovery in search")
}

func getPermissionManager() (*permissions.Manager, error) {
	flags := GetGlobalFlags()

	configDir := getConfigDir()
	authMgr := auth.NewManager(configDir)
	creds, err := authMgr.LoadCredentials(flags.Profile)
	if err != nil {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeAuthRequired,
			"Authentication required. Run 'gdrive auth login' first.").Build())
	}

	service, err := authMgr.GetDriveService(context.Background(), creds)
	if err != nil {
		return nil, err
	}

	client := api.NewClient(service, utils.DefaultMaxRetries, utils.DefaultRetryDelayMs, GetLogger())
	return permissions.NewManager(client), nil
}

func runPermList(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	writer := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	mgr, err := getPermissionManager()
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return writer.WriteError("permission.list", appErr.CLIError)
		}
		return writer.WriteError("permission.list", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	reqCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypePermissionOp)
	fileID := args[0]

	result, err := mgr.List(context.Background(), reqCtx, fileID, permissions.ListOptions{})
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			os.Exit(utils.GetExitCode(appErr.CLIError.Code))
			return writer.WriteError("permission.list", appErr.CLIError)
		}
		return writer.WriteError("permission.list", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return writer.WriteSuccess("permission.list", result)
}

func runPermCreate(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	writer := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	// Validate type
	validTypes := map[string]bool{"user": true, "group": true, "domain": true, "anyone": true}
	if !validTypes[permType] {
		return writer.WriteError("permissions.create", utils.NewCLIError(utils.ErrCodeInvalidArgument,
			"Invalid permission type. Must be one of: user, group, domain, anyone").Build())
	}

	// Validate role
	validRoles := map[string]bool{"reader": true, "commenter": true, "writer": true, "organizer": true, "owner": true}
	if !validRoles[permRole] {
		return writer.WriteError("permissions.create", utils.NewCLIError(utils.ErrCodeInvalidArgument,
			"Invalid permission role. Must be one of: reader, commenter, writer, organizer, owner").Build())
	}

	// Validate email for user/group type
	if (permType == "user" || permType == "group") && permEmail == "" {
		return writer.WriteError("permissions.create", utils.NewCLIError(utils.ErrCodeInvalidArgument,
			"Email address is required for user or group permission type").Build())
	}

	// Validate domain for domain type
	if permType == "domain" && permDomain == "" {
		return writer.WriteError("permissions.create", utils.NewCLIError(utils.ErrCodeInvalidArgument,
			"Domain is required for domain permission type").Build())
	}

	mgr, err := getPermissionManager()
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return writer.WriteError("permissions.create", appErr.CLIError)
		}
		return writer.WriteError("permissions.create", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	reqCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypePermissionOp)
	fileID := args[0]

	opts := permissions.CreateOptions{
		Type:                  permType,
		Role:                  permRole,
		EmailAddress:          permEmail,
		Domain:                permDomain,
		SendNotificationEmail: permSendNotification,
		EmailMessage:          permEmailMessage,
		TransferOwnership:     permTransferOwnership,
		AllowFileDiscovery:    permAllowFileDiscovery,
	}

	result, err := mgr.Create(context.Background(), reqCtx, fileID, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			os.Exit(utils.GetExitCode(appErr.CLIError.Code))
			return writer.WriteError("permissions.create", appErr.CLIError)
		}
		return writer.WriteError("permissions.create", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return writer.WriteSuccess("permissions.create", result)
}

func runPermUpdate(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	writer := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	// Validate role
	validRoles := map[string]bool{"reader": true, "commenter": true, "writer": true, "organizer": true, "owner": true}
	if !validRoles[permRole] {
		return writer.WriteError("permission.update", utils.NewCLIError(utils.ErrCodeInvalidArgument,
			"Invalid permission role. Must be one of: reader, commenter, writer, organizer, owner").Build())
	}

	mgr, err := getPermissionManager()
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return writer.WriteError("permission.update", appErr.CLIError)
		}
		return writer.WriteError("permission.update", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	reqCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypePermissionOp)
	fileID := args[0]
	permissionID := args[1]

	result, err := mgr.Update(context.Background(), reqCtx, fileID, permissionID, permissions.UpdateOptions{Role: permRole})
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			os.Exit(utils.GetExitCode(appErr.CLIError.Code))
			return writer.WriteError("permission.update", appErr.CLIError)
		}
		return writer.WriteError("permission.update", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return writer.WriteSuccess("permission.update", result)
}

func runPermRemove(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	writer := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	mgr, err := getPermissionManager()
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return writer.WriteError("permission.remove", appErr.CLIError)
		}
		return writer.WriteError("permission.remove", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	reqCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypePermissionOp)
	fileID := args[0]
	permissionID := args[1]

	err = mgr.Delete(context.Background(), reqCtx, fileID, permissionID, permissions.DeleteOptions{})
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			os.Exit(utils.GetExitCode(appErr.CLIError.Code))
			return writer.WriteError("permission.remove", appErr.CLIError)
		}
		return writer.WriteError("permission.remove", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return writer.WriteSuccess("permission.remove", map[string]interface{}{
		"deleted":      true,
		"fileId":       fileID,
		"permissionId": permissionID,
	})
}

func runPermCreateLink(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	writer := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	// Validate role for public link
	validRoles := map[string]bool{"reader": true, "commenter": true, "writer": true}
	if !validRoles[permRole] {
		return writer.WriteError("permission.create-link", utils.NewCLIError(utils.ErrCodeInvalidArgument,
			"Invalid permission role for public link. Must be one of: reader, commenter, writer").Build())
	}

	mgr, err := getPermissionManager()
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return writer.WriteError("permission.create-link", appErr.CLIError)
		}
		return writer.WriteError("permission.create-link", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	reqCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypePermissionOp)
	fileID := args[0]

	result, err := mgr.CreatePublicLink(context.Background(), reqCtx, fileID, permRole, permAllowFileDiscovery)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			os.Exit(utils.GetExitCode(appErr.CLIError.Code))
			return writer.WriteError("permission.create-link", appErr.CLIError)
		}
		return writer.WriteError("permission.create-link", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return writer.WriteSuccess("permission.create-link", result)
}
