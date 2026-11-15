// Calendar API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package descriptors

// ServiceDescriptor contains metadata for the calendar API
type ServiceDescriptor struct {
	Name        string
	Version     string
	BaseURL     string
	RootURL     string
	ServicePath string
}

// Service returns the service descriptor for calendar
var Service = ServiceDescriptor{
	Name:        "calendar",
	Version:     "v3",
	BaseURL:     "https://www.googleapis.com/calendar/v3/",
	RootURL:     "https://www.googleapis.com/",
	ServicePath: "calendar/v3/",
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

	"acl": Acl,

	"calendarList": CalendarList,

	"calendars": Calendars,

	"channels": Channels,

	"colors": Colors,

	"events": Events,

	"freebusy": Freebusy,

	"settings": Settings,
}

// acl
var Acl = ResourceDescriptor{
	Name: "acl",
	Methods: map[string]MethodDescriptor{

		"calendar.acl.delete": {
			Name:       "calendar.acl.delete",
			HTTPMethod: "DELETE",
			Path:       "calendars/{calendarId}/acl/{ruleId}",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},

				{Name: "ruleId", Type: "string", Location: "path", Required: true},
			},
		},

		"calendar.acl.get": {
			Name:       "calendar.acl.get",
			HTTPMethod: "GET",
			Path:       "calendars/{calendarId}/acl/{ruleId}",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},

				{Name: "ruleId", Type: "string", Location: "path", Required: true},
			},
		},

		"calendar.acl.insert": {
			Name:       "calendar.acl.insert",
			HTTPMethod: "POST",
			Path:       "calendars/{calendarId}/acl",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},

				{Name: "sendNotifications", Type: "boolean", Location: "query", Required: false},
			},
		},

		"calendar.acl.list": {
			Name:       "calendar.acl.list",
			HTTPMethod: "GET",
			Path:       "calendars/{calendarId}/acl",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},

				{Name: "maxResults", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "showDeleted", Type: "boolean", Location: "query", Required: false},

				{Name: "syncToken", Type: "string", Location: "query", Required: false},
			},
		},

		"calendar.acl.patch": {
			Name:       "calendar.acl.patch",
			HTTPMethod: "PATCH",
			Path:       "calendars/{calendarId}/acl/{ruleId}",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},

				{Name: "ruleId", Type: "string", Location: "path", Required: true},

				{Name: "sendNotifications", Type: "boolean", Location: "query", Required: false},
			},
		},

		"calendar.acl.update": {
			Name:       "calendar.acl.update",
			HTTPMethod: "PUT",
			Path:       "calendars/{calendarId}/acl/{ruleId}",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},

				{Name: "ruleId", Type: "string", Location: "path", Required: true},

				{Name: "sendNotifications", Type: "boolean", Location: "query", Required: false},
			},
		},

		"calendar.acl.watch": {
			Name:       "calendar.acl.watch",
			HTTPMethod: "POST",
			Path:       "calendars/{calendarId}/acl/watch",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},

				{Name: "maxResults", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "showDeleted", Type: "boolean", Location: "query", Required: false},

				{Name: "syncToken", Type: "string", Location: "query", Required: false},
			},
		},
	},
}

// calendarList
var CalendarList = ResourceDescriptor{
	Name: "calendarList",
	Methods: map[string]MethodDescriptor{

		"calendar.calendarList.delete": {
			Name:       "calendar.calendarList.delete",
			HTTPMethod: "DELETE",
			Path:       "users/me/calendarList/{calendarId}",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},
			},
		},

		"calendar.calendarList.get": {
			Name:       "calendar.calendarList.get",
			HTTPMethod: "GET",
			Path:       "users/me/calendarList/{calendarId}",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},
			},
		},

		"calendar.calendarList.insert": {
			Name:       "calendar.calendarList.insert",
			HTTPMethod: "POST",
			Path:       "users/me/calendarList",
			Parameters: []ParameterDescriptor{

				{Name: "colorRgbFormat", Type: "boolean", Location: "query", Required: false},
			},
		},

		"calendar.calendarList.list": {
			Name:       "calendar.calendarList.list",
			HTTPMethod: "GET",
			Path:       "users/me/calendarList",
			Parameters: []ParameterDescriptor{

				{Name: "maxResults", Type: "integer", Location: "query", Required: false},

				{Name: "minAccessRole", Type: "string", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "showDeleted", Type: "boolean", Location: "query", Required: false},

				{Name: "showHidden", Type: "boolean", Location: "query", Required: false},

				{Name: "showOwnOrganizationOnly", Type: "boolean", Location: "query", Required: false},

				{Name: "syncToken", Type: "string", Location: "query", Required: false},
			},
		},

		"calendar.calendarList.patch": {
			Name:       "calendar.calendarList.patch",
			HTTPMethod: "PATCH",
			Path:       "users/me/calendarList/{calendarId}",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},

				{Name: "colorRgbFormat", Type: "boolean", Location: "query", Required: false},
			},
		},

		"calendar.calendarList.update": {
			Name:       "calendar.calendarList.update",
			HTTPMethod: "PUT",
			Path:       "users/me/calendarList/{calendarId}",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},

				{Name: "colorRgbFormat", Type: "boolean", Location: "query", Required: false},
			},
		},

		"calendar.calendarList.watch": {
			Name:       "calendar.calendarList.watch",
			HTTPMethod: "POST",
			Path:       "users/me/calendarList/watch",
			Parameters: []ParameterDescriptor{

				{Name: "maxResults", Type: "integer", Location: "query", Required: false},

				{Name: "minAccessRole", Type: "string", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "showDeleted", Type: "boolean", Location: "query", Required: false},

				{Name: "showHidden", Type: "boolean", Location: "query", Required: false},

				{Name: "showOwnOrganizationOnly", Type: "boolean", Location: "query", Required: false},

				{Name: "syncToken", Type: "string", Location: "query", Required: false},
			},
		},
	},
}

// calendars
var Calendars = ResourceDescriptor{
	Name: "calendars",
	Methods: map[string]MethodDescriptor{

		"calendar.calendars.clear": {
			Name:       "calendar.calendars.clear",
			HTTPMethod: "POST",
			Path:       "calendars/{calendarId}/clear",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},
			},
		},

		"calendar.calendars.delete": {
			Name:       "calendar.calendars.delete",
			HTTPMethod: "DELETE",
			Path:       "calendars/{calendarId}",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},
			},
		},

		"calendar.calendars.get": {
			Name:       "calendar.calendars.get",
			HTTPMethod: "GET",
			Path:       "calendars/{calendarId}",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},
			},
		},

		"calendar.calendars.insert": {
			Name:       "calendar.calendars.insert",
			HTTPMethod: "POST",
			Path:       "calendars",
			Parameters: []ParameterDescriptor{},
		},

		"calendar.calendars.patch": {
			Name:       "calendar.calendars.patch",
			HTTPMethod: "PATCH",
			Path:       "calendars/{calendarId}",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},
			},
		},

		"calendar.calendars.transferOwnership": {
			Name:       "calendar.calendars.transferOwnership",
			HTTPMethod: "POST",
			Path:       "calendars/{calendarId}/transferOwnership",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},

				{Name: "newDataOwner", Type: "string", Location: "query", Required: true},

				{Name: "useAdminAccess", Type: "boolean", Location: "query", Required: true},
			},
		},

		"calendar.calendars.update": {
			Name:       "calendar.calendars.update",
			HTTPMethod: "PUT",
			Path:       "calendars/{calendarId}",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},
			},
		},
	},
}

// channels
var Channels = ResourceDescriptor{
	Name: "channels",
	Methods: map[string]MethodDescriptor{

		"calendar.channels.stop": {
			Name:       "calendar.channels.stop",
			HTTPMethod: "POST",
			Path:       "channels/stop",
			Parameters: []ParameterDescriptor{},
		},
	},
}

// colors
var Colors = ResourceDescriptor{
	Name: "colors",
	Methods: map[string]MethodDescriptor{

		"calendar.colors.get": {
			Name:       "calendar.colors.get",
			HTTPMethod: "GET",
			Path:       "colors",
			Parameters: []ParameterDescriptor{},
		},
	},
}

// events
var Events = ResourceDescriptor{
	Name: "events",
	Methods: map[string]MethodDescriptor{

		"calendar.events.delete": {
			Name:       "calendar.events.delete",
			HTTPMethod: "DELETE",
			Path:       "calendars/{calendarId}/events/{eventId}",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},

				{Name: "eventId", Type: "string", Location: "path", Required: true},

				{Name: "sendNotifications", Type: "boolean", Location: "query", Required: false},

				{Name: "sendUpdates", Type: "string", Location: "query", Required: false},
			},
		},

		"calendar.events.get": {
			Name:       "calendar.events.get",
			HTTPMethod: "GET",
			Path:       "calendars/{calendarId}/events/{eventId}",
			Parameters: []ParameterDescriptor{

				{Name: "alwaysIncludeEmail", Type: "boolean", Location: "query", Required: false},

				{Name: "calendarId", Type: "string", Location: "path", Required: true},

				{Name: "eventId", Type: "string", Location: "path", Required: true},

				{Name: "maxAttendees", Type: "integer", Location: "query", Required: false},

				{Name: "timeZone", Type: "string", Location: "query", Required: false},
			},
		},

		"calendar.events.import": {
			Name:       "calendar.events.import",
			HTTPMethod: "POST",
			Path:       "calendars/{calendarId}/events/import",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},

				{Name: "conferenceDataVersion", Type: "integer", Location: "query", Required: false},

				{Name: "eventLabelVersion", Type: "integer", Location: "query", Required: false},

				{Name: "supportsAttachments", Type: "boolean", Location: "query", Required: false},
			},
		},

		"calendar.events.insert": {
			Name:       "calendar.events.insert",
			HTTPMethod: "POST",
			Path:       "calendars/{calendarId}/events",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},

				{Name: "conferenceDataVersion", Type: "integer", Location: "query", Required: false},

				{Name: "eventLabelVersion", Type: "integer", Location: "query", Required: false},

				{Name: "maxAttendees", Type: "integer", Location: "query", Required: false},

				{Name: "sendNotifications", Type: "boolean", Location: "query", Required: false},

				{Name: "sendUpdates", Type: "string", Location: "query", Required: false},

				{Name: "supportsAttachments", Type: "boolean", Location: "query", Required: false},
			},
		},

		"calendar.events.instances": {
			Name:       "calendar.events.instances",
			HTTPMethod: "GET",
			Path:       "calendars/{calendarId}/events/{eventId}/instances",
			Parameters: []ParameterDescriptor{

				{Name: "alwaysIncludeEmail", Type: "boolean", Location: "query", Required: false},

				{Name: "calendarId", Type: "string", Location: "path", Required: true},

				{Name: "eventId", Type: "string", Location: "path", Required: true},

				{Name: "maxAttendees", Type: "integer", Location: "query", Required: false},

				{Name: "maxResults", Type: "integer", Location: "query", Required: false},

				{Name: "originalStart", Type: "string", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "showDeleted", Type: "boolean", Location: "query", Required: false},

				{Name: "timeMax", Type: "string", Location: "query", Required: false},

				{Name: "timeMin", Type: "string", Location: "query", Required: false},

				{Name: "timeZone", Type: "string", Location: "query", Required: false},
			},
		},

		"calendar.events.list": {
			Name:       "calendar.events.list",
			HTTPMethod: "GET",
			Path:       "calendars/{calendarId}/events",
			Parameters: []ParameterDescriptor{

				{Name: "alwaysIncludeEmail", Type: "boolean", Location: "query", Required: false},

				{Name: "calendarId", Type: "string", Location: "path", Required: true},

				{Name: "eventTypes", Type: "string", Location: "query", Required: false},

				{Name: "iCalUID", Type: "string", Location: "query", Required: false},

				{Name: "maxAttendees", Type: "integer", Location: "query", Required: false},

				{Name: "maxResults", Type: "integer", Location: "query", Required: false},

				{Name: "orderBy", Type: "string", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "privateExtendedProperty", Type: "string", Location: "query", Required: false},

				{Name: "q", Type: "string", Location: "query", Required: false},

				{Name: "sharedExtendedProperty", Type: "string", Location: "query", Required: false},

				{Name: "showDeleted", Type: "boolean", Location: "query", Required: false},

				{Name: "showHiddenInvitations", Type: "boolean", Location: "query", Required: false},

				{Name: "singleEvents", Type: "boolean", Location: "query", Required: false},

				{Name: "syncToken", Type: "string", Location: "query", Required: false},

				{Name: "timeMax", Type: "string", Location: "query", Required: false},

				{Name: "timeMin", Type: "string", Location: "query", Required: false},

				{Name: "timeZone", Type: "string", Location: "query", Required: false},

				{Name: "updatedMin", Type: "string", Location: "query", Required: false},
			},
		},

		"calendar.events.move": {
			Name:       "calendar.events.move",
			HTTPMethod: "POST",
			Path:       "calendars/{calendarId}/events/{eventId}/move",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},

				{Name: "destination", Type: "string", Location: "query", Required: true},

				{Name: "eventId", Type: "string", Location: "path", Required: true},

				{Name: "sendNotifications", Type: "boolean", Location: "query", Required: false},

				{Name: "sendUpdates", Type: "string", Location: "query", Required: false},
			},
		},

		"calendar.events.patch": {
			Name:       "calendar.events.patch",
			HTTPMethod: "PATCH",
			Path:       "calendars/{calendarId}/events/{eventId}",
			Parameters: []ParameterDescriptor{

				{Name: "alwaysIncludeEmail", Type: "boolean", Location: "query", Required: false},

				{Name: "calendarId", Type: "string", Location: "path", Required: true},

				{Name: "conferenceDataVersion", Type: "integer", Location: "query", Required: false},

				{Name: "eventId", Type: "string", Location: "path", Required: true},

				{Name: "eventLabelVersion", Type: "integer", Location: "query", Required: false},

				{Name: "maxAttendees", Type: "integer", Location: "query", Required: false},

				{Name: "sendNotifications", Type: "boolean", Location: "query", Required: false},

				{Name: "sendUpdates", Type: "string", Location: "query", Required: false},

				{Name: "supportsAttachments", Type: "boolean", Location: "query", Required: false},
			},
		},

		"calendar.events.quickAdd": {
			Name:       "calendar.events.quickAdd",
			HTTPMethod: "POST",
			Path:       "calendars/{calendarId}/events/quickAdd",
			Parameters: []ParameterDescriptor{

				{Name: "calendarId", Type: "string", Location: "path", Required: true},

				{Name: "sendNotifications", Type: "boolean", Location: "query", Required: false},

				{Name: "sendUpdates", Type: "string", Location: "query", Required: false},

				{Name: "text", Type: "string", Location: "query", Required: true},
			},
		},

		"calendar.events.update": {
			Name:       "calendar.events.update",
			HTTPMethod: "PUT",
			Path:       "calendars/{calendarId}/events/{eventId}",
			Parameters: []ParameterDescriptor{

				{Name: "alwaysIncludeEmail", Type: "boolean", Location: "query", Required: false},

				{Name: "calendarId", Type: "string", Location: "path", Required: true},

				{Name: "conferenceDataVersion", Type: "integer", Location: "query", Required: false},

				{Name: "eventId", Type: "string", Location: "path", Required: true},

				{Name: "eventLabelVersion", Type: "integer", Location: "query", Required: false},

				{Name: "maxAttendees", Type: "integer", Location: "query", Required: false},

				{Name: "sendNotifications", Type: "boolean", Location: "query", Required: false},

				{Name: "sendUpdates", Type: "string", Location: "query", Required: false},

				{Name: "supportsAttachments", Type: "boolean", Location: "query", Required: false},
			},
		},

		"calendar.events.watch": {
			Name:       "calendar.events.watch",
			HTTPMethod: "POST",
			Path:       "calendars/{calendarId}/events/watch",
			Parameters: []ParameterDescriptor{

				{Name: "alwaysIncludeEmail", Type: "boolean", Location: "query", Required: false},

				{Name: "calendarId", Type: "string", Location: "path", Required: true},

				{Name: "eventTypes", Type: "string", Location: "query", Required: false},

				{Name: "iCalUID", Type: "string", Location: "query", Required: false},

				{Name: "maxAttendees", Type: "integer", Location: "query", Required: false},

				{Name: "maxResults", Type: "integer", Location: "query", Required: false},

				{Name: "orderBy", Type: "string", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "privateExtendedProperty", Type: "string", Location: "query", Required: false},

				{Name: "q", Type: "string", Location: "query", Required: false},

				{Name: "sharedExtendedProperty", Type: "string", Location: "query", Required: false},

				{Name: "showDeleted", Type: "boolean", Location: "query", Required: false},

				{Name: "showHiddenInvitations", Type: "boolean", Location: "query", Required: false},

				{Name: "singleEvents", Type: "boolean", Location: "query", Required: false},

				{Name: "syncToken", Type: "string", Location: "query", Required: false},

				{Name: "timeMax", Type: "string", Location: "query", Required: false},

				{Name: "timeMin", Type: "string", Location: "query", Required: false},

				{Name: "timeZone", Type: "string", Location: "query", Required: false},

				{Name: "updatedMin", Type: "string", Location: "query", Required: false},
			},
		},
	},
}

// freebusy
var Freebusy = ResourceDescriptor{
	Name: "freebusy",
	Methods: map[string]MethodDescriptor{

		"calendar.freebusy.query": {
			Name:       "calendar.freebusy.query",
			HTTPMethod: "POST",
			Path:       "freeBusy",
			Parameters: []ParameterDescriptor{},
		},
	},
}

// settings
var Settings = ResourceDescriptor{
	Name: "settings",
	Methods: map[string]MethodDescriptor{

		"calendar.settings.get": {
			Name:       "calendar.settings.get",
			HTTPMethod: "GET",
			Path:       "users/me/settings/{setting}",
			Parameters: []ParameterDescriptor{

				{Name: "setting", Type: "string", Location: "path", Required: true},
			},
		},

		"calendar.settings.list": {
			Name:       "calendar.settings.list",
			HTTPMethod: "GET",
			Path:       "users/me/settings",
			Parameters: []ParameterDescriptor{

				{Name: "maxResults", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "syncToken", Type: "string", Location: "query", Required: false},
			},
		},

		"calendar.settings.watch": {
			Name:       "calendar.settings.watch",
			HTTPMethod: "POST",
			Path:       "users/me/settings/watch",
			Parameters: []ParameterDescriptor{

				{Name: "maxResults", Type: "integer", Location: "query", Required: false},

				{Name: "pageToken", Type: "string", Location: "query", Required: false},

				{Name: "syncToken", Type: "string", Location: "query", Required: false},
			},
		},
	},
}
