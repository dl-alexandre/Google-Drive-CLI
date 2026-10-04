// Gmail API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package descriptors

// ServiceDescriptor contains metadata for the gmail API
type ServiceDescriptor struct {
	Name        string
	Version     string
	BaseURL     string
	RootURL     string
	ServicePath string
}

// Service returns the service descriptor for gmail
var Service = ServiceDescriptor{
	Name:        "gmail",
	Version:     "v1",
	BaseURL:     "https://gmail.googleapis.com/",
	RootURL:     "https://gmail.googleapis.com/",
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

	"users": Users,
}

// users
var Users = ResourceDescriptor{
	Name: "users",
	Methods: map[string]MethodDescriptor{

		"gmail.users.getProfile": {
			Name:       "gmail.users.getProfile",
			HTTPMethod: "GET",
			Path:       "gmail/v1/users/{userId}/profile",
			Parameters: []ParameterDescriptor{

				{Name: "userId", Type: "string", Location: "path", Required: true},
			},
		},

		"gmail.users.stop": {
			Name:       "gmail.users.stop",
			HTTPMethod: "POST",
			Path:       "gmail/v1/users/{userId}/stop",
			Parameters: []ParameterDescriptor{

				{Name: "userId", Type: "string", Location: "path", Required: true},
			},
		},

		"gmail.users.watch": {
			Name:       "gmail.users.watch",
			HTTPMethod: "POST",
			Path:       "gmail/v1/users/{userId}/watch",
			Parameters: []ParameterDescriptor{

				{Name: "userId", Type: "string", Location: "path", Required: true},
			},
		},
	},
}
