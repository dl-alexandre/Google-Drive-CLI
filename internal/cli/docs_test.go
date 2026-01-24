package cli

import (
	"os"
	"testing"
)

func TestReadDocsRequests_EmptyInput(t *testing.T) {
	t.Cleanup(func() {
		docsUpdateRequests = ""
		docsUpdateFile = ""
	})
	docsUpdateRequests = ""
	docsUpdateFile = ""
	if _, err := readDocsRequests(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestReadDocsRequests_FromJSON(t *testing.T) {
	t.Cleanup(func() {
		docsUpdateRequests = ""
		docsUpdateFile = ""
	})
	docsUpdateRequests = `[{"insertText":{"location":{"index":1},"text":"Hello"}}]`
	docsUpdateFile = ""
	requests, err := readDocsRequests()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(requests))
	}
}

func TestReadDocsRequests_FromFile(t *testing.T) {
	t.Cleanup(func() {
		docsUpdateRequests = ""
		docsUpdateFile = ""
	})
	tmp, err := os.CreateTemp("", "requests-*.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Remove(tmp.Name())
	})
	if _, err := tmp.WriteString(`[{"insertText":{"location":{"index":1},"text":"Hello"}}]`); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
	if err := tmp.Close(); err != nil {
		t.Fatalf("failed to close temp file: %v", err)
	}
	docsUpdateRequests = ""
	docsUpdateFile = tmp.Name()
	requests, err := readDocsRequests()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(requests))
	}
}

func TestReadDocsRequests_InvalidJSON(t *testing.T) {
	t.Cleanup(func() {
		docsUpdateRequests = ""
		docsUpdateFile = ""
	})
	docsUpdateRequests = `invalid`
	docsUpdateFile = ""
	if _, err := readDocsRequests(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestReadDocsRequests_FileNotFound(t *testing.T) {
	t.Cleanup(func() {
		docsUpdateRequests = ""
		docsUpdateFile = ""
	})
	docsUpdateRequests = ""
	docsUpdateFile = "/nonexistent/file.json"
	if _, err := readDocsRequests(); err == nil {
		t.Fatalf("expected error")
	}
}
