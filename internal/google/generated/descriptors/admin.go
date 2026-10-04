// Admin SDK API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package descriptors

// ServiceDescriptor contains metadata for the admin API
type ServiceDescriptor struct {
	Name        string
	Version     string
	BaseURL     string
	RootURL     string
	ServicePath string
}

// Service returns the service descriptor for admin
var Service = ServiceDescriptor{
	Name:        "admin",
	Version:     "directory_v1",
	BaseURL:     "https://admin.googleapis.com/",
	RootURL:     "https://admin.googleapis.com/",
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

	"asps": Asps,

	"channels": Channels,

	"chromeosdevices": Chromeosdevices,

	"customer": Customer,

	"customers": Customers,

	"domainAliases": DomainAliases,

	"domains": Domains,

	"groups": Groups,

	"members": Members,

	"mobiledevices": Mobiledevices,

	"orgunits": Orgunits,

	"privileges": Privileges,

	"resources": Resources,

	"roleAssignments": RoleAssignments,

	"roles": Roles,

	"schemas": Schemas,

	"tokens": Tokens,

	"twoStepVerification": TwoStepVerification,

	"users": Users,

	"verificationCodes": VerificationCodes,
}

// asps
var Asps = ResourceDescriptor{
	Name: "asps",
	Methods: map[string]MethodDescriptor{

		"directory.asps.delete": {
			Name:       "directory.asps.delete",
			HTTPMethod: "DELETE",
			Path:       "admin/directory/v1/users/{userKey}/asps/{codeId}",
			Parameters: []ParameterDescriptor{

				{Name: "codeId", Type: "integer", Location: "path", Required: true},

				{Name: "userKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.asps.get": {
			Name:       "directory.asps.get",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/users/{userKey}/asps/{codeId}",
			Parameters: []ParameterDescriptor{

				{Name: "codeId", Type: "integer", Location: "path", Required: true},

				{Name: "userKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.asps.list": {
			Name:       "directory.asps.list",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/users/{userKey}/asps",
			Parameters: []ParameterDescriptor{

				{Name: "userKey", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// channels
var Channels = ResourceDescriptor{
	Name: "channels",
	Methods: map[string]MethodDescriptor{

		"admin.channels.stop": {
			Name:       "admin.channels.stop",
			HTTPMethod: "POST",
			Path:       "admin/directory_v1/channels/stop",
			Parameters: []ParameterDescriptor{},
		},
	},
}

// chromeosdevices
var Chromeosdevices = ResourceDescriptor{
	Name: "chromeosdevices",
	Methods: map[string]MethodDescriptor{

		"directory.chromeosdevices.action": {
			Name:       "directory.chromeosdevices.action",
			HTTPMethod: "POST",
			Path:       "admin/directory/v1/customer/{customerId}/devices/chromeos/{resourceId}/action",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},

				{Name: "resourceId", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.chromeosdevices.get": {
			Name:       "directory.chromeosdevices.get",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/customer/{customerId}/devices/chromeos/{deviceId}",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},

				{Name: "deviceId", Type: "string", Location: "path", Required: true},

				{Name: "projection", Type: "string", Location: "query", Required: false},
			},
		},

		"directory.chromeosdevices.list": {
			Name:       "directory.chromeosdevices.list",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/customer/{customerId}/devices/chromeos",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},

				{Name: "includeChildOrgunits", Type: "boolean", Location: "query", Required: false},

				{Name: "maxResults", Type: "integer", Location: "query", Required: false},

				{Name: "orderBy", Type: "string", Location: "query", Required: false},

				{Name: "orgUnitPath", Type: "string", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "projection", Type: "string", Location: "query", Required: false},

				{Name: "query", Type: "string", Location: "query", Required: false},

				{Name: "sortOrder", Type: "string", Location: "query", Required: false},
			},
		},

		"directory.chromeosdevices.moveDevicesToOu": {
			Name:       "directory.chromeosdevices.moveDevicesToOu",
			HTTPMethod: "POST",
			Path:       "admin/directory/v1/customer/{customerId}/devices/chromeos/moveDevicesToOu",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},

				{Name: "orgUnitPath", Type: "string", Location: "query", Required: true},
			},
		},

		"directory.chromeosdevices.patch": {
			Name:       "directory.chromeosdevices.patch",
			HTTPMethod: "PATCH",
			Path:       "admin/directory/v1/customer/{customerId}/devices/chromeos/{deviceId}",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},

				{Name: "deviceId", Type: "string", Location: "path", Required: true},

				{Name: "projection", Type: "string", Location: "query", Required: false},
			},
		},

		"directory.chromeosdevices.update": {
			Name:       "directory.chromeosdevices.update",
			HTTPMethod: "PUT",
			Path:       "admin/directory/v1/customer/{customerId}/devices/chromeos/{deviceId}",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},

				{Name: "deviceId", Type: "string", Location: "path", Required: true},

				{Name: "projection", Type: "string", Location: "query", Required: false},
			},
		},
	},
}

// customer
var Customer = ResourceDescriptor{
	Name:    "customer",
	Methods: map[string]MethodDescriptor{},
}

// customers
var Customers = ResourceDescriptor{
	Name: "customers",
	Methods: map[string]MethodDescriptor{

		"directory.customers.get": {
			Name:       "directory.customers.get",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/customers/{customerKey}",
			Parameters: []ParameterDescriptor{

				{Name: "customerKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.customers.patch": {
			Name:       "directory.customers.patch",
			HTTPMethod: "PATCH",
			Path:       "admin/directory/v1/customers/{customerKey}",
			Parameters: []ParameterDescriptor{

				{Name: "customerKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.customers.update": {
			Name:       "directory.customers.update",
			HTTPMethod: "PUT",
			Path:       "admin/directory/v1/customers/{customerKey}",
			Parameters: []ParameterDescriptor{

				{Name: "customerKey", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// domainAliases
var DomainAliases = ResourceDescriptor{
	Name: "domainAliases",
	Methods: map[string]MethodDescriptor{

		"directory.domainAliases.delete": {
			Name:       "directory.domainAliases.delete",
			HTTPMethod: "DELETE",
			Path:       "admin/directory/v1/customer/{customer}/domainaliases/{domainAliasName}",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "path", Required: true},

				{Name: "domainAliasName", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.domainAliases.get": {
			Name:       "directory.domainAliases.get",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/customer/{customer}/domainaliases/{domainAliasName}",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "path", Required: true},

				{Name: "domainAliasName", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.domainAliases.insert": {
			Name:       "directory.domainAliases.insert",
			HTTPMethod: "POST",
			Path:       "admin/directory/v1/customer/{customer}/domainaliases",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.domainAliases.list": {
			Name:       "directory.domainAliases.list",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/customer/{customer}/domainaliases",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "path", Required: true},

				{Name: "parentDomainName", Type: "string", Location: "query", Required: false},
			},
		},
	},
}

// domains
var Domains = ResourceDescriptor{
	Name: "domains",
	Methods: map[string]MethodDescriptor{

		"directory.domains.delete": {
			Name:       "directory.domains.delete",
			HTTPMethod: "DELETE",
			Path:       "admin/directory/v1/customer/{customer}/domains/{domainName}",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "path", Required: true},

				{Name: "domainName", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.domains.get": {
			Name:       "directory.domains.get",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/customer/{customer}/domains/{domainName}",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "path", Required: true},

				{Name: "domainName", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.domains.insert": {
			Name:       "directory.domains.insert",
			HTTPMethod: "POST",
			Path:       "admin/directory/v1/customer/{customer}/domains",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.domains.list": {
			Name:       "directory.domains.list",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/customer/{customer}/domains",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// groups
var Groups = ResourceDescriptor{
	Name: "groups",
	Methods: map[string]MethodDescriptor{

		"directory.groups.delete": {
			Name:       "directory.groups.delete",
			HTTPMethod: "DELETE",
			Path:       "admin/directory/v1/groups/{groupKey}",
			Parameters: []ParameterDescriptor{

				{Name: "groupKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.groups.get": {
			Name:       "directory.groups.get",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/groups/{groupKey}",
			Parameters: []ParameterDescriptor{

				{Name: "groupKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.groups.insert": {
			Name:       "directory.groups.insert",
			HTTPMethod: "POST",
			Path:       "admin/directory/v1/groups",
			Parameters: []ParameterDescriptor{},
		},

		"directory.groups.list": {
			Name:       "directory.groups.list",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/groups",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "query", Required: false},

				{Name: "domain", Type: "string", Location: "query", Required: false},

				{Name: "maxResults", Type: "integer", Location: "query", Required: false},

				{Name: "orderBy", Type: "string", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "query", Type: "string", Location: "query", Required: false},

				{Name: "sortOrder", Type: "string", Location: "query", Required: false},

				{Name: "userKey", Type: "string", Location: "query", Required: false},
			},
		},

		"directory.groups.patch": {
			Name:       "directory.groups.patch",
			HTTPMethod: "PATCH",
			Path:       "admin/directory/v1/groups/{groupKey}",
			Parameters: []ParameterDescriptor{

				{Name: "groupKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.groups.update": {
			Name:       "directory.groups.update",
			HTTPMethod: "PUT",
			Path:       "admin/directory/v1/groups/{groupKey}",
			Parameters: []ParameterDescriptor{

				{Name: "groupKey", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// members
var Members = ResourceDescriptor{
	Name: "members",
	Methods: map[string]MethodDescriptor{

		"directory.members.delete": {
			Name:       "directory.members.delete",
			HTTPMethod: "DELETE",
			Path:       "admin/directory/v1/groups/{groupKey}/members/{memberKey}",
			Parameters: []ParameterDescriptor{

				{Name: "groupKey", Type: "string", Location: "path", Required: true},

				{Name: "memberKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.members.get": {
			Name:       "directory.members.get",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/groups/{groupKey}/members/{memberKey}",
			Parameters: []ParameterDescriptor{

				{Name: "groupKey", Type: "string", Location: "path", Required: true},

				{Name: "memberKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.members.hasMember": {
			Name:       "directory.members.hasMember",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/groups/{groupKey}/hasMember/{memberKey}",
			Parameters: []ParameterDescriptor{

				{Name: "groupKey", Type: "string", Location: "path", Required: true},

				{Name: "memberKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.members.insert": {
			Name:       "directory.members.insert",
			HTTPMethod: "POST",
			Path:       "admin/directory/v1/groups/{groupKey}/members",
			Parameters: []ParameterDescriptor{

				{Name: "groupKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.members.list": {
			Name:       "directory.members.list",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/groups/{groupKey}/members",
			Parameters: []ParameterDescriptor{

				{Name: "groupKey", Type: "string", Location: "path", Required: true},

				{Name: "includeDerivedMembership", Type: "boolean", Location: "query", Required: false},

				{Name: "maxResults", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "roles", Type: "string", Location: "query", Required: false},
			},
		},

		"directory.members.patch": {
			Name:       "directory.members.patch",
			HTTPMethod: "PATCH",
			Path:       "admin/directory/v1/groups/{groupKey}/members/{memberKey}",
			Parameters: []ParameterDescriptor{

				{Name: "groupKey", Type: "string", Location: "path", Required: true},

				{Name: "memberKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.members.update": {
			Name:       "directory.members.update",
			HTTPMethod: "PUT",
			Path:       "admin/directory/v1/groups/{groupKey}/members/{memberKey}",
			Parameters: []ParameterDescriptor{

				{Name: "groupKey", Type: "string", Location: "path", Required: true},

				{Name: "memberKey", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// mobiledevices
var Mobiledevices = ResourceDescriptor{
	Name: "mobiledevices",
	Methods: map[string]MethodDescriptor{

		"directory.mobiledevices.action": {
			Name:       "directory.mobiledevices.action",
			HTTPMethod: "POST",
			Path:       "admin/directory/v1/customer/{customerId}/devices/mobile/{resourceId}/action",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},

				{Name: "resourceId", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.mobiledevices.delete": {
			Name:       "directory.mobiledevices.delete",
			HTTPMethod: "DELETE",
			Path:       "admin/directory/v1/customer/{customerId}/devices/mobile/{resourceId}",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},

				{Name: "resourceId", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.mobiledevices.get": {
			Name:       "directory.mobiledevices.get",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/customer/{customerId}/devices/mobile/{resourceId}",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},

				{Name: "projection", Type: "string", Location: "query", Required: false},

				{Name: "resourceId", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.mobiledevices.list": {
			Name:       "directory.mobiledevices.list",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/customer/{customerId}/devices/mobile",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},

				{Name: "maxResults", Type: "integer", Location: "query", Required: false},

				{Name: "orderBy", Type: "string", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "projection", Type: "string", Location: "query", Required: false},

				{Name: "query", Type: "string", Location: "query", Required: false},

				{Name: "sortOrder", Type: "string", Location: "query", Required: false},
			},
		},
	},
}

// orgunits
var Orgunits = ResourceDescriptor{
	Name: "orgunits",
	Methods: map[string]MethodDescriptor{

		"directory.orgunits.delete": {
			Name:       "directory.orgunits.delete",
			HTTPMethod: "DELETE",
			Path:       "admin/directory/v1/customer/{customerId}/orgunits/{+orgUnitPath}",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},

				{Name: "orgUnitPath", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.orgunits.get": {
			Name:       "directory.orgunits.get",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/customer/{customerId}/orgunits/{+orgUnitPath}",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},

				{Name: "orgUnitPath", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.orgunits.insert": {
			Name:       "directory.orgunits.insert",
			HTTPMethod: "POST",
			Path:       "admin/directory/v1/customer/{customerId}/orgunits",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.orgunits.list": {
			Name:       "directory.orgunits.list",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/customer/{customerId}/orgunits",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},

				{Name: "orgUnitPath", Type: "string", Location: "query", Required: false},

				{Name: "type", Type: "string", Location: "query", Required: false},
			},
		},

		"directory.orgunits.patch": {
			Name:       "directory.orgunits.patch",
			HTTPMethod: "PATCH",
			Path:       "admin/directory/v1/customer/{customerId}/orgunits/{+orgUnitPath}",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},

				{Name: "orgUnitPath", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.orgunits.update": {
			Name:       "directory.orgunits.update",
			HTTPMethod: "PUT",
			Path:       "admin/directory/v1/customer/{customerId}/orgunits/{+orgUnitPath}",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},

				{Name: "orgUnitPath", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// privileges
var Privileges = ResourceDescriptor{
	Name: "privileges",
	Methods: map[string]MethodDescriptor{

		"directory.privileges.list": {
			Name:       "directory.privileges.list",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/customer/{customer}/roles/ALL/privileges",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// resources
var Resources = ResourceDescriptor{
	Name:    "resources",
	Methods: map[string]MethodDescriptor{},
}

// roleAssignments
var RoleAssignments = ResourceDescriptor{
	Name: "roleAssignments",
	Methods: map[string]MethodDescriptor{

		"directory.roleAssignments.delete": {
			Name:       "directory.roleAssignments.delete",
			HTTPMethod: "DELETE",
			Path:       "admin/directory/v1/customer/{customer}/roleassignments/{roleAssignmentId}",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "path", Required: true},

				{Name: "roleAssignmentId", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.roleAssignments.get": {
			Name:       "directory.roleAssignments.get",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/customer/{customer}/roleassignments/{roleAssignmentId}",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "path", Required: true},

				{Name: "roleAssignmentId", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.roleAssignments.insert": {
			Name:       "directory.roleAssignments.insert",
			HTTPMethod: "POST",
			Path:       "admin/directory/v1/customer/{customer}/roleassignments",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.roleAssignments.list": {
			Name:       "directory.roleAssignments.list",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/customer/{customer}/roleassignments",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "path", Required: true},

				{Name: "includeIndirectRoleAssignments", Type: "boolean", Location: "query", Required: false},

				{Name: "maxResults", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "roleId", Type: "string", Location: "query", Required: false},

				{Name: "userKey", Type: "string", Location: "query", Required: false},
			},
		},
	},
}

// roles
var Roles = ResourceDescriptor{
	Name: "roles",
	Methods: map[string]MethodDescriptor{

		"directory.roles.delete": {
			Name:       "directory.roles.delete",
			HTTPMethod: "DELETE",
			Path:       "admin/directory/v1/customer/{customer}/roles/{roleId}",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "path", Required: true},

				{Name: "roleId", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.roles.get": {
			Name:       "directory.roles.get",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/customer/{customer}/roles/{roleId}",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "path", Required: true},

				{Name: "roleId", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.roles.insert": {
			Name:       "directory.roles.insert",
			HTTPMethod: "POST",
			Path:       "admin/directory/v1/customer/{customer}/roles",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.roles.list": {
			Name:       "directory.roles.list",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/customer/{customer}/roles",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "path", Required: true},

				{Name: "maxResults", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},
			},
		},

		"directory.roles.patch": {
			Name:       "directory.roles.patch",
			HTTPMethod: "PATCH",
			Path:       "admin/directory/v1/customer/{customer}/roles/{roleId}",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "path", Required: true},

				{Name: "roleId", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.roles.update": {
			Name:       "directory.roles.update",
			HTTPMethod: "PUT",
			Path:       "admin/directory/v1/customer/{customer}/roles/{roleId}",
			Parameters: []ParameterDescriptor{

				{Name: "customer", Type: "string", Location: "path", Required: true},

				{Name: "roleId", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// schemas
var Schemas = ResourceDescriptor{
	Name: "schemas",
	Methods: map[string]MethodDescriptor{

		"directory.schemas.delete": {
			Name:       "directory.schemas.delete",
			HTTPMethod: "DELETE",
			Path:       "admin/directory/v1/customer/{customerId}/schemas/{schemaKey}",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},

				{Name: "schemaKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.schemas.get": {
			Name:       "directory.schemas.get",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/customer/{customerId}/schemas/{schemaKey}",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},

				{Name: "schemaKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.schemas.insert": {
			Name:       "directory.schemas.insert",
			HTTPMethod: "POST",
			Path:       "admin/directory/v1/customer/{customerId}/schemas",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.schemas.list": {
			Name:       "directory.schemas.list",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/customer/{customerId}/schemas",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.schemas.patch": {
			Name:       "directory.schemas.patch",
			HTTPMethod: "PATCH",
			Path:       "admin/directory/v1/customer/{customerId}/schemas/{schemaKey}",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},

				{Name: "schemaKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.schemas.update": {
			Name:       "directory.schemas.update",
			HTTPMethod: "PUT",
			Path:       "admin/directory/v1/customer/{customerId}/schemas/{schemaKey}",
			Parameters: []ParameterDescriptor{

				{Name: "customerId", Type: "string", Location: "path", Required: true},

				{Name: "schemaKey", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// tokens
var Tokens = ResourceDescriptor{
	Name: "tokens",
	Methods: map[string]MethodDescriptor{

		"directory.tokens.delete": {
			Name:       "directory.tokens.delete",
			HTTPMethod: "DELETE",
			Path:       "admin/directory/v1/users/{userKey}/tokens/{clientId}",
			Parameters: []ParameterDescriptor{

				{Name: "clientId", Type: "string", Location: "path", Required: true},

				{Name: "userKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.tokens.get": {
			Name:       "directory.tokens.get",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/users/{userKey}/tokens/{clientId}",
			Parameters: []ParameterDescriptor{

				{Name: "clientId", Type: "string", Location: "path", Required: true},

				{Name: "userKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.tokens.list": {
			Name:       "directory.tokens.list",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/users/{userKey}/tokens",
			Parameters: []ParameterDescriptor{

				{Name: "userKey", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// twoStepVerification
var TwoStepVerification = ResourceDescriptor{
	Name: "twoStepVerification",
	Methods: map[string]MethodDescriptor{

		"directory.twoStepVerification.turnOff": {
			Name:       "directory.twoStepVerification.turnOff",
			HTTPMethod: "POST",
			Path:       "admin/directory/v1/users/{userKey}/twoStepVerification/turnOff",
			Parameters: []ParameterDescriptor{

				{Name: "userKey", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// users
var Users = ResourceDescriptor{
	Name: "users",
	Methods: map[string]MethodDescriptor{

		"directory.users.createGuest": {
			Name:       "directory.users.createGuest",
			HTTPMethod: "POST",
			Path:       "admin/directory/v1/users:createGuest",
			Parameters: []ParameterDescriptor{},
		},

		"directory.users.delete": {
			Name:       "directory.users.delete",
			HTTPMethod: "DELETE",
			Path:       "admin/directory/v1/users/{userKey}",
			Parameters: []ParameterDescriptor{

				{Name: "userKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.users.get": {
			Name:       "directory.users.get",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/users/{userKey}",
			Parameters: []ParameterDescriptor{

				{Name: "customFieldMask", Type: "string", Location: "query", Required: false},

				{Name: "projection", Type: "string", Location: "query", Required: false},

				{Name: "userKey", Type: "string", Location: "path", Required: true},

				{Name: "viewType", Type: "string", Location: "query", Required: false},
			},
		},

		"directory.users.insert": {
			Name:       "directory.users.insert",
			HTTPMethod: "POST",
			Path:       "admin/directory/v1/users",
			Parameters: []ParameterDescriptor{

				{Name: "resolveConflictAccount", Type: "boolean", Location: "query", Required: false},
			},
		},

		"directory.users.list": {
			Name:       "directory.users.list",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/users",
			Parameters: []ParameterDescriptor{

				{Name: "customFieldMask", Type: "string", Location: "query", Required: false},

				{Name: "customer", Type: "string", Location: "query", Required: false},

				{Name: "domain", Type: "string", Location: "query", Required: false},

				{Name: "event", Type: "string", Location: "query", Required: false},

				{Name: "maxResults", Type: "integer", Location: "query", Required: false},

				{Name: "orderBy", Type: "string", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "projection", Type: "string", Location: "query", Required: false},

				{Name: "query", Type: "string", Location: "query", Required: false},

				{Name: "showDeleted", Type: "string", Location: "query", Required: false},

				{Name: "sortOrder", Type: "string", Location: "query", Required: false},

				{Name: "viewType", Type: "string", Location: "query", Required: false},
			},
		},

		"directory.users.makeAdmin": {
			Name:       "directory.users.makeAdmin",
			HTTPMethod: "POST",
			Path:       "admin/directory/v1/users/{userKey}/makeAdmin",
			Parameters: []ParameterDescriptor{

				{Name: "userKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.users.patch": {
			Name:       "directory.users.patch",
			HTTPMethod: "PATCH",
			Path:       "admin/directory/v1/users/{userKey}",
			Parameters: []ParameterDescriptor{

				{Name: "userKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.users.signOut": {
			Name:       "directory.users.signOut",
			HTTPMethod: "POST",
			Path:       "admin/directory/v1/users/{userKey}/signOut",
			Parameters: []ParameterDescriptor{

				{Name: "userKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.users.undelete": {
			Name:       "directory.users.undelete",
			HTTPMethod: "POST",
			Path:       "admin/directory/v1/users/{userKey}/undelete",
			Parameters: []ParameterDescriptor{

				{Name: "userKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.users.update": {
			Name:       "directory.users.update",
			HTTPMethod: "PUT",
			Path:       "admin/directory/v1/users/{userKey}",
			Parameters: []ParameterDescriptor{

				{Name: "userKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.users.watch": {
			Name:       "directory.users.watch",
			HTTPMethod: "POST",
			Path:       "admin/directory/v1/users/watch",
			Parameters: []ParameterDescriptor{

				{Name: "customFieldMask", Type: "string", Location: "query", Required: false},

				{Name: "customer", Type: "string", Location: "query", Required: false},

				{Name: "domain", Type: "string", Location: "query", Required: false},

				{Name: "event", Type: "string", Location: "query", Required: false},

				{Name: "maxResults", Type: "integer", Location: "query", Required: false},

				{Name: "orderBy", Type: "string", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "projection", Type: "string", Location: "query", Required: false},

				{Name: "query", Type: "string", Location: "query", Required: false},

				{Name: "showDeleted", Type: "string", Location: "query", Required: false},

				{Name: "sortOrder", Type: "string", Location: "query", Required: false},

				{Name: "viewType", Type: "string", Location: "query", Required: false},
			},
		},
	},
}

// verificationCodes
var VerificationCodes = ResourceDescriptor{
	Name: "verificationCodes",
	Methods: map[string]MethodDescriptor{

		"directory.verificationCodes.generate": {
			Name:       "directory.verificationCodes.generate",
			HTTPMethod: "POST",
			Path:       "admin/directory/v1/users/{userKey}/verificationCodes/generate",
			Parameters: []ParameterDescriptor{

				{Name: "userKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.verificationCodes.invalidate": {
			Name:       "directory.verificationCodes.invalidate",
			HTTPMethod: "POST",
			Path:       "admin/directory/v1/users/{userKey}/verificationCodes/invalidate",
			Parameters: []ParameterDescriptor{

				{Name: "userKey", Type: "string", Location: "path", Required: true},
			},
		},

		"directory.verificationCodes.list": {
			Name:       "directory.verificationCodes.list",
			HTTPMethod: "GET",
			Path:       "admin/directory/v1/users/{userKey}/verificationCodes",
			Parameters: []ParameterDescriptor{

				{Name: "userKey", Type: "string", Location: "path", Required: true},
			},
		},
	},
}
