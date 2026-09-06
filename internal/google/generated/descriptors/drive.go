// Google Drive API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package descriptors

// ServiceDescriptor contains metadata for the drive API
type ServiceDescriptor struct {
	Name        string
	Version     string
	BaseURL     string
	RootURL     string
	ServicePath string
}

// Service returns the service descriptor for drive
var Service = ServiceDescriptor{
	Name:        "drive",
	Version:     "v3",
	BaseURL:     "https://www.googleapis.com/drive/v3/",
	RootURL:     "https://www.googleapis.com/",
	ServicePath: "drive/v3/",
}

type MethodDescriptor struct {
	Name       string
	HTTPMethod string
	Path       string
	Parameters []ParameterDescriptor
}

type ParameterDescriptor struct {
	Name     string
	Type     string
	Location string
	Required bool
}

type ResourceDescriptor struct {
	Name    string
	Methods map[string]MethodDescriptor
}

var Resources = map[string]ResourceDescriptor{

	"about": About,

	"accessproposals": Accessproposals,

	"approvals": Approvals,

	"apps": Apps,

	"changes": Changes,

	"channels": Channels,

	"comments": Comments,

	"drives": Drives,

	"files": Files,

	"operations": Operations,

	"permissions": Permissions,

	"replies": Replies,

	"revisions": Revisions,

	"teamdrives": Teamdrives,
}

// about
var About = ResourceDescriptor{
	Name: "about",
	Methods: map[string]MethodDescriptor{

		"drive.about.get": {
			Name:       "drive.about.get",
			HTTPMethod: "GET",
			Path:       "about",
			Parameters: []ParameterDescriptor{},
		},
	},
}

// accessproposals
var Accessproposals = ResourceDescriptor{
	Name: "accessproposals",
	Methods: map[string]MethodDescriptor{

		"drive.accessproposals.get": {
			Name:       "drive.accessproposals.get",
			HTTPMethod: "GET",
			Path:       "files/{fileId}/accessproposals/{proposalId}",
			Parameters: []ParameterDescriptor{

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "proposalId", Type: "string", Location: "path", Required: true},
			},
		},

		"drive.accessproposals.list": {
			Name:       "drive.accessproposals.list",
			HTTPMethod: "GET",
			Path:       "files/{fileId}/accessproposals",
			Parameters: []ParameterDescriptor{

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},
			},
		},

		"drive.accessproposals.resolve": {
			Name:       "drive.accessproposals.resolve",
			HTTPMethod: "POST",
			Path:       "files/{fileId}/accessproposals/{proposalId}:resolve",
			Parameters: []ParameterDescriptor{

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "proposalId", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// approvals
var Approvals = ResourceDescriptor{
	Name: "approvals",
	Methods: map[string]MethodDescriptor{

		"drive.approvals.approve": {
			Name:       "drive.approvals.approve",
			HTTPMethod: "POST",
			Path:       "files/{fileId}/approvals/{approvalId}:approve",
			Parameters: []ParameterDescriptor{

				{Name: "approvalId", Type: "string", Location: "path", Required: true},

				{Name: "fileId", Type: "string", Location: "path", Required: true},
			},
		},

		"drive.approvals.cancel": {
			Name:       "drive.approvals.cancel",
			HTTPMethod: "POST",
			Path:       "files/{fileId}/approvals/{approvalId}:cancel",
			Parameters: []ParameterDescriptor{

				{Name: "approvalId", Type: "string", Location: "path", Required: true},

				{Name: "fileId", Type: "string", Location: "path", Required: true},
			},
		},

		"drive.approvals.comment": {
			Name:       "drive.approvals.comment",
			HTTPMethod: "POST",
			Path:       "files/{fileId}/approvals/{approvalId}:comment",
			Parameters: []ParameterDescriptor{

				{Name: "approvalId", Type: "string", Location: "path", Required: true},

				{Name: "fileId", Type: "string", Location: "path", Required: true},
			},
		},

		"drive.approvals.decline": {
			Name:       "drive.approvals.decline",
			HTTPMethod: "POST",
			Path:       "files/{fileId}/approvals/{approvalId}:decline",
			Parameters: []ParameterDescriptor{

				{Name: "approvalId", Type: "string", Location: "path", Required: true},

				{Name: "fileId", Type: "string", Location: "path", Required: true},
			},
		},

		"drive.approvals.get": {
			Name:       "drive.approvals.get",
			HTTPMethod: "GET",
			Path:       "files/{fileId}/approvals/{approvalId}",
			Parameters: []ParameterDescriptor{

				{Name: "approvalId", Type: "string", Location: "path", Required: true},

				{Name: "fileId", Type: "string", Location: "path", Required: true},
			},
		},

		"drive.approvals.list": {
			Name:       "drive.approvals.list",
			HTTPMethod: "GET",
			Path:       "files/{fileId}/approvals",
			Parameters: []ParameterDescriptor{

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},
			},
		},

		"drive.approvals.reassign": {
			Name:       "drive.approvals.reassign",
			HTTPMethod: "POST",
			Path:       "files/{fileId}/approvals/{approvalId}:reassign",
			Parameters: []ParameterDescriptor{

				{Name: "approvalId", Type: "string", Location: "path", Required: true},

				{Name: "fileId", Type: "string", Location: "path", Required: true},
			},
		},

		"drive.approvals.start": {
			Name:       "drive.approvals.start",
			HTTPMethod: "POST",
			Path:       "files/{fileId}/approvals:start",
			Parameters: []ParameterDescriptor{

				{Name: "fileId", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// apps
var Apps = ResourceDescriptor{
	Name: "apps",
	Methods: map[string]MethodDescriptor{

		"drive.apps.get": {
			Name:       "drive.apps.get",
			HTTPMethod: "GET",
			Path:       "apps/{appId}",
			Parameters: []ParameterDescriptor{

				{Name: "appId", Type: "string", Location: "path", Required: true},
			},
		},

		"drive.apps.list": {
			Name:       "drive.apps.list",
			HTTPMethod: "GET",
			Path:       "apps",
			Parameters: []ParameterDescriptor{

				{Name: "appFilterExtensions", Type: "string", Location: "query", Required: false},

				{Name: "appFilterMimeTypes", Type: "string", Location: "query", Required: false},

				{Name: "languageCode", Type: "string", Location: "query", Required: false},
			},
		},
	},
}

// changes
var Changes = ResourceDescriptor{
	Name: "changes",
	Methods: map[string]MethodDescriptor{

		"drive.changes.getStartPageToken": {
			Name:       "drive.changes.getStartPageToken",
			HTTPMethod: "GET",
			Path:       "changes/startPageToken",
			Parameters: []ParameterDescriptor{

				{Name: "driveId", Type: "string", Location: "query", Required: false},

				{Name: "supportsAllDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "supportsTeamDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "teamDriveId", Type: "string", Location: "query", Required: false},
			},
		},

		"drive.changes.list": {
			Name:       "drive.changes.list",
			HTTPMethod: "GET",
			Path:       "changes",
			Parameters: []ParameterDescriptor{

				{Name: "driveId", Type: "string", Location: "query", Required: false},

				{Name: "includeCorpusRemovals", Type: "boolean", Location: "query", Required: false},

				{Name: "includeItemsFromAllDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "includeLabels", Type: "string", Location: "query", Required: false},

				{Name: "includePermissionsForView", Type: "string", Location: "query", Required: false},

				{Name: "includeRemoved", Type: "boolean", Location: "query", Required: false},

				{Name: "includeTeamDriveItems", Type: "boolean", Location: "query", Required: false},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: true},

				{Name: "restrictToMyDrive", Type: "boolean", Location: "query", Required: false},

				{Name: "spaces", Type: "string", Location: "query", Required: false},

				{Name: "supportsAllDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "supportsTeamDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "teamDriveId", Type: "string", Location: "query", Required: false},
			},
		},

		"drive.changes.watch": {
			Name:       "drive.changes.watch",
			HTTPMethod: "POST",
			Path:       "changes/watch",
			Parameters: []ParameterDescriptor{

				{Name: "driveId", Type: "string", Location: "query", Required: false},

				{Name: "includeCorpusRemovals", Type: "boolean", Location: "query", Required: false},

				{Name: "includeItemsFromAllDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "includeLabels", Type: "string", Location: "query", Required: false},

				{Name: "includePermissionsForView", Type: "string", Location: "query", Required: false},

				{Name: "includeRemoved", Type: "boolean", Location: "query", Required: false},

				{Name: "includeTeamDriveItems", Type: "boolean", Location: "query", Required: false},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: true},

				{Name: "restrictToMyDrive", Type: "boolean", Location: "query", Required: false},

				{Name: "spaces", Type: "string", Location: "query", Required: false},

				{Name: "supportsAllDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "supportsTeamDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "teamDriveId", Type: "string", Location: "query", Required: false},
			},
		},
	},
}

// channels
var Channels = ResourceDescriptor{
	Name: "channels",
	Methods: map[string]MethodDescriptor{

		"drive.channels.stop": {
			Name:       "drive.channels.stop",
			HTTPMethod: "POST",
			Path:       "channels/stop",
			Parameters: []ParameterDescriptor{},
		},
	},
}

// comments
var Comments = ResourceDescriptor{
	Name: "comments",
	Methods: map[string]MethodDescriptor{

		"drive.comments.create": {
			Name:       "drive.comments.create",
			HTTPMethod: "POST",
			Path:       "files/{fileId}/comments",
			Parameters: []ParameterDescriptor{

				{Name: "fileId", Type: "string", Location: "path", Required: true},
			},
		},

		"drive.comments.delete": {
			Name:       "drive.comments.delete",
			HTTPMethod: "DELETE",
			Path:       "files/{fileId}/comments/{commentId}",
			Parameters: []ParameterDescriptor{

				{Name: "commentId", Type: "string", Location: "path", Required: true},

				{Name: "fileId", Type: "string", Location: "path", Required: true},
			},
		},

		"drive.comments.get": {
			Name:       "drive.comments.get",
			HTTPMethod: "GET",
			Path:       "files/{fileId}/comments/{commentId}",
			Parameters: []ParameterDescriptor{

				{Name: "commentId", Type: "string", Location: "path", Required: true},

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "includeDeleted", Type: "boolean", Location: "query", Required: false},
			},
		},

		"drive.comments.list": {
			Name:       "drive.comments.list",
			HTTPMethod: "GET",
			Path:       "files/{fileId}/comments",
			Parameters: []ParameterDescriptor{

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "includeDeleted", Type: "boolean", Location: "query", Required: false},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "startModifiedTime", Type: "string", Location: "query", Required: false},
			},
		},

		"drive.comments.update": {
			Name:       "drive.comments.update",
			HTTPMethod: "PATCH",
			Path:       "files/{fileId}/comments/{commentId}",
			Parameters: []ParameterDescriptor{

				{Name: "commentId", Type: "string", Location: "path", Required: true},

				{Name: "fileId", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// drives
var Drives = ResourceDescriptor{
	Name: "drives",
	Methods: map[string]MethodDescriptor{

		"drive.drives.create": {
			Name:       "drive.drives.create",
			HTTPMethod: "POST",
			Path:       "drives",
			Parameters: []ParameterDescriptor{

				{Name: "requestId", Type: "string", Location: "query", Required: true},
			},
		},

		"drive.drives.delete": {
			Name:       "drive.drives.delete",
			HTTPMethod: "DELETE",
			Path:       "drives/{driveId}",
			Parameters: []ParameterDescriptor{

				{Name: "allowItemDeletion", Type: "boolean", Location: "query", Required: false},

				{Name: "driveId", Type: "string", Location: "path", Required: true},

				{Name: "useDomainAdminAccess", Type: "boolean", Location: "query", Required: false},
			},
		},

		"drive.drives.get": {
			Name:       "drive.drives.get",
			HTTPMethod: "GET",
			Path:       "drives/{driveId}",
			Parameters: []ParameterDescriptor{

				{Name: "driveId", Type: "string", Location: "path", Required: true},

				{Name: "useDomainAdminAccess", Type: "boolean", Location: "query", Required: false},
			},
		},

		"drive.drives.hide": {
			Name:       "drive.drives.hide",
			HTTPMethod: "POST",
			Path:       "drives/{driveId}/hide",
			Parameters: []ParameterDescriptor{

				{Name: "driveId", Type: "string", Location: "path", Required: true},
			},
		},

		"drive.drives.list": {
			Name:       "drive.drives.list",
			HTTPMethod: "GET",
			Path:       "drives",
			Parameters: []ParameterDescriptor{

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "q", Type: "string", Location: "query", Required: false},

				{Name: "useDomainAdminAccess", Type: "boolean", Location: "query", Required: false},
			},
		},

		"drive.drives.unhide": {
			Name:       "drive.drives.unhide",
			HTTPMethod: "POST",
			Path:       "drives/{driveId}/unhide",
			Parameters: []ParameterDescriptor{

				{Name: "driveId", Type: "string", Location: "path", Required: true},
			},
		},

		"drive.drives.update": {
			Name:       "drive.drives.update",
			HTTPMethod: "PATCH",
			Path:       "drives/{driveId}",
			Parameters: []ParameterDescriptor{

				{Name: "driveId", Type: "string", Location: "path", Required: true},

				{Name: "useDomainAdminAccess", Type: "boolean", Location: "query", Required: false},
			},
		},
	},
}

// files
var Files = ResourceDescriptor{
	Name: "files",
	Methods: map[string]MethodDescriptor{

		"drive.files.copy": {
			Name:       "drive.files.copy",
			HTTPMethod: "POST",
			Path:       "files/{fileId}/copy",
			Parameters: []ParameterDescriptor{

				{Name: "copyComments", Type: "boolean", Location: "query", Required: false},

				{Name: "enforceSingleParent", Type: "boolean", Location: "query", Required: false},

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "ignoreDefaultVisibility", Type: "boolean", Location: "query", Required: false},

				{Name: "includeLabels", Type: "string", Location: "query", Required: false},

				{Name: "includePermissionsForView", Type: "string", Location: "query", Required: false},

				{Name: "keepRevisionForever", Type: "boolean", Location: "query", Required: false},

				{Name: "ocrLanguage", Type: "string", Location: "query", Required: false},

				{Name: "supportsAllDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "supportsTeamDrives", Type: "boolean", Location: "query", Required: false},
			},
		},

		"drive.files.create": {
			Name:       "drive.files.create",
			HTTPMethod: "POST",
			Path:       "files",
			Parameters: []ParameterDescriptor{

				{Name: "enforceSingleParent", Type: "boolean", Location: "query", Required: false},

				{Name: "ignoreDefaultVisibility", Type: "boolean", Location: "query", Required: false},

				{Name: "includeLabels", Type: "string", Location: "query", Required: false},

				{Name: "includePermissionsForView", Type: "string", Location: "query", Required: false},

				{Name: "keepRevisionForever", Type: "boolean", Location: "query", Required: false},

				{Name: "ocrLanguage", Type: "string", Location: "query", Required: false},

				{Name: "supportsAllDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "supportsTeamDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "useContentAsIndexableText", Type: "boolean", Location: "query", Required: false},
			},
		},

		"drive.files.delete": {
			Name:       "drive.files.delete",
			HTTPMethod: "DELETE",
			Path:       "files/{fileId}",
			Parameters: []ParameterDescriptor{

				{Name: "enforceSingleParent", Type: "boolean", Location: "query", Required: false},

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "supportsAllDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "supportsTeamDrives", Type: "boolean", Location: "query", Required: false},
			},
		},

		"drive.files.download": {
			Name:       "drive.files.download",
			HTTPMethod: "POST",
			Path:       "files/{fileId}/download",
			Parameters: []ParameterDescriptor{

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "mimeType", Type: "string", Location: "query", Required: false},

				{Name: "revisionId", Type: "string", Location: "query", Required: false},
			},
		},

		"drive.files.emptyTrash": {
			Name:       "drive.files.emptyTrash",
			HTTPMethod: "DELETE",
			Path:       "files/trash",
			Parameters: []ParameterDescriptor{

				{Name: "driveId", Type: "string", Location: "query", Required: false},

				{Name: "enforceSingleParent", Type: "boolean", Location: "query", Required: false},
			},
		},

		"drive.files.export": {
			Name:       "drive.files.export",
			HTTPMethod: "GET",
			Path:       "files/{fileId}/export",
			Parameters: []ParameterDescriptor{

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "mimeType", Type: "string", Location: "query", Required: true},
			},
		},

		"drive.files.generateCseToken": {
			Name:       "drive.files.generateCseToken",
			HTTPMethod: "GET",
			Path:       "files/generateCseToken",
			Parameters: []ParameterDescriptor{

				{Name: "fileId", Type: "string", Location: "query", Required: false},

				{Name: "parent", Type: "string", Location: "query", Required: false},
			},
		},

		"drive.files.generateIds": {
			Name:       "drive.files.generateIds",
			HTTPMethod: "GET",
			Path:       "files/generateIds",
			Parameters: []ParameterDescriptor{

				{Name: "count", Type: "integer", Location: "query", Required: false},

				{Name: "space", Type: "string", Location: "query", Required: false},

				{Name: "type", Type: "string", Location: "query", Required: false},
			},
		},

		"drive.files.get": {
			Name:       "drive.files.get",
			HTTPMethod: "GET",
			Path:       "files/{fileId}",
			Parameters: []ParameterDescriptor{

				{Name: "acknowledgeAbuse", Type: "boolean", Location: "query", Required: false},

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "includeLabels", Type: "string", Location: "query", Required: false},

				{Name: "includePermissionsForView", Type: "string", Location: "query", Required: false},

				{Name: "supportsAllDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "supportsTeamDrives", Type: "boolean", Location: "query", Required: false},
			},
		},

		"drive.files.list": {
			Name:       "drive.files.list",
			HTTPMethod: "GET",
			Path:       "files",
			Parameters: []ParameterDescriptor{

				{Name: "corpora", Type: "string", Location: "query", Required: false},

				{Name: "corpus", Type: "string", Location: "query", Required: false},

				{Name: "driveId", Type: "string", Location: "query", Required: false},

				{Name: "includeItemsFromAllDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "includeLabels", Type: "string", Location: "query", Required: false},

				{Name: "includePermissionsForView", Type: "string", Location: "query", Required: false},

				{Name: "includeTeamDriveItems", Type: "boolean", Location: "query", Required: false},

				{Name: "orderBy", Type: "string", Location: "query", Required: false},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "q", Type: "string", Location: "query", Required: false},

				{Name: "spaces", Type: "string", Location: "query", Required: false},

				{Name: "supportsAllDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "supportsTeamDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "teamDriveId", Type: "string", Location: "query", Required: false},
			},
		},

		"drive.files.listLabels": {
			Name:       "drive.files.listLabels",
			HTTPMethod: "GET",
			Path:       "files/{fileId}/listLabels",
			Parameters: []ParameterDescriptor{

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "maxResults", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},
			},
		},

		"drive.files.modifyLabels": {
			Name:       "drive.files.modifyLabels",
			HTTPMethod: "POST",
			Path:       "files/{fileId}/modifyLabels",
			Parameters: []ParameterDescriptor{

				{Name: "fileId", Type: "string", Location: "path", Required: true},
			},
		},

		"drive.files.update": {
			Name:       "drive.files.update",
			HTTPMethod: "PATCH",
			Path:       "files/{fileId}",
			Parameters: []ParameterDescriptor{

				{Name: "addParents", Type: "string", Location: "query", Required: false},

				{Name: "enforceSingleParent", Type: "boolean", Location: "query", Required: false},

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "includeLabels", Type: "string", Location: "query", Required: false},

				{Name: "includePermissionsForView", Type: "string", Location: "query", Required: false},

				{Name: "keepRevisionForever", Type: "boolean", Location: "query", Required: false},

				{Name: "ocrLanguage", Type: "string", Location: "query", Required: false},

				{Name: "removeParents", Type: "string", Location: "query", Required: false},

				{Name: "supportsAllDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "supportsTeamDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "useContentAsIndexableText", Type: "boolean", Location: "query", Required: false},
			},
		},

		"drive.files.watch": {
			Name:       "drive.files.watch",
			HTTPMethod: "POST",
			Path:       "files/{fileId}/watch",
			Parameters: []ParameterDescriptor{

				{Name: "acknowledgeAbuse", Type: "boolean", Location: "query", Required: false},

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "includeLabels", Type: "string", Location: "query", Required: false},

				{Name: "includePermissionsForView", Type: "string", Location: "query", Required: false},

				{Name: "supportsAllDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "supportsTeamDrives", Type: "boolean", Location: "query", Required: false},
			},
		},
	},
}

// operations
var Operations = ResourceDescriptor{
	Name: "operations",
	Methods: map[string]MethodDescriptor{

		"drive.operations.get": {
			Name:       "drive.operations.get",
			HTTPMethod: "GET",
			Path:       "operations/{name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// permissions
var Permissions = ResourceDescriptor{
	Name: "permissions",
	Methods: map[string]MethodDescriptor{

		"drive.permissions.create": {
			Name:       "drive.permissions.create",
			HTTPMethod: "POST",
			Path:       "files/{fileId}/permissions",
			Parameters: []ParameterDescriptor{

				{Name: "emailMessage", Type: "string", Location: "query", Required: false},

				{Name: "enforceExpansiveAccess", Type: "boolean", Location: "query", Required: false},

				{Name: "enforceSingleParent", Type: "boolean", Location: "query", Required: false},

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "moveToNewOwnersRoot", Type: "boolean", Location: "query", Required: false},

				{Name: "sendNotificationEmail", Type: "boolean", Location: "query", Required: false},

				{Name: "supportsAllDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "supportsTeamDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "transferOwnership", Type: "boolean", Location: "query", Required: false},

				{Name: "useDomainAdminAccess", Type: "boolean", Location: "query", Required: false},
			},
		},

		"drive.permissions.delete": {
			Name:       "drive.permissions.delete",
			HTTPMethod: "DELETE",
			Path:       "files/{fileId}/permissions/{permissionId}",
			Parameters: []ParameterDescriptor{

				{Name: "enforceExpansiveAccess", Type: "boolean", Location: "query", Required: false},

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "permissionId", Type: "string", Location: "path", Required: true},

				{Name: "supportsAllDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "supportsTeamDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "useDomainAdminAccess", Type: "boolean", Location: "query", Required: false},
			},
		},

		"drive.permissions.get": {
			Name:       "drive.permissions.get",
			HTTPMethod: "GET",
			Path:       "files/{fileId}/permissions/{permissionId}",
			Parameters: []ParameterDescriptor{

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "permissionId", Type: "string", Location: "path", Required: true},

				{Name: "supportsAllDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "supportsTeamDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "useDomainAdminAccess", Type: "boolean", Location: "query", Required: false},
			},
		},

		"drive.permissions.list": {
			Name:       "drive.permissions.list",
			HTTPMethod: "GET",
			Path:       "files/{fileId}/permissions",
			Parameters: []ParameterDescriptor{

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "includePermissionsForView", Type: "string", Location: "query", Required: false},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "supportsAllDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "supportsTeamDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "useDomainAdminAccess", Type: "boolean", Location: "query", Required: false},
			},
		},

		"drive.permissions.update": {
			Name:       "drive.permissions.update",
			HTTPMethod: "PATCH",
			Path:       "files/{fileId}/permissions/{permissionId}",
			Parameters: []ParameterDescriptor{

				{Name: "enforceExpansiveAccess", Type: "boolean", Location: "query", Required: false},

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "permissionId", Type: "string", Location: "path", Required: true},

				{Name: "removeExpiration", Type: "boolean", Location: "query", Required: false},

				{Name: "supportsAllDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "supportsTeamDrives", Type: "boolean", Location: "query", Required: false},

				{Name: "transferOwnership", Type: "boolean", Location: "query", Required: false},

				{Name: "useDomainAdminAccess", Type: "boolean", Location: "query", Required: false},
			},
		},
	},
}

// replies
var Replies = ResourceDescriptor{
	Name: "replies",
	Methods: map[string]MethodDescriptor{

		"drive.replies.create": {
			Name:       "drive.replies.create",
			HTTPMethod: "POST",
			Path:       "files/{fileId}/comments/{commentId}/replies",
			Parameters: []ParameterDescriptor{

				{Name: "commentId", Type: "string", Location: "path", Required: true},

				{Name: "fileId", Type: "string", Location: "path", Required: true},
			},
		},

		"drive.replies.delete": {
			Name:       "drive.replies.delete",
			HTTPMethod: "DELETE",
			Path:       "files/{fileId}/comments/{commentId}/replies/{replyId}",
			Parameters: []ParameterDescriptor{

				{Name: "commentId", Type: "string", Location: "path", Required: true},

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "replyId", Type: "string", Location: "path", Required: true},
			},
		},

		"drive.replies.get": {
			Name:       "drive.replies.get",
			HTTPMethod: "GET",
			Path:       "files/{fileId}/comments/{commentId}/replies/{replyId}",
			Parameters: []ParameterDescriptor{

				{Name: "commentId", Type: "string", Location: "path", Required: true},

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "includeDeleted", Type: "boolean", Location: "query", Required: false},

				{Name: "replyId", Type: "string", Location: "path", Required: true},
			},
		},

		"drive.replies.list": {
			Name:       "drive.replies.list",
			HTTPMethod: "GET",
			Path:       "files/{fileId}/comments/{commentId}/replies",
			Parameters: []ParameterDescriptor{

				{Name: "commentId", Type: "string", Location: "path", Required: true},

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "includeDeleted", Type: "boolean", Location: "query", Required: false},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},
			},
		},

		"drive.replies.update": {
			Name:       "drive.replies.update",
			HTTPMethod: "PATCH",
			Path:       "files/{fileId}/comments/{commentId}/replies/{replyId}",
			Parameters: []ParameterDescriptor{

				{Name: "commentId", Type: "string", Location: "path", Required: true},

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "replyId", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// revisions
var Revisions = ResourceDescriptor{
	Name: "revisions",
	Methods: map[string]MethodDescriptor{

		"drive.revisions.delete": {
			Name:       "drive.revisions.delete",
			HTTPMethod: "DELETE",
			Path:       "files/{fileId}/revisions/{revisionId}",
			Parameters: []ParameterDescriptor{

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "revisionId", Type: "string", Location: "path", Required: true},
			},
		},

		"drive.revisions.get": {
			Name:       "drive.revisions.get",
			HTTPMethod: "GET",
			Path:       "files/{fileId}/revisions/{revisionId}",
			Parameters: []ParameterDescriptor{

				{Name: "acknowledgeAbuse", Type: "boolean", Location: "query", Required: false},

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "revisionId", Type: "string", Location: "path", Required: true},
			},
		},

		"drive.revisions.list": {
			Name:       "drive.revisions.list",
			HTTPMethod: "GET",
			Path:       "files/{fileId}/revisions",
			Parameters: []ParameterDescriptor{

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},
			},
		},

		"drive.revisions.update": {
			Name:       "drive.revisions.update",
			HTTPMethod: "PATCH",
			Path:       "files/{fileId}/revisions/{revisionId}",
			Parameters: []ParameterDescriptor{

				{Name: "fileId", Type: "string", Location: "path", Required: true},

				{Name: "revisionId", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// teamdrives
var Teamdrives = ResourceDescriptor{
	Name: "teamdrives",
	Methods: map[string]MethodDescriptor{

		"drive.teamdrives.create": {
			Name:       "drive.teamdrives.create",
			HTTPMethod: "POST",
			Path:       "teamdrives",
			Parameters: []ParameterDescriptor{

				{Name: "requestId", Type: "string", Location: "query", Required: true},
			},
		},

		"drive.teamdrives.delete": {
			Name:       "drive.teamdrives.delete",
			HTTPMethod: "DELETE",
			Path:       "teamdrives/{teamDriveId}",
			Parameters: []ParameterDescriptor{

				{Name: "teamDriveId", Type: "string", Location: "path", Required: true},
			},
		},

		"drive.teamdrives.get": {
			Name:       "drive.teamdrives.get",
			HTTPMethod: "GET",
			Path:       "teamdrives/{teamDriveId}",
			Parameters: []ParameterDescriptor{

				{Name: "teamDriveId", Type: "string", Location: "path", Required: true},

				{Name: "useDomainAdminAccess", Type: "boolean", Location: "query", Required: false},
			},
		},

		"drive.teamdrives.list": {
			Name:       "drive.teamdrives.list",
			HTTPMethod: "GET",
			Path:       "teamdrives",
			Parameters: []ParameterDescriptor{

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "q", Type: "string", Location: "query", Required: false},

				{Name: "useDomainAdminAccess", Type: "boolean", Location: "query", Required: false},
			},
		},

		"drive.teamdrives.update": {
			Name:       "drive.teamdrives.update",
			HTTPMethod: "PATCH",
			Path:       "teamdrives/{teamDriveId}",
			Parameters: []ParameterDescriptor{

				{Name: "teamDriveId", Type: "string", Location: "path", Required: true},

				{Name: "useDomainAdminAccess", Type: "boolean", Location: "query", Required: false},
			},
		},
	},
}
