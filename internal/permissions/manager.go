// Package permissions provides Drive API permission management functionality.
// It handles creating, listing, updating, and deleting permissions on Drive files and folders.
// The package supports all permission types (user, group, domain, anyone) and roles
// (reader, commenter, writer, organizer, owner), including Shared Drive-specific behaviors.
package permissions

import (
	"context"
	"fmt"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/safety"
	"github.com/dl-alexandre/gdrv/internal/types"
	"github.com/dl-alexandre/gdrv/internal/utils"
	"google.golang.org/api/drive/v3"
)

// Manager handles permission operations for Google Drive files and folders.
// It provides methods for creating, listing, updating, and deleting permissions
// with support for:
//   - All permission types: user, group, domain, anyone
//   - All permission roles: reader, commenter, writer, organizer, owner
//   - Notification emails and custom messages
//   - Ownership transfer
//   - Public link creation with discoverability control
//   - Shared Drive-specific permission behaviors
//   - Domain admin access for Workspace environments
//   - Resource key handling for link-shared files
type Manager struct {
	client *api.Client
	shaper *api.RequestShaper
}

// NewManager creates a new permission manager
func NewManager(client *api.Client) *Manager {
	return &Manager{
		client: client,
		shaper: api.NewRequestShaper(client),
	}
}

// CreateOptions configures permission creation.
//
// Type specifies the permission type:
//   - "user": Permission for a specific user (requires EmailAddress)
//   - "group": Permission for a group (requires EmailAddress)
//   - "domain": Permission for an entire domain (requires Domain)
//   - "anyone": Public permission (link-sharing)
//
// Role specifies the permission level:
//   - "reader": Can view and download
//   - "commenter": Can view, download, and comment
//   - "writer": Can view, download, comment, and edit
//   - "organizer": Can organize files in Shared Drives (Shared Drive only)
//   - "owner": Full ownership (requires TransferOwnership=true)
//
// Requirements:
//   - Requirement 4.1: Support user/group/domain/anyone permission types
//   - Requirement 4.2: Support reader/commenter/writer/organizer roles
//   - Requirement 4.3: Support sendNotificationEmail and emailMessage parameters
//   - Requirement 4.4: Support transferOwnership parameter
//   - Requirement 4.5: Support allowFileDiscovery for public permissions
//   - Requirement 4.6: Support useDomainAdminAccess for Workspace environments
type CreateOptions struct {
	Type                  string // user, group, domain, anyone
	Role                  string // reader, commenter, writer, organizer, owner
	EmailAddress          string // Required for user and group types
	Domain                string // Required for domain type
	SendNotificationEmail bool   // Send email notification to recipients
	EmailMessage          string // Custom message to include in notification email
	TransferOwnership     bool   // Transfer ownership (only valid when Role="owner")
	AllowFileDiscovery    bool   // Allow file to be discovered via search (anyone type only)
	UseDomainAdminAccess  bool   // Use domain administrator access for Workspace environments
}

// UpdateOptions configures permission updates.
//
// Requirements:
//   - Requirement 4.10: Modify existing permission levels
//   - Requirement 4.6: Support useDomainAdminAccess for Workspace environments
type UpdateOptions struct {
	Role                 string // New role: reader, commenter, writer, organizer
	UseDomainAdminAccess bool   // Use domain administrator access
}

// DeleteOptions configures permission deletion.
//
// Requirements:
//   - Requirement 4.9: Revoke access for specified user or group
//   - Requirement 4.6: Support useDomainAdminAccess for Workspace environments
type DeleteOptions struct {
	UseDomainAdminAccess bool // Use domain administrator access
}

// ListOptions configures permission listing.
//
// Requirements:
//   - Requirement 4.7: Return all current permissions
//   - Requirement 4.8: Support pagination with pageToken and nextPageToken
//   - Requirement 4.6: Support useDomainAdminAccess for Workspace environments
type ListOptions struct {
	UseDomainAdminAccess bool // Use domain administrator access
	PageSize             int  // Number of permissions per page (0 = API default)
}

// List lists all permissions for a file or folder.
// It automatically handles pagination to retrieve all permissions.
//
// Parameters:
//   - ctx: Context for request cancellation
//   - reqCtx: Request context with profile, drive context, and trace ID
//   - fileID: The ID of the file or folder
//   - opts: Options for listing (domain admin access, page size)
//
// Returns all permissions or an error. Supports Shared Drive files when
// reqCtx.DriveID is set.
//
// Requirements:
//   - Requirement 4.7: Return all current permissions
//   - Requirement 4.8: Iterate through all pages using pageToken
//   - Requirement 4.14: Include supportsAllDrives=true parameter
func (m *Manager) List(ctx context.Context, reqCtx *types.RequestContext, fileID string, opts ListOptions) ([]*types.Permission, error) {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, fileID)

	call := m.client.Service().Permissions.List(fileID)
	call = m.shaper.ShapePermissionsList(call, reqCtx)
	call = call.Fields("permissions(id,type,role,emailAddress,domain,displayName),nextPageToken")

	if opts.UseDomainAdminAccess {
		call = call.UseDomainAdminAccess(true)
	}
	if opts.PageSize > 0 {
		call = call.PageSize(int64(opts.PageSize))
	}

	var allPerms []*types.Permission
	pageToken := ""

	for {
		if pageToken != "" {
			call = call.PageToken(pageToken)
		}

		result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.PermissionList, error) {
			return call.Do()
		})
		if err != nil {
			return nil, err
		}

		for _, p := range result.Permissions {
			allPerms = append(allPerms, convertPermission(p))
		}

		if result.NextPageToken == "" {
			break
		}
		pageToken = result.NextPageToken
	}

	return allPerms, nil
}

// Create creates a new permission on a file or folder.
//
// Supports all permission types (user, group, domain, anyone) and roles
// (reader, commenter, writer, organizer, owner). Can send notification emails,
// transfer ownership, and control discoverability.
//
// Parameters:
//   - ctx: Context for request cancellation
//   - reqCtx: Request context with profile, drive context, and trace ID
//   - fileID: The ID of the file or folder
//   - opts: Permission creation options (type, role, email, etc.)
//
// Returns the created permission or an error. Handles Shared Drive-specific
// behaviors when reqCtx.DriveID is set.
//
// Error Handling:
//   - Returns ErrCodePolicyViolation for domain policy restrictions
//   - Returns ErrCodeSharingRestricted for invalid sharing requests
//   - Returns structured errors with capability indicators
//
// Requirements:
//   - Requirement 4.1: Support user/group/domain/anyone permission types
//   - Requirement 4.2: Support reader/commenter/writer/organizer roles
//   - Requirement 4.3: Support sendNotificationEmail and emailMessage
//   - Requirement 4.4: Support transferOwnership parameter
//   - Requirement 4.5: Support allowFileDiscovery parameter
//   - Requirement 4.6: Support useDomainAdminAccess parameter
//   - Requirement 4.12: Return policy violation errors with capability indicators
//   - Requirement 4.13: Validate ownership transfer restrictions
//   - Requirement 4.14: Include supportsAllDrives=true parameter
func (m *Manager) Create(ctx context.Context, reqCtx *types.RequestContext, fileID string, opts CreateOptions) (*types.Permission, error) {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, fileID)

	perm := &drive.Permission{
		Type: opts.Type,
		Role: opts.Role,
	}
	if opts.EmailAddress != "" {
		perm.EmailAddress = opts.EmailAddress
	}
	if opts.Domain != "" {
		perm.Domain = opts.Domain
	}
	if opts.Type == "anyone" {
		perm.AllowFileDiscovery = opts.AllowFileDiscovery
	}

	call := m.client.Service().Permissions.Create(fileID, perm)
	call = m.shaper.ShapePermissionsCreate(call, reqCtx)
	call = call.SendNotificationEmail(opts.SendNotificationEmail)
	call = call.Fields("id,type,role,emailAddress,domain,displayName")

	if opts.EmailMessage != "" {
		call = call.EmailMessage(opts.EmailMessage)
	}
	if opts.TransferOwnership {
		call = call.TransferOwnership(true)
	}
	if opts.UseDomainAdminAccess {
		call = call.UseDomainAdminAccess(true)
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.Permission, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	return convertPermission(result), nil
}

// Update updates an existing permission's role.
//
// Parameters:
//   - ctx: Context for request cancellation
//   - reqCtx: Request context with profile, drive context, and trace ID
//   - fileID: The ID of the file or folder
//   - permissionID: The ID of the permission to update
//   - opts: Update options (new role, domain admin access)
//
// Returns the updated permission or an error.
//
// Requirements:
//   - Requirement 4.10: Modify existing permission levels
//   - Requirement 4.14: Include supportsAllDrives=true parameter
func (m *Manager) Update(ctx context.Context, reqCtx *types.RequestContext, fileID, permissionID string, opts UpdateOptions) (*types.Permission, error) {
	return m.UpdateWithSafety(ctx, reqCtx, fileID, permissionID, opts, safety.Default(), nil)
}

// UpdateWithSafety updates an existing permission's role with safety controls.
// Supports dry-run mode and confirmation.
//
// Requirements:
//   - Requirement 13.1: Support --dry-run mode for destructive operations
func (m *Manager) UpdateWithSafety(ctx context.Context, reqCtx *types.RequestContext, fileID, permissionID string, opts UpdateOptions, safetyOpts safety.SafetyOptions, recorder safety.DryRunRecorder) (*types.Permission, error) {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, fileID)

	// Get current permission for dry-run display
	if safetyOpts.DryRun && recorder != nil {
		safety.RecordPermissionUpdate(recorder, fileID, fileID, permissionID, opts.Role)
		// Return a placeholder permission
		return &types.Permission{
			ID:   permissionID,
			Role: opts.Role,
		}, nil
	}

	perm := &drive.Permission{
		Role: opts.Role,
	}

	call := m.client.Service().Permissions.Update(fileID, permissionID, perm)
	call = call.SupportsAllDrives(true)
	call = call.Fields("id,type,role,emailAddress,domain,displayName")

	if opts.UseDomainAdminAccess {
		call = call.UseDomainAdminAccess(true)
	}

	header := m.client.ResourceKeys().BuildHeader(reqCtx.InvolvedFileIDs)
	if header != "" {
		call.Header().Set("X-Goog-Drive-Resource-Keys", header)
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.Permission, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	return convertPermission(result), nil
}

// Delete removes a permission from a file or folder.
//
// Parameters:
//   - ctx: Context for request cancellation
//   - reqCtx: Request context with profile, drive context, and trace ID
//   - fileID: The ID of the file or folder
//   - permissionID: The ID of the permission to delete
//   - opts: Delete options (domain admin access)
//
// Returns an error if the deletion fails.
//
// Requirements:
//   - Requirement 4.9: Revoke access for specified user or group
//   - Requirement 4.14: Include supportsAllDrives=true parameter
func (m *Manager) Delete(ctx context.Context, reqCtx *types.RequestContext, fileID, permissionID string, opts DeleteOptions) error {
	return m.DeleteWithSafety(ctx, reqCtx, fileID, permissionID, opts, safety.Default(), nil)
}

// DeleteWithSafety removes a permission from a file or folder with safety controls.
// Supports dry-run mode, confirmation, and idempotency.
//
// Requirements:
//   - Requirement 13.1: Support --dry-run mode for destructive operations
//   - Requirement 13.2: Support --force flag to skip confirmations
func (m *Manager) DeleteWithSafety(ctx context.Context, reqCtx *types.RequestContext, fileID, permissionID string, opts DeleteOptions, safetyOpts safety.SafetyOptions, recorder safety.DryRunRecorder) error {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, fileID)

	// Get permission details for confirmation
	perm, err := m.Get(ctx, reqCtx, fileID, permissionID)
	if err != nil && !safetyOpts.DryRun {
		return err
	}

	// Dry-run mode: record operation without executing
	if safetyOpts.DryRun && recorder != nil {
		safety.RecordPermissionDelete(recorder, fileID, fileID, permissionID)
		return nil
	}

	// Confirmation for destructive operations
	if safetyOpts.ShouldConfirm() {
		displayName := permissionID
		if perm != nil && perm.EmailAddress != "" {
			displayName = perm.EmailAddress
		} else if perm != nil && perm.DisplayName != "" {
			displayName = perm.DisplayName
		}

		confirmed, err := safety.Confirm(
			fmt.Sprintf("About to revoke permission for '%s'. Continue?", displayName),
			safetyOpts,
		)
		if err != nil {
			return err
		}
		if !confirmed {
			return utils.NewAppError(utils.NewCLIError(utils.ErrCodeCancelled, "Operation cancelled by user").Build())
		}
	}

	call := m.client.Service().Permissions.Delete(fileID, permissionID)
	call = call.SupportsAllDrives(true)

	if opts.UseDomainAdminAccess {
		call = call.UseDomainAdminAccess(true)
	}

	header := m.client.ResourceKeys().BuildHeader(reqCtx.InvolvedFileIDs)
	if header != "" {
		call.Header().Set("X-Goog-Drive-Resource-Keys", header)
	}

	_, err = api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (interface{}, error) {
		return nil, call.Do()
	})
	return err
}

// CreatePublicLink creates a public "anyone with link" permission.
//
// This is a convenience method for creating public sharing links.
// For domain-only sharing, use Create with Type="domain".
//
// Parameters:
//   - ctx: Context for request cancellation
//   - reqCtx: Request context with profile, drive context, and trace ID
//   - fileID: The ID of the file or folder
//   - role: Permission role (reader, commenter, writer)
//   - allowDiscovery: If true, file can be discovered via search
//
// Returns the created permission or an error.
//
// Requirements:
//   - Requirement 4.11: Support "anyone with link" sharing
//   - Requirement 4.5: Support allowFileDiscovery for discoverability control
func (m *Manager) CreatePublicLink(ctx context.Context, reqCtx *types.RequestContext, fileID string, role string, allowDiscovery bool) (*types.Permission, error) {
	return m.Create(ctx, reqCtx, fileID, CreateOptions{
		Type:               "anyone",
		Role:               role,
		AllowFileDiscovery: allowDiscovery,
	})
}

// Get retrieves a specific permission by ID.
//
// Parameters:
//   - ctx: Context for request cancellation
//   - reqCtx: Request context with profile, drive context, and trace ID
//   - fileID: The ID of the file or folder
//   - permissionID: The ID of the permission to retrieve
//
// Returns the permission or an error if not found.
//
// Requirements:
//   - Requirement 4.14: Include supportsAllDrives=true parameter
func (m *Manager) Get(ctx context.Context, reqCtx *types.RequestContext, fileID, permissionID string) (*types.Permission, error) {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, fileID)

	call := m.client.Service().Permissions.Get(fileID, permissionID)
	call = call.SupportsAllDrives(true)
	call = call.Fields("id,type,role,emailAddress,domain,displayName")

	header := m.client.ResourceKeys().BuildHeader(reqCtx.InvolvedFileIDs)
	if header != "" {
		call.Header().Set("X-Goog-Drive-Resource-Keys", header)
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.Permission, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	return convertPermission(result), nil
}

func convertPermission(p *drive.Permission) *types.Permission {
	return &types.Permission{
		ID:           p.Id,
		Type:         p.Type,
		Role:         p.Role,
		EmailAddress: p.EmailAddress,
		Domain:       p.Domain,
		DisplayName:  p.DisplayName,
	}
}

// AuditPublic finds all files with public access (type="anyone")
func (m *Manager) AuditPublic(ctx context.Context, reqCtx *types.RequestContext, opts types.AuditOptions) (*types.AuditResult, error) {
	query := "visibility = 'anyoneCanFind' or visibility = 'anyoneWithLink'"
	return m.auditByQuery(ctx, reqCtx, query, opts, func(perms []*types.Permission) bool {
		for _, p := range perms {
			if p.Type == "anyone" {
				return true
			}
		}
		return false
	})
}

// AuditExternal finds all files shared with external domains
func (m *Manager) AuditExternal(ctx context.Context, reqCtx *types.RequestContext, opts types.AuditOptions) (*types.AuditResult, error) {
	if opts.InternalDomain == "" {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidArgument,
			"InternalDomain is required for external audit").Build())
	}

	return m.auditByQuery(ctx, reqCtx, "", opts, func(perms []*types.Permission) bool {
		for _, p := range perms {
			switch p.Type {
			case "user", "group":
				if p.EmailAddress != "" && !isInternalEmail(p.EmailAddress, opts.InternalDomain) {
					return true
				}
			case "domain":
				if p.Domain != "" && p.Domain != opts.InternalDomain {
					return true
				}
			}
		}
		return false
	})
}

// AuditAnyoneWithLink finds all files with "anyone with link" access
func (m *Manager) AuditAnyoneWithLink(ctx context.Context, reqCtx *types.RequestContext, opts types.AuditOptions) (*types.AuditResult, error) {
	query := "visibility = 'anyoneWithLink'"
	return m.auditByQuery(ctx, reqCtx, query, opts, func(perms []*types.Permission) bool {
		for _, p := range perms {
			if p.Type == "anyone" {
				return true
			}
		}
		return false
	})
}

// AuditUser finds all files accessible by a specific user email
func (m *Manager) AuditUser(ctx context.Context, reqCtx *types.RequestContext, email string, opts types.AuditOptions) (*types.AuditResult, error) {
	if email == "" {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidArgument,
			"Email is required for user audit").Build())
	}

	return m.auditByQuery(ctx, reqCtx, "", opts, func(perms []*types.Permission) bool {
		for _, p := range perms {
			if p.EmailAddress == email {
				return true
			}
		}
		return false
	})
}

// AnalyzeFolder analyzes permissions for a folder and optionally its descendants
func (m *Manager) AnalyzeFolder(ctx context.Context, reqCtx *types.RequestContext, folderID string, opts types.AnalyzeOptions) (*types.PermissionAnalysis, error) {
	reqCtx.InvolvedParentIDs = append(reqCtx.InvolvedParentIDs, folderID)

	filesManager := m.client.Service().Files

	folderCall := filesManager.Get(folderID).Fields("id,name,mimeType")
	folderCall = folderCall.SupportsAllDrives(true)
	header := m.client.ResourceKeys().BuildHeader(reqCtx.InvolvedParentIDs)
	if header != "" {
		folderCall.Header().Set("X-Goog-Drive-Resource-Keys", header)
	}

	folder, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.File, error) {
		return folderCall.Do()
	})
	if err != nil {
		return nil, err
	}

	analysis := &types.PermissionAnalysis{
		FolderID:         folderID,
		FolderName:       folder.Name,
		Recursive:        opts.Recursive,
		RiskDistribution: make(map[string]int),
		PermissionTypes:  make(map[string]int),
		RoleDistribution: make(map[string]int),
	}

	query := fmt.Sprintf("'%s' in parents", folderID)
	if !opts.IncludeTrashed {
		query += " and trashed = false"
	}

	listCall := filesManager.List().Q(query).Fields("files(id,name,mimeType,webViewLink,createdTime,modifiedTime)")
	listCall = m.shaper.ShapeFilesList(listCall, reqCtx)

	fileList, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.FileList, error) {
		return listCall.Do()
	})
	if err != nil {
		return nil, err
	}

	for _, file := range fileList.Files {
		if file.MimeType == "application/vnd.google-apps.folder" {
			analysis.TotalFolders++
		} else {
			analysis.TotalFiles++
		}

		perms, err := m.List(ctx, reqCtx, file.Id, ListOptions{})
		if err != nil {
			continue
		}

		fileInfo := analyzeFilePermissions(file, perms, opts.InternalDomain)

		for _, p := range perms {
			analysis.PermissionTypes[p.Type]++
			analysis.RoleDistribution[p.Role]++
		}

		analysis.RiskDistribution[fileInfo.RiskLevel]++

		if fileInfo.RiskLevel == types.RiskLevelHigh || fileInfo.RiskLevel == types.RiskLevelCritical {
			if file.MimeType == "application/vnd.google-apps.folder" {
				analysis.FoldersWithRisks++
			} else {
				analysis.FilesWithRisks++
			}
		}

		if opts.IncludeDetails {
			if fileInfo.HasPublicAccess {
				analysis.PublicFiles = append(analysis.PublicFiles, fileInfo)
			}
			if fileInfo.HasExternalAccess {
				analysis.ExternalShares = append(analysis.ExternalShares, fileInfo)
			}
			if fileInfo.HasAnyoneWithLink {
				analysis.AnyoneWithLink = append(analysis.AnyoneWithLink, fileInfo)
			}
			if fileInfo.RiskLevel == types.RiskLevelHigh || fileInfo.RiskLevel == types.RiskLevelCritical {
				analysis.HighRiskFiles = append(analysis.HighRiskFiles, fileInfo)
			}
		}

		if opts.MaxFiles > 0 && (analysis.TotalFiles+analysis.TotalFolders) >= opts.MaxFiles {
			break
		}
	}

	if opts.Recursive {
		for _, file := range fileList.Files {
			if file.MimeType == "application/vnd.google-apps.folder" {
				if opts.MaxDepth > 0 {
					opts.MaxDepth--
					if opts.MaxDepth == 0 {
						break
					}
				}
				subAnalysis, err := m.AnalyzeFolder(ctx, reqCtx, file.Id, opts)
				if err == nil {
					analysis.Subfolders = append(analysis.Subfolders, subAnalysis)
				}
			}
		}
	}

	return analysis, nil
}

// GenerateReport generates a detailed permission report for a file or folder
func (m *Manager) GenerateReport(ctx context.Context, reqCtx *types.RequestContext, fileID string, internalDomain string) (*types.PermissionReport, error) {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, fileID)

	fileCall := m.client.Service().Files.Get(fileID).Fields("id,name,mimeType,webViewLink,createdTime,modifiedTime,owners")
	fileCall = m.shaper.ShapeFilesGet(fileCall, reqCtx)

	file, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.File, error) {
		return fileCall.Do()
	})
	if err != nil {
		return nil, err
	}

	perms, err := m.List(ctx, reqCtx, fileID, ListOptions{})
	if err != nil {
		return nil, err
	}

	report := &types.PermissionReport{
		ResourceID:      fileID,
		ResourceName:    file.Name,
		ResourceType:    "file",
		MimeType:        file.MimeType,
		WebViewLink:     file.WebViewLink,
		CreatedTime:     file.CreatedTime,
		ModifiedTime:    file.ModifiedTime,
		InternalDomain:  internalDomain,
		PermissionCount: len(perms),
		Permissions:     make([]*types.PermissionDetail, 0, len(perms)),
	}

	if file.MimeType == "application/vnd.google-apps.folder" {
		report.ResourceType = "folder"
	}

	if len(file.Owners) > 0 {
		report.Owner = file.Owners[0].EmailAddress
	}

	externalDomains := make(map[string]bool)
	riskScore := 0

	for _, p := range perms {
		detail := &types.PermissionDetail{
			ID:           p.ID,
			Type:         p.Type,
			Role:         p.Role,
			EmailAddress: p.EmailAddress,
			Domain:       p.Domain,
			DisplayName:  p.DisplayName,
		}

		if p.Type == "anyone" {
			detail.IsPublic = true
			report.HasPublicAccess = true
			report.HasAnyoneWithLink = true
			detail.RiskLevel = types.RiskLevelCritical
			riskScore += 40
		} else if p.Type == "domain" && p.Domain != internalDomain {
			detail.IsExternal = true
			report.HasExternalAccess = true
			externalDomains[p.Domain] = true
			detail.RiskLevel = types.RiskLevelHigh
			riskScore += 20
		} else if (p.Type == "user" || p.Type == "group") && p.EmailAddress != "" {
			if !isInternalEmail(p.EmailAddress, internalDomain) {
				detail.IsExternal = true
				report.HasExternalAccess = true
				domain := extractDomain(p.EmailAddress)
				if domain != "" {
					externalDomains[domain] = true
				}
				detail.RiskLevel = types.RiskLevelMedium
				riskScore += 10
			} else {
				detail.RiskLevel = types.RiskLevelLow
			}
		}

		switch p.Role {
		case "writer", "organizer":
			if detail.IsPublic || detail.IsExternal {
				riskScore += 10
			}
		case "owner":
			if detail.IsExternal {
				riskScore += 20
			}
		}

		report.Permissions = append(report.Permissions, detail)
	}

	for domain := range externalDomains {
		report.ExternalDomains = append(report.ExternalDomains, domain)
	}

	report.RiskScore = riskScore
	if riskScore >= 60 {
		report.RiskLevel = types.RiskLevelCritical
		report.RiskReasons = append(report.RiskReasons, "Multiple high-risk permissions detected")
	} else if riskScore >= 40 {
		report.RiskLevel = types.RiskLevelHigh
		report.RiskReasons = append(report.RiskReasons, "Public or external access with elevated permissions")
	} else if riskScore >= 20 {
		report.RiskLevel = types.RiskLevelMedium
		report.RiskReasons = append(report.RiskReasons, "External access detected")
	} else {
		report.RiskLevel = types.RiskLevelLow
	}

	if report.HasPublicAccess {
		report.Recommendations = append(report.Recommendations, "Consider removing public access and sharing with specific users")
	}
	if report.HasExternalAccess {
		report.Recommendations = append(report.Recommendations, "Review external shares and ensure they are necessary")
	}
	if len(report.Permissions) > 10 {
		report.Recommendations = append(report.Recommendations, "Consider using groups to simplify permission management")
	}

	return report, nil
}

// BulkRemovePublic removes public access from all files in a folder
func (m *Manager) BulkRemovePublic(ctx context.Context, reqCtx *types.RequestContext, opts types.BulkOptions) (*types.BulkOperationResult, error) {
	if opts.FolderID == "" {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidArgument,
			"FolderID is required for bulk operations").Build())
	}

	result := &types.BulkOperationResult{
		DryRun: opts.DryRun,
	}

	files, err := m.findFilesInFolder(ctx, reqCtx, opts)
	if err != nil {
		return nil, err
	}

	if opts.MaxFiles > 0 && len(files) > opts.MaxFiles {
		files = files[:opts.MaxFiles]
	}

	result.TotalFiles = len(files)

	for _, file := range files {
		perms, err := m.List(ctx, reqCtx, file.Id, ListOptions{})
		if err != nil {
			result.FailureCount++
			result.FailedFiles = append(result.FailedFiles, &types.BulkOperationItem{
				FileID:       file.Id,
				FileName:     file.Name,
				Operation:    "remove_public",
				Status:       "failure",
				ErrorMessage: err.Error(),
			})
			if !opts.ContinueOnError {
				return result, err
			}
			continue
		}

		hasPublic := false
		for _, p := range perms {
			if p.Type == "anyone" {
				hasPublic = true
				if opts.DryRun {
					result.SuccessCount++
					result.SuccessfulFiles = append(result.SuccessfulFiles, &types.BulkOperationItem{
						FileID:    file.Id,
						FileName:  file.Name,
						Operation: "remove_public",
						Status:    "success",
					})
				} else {
					err := m.Delete(ctx, reqCtx, file.Id, p.ID, DeleteOptions{})
					if err != nil {
						result.FailureCount++
						result.FailedFiles = append(result.FailedFiles, &types.BulkOperationItem{
							FileID:       file.Id,
							FileName:     file.Name,
							Operation:    "remove_public",
							Status:       "failure",
							ErrorMessage: err.Error(),
						})
						if !opts.ContinueOnError {
							return result, err
						}
					} else {
						result.SuccessCount++
						result.SuccessfulFiles = append(result.SuccessfulFiles, &types.BulkOperationItem{
							FileID:    file.Id,
							FileName:  file.Name,
							Operation: "remove_public",
							Status:    "success",
						})
					}
				}
				break
			}
		}

		if !hasPublic {
			result.SkippedCount++
			result.SkippedFiles = append(result.SkippedFiles, &types.BulkOperationItem{
				FileID:    file.Id,
				FileName:  file.Name,
				Operation: "remove_public",
				Status:    "skipped",
			})
		}
	}

	return result, nil
}

// BulkUpdateRole updates permissions from one role to another in a folder
func (m *Manager) BulkUpdateRole(ctx context.Context, reqCtx *types.RequestContext, fromRole, toRole string, opts types.BulkOptions) (*types.BulkOperationResult, error) {
	if opts.FolderID == "" {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidArgument,
			"FolderID is required for bulk operations").Build())
	}

	if fromRole == "" || toRole == "" {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidArgument,
			"Both fromRole and toRole are required").Build())
	}

	result := &types.BulkOperationResult{
		DryRun: opts.DryRun,
	}

	files, err := m.findFilesInFolder(ctx, reqCtx, opts)
	if err != nil {
		return nil, err
	}

	if opts.MaxFiles > 0 && len(files) > opts.MaxFiles {
		files = files[:opts.MaxFiles]
	}

	result.TotalFiles = len(files)

	for _, file := range files {
		perms, err := m.List(ctx, reqCtx, file.Id, ListOptions{})
		if err != nil {
			result.FailureCount++
			result.FailedFiles = append(result.FailedFiles, &types.BulkOperationItem{
				FileID:       file.Id,
				FileName:     file.Name,
				Operation:    "update_role",
				Status:       "failure",
				ErrorMessage: err.Error(),
			})
			if !opts.ContinueOnError {
				return result, err
			}
			continue
		}

		updated := false
		for _, p := range perms {
			if p.Role == fromRole {
				if opts.DryRun {
					result.SuccessCount++
					result.SuccessfulFiles = append(result.SuccessfulFiles, &types.BulkOperationItem{
						FileID:    file.Id,
						FileName:  file.Name,
						Operation: "update_role",
						Status:    "success",
					})
					updated = true
				} else {
					_, err := m.Update(ctx, reqCtx, file.Id, p.ID, UpdateOptions{Role: toRole})
					if err != nil {
						result.FailureCount++
						result.FailedFiles = append(result.FailedFiles, &types.BulkOperationItem{
							FileID:       file.Id,
							FileName:     file.Name,
							Operation:    "update_role",
							Status:       "failure",
							ErrorMessage: err.Error(),
						})
						if !opts.ContinueOnError {
							return result, err
						}
					} else {
						result.SuccessCount++
						result.SuccessfulFiles = append(result.SuccessfulFiles, &types.BulkOperationItem{
							FileID:    file.Id,
							FileName:  file.Name,
							Operation: "update_role",
							Status:    "success",
						})
						updated = true
					}
				}
			}
		}

		if !updated {
			result.SkippedCount++
			result.SkippedFiles = append(result.SkippedFiles, &types.BulkOperationItem{
				FileID:    file.Id,
				FileName:  file.Name,
				Operation: "update_role",
				Status:    "skipped",
			})
		}
	}

	return result, nil
}

// SearchByEmail finds all files accessible by a specific email address
func (m *Manager) SearchByEmail(ctx context.Context, reqCtx *types.RequestContext, opts types.SearchOptions) (*types.AuditResult, error) {
	if opts.Email == "" {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidArgument,
			"Email is required for search").Build())
	}

	auditOpts := types.AuditOptions{
		FolderID:           opts.FolderID,
		Recursive:          opts.Recursive,
		IncludeTrashed:     opts.IncludeTrashed,
		MimeType:           opts.MimeType,
		Query:              opts.Query,
		PageSize:           opts.PageSize,
		PageToken:          opts.PageToken,
		IncludePermissions: opts.IncludePermissions,
	}

	return m.AuditUser(ctx, reqCtx, opts.Email, auditOpts)
}

// SearchByRole finds all files with a specific permission role
func (m *Manager) SearchByRole(ctx context.Context, reqCtx *types.RequestContext, opts types.SearchOptions) (*types.AuditResult, error) {
	if opts.Role == "" {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidArgument,
			"Role is required for search").Build())
	}

	auditOpts := types.AuditOptions{
		FolderID:           opts.FolderID,
		Recursive:          opts.Recursive,
		IncludeTrashed:     opts.IncludeTrashed,
		MimeType:           opts.MimeType,
		Query:              opts.Query,
		PageSize:           opts.PageSize,
		PageToken:          opts.PageToken,
		IncludePermissions: opts.IncludePermissions,
	}

	return m.auditByQuery(ctx, reqCtx, "", auditOpts, func(perms []*types.Permission) bool {
		for _, p := range perms {
			if p.Role == opts.Role {
				return true
			}
		}
		return false
	})
}

func (m *Manager) auditByQuery(ctx context.Context, reqCtx *types.RequestContext, baseQuery string, opts types.AuditOptions, filter func([]*types.Permission) bool) (*types.AuditResult, error) {
	query := baseQuery
	if opts.FolderID != "" {
		if query != "" {
			query += " and "
		}
		query += fmt.Sprintf("'%s' in parents", opts.FolderID)
		reqCtx.InvolvedParentIDs = append(reqCtx.InvolvedParentIDs, opts.FolderID)
	}
	if !opts.IncludeTrashed {
		if query != "" {
			query += " and "
		}
		query += "trashed = false"
	}
	if opts.MimeType != "" {
		if query != "" {
			query += " and "
		}
		query += fmt.Sprintf("mimeType = '%s'", opts.MimeType)
	}
	if opts.Query != "" {
		if query != "" {
			query += " and "
		}
		query += opts.Query
	}

	listCall := m.client.Service().Files.List()
	listCall = m.shaper.ShapeFilesList(listCall, reqCtx)
	if query != "" {
		listCall = listCall.Q(query)
	}
	listCall = listCall.Fields("files(id,name,mimeType,webViewLink,createdTime,modifiedTime)")
	if opts.PageSize > 0 {
		listCall = listCall.PageSize(int64(opts.PageSize))
	}
	if opts.PageToken != "" {
		listCall = listCall.PageToken(opts.PageToken)
	}

	fileList, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.FileList, error) {
		return listCall.Do()
	})
	if err != nil {
		return nil, err
	}

	result := &types.AuditResult{
		Files:   make([]*types.FilePermissionInfo, 0),
		Summary: make(map[string]int),
	}

	for _, file := range fileList.Files {
		perms, err := m.List(ctx, reqCtx, file.Id, ListOptions{})
		if err != nil {
			continue
		}

		if filter(perms) {
			fileInfo := &types.FilePermissionInfo{
				FileID:          file.Id,
				FileName:        file.Name,
				MimeType:        file.MimeType,
				WebViewLink:     file.WebViewLink,
				CreatedTime:     file.CreatedTime,
				ModifiedTime:    file.ModifiedTime,
				PermissionCount: len(perms),
			}

			if opts.IncludePermissions {
				fileInfo.Permissions = perms
			}

			fileInfo = analyzeFilePermissions(file, perms, opts.InternalDomain)
			result.Files = append(result.Files, fileInfo)
			result.Summary[fileInfo.RiskLevel]++
		}
	}

	result.TotalCount = len(result.Files)

	if result.TotalCount == 0 {
		result.RiskLevel = types.RiskLevelLow
	} else {
		criticalCount := result.Summary[types.RiskLevelCritical]
		highCount := result.Summary[types.RiskLevelHigh]
		if criticalCount > 0 {
			result.RiskLevel = types.RiskLevelCritical
		} else if highCount > 0 {
			result.RiskLevel = types.RiskLevelHigh
		} else if result.Summary[types.RiskLevelMedium] > 0 {
			result.RiskLevel = types.RiskLevelMedium
		} else {
			result.RiskLevel = types.RiskLevelLow
		}
	}

	return result, nil
}

func (m *Manager) findFilesInFolder(ctx context.Context, reqCtx *types.RequestContext, opts types.BulkOptions) ([]*drive.File, error) {
	query := fmt.Sprintf("'%s' in parents", opts.FolderID)
	if !opts.IncludeTrashed {
		query += " and trashed = false"
	}
	if opts.MimeType != "" {
		query += fmt.Sprintf(" and mimeType = '%s'", opts.MimeType)
	}
	if opts.Query != "" {
		query += " and " + opts.Query
	}

	listCall := m.client.Service().Files.List().Q(query).Fields("files(id,name,mimeType)")
	listCall = m.shaper.ShapeFilesList(listCall, reqCtx)

	fileList, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.FileList, error) {
		return listCall.Do()
	})
	if err != nil {
		return nil, err
	}

	files := fileList.Files

	if opts.Recursive {
		for _, file := range fileList.Files {
			if file.MimeType == "application/vnd.google-apps.folder" {
				subOpts := opts
				subOpts.FolderID = file.Id
				subFiles, err := m.findFilesInFolder(ctx, reqCtx, subOpts)
				if err == nil {
					files = append(files, subFiles...)
				}
			}
		}
	}

	return files, nil
}

func analyzeFilePermissions(file *drive.File, perms []*types.Permission, internalDomain string) *types.FilePermissionInfo {
	info := &types.FilePermissionInfo{
		FileID:          file.Id,
		FileName:        file.Name,
		MimeType:        file.MimeType,
		WebViewLink:     file.WebViewLink,
		CreatedTime:     file.CreatedTime,
		ModifiedTime:    file.ModifiedTime,
		Permissions:     perms,
		PermissionCount: len(perms),
		RiskReasons:     make([]string, 0),
		ExternalDomains: make([]string, 0),
	}

	externalDomains := make(map[string]bool)
	riskScore := 0

	for _, p := range perms {
		if p.Type == "anyone" {
			info.HasPublicAccess = true
			info.HasAnyoneWithLink = true
			info.RiskReasons = append(info.RiskReasons, "Public access enabled")
			riskScore += 40
		} else if p.Type == "domain" && p.Domain != internalDomain {
			info.HasExternalAccess = true
			externalDomains[p.Domain] = true
			info.RiskReasons = append(info.RiskReasons, fmt.Sprintf("Shared with external domain: %s", p.Domain))
			riskScore += 20
		} else if (p.Type == "user" || p.Type == "group") && p.EmailAddress != "" {
			if !isInternalEmail(p.EmailAddress, internalDomain) {
				info.HasExternalAccess = true
				domain := extractDomain(p.EmailAddress)
				if domain != "" {
					externalDomains[domain] = true
				}
				info.RiskReasons = append(info.RiskReasons, fmt.Sprintf("Shared with external user: %s", p.EmailAddress))
				riskScore += 10
			}
		}

		if p.Role == "writer" || p.Role == "organizer" {
			if info.HasPublicAccess || info.HasExternalAccess {
				riskScore += 10
			}
		}
	}

	for domain := range externalDomains {
		info.ExternalDomains = append(info.ExternalDomains, domain)
	}

	if riskScore >= 60 {
		info.RiskLevel = types.RiskLevelCritical
	} else if riskScore >= 40 {
		info.RiskLevel = types.RiskLevelHigh
	} else if riskScore >= 20 {
		info.RiskLevel = types.RiskLevelMedium
	} else {
		info.RiskLevel = types.RiskLevelLow
	}

	return info
}

func isInternalEmail(email, internalDomain string) bool {
	if internalDomain == "" {
		return true
	}
	domain := extractDomain(email)
	return domain == internalDomain
}

func extractDomain(email string) string {
	parts := []rune(email)
	for i := len(parts) - 1; i >= 0; i-- {
		if parts[i] == '@' {
			return string(parts[i+1:])
		}
	}
	return ""
}
