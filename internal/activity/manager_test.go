package activity

import (
	"testing"
	"time"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/types"
	"google.golang.org/api/driveactivity/v2"
)

func TestManager_Creation(t *testing.T) {
	client := &api.Client{}
	manager := NewManager(client)

	if manager == nil {
		t.Error("NewManager returned nil")
	}

	if manager.client != client {
		t.Error("Manager client not set correctly")
	}
}

func TestBuildTimeFilter(t *testing.T) {
	tests := []struct {
		name      string
		startTime string
		endTime   string
		expected  string
	}{
		{
			name:      "both times provided",
			startTime: "2026-01-01T00:00:00Z",
			endTime:   "2026-01-31T23:59:59Z",
			expected:  "time >= '2026-01-01T00:00:00Z' AND time <= '2026-01-31T23:59:59Z'",
		},
		{
			name:      "only start time",
			startTime: "2026-01-01T00:00:00Z",
			endTime:   "",
			expected:  "time >= '2026-01-01T00:00:00Z'",
		},
		{
			name:      "only end time",
			startTime: "",
			endTime:   "2026-01-31T23:59:59Z",
			expected:  "time <= '2026-01-31T23:59:59Z'",
		},
		{
			name:      "no times provided",
			startTime: "",
			endTime:   "",
			expected:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildTimeFilter(tt.startTime, tt.endTime)
			if result != tt.expected {
				t.Errorf("buildTimeFilter(%q, %q) = %q, want %q", tt.startTime, tt.endTime, result, tt.expected)
			}
		})
	}
}

func TestBuildActionFilter(t *testing.T) {
	tests := []struct {
		name        string
		actionTypes string
		expected    string
	}{
		{
			name:        "single action type",
			actionTypes: "edit",
			expected:    "detail.action_detail_case:EDIT",
		},
		{
			name:        "multiple action types",
			actionTypes: "edit,share,delete",
			expected:    "detail.action_detail_case:EDIT OR detail.action_detail_case:PERMISSION_CHANGE OR detail.action_detail_case:DELETE",
		},
		{
			name:        "action types with spaces",
			actionTypes: "edit, comment, move",
			expected:    "detail.action_detail_case:EDIT OR detail.action_detail_case:COMMENT OR detail.action_detail_case:MOVE",
		},
		{
			name:        "permission_change alias",
			actionTypes: "permission_change",
			expected:    "detail.action_detail_case:PERMISSION_CHANGE",
		},
		{
			name:        "share alias for permission_change",
			actionTypes: "share",
			expected:    "detail.action_detail_case:PERMISSION_CHANGE",
		},
		{
			name:        "all action types",
			actionTypes: "edit,comment,share,move,delete,restore,create,rename",
			expected:    "detail.action_detail_case:EDIT OR detail.action_detail_case:COMMENT OR detail.action_detail_case:PERMISSION_CHANGE OR detail.action_detail_case:MOVE OR detail.action_detail_case:DELETE OR detail.action_detail_case:RESTORE OR detail.action_detail_case:CREATE OR detail.action_detail_case:RENAME",
		},
		{
			name:        "empty string",
			actionTypes: "",
			expected:    "",
		},
		{
			name:        "unknown action type",
			actionTypes: "unknown",
			expected:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildActionFilter(tt.actionTypes)
			if result != tt.expected {
				t.Errorf("buildActionFilter(%q) = %q, want %q", tt.actionTypes, result, tt.expected)
			}
		})
	}
}

func TestConvertActivity(t *testing.T) {
	tests := []struct {
		name     string
		input    *driveactivity.DriveActivity
		validate func(*types.Activity) error
	}{
		{
			name: "activity with timestamp",
			input: &driveactivity.DriveActivity{
				Timestamp: "2026-01-15T10:30:00Z",
				Actors: []*driveactivity.Actor{
					{
						User: &driveactivity.User{
							KnownUser: &driveactivity.KnownUser{
								PersonName: "people/user123",
							},
						},
					},
				},
				Targets: []*driveactivity.Target{
					{
						DriveItem: &driveactivity.DriveItem{
							Name:     "items/file123",
							Title:    "test.txt",
							MimeType: "text/plain",
						},
					},
				},
				Actions: []*driveactivity.Action{
					{
						Detail: &driveactivity.ActionDetail{
							Edit: &driveactivity.Edit{},
						},
					},
				},
			},
			validate: func(a *types.Activity) error {
				if a.Timestamp.IsZero() {
					return errorf("timestamp should not be zero")
				}
				if len(a.Actors) == 0 {
					return errorf("actors should not be empty")
				}
				if len(a.Targets) == 0 {
					return errorf("targets should not be empty")
				}
				if len(a.Actions) == 0 {
					return errorf("actions should not be empty")
				}
				return nil
			},
		},
		{
			name: "activity with multiple actors",
			input: &driveactivity.DriveActivity{
				Actors: []*driveactivity.Actor{
					{
						User: &driveactivity.User{
							KnownUser: &driveactivity.KnownUser{
								PersonName: "people/user1",
							},
						},
					},
					{
						Administrator: &driveactivity.Administrator{},
					},
				},
			},
			validate: func(a *types.Activity) error {
				if len(a.Actors) != 2 {
					return errorf("expected 2 actors, got %d", len(a.Actors))
				}
				if a.Actors[0].Type != "user" {
					return errorf("first actor should be user, got %s", a.Actors[0].Type)
				}
				if a.Actors[1].Type != "administrator" {
					return errorf("second actor should be administrator, got %s", a.Actors[1].Type)
				}
				return nil
			},
		},
		{
			name: "activity with nil actors",
			input: &driveactivity.DriveActivity{
				Actors: nil,
			},
			validate: func(a *types.Activity) error {
				if len(a.Actors) != 0 {
					return errorf("actors should be empty for nil input")
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertActivity(tt.input)
			if err := tt.validate(&result); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestConvertActors(t *testing.T) {
	tests := []struct {
		name     string
		input    []*driveactivity.Actor
		expected int
		validate func([]types.Actor) error
	}{
		{
			name: "user actor",
			input: []*driveactivity.Actor{
				{
					User: &driveactivity.User{
						KnownUser: &driveactivity.KnownUser{
							PersonName: "people/user123",
						},
					},
				},
			},
			expected: 1,
			validate: func(actors []types.Actor) error {
				if actors[0].Type != "user" {
					return errorf("expected user type, got %s", actors[0].Type)
				}
				if actors[0].User == nil {
					return errorf("user should not be nil")
				}
				return nil
			},
		},
		{
			name: "administrator actor",
			input: []*driveactivity.Actor{
				{
					Administrator: &driveactivity.Administrator{},
				},
			},
			expected: 1,
			validate: func(actors []types.Actor) error {
				if actors[0].Type != "administrator" {
					return errorf("expected administrator type, got %s", actors[0].Type)
				}
				return nil
			},
		},
		{
			name: "anonymous actor",
			input: []*driveactivity.Actor{
				{
					Anonymous: &driveactivity.AnonymousUser{},
				},
			},
			expected: 1,
			validate: func(actors []types.Actor) error {
				if actors[0].Type != "anonymous" {
					return errorf("expected anonymous type, got %s", actors[0].Type)
				}
				return nil
			},
		},
		{
			name:     "empty actors",
			input:    []*driveactivity.Actor{},
			expected: 0,
			validate: func(actors []types.Actor) error {
				return nil
			},
		},
		{
			name:     "nil actors",
			input:    nil,
			expected: 0,
			validate: func(actors []types.Actor) error {
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertActors(tt.input)
			if len(result) != tt.expected {
				t.Errorf("expected %d actors, got %d", tt.expected, len(result))
			}
			if err := tt.validate(result); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestConvertTargets(t *testing.T) {
	tests := []struct {
		name     string
		input    []*driveactivity.Target
		expected int
		validate func([]types.Target) error
	}{
		{
			name: "drive item target",
			input: []*driveactivity.Target{
				{
					DriveItem: &driveactivity.DriveItem{
						Name:     "items/file123",
						Title:    "test.txt",
						MimeType: "text/plain",
					},
				},
			},
			expected: 1,
			validate: func(targets []types.Target) error {
				if targets[0].Type != "driveItem" {
					return errorf("expected driveItem type, got %s", targets[0].Type)
				}
				if targets[0].DriveItem == nil {
					return errorf("driveItem should not be nil")
				}
				if targets[0].DriveItem.Title != "test.txt" {
					return errorf("expected title 'test.txt', got %s", targets[0].DriveItem.Title)
				}
				return nil
			},
		},
		{
			name: "drive target",
			input: []*driveactivity.Target{
				{
					Drive: &driveactivity.Drive{},
				},
			},
			expected: 1,
			validate: func(targets []types.Target) error {
				if targets[0].Type != "drive" {
					return errorf("expected drive type, got %s", targets[0].Type)
				}
				return nil
			},
		},
		{
			name: "file comment target",
			input: []*driveactivity.Target{
				{
					FileComment: &driveactivity.FileComment{},
				},
			},
			expected: 1,
			validate: func(targets []types.Target) error {
				if targets[0].Type != "fileComment" {
					return errorf("expected fileComment type, got %s", targets[0].Type)
				}
				return nil
			},
		},
		{
			name:     "empty targets",
			input:    []*driveactivity.Target{},
			expected: 0,
			validate: func(targets []types.Target) error {
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertTargets(tt.input)
			if len(result) != tt.expected {
				t.Errorf("expected %d targets, got %d", tt.expected, len(result))
			}
			if err := tt.validate(result); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestConvertActions(t *testing.T) {
	tests := []struct {
		name     string
		input    []*driveactivity.Action
		expected int
		validate func([]types.Action) error
	}{
		{
			name: "edit action",
			input: []*driveactivity.Action{
				{
					Detail: &driveactivity.ActionDetail{
						Edit: &driveactivity.Edit{},
					},
				},
			},
			expected: 1,
			validate: func(actions []types.Action) error {
				if actions[0].Type != "edit" {
					return errorf("expected edit type, got %s", actions[0].Type)
				}
				return nil
			},
		},
		{
			name: "comment action",
			input: []*driveactivity.Action{
				{
					Detail: &driveactivity.ActionDetail{
						Comment: &driveactivity.Comment{},
					},
				},
			},
			expected: 1,
			validate: func(actions []types.Action) error {
				if actions[0].Type != "comment" {
					return errorf("expected comment type, got %s", actions[0].Type)
				}
				return nil
			},
		},
		{
			name: "permission change action",
			input: []*driveactivity.Action{
				{
					Detail: &driveactivity.ActionDetail{
						PermissionChange: &driveactivity.PermissionChange{},
					},
				},
			},
			expected: 1,
			validate: func(actions []types.Action) error {
				if actions[0].Type != "permission_change" {
					return errorf("expected permission_change type, got %s", actions[0].Type)
				}
				return nil
			},
		},
		{
			name:     "empty actions",
			input:    []*driveactivity.Action{},
			expected: 0,
			validate: func(actions []types.Action) error {
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertActions(tt.input)
			if len(result) != tt.expected {
				t.Errorf("expected %d actions, got %d", tt.expected, len(result))
			}
			if err := tt.validate(result); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestParseTimestamp(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		shouldErr bool
		validate  func(time.Time) error
	}{
		{
			name:      "valid RFC3339 timestamp",
			input:     "2026-01-15T10:30:00Z",
			shouldErr: false,
			validate: func(ts time.Time) error {
				if ts.IsZero() {
					return errorf("timestamp should not be zero")
				}
				if ts.Year() != 2026 {
					return errorf("expected year 2026, got %d", ts.Year())
				}
				return nil
			},
		},
		{
			name:      "valid RFC3339 with timezone",
			input:     "2026-01-15T10:30:00-05:00",
			shouldErr: false,
			validate: func(ts time.Time) error {
				if ts.IsZero() {
					return errorf("timestamp should not be zero")
				}
				return nil
			},
		},
		{
			name:      "invalid timestamp",
			input:     "invalid",
			shouldErr: true,
			validate: func(ts time.Time) error {
				return nil
			},
		},
		{
			name:      "empty string",
			input:     "",
			shouldErr: true,
			validate: func(ts time.Time) error {
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseTimestamp(tt.input)
			if (err != nil) != tt.shouldErr {
				t.Errorf("parseTimestamp(%q) error = %v, shouldErr %v", tt.input, err, tt.shouldErr)
			}
			if !tt.shouldErr {
				if err := tt.validate(result); err != nil {
					t.Error(err)
				}
			}
		})
	}
}

func errorf(format string, args ...interface{}) error {
	return &testError{msg: format, args: args}
}

type testError struct {
	msg  string
	args []interface{}
}

func (e *testError) Error() string {
	if len(e.args) == 0 {
		return e.msg
	}
	return sprintf(e.msg, e.args...)
}

func sprintf(format string, args ...interface{}) string {
	return format
}
