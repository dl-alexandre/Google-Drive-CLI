package utils

// Upload thresholds (binary units)
const (
	UploadSimpleMaxBytes = 5 * 1024 * 1024  // 5 MiB
	UploadChunkSize      = 8 * 1024 * 1024  // 8 MiB
	ExportMaxBytes       = 10 * 1024 * 1024 // 10 MiB
)

// Revision limits
const RevisionKeepForeverLimit = 200

// OAuth scopes
const (
	ScopeFull                        = "https://www.googleapis.com/auth/drive"
	ScopeFile                        = "https://www.googleapis.com/auth/drive.file"
	ScopeReadonly                    = "https://www.googleapis.com/auth/drive.readonly"
	ScopeMetadataReadonly            = "https://www.googleapis.com/auth/drive.metadata.readonly"
	ScopeAppdata                     = "https://www.googleapis.com/auth/drive.appdata"
	ScopeSheets                      = "https://www.googleapis.com/auth/spreadsheets"
	ScopeSheetsReadonly              = "https://www.googleapis.com/auth/spreadsheets.readonly"
	ScopeDocs                        = "https://www.googleapis.com/auth/documents"
	ScopeDocsReadonly                = "https://www.googleapis.com/auth/documents.readonly"
	ScopeSlides                      = "https://www.googleapis.com/auth/presentations"
	ScopeSlidesReadonly              = "https://www.googleapis.com/auth/presentations.readonly"
	ScopeAdminDirectoryUser          = "https://www.googleapis.com/auth/admin.directory.user"
	ScopeAdminDirectoryUserReadonly  = "https://www.googleapis.com/auth/admin.directory.user.readonly"
	ScopeAdminDirectoryGroup         = "https://www.googleapis.com/auth/admin.directory.group"
	ScopeAdminDirectoryGroupReadonly = "https://www.googleapis.com/auth/admin.directory.group.readonly"
	ScopeLabels                      = "https://www.googleapis.com/auth/drive.labels"
	ScopeLabelsReadonly              = "https://www.googleapis.com/auth/drive.labels.readonly"
	ScopeAdminLabels                 = "https://www.googleapis.com/auth/drive.admin.labels"
	ScopeAdminLabelsReadonly         = "https://www.googleapis.com/auth/drive.admin.labels.readonly"
	ScopeActivity                    = "https://www.googleapis.com/auth/drive.activity"
	ScopeActivityReadonly            = "https://www.googleapis.com/auth/drive.activity.readonly"
)

var (
	ScopesWorkspaceBasic = []string{
		ScopeFile,
		ScopeReadonly,
		ScopeMetadataReadonly,
		ScopeSheetsReadonly,
		ScopeDocsReadonly,
		ScopeSlidesReadonly,
		ScopeLabelsReadonly,
	}
	ScopesWorkspaceFull = []string{
		ScopeFull,
		ScopeSheets,
		ScopeDocs,
		ScopeSlides,
		ScopeLabels,
	}
	ScopesAdmin = []string{
		ScopeAdminDirectoryUser,
		ScopeAdminDirectoryGroup,
		ScopeAdminLabels,
	}
	ScopesWorkspaceWithAdmin = []string{
		ScopeFull,
		ScopeSheets,
		ScopeDocs,
		ScopeSlides,
		ScopeAdminDirectoryUser,
		ScopeAdminDirectoryGroup,
		ScopeLabels,
		ScopeAdminLabels,
	}
	// New presets for additional APIs
	ScopesWorkspaceActivity = []string{
		ScopeFile,
		ScopeReadonly,
		ScopeMetadataReadonly,
		ScopeSheetsReadonly,
		ScopeDocsReadonly,
		ScopeSlidesReadonly,
		ScopeLabelsReadonly,
		ScopeActivityReadonly,
	}
	ScopesWorkspaceLabels = []string{
		ScopeFull,
		ScopeSheets,
		ScopeDocs,
		ScopeSlides,
		ScopeLabels,
	}
	ScopesWorkspaceSync = []string{
		ScopeFull,
		ScopeSheets,
		ScopeDocs,
		ScopeSlides,
		ScopeLabels,
		// Changes API uses standard Drive scopes
	}
	ScopesWorkspaceComplete = []string{
		ScopeFull,
		ScopeSheets,
		ScopeDocs,
		ScopeSlides,
		ScopeLabels,
		ScopeActivity,
		// Changes API uses standard Drive scopes
	}
)

// Drive API base URLs
const (
	DriveAPIBase    = "https://www.googleapis.com/drive/v3"
	DriveUploadBase = "https://www.googleapis.com/upload/drive/v3"
)

// Retry configuration
const (
	DefaultMaxRetries   = 3
	DefaultRetryDelayMs = 1000
	MaxRetryDelayMs     = 32000
)

// Cache TTL
const DefaultCacheTTLSeconds = 300

// Schema version
const SchemaVersion = "1.0"

// Google Workspace MIME types
const (
	MimeTypeDocument     = "application/vnd.google-apps.document"
	MimeTypeSpreadsheet  = "application/vnd.google-apps.spreadsheet"
	MimeTypePresentation = "application/vnd.google-apps.presentation"
	MimeTypeDrawing      = "application/vnd.google-apps.drawing"
	MimeTypeForm         = "application/vnd.google-apps.form"
	MimeTypeScript       = "application/vnd.google-apps.script"
	MimeTypeFolder       = "application/vnd.google-apps.folder"
	MimeTypeShortcut     = "application/vnd.google-apps.shortcut"
)

// FormatMappings maps convenience format names to MIME types
var FormatMappings = map[string]string{
	"pdf":  "application/pdf",
	"docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	"xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	"pptx": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	"txt":  "text/plain",
	"html": "text/html",
	"csv":  "text/csv",
	"png":  "image/png",
	"jpg":  "image/jpeg",
	"svg":  "image/svg+xml",
}

// IsWorkspaceMimeType checks if a MIME type is a Google Workspace type
func IsWorkspaceMimeType(mimeType string) bool {
	switch mimeType {
	case MimeTypeDocument, MimeTypeSpreadsheet, MimeTypePresentation,
		MimeTypeDrawing, MimeTypeForm, MimeTypeScript:
		return true
	}
	return false
}
