// Google Slides API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package descriptors

// ServiceDescriptor contains metadata for the slides API
type ServiceDescriptor struct {
	Name        string
	Version     string
	BaseURL     string
	RootURL     string
	ServicePath string
}

// Service returns the service descriptor for slides
var Service = ServiceDescriptor{
	Name:        "slides",
	Version:     "v1",
	BaseURL:     "https://slides.googleapis.com/",
	RootURL:     "https://slides.googleapis.com/",
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

	"presentations": Presentations,
}

// presentations
var Presentations = ResourceDescriptor{
	Name: "presentations",
	Methods: map[string]MethodDescriptor{

		"slides.presentations.batchUpdate": {
			Name:       "slides.presentations.batchUpdate",
			HTTPMethod: "POST",
			Path:       "v1/presentations/{presentationId}:batchUpdate",
			Parameters: []ParameterDescriptor{

				{Name: "presentationId", Type: "string", Location: "path", Required: true},
			},
		},

		"slides.presentations.create": {
			Name:       "slides.presentations.create",
			HTTPMethod: "POST",
			Path:       "v1/presentations",
			Parameters: []ParameterDescriptor{},
		},

		"slides.presentations.get": {
			Name:       "slides.presentations.get",
			HTTPMethod: "GET",
			Path:       "v1/presentations/{+presentationId}",
			Parameters: []ParameterDescriptor{

				{Name: "presentationId", Type: "string", Location: "path", Required: true},
			},
		},
	},
}
