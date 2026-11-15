// Google Forms API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package descriptors

// ServiceDescriptor contains metadata for the forms API
type ServiceDescriptor struct {
	Name        string
	Version     string
	BaseURL     string
	RootURL     string
	ServicePath string
}

// Service returns the service descriptor for forms
var Service = ServiceDescriptor{
	Name:        "forms",
	Version:     "v1",
	BaseURL:     "https://forms.googleapis.com/",
	RootURL:     "https://forms.googleapis.com/",
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

	"forms": Forms,
}

// forms
var Forms = ResourceDescriptor{
	Name: "forms",
	Methods: map[string]MethodDescriptor{

		"forms.forms.batchUpdate": {
			Name:       "forms.forms.batchUpdate",
			HTTPMethod: "POST",
			Path:       "v1/forms/{formId}:batchUpdate",
			Parameters: []ParameterDescriptor{

				{Name: "formId", Type: "string", Location: "path", Required: true},
			},
		},

		"forms.forms.create": {
			Name:       "forms.forms.create",
			HTTPMethod: "POST",
			Path:       "v1/forms",
			Parameters: []ParameterDescriptor{

				{Name: "unpublished", Type: "boolean", Location: "query", Required: false},
			},
		},

		"forms.forms.get": {
			Name:       "forms.forms.get",
			HTTPMethod: "GET",
			Path:       "v1/forms/{formId}",
			Parameters: []ParameterDescriptor{

				{Name: "formId", Type: "string", Location: "path", Required: true},
			},
		},

		"forms.forms.setPublishSettings": {
			Name:       "forms.forms.setPublishSettings",
			HTTPMethod: "POST",
			Path:       "v1/forms/{formId}:setPublishSettings",
			Parameters: []ParameterDescriptor{

				{Name: "formId", Type: "string", Location: "path", Required: true},
			},
		},
	},
}
