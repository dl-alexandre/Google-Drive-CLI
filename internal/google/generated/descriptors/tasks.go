// Google Tasks API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package descriptors

// ServiceDescriptor contains metadata for the tasks API
type ServiceDescriptor struct {
	Name        string
	Version     string
	BaseURL     string
	RootURL     string
	ServicePath string
}

// Service returns the service descriptor for tasks
var Service = ServiceDescriptor{
	Name:        "tasks",
	Version:     "v1",
	BaseURL:     "https://tasks.googleapis.com/",
	RootURL:     "https://tasks.googleapis.com/",
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

	"tasklists": Tasklists,

	"tasks": Tasks,
}

// tasklists
var Tasklists = ResourceDescriptor{
	Name: "tasklists",
	Methods: map[string]MethodDescriptor{

		"tasks.tasklists.delete": {
			Name:       "tasks.tasklists.delete",
			HTTPMethod: "DELETE",
			Path:       "tasks/v1/users/@me/lists/{tasklist}",
			Parameters: []ParameterDescriptor{

				{Name: "tasklist", Type: "string", Location: "path", Required: true},
			},
		},

		"tasks.tasklists.get": {
			Name:       "tasks.tasklists.get",
			HTTPMethod: "GET",
			Path:       "tasks/v1/users/@me/lists/{tasklist}",
			Parameters: []ParameterDescriptor{

				{Name: "tasklist", Type: "string", Location: "path", Required: true},
			},
		},

		"tasks.tasklists.insert": {
			Name:       "tasks.tasklists.insert",
			HTTPMethod: "POST",
			Path:       "tasks/v1/users/@me/lists",
			Parameters: []ParameterDescriptor{},
		},

		"tasks.tasklists.list": {
			Name:       "tasks.tasklists.list",
			HTTPMethod: "GET",
			Path:       "tasks/v1/users/@me/lists",
			Parameters: []ParameterDescriptor{

				{Name: "maxResults", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},
			},
		},

		"tasks.tasklists.patch": {
			Name:       "tasks.tasklists.patch",
			HTTPMethod: "PATCH",
			Path:       "tasks/v1/users/@me/lists/{tasklist}",
			Parameters: []ParameterDescriptor{

				{Name: "tasklist", Type: "string", Location: "path", Required: true},
			},
		},

		"tasks.tasklists.update": {
			Name:       "tasks.tasklists.update",
			HTTPMethod: "PUT",
			Path:       "tasks/v1/users/@me/lists/{tasklist}",
			Parameters: []ParameterDescriptor{

				{Name: "tasklist", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// tasks
var Tasks = ResourceDescriptor{
	Name: "tasks",
	Methods: map[string]MethodDescriptor{

		"tasks.tasks.clear": {
			Name:       "tasks.tasks.clear",
			HTTPMethod: "POST",
			Path:       "tasks/v1/lists/{tasklist}/clear",
			Parameters: []ParameterDescriptor{

				{Name: "tasklist", Type: "string", Location: "path", Required: true},
			},
		},

		"tasks.tasks.delete": {
			Name:       "tasks.tasks.delete",
			HTTPMethod: "DELETE",
			Path:       "tasks/v1/lists/{tasklist}/tasks/{task}",
			Parameters: []ParameterDescriptor{

				{Name: "task", Type: "string", Location: "path", Required: true},

				{Name: "tasklist", Type: "string", Location: "path", Required: true},
			},
		},

		"tasks.tasks.get": {
			Name:       "tasks.tasks.get",
			HTTPMethod: "GET",
			Path:       "tasks/v1/lists/{tasklist}/tasks/{task}",
			Parameters: []ParameterDescriptor{

				{Name: "task", Type: "string", Location: "path", Required: true},

				{Name: "tasklist", Type: "string", Location: "path", Required: true},
			},
		},

		"tasks.tasks.insert": {
			Name:       "tasks.tasks.insert",
			HTTPMethod: "POST",
			Path:       "tasks/v1/lists/{tasklist}/tasks",
			Parameters: []ParameterDescriptor{

				{Name: "parent", Type: "string", Location: "query", Required: false},

				{Name: "previous", Type: "string", Location: "query", Required: false},

				{Name: "tasklist", Type: "string", Location: "path", Required: true},
			},
		},

		"tasks.tasks.list": {
			Name:       "tasks.tasks.list",
			HTTPMethod: "GET",
			Path:       "tasks/v1/lists/{tasklist}/tasks",
			Parameters: []ParameterDescriptor{

				{Name: "completedMax", Type: "string", Location: "query", Required: false},

				{Name: "completedMin", Type: "string", Location: "query", Required: false},

				{Name: "dueMax", Type: "string", Location: "query", Required: false},

				{Name: "dueMin", Type: "string", Location: "query", Required: false},

				{Name: "maxResults", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "showAssigned", Type: "boolean", Location: "query", Required: false},

				{Name: "showCompleted", Type: "boolean", Location: "query", Required: false},

				{Name: "showDeleted", Type: "boolean", Location: "query", Required: false},

				{Name: "showHidden", Type: "boolean", Location: "query", Required: false},

				{Name: "tasklist", Type: "string", Location: "path", Required: true},

				{Name: "updatedMin", Type: "string", Location: "query", Required: false},
			},
		},

		"tasks.tasks.move": {
			Name:       "tasks.tasks.move",
			HTTPMethod: "POST",
			Path:       "tasks/v1/lists/{tasklist}/tasks/{task}/move",
			Parameters: []ParameterDescriptor{

				{Name: "destinationTasklist", Type: "string", Location: "query", Required: false},

				{Name: "parent", Type: "string", Location: "query", Required: false},

				{Name: "previous", Type: "string", Location: "query", Required: false},

				{Name: "task", Type: "string", Location: "path", Required: true},

				{Name: "tasklist", Type: "string", Location: "path", Required: true},
			},
		},

		"tasks.tasks.patch": {
			Name:       "tasks.tasks.patch",
			HTTPMethod: "PATCH",
			Path:       "tasks/v1/lists/{tasklist}/tasks/{task}",
			Parameters: []ParameterDescriptor{

				{Name: "task", Type: "string", Location: "path", Required: true},

				{Name: "tasklist", Type: "string", Location: "path", Required: true},
			},
		},

		"tasks.tasks.update": {
			Name:       "tasks.tasks.update",
			HTTPMethod: "PUT",
			Path:       "tasks/v1/lists/{tasklist}/tasks/{task}",
			Parameters: []ParameterDescriptor{

				{Name: "task", Type: "string", Location: "path", Required: true},

				{Name: "tasklist", Type: "string", Location: "path", Required: true},
			},
		},
	},
}
