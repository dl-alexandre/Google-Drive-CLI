package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/dl-alexandre/gdrv/internal/types"
	"github.com/dl-alexandre/gdrv/internal/utils"
	"github.com/google/uuid"
	"github.com/olekukonko/tablewriter"
)

// OutputFormatter handles output formatting for CLI commands
type OutputFormatter struct {
	format          types.OutputFormat
	quiet           bool
	verbose         bool
	includeTraceID  bool
	colorOutput     bool
	writer          io.Writer
	errorWriter     io.Writer
	warnings        []types.CLIWarning
}

// OutputOptions configures the output formatter
type OutputOptions struct {
	Format         types.OutputFormat
	Quiet          bool
	Verbose        bool
	IncludeTraceID bool
	ColorOutput    bool
}

// NewOutputFormatter creates a new output formatter
func NewOutputFormatter(opts OutputOptions) *OutputFormatter {
	return &OutputFormatter{
		format:          opts.Format,
		quiet:           opts.Quiet,
		verbose:         opts.Verbose,
		includeTraceID:  opts.IncludeTraceID,
		colorOutput:     opts.ColorOutput,
		writer:          os.Stdout,
		errorWriter:     os.Stderr,
		warnings:        []types.CLIWarning{},
	}
}

// AddWarning adds a warning to be included in output
func (f *OutputFormatter) AddWarning(code, message, severity string) {
	f.warnings = append(f.warnings, types.CLIWarning{
		Code:     code,
		Message:  message,
		Severity: severity,
	})
}

// WriteSuccess writes a successful result
func (f *OutputFormatter) WriteSuccess(command string, data interface{}) error {
	traceID := ""
	if f.verbose || f.includeTraceID {
		traceID = uuid.New().String()
	}

	output := types.CLIOutput{
		SchemaVersion: utils.SchemaVersion,
		TraceID:       traceID,
		Command:       command,
		Data:          data,
		Warnings:      f.warnings,
		Errors:        []types.CLIError{},
	}

	// In verbose mode, log trace ID to stderr
	if f.verbose && traceID != "" {
		f.Verbose("Trace ID: %s", traceID)
	}

	switch f.format {
	case types.OutputFormatJSON:
		return f.writeJSON(output)
	case types.OutputFormatTable:
		return f.writeTable(data)
	default:
		return fmt.Errorf("unsupported output format: %s", f.format)
	}
}

// WriteError writes an error result
func (f *OutputFormatter) WriteError(command string, cliErr types.CLIError) error {
	traceID := uuid.New().String()

	output := types.CLIOutput{
		SchemaVersion: utils.SchemaVersion,
		TraceID:       traceID,
		Command:       command,
		Data:          nil,
		Warnings:      f.warnings,
		Errors:        []types.CLIError{cliErr},
	}

	// Always output errors as JSON for structured parsing
	if err := f.writeJSON(output); err != nil {
		return err
	}

	// In verbose mode, also log to stderr
	if f.verbose {
		f.Verbose("Error occurred - Trace ID: %s", traceID)
	}

	return nil
}

// writeJSON writes data as JSON
func (f *OutputFormatter) writeJSON(data interface{}) error {
	encoder := json.NewEncoder(f.writer)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

// writeTable writes data in table format
func (f *OutputFormatter) writeTable(data interface{}) error {
	// Display warnings if any (to stderr)
	if len(f.warnings) > 0 && !f.quiet {
		for _, warning := range f.warnings {
			if _, err := fmt.Fprintf(f.errorWriter, "Warning [%s]: %s\n", warning.Code, warning.Message); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprintln(f.errorWriter); err != nil {
			return err
		}
	}

	if renderable, ok := data.(types.TableRenderable); ok {
		return f.renderTable(renderable.AsTableRenderer())
	}
	if renderer, ok := data.(types.TableRenderer); ok {
		return f.renderTable(renderer)
	}

	switch v := data.(type) {
	case []*types.DriveFile:
		return f.writeFileTable(v)
	case *types.DriveFile:
		return f.writeFileTable([]*types.DriveFile{v})
	case []*types.Permission:
		return f.writePermissionTable(v)
	case *types.Permission:
		return f.writePermissionTable([]*types.Permission{v})
	case *types.FileListResult:
		if err := f.writeFileTable(v.Files); err != nil {
			return err
		}
		if v.NextPageToken != "" {
			if _, err := fmt.Fprintf(f.errorWriter, "\nMore results available. Use --page-token %s to continue.\n", v.NextPageToken); err != nil {
				return err
			}
		}
		if v.IncompleteSearch {
			if _, err := fmt.Fprintln(f.errorWriter, "\nWarning: Search results may be incomplete."); err != nil {
				return err
			}
		}
		return nil
	case map[string]interface{}:
		// Generic key-value output for about/info commands
		return f.writeKeyValueTable(v)
	default:
		// Fallback to JSON for unknown types
		return f.writeJSON(types.CLIOutput{
			SchemaVersion: utils.SchemaVersion,
			TraceID:       "",
			Command:       "unknown",
			Data:          data,
			Warnings:      f.warnings,
			Errors:        []types.CLIError{},
		})
	}
}

func (f *OutputFormatter) renderTable(renderer types.TableRenderer) error {
	rows := renderer.Rows()
	if len(rows) == 0 {
		if !f.quiet {
			if _, err := fmt.Fprintln(f.writer, renderer.EmptyMessage()); err != nil {
				return err
			}
		}
		return nil
	}

	table := tablewriter.NewWriter(f.writer)
	table.SetHeader(renderer.Headers())
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)

	for _, row := range rows {
		table.Append(row)
	}

	table.Render()
	return nil
}

// writeFileTable writes file data as a table
func (f *OutputFormatter) writeFileTable(files []*types.DriveFile) error {
	if len(files) == 0 {
		if !f.quiet {
			if _, err := fmt.Fprintln(f.writer, "No files found."); err != nil {
				return err
			}
		}
		return nil
	}

	table := tablewriter.NewWriter(f.writer)
	
	// Configure table appearance
	table.SetHeader([]string{"ID", "Name", "Type", "Size", "Modified"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)

	for _, file := range files {
		size := formatFileSize(file.Size, file.MimeType)
		modTime := formatTime(file.ModifiedTime)
		
		table.Append([]string{
			truncateString(file.ID, 20),
			truncateString(file.Name, 50),
			formatMimeType(file.MimeType),
			size,
			modTime,
		})
	}

	table.Render()
	return nil
}

// writePermissionTable writes permission data as a table
func (f *OutputFormatter) writePermissionTable(perms []*types.Permission) error {
	if len(perms) == 0 {
		if !f.quiet {
			if _, err := fmt.Fprintln(f.writer, "No permissions found."); err != nil {
				return err
			}
		}
		return nil
	}

	table := tablewriter.NewWriter(f.writer)
	table.SetHeader([]string{"ID", "Type", "Role", "Identity", "Display Name"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, perm := range perms {
		identity := getPermissionIdentity(perm)
		displayName := perm.DisplayName
		if displayName == "" {
			displayName = "-"
		}
		
		table.Append([]string{
			truncateString(perm.ID, 20),
			perm.Type,
			perm.Role,
			identity,
			truncateString(displayName, 30),
		})
	}

	table.Render()
	return nil
}

// writeKeyValueTable writes a generic key-value table
func (f *OutputFormatter) writeKeyValueTable(data map[string]interface{}) error {
	table := tablewriter.NewWriter(f.writer)
	table.SetHeader([]string{"Key", "Value"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for key, value := range data {
		valueStr := fmt.Sprintf("%v", value)
		table.Append([]string{key, valueStr})
	}

	table.Render()
	return nil
}

// Log writes a message to stderr unless quiet mode is enabled
func (f *OutputFormatter) Log(format string, args ...interface{}) {
	if !f.quiet {
		if _, err := fmt.Fprintf(f.errorWriter, format+"\n", args...); err != nil {
			return
		}
	}
}

// Verbose writes a message to stderr only in verbose mode
func (f *OutputFormatter) Verbose(format string, args ...interface{}) {
	if f.verbose {
		if _, err := fmt.Fprintf(f.errorWriter, "[VERBOSE] "+format+"\n", args...); err != nil {
			return
		}
	}
}

// Debug writes a message to stderr only in debug mode
func (f *OutputFormatter) Debug(format string, args ...interface{}) {
	// Debug is treated as more verbose than verbose
	if f.verbose {
		if _, err := fmt.Fprintf(f.errorWriter, "[DEBUG] "+format+"\n", args...); err != nil {
			return
		}
	}
}

// Helper functions

// formatFileSize formats file size for display
func formatFileSize(bytes int64, mimeType string) string {
	// Google Docs/Sheets/Slides don't have meaningful size
	if isGoogleWorkspaceType(mimeType) {
		return "-"
	}

	if bytes == 0 {
		return "-"
	}

	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}

	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// formatTime formats timestamp for display
func formatTime(timestamp string) string {
	if timestamp == "" {
		return "-"
	}

	t, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return timestamp
	}

	// Show relative time for recent files
	now := time.Now()
	diff := now.Sub(t)

	if diff < 24*time.Hour {
		return t.Format("15:04 Today")
	} else if diff < 48*time.Hour {
		return t.Format("15:04 Yesterday")
	} else if diff < 7*24*time.Hour {
		return t.Format("Mon 15:04")
	}

	return t.Format("2006-01-02")
}

// formatMimeType formats MIME type for display
func formatMimeType(mimeType string) string {
	// Map common MIME types to readable names
	typeMap := map[string]string{
		"application/vnd.google-apps.folder":       "Folder",
		"application/vnd.google-apps.document":     "Doc",
		"application/vnd.google-apps.spreadsheet":  "Sheet",
		"application/vnd.google-apps.presentation": "Slides",
		"application/vnd.google-apps.form":         "Form",
		"application/vnd.google-apps.drawing":      "Drawing",
		"application/vnd.google-apps.script":       "Script",
		"application/vnd.google-apps.site":         "Site",
		"application/pdf":                          "PDF",
		"image/jpeg":                               "JPEG",
		"image/png":                                "PNG",
		"image/gif":                                "GIF",
		"text/plain":                               "Text",
		"text/html":                                "HTML",
		"application/zip":                          "ZIP",
	}

	if readable, ok := typeMap[mimeType]; ok {
		return readable
	}

	// For other types, show simplified version
	if strings.HasPrefix(mimeType, "application/vnd.google-apps.") {
		return strings.TrimPrefix(mimeType, "application/vnd.google-apps.")
	}

	// Extract main type
	parts := strings.Split(mimeType, "/")
	if len(parts) > 1 {
		return parts[1]
	}

	return mimeType
}

// isGoogleWorkspaceType checks if MIME type is a Google Workspace type
func isGoogleWorkspaceType(mimeType string) bool {
	return strings.HasPrefix(mimeType, "application/vnd.google-apps.")
}

// getPermissionIdentity returns the identity string for a permission
func getPermissionIdentity(perm *types.Permission) string {
	if perm.EmailAddress != "" {
		return perm.EmailAddress
	}
	if perm.Domain != "" {
		return perm.Domain
	}
	if perm.Type == "anyone" {
		return "Anyone"
	}
	return "-"
}

// truncateString truncates a string to maxLen with ellipsis
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return s[:maxLen]
	}
	return s[:maxLen-3] + "..."
}
