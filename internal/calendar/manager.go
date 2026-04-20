package calendar

import (
	"context"
	"fmt"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/types"
	calendar "google.golang.org/api/calendar/v3"
)

// Manager handles Google Calendar API operations.
type Manager struct {
	client  *api.Client
	service *calendar.Service
}

// NewManager creates a new Calendar manager.
func NewManager(client *api.Client, service *calendar.Service) *Manager {
	return &Manager{
		client:  client,
		service: service,
	}
}

// ListEvents lists events from the specified calendar.
// Default calendarID is "primary". singleEvents expands recurring events.
func (m *Manager) ListEvents(ctx context.Context, reqCtx *types.RequestContext, calendarID, timeMin, timeMax string, maxResults int64, pageToken, query string, singleEvents bool, orderBy string) (*types.CalendarEventList, string, error) {
	if calendarID == "" {
		calendarID = "primary"
	}

	call := m.service.Events.List(calendarID)
	if timeMin != "" {
		call = call.TimeMin(timeMin)
	}
	if timeMax != "" {
		call = call.TimeMax(timeMax)
	}
	if maxResults > 0 {
		call = call.MaxResults(maxResults)
	}
	if pageToken != "" {
		call = call.PageToken(pageToken)
	}
	if query != "" {
		call = call.Q(query)
	}
	if singleEvents {
		call = call.SingleEvents(true)
	}
	if orderBy != "" {
		call = call.OrderBy(orderBy)
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*calendar.Events, error) {
		return call.Do()
	})
	if err != nil {
		return nil, "", err
	}

	events := make([]types.CalendarEvent, len(result.Items))
	for i, item := range result.Items {
		events[i] = convertEvent(item)
	}

	return &types.CalendarEventList{
		Events:  events,
		Summary: result.Summary,
	}, result.NextPageToken, nil
}

// GetEvent retrieves a single event by ID.
func (m *Manager) GetEvent(ctx context.Context, reqCtx *types.RequestContext, calendarID, eventID string) (*types.CalendarEvent, error) {
	if calendarID == "" {
		calendarID = "primary"
	}

	call := m.service.Events.Get(calendarID, eventID)

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*calendar.Event, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	event := convertEvent(result)
	return &event, nil
}

// SearchEvents searches for events matching a query string.
func (m *Manager) SearchEvents(ctx context.Context, reqCtx *types.RequestContext, calendarID, query, timeMin, timeMax string) (*types.CalendarEventList, string, error) {
	if calendarID == "" {
		calendarID = "primary"
	}

	call := m.service.Events.List(calendarID).Q(query)
	if timeMin != "" {
		call = call.TimeMin(timeMin)
	}
	if timeMax != "" {
		call = call.TimeMax(timeMax)
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*calendar.Events, error) {
		return call.Do()
	})
	if err != nil {
		return nil, "", err
	}

	events := make([]types.CalendarEvent, len(result.Items))
	for i, item := range result.Items {
		events[i] = convertEvent(item)
	}

	return &types.CalendarEventList{
		Events:  events,
		Summary: result.Summary,
	}, result.NextPageToken, nil
}

// CreateEvent creates a new calendar event.
// For all-day events, set allDay to true and pass date strings (YYYY-MM-DD) for startTime/endTime.
// For timed events, pass RFC3339 timestamps.
// Date and DateTime are mutually exclusive on EventDateTime — never set both.
func (m *Manager) CreateEvent(ctx context.Context, reqCtx *types.RequestContext, calendarID, summary, description, location, startTime, endTime string, attendees []string, allDay bool, recurrence []string, sendUpdates string) (*types.CalendarEventResult, error) {
	if calendarID == "" {
		calendarID = "primary"
	}

	event := &calendar.Event{
		Summary:     summary,
		Description: description,
		Location:    location,
		Start:       buildEventDateTime(startTime, allDay),
		End:         buildEventDateTime(endTime, allDay),
	}

	if len(attendees) > 0 {
		eventAttendees := make([]*calendar.EventAttendee, len(attendees))
		for i, email := range attendees {
			eventAttendees[i] = &calendar.EventAttendee{Email: email}
		}
		event.Attendees = eventAttendees
	}

	if len(recurrence) > 0 {
		event.Recurrence = recurrence
	}

	call := m.service.Events.Insert(calendarID, event)
	if sendUpdates != "" {
		call = call.SendUpdates(sendUpdates)
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*calendar.Event, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	return &types.CalendarEventResult{
		ID:       result.Id,
		Summary:  result.Summary,
		HtmlLink: result.HtmlLink,
	}, nil
}

// UpdateEvent patches an existing calendar event. Only non-nil fields are updated.
func (m *Manager) UpdateEvent(ctx context.Context, reqCtx *types.RequestContext, calendarID, eventID string, summary, description, location, startTime, endTime *string, sendUpdates string) (*types.CalendarEventResult, error) {
	if calendarID == "" {
		calendarID = "primary"
	}

	patch := &calendar.Event{}
	if summary != nil {
		patch.Summary = *summary
	}
	if description != nil {
		patch.Description = *description
	}
	if location != nil {
		patch.Location = *location
	}
	if startTime != nil {
		patch.Start = buildEventDateTime(*startTime, false)
	}
	if endTime != nil {
		patch.End = buildEventDateTime(*endTime, false)
	}

	call := m.service.Events.Patch(calendarID, eventID, patch)
	if sendUpdates != "" {
		call = call.SendUpdates(sendUpdates)
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*calendar.Event, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	return &types.CalendarEventResult{
		ID:       result.Id,
		Summary:  result.Summary,
		HtmlLink: result.HtmlLink,
	}, nil
}

// DeleteEvent deletes an event from the calendar.
func (m *Manager) DeleteEvent(ctx context.Context, reqCtx *types.RequestContext, calendarID, eventID, sendUpdates string) error {
	if calendarID == "" {
		calendarID = "primary"
	}

	call := m.service.Events.Delete(calendarID, eventID)
	if sendUpdates != "" {
		call = call.SendUpdates(sendUpdates)
	}

	_, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*struct{}, error) {
		return nil, call.Do()
	})
	return err
}

// RespondToEvent updates the authenticated user's attendance response for an event.
// response should be one of: "accepted", "declined", "tentative".
func (m *Manager) RespondToEvent(ctx context.Context, reqCtx *types.RequestContext, calendarID, eventID, response, sendUpdates string) (*types.CalendarEventResult, error) {
	if calendarID == "" {
		calendarID = "primary"
	}

	// First, get the current event to find the self attendee
	getCall := m.service.Events.Get(calendarID, eventID)
	event, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*calendar.Event, error) {
		return getCall.Do()
	})
	if err != nil {
		return nil, err
	}

	// Find self in attendees and update response status
	found := false
	for _, attendee := range event.Attendees {
		if attendee.Self {
			attendee.ResponseStatus = response
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("current user not found in event attendees")
	}

	// Patch the event with updated attendees
	patch := &calendar.Event{
		Attendees: event.Attendees,
	}

	patchCall := m.service.Events.Patch(calendarID, eventID, patch)
	if sendUpdates != "" {
		patchCall = patchCall.SendUpdates(sendUpdates)
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*calendar.Event, error) {
		return patchCall.Do()
	})
	if err != nil {
		return nil, err
	}

	return &types.CalendarEventResult{
		ID:       result.Id,
		Summary:  result.Summary,
		HtmlLink: result.HtmlLink,
	}, nil
}

// FreeBusy queries free/busy information for the specified calendars.
func (m *Manager) FreeBusy(ctx context.Context, reqCtx *types.RequestContext, calendarIDs []string, timeMin, timeMax string) (*types.CalendarFreeBusy, error) {
	items := make([]*calendar.FreeBusyRequestItem, len(calendarIDs))
	for i, id := range calendarIDs {
		items[i] = &calendar.FreeBusyRequestItem{Id: id}
	}

	req := &calendar.FreeBusyRequest{
		TimeMin: timeMin,
		TimeMax: timeMax,
		Items:   items,
	}

	call := m.service.Freebusy.Query(req)

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*calendar.FreeBusyResponse, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	calendars := make(map[string][]types.CalendarBusyPeriod)
	for calID, busyCal := range result.Calendars {
		periods := make([]types.CalendarBusyPeriod, len(busyCal.Busy))
		for i, busy := range busyCal.Busy {
			periods[i] = types.CalendarBusyPeriod{
				Start: busy.Start,
				End:   busy.End,
			}
		}
		calendars[calID] = periods
	}

	return &types.CalendarFreeBusy{
		Calendars: calendars,
	}, nil
}

// FindConflicts lists events in the given time range and identifies overlapping events.
func (m *Manager) FindConflicts(ctx context.Context, reqCtx *types.RequestContext, calendarID, startTime, endTime string) ([]types.CalendarConflict, error) {
	if calendarID == "" {
		calendarID = "primary"
	}

	call := m.service.Events.List(calendarID).
		TimeMin(startTime).
		TimeMax(endTime).
		SingleEvents(true).
		OrderBy("startTime")

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*calendar.Events, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	events := make([]types.CalendarEvent, len(result.Items))
	for i, item := range result.Items {
		events[i] = convertEvent(item)
	}

	return detectConflicts(events), nil
}

// detectConflicts finds overlapping events in a sorted slice.
// Two events conflict if one starts before the other ends.
func detectConflicts(events []types.CalendarEvent) []types.CalendarConflict {
	var conflicts []types.CalendarConflict

	for i := 0; i < len(events); i++ {
		var overlaps []types.CalendarEvent
		for j := i + 1; j < len(events); j++ {
			// Events are sorted by start time. Event j starts after event i.
			// Conflict exists if event j starts before event i ends.
			if events[j].Start < events[i].End {
				overlaps = append(overlaps, events[j])
			}
		}
		if len(overlaps) > 0 {
			conflicts = append(conflicts, types.CalendarConflict{
				Event:         events[i],
				ConflictsWith: overlaps,
			})
		}
	}

	return conflicts
}

// buildEventDateTime constructs an EventDateTime for the Calendar API.
// If allDay is true, the Date field is set (YYYY-MM-DD format) and DateTime is left empty.
// If allDay is false, the DateTime field is set (RFC3339 format) and Date is left empty.
// Date and DateTime are mutually exclusive — never set both.
func buildEventDateTime(dateTimeStr string, allDay bool) *calendar.EventDateTime {
	if allDay {
		return &calendar.EventDateTime{
			Date: dateTimeStr,
		}
	}
	return &calendar.EventDateTime{
		DateTime: dateTimeStr,
	}
}

// convertEvent converts a Google Calendar API Event to the domain CalendarEvent type.
// For Start/End, it prefers DateTime if set, otherwise falls back to Date.
func convertEvent(e *calendar.Event) types.CalendarEvent {
	if e == nil {
		return types.CalendarEvent{}
	}

	event := types.CalendarEvent{
		ID:          e.Id,
		Summary:     e.Summary,
		Description: e.Description,
		Location:    e.Location,
		Status:      e.Status,
		HangoutLink: e.HangoutLink,
		HtmlLink:    e.HtmlLink,
		Recurrence:  e.Recurrence,
	}

	if e.Start != nil {
		if e.Start.DateTime != "" {
			event.Start = e.Start.DateTime
		} else {
			event.Start = e.Start.Date
		}
	}

	if e.End != nil {
		if e.End.DateTime != "" {
			event.End = e.End.DateTime
		} else {
			event.End = e.End.Date
		}
	}

	if e.Creator != nil {
		event.Creator = e.Creator.Email
	}

	if e.Organizer != nil {
		event.Organizer = e.Organizer.Email
	}

	if e.Attendees != nil {
		event.Attendees = convertAttendees(e.Attendees)
	}

	return event
}

// convertAttendees converts a slice of Google Calendar API EventAttendee to domain CalendarAttendee.
func convertAttendees(attendees []*calendar.EventAttendee) []types.CalendarAttendee {
	result := make([]types.CalendarAttendee, len(attendees))
	for i, a := range attendees {
		result[i] = types.CalendarAttendee{
			Email:          a.Email,
			DisplayName:    a.DisplayName,
			ResponseStatus: a.ResponseStatus,
			Self:           a.Self,
			Optional:       a.Optional,
		}
	}
	return result
}
