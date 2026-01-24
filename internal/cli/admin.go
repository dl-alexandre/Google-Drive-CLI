package cli

import (
	"context"
	"fmt"

	"github.com/dl-alexandre/gdrive/internal/admin"
	"github.com/dl-alexandre/gdrive/internal/api"
	"github.com/dl-alexandre/gdrive/internal/auth"
	"github.com/dl-alexandre/gdrive/internal/types"
	"github.com/dl-alexandre/gdrive/internal/utils"
	"github.com/spf13/cobra"
	adminapi "google.golang.org/api/admin/directory/v1"
)

var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Google Workspace Admin SDK operations",
	Long:  "Commands for managing users and groups",
}

var adminUsersCmd = &cobra.Command{
	Use:   "users",
	Short: "User management operations",
}

var adminGroupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "Group management operations",
}

var adminMembersCmd = &cobra.Command{
	Use:   "members",
	Short: "Group membership management",
}

var adminMembersListCmd = &cobra.Command{
	Use:   "list <group-key>",
	Short: "List group members",
	Long:  "List members of a group",
	Args:  cobra.ExactArgs(1),
	RunE:  runAdminMembersList,
}

var adminMembersAddCmd = &cobra.Command{
	Use:   "add <group-key> <member-email>",
	Short: "Add member to group",
	Args:  cobra.ExactArgs(2),
	RunE:  runAdminMembersAdd,
}

var adminMembersRemoveCmd = &cobra.Command{
	Use:   "remove <group-key> <member-email>",
	Short: "Remove member from group",
	Args:  cobra.ExactArgs(2),
	RunE:  runAdminMembersRemove,
}

var adminUsersListCmd = &cobra.Command{
	Use:   "list",
	Short: "List users",
	Long:  "List users in the domain",
	RunE:  runAdminUsersList,
}

var adminUsersGetCmd = &cobra.Command{
	Use:   "get <user-key>",
	Short: "Get user details",
	Long:  "Get details for a user by email or ID",
	Args:  cobra.ExactArgs(1),
	RunE:  runAdminUsersGet,
}

var adminUsersCreateCmd = &cobra.Command{
	Use:   "create <email>",
	Short: "Create a new user",
	Long:  "Create a new user in the domain",
	Args:  cobra.ExactArgs(1),
	RunE:  runAdminUsersCreate,
}

var adminUsersDeleteCmd = &cobra.Command{
	Use:   "delete <user-key>",
	Short: "Delete a user",
	Long:  "Delete a user by email or ID",
	Args:  cobra.ExactArgs(1),
	RunE:  runAdminUsersDelete,
}

var adminUsersUpdateCmd = &cobra.Command{
	Use:   "update <user-key>",
	Short: "Update a user",
	Long:  "Update user properties",
	Args:  cobra.ExactArgs(1),
	RunE:  runAdminUsersUpdate,
}

var adminUsersSuspendCmd = &cobra.Command{
	Use:   "suspend <user-key>",
	Short: "Suspend a user",
	Long:  "Suspend a user account",
	Args:  cobra.ExactArgs(1),
	RunE:  runAdminUsersSuspend,
}

var adminUsersUnsuspendCmd = &cobra.Command{
	Use:   "unsuspend <user-key>",
	Short: "Unsuspend a user",
	Long:  "Unsuspend a user account",
	Args:  cobra.ExactArgs(1),
	RunE:  runAdminUsersUnsuspend,
}
var adminGroupsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List groups",
	Long:  "List groups in the domain",
	RunE:  runAdminGroupsList,
}

var adminGroupsGetCmd = &cobra.Command{
	Use:   "get <group-key>",
	Short: "Get group details",
	Long:  "Get details of a group by email, alias, or ID",
	Args:  cobra.ExactArgs(1),
	RunE:  runAdminGroupsGet,
}

var adminGroupsCreateCmd = &cobra.Command{
	Use:   "create <email> <name>",
	Short: "Create a new group",
	Long:  "Create a new Google Workspace group",
	Args:  cobra.ExactArgs(2),
	RunE:  runAdminGroupsCreate,
}

var adminGroupsDeleteCmd = &cobra.Command{
	Use:   "delete <group-key>",
	Short: "Delete a group",
	Long:  "Delete a group by email, alias, or ID",
	Args:  cobra.ExactArgs(1),
	RunE:  runAdminGroupsDelete,
}

var adminGroupsUpdateCmd = &cobra.Command{
	Use:   "update <group-key>",
	Short: "Update a group",
	Long:  "Update group properties",
	Args:  cobra.ExactArgs(1),
	RunE:  runAdminGroupsUpdate,
}

var (
	adminUsersListDomain    string
	adminUsersListCustomer  string
	adminUsersListQuery     string
	adminUsersListLimit     int
	adminUsersListPageToken string
	adminUsersListFields    string
	adminUsersListPaginate  bool
	adminUsersListOrderBy   string
	adminUsersGetFields     string
	adminUsersCreateGiven   string
	adminUsersCreateFamily  string
	adminUsersCreatePass    string
	adminUsersUpdateGiven   string
	adminUsersUpdateFamily  string
	adminUsersUpdateSuspend string
	adminUsersUpdateOrgUnit string

	adminGroupsListDomain    string
	adminGroupsListCustomer  string
	adminGroupsListQuery     string
	adminGroupsListLimit     int
	adminGroupsListPageToken string
	adminGroupsListFields    string
	adminGroupsListPaginate  bool
	adminGroupsListOrderBy   string
	adminGroupsGetFields     string
	adminGroupsCreateDesc    string
	adminGroupsUpdateName    string
	adminGroupsUpdateDesc    string

	adminMembersListLimit     int
	adminMembersListPageToken string
	adminMembersListRoles     string
	adminMembersListFields    string
	adminMembersListPaginate  bool
	adminMembersAddRole       string
)

func init() {
	adminUsersListCmd.Flags().StringVar(&adminUsersListDomain, "domain", "", "Domain to list users from")
	adminUsersListCmd.Flags().StringVar(&adminUsersListCustomer, "customer", "", "Customer ID")
	adminUsersListCmd.Flags().StringVar(&adminUsersListQuery, "query", "", "Search query")
	adminUsersListCmd.Flags().IntVar(&adminUsersListLimit, "limit", 100, "Maximum results per page")
	adminUsersListCmd.Flags().StringVar(&adminUsersListPageToken, "page-token", "", "Page token for pagination")
	adminUsersListCmd.Flags().StringVar(&adminUsersListFields, "fields", "", "Fields to return")
	adminUsersListCmd.Flags().BoolVar(&adminUsersListPaginate, "paginate", false, "Automatically fetch all pages")
	adminUsersListCmd.Flags().StringVar(&adminUsersListOrderBy, "order-by", "", "Sort order")
	adminUsersGetCmd.Flags().StringVar(&adminUsersGetFields, "fields", "", "Fields to return")
	adminUsersCreateCmd.Flags().StringVar(&adminUsersCreateGiven, "given-name", "", "First name")
	adminUsersCreateCmd.Flags().StringVar(&adminUsersCreateFamily, "family-name", "", "Last name")
	adminUsersCreateCmd.Flags().StringVar(&adminUsersCreatePass, "password", "", "Password")
	_ = adminUsersCreateCmd.MarkFlagRequired("given-name")
	_ = adminUsersCreateCmd.MarkFlagRequired("family-name")
	_ = adminUsersCreateCmd.MarkFlagRequired("password")
	adminUsersUpdateCmd.Flags().StringVar(&adminUsersUpdateGiven, "given-name", "", "Update first name")
	adminUsersUpdateCmd.Flags().StringVar(&adminUsersUpdateFamily, "family-name", "", "Update last name")
	adminUsersUpdateCmd.Flags().StringVar(&adminUsersUpdateSuspend, "suspended", "", "Set suspension status (true/false)")
	adminUsersUpdateCmd.Flags().StringVar(&adminUsersUpdateOrgUnit, "org-unit-path", "", "Update organizational unit path")

	adminGroupsListCmd.Flags().StringVar(&adminGroupsListDomain, "domain", "", "Domain to list groups from")
	adminGroupsListCmd.Flags().StringVar(&adminGroupsListCustomer, "customer", "", "Customer ID")
	adminGroupsListCmd.Flags().StringVar(&adminGroupsListQuery, "query", "", "Search query")
	adminGroupsListCmd.Flags().IntVar(&adminGroupsListLimit, "limit", 100, "Maximum results per page")
	adminGroupsListCmd.Flags().StringVar(&adminGroupsListPageToken, "page-token", "", "Page token for pagination")
	adminGroupsListCmd.Flags().StringVar(&adminGroupsListFields, "fields", "", "Fields to return")
	adminGroupsListCmd.Flags().BoolVar(&adminGroupsListPaginate, "paginate", false, "Automatically fetch all pages")
	adminGroupsListCmd.Flags().StringVar(&adminGroupsListOrderBy, "order-by", "", "Sort order")
	adminGroupsGetCmd.Flags().StringVar(&adminGroupsGetFields, "fields", "", "Fields to return")
	adminGroupsCreateCmd.Flags().StringVar(&adminGroupsCreateDesc, "description", "", "Group description")
	adminGroupsUpdateCmd.Flags().StringVar(&adminGroupsUpdateName, "name", "", "Update group name")
	adminGroupsUpdateCmd.Flags().StringVar(&adminGroupsUpdateDesc, "description", "", "Update group description")

	adminMembersListCmd.Flags().IntVar(&adminMembersListLimit, "limit", 200, "Maximum results per page")
	adminMembersListCmd.Flags().StringVar(&adminMembersListPageToken, "page-token", "", "Page token for pagination")
	adminMembersListCmd.Flags().StringVar(&adminMembersListRoles, "roles", "", "Filter by role (OWNER, MANAGER, MEMBER)")
	adminMembersListCmd.Flags().StringVar(&adminMembersListFields, "fields", "", "Fields to return")
	adminMembersListCmd.Flags().BoolVar(&adminMembersListPaginate, "paginate", false, "Automatically fetch all pages")
	adminMembersAddCmd.Flags().StringVar(&adminMembersAddRole, "role", "MEMBER", "Member role: OWNER, MANAGER, or MEMBER")

	adminUsersCmd.AddCommand(adminUsersListCmd)
	adminUsersCmd.AddCommand(adminUsersGetCmd)
	adminUsersCmd.AddCommand(adminUsersCreateCmd)
	adminUsersCmd.AddCommand(adminUsersDeleteCmd)
	adminUsersCmd.AddCommand(adminUsersUpdateCmd)
	adminUsersCmd.AddCommand(adminUsersSuspendCmd)
	adminUsersCmd.AddCommand(adminUsersUnsuspendCmd)
	adminGroupsCmd.AddCommand(adminGroupsListCmd)
	adminGroupsCmd.AddCommand(adminGroupsGetCmd)
	adminGroupsCmd.AddCommand(adminGroupsCreateCmd)
	adminGroupsCmd.AddCommand(adminGroupsDeleteCmd)
	adminGroupsCmd.AddCommand(adminGroupsUpdateCmd)
	adminMembersCmd.AddCommand(adminMembersListCmd)
	adminMembersCmd.AddCommand(adminMembersAddCmd)
	adminMembersCmd.AddCommand(adminMembersRemoveCmd)
	adminGroupsCmd.AddCommand(adminMembersCmd)
	adminCmd.AddCommand(adminUsersCmd)
	adminCmd.AddCommand(adminGroupsCmd)
	rootCmd.AddCommand(adminCmd)
}

func runAdminUsersList(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getAdminService(ctx, flags)
	if err != nil {
		return out.WriteError("admin.users.list", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	if adminUsersListDomain == "" && adminUsersListCustomer == "" {
		return out.WriteError("admin.users.list", utils.NewCLIError(utils.ErrCodeInvalidArgument, "domain or customer is required").Build())
	}

	mgr := admin.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeListOrSearch
	result, err := mgr.ListUsers(ctx, reqCtx, &admin.ListUsersOptions{
		Domain:     adminUsersListDomain,
		Customer:   adminUsersListCustomer,
		Query:      adminUsersListQuery,
		MaxResults: int64(adminUsersListLimit),
		PageToken:  adminUsersListPageToken,
		OrderBy:    adminUsersListOrderBy,
		Fields:     adminUsersListFields,
		Paginate:   adminUsersListPaginate,
	})
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("admin.users.list", appErr.CLIError)
		}
		return out.WriteError("admin.users.list", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("admin.users.list", result)
}

func runAdminUsersGet(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getAdminService(ctx, flags)
	if err != nil {
		return out.WriteError("admin.users.get", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	userKey := args[0]
	mgr := admin.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeGetByID
	result, err := mgr.GetUser(ctx, reqCtx, userKey, adminUsersGetFields)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("admin.users.get", appErr.CLIError)
		}
		return out.WriteError("admin.users.get", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("admin.users.get", result)
}

func runAdminUsersCreate(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getAdminService(ctx, flags)
	if err != nil {
		return out.WriteError("admin.users.create", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	email := args[0]
	mgr := admin.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeMutation
	result, err := mgr.CreateUser(ctx, reqCtx, &types.CreateUserRequest{
		Email:      email,
		GivenName:  adminUsersCreateGiven,
		FamilyName: adminUsersCreateFamily,
		Password:   adminUsersCreatePass,
	})
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("admin.users.create", appErr.CLIError)
		}
		return out.WriteError("admin.users.create", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("admin.users.create", result)
}

func runAdminUsersDelete(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getAdminService(ctx, flags)
	if err != nil {
		return out.WriteError("admin.users.delete", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	userKey := args[0]
	mgr := admin.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeMutation
	if err := mgr.DeleteUser(ctx, reqCtx, userKey); err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("admin.users.delete", appErr.CLIError)
		}
		return out.WriteError("admin.users.delete", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("admin.users.delete", map[string]string{
		"message": fmt.Sprintf("User %s deleted successfully", userKey),
	})
}

func runAdminUsersUpdate(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getAdminService(ctx, flags)
	if err != nil {
		return out.WriteError("admin.users.update", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	req := &types.UpdateUserRequest{}
	if adminUsersUpdateGiven != "" {
		req.GivenName = &adminUsersUpdateGiven
	}
	if adminUsersUpdateFamily != "" {
		req.FamilyName = &adminUsersUpdateFamily
	}
	if adminUsersUpdateSuspend != "" {
		value, err := parseSuspendedFlag(adminUsersUpdateSuspend)
		if err != nil {
			return out.WriteError("admin.users.update", utils.NewCLIError(utils.ErrCodeInvalidArgument, err.Error()).Build())
		}
		req.Suspended = value
	}
	if adminUsersUpdateOrgUnit != "" {
		req.OrgUnitPath = &adminUsersUpdateOrgUnit
	}
	if req.GivenName == nil && req.FamilyName == nil && req.Suspended == nil && req.OrgUnitPath == nil {
		return out.WriteError("admin.users.update", utils.NewCLIError(utils.ErrCodeInvalidArgument, "at least one field must be provided").Build())
	}

	userKey := args[0]
	mgr := admin.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeMutation
	result, err := mgr.UpdateUser(ctx, reqCtx, userKey, req)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("admin.users.update", appErr.CLIError)
		}
		return out.WriteError("admin.users.update", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("admin.users.update", result)
}

func runAdminUsersSuspend(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getAdminService(ctx, flags)
	if err != nil {
		return out.WriteError("admin.users.suspend", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	userKey := args[0]
	mgr := admin.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeMutation
	result, err := mgr.SuspendUser(ctx, reqCtx, userKey)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("admin.users.suspend", appErr.CLIError)
		}
		return out.WriteError("admin.users.suspend", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("admin.users.suspend", result)
}

func runAdminUsersUnsuspend(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getAdminService(ctx, flags)
	if err != nil {
		return out.WriteError("admin.users.unsuspend", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	userKey := args[0]
	mgr := admin.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeMutation
	result, err := mgr.UnsuspendUser(ctx, reqCtx, userKey)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("admin.users.unsuspend", appErr.CLIError)
		}
		return out.WriteError("admin.users.unsuspend", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("admin.users.unsuspend", result)
}

func runAdminGroupsList(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getAdminService(ctx, flags)
	if err != nil {
		return out.WriteError("admin.groups.list", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	if adminGroupsListDomain == "" && adminGroupsListCustomer == "" {
		return out.WriteError("admin.groups.list", utils.NewCLIError(utils.ErrCodeInvalidArgument, "domain or customer is required").Build())
	}

	mgr := admin.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeListOrSearch
	result, err := mgr.ListGroups(ctx, reqCtx, &admin.ListGroupsOptions{
		Domain:     adminGroupsListDomain,
		Customer:   adminGroupsListCustomer,
		Query:      adminGroupsListQuery,
		MaxResults: int64(adminGroupsListLimit),
		PageToken:  adminGroupsListPageToken,
		OrderBy:    adminGroupsListOrderBy,
		Fields:     adminGroupsListFields,
		Paginate:   adminGroupsListPaginate,
	})
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("admin.groups.list", appErr.CLIError)
		}
		return out.WriteError("admin.groups.list", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("admin.groups.list", result)
}

func runAdminGroupsGet(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getAdminService(ctx, flags)
	if err != nil {
		return out.WriteError("admin.groups.get", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	groupKey := args[0]
	mgr := admin.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeGetByID
	result, err := mgr.GetGroup(ctx, reqCtx, groupKey, &admin.GetGroupOptions{Fields: adminGroupsGetFields})
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("admin.groups.get", appErr.CLIError)
		}
		return out.WriteError("admin.groups.get", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("admin.groups.get", result)
}

func runAdminGroupsCreate(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getAdminService(ctx, flags)
	if err != nil {
		return out.WriteError("admin.groups.create", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	email := args[0]
	name := args[1]
	mgr := admin.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeMutation
	result, err := mgr.CreateGroup(ctx, reqCtx, &types.CreateGroupRequest{
		Email:       email,
		Name:        name,
		Description: adminGroupsCreateDesc,
	})
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("admin.groups.create", appErr.CLIError)
		}
		return out.WriteError("admin.groups.create", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("admin.groups.create", result)
}

func runAdminGroupsDelete(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getAdminService(ctx, flags)
	if err != nil {
		return out.WriteError("admin.groups.delete", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	groupKey := args[0]
	mgr := admin.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeMutation
	if err := mgr.DeleteGroup(ctx, reqCtx, groupKey); err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("admin.groups.delete", appErr.CLIError)
		}
		return out.WriteError("admin.groups.delete", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("admin.groups.delete", map[string]string{
		"message": fmt.Sprintf("Group %s deleted successfully", groupKey),
	})
}

func runAdminGroupsUpdate(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getAdminService(ctx, flags)
	if err != nil {
		return out.WriteError("admin.groups.update", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	req := &types.UpdateGroupRequest{}
	if adminGroupsUpdateName != "" {
		req.Name = &adminGroupsUpdateName
	}
	if adminGroupsUpdateDesc != "" {
		req.Description = &adminGroupsUpdateDesc
	}
	if req.Name == nil && req.Description == nil {
		return out.WriteError("admin.groups.update", utils.NewCLIError(utils.ErrCodeInvalidArgument, "at least one field must be provided").Build())
	}

	groupKey := args[0]
	mgr := admin.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeMutation
	result, err := mgr.UpdateGroup(ctx, reqCtx, groupKey, req)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("admin.groups.update", appErr.CLIError)
		}
		return out.WriteError("admin.groups.update", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("admin.groups.update", result)
}

func parseSuspendedFlag(value string) (*bool, error) {
	switch value {
	case "true":
		result := true
		return &result, nil
	case "false":
		result := false
		return &result, nil
	default:
		return nil, fmt.Errorf("suspended must be true or false")
	}
}

func runAdminMembersList(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getAdminService(ctx, flags)
	if err != nil {
		return out.WriteError("admin.groups.members.list", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	groupKey := args[0]
	mgr := admin.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeListOrSearch
	result, err := mgr.ListMembers(ctx, reqCtx, groupKey, &admin.ListMembersOptions{
		MaxResults: int64(adminMembersListLimit),
		PageToken:  adminMembersListPageToken,
		Roles:      adminMembersListRoles,
		Fields:     adminMembersListFields,
		Paginate:   adminMembersListPaginate,
	})
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("admin.groups.members.list", appErr.CLIError)
		}
		return out.WriteError("admin.groups.members.list", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("admin.groups.members.list", result)
}

func runAdminMembersAdd(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getAdminService(ctx, flags)
	if err != nil {
		return out.WriteError("admin.groups.members.add", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	role := adminMembersAddRole
	if role != "OWNER" && role != "MANAGER" && role != "MEMBER" {
		return out.WriteError("admin.groups.members.add", utils.NewCLIError(utils.ErrCodeInvalidArgument, "role must be OWNER, MANAGER, or MEMBER").Build())
	}

	groupKey := args[0]
	memberEmail := args[1]

	mgr := admin.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeMutation
	result, err := mgr.AddMember(ctx, reqCtx, groupKey, &types.AddMemberRequest{
		Email: memberEmail,
		Role:  role,
	})
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("admin.groups.members.add", appErr.CLIError)
		}
		return out.WriteError("admin.groups.members.add", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("admin.groups.members.add", result)
}

func runAdminMembersRemove(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getAdminService(ctx, flags)
	if err != nil {
		return out.WriteError("admin.groups.members.remove", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	groupKey := args[0]
	memberKey := args[1]

	mgr := admin.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeMutation
	if err := mgr.RemoveMember(ctx, reqCtx, groupKey, memberKey); err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("admin.groups.members.remove", appErr.CLIError)
		}
		return out.WriteError("admin.groups.members.remove", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("admin.groups.members.remove", map[string]string{
		"message": fmt.Sprintf("Member %s removed from group %s", memberKey, groupKey),
	})
}

func getAdminService(ctx context.Context, flags types.GlobalFlags) (*adminapi.Service, *api.Client, *types.RequestContext, error) {
	configDir := getConfigDir()
	authMgr := auth.NewManager(configDir)

	creds, err := authMgr.GetValidCredentials(ctx, flags.Profile)
	if err != nil {
		return nil, nil, nil, err
	}
	if creds.Type != types.AuthTypeServiceAccount && creds.Type != types.AuthTypeImpersonated {
		return nil, nil, nil, fmt.Errorf("admin operations require service account authentication")
	}

	if err := authMgr.ValidateServiceScopes(creds, auth.ServiceAdminDir); err != nil {
		return nil, nil, nil, err
	}

	svc, err := authMgr.GetAdminService(ctx, creds)
	if err != nil {
		return nil, nil, nil, err
	}

	driveSvc, err := authMgr.GetDriveService(ctx, creds)
	if err != nil {
		return nil, nil, nil, err
	}

	client := api.NewClient(driveSvc, utils.DefaultMaxRetries, utils.DefaultRetryDelayMs, GetLogger())
	reqCtx := api.NewRequestContext(flags.Profile, "", types.RequestTypeListOrSearch)
	return svc, client, reqCtx, nil
}
