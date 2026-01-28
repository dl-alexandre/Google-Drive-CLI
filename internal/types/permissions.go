package types

// AuditResult represents the result of a permission audit operation.
// It contains files that match specific permission criteria (public, external, etc.)
type AuditResult struct {
	// Files is the list of files matching the audit criteria
	Files []*FilePermissionInfo `json:"files"`

	// TotalCount is the total number of files found
	TotalCount int `json:"totalCount"`

	// RiskLevel categorizes the overall risk (low, medium, high, critical)
	RiskLevel string `json:"riskLevel,omitempty"`

	// Summary provides counts by permission type or risk category
	Summary map[string]int `json:"summary,omitempty"`

	// Warnings contains any warnings generated during the audit
	Warnings []string `json:"warnings,omitempty"`
}

// FilePermissionInfo represents a file with its permission details for auditing
type FilePermissionInfo struct {
	// File metadata
	FileID       string `json:"fileId"`
	FileName     string `json:"fileName"`
	MimeType     string `json:"mimeType,omitempty"`
	WebViewLink  string `json:"webViewLink,omitempty"`
	CreatedTime  string `json:"createdTime,omitempty"`
	ModifiedTime string `json:"modifiedTime,omitempty"`

	// Permission details
	Permissions []*Permission `json:"permissions"`

	// Risk assessment
	RiskLevel   string   `json:"riskLevel,omitempty"`   // low, medium, high, critical
	RiskReasons []string `json:"riskReasons,omitempty"` // Reasons for risk classification

	// Access summary
	HasPublicAccess   bool     `json:"hasPublicAccess"`
	HasExternalAccess bool     `json:"hasExternalAccess"`
	HasAnyoneWithLink bool     `json:"hasAnyoneWithLink"`
	ExternalDomains   []string `json:"externalDomains,omitempty"`
	PermissionCount   int      `json:"permissionCount"`
}

// PermissionAnalysis represents a hierarchical analysis of folder permissions
type PermissionAnalysis struct {
	// Folder metadata
	FolderID   string `json:"folderId"`
	FolderName string `json:"folderName"`
	FolderPath string `json:"folderPath,omitempty"`

	// Analysis results
	TotalFiles       int            `json:"totalFiles"`
	TotalFolders     int            `json:"totalFolders"`
	FilesWithRisks   int            `json:"filesWithRisks"`
	FoldersWithRisks int            `json:"foldersWithRisks"`
	RiskDistribution map[string]int `json:"riskDistribution"` // low, medium, high, critical counts
	PermissionTypes  map[string]int `json:"permissionTypes"`  // user, group, domain, anyone counts
	RoleDistribution map[string]int `json:"roleDistribution"` // reader, writer, etc. counts

	// Detailed findings
	PublicFiles    []*FilePermissionInfo `json:"publicFiles,omitempty"`
	ExternalShares []*FilePermissionInfo `json:"externalShares,omitempty"`
	AnyoneWithLink []*FilePermissionInfo `json:"anyoneWithLink,omitempty"`
	HighRiskFiles  []*FilePermissionInfo `json:"highRiskFiles,omitempty"`

	// Recursive analysis (if enabled)
	Recursive  bool                  `json:"recursive"`
	Subfolders []*PermissionAnalysis `json:"subfolders,omitempty"`
}

// PermissionReport represents a detailed permission report for a single file or folder
type PermissionReport struct {
	// Resource metadata
	ResourceID   string `json:"resourceId"`
	ResourceName string `json:"resourceName"`
	ResourceType string `json:"resourceType"` // file or folder
	MimeType     string `json:"mimeType,omitempty"`
	WebViewLink  string `json:"webViewLink,omitempty"`
	CreatedTime  string `json:"createdTime,omitempty"`
	ModifiedTime string `json:"modifiedTime,omitempty"`
	Owner        string `json:"owner,omitempty"`

	// Permission details
	Permissions     []*PermissionDetail `json:"permissions"`
	PermissionCount int                 `json:"permissionCount"`

	// Access analysis
	HasPublicAccess   bool     `json:"hasPublicAccess"`
	HasExternalAccess bool     `json:"hasExternalAccess"`
	HasAnyoneWithLink bool     `json:"hasAnyoneWithLink"`
	ExternalDomains   []string `json:"externalDomains,omitempty"`
	InternalDomain    string   `json:"internalDomain,omitempty"`

	// Risk assessment
	RiskLevel   string   `json:"riskLevel"`
	RiskReasons []string `json:"riskReasons,omitempty"`
	RiskScore   int      `json:"riskScore,omitempty"` // 0-100

	// Recommendations
	Recommendations []string `json:"recommendations,omitempty"`
}

// PermissionDetail represents detailed information about a single permission
type PermissionDetail struct {
	// Permission fields
	ID           string `json:"id"`
	Type         string `json:"type"` // user, group, domain, anyone
	Role         string `json:"role"` // reader, commenter, writer, organizer, owner
	EmailAddress string `json:"emailAddress,omitempty"`
	Domain       string `json:"domain,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`

	// Classification
	IsExternal bool   `json:"isExternal"`
	IsPublic   bool   `json:"isPublic"`
	RiskLevel  string `json:"riskLevel,omitempty"` // low, medium, high, critical
}

// AuditOptions configures permission audit operations
type AuditOptions struct {
	// Scope options
	FolderID  string // Limit audit to specific folder
	Recursive bool   // Include subfolders

	// Filter options
	IncludeTrashed bool   // Include trashed files
	MimeType       string // Filter by MIME type
	Query          string // Additional Drive API query

	// Domain options
	InternalDomain string // Domain to consider as internal (for external detection)

	// Pagination
	PageSize  int    // Number of results per page
	PageToken string // Token for pagination

	// Output options
	IncludePermissions  bool // Include full permission details in results
	IncludeRiskAnalysis bool // Include risk assessment
}

// AnalyzeOptions configures folder permission analysis
type AnalyzeOptions struct {
	// Scope
	Recursive bool // Analyze subfolders recursively
	MaxDepth  int  // Maximum recursion depth (0 = unlimited)

	// Filters
	IncludeTrashed bool // Include trashed items

	// Analysis options
	IncludeDetails bool   // Include detailed file lists
	InternalDomain string // Domain to consider as internal
	RiskThreshold  string // Minimum risk level to include (low, medium, high, critical)

	// Performance
	MaxFiles int // Maximum files to analyze (0 = unlimited)
}

// BulkOptions configures bulk permission operations
type BulkOptions struct {
	// Scope
	FolderID  string // Folder to operate on
	Recursive bool   // Include subfolders

	// Filters
	Query          string // Drive API query to filter files
	MimeType       string // Filter by MIME type
	IncludeTrashed bool   // Include trashed files

	// Safety
	DryRun          bool // Preview operations without executing
	MaxFiles        int  // Maximum files to process (safety limit)
	BatchSize       int  // Number of operations per batch
	ContinueOnError bool // Continue processing if individual operations fail

	// Progress
	ShowProgress bool // Show progress during bulk operations
}

// SearchOptions configures permission search operations
type SearchOptions struct {
	// Search criteria
	Email string // Search by email address
	Role  string // Search by role (reader, writer, etc.)
	Type  string // Search by permission type (user, group, domain, anyone)

	// Scope
	FolderID  string // Limit search to specific folder
	Recursive bool   // Include subfolders

	// Filters
	IncludeTrashed bool   // Include trashed files
	MimeType       string // Filter by MIME type
	Query          string // Additional Drive API query

	// Pagination
	PageSize  int    // Number of results per page
	PageToken string // Token for pagination

	// Output
	IncludePermissions bool // Include full permission details
}

// BulkOperationResult represents the result of a bulk permission operation
type BulkOperationResult struct {
	// Summary
	TotalFiles   int `json:"totalFiles"`
	SuccessCount int `json:"successCount"`
	FailureCount int `json:"failureCount"`
	SkippedCount int `json:"skippedCount"`

	// Details
	SuccessfulFiles []*BulkOperationItem `json:"successfulFiles,omitempty"`
	FailedFiles     []*BulkOperationItem `json:"failedFiles,omitempty"`
	SkippedFiles    []*BulkOperationItem `json:"skippedFiles,omitempty"`

	// Errors
	Errors []string `json:"errors,omitempty"`

	// Dry run
	DryRun bool `json:"dryRun"`
}

// BulkOperationItem represents a single item in a bulk operation
type BulkOperationItem struct {
	FileID       string `json:"fileId"`
	FileName     string `json:"fileName"`
	Operation    string `json:"operation"` // remove, update, etc.
	Status       string `json:"status"`    // success, failure, skipped
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// RiskLevel constants
const (
	RiskLevelLow      = "low"
	RiskLevelMedium   = "medium"
	RiskLevelHigh     = "high"
	RiskLevelCritical = "critical"
)

// PermissionType constants (already defined in Permission, but useful for filtering)
const (
	PermissionTypeUser   = "user"
	PermissionTypeGroup  = "group"
	PermissionTypeDomain = "domain"
	PermissionTypeAnyone = "anyone"
)

// PermissionRole constants (already defined in Permission, but useful for filtering)
const (
	PermissionRoleReader    = "reader"
	PermissionRoleCommenter = "commenter"
	PermissionRoleWriter    = "writer"
	PermissionRoleOrganizer = "organizer"
	PermissionRoleOwner     = "owner"
)
