package tui

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/dl-alexandre/gdrv/internal/tui/editor"
	"github.com/dl-alexandre/gdrv/internal/tui/vfs"
)

type commandTestVFS struct {
	exportCSV      string
	exportRevision string
	exportErr      error
	hasFormulas    bool
	importErr      error
	importCalls    int
	importData     string
	importRevision string
}

func (v *commandTestVFS) Roots(context.Context) ([]vfs.Node, error)        { return nil, nil }
func (v *commandTestVFS) List(context.Context, string) ([]vfs.Node, error) { return nil, nil }
func (v *commandTestVFS) Stat(context.Context, string) (vfs.Node, error)   { return vfs.Node{}, nil }
func (v *commandTestVFS) Rename(context.Context, string, string) error     { return nil }
func (v *commandTestVFS) Move(context.Context, string, string) error       { return nil }
func (v *commandTestVFS) Delete(context.Context, string) error             { return nil }
func (v *commandTestVFS) ListSheetTabs(context.Context, string) ([]vfs.SheetTab, error) {
	return nil, nil
}
func (v *commandTestVFS) ExportSheetTab(context.Context, string, string) (string, string, error) {
	return v.exportCSV, v.exportRevision, v.exportErr
}
func (v *commandTestVFS) ImportSheetTab(_ context.Context, _, _ string, csvData string, expectedRevision string) error {
	v.importCalls++
	v.importData = csvData
	v.importRevision = expectedRevision
	return v.importErr
}
func (v *commandTestVFS) GetFileRevision(context.Context, string) (string, error) { return "", nil }
func (v *commandTestVFS) TabHasFormulas(context.Context, string, string) (bool, error) {
	return v.hasFormulas, nil
}

func TestFormatNodePreview_ShortIDDoesNotPanic(t *testing.T) {
	node := vfs.Node{
		ID:       "doc1",
		Name:     "notes.txt",
		Kind:     vfs.KindBinary,
		Modified: time.Date(2024, 1, 2, 3, 4, 0, 0, time.UTC),
		Size:     42,
	}

	got := formatNodePreview(node)

	if !strings.Contains(got, "ID: doc1") {
		t.Fatalf("expected full short id in preview, got %q", got)
	}
}

func TestFormatSheetPreview_TruncatesRowsAndColumns(t *testing.T) {
	csvData := "A,B,C,D,E,F,G\n1,2,3,4,5,6,7\n8,9,10,11,12,13,14\n"

	got, err := formatSheetPreview(csvData)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(got, "A") || !strings.Contains(got, "F") {
		t.Fatalf("expected preview header cells, got %q", got)
	}
	if strings.Contains(got, "G") {
		t.Fatalf("expected extra columns to be trimmed, got %q", got)
	}
	if !strings.Contains(got, "...") {
		t.Fatalf("expected truncation marker, got %q", got)
	}
}

func TestPrepareSheetEditCmd_AddsFormulaWarningOnlyWhenNeeded(t *testing.T) {
	originalCheckEditor := checkEditor
	checkEditor = func(string) error { return nil }
	defer func() { checkEditor = originalCheckEditor }()

	v := &commandTestVFS{
		exportCSV:      "A,B\n=SUM(1,2),3\n",
		exportRevision: "rev-1",
		hasFormulas:    true,
	}

	msg := prepareSheetEditCmd(v, "file123", "Sheet1")()
	prepared, ok := msg.(EditPreparedMsg)
	if !ok {
		t.Fatalf("expected EditPreparedMsg, got %T", msg)
	}
	defer os.Remove(prepared.Session.TempPath)

	joined := strings.Join(prepared.Warnings, " ")
	if !strings.Contains(joined, "contains formulas") {
		t.Fatalf("expected formula warning, got %q", joined)
	}

	v.hasFormulas = false
	msg = prepareSheetEditCmd(v, "file123", "Sheet1")()
	prepared, ok = msg.(EditPreparedMsg)
	if !ok {
		t.Fatalf("expected EditPreparedMsg, got %T", msg)
	}
	defer os.Remove(prepared.Session.TempPath)

	joined = strings.Join(prepared.Warnings, " ")
	if strings.Contains(joined, "contains formulas") {
		t.Fatalf("did not expect formula warning, got %q", joined)
	}
}

func TestFinalizeSheetEditCmd_NoChanges(t *testing.T) {
	tempPath := filepath.Join(t.TempDir(), "sheet.csv")
	content := []byte("A,B\n1,2\n")
	if err := os.WriteFile(tempPath, content, 0o644); err != nil {
		t.Fatal(err)
	}

	hash, err := editor.ComputeFileHash(tempPath)
	if err != nil {
		t.Fatal(err)
	}

	v := &commandTestVFS{}
	msg := finalizeSheetEditCmd(v, editor.EditSession{
		FileID:     "file123",
		TabName:    "Sheet1",
		TempPath:   tempPath,
		BeforeHash: hash,
	}, time.Second)()

	complete, ok := msg.(EditCompleteMsg)
	if !ok {
		t.Fatalf("expected EditCompleteMsg, got %T", msg)
	}
	if complete.Changed {
		t.Fatal("expected no-change result")
	}
	if v.importCalls != 0 {
		t.Fatalf("expected no import call, got %d", v.importCalls)
	}
}

func TestFinalizeSheetEditCmd_Conflict(t *testing.T) {
	tempPath := filepath.Join(t.TempDir(), "sheet.csv")
	if err := os.WriteFile(tempPath, []byte("A,B\n1,2\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	v := &commandTestVFS{importErr: errors.New("conflict: spreadsheet changed remotely")}
	msg := finalizeSheetEditCmd(v, editor.EditSession{
		FileID:     "file123",
		TabName:    "Sheet1",
		TempPath:   tempPath,
		BeforeHash: "different",
		Revision:   "rev-1",
	}, time.Second)()

	errMsg, ok := msg.(ErrMsg)
	if !ok {
		t.Fatalf("expected ErrMsg, got %T", msg)
	}
	if !strings.Contains(errMsg.Err.Error(), "conflict detected") {
		t.Fatalf("expected conflict message, got %v", errMsg.Err)
	}
	if v.importCalls != 1 {
		t.Fatalf("expected import call, got %d", v.importCalls)
	}
}
