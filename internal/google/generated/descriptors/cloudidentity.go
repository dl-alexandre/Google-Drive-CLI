// Cloud Identity API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package descriptors

// ServiceDescriptor contains metadata for the cloudidentity API
type ServiceDescriptor struct {
	Name        string
	Version     string
	BaseURL     string
	RootURL     string
	ServicePath string
}

// Service returns the service descriptor for cloudidentity
var Service = ServiceDescriptor{
	Name:        "cloudidentity",
	Version:     "v1",
	BaseURL:     "https://cloudidentity.googleapis.com/",
	RootURL:     "https://cloudidentity.googleapis.com/",
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

	"allowlistedDomains": AllowlistedDomains,

	"customers": Customers,

	"devices": Devices,

	"groups": Groups,

	"inboundOidcSsoProfiles": InboundOidcSsoProfiles,

	"inboundSamlSsoProfiles": InboundSamlSsoProfiles,

	"inboundSsoAssignments": InboundSsoAssignments,

	"policies": Policies,
}

// allowlistedDomains
var AllowlistedDomains = ResourceDescriptor{
	Name: "allowlistedDomains",
	Methods: map[string]MethodDescriptor{

		"cloudidentity.allowlistedDomains.create": {
			Name:       "cloudidentity.allowlistedDomains.create",
			HTTPMethod: "POST",
			Path:       "v1/allowlistedDomains",
			Parameters: []ParameterDescriptor{},
		},

		"cloudidentity.allowlistedDomains.delete": {
			Name:       "cloudidentity.allowlistedDomains.delete",
			HTTPMethod: "DELETE",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"cloudidentity.allowlistedDomains.get": {
			Name:       "cloudidentity.allowlistedDomains.get",
			HTTPMethod: "GET",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"cloudidentity.allowlistedDomains.list": {
			Name:       "cloudidentity.allowlistedDomains.list",
			HTTPMethod: "GET",
			Path:       "v1/allowlistedDomains",
			Parameters: []ParameterDescriptor{

				{Name: "filter", Type: "string", Location: "query", Required: false},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},
			},
		},
	},
}

// customers
var Customers = ResourceDescriptor{
	Name:    "customers",
	Methods: map[string]MethodDescriptor{},
}

// devices
var Devices = ResourceDescriptor{
	Name: "devices",
	Methods: map[string]MethodDescriptor{

		"cloudidentity.devices.cancelWipe": {
			Name:       "cloudidentity.devices.cancelWipe",
			HTTPMethod: "POST",
			Path:       "v1/{+name}:cancelWipe",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"cloudidentity.devices.create": {
			Name:       "cloudidentity.devices.create",
			HTTPMethod: "POST",
			Path:       "v1/devices",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "query", Required: false},
			},
		},

		"cloudidentity.devices.delete": {
			Name:       "cloudidentity.devices.delete",
			HTTPMethod: "DELETE",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "query", Required: false},

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"cloudidentity.devices.get": {
			Name:       "cloudidentity.devices.get",
			HTTPMethod: "GET",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "query", Required: false},

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"cloudidentity.devices.list": {
			Name:       "cloudidentity.devices.list",
			HTTPMethod: "GET",
			Path:       "v1/devices",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "query", Required: false},

				{Name: "filter", Type: "string", Location: "query", Required: false},

				{Name: "orderBy", Type: "string", Location: "query", Required: false},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "view", Type: "string", Location: "query", Required: false},
			},
		},

		"cloudidentity.devices.wipe": {
			Name:       "cloudidentity.devices.wipe",
			HTTPMethod: "POST",
			Path:       "v1/{+name}:wipe",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// groups
var Groups = ResourceDescriptor{
	Name: "groups",
	Methods: map[string]MethodDescriptor{

		"cloudidentity.groups.create": {
			Name:       "cloudidentity.groups.create",
			HTTPMethod: "POST",
			Path:       "v1/groups",
			Parameters: []ParameterDescriptor{

				{Name: "initialGroupConfig", Type: "string", Location: "query", Required: false},
			},
		},

		"cloudidentity.groups.delete": {
			Name:       "cloudidentity.groups.delete",
			HTTPMethod: "DELETE",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"cloudidentity.groups.get": {
			Name:       "cloudidentity.groups.get",
			HTTPMethod: "GET",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"cloudidentity.groups.getSecuritySettings": {
			Name:       "cloudidentity.groups.getSecuritySettings",
			HTTPMethod: "GET",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},

				{Name: "readMask", Type: "string", Location: "query", Required: false},
			},
		},

		"cloudidentity.groups.list": {
			Name:       "cloudidentity.groups.list",
			HTTPMethod: "GET",
			Path:       "v1/groups",
			Parameters: []ParameterDescriptor{

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "parent", Type: "string", Location: "query", Required: false},

				{Name: "view", Type: "string", Location: "query", Required: false},
			},
		},

		"cloudidentity.groups.lookup": {
			Name:       "cloudidentity.groups.lookup",
			HTTPMethod: "GET",
			Path:       "v1/groups:lookup",
			Parameters: []ParameterDescriptor{

				{Name: "groupKey.id", Type: "string", Location: "query", Required: false},

				{Name: "groupKey.namespace", Type: "string", Location: "query", Required: false},
			},
		},

		"cloudidentity.groups.patch": {
			Name:       "cloudidentity.groups.patch",
			HTTPMethod: "PATCH",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},

				{Name: "updateMask", Type: "string", Location: "query", Required: false},
			},
		},

		"cloudidentity.groups.search": {
			Name:       "cloudidentity.groups.search",
			HTTPMethod: "GET",
			Path:       "v1/groups:search",
			Parameters: []ParameterDescriptor{

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "query", Type: "string", Location: "query", Required: false},

				{Name: "view", Type: "string", Location: "query", Required: false},
			},
		},

		"cloudidentity.groups.updateSecuritySettings": {
			Name:       "cloudidentity.groups.updateSecuritySettings",
			HTTPMethod: "PATCH",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},

				{Name: "updateMask", Type: "string", Location: "query", Required: false},
			},
		},
	},
}

// inboundOidcSsoProfiles
var InboundOidcSsoProfiles = ResourceDescriptor{
	Name: "inboundOidcSsoProfiles",
	Methods: map[string]MethodDescriptor{

		"cloudidentity.inboundOidcSsoProfiles.create": {
			Name:       "cloudidentity.inboundOidcSsoProfiles.create",
			HTTPMethod: "POST",
			Path:       "v1/inboundOidcSsoProfiles",
			Parameters: []ParameterDescriptor{},
		},

		"cloudidentity.inboundOidcSsoProfiles.delete": {
			Name:       "cloudidentity.inboundOidcSsoProfiles.delete",
			HTTPMethod: "DELETE",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"cloudidentity.inboundOidcSsoProfiles.get": {
			Name:       "cloudidentity.inboundOidcSsoProfiles.get",
			HTTPMethod: "GET",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"cloudidentity.inboundOidcSsoProfiles.list": {
			Name:       "cloudidentity.inboundOidcSsoProfiles.list",
			HTTPMethod: "GET",
			Path:       "v1/inboundOidcSsoProfiles",
			Parameters: []ParameterDescriptor{

				{Name: "filter", Type: "string", Location: "query", Required: false},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},
			},
		},

		"cloudidentity.inboundOidcSsoProfiles.patch": {
			Name:       "cloudidentity.inboundOidcSsoProfiles.patch",
			HTTPMethod: "PATCH",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},

				{Name: "updateMask", Type: "string", Location: "query", Required: false},
			},
		},
	},
}

// inboundSamlSsoProfiles
var InboundSamlSsoProfiles = ResourceDescriptor{
	Name: "inboundSamlSsoProfiles",
	Methods: map[string]MethodDescriptor{

		"cloudidentity.inboundSamlSsoProfiles.create": {
			Name:       "cloudidentity.inboundSamlSsoProfiles.create",
			HTTPMethod: "POST",
			Path:       "v1/inboundSamlSsoProfiles",
			Parameters: []ParameterDescriptor{},
		},

		"cloudidentity.inboundSamlSsoProfiles.delete": {
			Name:       "cloudidentity.inboundSamlSsoProfiles.delete",
			HTTPMethod: "DELETE",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"cloudidentity.inboundSamlSsoProfiles.get": {
			Name:       "cloudidentity.inboundSamlSsoProfiles.get",
			HTTPMethod: "GET",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"cloudidentity.inboundSamlSsoProfiles.list": {
			Name:       "cloudidentity.inboundSamlSsoProfiles.list",
			HTTPMethod: "GET",
			Path:       "v1/inboundSamlSsoProfiles",
			Parameters: []ParameterDescriptor{

				{Name: "filter", Type: "string", Location: "query", Required: false},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},
			},
		},

		"cloudidentity.inboundSamlSsoProfiles.patch": {
			Name:       "cloudidentity.inboundSamlSsoProfiles.patch",
			HTTPMethod: "PATCH",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},

				{Name: "updateMask", Type: "string", Location: "query", Required: false},
			},
		},
	},
}

// inboundSsoAssignments
var InboundSsoAssignments = ResourceDescriptor{
	Name: "inboundSsoAssignments",
	Methods: map[string]MethodDescriptor{

		"cloudidentity.inboundSsoAssignments.create": {
			Name:       "cloudidentity.inboundSsoAssignments.create",
			HTTPMethod: "POST",
			Path:       "v1/inboundSsoAssignments",
			Parameters: []ParameterDescriptor{},
		},

		"cloudidentity.inboundSsoAssignments.delete": {
			Name:       "cloudidentity.inboundSsoAssignments.delete",
			HTTPMethod: "DELETE",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"cloudidentity.inboundSsoAssignments.get": {
			Name:       "cloudidentity.inboundSsoAssignments.get",
			HTTPMethod: "GET",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"cloudidentity.inboundSsoAssignments.list": {
			Name:       "cloudidentity.inboundSsoAssignments.list",
			HTTPMethod: "GET",
			Path:       "v1/inboundSsoAssignments",
			Parameters: []ParameterDescriptor{

				{Name: "filter", Type: "string", Location: "query", Required: false},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},
			},
		},

		"cloudidentity.inboundSsoAssignments.patch": {
			Name:       "cloudidentity.inboundSsoAssignments.patch",
			HTTPMethod: "PATCH",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},

				{Name: "updateMask", Type: "string", Location: "query", Required: false},
			},
		},
	},
}

// policies
var Policies = ResourceDescriptor{
	Name: "policies",
	Methods: map[string]MethodDescriptor{

		"cloudidentity.policies.create": {
			Name:       "cloudidentity.policies.create",
			HTTPMethod: "POST",
			Path:       "v1/policies",
			Parameters: []ParameterDescriptor{},
		},

		"cloudidentity.policies.delete": {
			Name:       "cloudidentity.policies.delete",
			HTTPMethod: "DELETE",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"cloudidentity.policies.get": {
			Name:       "cloudidentity.policies.get",
			HTTPMethod: "GET",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},

		"cloudidentity.policies.list": {
			Name:       "cloudidentity.policies.list",
			HTTPMethod: "GET",
			Path:       "v1/policies",
			Parameters: []ParameterDescriptor{

				{Name: "filter", Type: "string", Location: "query", Required: false},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},
			},
		},

		"cloudidentity.policies.patch": {
			Name:       "cloudidentity.policies.patch",
			HTTPMethod: "PATCH",
			Path:       "v1/{+name}",
			Parameters: []ParameterDescriptor{

				{Name: "name", Type: "string", Location: "path", Required: true},
			},
		},
	},
}
