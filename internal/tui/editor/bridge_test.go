package editor

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSanitizeFilename(t *testing.T) {
	got := SanitizeFilename("Q1/Budget:*?<>|Report")
	if strings.ContainsAny(got, `<>:"/\\|?*`) {
		t.Fatalf("expected sanitized filename, got %q", got)
	}
}

func TestMakeEditPath(t *testing.T) {
	base := t.TempDir()
	got := MakeEditPath(base, "file123", "Q1/Budget")

	if !strings.HasPrefix(got, filepath.Join(base, "file123")) {
		t.Fatalf("expected path under temp base, got %q", got)
	}
	if filepath.Ext(got) != ".csv" {
		t.Fatalf("expected csv path, got %q", got)
	}
}

func TestComputeFileHash_ChangesWhenFileChanges(t *testing.T) {
	path := filepath.Join(t.TempDir(), "sheet.csv")
	if err := os.WriteFile(path, []byte("a,b\n1,2\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	before, err := ComputeFileHash(path)
	if err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(path, []byte("a,b\n1,3\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	after, err := ComputeFileHash(path)
	if err != nil {
		t.Fatal(err)
	}

	if before == after {
		t.Fatal("expected file hash to change after editing")
	}
}

func TestCheckEditor_MissingBinary(t *testing.T) {
	err := CheckEditor("definitely-not-a-real-editor-binary")
	if err == nil {
		t.Fatal("expected missing editor error")
	}
	if !errors.Is(err, ErrEditorNotFound) {
		t.Fatalf("expected ErrEditorNotFound, got %v", err)
	}
}
