// Package editor provides integration with external editor programs.
package editor

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// Bridge defines the interface for launching external editors.
type Bridge interface {
	// EditCSV opens the given CSV file in an external editor.
	// Returns when the editor exits.
	EditCSV(ctx context.Context, path string) (Result, error)
}

// Result contains information about the editing session.
type Result struct {
	// Changed indicates if the file was modified.
	Changed bool
	// BeforeHash is the file hash before editing (if computed).
	BeforeHash string
	// AfterHash is the file hash after editing.
	AfterHash string
	// ExitCode from the editor process.
	ExitCode int
	// Duration of the editing session.
	Duration time.Duration
	// EditorError if the editor failed to launch or crashed.
	EditorError error
	// RowCount of the sheet (for UI feedback)
	RowCount int
	// LargeSheet warning flag
	LargeSheet bool
}

// EditSession tracks an in-progress or completed edit.
type EditSession struct {
	FileID     string
	TabName    string
	TempPath   string
	Revision   string // Google Drive revision ID at export time
	BeforeHash string
	ExportedAt time.Time
}

// SheetsBridge launches the `sheets` terminal spreadsheet editor.
type SheetsBridge struct {
	// EditorPath is the command to launch sheets (default: "sheets")
	EditorPath string
	// Timeout for the editing session (0 = no timeout)
	Timeout time.Duration
}

// NewSheetsBridge creates a new sheets editor bridge.
func NewSheetsBridge() *SheetsBridge {
	return &SheetsBridge{
		EditorPath: "sheets",
		Timeout:    0, // No timeout by default
	}
}

// ErrEditorNotFound is returned when the editor binary is not in PATH.
var ErrEditorNotFound = errors.New("editor not found in PATH")

// PreEditCheck performs pre-flight checks before editing.
// Returns row count and warnings if any.
func (b *SheetsBridge) PreEditCheck(ctx context.Context, path string) (rowCount int, warnings []string, err error) {
	// Check editor exists
	if err := CheckEditor(b.EditorPath); err != nil {
		return 0, nil, err
	}

	// Check file exists and get stats
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, nil, fmt.Errorf("cannot read file: %w", err)
	}

	// Count rows (approximate from newlines)
	rowCount = strings.Count(string(data), "\n")

	// Warn if large
	if rowCount > 5000 {
		warnings = append(warnings, fmt.Sprintf("Large sheet: %d rows (editing may be slow)", rowCount))
	}

	return rowCount, warnings, nil
}

// EditCSV opens the CSV file in the sheets editor.
func (b *SheetsBridge) EditCSV(ctx context.Context, path string) (Result, error) {
	startTime := time.Now()

	// Pre-flight checks
	rowCount, warnings, err := b.PreEditCheck(ctx, path)
	if err != nil {
		return Result{
			Changed:     false,
			EditorError: err,
		}, err
	}

	// Compute initial hash
	beforeHash, err := computeFileHash(path)
	if err != nil {
		// Non-fatal, just track that we couldn't compute hash
		beforeHash = ""
	}

	// Prepare command
	cmd := exec.CommandContext(ctx, b.EditorPath, path)

	// Set up I/O
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Launch editor
	err = cmd.Run()
	duration := time.Since(startTime)

	exitCode := 0
	var editorError error
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
			// Exit code 0 is success, anything else may indicate error
			// But some editors return non-zero on normal quit, so we don't treat it as fatal
		} else {
			editorError = fmt.Errorf("editor failed: %w", err)
		}
	}

	// Compute after hash
	afterHash, err := computeFileHash(path)
	if err != nil {
		afterHash = ""
	}

	// Detect changes
	changed := beforeHash != afterHash && beforeHash != "" && afterHash != ""

	result := Result{
		Changed:     changed,
		BeforeHash:  beforeHash,
		AfterHash:   afterHash,
		ExitCode:    exitCode,
		Duration:    duration,
		EditorError: editorError,
		RowCount:    rowCount,
		LargeSheet:  rowCount > 5000,
	}

	// Include warnings in error if present
	if len(warnings) > 0 && result.EditorError == nil {
		result.EditorError = fmt.Errorf("warnings: %s", strings.Join(warnings, "; "))
	}

	return result, nil
}

// Command builds the external editor command for Bubble Tea process handoff.
func (b *SheetsBridge) Command(ctx context.Context, path string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, b.EditorPath, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

// ComputeFileHash computes a SHA256 hash of the file contents.
func ComputeFileHash(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}

func computeFileHash(path string) (string, error) {
	return ComputeFileHash(path)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// CheckEditor verifies the editor binary exists in PATH.
func CheckEditor(editorPath string) error {
	_, err := exec.LookPath(editorPath)
	if err != nil {
		return fmt.Errorf("%w: '%s' not found in PATH. Install from: https://github.com/maaslalani/sheets", ErrEditorNotFound, editorPath)
	}
	return nil
}

// SanitizeFilename creates a safe filename from a sheet tab name.
func SanitizeFilename(name string) string {
	// Replace unsafe characters with underscores
	unsafe := regexp.MustCompile(`[<>:"/\\|?*\x00-\x1f]`)
	safe := unsafe.ReplaceAllString(name, "_")
	// Limit length
	if len(safe) > 50 {
		safe = safe[:50]
	}
	// Ensure not empty
	if strings.TrimSpace(safe) == "" {
		safe = "sheet"
	}
	return safe
}

// MakeEditPath creates a safe temp path for editing.
func MakeEditPath(baseDir, fileID, tabName string) string {
	safeTab := SanitizeFilename(tabName)
	timestamp := time.Now().Unix()
	return filepath.Join(baseDir, fileID, fmt.Sprintf("%s-%d.csv", safeTab, timestamp))
}
