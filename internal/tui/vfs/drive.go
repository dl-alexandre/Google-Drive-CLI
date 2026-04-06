// Package vfs provides Google Drive implementation of the VFS interface.
package vfs

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"strings"
	"time"

	"github.com/dl-alexandre/gdrv/internal/files"
	"github.com/dl-alexandre/gdrv/internal/sheets"
	"github.com/dl-alexandre/gdrv/internal/types"
	"google.golang.org/api/drive/v3"
)

// DriveVFS wraps the existing managers to implement VFS.
type DriveVFS struct {
	files   *files.Manager
	sheets  *sheets.Manager
	driveID string
}

// NewDriveVFS creates a VFS backed by Google Drive.
func NewDriveVFS(files *files.Manager, sheets *sheets.Manager, driveID string) *DriveVFS {
	return &DriveVFS{
		files:   files,
		sheets:  sheets,
		driveID: driveID,
	}
}

func (d *DriveVFS) requestContext() *types.RequestContext {
	return &types.RequestContext{DriveID: d.driveID}
}

// Roots returns My Drive as the initial root.
func (d *DriveVFS) Roots(ctx context.Context) ([]Node, error) {
	if d.driveID != "" {
		return []Node{{
			ID:    d.driveID,
			Name:  "Shared Drive",
			Kind:  KindFolder,
			IsDir: true,
		}}, nil
	}

	// "root" is the special ID for My Drive
	return []Node{{
		ID:    "root",
		Name:  "My Drive",
		Kind:  KindFolder,
		IsDir: true,
	}}, nil
}

// List retrieves children using the files manager.
func (d *DriveVFS) List(ctx context.Context, parentID string) ([]Node, error) {
	reqCtx := d.requestContext()
	opts := files.ListOptions{
		ParentID:       parentID,
		IncludeTrashed: false,
		PageSize:       100,
		Fields:         "id,name,mimeType,modifiedTime,size,owners,starred,trashed,webViewLink,driveId,parents",
	}

	result, err := d.files.List(ctx, reqCtx, opts)
	if err != nil {
		return nil, err
	}

	nodes := make([]Node, len(result.Files))
	for i, f := range result.Files {
		nodes[i] = convertDriveFile(f)
	}
	return nodes, nil
}

// Stat retrieves metadata for a single file.
func (d *DriveVFS) Stat(ctx context.Context, id string) (Node, error) {
	reqCtx := d.requestContext()
	f, err := d.files.Get(ctx, reqCtx, id, "id,name,mimeType,modifiedTime,size,owners,starred,trashed,webViewLink,driveId,parents")
	if err != nil {
		return Node{}, err
	}
	return convertDriveFile(f), nil
}

// Rename changes a file's name.
func (d *DriveVFS) Rename(ctx context.Context, id, name string) error {
	reqCtx := d.requestContext()
	metadata := &drive.File{Name: name}
	_, err := d.files.Update(ctx, reqCtx, id, metadata, "name")
	return err
}

// Move changes a file's parent.
func (d *DriveVFS) Move(ctx context.Context, id, newParent string) error {
	reqCtx := d.requestContext()
	// TODO: Need proper error handling wrapper
	_, err := d.files.Move(ctx, reqCtx, id, newParent)
	return err
}

// Delete permanently removes a file.
func (d *DriveVFS) Delete(ctx context.Context, id string) error {
	reqCtx := d.requestContext()
	return d.files.Delete(ctx, reqCtx, id, true)
}

// ListSheetTabs retrieves worksheets from a Google Spreadsheet.
func (d *DriveVFS) ListSheetTabs(ctx context.Context, spreadsheetID string) ([]SheetTab, error) {
	if d.sheets == nil {
		return nil, nil // No sheets service available
	}

	reqCtx := d.requestContext()
	spreadsheet, err := d.sheets.GetSpreadsheet(ctx, reqCtx, spreadsheetID)
	if err != nil {
		return nil, err
	}

	if spreadsheet == nil || len(spreadsheet.Sheets) == 0 {
		return []SheetTab{}, nil
	}

	tabs := make([]SheetTab, 0, len(spreadsheet.Sheets))
	for i, sheet := range spreadsheet.Sheets {
		tabs = append(tabs, SheetTab{
			SheetID:   sheet.ID,
			Title:     sheet.Title,
			Index:     i,
			SheetType: sheet.Type,
		})
	}

	return tabs, nil
}

// convertDriveFile converts API type to VFS Node.
func convertDriveFile(f *types.DriveFile) Node {
	modified, _ := time.Parse(time.RFC3339, f.ModifiedTime)

	return Node{
		ID:       f.ID,
		ParentID: firstOrEmpty(f.Parents),
		Name:     f.Name,
		Kind:     DetectKind(f.MimeType),
		MimeType: f.MimeType,
		Modified: modified,
		Size:     f.Size,
		Owner:    "", // Not available in current DriveFile type
		Trashed:  f.Trashed,
		WebURL:   f.WebViewLink,
		IsDir:    f.MimeType == "application/vnd.google-apps.folder",
	}
}

func firstOrEmpty(arr []string) string {
	if len(arr) > 0 {
		return arr[0]
	}
	return ""
}

// ExportSheetTab exports a worksheet to CSV using the Drive API export functionality.
// Returns the CSV data and the file's current revision (modified time) for conflict detection.
func (d *DriveVFS) ExportSheetTab(ctx context.Context, spreadsheetID, tabName string) (string, string, error) {
	reqCtx := d.requestContext()

	// Get current revision (modified time) for conflict detection
	revision, err := d.GetFileRevision(ctx, spreadsheetID)
	if err != nil {
		// Non-fatal: continue without revision tracking
		revision = ""
	}

	// Verify the tab exists before exporting.
	found := false
	tabs, err := d.ListSheetTabs(ctx, spreadsheetID)
	if err != nil {
		return "", "", fmt.Errorf("failed to list tabs: %w", err)
	}
	for _, tab := range tabs {
		if tab.Title == tabName {
			found = true
			break
		}
	}
	if !found {
		return "", "", fmt.Errorf("tab not found: %s", tabName)
	}

	csv, err := d.exportViaSheetsAPI(ctx, reqCtx, spreadsheetID, tabName)
	return csv, revision, err
}

// exportViaSheetsAPI reads sheet values and serializes them to CSV.
func (d *DriveVFS) exportViaSheetsAPI(ctx context.Context, reqCtx *types.RequestContext, spreadsheetID, tabName string) (string, error) {
	if d.sheets == nil {
		return "", fmt.Errorf("sheets service not available")
	}

	values, err := d.sheets.GetValues(ctx, reqCtx, spreadsheetID, quotedTabName(tabName))
	if err != nil {
		return "", fmt.Errorf("failed to read sheet values: %w", err)
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	for _, row := range values.Rows() {
		if err := writer.Write(row); err != nil {
			return "", fmt.Errorf("failed to write csv row: %w", err)
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return "", fmt.Errorf("failed to finalize csv export: %w", err)
	}

	return buf.String(), nil
}

// ImportSheetTab imports CSV data into a worksheet.
// If expectedRevision is provided, checks if file changed remotely first.
func (d *DriveVFS) ImportSheetTab(ctx context.Context, spreadsheetID, tabName string, csvData string, expectedRevision string) error {
	if d.sheets == nil {
		return fmt.Errorf("sheets service not available")
	}

	// Check for conflicts if revision tracking is enabled
	if expectedRevision != "" {
		currentRevision, err := d.GetFileRevision(ctx, spreadsheetID)
		if err == nil && currentRevision != expectedRevision {
			return fmt.Errorf("conflict: spreadsheet changed remotely since export (was %s, now %s)", expectedRevision, currentRevision)
		}
	}

	reqCtx := d.requestContext()
	oldValues, err := d.sheets.GetValues(ctx, reqCtx, spreadsheetID, quotedTabName(tabName))
	if err != nil {
		return fmt.Errorf("failed to read existing sheet values: %w", err)
	}
	oldRows, oldCols := sheetDimensions(oldValues.Rows())

	// Parse CSV data
	records, err := parseCSV(csvData)
	if err != nil {
		return fmt.Errorf("failed to parse CSV: %w", err)
	}

	if len(records) == 0 {
		if _, err := d.sheets.ClearValues(ctx, reqCtx, spreadsheetID, quotedTabName(tabName)); err != nil {
			return fmt.Errorf("failed to clear empty sheet values: %w", err)
		}
		return nil
	}

	// Convert to Sheets API value range
	values := make([][]interface{}, len(records))
	for i, record := range records {
		row := make([]interface{}, len(record))
		for j, cell := range record {
			row[j] = cell
		}
		values[i] = row
	}

	// Update values first so a failed write does not wipe the tab.
	range_ := fmt.Sprintf("%s!A1", quotedTabName(tabName))

	// Use USER_ENTERED to parse values as if user typed them (handles numbers, formulas, etc.)
	_, err = d.sheets.UpdateValues(ctx, reqCtx, spreadsheetID, range_, values, "USER_ENTERED")
	if err != nil {
		return fmt.Errorf("failed to update sheet values: %w", err)
	}

	newRows, newCols := sheetDimensions(records)
	if err := d.clearStaleCells(ctx, reqCtx, spreadsheetID, tabName, oldRows, oldCols, newRows, newCols); err != nil {
		return err
	}

	return nil
}

// parseCSV parses CSV string into records.
func parseCSV(data string) ([][]string, error) {
	reader := csv.NewReader(strings.NewReader(data))
	return reader.ReadAll()
}

func quotedTabName(tabName string) string {
	escaped := strings.ReplaceAll(tabName, "'", "''")
	return fmt.Sprintf("'%s'", escaped)
}

func sheetDimensions[T any](rows [][]T) (int, int) {
	rowCount := len(rows)
	colCount := 0
	for _, row := range rows {
		if len(row) > colCount {
			colCount = len(row)
		}
	}
	return rowCount, colCount
}

func (d *DriveVFS) clearStaleCells(ctx context.Context, reqCtx *types.RequestContext, spreadsheetID, tabName string, oldRows, oldCols, newRows, newCols int) error {
	quoted := quotedTabName(tabName)

	if oldCols > newCols && newCols > 0 && newRows > 0 {
		clearRange := fmt.Sprintf("%s!%s1:%s%d", quoted, columnLetter(newCols), columnLetter(oldCols-1), newRows)
		if _, err := d.sheets.ClearValues(ctx, reqCtx, spreadsheetID, clearRange); err != nil {
			return fmt.Errorf("failed to clear trailing columns: %w", err)
		}
	}

	if oldRows > newRows {
		endCol := oldCols
		if newCols > endCol {
			endCol = newCols
		}
		if endCol > 0 {
			clearRange := fmt.Sprintf("%s!A%d:%s%d", quoted, newRows+1, columnLetter(endCol-1), oldRows)
			if _, err := d.sheets.ClearValues(ctx, reqCtx, spreadsheetID, clearRange); err != nil {
				return fmt.Errorf("failed to clear trailing rows: %w", err)
			}
		}
	}

	return nil
}

func columnLetter(col int) string {
	result := ""
	for col >= 0 {
		result = string(rune('A'+(col%26))) + result
		col = col/26 - 1
	}
	return result
}

// GetFileRevision returns the current revision ID (ModifiedTime) of a file.
// For Google Workspace files, we use ModifiedTime as a proxy for revision since
// they don't have traditional binary revision IDs.
func (d *DriveVFS) GetFileRevision(ctx context.Context, fileID string) (string, error) {
	reqCtx := d.requestContext()
	file, err := d.files.Get(ctx, reqCtx, fileID, "id,modifiedTime")
	if err != nil {
		return "", err
	}
	// Use ModifiedTime as revision indicator - it changes on every edit
	return file.ModifiedTime, nil
}

// TabHasFormulas reports whether the worksheet contains formula cells.
func (d *DriveVFS) TabHasFormulas(ctx context.Context, spreadsheetID, tabName string) (bool, error) {
	if d.sheets == nil {
		return false, nil
	}
	reqCtx := d.requestContext()
	return d.sheets.HasFormulas(ctx, reqCtx, spreadsheetID, quotedTabName(tabName))
}
