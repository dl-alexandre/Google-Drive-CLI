// People API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package descriptors

// ServiceDescriptor contains metadata for the people API
type ServiceDescriptor struct {
	Name        string
	Version     string
	BaseURL     string
	RootURL     string
	ServicePath string
}

// Service returns the service descriptor for people
var Service = ServiceDescriptor{
	Name:        "people",
	Version:     "v1",
	BaseURL:     "https://people.googleapis.com/",
	RootURL:     "https://people.googleapis.com/",
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

	"contactGroups": ContactGroups,

	"otherContacts": OtherContacts,

	"people": People,
}

// contactGroups
var ContactGroups = ResourceDescriptor{
	Name: "contactGroups",
	Methods: map[string]MethodDescriptor{

		"people.contactGroups.batchGet": {
			Name:       "people.contactGroups.batchGet",
			HTTPMethod: "GET",
			Path:       "v1/contactGroups:batchGet",
			Parameters: []ParameterDescriptor{

				{Name: "groupFields", Type: "string", Location: "query", Required: false},

				{Name: "maxMembers", Type: "integer", Location: "query", Required: false},

				{Name: "resourceNames", Type: "string", Location: "query", Required: false},
			},
		},

		"people.contactGroups.create": {
			Name:       "people.contactGroups.create",
			HTTPMethod: "POST",
			Path:       "v1/contactGroups",
			Parameters: []ParameterDescriptor{},
		},

		"people.contactGroups.delete": {
			Name:       "people.contactGroups.delete",
			HTTPMethod: "DELETE",
			Path:       "v1/{+resourceName}",
			Parameters: []ParameterDescriptor{

				{Name: "deleteContacts", Type: "boolean", Location: "query", Required: false},

				{Name: "resourceName", Type: "string", Location: "path", Required: true},
			},
		},

		"people.contactGroups.get": {
			Name:       "people.contactGroups.get",
			HTTPMethod: "GET",
			Path:       "v1/{+resourceName}",
			Parameters: []ParameterDescriptor{

				{Name: "groupFields", Type: "string", Location: "query", Required: false},

				{Name: "maxMembers", Type: "integer", Location: "query", Required: false},

				{Name: "resourceName", Type: "string", Location: "path", Required: true},
			},
		},

		"people.contactGroups.list": {
			Name:       "people.contactGroups.list",
			HTTPMethod: "GET",
			Path:       "v1/contactGroups",
			Parameters: []ParameterDescriptor{

				{Name: "groupFields", Type: "string", Location: "query", Required: false},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "syncToken", Type: "string", Location: "query", Required: false},
			},
		},

		"people.contactGroups.update": {
			Name:       "people.contactGroups.update",
			HTTPMethod: "PUT",
			Path:       "v1/{+resourceName}",
			Parameters: []ParameterDescriptor{

				{Name: "resourceName", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// otherContacts
var OtherContacts = ResourceDescriptor{
	Name: "otherContacts",
	Methods: map[string]MethodDescriptor{

		"people.otherContacts.copyOtherContactToMyContactsGroup": {
			Name:       "people.otherContacts.copyOtherContactToMyContactsGroup",
			HTTPMethod: "POST",
			Path:       "v1/{+resourceName}:copyOtherContactToMyContactsGroup",
			Parameters: []ParameterDescriptor{

				{Name: "resourceName", Type: "string", Location: "path", Required: true},
			},
		},

		"people.otherContacts.list": {
			Name:       "people.otherContacts.list",
			HTTPMethod: "GET",
			Path:       "v1/otherContacts",
			Parameters: []ParameterDescriptor{

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "readMask", Type: "string", Location: "query", Required: false},

				{Name: "requestSyncToken", Type: "boolean", Location: "query", Required: false},

				{Name: "sources", Type: "string", Location: "query", Required: false},

				{Name: "syncToken", Type: "string", Location: "query", Required: false},
			},
		},

		"people.otherContacts.search": {
			Name:       "people.otherContacts.search",
			HTTPMethod: "GET",
			Path:       "v1/otherContacts:search",
			Parameters: []ParameterDescriptor{

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "query", Type: "string", Location: "query", Required: false},

				{Name: "readMask", Type: "string", Location: "query", Required: false},
			},
		},
	},
}

// people
var People = ResourceDescriptor{
	Name: "people",
	Methods: map[string]MethodDescriptor{

		"people.people.batchCreateContacts": {
			Name:       "people.people.batchCreateContacts",
			HTTPMethod: "POST",
			Path:       "v1/people:batchCreateContacts",
			Parameters: []ParameterDescriptor{},
		},

		"people.people.batchDeleteContacts": {
			Name:       "people.people.batchDeleteContacts",
			HTTPMethod: "POST",
			Path:       "v1/people:batchDeleteContacts",
			Parameters: []ParameterDescriptor{},
		},

		"people.people.batchUpdateContacts": {
			Name:       "people.people.batchUpdateContacts",
			HTTPMethod: "POST",
			Path:       "v1/people:batchUpdateContacts",
			Parameters: []ParameterDescriptor{},
		},

		"people.people.createContact": {
			Name:       "people.people.createContact",
			HTTPMethod: "POST",
			Path:       "v1/people:createContact",
			Parameters: []ParameterDescriptor{

				{Name: "personFields", Type: "string", Location: "query", Required: false},

				{Name: "sources", Type: "string", Location: "query", Required: false},
			},
		},

		"people.people.deleteContact": {
			Name:       "people.people.deleteContact",
			HTTPMethod: "DELETE",
			Path:       "v1/{+resourceName}:deleteContact",
			Parameters: []ParameterDescriptor{

				{Name: "resourceName", Type: "string", Location: "path", Required: true},
			},
		},

		"people.people.deleteContactPhoto": {
			Name:       "people.people.deleteContactPhoto",
			HTTPMethod: "DELETE",
			Path:       "v1/{+resourceName}:deleteContactPhoto",
			Parameters: []ParameterDescriptor{

				{Name: "personFields", Type: "string", Location: "query", Required: false},

				{Name: "resourceName", Type: "string", Location: "path", Required: true},

				{Name: "sources", Type: "string", Location: "query", Required: false},
			},
		},

		"people.people.get": {
			Name:       "people.people.get",
			HTTPMethod: "GET",
			Path:       "v1/{+resourceName}",
			Parameters: []ParameterDescriptor{

				{Name: "personFields", Type: "string", Location: "query", Required: false},

				{Name: "requestMask.includeField", Type: "string", Location: "query", Required: false},

				{Name: "resourceName", Type: "string", Location: "path", Required: true},

				{Name: "sources", Type: "string", Location: "query", Required: false},
			},
		},

		"people.people.getBatchGet": {
			Name:       "people.people.getBatchGet",
			HTTPMethod: "GET",
			Path:       "v1/people:batchGet",
			Parameters: []ParameterDescriptor{

				{Name: "personFields", Type: "string", Location: "query", Required: false},

				{Name: "requestMask.includeField", Type: "string", Location: "query", Required: false},

				{Name: "resourceNames", Type: "string", Location: "query", Required: false},

				{Name: "sources", Type: "string", Location: "query", Required: false},
			},
		},

		"people.people.listDirectoryPeople": {
			Name:       "people.people.listDirectoryPeople",
			HTTPMethod: "GET",
			Path:       "v1/people:listDirectoryPeople",
			Parameters: []ParameterDescriptor{

				{Name: "mergeSources", Type: "string", Location: "query", Required: false},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "readMask", Type: "string", Location: "query", Required: false},

				{Name: "requestSyncToken", Type: "boolean", Location: "query", Required: false},

				{Name: "sources", Type: "string", Location: "query", Required: false},

				{Name: "syncToken", Type: "string", Location: "query", Required: false},
			},
		},

		"people.people.searchContacts": {
			Name:       "people.people.searchContacts",
			HTTPMethod: "GET",
			Path:       "v1/people:searchContacts",
			Parameters: []ParameterDescriptor{

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "query", Type: "string", Location: "query", Required: false},

				{Name: "readMask", Type: "string", Location: "query", Required: false},

				{Name: "sources", Type: "string", Location: "query", Required: false},
			},
		},

		"people.people.searchDirectoryPeople": {
			Name:       "people.people.searchDirectoryPeople",
			HTTPMethod: "GET",
			Path:       "v1/people:searchDirectoryPeople",
			Parameters: []ParameterDescriptor{

				{Name: "mergeSources", Type: "string", Location: "query", Required: false},

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "query", Type: "string", Location: "query", Required: false},

				{Name: "readMask", Type: "string", Location: "query", Required: false},

				{Name: "sources", Type: "string", Location: "query", Required: false},
			},
		},

		"people.people.updateContact": {
			Name:       "people.people.updateContact",
			HTTPMethod: "PATCH",
			Path:       "v1/{+resourceName}:updateContact",
			Parameters: []ParameterDescriptor{

				{Name: "personFields", Type: "string", Location: "query", Required: false},

				{Name: "resourceName", Type: "string", Location: "path", Required: true},

				{Name: "sources", Type: "string", Location: "query", Required: false},

				{Name: "updatePersonFields", Type: "string", Location: "query", Required: false},
			},
		},

		"people.people.updateContactPhoto": {
			Name:       "people.people.updateContactPhoto",
			HTTPMethod: "PATCH",
			Path:       "v1/{+resourceName}:updateContactPhoto",
			Parameters: []ParameterDescriptor{

				{Name: "resourceName", Type: "string", Location: "path", Required: true},
			},
		},
	},
}
