// Package vfs provides a virtual filesystem abstraction for the TUI.
package vfs

import (
	"context"
	"time"
)

// Node represents a file or folder in the VFS.
type Node struct {
	ID       string
	ParentID string
	Name     string
	Kind     NodeKind
	MimeType string
	Modified time.Time
	Size     int64
	Owner    string
	Starred  bool
	Trashed  bool
	WebURL   string
	DriveID  string
	IsDir    bool
}

// NodeKind categorizes file types for dispatch and icon selection.
type NodeKind string

const (
	KindFolder  NodeKind = "folder"
	KindSheet   NodeKind = "sheet"
	KindDoc     NodeKind = "doc"
	KindSlide   NodeKind = "slide"
	KindPDF     NodeKind = "pdf"
	KindImage   NodeKind = "image"
	KindVideo   NodeKind = "video"
	KindBinary  NodeKind = "binary"
	KindUnknown NodeKind = "unknown"
)

// VFS defines the interface for file system operations.
type VFS interface {
	// Roots returns the top-level locations (e.g., My Drive).
	Roots(ctx context.Context) ([]Node, error)

	// List returns children of a parent ID.
	List(ctx context.Context, parentID string) ([]Node, error)

	// Stat returns detailed metadata for a single node.
	Stat(ctx context.Context, id string) (Node, error)

	// Rename changes a node's name.
	Rename(ctx context.Context, id, name string) error

	// Move changes a node's parent.
	Move(ctx context.Context, id, newParent string) error

	// Delete permanently removes a node.
	Delete(ctx context.Context, id string) error

	// ListSheetTabs returns worksheets in a Google Spreadsheet.
	ListSheetTabs(ctx context.Context, spreadsheetID string) ([]SheetTab, error)

	// ExportSheetTab exports a worksheet to CSV format.
	// Returns the CSV data and the current revision ID for conflict detection.
	ExportSheetTab(ctx context.Context, spreadsheetID, tabName string) (data string, revision string, err error)

	// ImportSheetTab imports CSV data into a worksheet.
	// The data should replace existing content in the tab.
	// Returns error if revision mismatch (sheet changed remotely).
	ImportSheetTab(ctx context.Context, spreadsheetID, tabName string, csvData string, expectedRevision string) error

	// GetFileRevision returns the current revision ID of a file.
	GetFileRevision(ctx context.Context, fileID string) (string, error)

	// TabHasFormulas reports whether the worksheet contains formula cells.
	TabHasFormulas(ctx context.Context, spreadsheetID, tabName string) (bool, error)
}

// DetectKind maps MIME types to NodeKind.
func DetectKind(mimeType string) NodeKind {
	switch mimeType {
	case "application/vnd.google-apps.folder":
		return KindFolder
	case "application/vnd.google-apps.spreadsheet":
		return KindSheet
	case "application/vnd.google-apps.document":
		return KindDoc
	case "application/vnd.google-apps.presentation":
		return KindSlide
	case "application/pdf":
		return KindPDF
	default:
		if isImageMime(mimeType) {
			return KindImage
		}
		if isVideoMime(mimeType) {
			return KindVideo
		}
		return KindBinary
	}
}

func isImageMime(m string) bool {
	return len(m) > 6 && m[:6] == "image/"
}

func isVideoMime(m string) bool {
	return len(m) > 6 && m[:6] == "video/"
}
