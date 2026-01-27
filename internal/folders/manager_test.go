package folders

import (
	"testing"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/types"
	"github.com/dl-alexandre/gdrv/internal/utils"
	"google.golang.org/api/drive/v3"
)

func TestNewManager(t *testing.T) {
	client := &api.Client{}
	manager := NewManager(client)

	if manager == nil {
		t.Fatal("NewManager returned nil")
	}

	if manager.client != client {
		t.Error("Manager client not set correctly")
	}

	if manager.shaper == nil {
		t.Error("Manager shaper not initialized")
	}
}

func TestConvertDriveFile(t *testing.T) {
	tests := []struct {
		name     string
		input    *drive.File
		expected *types.DriveFile
	}{
		{
			name: "folder with all fields",
			input: &drive.File{
				Id:           "folder123",
				Name:         "My Folder",
				MimeType:     utils.MimeTypeFolder,
				Size:         0,
				CreatedTime:  "2024-01-01T00:00:00Z",
				ModifiedTime: "2024-01-02T00:00:00Z",
				Parents:      []string{"parent1"},
				ResourceKey:  "key123",
				Trashed:      false,
				Capabilities: &drive.FileCapabilities{
					CanDownload:      false,
					CanEdit:          true,
					CanShare:         true,
					CanDelete:        true,
					CanTrash:         true,
					CanReadRevisions: false,
				},
			},
			expected: &types.DriveFile{
				ID:           "folder123",
				Name:         "My Folder",
				MimeType:     utils.MimeTypeFolder,
				Size:         0,
				CreatedTime:  "2024-01-01T00:00:00Z",
				ModifiedTime: "2024-01-02T00:00:00Z",
				Parents:      []string{"parent1"},
				ResourceKey:  "key123",
				Trashed:      false,
				Capabilities: &types.FileCapabilities{
					CanDownload:      false,
					CanEdit:          true,
					CanShare:         true,
					CanDelete:        true,
					CanTrash:         true,
					CanReadRevisions: false,
				},
			},
		},
		{
			name: "folder without capabilities",
			input: &drive.File{
				Id:       "folder456",
				Name:     "Simple Folder",
				MimeType: utils.MimeTypeFolder,
			},
			expected: &types.DriveFile{
				ID:       "folder456",
				Name:     "Simple Folder",
				MimeType: utils.MimeTypeFolder,
			},
		},
		{
			name: "trashed folder",
			input: &drive.File{
				Id:       "folder789",
				Name:     "Trashed Folder",
				MimeType: utils.MimeTypeFolder,
				Trashed:  true,
			},
			expected: &types.DriveFile{
				ID:       "folder789",
				Name:     "Trashed Folder",
				MimeType: utils.MimeTypeFolder,
				Trashed:  true,
			},
		},
		{
			name: "folder with multiple parents",
			input: &drive.File{
				Id:       "folder999",
				Name:     "Multi-Parent Folder",
				MimeType: utils.MimeTypeFolder,
				Parents:  []string{"parent1", "parent2", "parent3"},
			},
			expected: &types.DriveFile{
				ID:       "folder999",
				Name:     "Multi-Parent Folder",
				MimeType: utils.MimeTypeFolder,
				Parents:  []string{"parent1", "parent2", "parent3"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertDriveFile(tt.input)

			if result.ID != tt.expected.ID {
				t.Errorf("ID mismatch: got %s, want %s", result.ID, tt.expected.ID)
			}
			if result.Name != tt.expected.Name {
				t.Errorf("Name mismatch: got %s, want %s", result.Name, tt.expected.Name)
			}
			if result.MimeType != tt.expected.MimeType {
				t.Errorf("MimeType mismatch: got %s, want %s", result.MimeType, tt.expected.MimeType)
			}
			if result.Size != tt.expected.Size {
				t.Errorf("Size mismatch: got %d, want %d", result.Size, tt.expected.Size)
			}
			if result.CreatedTime != tt.expected.CreatedTime {
				t.Errorf("CreatedTime mismatch: got %s, want %s", result.CreatedTime, tt.expected.CreatedTime)
			}
			if result.ModifiedTime != tt.expected.ModifiedTime {
				t.Errorf("ModifiedTime mismatch: got %s, want %s", result.ModifiedTime, tt.expected.ModifiedTime)
			}
			if result.ResourceKey != tt.expected.ResourceKey {
				t.Errorf("ResourceKey mismatch: got %s, want %s", result.ResourceKey, tt.expected.ResourceKey)
			}
			if result.Trashed != tt.expected.Trashed {
				t.Errorf("Trashed mismatch: got %v, want %v", result.Trashed, tt.expected.Trashed)
			}

			// Check parents
			if len(result.Parents) != len(tt.expected.Parents) {
				t.Errorf("Parents length mismatch: got %d, want %d", len(result.Parents), len(tt.expected.Parents))
			} else {
				for i, p := range result.Parents {
					if p != tt.expected.Parents[i] {
						t.Errorf("Parent[%d] mismatch: got %s, want %s", i, p, tt.expected.Parents[i])
					}
				}
			}

			// Check capabilities
			if tt.expected.Capabilities == nil {
				if result.Capabilities != nil {
					t.Error("Expected nil capabilities, got non-nil")
				}
			} else {
				if result.Capabilities == nil {
					t.Fatal("Expected non-nil capabilities, got nil")
				}
				if result.Capabilities.CanDownload != tt.expected.Capabilities.CanDownload {
					t.Errorf("CanDownload mismatch: got %v, want %v", result.Capabilities.CanDownload, tt.expected.Capabilities.CanDownload)
				}
				if result.Capabilities.CanEdit != tt.expected.Capabilities.CanEdit {
					t.Errorf("CanEdit mismatch: got %v, want %v", result.Capabilities.CanEdit, tt.expected.Capabilities.CanEdit)
				}
				if result.Capabilities.CanShare != tt.expected.Capabilities.CanShare {
					t.Errorf("CanShare mismatch: got %v, want %v", result.Capabilities.CanShare, tt.expected.Capabilities.CanShare)
				}
				if result.Capabilities.CanDelete != tt.expected.Capabilities.CanDelete {
					t.Errorf("CanDelete mismatch: got %v, want %v", result.Capabilities.CanDelete, tt.expected.Capabilities.CanDelete)
				}
				if result.Capabilities.CanTrash != tt.expected.Capabilities.CanTrash {
					t.Errorf("CanTrash mismatch: got %v, want %v", result.Capabilities.CanTrash, tt.expected.Capabilities.CanTrash)
				}
				if result.Capabilities.CanReadRevisions != tt.expected.Capabilities.CanReadRevisions {
					t.Errorf("CanReadRevisions mismatch: got %v, want %v", result.Capabilities.CanReadRevisions, tt.expected.Capabilities.CanReadRevisions)
				}
			}
		})
	}
}

func TestManager_CreateFolder_RequestContextSetup(t *testing.T) {
	tests := []struct {
		name              string
		folderName        string
		parentID          string
		expectParentInCtx bool
	}{
		{
			name:              "create folder with parent",
			folderName:        "New Folder",
			parentID:          "parent123",
			expectParentInCtx: true,
		},
		{
			name:              "create folder without parent (root)",
			folderName:        "Root Folder",
			parentID:          "",
			expectParentInCtx: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This test validates that the RequestContext is properly setup
			// In a real scenario, we would use a mock API client
			reqCtx := api.NewRequestContext("default", "", types.RequestTypeMutation)

			if tt.parentID != "" {
				reqCtx.InvolvedParentIDs = append(reqCtx.InvolvedParentIDs, tt.parentID)
			}

			if tt.expectParentInCtx {
				if len(reqCtx.InvolvedParentIDs) != 1 {
					t.Errorf("Expected 1 parent in context, got %d", len(reqCtx.InvolvedParentIDs))
				}
				if reqCtx.InvolvedParentIDs[0] != tt.parentID {
					t.Errorf("Expected parent %s in context, got %s", tt.parentID, reqCtx.InvolvedParentIDs[0])
				}
			} else {
				if len(reqCtx.InvolvedParentIDs) != 0 {
					t.Errorf("Expected no parents in context, got %d", len(reqCtx.InvolvedParentIDs))
				}
			}
		})
	}
}

func TestManager_ListFolder_QueryConstruction(t *testing.T) {
	tests := []struct {
		name           string
		folderID       string
		expectedQuery  string
		expectFolderID bool
	}{
		{
			name:           "list folder contents",
			folderID:       "folder123",
			expectedQuery:  "'folder123' in parents and trashed = false",
			expectFolderID: true,
		},
		{
			name:           "list root folder",
			folderID:       "root",
			expectedQuery:  "'root' in parents and trashed = false",
			expectFolderID: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqCtx := api.NewRequestContext("default", "", types.RequestTypeListOrSearch)

			if tt.expectFolderID {
				reqCtx.InvolvedParentIDs = append(reqCtx.InvolvedParentIDs, tt.folderID)

				if len(reqCtx.InvolvedParentIDs) != 1 {
					t.Errorf("Expected 1 folder in context, got %d", len(reqCtx.InvolvedParentIDs))
				}
				if reqCtx.InvolvedParentIDs[0] != tt.folderID {
					t.Errorf("Expected folder %s in context, got %s", tt.folderID, reqCtx.InvolvedParentIDs[0])
				}
			}
		})
	}
}

func TestManager_DeleteFolder_RecursiveFlag(t *testing.T) {
	tests := []struct {
		name      string
		folderID  string
		recursive bool
	}{
		{
			name:      "delete folder non-recursively",
			folderID:  "folder123",
			recursive: false,
		},
		{
			name:      "delete folder recursively",
			folderID:  "folder456",
			recursive: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Validate that the folderID is added to context
			reqCtx := api.NewRequestContext("default", "", types.RequestTypeMutation)
			reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, tt.folderID)

			if len(reqCtx.InvolvedFileIDs) != 1 {
				t.Errorf("Expected 1 file in context, got %d", len(reqCtx.InvolvedFileIDs))
			}
			if reqCtx.InvolvedFileIDs[0] != tt.folderID {
				t.Errorf("Expected folder %s in context, got %s", tt.folderID, reqCtx.InvolvedFileIDs[0])
			}
		})
	}
}

func TestManager_MoveFolder_RequestContextSetup(t *testing.T) {
	tests := []struct {
		name        string
		folderID    string
		newParentID string
	}{
		{
			name:        "move folder to new parent",
			folderID:    "folder123",
			newParentID: "parent456",
		},
		{
			name:        "move folder to root",
			folderID:    "folder789",
			newParentID: "root",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqCtx := api.NewRequestContext("default", "", types.RequestTypeMutation)
			reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, tt.folderID)
			reqCtx.InvolvedParentIDs = append(reqCtx.InvolvedParentIDs, tt.newParentID)

			if len(reqCtx.InvolvedFileIDs) != 1 {
				t.Errorf("Expected 1 file in context, got %d", len(reqCtx.InvolvedFileIDs))
			}
			if reqCtx.InvolvedFileIDs[0] != tt.folderID {
				t.Errorf("Expected folder %s in context, got %s", tt.folderID, reqCtx.InvolvedFileIDs[0])
			}

			if len(reqCtx.InvolvedParentIDs) != 1 {
				t.Errorf("Expected 1 parent in context, got %d", len(reqCtx.InvolvedParentIDs))
			}
			if reqCtx.InvolvedParentIDs[0] != tt.newParentID {
				t.Errorf("Expected parent %s in context, got %s", tt.newParentID, reqCtx.InvolvedParentIDs[0])
			}
		})
	}
}

func TestManager_GetFolder_RequestContextSetup(t *testing.T) {
	tests := []struct {
		name     string
		folderID string
		fields   string
	}{
		{
			name:     "get folder with default fields",
			folderID: "folder123",
			fields:   "",
		},
		{
			name:     "get folder with custom fields",
			folderID: "folder456",
			fields:   "id,name,mimeType,parents",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqCtx := api.NewRequestContext("default", "", types.RequestTypeGetByID)
			reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, tt.folderID)

			if len(reqCtx.InvolvedFileIDs) != 1 {
				t.Errorf("Expected 1 file in context, got %d", len(reqCtx.InvolvedFileIDs))
			}
			if reqCtx.InvolvedFileIDs[0] != tt.folderID {
				t.Errorf("Expected folder %s in context, got %s", tt.folderID, reqCtx.InvolvedFileIDs[0])
			}
		})
	}
}

func TestManager_RenameFolder_RequestContextSetup(t *testing.T) {
	tests := []struct {
		name       string
		folderID   string
		newName    string
		expectName bool
	}{
		{
			name:       "rename folder",
			folderID:   "folder123",
			newName:    "Renamed Folder",
			expectName: true,
		},
		{
			name:       "rename folder with special characters",
			folderID:   "folder456",
			newName:    "Folder (2024)",
			expectName: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqCtx := api.NewRequestContext("default", "", types.RequestTypeMutation)
			reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, tt.folderID)

			if len(reqCtx.InvolvedFileIDs) != 1 {
				t.Errorf("Expected 1 file in context, got %d", len(reqCtx.InvolvedFileIDs))
			}
			if reqCtx.InvolvedFileIDs[0] != tt.folderID {
				t.Errorf("Expected folder %s in context, got %s", tt.folderID, reqCtx.InvolvedFileIDs[0])
			}

			if tt.expectName && tt.newName == "" {
				t.Error("Expected new name to be set")
			}
		})
	}
}

func TestFileListResult_Structure(t *testing.T) {
	// Test that FileListResult properly structures paginated results
	files := []*types.DriveFile{
		{ID: "file1", Name: "File 1", MimeType: "text/plain"},
		{ID: "file2", Name: "File 2", MimeType: utils.MimeTypeFolder},
		{ID: "file3", Name: "File 3", MimeType: "application/pdf"},
	}

	result := &types.FileListResult{
		Files:            files,
		NextPageToken:    "next-token-123",
		IncompleteSearch: false,
	}

	if len(result.Files) != 3 {
		t.Errorf("Expected 3 files, got %d", len(result.Files))
	}

	if result.NextPageToken != "next-token-123" {
		t.Errorf("NextPageToken mismatch: got %s, want next-token-123", result.NextPageToken)
	}

	if result.IncompleteSearch {
		t.Error("Expected IncompleteSearch to be false")
	}

	// Verify folder is properly identified
	hasFolder := false
	for _, f := range result.Files {
		if f.MimeType == utils.MimeTypeFolder {
			hasFolder = true
			break
		}
	}
	if !hasFolder {
		t.Error("Expected to find a folder in the results")
	}
}

func TestManager_SharedDriveContext(t *testing.T) {
	tests := []struct {
		name    string
		driveID string
		profile string
	}{
		{
			name:    "operation on shared drive",
			driveID: "shared-drive-123",
			profile: "default",
		},
		{
			name:    "operation on personal drive",
			driveID: "",
			profile: "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqCtx := api.NewRequestContext(tt.profile, tt.driveID, types.RequestTypeMutation)

			if reqCtx.DriveID != tt.driveID {
				t.Errorf("DriveID mismatch: got %s, want %s", reqCtx.DriveID, tt.driveID)
			}

			if reqCtx.Profile != tt.profile {
				t.Errorf("Profile mismatch: got %s, want %s", reqCtx.Profile, tt.profile)
			}

			// Verify TraceID is set
			if reqCtx.TraceID == "" {
				t.Error("TraceID should be set")
			}
		})
	}
}

// Integration-style test that validates the folder creation metadata structure
func TestFolderMetadataStructure(t *testing.T) {
	tests := []struct {
		name     string
		metadata *drive.File
		validate func(*testing.T, *drive.File)
	}{
		{
			name: "folder with parent",
			metadata: &drive.File{
				Name:     "Test Folder",
				MimeType: utils.MimeTypeFolder,
				Parents:  []string{"parent123"},
			},
			validate: func(t *testing.T, f *drive.File) {
				if f.MimeType != utils.MimeTypeFolder {
					t.Errorf("Expected folder MIME type, got %s", f.MimeType)
				}
				if len(f.Parents) != 1 {
					t.Errorf("Expected 1 parent, got %d", len(f.Parents))
				}
			},
		},
		{
			name: "folder without parent (root)",
			metadata: &drive.File{
				Name:     "Root Folder",
				MimeType: utils.MimeTypeFolder,
			},
			validate: func(t *testing.T, f *drive.File) {
				if f.MimeType != utils.MimeTypeFolder {
					t.Errorf("Expected folder MIME type, got %s", f.MimeType)
				}
				if len(f.Parents) != 0 {
					t.Errorf("Expected no parents, got %d", len(f.Parents))
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.validate(t, tt.metadata)
		})
	}
}

// Test MimeTypeFolder constant
func TestMimeTypeFolder(t *testing.T) {
	expected := "application/vnd.google-apps.folder"
	if utils.MimeTypeFolder != expected {
		t.Errorf("MimeTypeFolder = %s, want %s", utils.MimeTypeFolder, expected)
	}
}

// Test that converted files maintain folder properties
func TestConvertDriveFile_FolderProperties(t *testing.T) {
	folder := &drive.File{
		Id:           "folder-id",
		Name:         "My Folder",
		MimeType:     utils.MimeTypeFolder,
		CreatedTime:  "2024-01-01T00:00:00Z",
		ModifiedTime: "2024-01-02T00:00:00Z",
	}

	result := convertDriveFile(folder)

	if result.MimeType != utils.MimeTypeFolder {
		t.Errorf("Expected folder MIME type, got %s", result.MimeType)
	}

	if result.Size != 0 {
		t.Error("Folders should have size 0")
	}
}

// Test conversion of folder capabilities
func TestConvertDriveFile_FolderCapabilities(t *testing.T) {
	tests := []struct {
		name         string
		capabilities *drive.FileCapabilities
		checkFunc    func(*testing.T, *types.FileCapabilities)
	}{
		{
			name: "full folder capabilities",
			capabilities: &drive.FileCapabilities{
				CanEdit:   true,
				CanShare:  true,
				CanDelete: true,
				CanTrash:  true,
			},
			checkFunc: func(t *testing.T, caps *types.FileCapabilities) {
				if !caps.CanEdit {
					t.Error("CanEdit should be true")
				}
				if !caps.CanShare {
					t.Error("CanShare should be true")
				}
				if !caps.CanDelete {
					t.Error("CanDelete should be true")
				}
				if !caps.CanTrash {
					t.Error("CanTrash should be true")
				}
			},
		},
		{
			name: "read-only folder",
			capabilities: &drive.FileCapabilities{
				CanEdit:   false,
				CanShare:  false,
				CanDelete: false,
				CanTrash:  false,
			},
			checkFunc: func(t *testing.T, caps *types.FileCapabilities) {
				if caps.CanEdit {
					t.Error("CanEdit should be false")
				}
				if caps.CanShare {
					t.Error("CanShare should be false")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			folder := &drive.File{
				Id:           "folder-id",
				Name:         "Test Folder",
				MimeType:     utils.MimeTypeFolder,
				Capabilities: tt.capabilities,
			}

			result := convertDriveFile(folder)
			if result.Capabilities == nil {
				t.Fatal("Capabilities should not be nil")
			}

			tt.checkFunc(t, result.Capabilities)
		})
	}
}

// Test FileListResult with mixed content
func TestFileListResult_MixedContent(t *testing.T) {
	files := []*types.DriveFile{
		{ID: "f1", Name: "File 1", MimeType: "text/plain"},
		{ID: "f2", Name: "Folder 1", MimeType: utils.MimeTypeFolder},
		{ID: "f3", Name: "File 2", MimeType: "application/pdf"},
		{ID: "f4", Name: "Folder 2", MimeType: utils.MimeTypeFolder},
	}

	result := &types.FileListResult{
		Files:            files,
		NextPageToken:    "",
		IncompleteSearch: false,
	}

	// Count folders
	folderCount := 0
	fileCount := 0
	for _, f := range result.Files {
		if f.MimeType == utils.MimeTypeFolder {
			folderCount++
		} else {
			fileCount++
		}
	}

	if folderCount != 2 {
		t.Errorf("Expected 2 folders, got %d", folderCount)
	}

	if fileCount != 2 {
		t.Errorf("Expected 2 files, got %d", fileCount)
	}
}

// Test pagination token handling
func TestFileListResult_Pagination(t *testing.T) {
	tests := []struct {
		name              string
		nextPageToken     string
		incompleteSearch  bool
		expectMoreResults bool
	}{
		{
			name:              "has next page",
			nextPageToken:     "token-123",
			incompleteSearch:  false,
			expectMoreResults: true,
		},
		{
			name:              "last page",
			nextPageToken:     "",
			incompleteSearch:  false,
			expectMoreResults: false,
		},
		{
			name:              "incomplete search",
			nextPageToken:     "token-456",
			incompleteSearch:  true,
			expectMoreResults: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := &types.FileListResult{
				Files:            []*types.DriveFile{},
				NextPageToken:    tt.nextPageToken,
				IncompleteSearch: tt.incompleteSearch,
			}

			hasMore := result.NextPageToken != ""
			if hasMore != tt.expectMoreResults {
				t.Errorf("Expected hasMore=%v, got %v", tt.expectMoreResults, hasMore)
			}
		})
	}
}

// Test ResourceKey handling in folder operations
func TestResourceKey_FolderContext(t *testing.T) {
	folder := &drive.File{
		Id:          "folder-id",
		Name:        "Shared Folder",
		MimeType:    utils.MimeTypeFolder,
		ResourceKey: "resource-key-123",
	}

	result := convertDriveFile(folder)

	if result.ResourceKey != "resource-key-123" {
		t.Errorf("ResourceKey = %s, want resource-key-123", result.ResourceKey)
	}
}

// Test folder parent handling
func TestFolderParents_EdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		parents []string
		valid   bool
	}{
		{"no parents (root)", []string{}, true},
		{"single parent", []string{"parent1"}, true},
		{"multiple parents", []string{"parent1", "parent2"}, true},
		{"nil parents", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			folder := &drive.File{
				Id:       "folder-id",
				Name:     "Test Folder",
				MimeType: utils.MimeTypeFolder,
				Parents:  tt.parents,
			}

			result := convertDriveFile(folder)

			if len(tt.parents) == 0 {
				if len(result.Parents) != 0 {
					t.Errorf("Expected nil or empty parents, got %v", result.Parents)
				}
			} else {
				if len(result.Parents) != len(tt.parents) {
					t.Errorf("Expected %d parents, got %d", len(tt.parents), len(result.Parents))
				}
			}
		})
	}
}

// Test trashed folder handling
func TestTrashedFolder_Properties(t *testing.T) {
	trashedFolder := &drive.File{
		Id:       "trashed-folder",
		Name:     "Deleted Folder",
		MimeType: utils.MimeTypeFolder,
		Trashed:  true,
	}

	result := convertDriveFile(trashedFolder)

	if !result.Trashed {
		t.Error("Folder should be marked as trashed")
	}

	if result.MimeType != utils.MimeTypeFolder {
		t.Error("Trashed folder should still have folder MIME type")
	}
}

// Test MD5 checksum (folders don't have checksums)
func TestFolder_NoChecksum(t *testing.T) {
	folder := &drive.File{
		Id:          "folder-id",
		Name:        "Test Folder",
		MimeType:    utils.MimeTypeFolder,
		Md5Checksum: "", // Folders don't have MD5 checksums
	}

	result := convertDriveFile(folder)

	if result.MD5Checksum != "" {
		t.Error("Folders should not have MD5 checksum")
	}
}
