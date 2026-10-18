// Google Chat API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package descriptors

// ServiceDescriptor contains metadata for the chat API
type ServiceDescriptor struct {
	Name        string
	Version     string
	BaseURL     string
	RootURL     string
	ServicePath string
}

// Service returns the service descriptor for chat
var Service = ServiceDescriptor{
	Name:        "chat",
	Version:     "v1",
	BaseURL:     "https://chat.googleapis.com/",
	RootURL:     "https://chat.googleapis.com/",
	ServicePath: "",
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

	"customEmojis": CustomEmojis,

	"media": Media,

	"spaces": Spaces,

	"users": Users,
}

// customEmojis
var CustomEmojis = ResourceDescriptor{
	Name: "customEmojis",
	Methods: map[string]MethodDescriptor{

		"chat.customEmojis.create": {
			Name:       "chat.customEmojis.create",
			HTTPMethod: "POST",
			Path:       "v1/customEmojis",
			Parameters: []ParameterDescriptor{},
		},

		"chat.customEmojis.delete": {
			Name:       "chat.customEmojis.delete",
			HTTPMethod: "DELETE",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"chat.customEmojis.get": {
			Name:       "chat.customEmojis.get",
			HTTPMethod: "GET",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"chat.customEmojis.list": {
			Name:       "chat.customEmojis.list",
			HTTPMethod: "GET",
			Path:       "v1/customEmojis",
			Parameters: []ParameterDescriptor{

				{Name: "filter", Type: "string", Location: "query", Required: false},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},
			},
		},
	},
}

// media
var Media = ResourceDescriptor{
	Name: "media",
	Methods: map[string]MethodDescriptor{

		"chat.media.download": {
			Name:       "chat.media.download",
			HTTPMethod: "GET",
			Path:       "v1/media/{+resourceName}",
			Parameters: []ParameterDescriptor{

				{Name: "resourceName", Type: "string", Location: "path", Required: true},
			},
		},

		"chat.media.upload": {
			Name:       "chat.media.upload",
			HTTPMethod: "POST",
			Path:       "v1/{+parent}/attachments:upload",
			Parameters: []ParameterDescriptor{

				{Name: "parent", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// spaces
var Spaces = ResourceDescriptor{
	Name: "spaces",
	Methods: map[string]MethodDescriptor{

		"chat.spaces.completeImport": {
			Name:       "chat.spaces.completeImport",
			HTTPMethod: "POST",
			Path:       "v1/{+name}:completeImport",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"chat.spaces.create": {
			Name:       "chat.spaces.create",
			HTTPMethod: "POST",
			Path:       "v1/spaces",
			Parameters: []ParameterDescriptor{

				{Name: "requestId", Type: "string", Location: "query", Required: false},
			},
		},

		"chat.spaces.delete": {
			Name:       "chat.spaces.delete",
			HTTPMethod: "DELETE",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},

				{Name: "useAdminAccess", Type: "boolean", Location: "query", Required: false},
			},
		},

		"chat.spaces.findDirectMessage": {
			Name:       "chat.spaces.findDirectMessage",
			HTTPMethod: "GET",
			Path:       "v1/spaces:findDirectMessage",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "query", Required: false},
			},
		},

		"chat.spaces.findGroupChats": {
			Name:       "chat.spaces.findGroupChats",
			HTTPMethod: "GET",
			Path:       "v1/spaces:findGroupChats",
			Parameters: []ParameterDescriptor{

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "spaceView", Type: "string", Location: "query", Required: false},

				{Name: "users", Type: "string", Location: "query", Required: false},
			},
		},

		"chat.spaces.get": {
			Name:       "chat.spaces.get",
			HTTPMethod: "GET",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},

				{Name: "useAdminAccess", Type: "boolean", Location: "query", Required: false},
			},
		},

		"chat.spaces.list": {
			Name:       "chat.spaces.list",
			HTTPMethod: "GET",
			Path:       "v1/spaces",
			Parameters: []ParameterDescriptor{

				{Name: "filter", Type: "string", Location: "query", Required: false},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},
			},
		},

		"chat.spaces.patch": {
			Name:       "chat.spaces.patch",
			HTTPMethod: "PATCH",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},

				{Name: "updateMask", Type: "string", Location: "query", Required: false},

				{Name: "useAdminAccess", Type: "boolean", Location: "query", Required: false},
			},
		},

		"chat.spaces.search": {
			Name:       "chat.spaces.search",
			HTTPMethod: "GET",
			Path:       "v1/spaces:search",
			Parameters: []ParameterDescriptor{

				{Name: "orderBy", Type: "string", Location: "query", Required: false},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "query", Type: "string", Location: "query", Required: false},

				{Name: "useAdminAccess", Type: "boolean", Location: "query", Required: false},
			},
		},

		"chat.spaces.setup": {
			Name:       "chat.spaces.setup",
			HTTPMethod: "POST",
			Path:       "v1/spaces:setup",
			Parameters: []ParameterDescriptor{},
		},
	},
}

// users
var Users = ResourceDescriptor{
	Name:    "users",
	Methods: map[string]MethodDescriptor{},
}
