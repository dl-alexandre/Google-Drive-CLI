package types

// DriveFile represents a Google Drive file
type DriveFile struct {
	ID             string            `json:"id"`
	Name           string            `json:"name"`
	MimeType       string            `json:"mimeType"`
	Size           int64             `json:"size,omitempty"`
	MD5Checksum    string            `json:"md5Checksum,omitempty"`
	CreatedTime    string            `json:"createdTime,omitempty"`
	ModifiedTime   string            `json:"modifiedTime,omitempty"`
	Parents        []string          `json:"parents,omitempty"`
	Capabilities   *FileCapabilities `json:"capabilities,omitempty"`
	ResourceKey    string            `json:"resourceKey,omitempty"`
	ExportLinks    map[string]string `json:"exportLinks,omitempty"`
	WebViewLink    string            `json:"webViewLink,omitempty"`
	WebContentLink string            `json:"webContentLink,omitempty"`
	Trashed        bool              `json:"trashed,omitempty"`
}

// FileCapabilities represents what actions can be performed on a file
type FileCapabilities struct {
	CanDownload      bool `json:"canDownload"`
	CanEdit          bool `json:"canEdit"`
	CanShare         bool `json:"canShare"`
	CanDelete        bool `json:"canDelete"`
	CanTrash         bool `json:"canTrash"`
	CanReadRevisions bool `json:"canReadRevisions"`
}

// FileListResult represents paginated file list response
type FileListResult struct {
	Files            []*DriveFile `json:"files"`
	NextPageToken    string       `json:"nextPageToken,omitempty"`
	IncompleteSearch bool         `json:"incompleteSearch,omitempty"`
}

// Permission represents a Drive permission
type Permission struct {
	ID           string `json:"id"`
	Type         string `json:"type"` // user, group, domain, anyone
	Role         string `json:"role"` // reader, commenter, writer, organizer, owner
	EmailAddress string `json:"emailAddress,omitempty"`
	Domain       string `json:"domain,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
}

// Revision represents a file revision
type Revision struct {
	ID               string `json:"id"`
	ModifiedTime     string `json:"modifiedTime"`
	KeepForever      bool   `json:"keepForever"`
	Size             int64  `json:"size,omitempty"`
	MimeType         string `json:"mimeType,omitempty"`
	OriginalFilename string `json:"originalFilename,omitempty"`
}
