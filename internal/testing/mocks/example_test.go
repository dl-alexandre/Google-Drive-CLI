package mocks_test

import (
	"errors"
	"testing"

	"github.com/dl-alexandre/gdrv/internal/testing/mocks"
	testhelpers "github.com/dl-alexandre/gdrv/internal/testing"
	"google.golang.org/api/drive/v3"
)

// TestMockClient shows how to use MockClient in tests
func TestMockClient(t *testing.T) {
	// Create a mock client
	client := mocks.NewMockClient()
	mockService := client.GetMockService()

	// Configure mock behavior
	mockService.Files.GetFunc = func(fileID string) (*drive.File, error) {
		return &drive.File{
			Id:       fileID,
			Name:     "test-file.txt",
			MimeType: "text/plain",
			Size:     1024,
		}, nil
	}

	// Now use the mock in your tests
	file, err := mockService.Files.Get("test-file-id")
	testhelpers.AssertNoError(t, err, "getting file")
	testhelpers.AssertEqual(t, file.Name, "test-file.txt", "file name")
}

// Example: Testing file operations with mocks
func TestMockFilesService_Get(t *testing.T) {
	mockService := mocks.NewMockDriveService()

	// Test default behavior
	file, err := mockService.Files.Get("file123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if file.Id != "file123" {
		t.Errorf("file ID = %s, want file123", file.Id)
	}

	// Test custom behavior
	mockService.Files.GetFunc = func(fileID string) (*drive.File, error) {
		if fileID == "not-found" {
			return nil, errors.New("file not found")
		}
		return &drive.File{
			Id:       fileID,
			Name:     "custom-file.txt",
			MimeType: "text/plain",
		}, nil
	}

	// Test error case
	_, err = mockService.Files.Get("not-found")
	if err == nil {
		t.Error("expected error for not-found file")
	}

	// Test success case with custom data
	file, err = mockService.Files.Get("custom-id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if file.Name != "custom-file.txt" {
		t.Errorf("file name = %s, want custom-file.txt", file.Name)
	}
}

// Example: Testing with helper functions
func TestWithHelpers(t *testing.T) {
	mockService := mocks.NewMockDriveService()

	// Use test helpers to create test data
	testFile := testhelpers.TestFile("file1", "document.pdf", "application/pdf")
	testFolder := testhelpers.TestFolder("folder1", "My Folder")

	mockService.Files.GetFunc = func(fileID string) (*drive.File, error) {
		switch fileID {
		case "file1":
			return testFile, nil
		case "folder1":
			return testFolder, nil
		default:
			return nil, errors.New("not found")
		}
	}

	// Test file retrieval
	file, err := mockService.Files.Get("file1")
	testhelpers.AssertNoError(t, err, "getting file1")
	testhelpers.AssertEqual(t, file.Name, "document.pdf", "file name")

	// Test folder retrieval
	folder, err := mockService.Files.Get("folder1")
	testhelpers.AssertNoError(t, err, "getting folder1")
	testhelpers.AssertEqual(t, folder.MimeType, "application/vnd.google-apps.folder", "folder mime type")
}

// Example: Testing list operations
func TestMockFilesService_List(t *testing.T) {
	mockService := mocks.NewMockDriveService()

	// Test default behavior
	fileList, err := mockService.Files.List()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(fileList.Files) != 2 {
		t.Errorf("file list length = %d, want 2", len(fileList.Files))
	}

	// Test custom list with pagination
	mockService.Files.ListFunc = func() (*drive.FileList, error) {
		return &drive.FileList{
			Files: []*drive.File{
				testhelpers.TestFile("f1", "file1.txt", "text/plain"),
				testhelpers.TestFile("f2", "file2.txt", "text/plain"),
				testhelpers.TestFile("f3", "file3.txt", "text/plain"),
			},
			NextPageToken: "next-page-token",
		}, nil
	}

	fileList, err = mockService.Files.List()
	testhelpers.AssertNoError(t, err, "listing files")
	testhelpers.AssertEqual(t, len(fileList.Files), 3, "file count")
	testhelpers.AssertEqual(t, fileList.NextPageToken, "next-page-token", "page token")
}

// Example: Testing permissions
func TestMockPermissionsService(t *testing.T) {
	mockService := mocks.NewMockDriveService()

	// Configure mock permissions
	mockService.Permissions.ListFunc = func(fileID string) ([]*drive.Permission, error) {
		return []*drive.Permission{
			testhelpers.TestPermission("perm1", "user", "reader", "user1@example.com"),
			testhelpers.TestPermission("perm2", "user", "writer", "user2@example.com"),
		}, nil
	}

	// Test listing permissions
	perms, err := mockService.Permissions.List("file123")
	testhelpers.AssertNoError(t, err, "listing permissions")
	testhelpers.AssertEqual(t, len(perms), 2, "permission count")
	testhelpers.AssertEqual(t, perms[0].Role, "reader", "first permission role")
}

// Example: Testing create operations
func TestMockFilesService_Create(t *testing.T) {
	mockService := mocks.NewMockDriveService()

	// Test default create behavior
	newFile := &drive.File{
		Name:     "new-file.txt",
		MimeType: "text/plain",
	}

	created, err := mockService.Files.Create(newFile)
	testhelpers.AssertNoError(t, err, "creating file")
	testhelpers.AssertEqual(t, created.Id, "new-file-id", "created file ID")
	testhelpers.AssertEqual(t, created.Name, "new-file.txt", "created file name")

	// Test custom create with validation
	mockService.Files.CreateFunc = func(file *drive.File) (*drive.File, error) {
		if file.Name == "" {
			return nil, errors.New("file name is required")
		}
		file.Id = "custom-id-" + file.Name
		return file, nil
	}

	// Test error case
	_, err = mockService.Files.Create(&drive.File{})
	testhelpers.AssertError(t, err, "creating file without name")

	// Test success case
	created, err = mockService.Files.Create(&drive.File{Name: "test.txt"})
	testhelpers.AssertNoError(t, err, "creating file with name")
	testhelpers.AssertEqual(t, created.Id, "custom-id-test.txt", "custom ID generation")
}

// Example: Testing error scenarios
func TestMockFilesService_ErrorScenarios(t *testing.T) {
	tests := []struct {
		name      string
		fileID    string
		setupFunc func(*mocks.MockFilesService)
		wantError bool
	}{
		{
			name:   "file not found",
			fileID: "not-found",
			setupFunc: func(m *mocks.MockFilesService) {
				m.GetFunc = func(fileID string) (*drive.File, error) {
					return nil, errors.New("404: File not found")
				}
			},
			wantError: true,
		},
		{
			name:   "permission denied",
			fileID: "forbidden",
			setupFunc: func(m *mocks.MockFilesService) {
				m.GetFunc = func(fileID string) (*drive.File, error) {
					return nil, errors.New("403: Permission denied")
				}
			},
			wantError: true,
		},
		{
			name:   "success",
			fileID: "valid-file",
			setupFunc: func(m *mocks.MockFilesService) {
				m.GetFunc = func(fileID string) (*drive.File, error) {
					return testhelpers.TestFile(fileID, "file.txt", "text/plain"), nil
				}
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := mocks.NewMockDriveService()
			tt.setupFunc(mockService.Files)

			_, err := mockService.Files.Get(tt.fileID)
			if tt.wantError {
				testhelpers.AssertError(t, err, tt.name)
			} else {
				testhelpers.AssertNoError(t, err, tt.name)
			}
		})
	}
}
