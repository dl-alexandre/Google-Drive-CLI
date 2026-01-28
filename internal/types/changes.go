package types

import "time"

// Change represents a change to a file or shared drive
type Change struct {
	// ChangeType indicates the type of change (file, drive)
	ChangeType string `json:"changeType"`

	// FileID is the ID of the file that changed
	FileID string `json:"fileId,omitempty"`

	// File is the file resource (if not removed)
	File *DriveFile `json:"file,omitempty"`

	// Removed indicates if the file was removed
	Removed bool `json:"removed"`

	// Time when the change occurred
	Time time.Time `json:"time"`

	// DriveID is the ID of the shared drive
	DriveID string `json:"driveId,omitempty"`

	// Drive is the shared drive resource
	Drive *SharedDrive `json:"drive,omitempty"`
}

// SharedDrive represents a shared drive
type SharedDrive struct {
	// ID of the shared drive
	ID string `json:"id"`

	// Name of the shared drive
	Name string `json:"name"`

	// ColorRgb is the color of the shared drive as an RGB hex string
	ColorRgb string `json:"colorRgb,omitempty"`

	// BackgroundImageLink is a link to the background image
	BackgroundImageLink string `json:"backgroundImageLink,omitempty"`

	// Capabilities describes the capabilities the current user has on this shared drive
	Capabilities *DriveCapabilities `json:"capabilities,omitempty"`

	// CreatedTime is when the shared drive was created
	CreatedTime time.Time `json:"createdTime,omitempty"`

	// Hidden indicates if the shared drive is hidden from default view
	Hidden bool `json:"hidden"`

	// ThemeID is the ID of the theme from which the background image and color are set
	ThemeID string `json:"themeId,omitempty"`
}

// DriveCapabilities describes capabilities on a shared drive
type DriveCapabilities struct {
	// CanAddChildren indicates if the user can add children to folders in this drive
	CanAddChildren bool `json:"canAddChildren"`

	// CanComment indicates if the user can comment on files in this drive
	CanComment bool `json:"canComment"`

	// CanCopy indicates if the user can copy files in this drive
	CanCopy bool `json:"canCopy"`

	// CanDeleteDrive indicates if the user can delete this drive
	CanDeleteDrive bool `json:"canDeleteDrive"`

	// CanDownload indicates if the user can download files in this drive
	CanDownload bool `json:"canDownload"`

	// CanEdit indicates if the user can edit files in this drive
	CanEdit bool `json:"canEdit"`

	// CanListChildren indicates if the user can list children of folders in this drive
	CanListChildren bool `json:"canListChildren"`

	// CanManageMembers indicates if the user can add members to or remove members from this drive
	CanManageMembers bool `json:"canManageMembers"`

	// CanReadRevisions indicates if the user can read revisions of files in this drive
	CanReadRevisions bool `json:"canReadRevisions"`

	// CanRename indicates if the user can rename files or folders in this drive
	CanRename bool `json:"canRename"`

	// CanRenameDrive indicates if the user can rename this drive
	CanRenameDrive bool `json:"canRenameDrive"`

	// CanShare indicates if the user can share files or folders in this drive
	CanShare bool `json:"canShare"`

	// CanTrashChildren indicates if the user can trash children from folders in this drive
	CanTrashChildren bool `json:"canTrashChildren"`
}

// ChangeList represents a list of changes
type ChangeList struct {
	// Changes is the list of changes
	Changes []Change `json:"changes"`

	// NextPageToken is the page token for the next page of changes
	NextPageToken string `json:"nextPageToken,omitempty"`

	// NewStartPageToken is the starting page token for future changes
	NewStartPageToken string `json:"newStartPageToken,omitempty"`
}

// Channel represents a notification channel for watching changes
type Channel struct {
	// ID is a UUID or similar unique string that identifies this channel
	ID string `json:"id"`

	// ResourceID is an opaque ID that identifies the resource being watched
	ResourceID string `json:"resourceId"`

	// ResourceURI is a version-specific identifier for the watched resource
	ResourceURI string `json:"resourceUri,omitempty"`

	// Token is an arbitrary string delivered to the target address with each notification
	Token string `json:"token,omitempty"`

	// Expiration is the time when this channel will expire (Unix timestamp in milliseconds)
	Expiration int64 `json:"expiration,omitempty"`

	// Type is the type of delivery mechanism used for this channel
	Type string `json:"type,omitempty"`

	// Address is the address where notifications are delivered for this channel
	Address string `json:"address,omitempty"`

	// Params are additional parameters controlling delivery channel behavior
	Params map[string]string `json:"params,omitempty"`
}

// ListOptions configures change list parameters
type ListOptions struct {
	// PageToken is the token for continuing a previous list request
	PageToken string

	// DriveID is the shared drive from which changes are returned
	DriveID string

	// IncludeCorpusRemovals indicates whether changes should include the file resource if the file is still accessible
	IncludeCorpusRemovals bool

	// IncludeItemsFromAllDrives indicates whether both My Drive and shared drive items should be included
	IncludeItemsFromAllDrives bool

	// IncludePermissionsForView specifies which additional view's permissions to include in the response
	IncludePermissionsForView string

	// IncludeRemoved indicates whether to include changes indicating that items have been removed
	IncludeRemoved bool

	// RestrictToMyDrive indicates whether to restrict the results to changes inside the My Drive hierarchy
	RestrictToMyDrive bool

	// SupportsAllDrives indicates whether the requesting application supports both My Drives and shared drives
	SupportsAllDrives bool

	// Limit is the maximum number of changes to return per page
	Limit int

	// Fields specifies which fields to include in the response
	Fields string

	// Spaces is a comma-separated list of spaces to query (drive, appDataFolder, photos)
	Spaces string
}

// WatchOptions configures change watch parameters
type WatchOptions struct {
	// PageToken is the token for continuing a previous list request
	PageToken string

	// DriveID is the shared drive from which changes are returned
	DriveID string

	// IncludeCorpusRemovals indicates whether changes should include the file resource if the file is still accessible
	IncludeCorpusRemovals bool

	// IncludeItemsFromAllDrives indicates whether both My Drive and shared drive items should be included
	IncludeItemsFromAllDrives bool

	// IncludePermissionsForView specifies which additional view's permissions to include in the response
	IncludePermissionsForView string

	// IncludeRemoved indicates whether to include changes indicating that items have been removed
	IncludeRemoved bool

	// RestrictToMyDrive indicates whether to restrict the results to changes inside the My Drive hierarchy
	RestrictToMyDrive bool

	// SupportsAllDrives indicates whether the requesting application supports both My Drives and shared drives
	SupportsAllDrives bool

	// Spaces is a comma-separated list of spaces to query (drive, appDataFolder, photos)
	Spaces string

	// WebhookURL is the URL where notifications should be delivered
	WebhookURL string

	// Expiration is the expiration time for the channel (Unix timestamp in milliseconds)
	Expiration int64

	// Token is an arbitrary string delivered to the target address with each notification
	Token string
}
