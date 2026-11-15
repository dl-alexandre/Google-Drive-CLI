// Drive Labels API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package descriptors

// ServiceDescriptor contains metadata for the drivelabels API
type ServiceDescriptor struct {
	Name        string
	Version     string
	BaseURL     string
	RootURL     string
	ServicePath string
}

// Service returns the service descriptor for drivelabels
var Service = ServiceDescriptor{
	Name:        "drivelabels",
	Version:     "v2",
	BaseURL:     "https://drivelabels.googleapis.com/",
	RootURL:     "https://drivelabels.googleapis.com/",
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

	"labels": Labels,

	"limits": Limits,

	"users": Users,
}

// labels
var Labels = ResourceDescriptor{
	Name: "labels",
	Methods: map[string]MethodDescriptor{

		"drivelabels.labels.create": {
			Name:       "drivelabels.labels.create",
			HTTPMethod: "POST",
			Path:       "v2/labels",
			Parameters: []ParameterDescriptor{

				{Name: "languageCode", Type: "string", Location: "query", Required: false},

				{Name: "useAdminAccess", Type: "boolean", Location: "query", Required: false},
			},
		},

		"drivelabels.labels.delete": {
			Name:       "drivelabels.labels.delete",
			HTTPMethod: "DELETE",
			Path:       "v2/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},

				{Name: "useAdminAccess", Type: "boolean", Location: "query", Required: false},

				{Name: "writeControl.requiredRevisionId", Type: "string", Location: "query", Required: false},
			},
		},

		"drivelabels.labels.delta": {
			Name:       "drivelabels.labels.delta",
			HTTPMethod: "POST",
			Path:       "v2/{+name}:delta",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"drivelabels.labels.disable": {
			Name:       "drivelabels.labels.disable",
			HTTPMethod: "POST",
			Path:       "v2/{+name}:disable",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"drivelabels.labels.enable": {
			Name:       "drivelabels.labels.enable",
			HTTPMethod: "POST",
			Path:       "v2/{+name}:enable",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"drivelabels.labels.get": {
			Name:       "drivelabels.labels.get",
			HTTPMethod: "GET",
			Path:       "v2/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "languageCode", Type: "string", Location: "query", Required: false},

				{Name: "name", Type: "string", Location: "path", Required: true},

				{Name: "useAdminAccess", Type: "boolean", Location: "query", Required: false},

				{Name: "view", Type: "string", Location: "query", Required: false},
			},
		},

		"drivelabels.labels.list": {
			Name:       "drivelabels.labels.list",
			HTTPMethod: "GET",
			Path:       "v2/labels",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "query", Required: false},

				{Name: "languageCode", Type: "string", Location: "query", Required: false},

				{Name: "minimumRole", Type: "string", Location: "query", Required: false},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "publishedOnly", Type: "boolean", Location: "query", Required: false},

				{Name: "useAdminAccess", Type: "boolean", Location: "query", Required: false},

				{Name: "view", Type: "string", Location: "query", Required: false},
			},
		},

		"drivelabels.labels.publish": {
			Name:       "drivelabels.labels.publish",
			HTTPMethod: "POST",
			Path:       "v2/{+name}:publish",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"drivelabels.labels.updateLabelCopyMode": {
			Name:       "drivelabels.labels.updateLabelCopyMode",
			HTTPMethod: "POST",
			Path:       "v2/{+name}:updateLabelCopyMode",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"drivelabels.labels.updateLabelEnabledAppSettings": {
			Name:       "drivelabels.labels.updateLabelEnabledAppSettings",
			HTTPMethod: "POST",
			Path:       "v2/{+name}:updateLabelEnabledAppSettings",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"drivelabels.labels.updatePermissions": {
			Name:       "drivelabels.labels.updatePermissions",
			HTTPMethod: "PATCH",
			Path:       "v2/{+parent}/permissions",
			Parameters: []ParameterDescriptor{

				{Name: "parent", Type: "string", Location: "path", Required: true},

				{Name: "useAdminAccess", Type: "boolean", Location: "query", Required: false},
			},
		},
	},
}

// limits
var Limits = ResourceDescriptor{
	Name: "limits",
	Methods: map[string]MethodDescriptor{

		"drivelabels.limits.getLabel": {
			Name:       "drivelabels.limits.getLabel",
			HTTPMethod: "GET",
			Path:       "v2/limits/label",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "query", Required: false},
			},
		},
	},
}

// users
var Users = ResourceDescriptor{
	Name: "users",
	Methods: map[string]MethodDescriptor{

		"drivelabels.users.getCapabilities": {
			Name:       "drivelabels.users.getCapabilities",
			HTTPMethod: "GET",
			Path:       "v2/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "query", Required: false},

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},
	},
}
