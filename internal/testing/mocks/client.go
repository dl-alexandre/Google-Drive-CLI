package mocks

import (
	"context"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/logging"
	"github.com/dl-alexandre/gdrv/internal/types"
	"google.golang.org/api/drive/v3"
)

// MockClient is a mock implementation of api.Client for testing
type MockClient struct {
	service        *MockDriveService
	resourceKeyMgr *api.ResourceKeyManager
	logger         logging.Logger
}

// NewMockClient creates a new mock client
func NewMockClient() *MockClient {
	return &MockClient{
		service:        NewMockDriveService(),
		resourceKeyMgr: api.NewResourceKeyManager(),
		logger:         logging.NewNoOpLogger(),
	}
}

// Service returns the mock Drive service
func (c *MockClient) Service() *drive.Service {
	// Note: This returns nil because we mock at a lower level
	// Tests should use the mock service directly via GetMockService()
	return nil
}

// GetMockService returns the underlying mock service
func (c *MockClient) GetMockService() *MockDriveService {
	return c.service
}

// ResourceKeys returns the resource key manager
func (c *MockClient) ResourceKeys() *api.ResourceKeyManager {
	return c.resourceKeyMgr
}

// MockDriveService mocks the Google Drive service
type MockDriveService struct {
	Files       *MockFilesService
	Permissions *MockPermissionsService
	Drives      *MockDrivesService
}

// NewMockDriveService creates a new mock Drive service
func NewMockDriveService() *MockDriveService {
	return &MockDriveService{
		Files:       &MockFilesService{},
		Permissions: &MockPermissionsService{},
		Drives:      &MockDrivesService{},
	}
}

// MockFilesService mocks the Files service
type MockFilesService struct {
	GetFunc        func(fileID string) (*drive.File, error)
	ListFunc       func() (*drive.FileList, error)
	CreateFunc     func(file *drive.File) (*drive.File, error)
	UpdateFunc     func(fileID string, file *drive.File) (*drive.File, error)
	DeleteFunc     func(fileID string) error
	CopyFunc       func(fileID string, file *drive.File) (*drive.File, error)
	ExportFunc     func(fileID string, mimeType string) ([]byte, error)
	EmptyTrashFunc func() error
}

// Get mocks getting a file
func (m *MockFilesService) Get(fileID string) (*drive.File, error) {
	if m.GetFunc != nil {
		return m.GetFunc(fileID)
	}
	return &drive.File{
		Id:       fileID,
		Name:     "mock-file.txt",
		MimeType: "text/plain",
	}, nil
}

// List mocks listing files
func (m *MockFilesService) List() (*drive.FileList, error) {
	if m.ListFunc != nil {
		return m.ListFunc()
	}
	return &drive.FileList{
		Files: []*drive.File{
			{Id: "file1", Name: "file1.txt", MimeType: "text/plain"},
			{Id: "file2", Name: "file2.txt", MimeType: "text/plain"},
		},
		NextPageToken: "",
	}, nil
}

// Create mocks creating a file
func (m *MockFilesService) Create(file *drive.File) (*drive.File, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(file)
	}
	// Return the file with an ID assigned
	file.Id = "new-file-id"
	return file, nil
}

// Update mocks updating a file
func (m *MockFilesService) Update(fileID string, file *drive.File) (*drive.File, error) {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(fileID, file)
	}
	file.Id = fileID
	return file, nil
}

// Delete mocks deleting a file
func (m *MockFilesService) Delete(fileID string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(fileID)
	}
	return nil
}

// Copy mocks copying a file
func (m *MockFilesService) Copy(fileID string, file *drive.File) (*drive.File, error) {
	if m.CopyFunc != nil {
		return m.CopyFunc(fileID, file)
	}
	file.Id = "copied-file-id"
	return file, nil
}

// Export mocks exporting a file
func (m *MockFilesService) Export(fileID string, mimeType string) ([]byte, error) {
	if m.ExportFunc != nil {
		return m.ExportFunc(fileID, mimeType)
	}
	return []byte("exported content"), nil
}

// EmptyTrash mocks emptying trash
func (m *MockFilesService) EmptyTrash() error {
	if m.EmptyTrashFunc != nil {
		return m.EmptyTrashFunc()
	}
	return nil
}

// MockPermissionsService mocks the Permissions service
type MockPermissionsService struct {
	ListFunc   func(fileID string) ([]*drive.Permission, error)
	GetFunc    func(fileID, permissionID string) (*drive.Permission, error)
	CreateFunc func(fileID string, perm *drive.Permission) (*drive.Permission, error)
	UpdateFunc func(fileID, permissionID string, perm *drive.Permission) (*drive.Permission, error)
	DeleteFunc func(fileID, permissionID string) error
}

// List mocks listing permissions
func (m *MockPermissionsService) List(fileID string) ([]*drive.Permission, error) {
	if m.ListFunc != nil {
		return m.ListFunc(fileID)
	}
	return []*drive.Permission{
		{Id: "perm1", Type: "user", Role: "reader", EmailAddress: "user@example.com"},
	}, nil
}

// Get mocks getting a permission
func (m *MockPermissionsService) Get(fileID, permissionID string) (*drive.Permission, error) {
	if m.GetFunc != nil {
		return m.GetFunc(fileID, permissionID)
	}
	return &drive.Permission{
		Id:           permissionID,
		Type:         "user",
		Role:         "reader",
		EmailAddress: "user@example.com",
	}, nil
}

// Create mocks creating a permission
func (m *MockPermissionsService) Create(fileID string, perm *drive.Permission) (*drive.Permission, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(fileID, perm)
	}
	perm.Id = "new-perm-id"
	return perm, nil
}

// Update mocks updating a permission
func (m *MockPermissionsService) Update(fileID, permissionID string, perm *drive.Permission) (*drive.Permission, error) {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(fileID, permissionID, perm)
	}
	perm.Id = permissionID
	return perm, nil
}

// Delete mocks deleting a permission
func (m *MockPermissionsService) Delete(fileID, permissionID string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(fileID, permissionID)
	}
	return nil
}

// MockDrivesService mocks the Drives (Shared Drives) service
type MockDrivesService struct {
	ListFunc func() ([]*drive.Drive, error)
	GetFunc  func(driveID string) (*drive.Drive, error)
}

// List mocks listing drives
func (m *MockDrivesService) List() ([]*drive.Drive, error) {
	if m.ListFunc != nil {
		return m.ListFunc()
	}
	return []*drive.Drive{
		{Id: "drive1", Name: "Shared Drive 1"},
	}, nil
}

// Get mocks getting a drive
func (m *MockDrivesService) Get(driveID string) (*drive.Drive, error) {
	if m.GetFunc != nil {
		return m.GetFunc(driveID)
	}
	return &drive.Drive{
		Id:   driveID,
		Name: "Mock Shared Drive",
	}, nil
}

// ExecuteWithRetryFunc is a helper type for mocking ExecuteWithRetry
type ExecuteWithRetryFunc func(ctx context.Context, client *api.Client, reqCtx *types.RequestContext, fn interface{}) (interface{}, error)

// MockExecutor helps mock api.ExecuteWithRetry calls
type MockExecutor struct {
	ExecuteFunc ExecuteWithRetryFunc
	CallCount   int
	LastReqCtx  *types.RequestContext
}

// Execute mocks the ExecuteWithRetry function
func (m *MockExecutor) Execute(ctx context.Context, client *api.Client, reqCtx *types.RequestContext, fn interface{}) (interface{}, error) {
	m.CallCount++
	m.LastReqCtx = reqCtx

	if m.ExecuteFunc != nil {
		return m.ExecuteFunc(ctx, client, reqCtx, fn)
	}

	// Default: execute the function directly
	switch f := fn.(type) {
	case func() (*drive.File, error):
		return f()
	case func() (*drive.FileList, error):
		return f()
	case func() error:
		return nil, f()
	default:
		return nil, nil
	}
}

// Reset resets the executor state
func (m *MockExecutor) Reset() {
	m.CallCount = 0
	m.LastReqCtx = nil
}
