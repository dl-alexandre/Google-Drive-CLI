// Google Sheets API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package descriptors

// ServiceDescriptor contains metadata for the sheets API
type ServiceDescriptor struct {
	Name        string
	Version     string
	BaseURL     string
	RootURL     string
	ServicePath string
}

// Service returns the service descriptor for sheets
var Service = ServiceDescriptor{
	Name:        "sheets",
	Version:     "v4",
	BaseURL:     "https://sheets.googleapis.com/",
	RootURL:     "https://sheets.googleapis.com/",
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

	"spreadsheets": Spreadsheets,
}

// spreadsheets
var Spreadsheets = ResourceDescriptor{
	Name: "spreadsheets",
	Methods: map[string]MethodDescriptor{

		"sheets.spreadsheets.batchUpdate": {
			Name:       "sheets.spreadsheets.batchUpdate",
			HTTPMethod: "POST",
			Path:       "v4/spreadsheets/{spreadsheetId}:batchUpdate",
			Parameters: []ParameterDescriptor{

				{Name: "spreadsheetId", Type: "string", Location: "path", Required: true},
			},
		},

		"sheets.spreadsheets.create": {
			Name:       "sheets.spreadsheets.create",
			HTTPMethod: "POST",
			Path:       "v4/spreadsheets",
			Parameters: []ParameterDescriptor{},
		},

		"sheets.spreadsheets.get": {
			Name:       "sheets.spreadsheets.get",
			HTTPMethod: "GET",
			Path:       "v4/spreadsheets/{spreadsheetId}",
			Parameters: []ParameterDescriptor{

				{Name: "excludeTablesInBandedRanges", Type: "boolean", Location: "query", Required: false},

				{Name: "includeGridData", Type: "boolean", Location: "query", Required: false},

				{Name: "ranges", Type: "string", Location: "query", Required: false},

				{Name: "spreadsheetId", Type: "string", Location: "path", Required: true},
			},
		},

		"sheets.spreadsheets.getByDataFilter": {
			Name:       "sheets.spreadsheets.getByDataFilter",
			HTTPMethod: "POST",
			Path:       "v4/spreadsheets/{spreadsheetId}:getByDataFilter",
			Parameters: []ParameterDescriptor{

				{Name: "spreadsheetId", Type: "string", Location: "path", Required: true},
			},
		},
	},
}
