package tui

import (
	"context"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dl-alexandre/gdrv/internal/tui/editor"
	"github.com/dl-alexandre/gdrv/internal/tui/vfs"
)

// Update handles all messages and dispatches to appropriate handlers.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKey(msg)

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case RootsMsg:
		return m.handleRoots(msg)

	case LoadDirMsg:
		return m.handleLoadDir(msg)

	case PreviewMsg:
		return m.handlePreview(msg)

	case ErrMsg:
		return m.handleError(msg)

	case DebouncedPreviewMsg:
		// Only load preview if cursor hasn't moved since we scheduled this
		if len(m.items) > 0 && m.cursor < len(m.items) {
			node := &m.items[m.cursor]
			if node.ID == msg.NodeID {
				return m, m.loadPreview()
			}
		}
		return m, nil

	case LoadSheetTabsMsg:
		return m.handleLoadSheetTabs(msg)

	case SheetPreviewMsg:
		return m.handleSheetPreview(msg)

	case EditPreparedMsg:
		m.notice = compactWarnings(msg.Warnings)
		bridge := editor.NewSheetsBridge()
		cmd := bridge.Command(context.Background(), msg.Session.TempPath)
		return m, tea.ExecProcess(cmd, func(err error) tea.Msg {
			return EditorExitedMsg{Session: msg.Session, Err: err}
		})

	case EditorExitedMsg:
		return m, finalizeSheetEditCmd(m.vfs, msg.Session, time.Since(msg.Session.ExportedAt))

	case EditCompleteMsg:
		// Close modal and refresh after edit
		m.modal = ModalNone
		m.sheetModal = nil
		m.loading = true
		m.notice = msg.Message
		return m, loadDirCmd(m.vfs, m.currentID)

	default:
		return m, nil
	}
}

// handleKey processes keyboard input.
func (m Model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Modal takes precedence
	if m.modal != ModalNone {
		return m.handleModalKey(msg)
	}

	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit

	case "j", "down":
		return m.moveCursor(1)

	case "k", "up":
		return m.moveCursor(-1)

	case "l", "enter":
		return m.activateSelection()

	case "h":
		return m.goToParent()

	case "p":
		m.previewVisible = !m.previewVisible
		return m, nil

	case "R":
		return m, m.refresh()

	default:
		return m, nil
	}
}

// moveCursor moves the cursor up or down, with bounds checking and debounced preview.
func (m Model) moveCursor(delta int) (tea.Model, tea.Cmd) {
	if len(m.items) == 0 {
		return m, nil
	}

	newCursor := m.cursor + delta
	if newCursor < 0 {
		newCursor = 0
	}
	if newCursor >= len(m.items) {
		newCursor = len(m.items) - 1
	}

	// Cursor didn't change
	if newCursor == m.cursor {
		return m, nil
	}

	m.cursor = newCursor

	// Schedule debounced preview if visible
	if m.previewVisible && len(m.items) > 0 {
		node := &m.items[m.cursor]
		m.previewNode = node
		return m, tea.Tick(PreviewDebounceDelay, func(t time.Time) tea.Msg {
			return DebouncedPreviewMsg{NodeID: node.ID}
		})
	}
	return m, nil
}

// activateSelection opens a folder or activates a file.
func (m Model) activateSelection() (tea.Model, tea.Cmd) {
	if len(m.items) == 0 || m.cursor >= len(m.items) {
		return m, nil
	}

	node := &m.items[m.cursor]

	if node.IsDir {
		// Navigate into folder
		m.breadcrumbs = append(m.breadcrumbs, *node)
		m.currentID = node.ID
		m.loading = true
		m.clearError()
		return m, loadDirCmd(m.vfs, node.ID)
	}

	// Handle Google Sheets - open tab picker
	if node.Kind == vfs.KindSheet {
		m.modal = ModalSheetPicker
		m.sheetModal = &SheetPickerModal{
			FileID:   node.ID,
			FileName: node.Name,
			Tabs:     []vfs.SheetTab{},
			Cursor:   0,
			Loading:  true,
		}
		return m, loadSheetTabsCmd(m.vfs, node.ID)
	}

	// Phase 2: other file types - no-op for now
	return m, nil
}

// goToParent navigates to the parent directory.
func (m Model) goToParent() (tea.Model, tea.Cmd) {
	if len(m.breadcrumbs) <= 1 {
		return m, nil // At root
	}

	// Pop breadcrumb
	m.breadcrumbs = m.breadcrumbs[:len(m.breadcrumbs)-1]
	parent := m.breadcrumbs[len(m.breadcrumbs)-1]
	m.currentID = parent.ID
	m.loading = true
	m.clearError()

	return m, loadDirCmd(m.vfs, parent.ID)
}

// refresh reloads the current directory.
func (m Model) refresh() tea.Cmd {
	m.loading = true
	m.clearError()
	return loadDirCmd(m.vfs, m.currentID)
}

// loadPreview triggers async preview loading.
func (m Model) loadPreview() tea.Cmd {
	if m.cursor >= len(m.items) {
		return nil
	}
	node := &m.items[m.cursor]
	return statPreviewCmd(m.vfs, node.ID)
}

// handleRoots processes the initial root loading.
func (m Model) handleRoots(msg RootsMsg) (tea.Model, tea.Cmd) {
	if len(msg.Roots) == 0 {
		return m, nil
	}

	root := msg.Roots[0]
	m.breadcrumbs = []vfs.Node{root}
	m.currentID = root.ID
	m.loading = true

	return m, loadDirCmd(m.vfs, root.ID)
}

// handleLoadDir processes directory contents loading.
func (m Model) handleLoadDir(msg LoadDirMsg) (tea.Model, tea.Cmd) {
	m.items = msg.Items
	m.cursor = 0
	m.loading = false
	m.selected = make(map[string]bool)
	m.previewNode = nil
	m.previewContent = ""
	m.clearError()

	// Load initial preview if visible
	if m.previewVisible && len(m.items) > 0 {
		m.previewNode = &m.items[m.cursor]
		return m, m.loadPreview()
	}
	return m, nil
}

// handlePreview processes preview content.
func (m Model) handlePreview(msg PreviewMsg) (tea.Model, tea.Cmd) {
	if m.previewNode == nil || m.previewNode.ID != msg.NodeID {
		return m, nil
	}
	m.previewContent = msg.Content
	return m, nil
}

// handleError processes errors.
func (m Model) handleError(msg ErrMsg) (tea.Model, tea.Cmd) {
	m.loading = false
	m.err = msg.Err
	m.notice = ""
	return m, nil
}

// handleLoadSheetTabs processes loaded worksheet tabs.
func (m Model) handleLoadSheetTabs(msg LoadSheetTabsMsg) (tea.Model, tea.Cmd) {
	if m.sheetModal != nil && m.sheetModal.FileID == msg.SpreadsheetID {
		m.sheetModal.Tabs = msg.Tabs
		m.sheetModal.Loading = false
	}
	return m, nil
}

func (m Model) handleSheetPreview(msg SheetPreviewMsg) (tea.Model, tea.Cmd) {
	if m.sheetModal != nil && m.sheetModal.FileID == msg.SpreadsheetID && m.sheetModal.PreviewTab == msg.TabName {
		m.sheetModal.PreviewContent = msg.Content
		m.sheetModal.PreviewLoading = false
	}
	return m, nil
}

// handleModalKey handles keys when modal is open.
func (m Model) handleModalKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch m.modal {
	case ModalSheetPicker:
		return m.handleSheetPickerKey(msg)
	default:
		// Default: close modal
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.modal = ModalNone
			m.sheetModal = nil
			return m, nil
		default:
			return m, nil
		}
	}
}

// handleSheetPickerKey handles keys in the sheet picker modal.
func (m Model) handleSheetPickerKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.sheetModal == nil {
		m.modal = ModalNone
		return m, nil
	}

	switch msg.String() {
	case "q", "esc", "ctrl+c":
		// Close modal
		m.modal = ModalNone
		m.sheetModal = nil
		return m, nil

	case "j", "down":
		// Move down
		if m.sheetModal.Cursor < len(m.sheetModal.Tabs)-1 {
			m.sheetModal.Cursor++
		}
		return m, nil

	case "k", "up":
		// Move up
		if m.sheetModal.Cursor > 0 {
			m.sheetModal.Cursor--
		}
		return m, nil

	case "enter":
		if len(m.sheetModal.Tabs) == 0 || m.sheetModal.Cursor >= len(m.sheetModal.Tabs) {
			return m, nil
		}
		tab := m.sheetModal.Tabs[m.sheetModal.Cursor]
		m.sheetModal.PreviewTab = tab.Title
		m.sheetModal.PreviewContent = ""
		m.sheetModal.PreviewLoading = true
		return m, previewSheetTabCmd(m.vfs, m.sheetModal.FileID, tab.Title)

	case "e":
		// Edit selected tab in external editor (Phase 3)
		if len(m.sheetModal.Tabs) == 0 || m.sheetModal.Cursor >= len(m.sheetModal.Tabs) {
			return m, nil
		}
		tab := m.sheetModal.Tabs[m.sheetModal.Cursor]
		return m, prepareSheetEditCmd(m.vfs, m.sheetModal.FileID, tab.Title)

	default:
		return m, nil
	}
}

// isConflictError checks if an error is a revision conflict.
func isConflictError(err error) bool {
	if err == nil {
		return false
	}
	return len(err.Error()) > 10 && err.Error()[:10] == "conflict: "
}

// Helper for import resolution
var _ = lipgloss.NewStyle // Ensure lipgloss import is used
