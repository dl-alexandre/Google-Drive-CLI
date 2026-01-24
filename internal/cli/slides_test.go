package cli

import (
	"os"
	"testing"
)

func TestReadSlidesRequests_EmptyInput(t *testing.T) {
	t.Cleanup(func() {
		slidesUpdateRequests = ""
		slidesUpdateFile = ""
	})
	slidesUpdateRequests = ""
	slidesUpdateFile = ""
	if _, err := readSlidesRequests(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestReadSlidesRequests_FromJSON(t *testing.T) {
	t.Cleanup(func() {
		slidesUpdateRequests = ""
		slidesUpdateFile = ""
	})
	slidesUpdateRequests = `[{"insertText":{"objectId":"slide1","insertionIndex":0,"text":"Hello"}}]`
	slidesUpdateFile = ""
	requests, err := readSlidesRequests()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(requests))
	}
}

func TestReadSlidesRequests_FromFile(t *testing.T) {
	t.Cleanup(func() {
		slidesUpdateRequests = ""
		slidesUpdateFile = ""
	})
	tmp, err := os.CreateTemp("", "slides-requests-*.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Remove(tmp.Name())
	})
	if _, err := tmp.WriteString(`[{"insertText":{"objectId":"slide1","insertionIndex":0,"text":"Hello"}}]`); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
	if err := tmp.Close(); err != nil {
		t.Fatalf("failed to close temp file: %v", err)
	}
	slidesUpdateRequests = ""
	slidesUpdateFile = tmp.Name()
	requests, err := readSlidesRequests()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(requests))
	}
}

func TestReadSlidesRequests_InvalidJSON(t *testing.T) {
	t.Cleanup(func() {
		slidesUpdateRequests = ""
		slidesUpdateFile = ""
	})
	slidesUpdateRequests = `invalid`
	slidesUpdateFile = ""
	if _, err := readSlidesRequests(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestReadSlidesRequests_FileNotFound(t *testing.T) {
	t.Cleanup(func() {
		slidesUpdateRequests = ""
		slidesUpdateFile = ""
	})
	slidesUpdateRequests = ""
	slidesUpdateFile = "/nonexistent/file.json"
	if _, err := readSlidesRequests(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestReadSlidesReplacements_EmptyInput(t *testing.T) {
	t.Cleanup(func() {
		slidesReplaceData = ""
		slidesReplaceFile = ""
	})
	slidesReplaceData = ""
	slidesReplaceFile = ""
	if _, err := readSlidesReplacements(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestReadSlidesReplacements_FromJSON(t *testing.T) {
	t.Cleanup(func() {
		slidesReplaceData = ""
		slidesReplaceFile = ""
	})
	slidesReplaceData = `{"{{name}}":"John","{{date}}":"2024-01-01"}`
	slidesReplaceFile = ""
	replacements, err := readSlidesReplacements()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(replacements) != 2 {
		t.Fatalf("expected 2 replacements, got %d", len(replacements))
	}
}

func TestReadSlidesReplacements_FromFile(t *testing.T) {
	t.Cleanup(func() {
		slidesReplaceData = ""
		slidesReplaceFile = ""
	})
	tmp, err := os.CreateTemp("", "slides-replacements-*.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Remove(tmp.Name())
	})
	if _, err := tmp.WriteString(`{"{{name}}":"John"}`); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
	if err := tmp.Close(); err != nil {
		t.Fatalf("failed to close temp file: %v", err)
	}
	slidesReplaceData = ""
	slidesReplaceFile = tmp.Name()
	replacements, err := readSlidesReplacements()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(replacements) != 1 {
		t.Fatalf("expected 1 replacement, got %d", len(replacements))
	}
}

func TestReadSlidesReplacements_InvalidJSON(t *testing.T) {
	t.Cleanup(func() {
		slidesReplaceData = ""
		slidesReplaceFile = ""
	})
	slidesReplaceData = `invalid`
	slidesReplaceFile = ""
	if _, err := readSlidesReplacements(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestReadSlidesReplacements_FileNotFound(t *testing.T) {
	t.Cleanup(func() {
		slidesReplaceData = ""
		slidesReplaceFile = ""
	})
	slidesReplaceData = ""
	slidesReplaceFile = "/nonexistent/file.json"
	if _, err := readSlidesReplacements(); err == nil {
		t.Fatalf("expected error")
	}
}
