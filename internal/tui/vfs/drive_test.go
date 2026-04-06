package vfs

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/files"
	"github.com/dl-alexandre/gdrv/internal/logging"
	gsheets "github.com/dl-alexandre/gdrv/internal/sheets"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	sheetsapi "google.golang.org/api/sheets/v4"
)

func TestDriveVFSRoots_UsesSharedDriveRootWhenConfigured(t *testing.T) {
	v := NewDriveVFS(nil, nil, "shared123")
	roots, err := v.Roots(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(roots) != 1 || roots[0].ID != "shared123" {
		t.Fatalf("expected shared drive root, got %+v", roots)
	}
}

func TestDriveVFSList_PropagatesDriveID(t *testing.T) {
	var seenDriveID string
	server, driveSvc, sheetsSvc := newDriveVFSTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/drive/v3/files":
			seenDriveID = r.URL.Query().Get("driveId")
			writeJSON(t, w, http.StatusOK, map[string]any{"files": []map[string]any{}})
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
	})
	defer server.Close()

	v := newTestDriveVFS(driveSvc, sheetsSvc, "shared123")
	if _, err := v.List(context.Background(), "shared123"); err != nil {
		t.Fatal(err)
	}
	if seenDriveID != "shared123" {
		t.Fatalf("expected driveId to be propagated, got %q", seenDriveID)
	}
}

func TestDriveVFSExportSheetTab_ReturnsCSV(t *testing.T) {
	server, driveSvc, sheetsSvc := newDriveVFSTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/drive/v3/files/file123":
			writeJSON(t, w, http.StatusOK, map[string]any{"id": "file123", "modifiedTime": "2024-01-01T00:00:00Z"})
		case r.Method == http.MethodGet && r.URL.Path == "/v4/spreadsheets/file123":
			writeJSON(t, w, http.StatusOK, map[string]any{
				"spreadsheetId": "file123",
				"properties":    map[string]any{"title": "Test"},
				"sheets":        []map[string]any{{"properties": map[string]any{"sheetId": 1, "title": "Sheet1", "index": 0, "sheetType": "GRID"}}},
			})
		case r.Method == http.MethodGet && r.URL.Path == "/v4/spreadsheets/file123/values/'Sheet1'":
			writeJSON(t, w, http.StatusOK, map[string]any{
				"range":          "'Sheet1'!A1:B2",
				"majorDimension": "ROWS",
				"values":         [][]any{{"A", "B"}, {"1", "2"}},
			})
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
	})
	defer server.Close()

	v := newTestDriveVFS(driveSvc, sheetsSvc, "")
	data, revision, err := v.ExportSheetTab(context.Background(), "file123", "Sheet1")
	if err != nil {
		t.Fatal(err)
	}
	if revision != "2024-01-01T00:00:00Z" {
		t.Fatalf("unexpected revision %q", revision)
	}
	if data != "A,B\n1,2\n" {
		t.Fatalf("unexpected csv export %q", data)
	}
}

func TestDriveVFSImportSheetTab_EmptyCSVClearsSheet(t *testing.T) {
	cleared := false
	server, driveSvc, sheetsSvc := newDriveVFSTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/drive/v3/files/file123":
			writeJSON(t, w, http.StatusOK, map[string]any{"id": "file123", "modifiedTime": "2024-01-01T00:00:00Z"})
		case r.Method == http.MethodGet && r.URL.Path == "/v4/spreadsheets/file123/values/'Sheet1'":
			writeJSON(t, w, http.StatusOK, map[string]any{"range": "'Sheet1'!A1:B2", "majorDimension": "ROWS", "values": [][]any{{"A", "B"}, {"1", "2"}}})
		case r.Method == http.MethodPost && strings.Contains(r.URL.Path, "/values/'Sheet1':clear"):
			cleared = true
			writeJSON(t, w, http.StatusOK, map[string]any{"spreadsheetId": "file123", "clearedRange": "'Sheet1'"})
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
	})
	defer server.Close()

	v := newTestDriveVFS(driveSvc, sheetsSvc, "")
	if err := v.ImportSheetTab(context.Background(), "file123", "Sheet1", "", "2024-01-01T00:00:00Z"); err != nil {
		t.Fatal(err)
	}
	if !cleared {
		t.Fatal("expected empty csv import to clear sheet")
	}
}

func TestDriveVFSImportSheetTab_ConflictOnRevisionMismatch(t *testing.T) {
	server, driveSvc, sheetsSvc := newDriveVFSTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/drive/v3/files/file123":
			writeJSON(t, w, http.StatusOK, map[string]any{"id": "file123", "modifiedTime": "2024-01-02T00:00:00Z"})
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
	})
	defer server.Close()

	v := newTestDriveVFS(driveSvc, sheetsSvc, "")
	err := v.ImportSheetTab(context.Background(), "file123", "Sheet1", "A,B\n1,2\n", "2024-01-01T00:00:00Z")
	if err == nil || !strings.Contains(err.Error(), "conflict:") {
		t.Fatalf("expected conflict error, got %v", err)
	}
}

func newDriveVFSTestServer(t *testing.T, handler http.HandlerFunc) (*httptest.Server, *drive.Service, *sheetsapi.Service) {
	t.Helper()
	server := httptest.NewServer(handler)

	ctx := context.Background()
	driveService, err := drive.NewService(ctx,
		option.WithEndpoint(server.URL+"/drive/v3/"),
		option.WithHTTPClient(server.Client()),
	)
	if err != nil {
		t.Fatalf("failed to create drive service: %v", err)
	}

	sheetsService, err := sheetsapi.NewService(ctx,
		option.WithEndpoint(server.URL+"/"),
		option.WithHTTPClient(server.Client()),
	)
	if err != nil {
		t.Fatalf("failed to create sheets service: %v", err)
	}

	return server, driveService, sheetsService
}

func newTestDriveVFS(driveService *drive.Service, sheetsService *sheetsapi.Service, driveID string) *DriveVFS {
	client := api.NewClient(driveService, 0, 1, logging.NewNoOpLogger())
	return NewDriveVFS(files.NewManager(client), gsheets.NewManager(client, sheetsService), driveID)
}

func writeJSON(t *testing.T, w http.ResponseWriter, status int, payload any) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		t.Fatalf("failed to encode response: %v", err)
	}
}
