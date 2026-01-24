package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dl-alexandre/gdrive/internal/types"
	"github.com/dl-alexandre/gdrive/internal/utils"
	"github.com/google/uuid"
	"github.com/olekukonko/tablewriter"
)

// OutputWriter handles CLI output formatting
type OutputWriter struct {
	format   types.OutputFormat
	quiet    bool
	verbose  bool
	warnings []types.CLIWarning
}

// NewOutputWriter creates a new output writer
func NewOutputWriter(format types.OutputFormat, quiet, verbose bool) *OutputWriter {
	return &OutputWriter{
		format:   format,
		quiet:    quiet,
		verbose:  verbose,
		warnings: []types.CLIWarning{},
	}
}

// AddWarning adds a warning to the output
func (w *OutputWriter) AddWarning(code, message, severity string) {
	w.warnings = append(w.warnings, types.CLIWarning{
		Code:     code,
		Message:  message,
		Severity: severity,
	})
}

// WriteSuccess writes a successful result
func (w *OutputWriter) WriteSuccess(command string, data interface{}) error {
	output := types.CLIOutput{
		SchemaVersion: utils.SchemaVersion,
		TraceID:       uuid.New().String(),
		Command:       command,
		Data:          data,
		Warnings:      w.warnings,
		Errors:        []types.CLIError{},
	}

	if w.format == types.OutputFormatJSON {
		return w.writeJSON(output)
	}
	return w.writeTable(data)
}

// WriteError writes an error result
func (w *OutputWriter) WriteError(command string, cliErr types.CLIError) error {
	output := types.CLIOutput{
		SchemaVersion: utils.SchemaVersion,
		TraceID:       uuid.New().String(),
		Command:       command,
		Data:          nil,
		Warnings:      w.warnings,
		Errors:        []types.CLIError{cliErr},
	}

	return w.writeJSON(output)
}

func (w *OutputWriter) writeJSON(output types.CLIOutput) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(output)
}

func (w *OutputWriter) writeTable(data interface{}) error {
	if renderable, ok := data.(types.TableRenderable); ok {
		return w.renderTable(renderable.AsTableRenderer())
	}
	if renderer, ok := data.(types.TableRenderer); ok {
		return w.renderTable(renderer)
	}
	switch v := data.(type) {
	case []*types.DriveFile:
		return w.writeFileTable(v)
	case *types.DriveFile:
		return w.writeFileTable([]*types.DriveFile{v})
	case []*types.Permission:
		return w.writePermissionTable(v)
	default:
		// Fallback to JSON for unknown types
		return w.writeJSON(types.CLIOutput{
			SchemaVersion: utils.SchemaVersion,
			TraceID:       uuid.New().String(),
			Command:       "unknown",
			Data:          data,
			Warnings:      w.warnings,
			Errors:        []types.CLIError{},
		})
	}
}

func (w *OutputWriter) renderTable(renderer types.TableRenderer) error {
	rows := renderer.Rows()
	if len(rows) == 0 {
		if !w.quiet {
			fmt.Fprintln(os.Stdout, renderer.EmptyMessage())
		}
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(renderer.Headers())
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, row := range rows {
		table.Append(row)
	}

	table.Render()
	return nil
}

func (w *OutputWriter) writeFileTable(files []*types.DriveFile) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Type", "Size", "Modified"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, f := range files {
		size := "-"
		if f.Size > 0 {
			size = formatSize(f.Size)
		}
		table.Append([]string{
			truncate(f.ID, 15),
			truncate(f.Name, 40),
			truncate(f.MimeType, 30),
			size,
			f.ModifiedTime,
		})
	}

	table.Render()
	return nil
}

func (w *OutputWriter) writePermissionTable(perms []*types.Permission) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Type", "Role", "Email/Domain"})
	table.SetBorder(false)

	for _, p := range perms {
		identity := p.EmailAddress
		if identity == "" {
			identity = p.Domain
		}
		if identity == "" {
			identity = "-"
		}
		table.Append([]string{p.ID, p.Type, p.Role, identity})
	}

	table.Render()
	return nil
}

// Log writes to stderr if not quiet
func (w *OutputWriter) Log(format string, args ...interface{}) {
	if !w.quiet {
		fmt.Fprintf(os.Stderr, format+"\n", args...)
	}
}

// Verbose writes to stderr if verbose is enabled
func (w *OutputWriter) Verbose(format string, args ...interface{}) {
	if w.verbose {
		fmt.Fprintf(os.Stderr, "[VERBOSE] "+format+"\n", args...)
	}
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}

func formatSize(bytes int64) string {
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
