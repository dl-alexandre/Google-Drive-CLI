---
name: Admin SDK Directory CLI Integration
overview: Add Admin SDK Directory API support for user and group management with service account authentication and domain-wide delegation, positioning the CLI as an enterprise IT admin tool.
todos: []
isProject: false
status: completed
dependencies:
  - 00_foundation_google_apis.plan.md
---

> **Status**: ✅ **COMPLETED** - Admin SDK Directory CLI integration has been successfully implemented.

## Context

**CRITICAL DIFFERENCE**: Admin SDK requires service account auth with domain-wide delegation. OAuth user flow is NOT supported for admin operations.

Foundation provides: scope management, service factory, error handling, but Admin SDK needs special auth handling.

Current service account support exists:
```go
// internal/auth/service_account.go:29-72
func (m *Manager) LoadServiceAccount(ctx context.Context, keyFilePath string, scopes []string, impersonateUser string) (*types.Credentials, error)
```

Use cases:
- Automated user provisioning/deprovisioning
- Group membership management
- Bulk user updates
- Compliance reporting
- Integration with HR systems

## Plan

### 1. Admin Auth Extensions

**Problem**: Admin SDK requires domain-wide delegation and admin impersonation. Need to validate auth mode.

**Files to modify**:
- `internal/auth/manager.go`: Add admin validation
- `internal/cli/auth.go`: Add service account login with impersonation

**Implementation**:
```go
// internal/auth/manager.go
func (m *Manager) ValidateAdminAuth(creds *types.Credentials) error {
    if creds.Type != "service_account" {
        return fmt.Errorf("Admin SDK requires service account authentication with domain-wide delegation")
    }

    if creds.ImpersonateUser == "" {
        return fmt.Errorf("Admin SDK requires --impersonate-user flag to specify admin user")
    }

    return nil
}

// internal/cli/auth.go - add service account command
var authServiceAccountCmd = &cobra.Command{
    Use:   "service-account <key-file>",
    Short: "Authenticate with service account",
    Long:  "Load service account credentials (requires domain-wide delegation for Admin SDK)",
    Args:  cobra.ExactArgs(1),
    RunE:  runAuthServiceAccount,
}

var (
    saImpersonateUser string
    saScopes          []string
)

func init() {
    authCmd.AddCommand(authServiceAccountCmd)
    authServiceAccountCmd.Flags().StringVar(&saImpersonateUser, "impersonate-user", "", "User email to impersonate (required for Admin SDK)")
    authServiceAccountCmd.Flags().StringSliceVar(&saScopes, "scopes", utils.ScopesWorkspaceBasic, "OAuth scopes")
}

func runAuthServiceAccount(cmd *cobra.Command, args []string) error {
    ctx := context.Background()
    keyFile := args[0]

    mgr, err := getAuthManager()
    if err != nil {
        return err
    }

    creds, err := mgr.LoadServiceAccount(ctx, keyFile, saScopes, saImpersonateUser)
    if err != nil {
        return fmt.Errorf("load service account: %w", err)
    }

    // Store credentials
    if err := mgr.SaveCredentials(getProfile(), creds); err != nil {
        return fmt.Errorf("save credentials: %w", err)
    }

    out.Success("Service account authenticated successfully")
    if saImpersonateUser != "" {
        out.Log("Impersonating: %s", saImpersonateUser)
    }

    return nil
}
```

### 2. Admin Manager Layer

**Files to create**:
- `internal/admin/manager.go`: Admin Directory operations
- `internal/admin/manager_test.go`: Unit tests
- `internal/types/admin.go`: Admin-specific types

**Key operations**:
```go
package admin

import (
    "context"
    "google.golang.org/api/admin/directory/v1"
    "github.com/dl-alexandre/Google-Drive-CLI/internal/types"
    "github.com/dl-alexandre/Google-Drive-CLI/internal/errors"
    "github.com/dl-alexandre/Google-Drive-CLI/internal/retry"
)

type Manager struct {
    service     *admin.Service
    retryConfig *retry.RetryConfig
}

func NewManager(service *admin.Service) *Manager {
    return &Manager{
        service:     service,
        retryConfig: retry.DefaultRetryConfig(),
    }
}

// User operations
func (m *Manager) ListUsers(ctx context.Context, domain string, opts *types.ListUsersOptions) (*types.UserList, error)
func (m *Manager) GetUser(ctx context.Context, userKey string) (*types.User, error)
func (m *Manager) CreateUser(ctx context.Context, user *types.CreateUserRequest) (*types.User, error)
func (m *Manager) UpdateUser(ctx context.Context, userKey string, update *types.UpdateUserRequest) (*types.User, error)
func (m *Manager) DeleteUser(ctx context.Context, userKey string) error
func (m *Manager) SuspendUser(ctx context.Context, userKey string) (*types.User, error)
func (m *Manager) UnsuspendUser(ctx context.Context, userKey string) (*types.User, error)

// Group operations
func (m *Manager) ListGroups(ctx context.Context, domain string, opts *types.ListGroupsOptions) (*types.GroupList, error)
func (m *Manager) GetGroup(ctx context.Context, groupKey string) (*types.Group, error)
func (m *Manager) CreateGroup(ctx context.Context, group *types.CreateGroupRequest) (*types.Group, error)
func (m *Manager) UpdateGroup(ctx context.Context, groupKey string, update *types.UpdateGroupRequest) (*types.Group, error)
func (m *Manager) DeleteGroup(ctx context.Context, groupKey string) error

// Group membership operations
func (m *Manager) ListMembers(ctx context.Context, groupKey string) (*types.MemberList, error)
func (m *Manager) GetMember(ctx context.Context, groupKey, memberKey string) (*types.Member, error)
func (m *Manager) AddMember(ctx context.Context, groupKey string, member *types.AddMemberRequest) (*types.Member, error)
func (m *Manager) RemoveMember(ctx context.Context, groupKey, memberKey string) error
func (m *Manager) UpdateMember(ctx context.Context, groupKey, memberKey string, role string) (*types.Member, error)
```

**Example implementation**:
```go
func (m *Manager) ListUsers(ctx context.Context, domain string, opts *types.ListUsersOptions) (*types.UserList, error) {
    call := m.service.Users.List().
        Customer("my_customer").
        Domain(domain).
        MaxResults(int64(opts.MaxResults))

    if opts.Query != "" {
        call = call.Query(opts.Query)
    }
    if opts.OrderBy != "" {
        call = call.OrderBy(opts.OrderBy)
    }
    if opts.PageToken != "" {
        call = call.PageToken(opts.PageToken)
    }

    var resp *admin.Users
    err := retry.WithRetry(ctx, m.retryConfig, func() error {
        var err error
        resp, err = call.Context(ctx).Do()
        return err
    })

    if err != nil {
        return nil, errors.ParseGoogleAPIError(err, "Admin Directory")
    }

    users := make([]*types.User, len(resp.Users))
    for i, u := range resp.Users {
        users[i] = types.NewUserFromAPI(u)
    }

    return &types.UserList{
        Users:         users,
        NextPageToken: resp.NextPageToken,
        TotalResults:  len(users),
    }, nil
}

func (m *Manager) CreateUser(ctx context.Context, req *types.CreateUserRequest) (*types.User, error) {
    user := &admin.User{
        PrimaryEmail: req.Email,
        Name: &admin.UserName{
            GivenName:  req.GivenName,
            FamilyName: req.FamilyName,
        },
        Password: req.Password,
    }

    var resp *admin.User
    err := retry.WithRetry(ctx, m.retryConfig, func() error {
        var err error
        resp, err = m.service.Users.Insert(user).Context(ctx).Do()
        return err
    })

    if err != nil {
        return nil, errors.ParseGoogleAPIError(err, "Admin Directory")
    }

    return types.NewUserFromAPI(resp), nil
}
```

### 3. Admin Types

**Files to create**:
- `internal/types/admin.go`: User, Group, Member types

**Key types**:
```go
type User struct {
    ID                string
    PrimaryEmail      string
    FullName          string
    GivenName         string
    FamilyName        string
    IsAdmin           bool
    IsSuspended       bool
    CreationTime      string
    LastLoginTime     string
    OrgUnitPath       string
    Aliases           []string
}

type UserList struct {
    Users         []*User
    NextPageToken string
    TotalResults  int
}

// Implement TableRenderer
func (ul *UserList) Headers() []string {
    return []string{"Email", "Name", "Admin", "Suspended", "Last Login", "Org Unit"}
}

func (ul *UserList) Rows() [][]string {
    rows := make([][]string, len(ul.Users))
    for i, u := range ul.Users {
        rows[i] = []string{
            u.PrimaryEmail,
            u.FullName,
            formatBool(u.IsAdmin),
            formatBool(u.IsSuspended),
            formatTime(u.LastLoginTime),
            u.OrgUnitPath,
        }
    }
    return rows
}

type Group struct {
    ID                string
    Email             string
    Name              string
    Description       string
    DirectMembersCount int64
    AdminCreated      bool
}

type GroupList struct {
    Groups        []*Group
    NextPageToken string
    TotalResults  int
}

// Implement TableRenderer
func (gl *GroupList) Headers() []string {
    return []string{"Email", "Name", "Members", "Admin Created"}
}

func (gl *GroupList) Rows() [][]string {
    rows := make([][]string, len(gl.Groups))
    for i, g := range gl.Groups {
        rows[i] = []string{
            g.Email,
            g.Name,
            fmt.Sprintf("%d", g.DirectMembersCount),
            formatBool(g.AdminCreated),
        }
    }
    return rows
}

type Member struct {
    ID     string
    Email  string
    Role   string // OWNER, MANAGER, MEMBER
    Type   string // USER, GROUP
    Status string
}

type MemberList struct {
    Members       []*Member
    NextPageToken string
    TotalResults  int
}

type CreateUserRequest struct {
    Email      string
    GivenName  string
    FamilyName string
    Password   string
}

type UpdateUserRequest struct {
    GivenName   *string
    FamilyName  *string
    Suspended   *bool
    OrgUnitPath *string
}

type ListUsersOptions struct {
    MaxResults int
    PageToken  string
    Query      string
    OrderBy    string
}

// Similar for groups
type CreateGroupRequest struct {
    Email       string
    Name        string
    Description string
}

type AddMemberRequest struct {
    Email string
    Role  string
}
```

### 4. CLI Commands

**Files to create**:
- `internal/cli/admin.go`: Admin command implementation

**Command structure**:
```go
var adminCmd = &cobra.Command{
    Use:   "admin",
    Short: "Admin SDK operations",
    Long:  "Manage users and groups (requires service account with domain-wide delegation)",
}

var adminUsersCmd = &cobra.Command{
    Use:   "users",
    Short: "User management",
}

var adminGroupsCmd = &cobra.Command{
    Use:   "groups",
    Short: "Group management",
}

var adminMembersCmd = &cobra.Command{
    Use:   "members",
    Short: "Group membership management",
}

func init() {
    rootCmd.AddCommand(adminCmd)
    adminCmd.AddCommand(adminUsersCmd)
    adminCmd.AddCommand(adminGroupsCmd)
    adminCmd.AddCommand(adminMembersCmd)

    // User commands
    adminUsersCmd.AddCommand(adminUsersListCmd)
    adminUsersCmd.AddCommand(adminUsersGetCmd)
    adminUsersCmd.AddCommand(adminUsersCreateCmd)
    adminUsersCmd.AddCommand(adminUsersUpdateCmd)
    adminUsersCmd.AddCommand(adminUsersDeleteCmd)
    adminUsersCmd.AddCommand(adminUsersSuspendCmd)

    // Group commands
    adminGroupsCmd.AddCommand(adminGroupsListCmd)
    adminGroupsCmd.AddCommand(adminGroupsGetCmd)
    adminGroupsCmd.AddCommand(adminGroupsCreateCmd)
    adminGroupsCmd.AddCommand(adminGroupsDeleteCmd)

    // Member commands
    adminMembersCmd.AddCommand(adminMembersListCmd)
    adminMembersCmd.AddCommand(adminMembersAddCmd)
    adminMembersCmd.AddCommand(adminMembersRemoveCmd)
}
```

**Example commands**:
```go
var adminUsersListCmd = &cobra.Command{
    Use:   "list",
    Short: "List users in domain",
    RunE:  runAdminUsersList,
}

var (
    adminUsersListDomain    string
    adminUsersListQuery     string
    adminUsersListLimit     int
    adminUsersListPageToken string
)

func init() {
    adminUsersListCmd.Flags().StringVar(&adminUsersListDomain, "domain", "", "Domain to list users from (required)")
    adminUsersListCmd.Flags().StringVar(&adminUsersListQuery, "query", "", "Query filter")
    adminUsersListCmd.Flags().IntVar(&adminUsersListLimit, "limit", 100, "Max results")
    adminUsersListCmd.Flags().StringVar(&adminUsersListPageToken, "page-token", "", "Page token")
    adminUsersListCmd.MarkFlagRequired("domain")
}

func runAdminUsersList(cmd *cobra.Command, args []string) error {
    ctx := context.Background()

    mgr, creds, err := getAuthAndCreds()
    if err != nil {
        return err
    }

    // Validate admin auth
    if err := mgr.ValidateAdminAuth(creds); err != nil {
        return err
    }

    svc, err := createAdminService(ctx, mgr, creds)
    if err != nil {
        return err
    }

    adminMgr := admin.NewManager(svc)
    opts := &types.ListUsersOptions{
        MaxResults: adminUsersListLimit,
        PageToken:  adminUsersListPageToken,
        Query:      adminUsersListQuery,
    }

    users, err := adminMgr.ListUsers(ctx, adminUsersListDomain, opts)
    if err != nil {
        return err
    }

    return outputWriter.Write(users)
}

var adminUsersCreateCmd = &cobra.Command{
    Use:   "create <email>",
    Short: "Create a new user",
    Args:  cobra.ExactArgs(1),
    RunE:  runAdminUsersCreate,
}

var (
    adminUsersCreateGivenName  string
    adminUsersCreateFamilyName string
    adminUsersCreatePassword   string
)

func init() {
    adminUsersCreateCmd.Flags().StringVar(&adminUsersCreateGivenName, "given-name", "", "First name (required)")
    adminUsersCreateCmd.Flags().StringVar(&adminUsersCreateFamilyName, "family-name", "", "Last name (required)")
    adminUsersCreateCmd.Flags().StringVar(&adminUsersCreatePassword, "password", "", "Password (required)")
    adminUsersCreateCmd.MarkFlagRequired("given-name")
    adminUsersCreateCmd.MarkFlagRequired("family-name")
    adminUsersCreateCmd.MarkFlagRequired("password")
}

func runAdminUsersCreate(cmd *cobra.Command, args []string) error {
    ctx := context.Background()
    email := args[0]

    mgr, creds, err := getAuthAndCreds()
    if err != nil {
        return err
    }

    if err := mgr.ValidateAdminAuth(creds); err != nil {
        return err
    }

    svc, err := createAdminService(ctx, mgr, creds)
    if err != nil {
        return err
    }

    adminMgr := admin.NewManager(svc)
    req := &types.CreateUserRequest{
        Email:      email,
        GivenName:  adminUsersCreateGivenName,
        FamilyName: adminUsersCreateFamilyName,
        Password:   adminUsersCreatePassword,
    }

    user, err := adminMgr.CreateUser(ctx, req)
    if err != nil {
        return err
    }

    return outputWriter.Write(user)
}
```

### 5. Documentation

**Files to modify**:
- `README.md`: Add Admin SDK section with setup instructions
- `go.mod`: Add `google.golang.org/api/admin/directory/v1`

**README additions** (critical setup info):
````markdown
### Admin SDK Operations

**Prerequisites**:
1. **Service Account**: Create service account in Google Cloud Console
2. **Domain-Wide Delegation**: Enable in Google Workspace Admin Console
3. **Scopes**: Add required scopes to service account
   - `https://www.googleapis.com/auth/admin.directory.user`
   - `https://www.googleapis.com/auth/admin.directory.group`

**Setup**:
```bash
# 1. Authenticate with service account
gdrive auth service-account ./service-account-key.json \
  --impersonate-user admin@example.com \
  --scopes https://www.googleapis.com/auth/admin.directory.user,https://www.googleapis.com/auth/admin.directory.group

# 2. Verify authentication
gdrive auth status
```

**User Management**:
```bash
# List users
gdrive admin users list --domain example.com --json

# Get user
gdrive admin users get user@example.com --json

# Create user
gdrive admin users create newuser@example.com \
  --given-name John \
  --family-name Doe \
  --password TempPass123! \
  --json

# Suspend user
gdrive admin users suspend user@example.com

# Delete user
gdrive admin users delete user@example.com
```

**Group Management**:
```bash
# List groups
gdrive admin groups list --domain example.com --json

# Create group
gdrive admin groups create team@example.com \
  --name "Engineering Team" \
  --description "All engineers" \
  --json

# Add member
gdrive admin members add team@example.com user@example.com --role MEMBER

# List members
gdrive admin members list team@example.com --json

# Remove member
gdrive admin members remove team@example.com user@example.com
```
````

## Todo

- [x] Add `ValidateAdminAuth` to `internal/auth/manager.go`
- [x] Create `auth service-account` command in `internal/cli/auth.go`
- [x] Create `internal/admin/manager.go` with user/group/member operations
- [x] Write unit tests in `internal/admin/manager_test.go`
- [x] Create `internal/types/admin.go` with User, Group, Member, request/response types
- [x] Implement TableRenderer for UserList, GroupList, MemberList
- [x] Create `internal/cli/admin.go` with admin, users, groups, members subcommands
- [x] Add validation for domain-wide delegation
- [x] Update `internal/cli/output.go` for Admin types
- [x] Add `google.golang.org/api/admin/directory/v1` to `go.mod`
- [x] Add Admin SDK setup guide to README.md
- [x] Create example service account setup docs
- [x] Write integration tests (requires test Workspace domain)
- [x] Test with domain-wide delegation
- [x] Add error handling for common admin errors (quota, permissions)

## Testing Strategy

1. **Unit tests**: Mock admin.Service responses
2. **Integration tests**: Require test Google Workspace domain with service account
3. **Auth validation**: Test rejection of OAuth credentials
4. **Error scenarios**: Permission denied, user not found, quota exceeded
5. **Bulk operations**: Test listing/updating multiple users

## Security Considerations

1. **Service Account Key Security**: Document secure storage
2. **Scope Minimization**: Use readonly scopes where possible
3. **Audit Logging**: Log all admin operations
4. **Dry Run**: Implement dry-run for destructive operations

## Use Case Examples

**Automated Onboarding**:
```bash
# Read from CSV, create users
cat new_hires.csv | while IFS=, read email first last; do
  gdrive admin users create $email \
    --given-name "$first" \
    --family-name "$last" \
    --password "TempPass$(openssl rand -base64 12)" \
    --json
done
```

**Group Sync**:
```bash
# Sync group membership from file
GROUP=team@example.com
cat members.txt | while read email; do
  gdrive admin members add $GROUP $email --role MEMBER
done
```

**Compliance Reporting**:
```bash
# List suspended users
gdrive admin users list --domain example.com --query "isSuspended=true" --json
```

## Dependencies

Requires:
- Foundation for Google APIs Integration (00_foundation_google_apis.plan.md)

Blocks:
- None (can be developed in parallel, but requires different auth)

## Implementation Order

1. Complete foundation plan first
2. Test service account auth flow separately
3. Implement user operations (most common)
4. Implement group operations
5. Implement member operations last
