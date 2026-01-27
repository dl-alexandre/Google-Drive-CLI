# Testing Infrastructure

This package provides testing utilities and mocks for the Google Drive CLI project.

## Overview

The testing infrastructure consists of:
- **Mocks**: Mock implementations of Drive API services
- **Helpers**: Test helper functions for common testing patterns
- **Examples**: Comprehensive examples showing how to use the mocks

## Quick Start

### Basic Usage

```go
import (
    "testing"
    "github.com/dl-alexandre/gdrv/internal/testing/mocks"
    testhelpers "github.com/dl-alexandre/gdrv/internal/testing"
)

func TestMyFunction(t *testing.T) {
    // Create mock client
    client := mocks.NewMockClient()
    mockService := client.GetMockService()

    // Configure mock behavior
    mockService.Files.GetFunc = func(fileID string) (*drive.File, error) {
        return testhelpers.TestFile(fileID, "test.txt", "text/plain"), nil
    }

    // Use in your test
    file, err := mockService.Files.Get("file123")
    testhelpers.AssertNoError(t, err, "getting file")
    testhelpers.AssertEqual(t, file.Name, "test.txt", "file name")
}
```

## Mock Services

### MockClient

The main mock client that wraps all services:

```go
client := mocks.NewMockClient()
mockService := client.GetMockService()

// Access mock services
mockService.Files        // MockFilesService
mockService.Permissions  // MockPermissionsService
mockService.Drives       // MockDrivesService
```

### MockFilesService

Mock for file operations:

```go
// Default behavior (returns mock data)
file, err := mockService.Files.Get("file123")

// Custom behavior
mockService.Files.GetFunc = func(fileID string) (*drive.File, error) {
    if fileID == "not-found" {
        return nil, errors.New("file not found")
    }
    return &drive.File{Id: fileID, Name: "custom.txt"}, nil
}
```

Available methods:
- `Get(fileID string) (*drive.File, error)`
- `List() (*drive.FileList, error)`
- `Create(file *drive.File) (*drive.File, error)`
- `Update(fileID string, file *drive.File) (*drive.File, error)`
- `Delete(fileID string) error`
- `Copy(fileID string, file *drive.File) (*drive.File, error)`
- `Export(fileID string, mimeType string) ([]byte, error)`
- `EmptyTrash() error`

### MockPermissionsService

Mock for permission operations:

```go
mockService.Permissions.ListFunc = func(fileID string) ([]*drive.Permission, error) {
    return []*drive.Permission{
        {Id: "perm1", Type: "user", Role: "reader"},
    }, nil
}
```

Available methods:
- `List(fileID string) ([]*drive.Permission, error)`
- `Get(fileID, permissionID string) (*drive.Permission, error)`
- `Create(fileID string, perm *drive.Permission) (*drive.Permission, error)`
- `Update(fileID, permissionID string, perm *drive.Permission) (*drive.Permission, error)`
- `Delete(fileID, permissionID string) error`

### MockDrivesService

Mock for Shared Drives operations:

```go
mockService.Drives.ListFunc = func() ([]*drive.Drive, error) {
    return []*drive.Drive{
        {Id: "drive1", Name: "My Shared Drive"},
    }, nil
}
```

Available methods:
- `List() ([]*drive.Drive, error)`
- `Get(driveID string) (*drive.Drive, error)`

## Test Helpers

### Context Helpers

```go
// Create standard test context
ctx := testhelpers.TestContext()

// Create request context
reqCtx := testhelpers.TestRequestContext()

// Request context with files
reqCtx := testhelpers.TestRequestContextWithFiles("file1", "file2")

// Request context with parents
reqCtx := testhelpers.TestRequestContextWithParents("parent1")
```

### Data Helpers

```go
// Create test file
file := testhelpers.TestFile("id123", "file.txt", "text/plain")

// Create test folder
folder := testhelpers.TestFolder("folder123", "My Folder")

// Create file with capabilities
file := testhelpers.TestFileWithCapabilities("id", "file.txt", true, true, false)

// Create test permission
perm := testhelpers.TestPermission("perm1", "user", "reader", "user@example.com")
```

### Assertion Helpers

```go
// Assert no error
testhelpers.AssertNoError(t, err, "optional message")

// Assert error exists
testhelpers.AssertError(t, err, "optional message")

// Assert equality
testhelpers.AssertEqual(t, got, want, "optional message")

// Assert not nil
testhelpers.AssertNotNil(t, value, "optional message")

// Assert nil
testhelpers.AssertNil(t, value, "optional message")
```

## Testing Patterns

### Table-Driven Tests

```go
func TestSomething(t *testing.T) {
    tests := []struct {
        name      string
        fileID    string
        setupFunc func(*mocks.MockFilesService)
        wantError bool
    }{
        {
            name: "success",
            fileID: "file123",
            setupFunc: func(m *mocks.MockFilesService) {
                m.GetFunc = func(fileID string) (*drive.File, error) {
                    return testhelpers.TestFile(fileID, "test.txt", "text/plain"), nil
                }
            },
            wantError: false,
        },
        {
            name: "file not found",
            fileID: "not-found",
            setupFunc: func(m *mocks.MockFilesService) {
                m.GetFunc = func(fileID string) (*drive.File, error) {
                    return nil, errors.New("404: Not Found")
                }
            },
            wantError: true,
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
```

### Testing Manager Methods

```go
func TestFilesManager_Get(t *testing.T) {
    // Create mock client (note: actual api.Client integration needs more work)
    mockClient := mocks.NewMockClient()
    mockService := mockClient.GetMockService()

    // Configure mock
    mockService.Files.GetFunc = func(fileID string) (*drive.File, error) {
        return testhelpers.TestFile(fileID, "test.txt", "text/plain"), nil
    }

    // Test your manager method
    // (This is a simplified example - actual usage depends on your manager implementation)
    // manager := files.NewManager(mockClient)
    // file, err := manager.Get(ctx, reqCtx, "file123", "id,name")
    // testhelpers.AssertNoError(t, err, "getting file")
}
```

### Testing Error Scenarios

```go
func TestErrorHandling(t *testing.T) {
    mockService := mocks.NewMockDriveService()

    mockService.Files.GetFunc = func(fileID string) (*drive.File, error) {
        switch fileID {
        case "not-found":
            return nil, &googleapi.Error{Code: 404, Message: "File not found"}
        case "forbidden":
            return nil, &googleapi.Error{Code: 403, Message: "Permission denied"}
        case "rate-limited":
            return nil, &googleapi.Error{Code: 429, Message: "Rate limit exceeded"}
        default:
            return testhelpers.TestFile(fileID, "file.txt", "text/plain"), nil
        }
    }

    // Test each error case
    _, err := mockService.Files.Get("not-found")
    testhelpers.AssertError(t, err, "not-found case")
}
```

## Examples

See `internal/testing/mocks/example_test.go` for comprehensive examples showing:
- Basic mock usage
- Custom mock behavior
- Error scenarios
- Pagination
- Permissions
- Table-driven tests

## Best Practices

1. **Use Table-Driven Tests**: Test multiple scenarios in a single test function
2. **Configure Mock Behavior**: Override default behavior with custom functions
3. **Use Test Helpers**: Leverage `testhelpers` for common assertions and data creation
4. **Test Error Paths**: Always test both success and failure scenarios
5. **Keep Tests Focused**: Each test should test one thing
6. **Use Subtests**: Use `t.Run()` for better test organization

## Limitations

- Mocks currently work at the service level, not at the full API client level
- Some integration between mocks and `api.ExecuteWithRetry()` may require additional work
- For testing actual retry logic, consider using a mock executor or integration tests

## Future Improvements

- [ ] Add mock for `api.ExecuteWithRetry()` function
- [ ] Add mocks for Admin SDK services
- [ ] Add mocks for Sheets, Docs, Slides services
- [ ] Create test fixtures for common scenarios
- [ ] Add helper for testing pagination
