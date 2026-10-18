// Apps Script API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package descriptors

// ServiceDescriptor contains metadata for the script API
type ServiceDescriptor struct {
	Name        string
	Version     string
	BaseURL     string
	RootURL     string
	ServicePath string
}

// Service returns the service descriptor for script
var Service = ServiceDescriptor{
	Name:        "script",
	Version:     "v1",
	BaseURL:     "https://script.googleapis.com/",
	RootURL:     "https://script.googleapis.com/",
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

	"processes": Processes,

	"projects": Projects,

	"scripts": Scripts,
}

// processes
var Processes = ResourceDescriptor{
	Name: "processes",
	Methods: map[string]MethodDescriptor{

		"script.processes.list": {
			Name:       "script.processes.list",
			HTTPMethod: "GET",
			Path:       "v1/processes",
			Parameters: []ParameterDescriptor{

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "userProcessFilter.deploymentId", Type: "string", Location: "query", Required: false},

				{Name: "userProcessFilter.endTime", Type: "string", Location: "query", Required: false},

				{Name: "userProcessFilter.functionName", Type: "string", Location: "query", Required: false},

				{Name: "userProcessFilter.projectName", Type: "string", Location: "query", Required: false},

				{Name: "userProcessFilter.scriptId", Type: "string", Location: "query", Required: false},

				{Name: "userProcessFilter.startTime", Type: "string", Location: "query", Required: false},

				{Name: "userProcessFilter.statuses", Type: "string", Location: "query", Required: false},

				{Name: "userProcessFilter.types", Type: "string", Location: "query", Required: false},

				{Name: "userProcessFilter.userAccessLevels", Type: "string", Location: "query", Required: false},
			},
		},

		"script.processes.listScriptProcesses": {
			Name:       "script.processes.listScriptProcesses",
			HTTPMethod: "GET",
			Path:       "v1/processes:listScriptProcesses",
			Parameters: []ParameterDescriptor{

				{Name: "pageSize", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "scriptId", Type: "string", Location: "query", Required: false},

				{Name: "scriptProcessFilter.deploymentId", Type: "string", Location: "query", Required: false},

				{Name: "scriptProcessFilter.endTime", Type: "string", Location: "query", Required: false},

				{Name: "scriptProcessFilter.functionName", Type: "string", Location: "query", Required: false},

				{Name: "scriptProcessFilter.startTime", Type: "string", Location: "query", Required: false},

				{Name: "scriptProcessFilter.statuses", Type: "string", Location: "query", Required: false},

				{Name: "scriptProcessFilter.types", Type: "string", Location: "query", Required: false},

				{Name: "scriptProcessFilter.userAccessLevels", Type: "string", Location: "query", Required: false},
			},
		},
	},
}

// projects
var Projects = ResourceDescriptor{
	Name: "projects",
	Methods: map[string]MethodDescriptor{

		"script.projects.create": {
			Name:       "script.projects.create",
			HTTPMethod: "POST",
			Path:       "v1/projects",
			Parameters: []ParameterDescriptor{},
		},

		"script.projects.get": {
			Name:       "script.projects.get",
			HTTPMethod: "GET",
			Path:       "v1/projects/{scriptId}",
			Parameters: []ParameterDescriptor{

				{Name: "scriptId", Type: "string", Location: "path", Required: true},
			},
		},

		"script.projects.getContent": {
			Name:       "script.projects.getContent",
			HTTPMethod: "GET",
			Path:       "v1/projects/{scriptId}/content",
			Parameters: []ParameterDescriptor{

				{Name: "scriptId", Type: "string", Location: "path", Required: true},

				{Name: "versionNumber", Type: "integer", Location: "query", Required: false},
			},
		},

		"script.projects.getMetrics": {
			Name:       "script.projects.getMetrics",
			HTTPMethod: "GET",
			Path:       "v1/projects/{scriptId}/metrics",
			Parameters: []ParameterDescriptor{

				{Name: "metricsFilter.deploymentId", Type: "string", Location: "query", Required: false},

				{Name: "metricsGranularity", Type: "string", Location: "query", Required: false},

				{Name: "scriptId", Type: "string", Location: "path", Required: true},
			},
		},

		"script.projects.updateContent": {
			Name:       "script.projects.updateContent",
			HTTPMethod: "PUT",
			Path:       "v1/projects/{scriptId}/content",
			Parameters: []ParameterDescriptor{

				{Name: "scriptId", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// scripts
var Scripts = ResourceDescriptor{
	Name: "scripts",
	Methods: map[string]MethodDescriptor{

		"script.scripts.run": {
			Name:       "script.scripts.run",
			HTTPMethod: "POST",
			Path:       "v1/scripts/{scriptId}:run",
			Parameters: []ParameterDescriptor{

				{Name: "scriptId", Type: "string", Location: "path", Required: true},
			},
		},
	},
}
