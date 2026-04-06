package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// Styles holds all Lipgloss styles for the TUI.
var Styles = struct {
	App           lipgloss.Style
	BrowserPane   lipgloss.Style
	PreviewPane   lipgloss.Style
	StatusBar     lipgloss.Style
	ErrorBar      lipgloss.Style
	SelectedItem  lipgloss.Style
	NormalItem    lipgloss.Style
	DirectoryItem lipgloss.Style
	SheetItem     lipgloss.Style
	DocItem       lipgloss.Style
	MediaItem     lipgloss.Style
	KindLabel     lipgloss.Style
	EmptyState    lipgloss.Style
	Header        lipgloss.Style
	Loading       lipgloss.Style
	Preview       lipgloss.Style
}{
	App: lipgloss.NewStyle(),

	BrowserPane: lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("241")).
		Padding(0, 1),

	PreviewPane: lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("241")).
		Padding(0, 1),

	StatusBar: lipgloss.NewStyle().
		Background(lipgloss.Color("236")).
		Foreground(lipgloss.Color("250")).
		Padding(0, 1).
		Height(1),

	ErrorBar: lipgloss.NewStyle().
		Background(lipgloss.Color("196")).
		Foreground(lipgloss.Color("231")).
		Padding(0, 1).
		Height(1),

	SelectedItem: lipgloss.NewStyle().
		Background(lipgloss.Color("238")).
		Foreground(lipgloss.Color("255")),

	NormalItem: lipgloss.NewStyle().
		Foreground(lipgloss.Color("250")),

	DirectoryItem: lipgloss.NewStyle().
		Foreground(lipgloss.Color("39")).
		Bold(true),

	SheetItem: lipgloss.NewStyle().
		Foreground(lipgloss.Color("76")), // Green for Sheets

	DocItem: lipgloss.NewStyle().
		Foreground(lipgloss.Color("33")), // Blue for Docs

	MediaItem: lipgloss.NewStyle().
		Foreground(lipgloss.Color("172")), // Orange for media/PDFs

	KindLabel: lipgloss.NewStyle().
		Foreground(lipgloss.Color("245")).
		Faint(true),

	EmptyState: lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		Italic(true),

	Header: lipgloss.NewStyle().
		Foreground(lipgloss.Color("39")).
		Bold(true).
		MarginBottom(1),

	Loading: lipgloss.NewStyle().
		Foreground(lipgloss.Color("208")).
		Bold(true),

	Preview: lipgloss.NewStyle().
		Foreground(lipgloss.Color("250")).
		Padding(1),
}

// PaneMinWidth is the minimum width before hiding preview.
const PaneMinWidth = 100
