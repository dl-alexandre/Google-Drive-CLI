package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/dl-alexandre/gdrv/internal/tui/vfs"
)

// View renders the complete TUI.
func (m Model) View() string {
	// Handle modal overlay first
	if m.modal != ModalNone {
		return m.renderModalOverlay()
	}

	// Handle initial loading
	if m.loading && len(m.items) == 0 {
		return m.renderLoading()
	}

	// Main content
	browser := m.renderBrowser()
	status := m.renderStatusBar()

	// Single pane if narrow or preview hidden
	if !m.previewVisible || m.width < PaneMinWidth {
		return lipgloss.JoinVertical(lipgloss.Left, browser, status)
	}

	// Split pane
	preview := m.renderPreview()
	body := lipgloss.JoinHorizontal(
		lipgloss.Top,
		Styles.BrowserPane.Width(m.width*60/100).Render(browser),
		Styles.PreviewPane.Width(m.width*40/100).Render(preview),
	)

	return lipgloss.JoinVertical(lipgloss.Left, body, status)
}

// renderModalOverlay renders the current modal on top of a dimmed background.
func (m Model) renderModalOverlay() string {
	// Get the background view (dimmed)
	var background string
	if m.loading && len(m.items) == 0 {
		background = m.renderLoading()
	} else {
		browser := m.renderBrowser()
		status := m.renderStatusBar()
		if !m.previewVisible || m.width < PaneMinWidth {
			background = lipgloss.JoinVertical(lipgloss.Left, browser, status)
		} else {
			preview := m.renderPreview()
			body := lipgloss.JoinHorizontal(
				lipgloss.Top,
				Styles.BrowserPane.Width(m.width*60/100).Render(browser),
				Styles.PreviewPane.Width(m.width*40/100).Render(preview),
			)
			background = lipgloss.JoinVertical(lipgloss.Left, body, status)
		}
	}

	// Render the modal
	var modalContent string
	switch m.modal {
	case ModalSheetPicker:
		modalContent = m.renderSheetPicker()
	default:
		modalContent = "Unknown modal"
	}

	// Center the modal
	modalStyle := lipgloss.NewStyle().
		Width(m.width-4).
		Height(m.height-4).
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("39")).
		Background(lipgloss.Color("235"))

	// Overlay modal on background
	// For simplicity, we'll just show the modal for now
	_ = background
	return modalStyle.Render(modalContent)
}

// renderLoading shows initial loading state.
func (m Model) renderLoading() string {
	return Styles.Loading.Render("Loading...")
}

// renderBrowser renders the file list.
func (m Model) renderBrowser() string {
	// Empty state handling
	if len(m.items) == 0 {
		if m.loading {
			return Styles.Loading.Render("Loading...")
		}
		// Show breadcrumb even when empty, then empty message
		var b strings.Builder
		b.WriteString(m.renderBreadcrumb())
		b.WriteString("\n\n")
		if m.err != nil {
			b.WriteString(Styles.ErrorBar.Render(fmt.Sprintf("⚠ %v", m.err)))
		} else {
			b.WriteString(Styles.EmptyState.Render("(empty folder)"))
		}
		return b.String()
	}

	var b strings.Builder

	// Breadcrumb header
	b.WriteString(m.renderBreadcrumb())
	b.WriteString("\n\n")

	// File list
	visibleHeight := m.height - 4 // Reserve space for header and status
	startIdx, endIdx := m.calculateVisibleRange(visibleHeight)

	for i := startIdx; i < endIdx && i < len(m.items); i++ {
		line := m.renderItem(i)
		b.WriteString(line)
		if i < endIdx-1 {
			b.WriteString("\n")
		}
	}

	// Show "more" indicator if needed
	if endIdx < len(m.items) {
		b.WriteString("\n")
		b.WriteString(Styles.NormalItem.Render("..."))
	}

	return b.String()
}

// renderBreadcrumb shows current path.
func (m Model) renderBreadcrumb() string {
	if len(m.breadcrumbs) == 0 {
		return Styles.Header.Render("My Drive")
	}

	var parts []string
	for _, b := range m.breadcrumbs {
		parts = append(parts, b.Name)
	}
	return Styles.Header.Render(strings.Join(parts, " / "))
}

// renderItem renders a single file list item.
func (m Model) renderItem(idx int) string {
	node := m.items[idx]
	icon := m.getIconForNode(node)
	kindLabel := m.getKindLabel(node)
	name := node.Name
	if node.IsDir {
		name += "/"
	}

	var line string
	if kindLabel != "" {
		line = fmt.Sprintf(" %s %s %s", icon, name, Styles.KindLabel.Render(kindLabel))
	} else {
		line = fmt.Sprintf(" %s %s", icon, name)
	}

	if idx == m.cursor {
		return Styles.SelectedItem.Render(line)
	}
	if node.IsDir {
		return Styles.DirectoryItem.Render(line)
	}
	// Use kind-specific styling
	return m.getItemStyle(node).Render(line)
}

// getIconForNode returns an icon for the node kind.
func (m Model) getIconForNode(node vfs.Node) string {
	if node.IsDir {
		return ""
	}
	switch node.Kind {
	case vfs.KindSheet:
		return ""
	case vfs.KindDoc:
		return ""
	case vfs.KindSlide:
		return ""
	case vfs.KindPDF:
		return ""
	case vfs.KindImage:
		return ""
	case vfs.KindVideo:
		return ""
	default:
		return ""
	}
}

// getKindLabel returns a short label for the node kind (shown in list).
func (m Model) getKindLabel(node vfs.Node) string {
	if node.IsDir {
		return ""
	}
	switch node.Kind {
	case vfs.KindSheet:
		return "[Sheet]"
	case vfs.KindDoc:
		return "[Doc]"
	case vfs.KindSlide:
		return "[Slide]"
	case vfs.KindPDF:
		return "[PDF]"
	case vfs.KindImage:
		return "[Img]"
	case vfs.KindVideo:
		return "[Vid]"
	default:
		return ""
	}
}

// getItemStyle returns the appropriate style for a node based on its kind.
func (m Model) getItemStyle(node vfs.Node) lipgloss.Style {
	switch node.Kind {
	case vfs.KindSheet:
		return Styles.SheetItem
	case vfs.KindDoc:
		return Styles.DocItem
	case vfs.KindPDF, vfs.KindImage, vfs.KindVideo:
		return Styles.MediaItem
	default:
		return Styles.NormalItem
	}
}

// calculateVisibleRange determines which items to show.
func (m Model) calculateVisibleRange(visibleHeight int) (int, int) {
	if len(m.items) <= visibleHeight {
		return 0, len(m.items)
	}

	// Center cursor in visible area
	half := visibleHeight / 2
	start := m.cursor - half
	if start < 0 {
		start = 0
	}
	end := start + visibleHeight
	if end > len(m.items) {
		end = len(m.items)
		start = end - visibleHeight
		if start < 0 {
			start = 0
		}
	}
	return start, end
}

// renderPreview renders the preview pane.
func (m Model) renderPreview() string {
	if m.previewContent == "" {
		return Styles.NormalItem.Render("No preview")
	}
	return Styles.Preview.Render(m.previewContent)
}

// renderStatusBar renders the bottom status bar.
func (m Model) renderStatusBar() string {
	if m.err != nil {
		errMsg := fmt.Sprintf("ERROR: %v", m.err)
		return Styles.ErrorBar.Render(errMsg)
	}

	var parts []string

	// Path
	if len(m.breadcrumbs) > 0 {
		parts = append(parts, m.breadcrumbs[len(m.breadcrumbs)-1].Name)
	} else {
		parts = append(parts, "My Drive")
	}

	// Item count
	parts = append(parts, fmt.Sprintf("%d items", len(m.items)))

	// Selection
	if len(m.selected) > 0 {
		parts = append(parts, fmt.Sprintf("%d selected", len(m.selected)))
	}

	// Loading indicator
	if m.loading {
		parts = append(parts, "loading...")
	}

	if m.notice != "" {
		parts = append(parts, m.notice)
	}

	// Help hint
	parts = append(parts, "q:quit j/k:nav l:open h:parent p:preview")

	return Styles.StatusBar.Render(strings.Join(parts, "  •  "))
}

// renderSheetPicker renders the sheet tab picker modal.
func (m Model) renderSheetPicker() string {
	if m.sheetModal == nil {
		return "Error: sheet modal not initialized"
	}

	modal := m.sheetModal

	// Header with spreadsheet name
	var b strings.Builder
	b.WriteString(Styles.Header.Render(fmt.Sprintf("📗 %s", modal.FileName)))
	b.WriteString("\n\n")
	b.WriteString("Worksheet Tabs:\n")
	b.WriteString("\n")

	if modal.Loading {
		b.WriteString(Styles.Loading.Render("Loading tabs..."))
	} else if len(modal.Tabs) == 0 {
		b.WriteString(Styles.EmptyState.Render("(no tabs found)"))
	} else {
		// List tabs
		for i, tab := range modal.Tabs {
			cursor := "  "
			if i == modal.Cursor {
				cursor = "> "
			}
			line := fmt.Sprintf("%s%s", cursor, tab.Title)
			if i == modal.Cursor {
				b.WriteString(Styles.SelectedItem.Render(line))
			} else {
				b.WriteString(line)
			}
			b.WriteString("\n")
		}
	}

	b.WriteString("\n")
	if modal.PreviewLoading {
		b.WriteString("\n")
		b.WriteString(Styles.Header.Render("Preview"))
		b.WriteString("\n")
		b.WriteString(Styles.Loading.Render("Loading preview..."))
		b.WriteString("\n")
	} else if modal.PreviewContent != "" {
		b.WriteString("\n")
		b.WriteString(Styles.Header.Render(fmt.Sprintf("Preview: %s", modal.PreviewTab)))
		b.WriteString("\n")
		b.WriteString(Styles.Preview.Render(modal.PreviewContent))
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(Styles.KindLabel.Render("CSV edit is value-based; formulas and formatting may not preserve."))
	b.WriteString("\n")
	b.WriteString(Styles.KindLabel.Render("Large tabs may open slowly in the external editor."))
	b.WriteString("\n\n")
	b.WriteString(Styles.KindLabel.Render("[j/k] nav • [enter] preview • [e] edit in sheets • [q/esc] close"))

	return b.String()
}
