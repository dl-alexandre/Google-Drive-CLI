package testing

import (
	"context"
	"testing"

	"github.com/dl-alexandre/gdrv/internal/types"
	"google.golang.org/api/drive/v3"
)

// TestContext creates a standard test context
func TestContext() context.Context {
	return context.Background()
}

// TestRequestContext creates a standard request context for testing
func TestRequestContext() *types.RequestContext {
	return &types.RequestContext{
		Profile:           "test-profile",
		DriveID:           "",
		InvolvedFileIDs:   []string{},
		InvolvedParentIDs: []string{},
		RequestType:       types.RequestTypeListOrSearch,
		TraceID:           "test-trace-id",
	}
}

// TestRequestContextWithFiles creates a request context with file IDs
func TestRequestContextWithFiles(fileIDs ...string) *types.RequestContext {
	ctx := TestRequestContext()
	ctx.InvolvedFileIDs = fileIDs
	return ctx
}

// TestRequestContextWithParents creates a request context with parent IDs
func TestRequestContextWithParents(parentIDs ...string) *types.RequestContext {
	ctx := TestRequestContext()
	ctx.InvolvedParentIDs = parentIDs
	return ctx
}

// TestFile creates a mock Drive file for testing
func TestFile(id, name, mimeType string) *drive.File {
	return &drive.File{
		Id:       id,
		Name:     name,
		MimeType: mimeType,
		Size:     1024,
	}
}

// TestFolder creates a mock Drive folder for testing
func TestFolder(id, name string) *drive.File {
	return &drive.File{
		Id:       id,
		Name:     name,
		MimeType: "application/vnd.google-apps.folder",
	}
}

// TestFileWithCapabilities creates a mock file with capabilities
func TestFileWithCapabilities(id, name string, canDownload, canEdit, canDelete bool) *drive.File {
	return &drive.File{
		Id:       id,
		Name:     name,
		MimeType: "text/plain",
		Capabilities: &drive.FileCapabilities{
			CanDownload: canDownload,
			CanEdit:     canEdit,
			CanDelete:   canDelete,
		},
	}
}

// TestPermission creates a mock permission for testing
func TestPermission(id, permType, role, email string) *drive.Permission {
	return &drive.Permission{
		Id:           id,
		Type:         permType,
		Role:         role,
		EmailAddress: email,
	}
}

// AssertNoError is a helper to fail the test if error is not nil
func AssertNoError(t *testing.T, err error, msgAndArgs ...interface{}) {
	t.Helper()
	if err != nil {
		if len(msgAndArgs) > 0 {
			t.Fatalf("%v: %v", msgAndArgs[0], err)
		} else {
			t.Fatalf("unexpected error: %v", err)
		}
	}
}

// AssertError is a helper to fail the test if error is nil
func AssertError(t *testing.T, err error, msgAndArgs ...interface{}) {
	t.Helper()
	if err == nil {
		if len(msgAndArgs) > 0 {
			t.Fatalf("%v: expected error but got nil", msgAndArgs[0])
		} else {
			t.Fatal("expected error but got nil")
		}
	}
}

// AssertEqual is a helper to fail the test if two values are not equal
func AssertEqual(t *testing.T, got, want interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if got != want {
		if len(msgAndArgs) > 0 {
			t.Fatalf("%v: got %v, want %v", msgAndArgs[0], got, want)
		} else {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

// AssertNotNil is a helper to fail the test if value is nil
func AssertNotNil(t *testing.T, value interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if value == nil {
		if len(msgAndArgs) > 0 {
			t.Fatalf("%v: expected non-nil value", msgAndArgs[0])
		} else {
			t.Fatal("expected non-nil value")
		}
	}
}

// AssertNil is a helper to fail the test if value is not nil
func AssertNil(t *testing.T, value interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if value != nil {
		if len(msgAndArgs) > 0 {
			t.Fatalf("%v: expected nil value but got %v", msgAndArgs[0], value)
		} else {
			t.Fatalf("expected nil value but got %v", value)
		}
	}
}
