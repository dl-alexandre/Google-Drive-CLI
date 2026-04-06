package tui

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dl-alexandre/gdrv/internal/tui/editor"
	"github.com/dl-alexandre/gdrv/internal/tui/vfs"
)

var checkEditor = editor.CheckEditor

// loadRootsCmd loads the root locations.
func loadRootsCmd(vfs vfs.VFS) tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		roots, err := vfs.Roots(ctx)
		if err != nil {
			return ErrMsg{Err: err}
		}
		return RootsMsg{Roots: roots}
	}
}

// loadDirCmd loads directory contents.
func loadDirCmd(vfs vfs.VFS, parentID string) tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		items, err := vfs.List(ctx, parentID)
		if err != nil {
			return ErrMsg{Err: err}
		}
		return LoadDirMsg{
			ParentID: parentID,
			Items:    items,
		}
	}
}

// statPreviewCmd loads preview for a node.
func statPreviewCmd(vfs vfs.VFS, id string) tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		node, err := vfs.Stat(ctx, id)
		if err != nil {
			return ErrMsg{Err: err}
		}

		content := formatNodePreview(node)
		return PreviewMsg{
			NodeID:  id,
			Content: content,
		}
	}
}

// formatNodePreview creates simple metadata preview.
func formatNodePreview(node vfs.Node) string {
	icon := getIcon(node.Kind)
	modified := node.Modified.Format("2006-01-02 15:04")
	if node.Modified.IsZero() {
		modified = "—"
	}

	size := humanSize(node.Size)
	if node.IsDir {
		size = "—"
	}

	id := node.ID
	if len(id) > 8 {
		id = id[:8] + "..."
	}

	return icon + " " + node.Name + "\n\n" +
		"Type: " + string(node.Kind) + "\n" +
		"Modified: " + modified + "\n" +
		"Size: " + size + "\n" +
		"ID: " + id
}

func getIcon(kind vfs.NodeKind) string {
	switch kind {
	case vfs.KindFolder:
		return ""
	case vfs.KindSheet:
		return ""
	case vfs.KindDoc:
		return ""
	case vfs.KindPDF:
		return ""
	case vfs.KindImage:
		return ""
	default:
		return ""
	}
}

func humanSize(bytes int64) string {
	if bytes == 0 {
		return "0 B"
	}
	if bytes < 1024 {
		return "1 KB"
	}
	return "1 KB" // Simplified for now
}

// loadSheetTabsCmd loads worksheet tabs from a spreadsheet.
func loadSheetTabsCmd(vfs vfs.VFS, spreadsheetID string) tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		tabs, err := vfs.ListSheetTabs(ctx, spreadsheetID)
		if err != nil {
			return ErrMsg{Err: err}
		}
		return LoadSheetTabsMsg{
			SpreadsheetID: spreadsheetID,
			Tabs:          tabs,
		}
	}
}

// previewSheetTabCmd loads a small read-only preview for a worksheet tab.
func previewSheetTabCmd(vfs vfs.VFS, spreadsheetID, tabName string) tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		csvData, _, err := vfs.ExportSheetTab(ctx, spreadsheetID, tabName)
		if err != nil {
			return ErrMsg{Err: fmt.Errorf("preview failed: %w", err)}
		}

		content, err := formatSheetPreview(csvData)
		if err != nil {
			return ErrMsg{Err: fmt.Errorf("preview parse failed: %w", err)}
		}

		return SheetPreviewMsg{
			SpreadsheetID: spreadsheetID,
			TabName:       tabName,
			Content:       content,
		}
	}
}

// prepareSheetEditCmd exports a sheet tab and prepares a temp editing session.
func prepareSheetEditCmd(vfs vfs.VFS, spreadsheetID, tabName string) tea.Cmd {
	return func() tea.Msg {
		if err := checkEditor("sheets"); err != nil {
			return ErrMsg{Err: fmt.Errorf("editor not available: %w", err)}
		}

		tempPath := editor.MakeEditPath(filepath.Join(os.TempDir(), "gdrv-edit"), spreadsheetID, tabName)
		tempDir := filepath.Dir(tempPath)
		if err := os.MkdirAll(tempDir, 0o755); err != nil {
			return ErrMsg{Err: fmt.Errorf("failed to create temp dir: %w", err)}
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		csvData, revision, err := vfs.ExportSheetTab(ctx, spreadsheetID, tabName)
		if err != nil {
			return ErrMsg{Err: fmt.Errorf("export failed: %w", err)}
		}

		if err := os.WriteFile(tempPath, []byte(csvData), 0o644); err != nil {
			return ErrMsg{Err: fmt.Errorf("failed to write temp file: %w", err)}
		}

		beforeHash, err := editor.ComputeFileHash(tempPath)
		if err != nil {
			_ = os.Remove(tempPath)
			return ErrMsg{Err: fmt.Errorf("failed to hash temp file: %w", err)}
		}

		bridge := editor.NewSheetsBridge()
		rowCount, warnings, err := bridge.PreEditCheck(context.Background(), tempPath)
		if err != nil {
			_ = os.Remove(tempPath)
			return ErrMsg{Err: err}
		}

		warnings = append([]string{"CSV round-trip edits values only; formatting may not preserve."}, warnings...)
		hasFormulas, err := vfs.TabHasFormulas(ctx, spreadsheetID, tabName)
		if err == nil && hasFormulas {
			warnings = append([]string{"This tab contains formulas; editing in CSV mode will convert them to values."}, warnings...)
		}

		return EditPreparedMsg{
			Session: editor.EditSession{
				FileID:     spreadsheetID,
				TabName:    tabName,
				TempPath:   tempPath,
				Revision:   revision,
				BeforeHash: beforeHash,
				ExportedAt: time.Now(),
			},
			RowCount: rowCount,
			Warnings: warnings,
		}
	}
}

// finalizeSheetEditCmd imports edited CSV back into the sheet if it changed.
func finalizeSheetEditCmd(vfs vfs.VFS, session editor.EditSession, duration time.Duration) tea.Cmd {
	return func() tea.Msg {
		defer func() { _ = os.Remove(session.TempPath) }()

		afterHash, err := editor.ComputeFileHash(session.TempPath)
		if err != nil {
			return ErrMsg{Err: fmt.Errorf("failed to read edited file: %w", err)}
		}

		if session.BeforeHash == afterHash {
			return EditCompleteMsg{
				SpreadsheetID: session.FileID,
				TabName:       session.TabName,
				Changed:       false,
				Duration:      duration,
				Message:       "No changes detected",
			}
		}

		editedData, err := os.ReadFile(session.TempPath)
		if err != nil {
			return ErrMsg{Err: fmt.Errorf("failed to read edited file: %w", err)}
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		if err := vfs.ImportSheetTab(ctx, session.FileID, session.TabName, string(editedData), session.Revision); err != nil {
			if isConflictError(err) {
				return ErrMsg{Err: fmt.Errorf("conflict detected: %w. Re-export and edit again, or cancel.", err)}
			}
			return ErrMsg{Err: fmt.Errorf("import failed: %w", err)}
		}

		return EditCompleteMsg{
			SpreadsheetID: session.FileID,
			TabName:       session.TabName,
			Changed:       true,
			Duration:      duration,
			Message:       fmt.Sprintf("Synced changes for %s", session.TabName),
		}
	}
}

func compactWarnings(warnings []string) string {
	return strings.Join(warnings, " ")
}

func formatSheetPreview(csvData string) (string, error) {
	reader := csv.NewReader(strings.NewReader(csvData))
	records, err := reader.ReadAll()
	if err != nil {
		return "", err
	}
	if len(records) == 0 {
		return "(empty tab)", nil
	}

	const maxRows = 10
	const maxCols = 6

	rows := records
	if len(rows) > maxRows {
		rows = rows[:maxRows]
	}

	widths := make([]int, maxCols)
	for _, row := range rows {
		for col := 0; col < len(row) && col < maxCols; col++ {
			cell := row[col]
			if len(cell) > 24 {
				cell = cell[:21] + "..."
			}
			if len(cell) > widths[col] {
				widths[col] = len(cell)
			}
		}
	}

	var b strings.Builder
	for rowIdx, row := range rows {
		for col := 0; col < len(row) && col < maxCols; col++ {
			cell := row[col]
			if len(cell) > 24 {
				cell = cell[:21] + "..."
			}
			b.WriteString(padRight(cell, widths[col]))
			if col < min(maxCols, len(row))-1 {
				b.WriteString("  ")
			}
		}
		if rowIdx == 0 {
			b.WriteString("\n")
			for col := 0; col < min(maxCols, len(row)); col++ {
				b.WriteString(strings.Repeat("-", widths[col]))
				if col < min(maxCols, len(row))-1 {
					b.WriteString("  ")
				}
			}
		}
		if rowIdx < len(rows)-1 {
			b.WriteString("\n")
		}
	}

	if len(records) > maxRows || maxColumnCount(records) > maxCols {
		b.WriteString("\n...")
	}

	return b.String(), nil
}

func padRight(s string, width int) string {
	if len(s) >= width {
		return s
	}
	return s + strings.Repeat(" ", width-len(s))
}

func maxColumnCount(records [][]string) int {
	maxCols := 0
	for _, row := range records {
		if len(row) > maxCols {
			maxCols = len(row)
		}
	}
	return maxCols
}
