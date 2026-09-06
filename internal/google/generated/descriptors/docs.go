// Google Docs API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package descriptors

// ServiceDescriptor contains metadata for the docs API
type ServiceDescriptor struct {
	Name        string
	Version     string
	BaseURL     string
	RootURL     string
	ServicePath string
}

// Service returns the service descriptor for docs
var Service = ServiceDescriptor{
	Name:        "docs",
	Version:     "v1",
	BaseURL:     "https://docs.googleapis.com/",
	RootURL:     "https://docs.googleapis.com/",
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

	"documents": Documents,
}

// documents
var Documents = ResourceDescriptor{
	Name: "documents",
	Methods: map[string]MethodDescriptor{

		"docs.documents.batchUpdate": {
			Name:       "docs.documents.batchUpdate",
			HTTPMethod: "POST",
			Path:       "v1/documents/{documentId}:batchUpdate",
			Parameters: []ParameterDescriptor{

				{Name: "documentId", Type: "string", Location: "path", Required: true},
			},
		},

		"docs.documents.create": {
			Name:       "docs.documents.create",
			HTTPMethod: "POST",
			Path:       "v1/documents",
			Parameters: []ParameterDescriptor{},
		},

		"docs.documents.get": {
			Name:       "docs.documents.get",
			HTTPMethod: "GET",
			Path:       "v1/documents/{documentId}",
			Parameters: []ParameterDescriptor{

				{Name: "documentId", Type: "string", Location: "path", Required: true},

				{Name: "includeTabsContent", Type: "boolean", Location: "query", Required: false},

				{Name: "suggestionsViewMode", Type: "string", Location: "query", Required: false},
			},
		},
	},
}
