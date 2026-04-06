package cli

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/auth"
	"github.com/dl-alexandre/gdrv/internal/config"
	"github.com/dl-alexandre/gdrv/internal/files"
	"github.com/dl-alexandre/gdrv/internal/sheets"
	"github.com/dl-alexandre/gdrv/internal/tui"
	"github.com/dl-alexandre/gdrv/internal/tui/vfs"
	"github.com/dl-alexandre/gdrv/internal/utils"
)

// TUICmd launches the interactive TUI file browser.
type TUICmd struct {
	// Stub flag for development/testing
	Stub bool `help:"Use stub filesystem (for testing)"`
}

// Run executes the TUI command.
func (c *TUICmd) Run(globals *Globals) error {
	var vfsImpl vfs.VFS
	var err error

	if c.Stub {
		// Use stub VFS for testing
		vfsImpl = &stubVFS{}
	} else {
		// Use real Google Drive
		vfsImpl, err = createDriveVFS(globals)
		if err != nil {
			return fmt.Errorf("failed to initialize Drive connection: %w", err)
		}
	}

	// Create and run TUI
	model := tui.NewModel(vfsImpl)
	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		return fmt.Errorf("tui error: %w", err)
	}

	return nil
}

// createDriveVFS creates a real Google Drive-backed VFS.
func createDriveVFS(globals *Globals) (vfs.VFS, error) {
	ctx := context.Background()
	flags := globals.ToGlobalFlags()

	configDir := getTUIConfigDir()
	authMgr := auth.NewManager(configDir)

	// Get valid credentials for the profile
	creds, err := authMgr.GetValidCredentials(ctx, flags.Profile)
	if err != nil {
		return nil, fmt.Errorf("authentication failed: %w", err)
	}

	// Create Drive service
	driveService, err := authMgr.GetDriveService(ctx, creds)
	if err != nil {
		return nil, fmt.Errorf("failed to create Drive service: %w", err)
	}

	// Create Sheets service (for later sheet editing)
	sheetsService, err := authMgr.GetSheetsService(ctx, creds)
	if err != nil {
		// Non-fatal - sheets just won't work
		sheetsService = nil
	}

	// Create API client with retry logic
	client := api.NewClient(driveService, utils.DefaultMaxRetries, utils.DefaultRetryDelayMs, globals.Logger)

	// Create managers
	filesMgr := files.NewManager(client)
	var sheetsMgr *sheets.Manager
	if sheetsService != nil {
		sheetsMgr = sheets.NewManager(client, sheetsService)
	}

	// Create VFS adapter
	return vfs.NewDriveVFS(filesMgr, sheetsMgr, flags.DriveID), nil
}

// getTUIConfigDir returns the configuration directory.
func getTUIConfigDir() string {
	dir, err := config.GetConfigDir()
	if err == nil {
		return dir
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "gdrv")
}

// stubVFS implements vfs.VFS for testing and development.
type stubVFS struct{}

func (s *stubVFS) Roots(ctx context.Context) ([]vfs.Node, error) {
	return []vfs.Node{{
		ID:    "root",
		Name:  "My Drive (TEST MODE)",
		Kind:  vfs.KindFolder,
		IsDir: true,
	}}, nil
}

func (s *stubVFS) List(ctx context.Context, parentID string) ([]vfs.Node, error) {
	// Return different items based on parent
	switch parentID {
	case "root":
		return []vfs.Node{
			{ID: "folder1", ParentID: "root", Name: "Documents", Kind: vfs.KindFolder, IsDir: true, Modified: time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)},
			{ID: "folder2", ParentID: "root", Name: "Projects", Kind: vfs.KindFolder, IsDir: true, Modified: time.Date(2024, 3, 10, 0, 0, 0, 0, time.UTC)},
			{ID: "sheet1", ParentID: "root", Name: "budget-2026", Kind: vfs.KindSheet, MimeType: "application/vnd.google-apps.spreadsheet", Modified: time.Date(2024, 4, 3, 0, 0, 0, 0, time.UTC)},
			{ID: "doc1", ParentID: "root", Name: "notes.txt", Kind: vfs.KindBinary, Size: 1024, Modified: time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC)},
			{ID: "pdf1", ParentID: "root", Name: "report.pdf", Kind: vfs.KindPDF, MimeType: "application/pdf", Size: 45056, Modified: time.Date(2024, 3, 28, 0, 0, 0, 0, time.UTC)},
		}, nil
	case "folder1":
		return []vfs.Node{
			{ID: "doc2", ParentID: "folder1", Name: "meeting-notes.doc", Kind: vfs.KindDoc, MimeType: "application/vnd.google-apps.document", Modified: time.Date(2024, 2, 20, 0, 0, 0, 0, time.UTC)},
			{ID: "sheet2", ParentID: "folder1", Name: "expenses", Kind: vfs.KindSheet, MimeType: "application/vnd.google-apps.spreadsheet", Modified: time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)},
		}, nil
	case "folder2":
		return []vfs.Node{
			{ID: "image1", ParentID: "folder2", Name: "screenshot.png", Kind: vfs.KindImage, MimeType: "image/png", Size: 204800, Modified: time.Date(2024, 4, 5, 0, 0, 0, 0, time.UTC)},
		}, nil
	default:
		return []vfs.Node{}, nil
	}
}

func (s *stubVFS) Stat(ctx context.Context, id string) (vfs.Node, error) {
	return vfs.Node{
		ID:       id,
		Name:     "test-file",
		Kind:     vfs.KindBinary,
		Modified: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}, nil
}

func (s *stubVFS) Rename(ctx context.Context, id, name string) error    { return nil }
func (s *stubVFS) Move(ctx context.Context, id, newParent string) error { return nil }
func (s *stubVFS) Delete(ctx context.Context, id string) error          { return nil }
func (s *stubVFS) ListSheetTabs(ctx context.Context, spreadsheetID string) ([]vfs.SheetTab, error) {
	// Return fake tabs for the budget spreadsheet
	return []vfs.SheetTab{
		{SheetID: 0, Title: "Summary", Index: 0, SheetType: "GRID"},
		{SheetID: 1, Title: "Q1 Budget", Index: 1, SheetType: "GRID"},
		{SheetID: 2, Title: "Q2 Budget", Index: 2, SheetType: "GRID"},
		{SheetID: 3, Title: "Raw Data", Index: 3, SheetType: "GRID"},
	}, nil
}

// ExportSheetTab returns fake CSV data for testing.
// Returns CSV data and a fake revision for conflict testing.
func (s *stubVFS) ExportSheetTab(ctx context.Context, spreadsheetID, tabName string) (string, string, error) {
	// Return sample CSV data based on tab name
	var csvData string
	switch tabName {
	case "Summary":
		csvData = "Category,Amount,Percentage\nIncome,100000,100%\nExpenses,75000,75%\nProfit,25000,25%\n"
	case "Q1 Budget":
		csvData = "Month,Revenue,Costs\nJanuary,25000,18000\nFebruary,28000,19000\nMarch,32000,21000\n"
	case "Q2 Budget":
		csvData = "Month,Revenue,Costs\nApril,30000,20000\nMay,35000,22000\nJune,38000,24000\n"
	case "Raw Data":
		csvData = "ID,Name,Value\n1,Item A,100\n2,Item B,200\n3,Item C,300\n"
	default:
		csvData = "A,B,C\n1,2,3\n4,5,6\n"
	}
	// Return fake revision for conflict detection testing
	return csvData, "rev-2024-01-01T00:00:00Z", nil
}

// ImportSheetTab accepts CSV data but doesn't actually store it (stub).
// In stub mode, ignores expectedRevision (no real conflict detection).
func (s *stubVFS) ImportSheetTab(ctx context.Context, spreadsheetID, tabName string, csvData string, expectedRevision string) error {
	// In stub mode, we just pretend it worked
	// The edited CSV would be lost, but that's fine for UI testing
	_ = expectedRevision // Ignored in stub mode
	return nil
}

// GetFileRevision returns a fake revision for stub testing.
func (s *stubVFS) GetFileRevision(ctx context.Context, fileID string) (string, error) {
	return "rev-2024-01-01T00:00:00Z", nil
}

func (s *stubVFS) TabHasFormulas(ctx context.Context, spreadsheetID, tabName string) (bool, error) {
	return tabName == "Summary", nil
}
