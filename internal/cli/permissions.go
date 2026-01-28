package cli

import (
	"context"
	"os"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/auth"
	"github.com/dl-alexandre/gdrv/internal/permissions"
	"github.com/dl-alexandre/gdrv/internal/types"
	"github.com/dl-alexandre/gdrv/internal/utils"
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

var permAuditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Audit permissions",
	Long:  "Audit file and folder permissions for security and compliance",
}

var permAuditPublicCmd = &cobra.Command{
	Use:   "public",
	Short: "Audit public files",
	Long:  "Find all files with public access (anyone)",
	RunE:  runPermAuditPublic,
}

var permAuditExternalCmd = &cobra.Command{
	Use:   "external",
	Short: "Audit external shares",
	Long:  "Find all files shared with external domains",
	RunE:  runPermAuditExternal,
}

var permAuditAnyoneWithLinkCmd = &cobra.Command{
	Use:   "anyone-with-link",
	Short: "Audit anyone-with-link files",
	Long:  "Find all files with 'anyone with link' access",
	RunE:  runPermAuditAnyoneWithLink,
}

var permAuditUserCmd = &cobra.Command{
	Use:   "user <email>",
	Short: "Audit user access",
	Long:  "Find all files accessible by a specific user",
	Args:  cobra.ExactArgs(1),
	RunE:  runPermAuditUser,
}

var permAnalyzeCmd = &cobra.Command{
	Use:   "analyze <folder-id>",
	Short: "Analyze folder permissions",
	Long:  "Analyze permissions for a folder and its contents",
	Args:  cobra.ExactArgs(1),
	RunE:  runPermAnalyze,
}

var permReportCmd = &cobra.Command{
	Use:   "report <file-id>",
	Short: "Generate permission report",
	Long:  "Generate a detailed permission report for a file or folder",
	Args:  cobra.ExactArgs(1),
	RunE:  runPermReport,
}

var permBulkCmd = &cobra.Command{
	Use:   "bulk",
	Short: "Bulk permission operations",
	Long:  "Perform bulk permission operations on multiple files",
}

var permBulkRemovePublicCmd = &cobra.Command{
	Use:   "remove-public",
	Short: "Bulk remove public access",
	Long:  "Remove public access from all files in a folder",
	RunE:  runPermBulkRemovePublic,
}

var permBulkUpdateRoleCmd = &cobra.Command{
	Use:   "update-role",
	Short: "Bulk update roles",
	Long:  "Update permission roles from one role to another in a folder",
	RunE:  runPermBulkUpdateRole,
}

var permSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search permissions",
	Long:  "Search for files by permission criteria",
	RunE:  runPermSearch,
}

var (
	auditFolderID       string
	auditRecursive      bool
	auditInternalDomain string
	auditIncludePerms   bool

	analyzeRecursive      bool
	analyzeMaxDepth       int
	analyzeIncludeDetails bool
	analyzeInternalDomain string

	bulkFolderID        string
	bulkRecursive       bool
	bulkFromRole        string
	bulkToRole          string
	bulkMaxFiles        int
	bulkContinueOnError bool

	searchEmail     string
	searchRole      string
	searchFolderID  string
	searchRecursive bool
)

func init() {
	rootCmd.AddCommand(permissionsCmd)

	permissionsCmd.AddCommand(permListCmd)
	permissionsCmd.AddCommand(permCreateCmd)
	permissionsCmd.AddCommand(permUpdateCmd)
	permissionsCmd.AddCommand(permRemoveCmd)
	permissionsCmd.AddCommand(permCreateLinkCmd)
	permissionsCmd.AddCommand(permAuditCmd)
	permissionsCmd.AddCommand(permAnalyzeCmd)
	permissionsCmd.AddCommand(permReportCmd)
	permissionsCmd.AddCommand(permBulkCmd)
	permissionsCmd.AddCommand(permSearchCmd)

	permAuditCmd.AddCommand(permAuditPublicCmd)
	permAuditCmd.AddCommand(permAuditExternalCmd)
	permAuditCmd.AddCommand(permAuditAnyoneWithLinkCmd)
	permAuditCmd.AddCommand(permAuditUserCmd)

	permBulkCmd.AddCommand(permBulkRemovePublicCmd)
	permBulkCmd.AddCommand(permBulkUpdateRoleCmd)

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

	// Audit flags
	permAuditPublicCmd.Flags().StringVar(&auditFolderID, "folder-id", "", "Limit audit to specific folder")
	permAuditPublicCmd.Flags().BoolVar(&auditRecursive, "recursive", false, "Include subfolders")
	permAuditPublicCmd.Flags().BoolVar(&auditIncludePerms, "include-permissions", false, "Include full permission details")

	permAuditExternalCmd.Flags().StringVar(&auditFolderID, "folder-id", "", "Limit audit to specific folder")
	permAuditExternalCmd.Flags().BoolVar(&auditRecursive, "recursive", false, "Include subfolders")
	permAuditExternalCmd.Flags().StringVar(&auditInternalDomain, "internal-domain", "", "Internal domain (required)")
	permAuditExternalCmd.Flags().BoolVar(&auditIncludePerms, "include-permissions", false, "Include full permission details")
	_ = permAuditExternalCmd.MarkFlagRequired("internal-domain")

	permAuditAnyoneWithLinkCmd.Flags().StringVar(&auditFolderID, "folder-id", "", "Limit audit to specific folder")
	permAuditAnyoneWithLinkCmd.Flags().BoolVar(&auditRecursive, "recursive", false, "Include subfolders")
	permAuditAnyoneWithLinkCmd.Flags().BoolVar(&auditIncludePerms, "include-permissions", false, "Include full permission details")

	permAuditUserCmd.Flags().StringVar(&auditFolderID, "folder-id", "", "Limit audit to specific folder")
	permAuditUserCmd.Flags().BoolVar(&auditRecursive, "recursive", false, "Include subfolders")
	permAuditUserCmd.Flags().BoolVar(&auditIncludePerms, "include-permissions", false, "Include full permission details")

	// Analyze flags
	permAnalyzeCmd.Flags().BoolVar(&analyzeRecursive, "recursive", false, "Analyze subfolders recursively")
	permAnalyzeCmd.Flags().IntVar(&analyzeMaxDepth, "max-depth", 0, "Maximum recursion depth (0 = unlimited)")
	permAnalyzeCmd.Flags().BoolVar(&analyzeIncludeDetails, "include-details", false, "Include detailed file lists")
	permAnalyzeCmd.Flags().StringVar(&analyzeInternalDomain, "internal-domain", "", "Internal domain for external detection")

	// Report flags
	permReportCmd.Flags().StringVar(&analyzeInternalDomain, "internal-domain", "", "Internal domain for external detection")

	// Bulk remove public flags
	permBulkRemovePublicCmd.Flags().StringVar(&bulkFolderID, "folder-id", "", "Folder to operate on (required)")
	permBulkRemovePublicCmd.Flags().BoolVar(&bulkRecursive, "recursive", false, "Include subfolders")
	permBulkRemovePublicCmd.Flags().IntVar(&bulkMaxFiles, "max-files", 0, "Maximum files to process (0 = unlimited)")
	permBulkRemovePublicCmd.Flags().BoolVar(&bulkContinueOnError, "continue-on-error", false, "Continue if individual operations fail")
	_ = permBulkRemovePublicCmd.MarkFlagRequired("folder-id")

	// Bulk update role flags
	permBulkUpdateRoleCmd.Flags().StringVar(&bulkFolderID, "folder-id", "", "Folder to operate on (required)")
	permBulkUpdateRoleCmd.Flags().BoolVar(&bulkRecursive, "recursive", false, "Include subfolders")
	permBulkUpdateRoleCmd.Flags().StringVar(&bulkFromRole, "from-role", "", "Source role (required)")
	permBulkUpdateRoleCmd.Flags().StringVar(&bulkToRole, "to-role", "", "Target role (required)")
	permBulkUpdateRoleCmd.Flags().IntVar(&bulkMaxFiles, "max-files", 0, "Maximum files to process (0 = unlimited)")
	permBulkUpdateRoleCmd.Flags().BoolVar(&bulkContinueOnError, "continue-on-error", false, "Continue if individual operations fail")
	_ = permBulkUpdateRoleCmd.MarkFlagRequired("folder-id")
	_ = permBulkUpdateRoleCmd.MarkFlagRequired("from-role")
	_ = permBulkUpdateRoleCmd.MarkFlagRequired("to-role")

	// Search flags
	permSearchCmd.Flags().StringVar(&searchEmail, "email", "", "Search by email address")
	permSearchCmd.Flags().StringVar(&searchRole, "role", "", "Search by role")
	permSearchCmd.Flags().StringVar(&searchFolderID, "folder-id", "", "Limit search to specific folder")
	permSearchCmd.Flags().BoolVar(&searchRecursive, "recursive", false, "Include subfolders")
}

func getPermissionManager() (*permissions.Manager, error) {
	flags := GetGlobalFlags()

	configDir := getConfigDir()
	authMgr := auth.NewManager(configDir)
	creds, err := authMgr.LoadCredentials(flags.Profile)
	if err != nil {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeAuthRequired,
			"Authentication required. Run 'gdrv auth login' first.").Build())
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

func runPermAuditPublic(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	writer := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	mgr, err := getPermissionManager()
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return writer.WriteError("permissions.audit.public", appErr.CLIError)
		}
		return writer.WriteError("permissions.audit.public", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	reqCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypePermissionOp)
	opts := types.AuditOptions{
		FolderID:           auditFolderID,
		Recursive:          auditRecursive,
		IncludePermissions: auditIncludePerms,
	}

	result, err := mgr.AuditPublic(context.Background(), reqCtx, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			os.Exit(utils.GetExitCode(appErr.CLIError.Code))
			return writer.WriteError("permissions.audit.public", appErr.CLIError)
		}
		return writer.WriteError("permissions.audit.public", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return writer.WriteSuccess("permissions.audit.public", result)
}

func runPermAuditExternal(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	writer := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	mgr, err := getPermissionManager()
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return writer.WriteError("permissions.audit.external", appErr.CLIError)
		}
		return writer.WriteError("permissions.audit.external", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	reqCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypePermissionOp)
	opts := types.AuditOptions{
		FolderID:           auditFolderID,
		Recursive:          auditRecursive,
		InternalDomain:     auditInternalDomain,
		IncludePermissions: auditIncludePerms,
	}

	result, err := mgr.AuditExternal(context.Background(), reqCtx, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			os.Exit(utils.GetExitCode(appErr.CLIError.Code))
			return writer.WriteError("permissions.audit.external", appErr.CLIError)
		}
		return writer.WriteError("permissions.audit.external", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return writer.WriteSuccess("permissions.audit.external", result)
}

func runPermAuditAnyoneWithLink(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	writer := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	mgr, err := getPermissionManager()
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return writer.WriteError("permissions.audit.anyone-with-link", appErr.CLIError)
		}
		return writer.WriteError("permissions.audit.anyone-with-link", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	reqCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypePermissionOp)
	opts := types.AuditOptions{
		FolderID:           auditFolderID,
		Recursive:          auditRecursive,
		IncludePermissions: auditIncludePerms,
	}

	result, err := mgr.AuditAnyoneWithLink(context.Background(), reqCtx, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			os.Exit(utils.GetExitCode(appErr.CLIError.Code))
			return writer.WriteError("permissions.audit.anyone-with-link", appErr.CLIError)
		}
		return writer.WriteError("permissions.audit.anyone-with-link", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return writer.WriteSuccess("permissions.audit.anyone-with-link", result)
}

func runPermAuditUser(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	writer := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	mgr, err := getPermissionManager()
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return writer.WriteError("permissions.audit.user", appErr.CLIError)
		}
		return writer.WriteError("permissions.audit.user", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	reqCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypePermissionOp)
	email := args[0]
	opts := types.AuditOptions{
		FolderID:           auditFolderID,
		Recursive:          auditRecursive,
		IncludePermissions: auditIncludePerms,
	}

	result, err := mgr.AuditUser(context.Background(), reqCtx, email, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			os.Exit(utils.GetExitCode(appErr.CLIError.Code))
			return writer.WriteError("permissions.audit.user", appErr.CLIError)
		}
		return writer.WriteError("permissions.audit.user", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return writer.WriteSuccess("permissions.audit.user", result)
}

func runPermAnalyze(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	writer := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	mgr, err := getPermissionManager()
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return writer.WriteError("permissions.analyze", appErr.CLIError)
		}
		return writer.WriteError("permissions.analyze", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	reqCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypePermissionOp)
	folderID := args[0]
	opts := types.AnalyzeOptions{
		Recursive:      analyzeRecursive,
		MaxDepth:       analyzeMaxDepth,
		IncludeDetails: analyzeIncludeDetails,
		InternalDomain: analyzeInternalDomain,
	}

	result, err := mgr.AnalyzeFolder(context.Background(), reqCtx, folderID, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			os.Exit(utils.GetExitCode(appErr.CLIError.Code))
			return writer.WriteError("permissions.analyze", appErr.CLIError)
		}
		return writer.WriteError("permissions.analyze", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return writer.WriteSuccess("permissions.analyze", result)
}

func runPermReport(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	writer := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	mgr, err := getPermissionManager()
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return writer.WriteError("permissions.report", appErr.CLIError)
		}
		return writer.WriteError("permissions.report", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	reqCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypePermissionOp)
	fileID := args[0]

	result, err := mgr.GenerateReport(context.Background(), reqCtx, fileID, analyzeInternalDomain)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			os.Exit(utils.GetExitCode(appErr.CLIError.Code))
			return writer.WriteError("permissions.report", appErr.CLIError)
		}
		return writer.WriteError("permissions.report", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return writer.WriteSuccess("permissions.report", result)
}

func runPermBulkRemovePublic(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	writer := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	mgr, err := getPermissionManager()
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return writer.WriteError("permissions.bulk.remove-public", appErr.CLIError)
		}
		return writer.WriteError("permissions.bulk.remove-public", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	reqCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypePermissionOp)
	opts := types.BulkOptions{
		FolderID:        bulkFolderID,
		Recursive:       bulkRecursive,
		DryRun:          flags.DryRun,
		MaxFiles:        bulkMaxFiles,
		ContinueOnError: bulkContinueOnError,
	}

	result, err := mgr.BulkRemovePublic(context.Background(), reqCtx, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			os.Exit(utils.GetExitCode(appErr.CLIError.Code))
			return writer.WriteError("permissions.bulk.remove-public", appErr.CLIError)
		}
		return writer.WriteError("permissions.bulk.remove-public", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return writer.WriteSuccess("permissions.bulk.remove-public", result)
}

func runPermBulkUpdateRole(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	writer := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	mgr, err := getPermissionManager()
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return writer.WriteError("permissions.bulk.update-role", appErr.CLIError)
		}
		return writer.WriteError("permissions.bulk.update-role", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	reqCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypePermissionOp)
	opts := types.BulkOptions{
		FolderID:        bulkFolderID,
		Recursive:       bulkRecursive,
		DryRun:          flags.DryRun,
		MaxFiles:        bulkMaxFiles,
		ContinueOnError: bulkContinueOnError,
	}

	result, err := mgr.BulkUpdateRole(context.Background(), reqCtx, bulkFromRole, bulkToRole, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			os.Exit(utils.GetExitCode(appErr.CLIError.Code))
			return writer.WriteError("permissions.bulk.update-role", appErr.CLIError)
		}
		return writer.WriteError("permissions.bulk.update-role", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return writer.WriteSuccess("permissions.bulk.update-role", result)
}

func runPermSearch(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	writer := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	if searchEmail == "" && searchRole == "" {
		return writer.WriteError("permissions.search", utils.NewCLIError(utils.ErrCodeInvalidArgument,
			"Either --email or --role must be specified").Build())
	}

	mgr, err := getPermissionManager()
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return writer.WriteError("permissions.search", appErr.CLIError)
		}
		return writer.WriteError("permissions.search", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	reqCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypePermissionOp)
	searchOpts := types.SearchOptions{
		Email:     searchEmail,
		Role:      searchRole,
		FolderID:  searchFolderID,
		Recursive: searchRecursive,
	}

	var result *types.AuditResult
	if searchEmail != "" {
		result, err = mgr.SearchByEmail(context.Background(), reqCtx, searchOpts)
	} else {
		result, err = mgr.SearchByRole(context.Background(), reqCtx, searchOpts)
	}

	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			os.Exit(utils.GetExitCode(appErr.CLIError.Code))
			return writer.WriteError("permissions.search", appErr.CLIError)
		}
		return writer.WriteError("permissions.search", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return writer.WriteSuccess("permissions.search", result)
}
