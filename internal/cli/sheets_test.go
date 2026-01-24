package cli

import (
	"os"
	"testing"
)

func TestReadSheetValues_EmptyInput(t *testing.T) {
	t.Cleanup(func() {
		sheetsValuesJSON = ""
		sheetsValuesFile = ""
	})
	sheetsValuesJSON = ""
	sheetsValuesFile = ""
	if _, err := readSheetValues(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestReadSheetValues_FromJSON(t *testing.T) {
	t.Cleanup(func() {
		sheetsValuesJSON = ""
		sheetsValuesFile = ""
	})
	sheetsValuesJSON = `[[1,2],[3,4]]`
	sheetsValuesFile = ""
	values, err := readSheetValues()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(values) != 2 {
		t.Fatalf("expected 2 rows, got %d", len(values))
	}
}

func TestReadSheetValues_FromFile(t *testing.T) {
	t.Cleanup(func() {
		sheetsValuesJSON = ""
		sheetsValuesFile = ""
	})
	tmp, err := os.CreateTemp("", "values-*.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Remove(tmp.Name())
	})
	if _, err := tmp.WriteString(`[[1,2]]`); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
	if err := tmp.Close(); err != nil {
		t.Fatalf("failed to close temp file: %v", err)
	}
	sheetsValuesJSON = ""
	sheetsValuesFile = tmp.Name()
	values, err := readSheetValues()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(values) != 1 {
		t.Fatalf("expected 1 row, got %d", len(values))
	}
}

func TestReadSheetValues_InvalidJSON(t *testing.T) {
	t.Cleanup(func() {
		sheetsValuesJSON = ""
		sheetsValuesFile = ""
	})
	sheetsValuesJSON = `invalid`
	sheetsValuesFile = ""
	if _, err := readSheetValues(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestReadSheetValues_FileNotFound(t *testing.T) {
	t.Cleanup(func() {
		sheetsValuesJSON = ""
		sheetsValuesFile = ""
	})
	sheetsValuesJSON = ""
	sheetsValuesFile = "/nonexistent/file.json"
	if _, err := readSheetValues(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestReadSheetsBatchRequests_EmptyInput(t *testing.T) {
	t.Cleanup(func() {
		sheetsBatchUpdateJSON = ""
		sheetsBatchUpdateFile = ""
	})
	sheetsBatchUpdateJSON = ""
	sheetsBatchUpdateFile = ""
	if _, err := readSheetsBatchRequests(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestReadSheetsBatchRequests_FromJSON(t *testing.T) {
	t.Cleanup(func() {
		sheetsBatchUpdateJSON = ""
		sheetsBatchUpdateFile = ""
	})
	sheetsBatchUpdateJSON = `[{"updateSpreadsheetProperties":{"properties":{"title":"New Title"},"fields":"title"}}]`
	sheetsBatchUpdateFile = ""
	requests, err := readSheetsBatchRequests()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(requests))
	}
}

func TestReadSheetsBatchRequests_FromFile(t *testing.T) {
	t.Cleanup(func() {
		sheetsBatchUpdateJSON = ""
		sheetsBatchUpdateFile = ""
	})
	tmp, err := os.CreateTemp("", "batch-requests-*.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Remove(tmp.Name())
	})
	if _, err := tmp.WriteString(`[{"updateSpreadsheetProperties":{"properties":{"title":"New Title"},"fields":"title"}}]`); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
	if err := tmp.Close(); err != nil {
		t.Fatalf("failed to close temp file: %v", err)
	}
	sheetsBatchUpdateJSON = ""
	sheetsBatchUpdateFile = tmp.Name()
	requests, err := readSheetsBatchRequests()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(requests))
	}
}

func TestReadSheetsBatchRequests_InvalidJSON(t *testing.T) {
	t.Cleanup(func() {
		sheetsBatchUpdateJSON = ""
		sheetsBatchUpdateFile = ""
	})
	sheetsBatchUpdateJSON = `invalid`
	sheetsBatchUpdateFile = ""
	if _, err := readSheetsBatchRequests(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestReadSheetsBatchRequests_FileNotFound(t *testing.T) {
	t.Cleanup(func() {
		sheetsBatchUpdateJSON = ""
		sheetsBatchUpdateFile = ""
	})
	sheetsBatchUpdateJSON = ""
	sheetsBatchUpdateFile = "/nonexistent/file.json"
	if _, err := readSheetsBatchRequests(); err == nil {
		t.Fatalf("expected error")
	}
}
