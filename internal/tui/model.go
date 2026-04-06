package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dl-alexandre/gdrv/internal/tui/vfs"
)

// PreviewDebounceDelay is the delay before fetching preview after cursor stops.
const PreviewDebounceDelay = 150 * time.Millisecond

// ModalType indicates which modal (if any) is active.
type ModalType int

const (
	ModalNone ModalType = iota
	ModalConfirm
	ModalSheetPicker
)

// SheetPickerModal holds the state for the sheet tab picker.
type SheetPickerModal struct {
	FileID         string         // Spreadsheet ID
	FileName       string         // Spreadsheet name
	Tabs           []vfs.SheetTab // List of worksheets
	Cursor         int            // Selected tab index
	Loading        bool           // Loading tabs async
	PreviewTab     string
	PreviewContent string
	PreviewLoading bool
}

// Model is the root Bubble Tea model for the TUI.
type Model struct {
	// Services
	vfs vfs.VFS

	// Navigation (ID-centric)
	currentID   string
	breadcrumbs []vfs.Node

	// Modal state
	modal      ModalType
	sheetModal *SheetPickerModal

	// Listing state
	items    []vfs.Node
	cursor   int
	selected map[string]bool

	// Preview
	previewVisible bool
	previewNode    *vfs.Node
	previewContent string

	// Status
	loading bool
	err     error
	notice  string

	// Layout
	width  int
	height int
}

// NewModel creates a new TUI model.
func NewModel(vfsImpl vfs.VFS) Model {
	return Model{
		vfs:      vfsImpl,
		selected: make(map[string]bool),
		items:    []vfs.Node{},
	}
}

// Title returns the terminal window title.
func (m Model) Title() string {
	return "gdrv"
}

// Init kicks off initial loading.
func (m Model) Init() tea.Cmd {
	return loadRootsCmd(m.vfs)
}

func (m *Model) clearError() {
	m.err = nil
}
