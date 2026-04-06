package vfs

// SheetTab represents a worksheet within a Google Spreadsheet.
type SheetTab struct {
	SheetID   int64  // Unique sheet ID within the spreadsheet
	Title     string // Human-readable name
	Index     int    // 0-based position in the tab bar
	SheetType string // Sheet type (GRID, etc.)
}

// Empty checks if the tab is a placeholder.
func (t SheetTab) Empty() bool {
	return t.Title == ""
}
