package tui

import (
	"time"

	"github.com/dl-alexandre/gdrv/internal/tui/editor"
	"github.com/dl-alexandre/gdrv/internal/tui/vfs"
)

// tea.Msg types for TUI communication.

// LoadDirMsg signals directory contents loaded.
type LoadDirMsg struct {
	ParentID string
	Items    []vfs.Node
}

// PreviewMsg signals preview content loaded.
type PreviewMsg struct {
	NodeID  string
	Content string
}

// ErrMsg signals an error occurred.
type ErrMsg struct {
	Err error
}

// RootsMsg signals root locations loaded.
type RootsMsg struct {
	Roots []vfs.Node
}

// DebouncedPreviewMsg triggers preview loading after cursor stops moving.
type DebouncedPreviewMsg struct {
	NodeID string
}

// LoadSheetTabsMsg signals spreadsheet worksheet tabs loaded.
type LoadSheetTabsMsg struct {
	SpreadsheetID string
	Tabs          []vfs.SheetTab
}

// SheetPreviewMsg carries a read-only preview for a worksheet tab.
type SheetPreviewMsg struct {
	SpreadsheetID string
	TabName       string
	Content       string
}

// EditPreparedMsg signals a sheet edit session is ready to launch.
type EditPreparedMsg struct {
	Session  editor.EditSession
	RowCount int
	Warnings []string
}

// EditorExitedMsg signals the external editor exited.
type EditorExitedMsg struct {
	Session editor.EditSession
	Err     error
}

// EditCompleteMsg signals a sheet edit session completed.
type EditCompleteMsg struct {
	SpreadsheetID string
	TabName       string
	Changed       bool
	Duration      time.Duration
	Message       string
}
