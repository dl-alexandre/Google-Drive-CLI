package tui

import (
	"context"
	"testing"

	"github.com/dl-alexandre/gdrv/internal/tui/vfs"
)

type previewTestVFS struct{}

func (previewTestVFS) Roots(context.Context) ([]vfs.Node, error)                     { return nil, nil }
func (previewTestVFS) List(context.Context, string) ([]vfs.Node, error)              { return nil, nil }
func (previewTestVFS) Stat(context.Context, string) (vfs.Node, error)                { return vfs.Node{}, nil }
func (previewTestVFS) Rename(context.Context, string, string) error                  { return nil }
func (previewTestVFS) Move(context.Context, string, string) error                    { return nil }
func (previewTestVFS) Delete(context.Context, string) error                          { return nil }
func (previewTestVFS) ListSheetTabs(context.Context, string) ([]vfs.SheetTab, error) { return nil, nil }
func (previewTestVFS) ExportSheetTab(context.Context, string, string) (string, string, error) {
	return "", "", nil
}
func (previewTestVFS) ImportSheetTab(context.Context, string, string, string, string) error {
	return nil
}
func (previewTestVFS) GetFileRevision(context.Context, string) (string, error) { return "", nil }
func (previewTestVFS) TabHasFormulas(context.Context, string, string) (bool, error) {
	return false, nil
}

func TestHandlePreview_IgnoresStalePreviewMessages(t *testing.T) {
	m := NewModel(previewTestVFS{})
	m.previewNode = &vfs.Node{ID: "current-node"}
	m.previewContent = "current preview"

	updated, _ := m.handlePreview(PreviewMsg{NodeID: "old-node", Content: "stale preview"})
	model := updated.(Model)

	if model.previewContent != "current preview" {
		t.Fatalf("expected stale preview to be ignored, got %q", model.previewContent)
	}
}

func TestHandlePreview_AppliesCurrentPreviewMessages(t *testing.T) {
	m := NewModel(previewTestVFS{})
	m.previewNode = &vfs.Node{ID: "current-node"}

	updated, _ := m.handlePreview(PreviewMsg{NodeID: "current-node", Content: "fresh preview"})
	model := updated.(Model)

	if model.previewContent != "fresh preview" {
		t.Fatalf("expected preview to update, got %q", model.previewContent)
	}
}
