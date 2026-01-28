package changes

import (
	"testing"
	"time"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/types"
	"google.golang.org/api/drive/v3"
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

	if manager.shaper == nil {
		t.Error("Manager shaper not initialized")
	}
}

func TestConvertChange(t *testing.T) {
	tests := []struct {
		name     string
		input    *drive.Change
		validate func(*types.Change) error
	}{
		{
			name: "change with file",
			input: &drive.Change{
				Type:    "file",
				FileId:  "file123",
				Removed: false,
				Time:    "2026-01-15T10:30:00Z",
				File: &drive.File{
					Id:       "file123",
					Name:     "test.txt",
					MimeType: "text/plain",
					Trashed:  false,
				},
			},
			validate: func(change *types.Change) error {
				if change.ChangeType != "file" {
					return errorf("ChangeType mismatch")
				}
				if change.FileID != "file123" {
					return errorf("FileID mismatch")
				}
				if change.Removed {
					return errorf("Removed should be false")
				}
				if change.File == nil {
					return errorf("File should not be nil")
				}
				if change.File.ID != "file123" {
					return errorf("File.ID mismatch")
				}
				return nil
			},
		},
		{
			name: "change with drive",
			input: &drive.Change{
				Type:    "drive",
				DriveId: "drive123",
				Removed: false,
				Drive: &drive.Drive{
					Id:   "drive123",
					Name: "Shared Drive",
				},
			},
			validate: func(change *types.Change) error {
				if change.ChangeType != "drive" {
					return errorf("ChangeType mismatch")
				}
				if change.DriveID != "drive123" {
					return errorf("DriveID mismatch")
				}
				if change.Drive == nil {
					return errorf("Drive should not be nil")
				}
				return nil
			},
		},
		{
			name: "removed change",
			input: &drive.Change{
				Type:    "file",
				FileId:  "file123",
				Removed: true,
			},
			validate: func(change *types.Change) error {
				if !change.Removed {
					return errorf("Removed should be true")
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertChange(tt.input)
			if err := tt.validate(&result); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestConvertDriveFile(t *testing.T) {
	tests := []struct {
		name     string
		input    *drive.File
		validate func(*types.DriveFile) error
	}{
		{
			name: "file with basic properties",
			input: &drive.File{
				Id:       "file123",
				Name:     "test.txt",
				MimeType: "text/plain",
				Trashed:  false,
			},
			validate: func(file *types.DriveFile) error {
				if file.ID != "file123" {
					return errorf("ID mismatch")
				}
				if file.Name != "test.txt" {
					return errorf("Name mismatch")
				}
				if file.MimeType != "text/plain" {
					return errorf("MimeType mismatch")
				}
				if file.Trashed {
					return errorf("Trashed should be false")
				}
				return nil
			},
		},
		{
			name: "file with size",
			input: &drive.File{
				Id:   "file123",
				Name: "test.txt",
				Size: 1024,
			},
			validate: func(file *types.DriveFile) error {
				if file.Size != 1024 {
					return errorf("Size mismatch")
				}
				return nil
			},
		},
		{
			name: "file with parents",
			input: &drive.File{
				Id:      "file123",
				Name:    "test.txt",
				Parents: []string{"parent1", "parent2"},
			},
			validate: func(file *types.DriveFile) error {
				if len(file.Parents) != 2 {
					return errorf("Parents count mismatch")
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertDriveFile(tt.input)
			if err := tt.validate(result); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestConvertSharedDrive(t *testing.T) {
	tests := []struct {
		name     string
		input    *drive.Drive
		validate func(*types.SharedDrive) error
	}{
		{
			name: "drive with basic properties",
			input: &drive.Drive{
				Id:   "drive123",
				Name: "Shared Drive",
			},
			validate: func(drive *types.SharedDrive) error {
				if drive.ID != "drive123" {
					return errorf("ID mismatch")
				}
				if drive.Name != "Shared Drive" {
					return errorf("Name mismatch")
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertSharedDrive(tt.input)
			if err := tt.validate(result); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestConvertChangeList(t *testing.T) {
	tests := []struct {
		name     string
		input    *drive.ChangeList
		validate func(*types.ChangeList) error
	}{
		{
			name: "change list with changes",
			input: &drive.ChangeList{
				Changes: []*drive.Change{
					{
						Type:   "file",
						FileId: "file1",
					},
					{
						Type:   "file",
						FileId: "file2",
					},
				},
				NextPageToken:     "token123",
				NewStartPageToken: "start456",
			},
			validate: func(list *types.ChangeList) error {
				if len(list.Changes) != 2 {
					return errorf("Changes count mismatch")
				}
				if list.NextPageToken != "token123" {
					return errorf("NextPageToken mismatch")
				}
				if list.NewStartPageToken != "start456" {
					return errorf("NewStartPageToken mismatch")
				}
				return nil
			},
		},
		{
			name: "empty change list",
			input: &drive.ChangeList{
				Changes: []*drive.Change{},
			},
			validate: func(list *types.ChangeList) error {
				if len(list.Changes) != 0 {
					return errorf("Changes count mismatch")
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertChangeList(tt.input)
			if err := tt.validate(result); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestConvertChannel(t *testing.T) {
	tests := []struct {
		name     string
		input    *drive.Channel
		validate func(*types.Channel) error
	}{
		{
			name: "channel with basic properties",
			input: &drive.Channel{
				Id:         "channel123",
				ResourceId: "resource456",
				Type:       "web_hook",
				Address:    "https://example.com/webhook",
			},
			validate: func(channel *types.Channel) error {
				if channel.ID != "channel123" {
					return errorf("ID mismatch")
				}
				if channel.ResourceID != "resource456" {
					return errorf("ResourceID mismatch")
				}
				if channel.Type != "web_hook" {
					return errorf("Type mismatch")
				}
				if channel.Address != "https://example.com/webhook" {
					return errorf("Address mismatch")
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertChannel(tt.input)
			if err := tt.validate(result); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestParseTime(t *testing.T) {
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
					return errorf("year mismatch")
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseTime(tt.input)
			if (err != nil) != tt.shouldErr {
				t.Errorf("parseTime(%q) error = %v, shouldErr %v", tt.input, err, tt.shouldErr)
			}
			if !tt.shouldErr {
				if err := tt.validate(result); err != nil {
					t.Error(err)
				}
			}
		})
	}
}

func errorf(msg string, args ...interface{}) error {
	return &testError{msg: msg}
}

type testError struct {
	msg string
}

func (e *testError) Error() string {
	return e.msg
}
