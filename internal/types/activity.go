package types

import "time"

// Activity represents a Drive activity event
type Activity struct {
	// Timestamp when the activity occurred
	Timestamp time.Time `json:"timestamp"`

	// PrimaryActionDetail describes the primary action
	PrimaryActionDetail ActionDetail `json:"primaryActionDetail"`

	// Actors who performed the activity
	Actors []Actor `json:"actors,omitempty"`

	// Targets affected by the activity
	Targets []Target `json:"targets,omitempty"`

	// Actions performed in this activity
	Actions []Action `json:"actions,omitempty"`
}

// ActionDetail describes a specific action
type ActionDetail struct {
	// Type of action (edit, comment, share, permission_change, move, delete, restore, etc.)
	Type string `json:"type"`

	// Description of the action
	Description string `json:"description,omitempty"`
}

// Actor represents an entity that performed an activity
type Actor struct {
	// Type of actor (user, administrator, system, anonymous)
	Type string `json:"type"`

	// User information (if type is user)
	User *ActivityUser `json:"user,omitempty"`
}

// ActivityUser represents a user actor in an activity
type ActivityUser struct {
	// Email address of the user
	Email string `json:"email,omitempty"`

	// Display name of the user
	DisplayName string `json:"displayName,omitempty"`
}

// Target represents an item affected by an activity
type Target struct {
	// Type of target (driveItem, drive, fileComment)
	Type string `json:"type"`

	// DriveItem information (if type is driveItem)
	DriveItem *DriveItem `json:"driveItem,omitempty"`
}

// DriveItem represents a file or folder target
type DriveItem struct {
	// Name of the item
	Name string `json:"name,omitempty"`

	// Title of the item
	Title string `json:"title,omitempty"`

	// MIME type of the item
	MimeType string `json:"mimeType,omitempty"`

	// Owner of the item
	Owner *ActivityUser `json:"owner,omitempty"`
}

// Action represents an action performed in an activity
type Action struct {
	// Type of action
	Type string `json:"type"`

	// Detail provides action-specific information
	Detail ActionDetail `json:"detail"`
}

// QueryOptions configures activity query parameters
type QueryOptions struct {
	// FileID filters activity for a specific file
	FileID string

	// FolderID filters activity for a specific folder
	FolderID string

	// AncestorName filters activity for items under an ancestor (e.g., "folders/123")
	AncestorName string

	// StartTime filters activity after this time (RFC3339 format)
	StartTime string

	// EndTime filters activity before this time (RFC3339 format)
	EndTime string

	// ActionTypes filters by action types (comma-separated: edit,comment,share,permission_change,move,delete,restore)
	ActionTypes string

	// User filters activity by user email
	User string

	// Limit is the maximum number of results per page
	Limit int

	// PageToken for pagination
	PageToken string

	// Fields to return (optional)
	Fields string
}

// ActivityQueryResult represents the result of an activity query
type ActivityQueryResult struct {
	// Activities returned by the query
	Activities []Activity `json:"activities"`

	// NextPageToken for pagination
	NextPageToken string `json:"nextPageToken,omitempty"`
}
