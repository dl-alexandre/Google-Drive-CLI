package files

import (
	"testing"

	"github.com/dl-alexandre/gdrive/internal/api"
	"github.com/dl-alexandre/gdrive/internal/types"
	"google.golang.org/api/drive/v3"
)

func TestManager_Creation(t *testing.T) {
	// Test that manager can be created
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

func TestConvertDriveFile(t *testing.T) {
	// Test conversion from Drive API file to internal type
	driveFile := &drive.File{
		Id:           "file123",
		Name:         "test.txt",
		MimeType:     "text/plain",
		Size:         1024,
		CreatedTime:  "2024-01-01T00:00:00Z",
		ModifiedTime: "2024-01-02T00:00:00Z",
		Parents:      []string{"parent1", "parent2"},
		ResourceKey:  "key123",
		WebViewLink:  "https://drive.google.com/file/d/file123/view",
		Trashed:      false,
		Capabilities: &drive.FileCapabilities{
			CanDownload:      true,
			CanEdit:          true,
			CanShare:         true,
			CanDelete:        true,
			CanTrash:         true,
			CanReadRevisions: true,
		},
		ExportLinks: map[string]string{
			"application/pdf": "https://export.link/pdf",
		},
	}

	converted := convertDriveFile(driveFile)

	if converted.ID != driveFile.Id {
		t.Errorf("ID mismatch: got %s, want %s", converted.ID, driveFile.Id)
	}

	if converted.Name != driveFile.Name {
		t.Errorf("Name mismatch: got %s, want %s", converted.Name, driveFile.Name)
	}

	if converted.MimeType != driveFile.MimeType {
		t.Errorf("MimeType mismatch: got %s, want %s", converted.MimeType, driveFile.MimeType)
	}

	if converted.Size != driveFile.Size {
		t.Errorf("Size mismatch: got %d, want %d", converted.Size, driveFile.Size)
	}

	if len(converted.Parents) != len(driveFile.Parents) {
		t.Errorf("Parents length mismatch: got %d, want %d", len(converted.Parents), len(driveFile.Parents))
	}

	if converted.ResourceKey != driveFile.ResourceKey {
		t.Errorf("ResourceKey mismatch: got %s, want %s", converted.ResourceKey, driveFile.ResourceKey)
	}

	if converted.Trashed != driveFile.Trashed {
		t.Errorf("Trashed mismatch: got %v, want %v", converted.Trashed, driveFile.Trashed)
	}

	if converted.Capabilities == nil {
		t.Error("Capabilities should not be nil")
	} else {
		if converted.Capabilities.CanDownload != driveFile.Capabilities.CanDownload {
			t.Errorf("CanDownload mismatch: got %v, want %v", converted.Capabilities.CanDownload, driveFile.Capabilities.CanDownload)
		}
		if converted.Capabilities.CanEdit != driveFile.Capabilities.CanEdit {
			t.Errorf("CanEdit mismatch: got %v, want %v", converted.Capabilities.CanEdit, driveFile.Capabilities.CanEdit)
		}
	}

	if len(converted.ExportLinks) != len(driveFile.ExportLinks) {
		t.Errorf("ExportLinks length mismatch: got %d, want %d", len(converted.ExportLinks), len(driveFile.ExportLinks))
	}
}

func TestConvertDriveFile_NilCapabilities(t *testing.T) {
	// Test conversion when capabilities are nil
	driveFile := &drive.File{
		Id:           "file123",
		Name:         "test.txt",
		MimeType:     "text/plain",
		Capabilities: nil,
	}

	converted := convertDriveFile(driveFile)

	if converted.Capabilities != nil {
		t.Error("Capabilities should be nil when source has nil capabilities")
	}
}

func TestListOptions_Defaults(t *testing.T) {
	// Test that ListOptions has sensible defaults
	opts := ListOptions{}

	if opts.ParentID != "" {
		t.Errorf("Default ParentID should be empty, got %s", opts.ParentID)
	}

	if opts.IncludeTrashed != false {
		t.Error("Default IncludeTrashed should be false")
	}
}

func TestUploadOptions_Defaults(t *testing.T) {
	// Test that UploadOptions has sensible defaults
	opts := UploadOptions{}

	if opts.Convert != false {
		t.Error("Default Convert should be false")
	}

	if opts.PinRevision != false {
		t.Error("Default PinRevision should be false")
	}
}

func TestDownloadOptions_Defaults(t *testing.T) {
	// Test that DownloadOptions has sensible defaults
	opts := DownloadOptions{}

	if opts.Wait != false {
		t.Error("Default Wait should be false")
	}

	if opts.Timeout != 0 {
		t.Errorf("Default Timeout should be 0, got %d", opts.Timeout)
	}

	if opts.PollInterval != 0 {
		t.Errorf("Default PollInterval should be 0, got %d", opts.PollInterval)
	}
}

// TestSelectUploadType_Integration tests upload type selection with realistic scenarios
func TestSelectUploadType_Integration(t *testing.T) {
	tests := []struct {
		name         string
		scenario     string
		size         int64
		metadata     *drive.File
		expectedType string
	}{
		{
			name:         "Quick file upload without metadata",
			scenario:     "User uploads small file without specifying name or parent",
			size:         1024,
			metadata:     &drive.File{},
			expectedType: "simple",
		},
		{
			name:         "File upload to specific folder",
			scenario:     "User uploads file to a specific parent folder",
			size:         1024,
			metadata:     &drive.File{Parents: []string{"folderID"}},
			expectedType: "multipart",
		},
		{
			name:         "Large file upload",
			scenario:     "User uploads 10MB video file",
			size:         10 * 1024 * 1024,
			metadata:     &drive.File{Name: "video.mp4", MimeType: "video/mp4"},
			expectedType: "resumable",
		},
		{
			name:         "Document with metadata",
			scenario:     "User uploads document with custom name",
			size:         500 * 1024,
			metadata:     &drive.File{Name: "report.pdf", MimeType: "application/pdf"},
			expectedType: "multipart",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := selectUploadType(tt.size, tt.metadata)
			if got != tt.expectedType {
				t.Errorf("Scenario: %s\nselectUploadType(%d, %+v) = %s, want %s",
					tt.scenario, tt.size, tt.metadata, got, tt.expectedType)
			}
		})
	}
}

// TestRequestContext_Construction tests that request contexts are properly constructed
func TestRequestContext_Construction(t *testing.T) {
	ctx := &types.RequestContext{
		Profile:           "default",
		DriveID:           "drive123",
		InvolvedFileIDs:   []string{"file1", "file2"},
		InvolvedParentIDs: []string{"parent1"},
		RequestType:       types.RequestTypeMutation,
		TraceID:           "trace123",
	}

	if ctx.Profile != "default" {
		t.Errorf("Profile = %s, want default", ctx.Profile)
	}

	if ctx.DriveID != "drive123" {
		t.Errorf("DriveID = %s, want drive123", ctx.DriveID)
	}

	if len(ctx.InvolvedFileIDs) != 2 {
		t.Errorf("InvolvedFileIDs length = %d, want 2", len(ctx.InvolvedFileIDs))
	}

	if len(ctx.InvolvedParentIDs) != 1 {
		t.Errorf("InvolvedParentIDs length = %d, want 1", len(ctx.InvolvedParentIDs))
	}
}
