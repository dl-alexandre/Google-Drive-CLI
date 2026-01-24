---
name: Sheets CLI Integration
overview: Add comprehensive Google Sheets API support for reading, writing, appending, and batch updating spreadsheet data with consistent CLI patterns and robust error handling.
todos: []
isProject: false
status: completed
dependencies:
  - 00_foundation_google_apis.plan.md
---

> **Status**: ✅ **COMPLETED** - Sheets CLI integration has been successfully implemented.

## Context

The foundation plan provides:
- Scope management (ScopeSheets, ScopeSheetsReadonly already defined)
- Service factory pattern for creating sheets.Service
- Unified error handling and retry logic
- TableRenderer interface for output

Existing patterns:
- Auth scopes validated per command in `internal/auth/manager.go:94-102`
- Output formatting uses `OutputWriter` in `internal/cli/output.go:78-96`
- Drive file operations use manager pattern (see `internal/drives/manager.go`)

## Plan

### 1. Sheets Manager Layer

**Files to create**:
- `internal/sheets/manager.go`: Core Sheets operations wrapper
- `internal/sheets/manager_test.go`: Unit tests with mocks
- `internal/types/sheets.go`: Sheets-specific types

**Manager structure**:
```go
// internal/sheets/manager.go
package sheets

import (
    "context"
    "fmt"
    "google.golang.org/api/sheets/v4"
    "github.com/dl-alexandre/Google-Drive-CLI/internal/types"
    "github.com/dl-alexandre/Google-Drive-CLI/internal/errors"
    "github.com/dl-alexandre/Google-Drive-CLI/internal/retry"
)

type Manager struct {
    service *sheets.Service
    retryConfig *retry.RetryConfig
}

func NewManager(service *sheets.Service) *Manager {
    return &Manager{
        service: service,
        retryConfig: retry.DefaultRetryConfig(),
    }
}

// GetValues reads values from a range
func (m *Manager) GetValues(ctx context.Context, spreadsheetID, rangeNotation string) (*types.SheetValues, error) {
    var resp *sheets.ValueRange
    err := retry.WithRetry(ctx, m.retryConfig, func() error {
        var err error
        resp, err = m.service.Spreadsheets.Values.Get(spreadsheetID, rangeNotation).Context(ctx).Do()
        return err
    })

    if err != nil {
        return nil, errors.ParseGoogleAPIError(err, "Sheets")
    }

    return &types.SheetValues{
        Range:          resp.Range,
        MajorDimension: resp.MajorDimension,
        Values:         resp.Values,
    }, nil
}

// UpdateValues updates a range with new values
func (m *Manager) UpdateValues(ctx context.Context, req *types.UpdateValuesRequest) (*types.UpdateValuesResponse, error) {
    valueRange := &sheets.ValueRange{
        Values: req.Values,
        MajorDimension: req.MajorDimension,
    }

    var resp *sheets.UpdateValuesResponse
    err := retry.WithRetry(ctx, m.retryConfig, func() error {
        var err error
        resp, err = m.service.Spreadsheets.Values.
            Update(req.SpreadsheetID, req.Range, valueRange).
            ValueInputOption(req.ValueInputOption).
            Context(ctx).
            Do()
        return err
    })

    if err != nil {
        return nil, errors.ParseGoogleAPIError(err, "Sheets")
    }

    return &types.UpdateValuesResponse{
        SpreadsheetID:  resp.SpreadsheetId,
        UpdatedRange:   resp.UpdatedRange,
        UpdatedRows:    int(resp.UpdatedRows),
        UpdatedColumns: int(resp.UpdatedColumns),
        UpdatedCells:   int(resp.UpdatedCells),
    }, nil
}

// AppendValues appends values to a range
func (m *Manager) AppendValues(ctx context.Context, req *types.AppendValuesRequest) (*types.UpdateValuesResponse, error) {
    valueRange := &sheets.ValueRange{
        Values: req.Values,
        MajorDimension: req.MajorDimension,
    }

    var resp *sheets.AppendValuesResponse
    err := retry.WithRetry(ctx, m.retryConfig, func() error {
        var err error
        resp, err = m.service.Spreadsheets.Values.
            Append(req.SpreadsheetID, req.Range, valueRange).
            ValueInputOption(req.ValueInputOption).
            Context(ctx).
            Do()
        return err
    })

    if err != nil {
        return nil, errors.ParseGoogleAPIError(err, "Sheets")
    }

    return &types.UpdateValuesResponse{
        SpreadsheetID:  resp.Updates.SpreadsheetId,
        UpdatedRange:   resp.Updates.UpdatedRange,
        UpdatedRows:    int(resp.Updates.UpdatedRows),
        UpdatedColumns: int(resp.Updates.UpdatedColumns),
        UpdatedCells:   int(resp.Updates.UpdatedCells),
    }, nil
}

// BatchUpdate performs batch updates (formatting, formulas, etc.)
func (m *Manager) BatchUpdate(ctx context.Context, spreadsheetID string, requests []*sheets.Request) (*types.BatchUpdateResponse, error) {
    batchUpdateRequest := &sheets.BatchUpdateSpreadsheetRequest{
        Requests: requests,
    }

    var resp *sheets.BatchUpdateSpreadsheetResponse
    err := retry.WithRetry(ctx, m.retryConfig, func() error {
        var err error
        resp, err = m.service.Spreadsheets.
            BatchUpdate(spreadsheetID, batchUpdateRequest).
            Context(ctx).
            Do()
        return err
    })

    if err != nil {
        return nil, errors.ParseGoogleAPIError(err, "Sheets")
    }

    return &types.BatchUpdateResponse{
        SpreadsheetID: resp.SpreadsheetId,
        RepliesCount:  len(resp.Replies),
        UpdatedCells:  resp.UpdatedCells,
    }, nil
}

// GetSpreadsheet gets metadata about a spreadsheet
func (m *Manager) GetSpreadsheet(ctx context.Context, spreadsheetID string) (*types.Spreadsheet, error) {
    var resp *sheets.Spreadsheet
    err := retry.WithRetry(ctx, m.retryConfig, func() error {
        var err error
        resp, err = m.service.Spreadsheets.Get(spreadsheetID).Context(ctx).Do()
        return err
    })

    if err != nil {
        return nil, errors.ParseGoogleAPIError(err, "Sheets")
    }

    return types.NewSpreadsheetFromAPI(resp), nil
}
```

### 2. Sheets Types

**Files to create**:
- `internal/types/sheets.go`: Request/response types with TableRenderer implementation

**Implementation**:
```go
// internal/types/sheets.go
package types

import (
    "fmt"
    "io"
    "strings"
    "google.golang.org/api/sheets/v4"
)

type SheetValues struct {
    Range          string
    MajorDimension string
    Values         [][]interface{}
}

// Implement TableRenderer
func (v *SheetValues) Headers() []string {
    if len(v.Values) == 0 {
        return []string{"(empty)"}
    }
    // Use first row as headers or generate column letters
    if len(v.Values) > 0 {
        headers := make([]string, len(v.Values[0]))
        for i := range headers {
            headers[i] = columnLetter(i)
        }
        return headers
    }
    return []string{}
}

func (v *SheetValues) Rows() [][]string {
    rows := make([][]string, len(v.Values))
    for i, row := range v.Values {
        rows[i] = make([]string, len(row))
        for j, cell := range row {
            rows[i][j] = fmt.Sprintf("%v", cell)
        }
    }
    return rows
}

func (v *SheetValues) RenderTable(w io.Writer) error {
    // Use tablewriter or simple formatting
    return nil
}

type UpdateValuesRequest struct {
    SpreadsheetID    string
    Range            string
    Values           [][]interface{}
    MajorDimension   string // ROWS or COLUMNS
    ValueInputOption string // RAW or USER_ENTERED
}

type UpdateValuesResponse struct {
    SpreadsheetID  string
    UpdatedRange   string
    UpdatedRows    int
    UpdatedColumns int
    UpdatedCells   int
}

// Implement TableRenderer
func (r *UpdateValuesResponse) Headers() []string {
    return []string{"Spreadsheet ID", "Range", "Rows", "Columns", "Cells"}
}

func (r *UpdateValuesResponse) Rows() [][]string {
    return [][]string{{
        r.SpreadsheetID,
        r.UpdatedRange,
        fmt.Sprintf("%d", r.UpdatedRows),
        fmt.Sprintf("%d", r.UpdatedColumns),
        fmt.Sprintf("%d", r.UpdatedCells),
    }}
}

type AppendValuesRequest struct {
    SpreadsheetID    string
    Range            string
    Values           [][]interface{}
    MajorDimension   string
    ValueInputOption string
}

type BatchUpdateResponse struct {
    SpreadsheetID string
    RepliesCount  int
    UpdatedCells  int64
}

type Spreadsheet struct {
    ID         string
    Title      string
    Locale     string
    TimeZone   string
    SheetCount int
    Sheets     []Sheet
}

type Sheet struct {
    ID    int64
    Title string
    Index int64
    Type  string
}

func NewSpreadsheetFromAPI(s *sheets.Spreadsheet) *Spreadsheet {
    spreadsheet := &Spreadsheet{
        ID:         s.SpreadsheetId,
        Title:      s.Properties.Title,
        Locale:     s.Properties.Locale,
        TimeZone:   s.Properties.TimeZone,
        SheetCount: len(s.Sheets),
        Sheets:     make([]Sheet, len(s.Sheets)),
    }

    for i, sheet := range s.Sheets {
        spreadsheet.Sheets[i] = Sheet{
            ID:    sheet.Properties.SheetId,
            Title: sheet.Properties.Title,
            Index: sheet.Properties.Index,
            Type:  sheet.Properties.SheetType,
        }
    }

    return spreadsheet
}

func columnLetter(col int) string {
    letter := ""
    for col >= 0 {
        letter = string(rune('A'+col%26)) + letter
        col = col/26 - 1
    }
    return letter
}
```

### 3. CLI Commands

**Files to create**:
- `internal/cli/sheets.go`: Sheets command implementation

**Command structure**:
```go
// internal/cli/sheets.go
package cli

import (
    "context"
    "encoding/json"
    "fmt"
    "io"
    "os"
    "strings"

    "github.com/spf13/cobra"
    "google.golang.org/api/sheets/v4"

    "github.com/dl-alexandre/Google-Drive-CLI/internal/auth"
    "github.com/dl-alexandre/Google-Drive-CLI/internal/sheets"
    "github.com/dl-alexandre/Google-Drive-CLI/internal/types"
    "github.com/dl-alexandre/Google-Drive-CLI/internal/utils"
)

var sheetsCmd = &cobra.Command{
    Use:   "sheets",
    Short: "Google Sheets operations",
    Long:  "Read, write, and manage Google Sheets",
}

// sheets list - list spreadsheet files via Drive API
var sheetsListCmd = &cobra.Command{
    Use:   "list",
    Short: "List spreadsheet files",
    Long:  "List all Google Sheets files accessible to you",
    RunE:  runSheetsList,
}

var (
    sheetsListLimit      int
    sheetsListPageToken  string
    sheetsListPaginate   bool
)

func init() {
    rootCmd.AddCommand(sheetsCmd)

    // list
    sheetsCmd.AddCommand(sheetsListCmd)
    sheetsListCmd.Flags().IntVar(&sheetsListLimit, "limit", 100, "Max results per page")
    sheetsListCmd.Flags().StringVar(&sheetsListPageToken, "page-token", "", "Page token")
    sheetsListCmd.Flags().BoolVar(&sheetsListPaginate, "paginate", false, "Auto-paginate all results")

    // get
    sheetsCmd.AddCommand(sheetsGetCmd)
    sheetsGetCmd.Flags().StringVar(&sheetsGetRange, "range", "Sheet1", "A1 notation range")

    // update
    sheetsCmd.AddCommand(sheetsUpdateCmd)
    sheetsUpdateCmd.Flags().StringVar(&sheetsUpdateRange, "range", "", "A1 notation range (required)")
    sheetsUpdateCmd.Flags().StringVar(&sheetsUpdateValues, "values", "", "JSON array of values")
    sheetsUpdateCmd.Flags().StringVar(&sheetsUpdateInput, "input", "", "File path (- for stdin)")
    sheetsUpdateCmd.Flags().StringVar(&sheetsUpdateDimension, "major-dimension", "ROWS", "ROWS or COLUMNS")
    sheetsUpdateCmd.Flags().StringVar(&sheetsUpdateInputOption, "value-input", "USER_ENTERED", "RAW or USER_ENTERED")
    sheetsUpdateCmd.MarkFlagRequired("range")

    // append
    sheetsCmd.AddCommand(sheetsAppendCmd)
    sheetsAppendCmd.Flags().StringVar(&sheetsAppendRange, "range", "", "A1 notation range (required)")
    sheetsAppendCmd.Flags().StringVar(&sheetsAppendValues, "values", "", "JSON array of values")
    sheetsAppendCmd.Flags().StringVar(&sheetsAppendInput, "input", "", "File path (- for stdin)")
    sheetsAppendCmd.Flags().StringVar(&sheetsAppendDimension, "major-dimension", "ROWS", "ROWS or COLUMNS")
    sheetsAppendCmd.Flags().StringVar(&sheetsAppendInputOption, "value-input", "USER_ENTERED", "RAW or USER_ENTERED")
    sheetsAppendCmd.MarkFlagRequired("range")

    // batch-update
    sheetsCmd.AddCommand(sheetsBatchUpdateCmd)
    sheetsBatchUpdateCmd.Flags().StringVar(&sheetsBatchUpdateInput, "input", "", "JSON file with requests (- for stdin)")
    sheetsBatchUpdateCmd.MarkFlagRequired("input")

    // metadata
    sheetsCmd.AddCommand(sheetsMetadataCmd)
}

var sheetsGetCmd = &cobra.Command{
    Use:   "get <spreadsheet-id>",
    Short: "Get values from a range",
    Args:  cobra.ExactArgs(1),
    RunE:  runSheetsGet,
}

var (
    sheetsGetRange string
)

func runSheetsGet(cmd *cobra.Command, args []string) error {
    ctx := context.Background()
    spreadsheetID := args[0]

    mgr, err := getAuthManager()
    if err != nil {
        return err
    }

    creds, err := mgr.LoadCredentials(getProfile())
    if err != nil {
        return fmt.Errorf("load credentials: %w", err)
    }

    svc, err := createSheetsService(ctx, mgr, creds)
    if err != nil {
        return err
    }

    sheetsMgr := sheets.NewManager(svc)
    values, err := sheetsMgr.GetValues(ctx, spreadsheetID, sheetsGetRange)
    if err != nil {
        return err
    }

    return outputWriter.Write(values)
}

var sheetsUpdateCmd = &cobra.Command{
    Use:   "update <spreadsheet-id>",
    Short: "Update values in a range",
    Args:  cobra.ExactArgs(1),
    RunE:  runSheetsUpdate,
}

var (
    sheetsUpdateRange       string
    sheetsUpdateValues      string
    sheetsUpdateInput       string
    sheetsUpdateDimension   string
    sheetsUpdateInputOption string
)

func runSheetsUpdate(cmd *cobra.Command, args []string) error {
    ctx := context.Background()
    spreadsheetID := args[0]

    // Parse values from --values or --input
    values, err := parseValuesInput(sheetsUpdateValues, sheetsUpdateInput)
    if err != nil {
        return fmt.Errorf("parse values: %w", err)
    }

    mgr, err := getAuthManager()
    if err != nil {
        return err
    }

    creds, err := mgr.LoadCredentials(getProfile())
    if err != nil {
        return fmt.Errorf("load credentials: %w", err)
    }

    svc, err := createSheetsService(ctx, mgr, creds)
    if err != nil {
        return err
    }

    sheetsMgr := sheets.NewManager(svc)
    req := &types.UpdateValuesRequest{
        SpreadsheetID:    spreadsheetID,
        Range:            sheetsUpdateRange,
        Values:           values,
        MajorDimension:   sheetsUpdateDimension,
        ValueInputOption: sheetsUpdateInputOption,
    }

    resp, err := sheetsMgr.UpdateValues(ctx, req)
    if err != nil {
        return err
    }

    return outputWriter.Write(resp)
}

var sheetsAppendCmd = &cobra.Command{
    Use:   "append <spreadsheet-id>",
    Short: "Append values to a range",
    Args:  cobra.ExactArgs(1),
    RunE:  runSheetsAppend,
}

var (
    sheetsAppendRange       string
    sheetsAppendValues      string
    sheetsAppendInput       string
    sheetsAppendDimension   string
    sheetsAppendInputOption string
)

func runSheetsAppend(cmd *cobra.Command, args []string) error {
    // Similar to update but calls AppendValues
    ctx := context.Background()
    spreadsheetID := args[0]

    values, err := parseValuesInput(sheetsAppendValues, sheetsAppendInput)
    if err != nil {
        return fmt.Errorf("parse values: %w", err)
    }

    mgr, err := getAuthManager()
    if err != nil {
        return err
    }

    creds, err := mgr.LoadCredentials(getProfile())
    if err != nil {
        return fmt.Errorf("load credentials: %w", err)
    }

    svc, err := createSheetsService(ctx, mgr, creds)
    if err != nil {
        return err
    }

    sheetsMgr := sheets.NewManager(svc)
    req := &types.AppendValuesRequest{
        SpreadsheetID:    spreadsheetID,
        Range:            sheetsAppendRange,
        Values:           values,
        MajorDimension:   sheetsAppendDimension,
        ValueInputOption: sheetsAppendInputOption,
    }

    resp, err := sheetsMgr.AppendValues(ctx, req)
    if err != nil {
        return err
    }

    return outputWriter.Write(resp)
}

var sheetsBatchUpdateCmd = &cobra.Command{
    Use:   "batch-update <spreadsheet-id>",
    Short: "Batch update spreadsheet (formatting, formulas, etc.)",
    Args:  cobra.ExactArgs(1),
    RunE:  runSheetsBatchUpdate,
}

var (
    sheetsBatchUpdateInput string
)

func runSheetsBatchUpdate(cmd *cobra.Command, args []string) error {
    ctx := context.Background()
    spreadsheetID := args[0]

    // Read batch update requests from file/stdin
    var reader io.Reader
    if sheetsBatchUpdateInput == "-" {
        reader = os.Stdin
    } else {
        f, err := os.Open(sheetsBatchUpdateInput)
        if err != nil {
            return fmt.Errorf("open input file: %w", err)
        }
        defer f.Close()
        reader = f
    }

    var requests []*sheets.Request
    if err := json.NewDecoder(reader).Decode(&requests); err != nil {
        return fmt.Errorf("parse batch requests: %w", err)
    }

    mgr, err := getAuthManager()
    if err != nil {
        return err
    }

    creds, err := mgr.LoadCredentials(getProfile())
    if err != nil {
        return fmt.Errorf("load credentials: %w", err)
    }

    svc, err := createSheetsService(ctx, mgr, creds)
    if err != nil {
        return err
    }

    sheetsMgr := sheets.NewManager(svc)
    resp, err := sheetsMgr.BatchUpdate(ctx, spreadsheetID, requests)
    if err != nil {
        return err
    }

    return outputWriter.Write(resp)
}

var sheetsMetadataCmd = &cobra.Command{
    Use:   "metadata <spreadsheet-id>",
    Short: "Get spreadsheet metadata",
    Args:  cobra.ExactArgs(1),
    RunE:  runSheetsMetadata,
}

func runSheetsMetadata(cmd *cobra.Command, args []string) error {
    ctx := context.Background()
    spreadsheetID := args[0]

    mgr, err := getAuthManager()
    if err != nil {
        return err
    }

    creds, err := mgr.LoadCredentials(getProfile())
    if err != nil {
        return fmt.Errorf("load credentials: %w", err)
    }

    svc, err := createSheetsService(ctx, mgr, creds)
    if err != nil {
        return err
    }

    sheetsMgr := sheets.NewManager(svc)
    metadata, err := sheetsMgr.GetSpreadsheet(ctx, spreadsheetID)
    if err != nil {
        return err
    }

    return outputWriter.Write(metadata)
}

// Helper functions

func createSheetsService(ctx context.Context, mgr *auth.Manager, creds *types.Credentials) (*sheets.Service, error) {
    client := mgr.GetHTTPClient(ctx, creds)
    return sheets.NewService(ctx, option.WithHTTPClient(client))
}

func parseValuesInput(valuesFlag, inputFlag string) ([][]interface{}, error) {
    var data [][]interface{}

    if valuesFlag != "" {
        if err := json.Unmarshal([]byte(valuesFlag), &data); err != nil {
            return nil, fmt.Errorf("parse --values JSON: %w", err)
        }
        return data, nil
    }

    if inputFlag != "" {
        var reader io.Reader
        if inputFlag == "-" {
            reader = os.Stdin
        } else {
            f, err := os.Open(inputFlag)
            if err != nil {
                return nil, fmt.Errorf("open input file: %w", err)
            }
            defer f.Close()
            reader = f
        }

        if err := json.NewDecoder(reader).Decode(&data); err != nil {
            return nil, fmt.Errorf("parse input JSON: %w", err)
        }
        return data, nil
    }

    return nil, fmt.Errorf("either --values or --input must be provided")
}

func runSheetsList(cmd *cobra.Command, args []string) error {
    // Use Drive API to list spreadsheets
    ctx := context.Background()

    mgr, err := getAuthManager()
    if err != nil {
        return err
    }

    creds, err := mgr.LoadCredentials(getProfile())
    if err != nil {
        return fmt.Errorf("load credentials: %w", err)
    }

    driveSvc, err := mgr.GetDriveService(ctx, creds)
    if err != nil {
        return fmt.Errorf("create drive service: %w", err)
    }

    query := fmt.Sprintf("mimeType='%s'", utils.MimeTypeSpreadsheet)

    // Use existing Drive list logic
    call := driveSvc.Files.List().
        Q(query).
        PageSize(int64(sheetsListLimit)).
        Fields("nextPageToken, files(id, name, mimeType, createdTime, modifiedTime, size)")

    if sheetsListPageToken != "" {
        call = call.PageToken(sheetsListPageToken)
    }

    resp, err := call.Context(ctx).Do()
    if err != nil {
        return fmt.Errorf("list spreadsheets: %w", err)
    }

    // Convert to DriveFile types for consistent output
    files := make([]*types.DriveFile, len(resp.Files))
    for i, f := range resp.Files {
        files[i] = types.NewDriveFileFromAPI(f)
    }

    return outputWriter.Write(files)
}
```

### 4. Update Output Writer

**Files to modify**:
- `internal/cli/output.go`: Add Sheets types to table rendering

**Changes**:
```go
func (w *OutputWriter) writeTable(data interface{}) error {
    // Try TableRenderer interface first
    if renderer, ok := data.(types.TableRenderer); ok {
        return w.renderTableFromInterface(renderer)
    }

    // Type switch for legacy types
    switch v := data.(type) {
    case []*types.DriveFile:
        return w.writeFileTable(v)
    case *types.SheetValues:
        return w.renderTableFromInterface(v)
    case *types.UpdateValuesResponse:
        return w.renderTableFromInterface(v)
    case *types.Spreadsheet:
        return w.writeSpreadsheetTable(v)
    // ... other cases
    default:
        return w.writeJSON(data)
    }
}
```

### 5. Documentation & Dependencies

**Files to modify**:
- `README.md`: Add Sheets section with examples
- `go.mod`: Add `google.golang.org/api/sheets/v4`

**README additions**:
````markdown
### Sheets Operations

```bash
# List all spreadsheets
gdrive sheets list --json

# Get values from a range
gdrive sheets get 1abc123... --range "Sheet1!A1:C10" --json

# Update values
gdrive sheets update 1abc123... --range "Sheet1!A1" --values '[[\"Name\",\"Age\"],[\"Alice\",30]]'

# Append values
gdrive sheets append 1abc123... --range "Sheet1!A:B" --values '[[\"Bob\",25]]'

# Batch update (formatting, formulas)
echo '[{"updateCells": {...}}]' | gdrive sheets batch-update 1abc123... --input -

# Get metadata
gdrive sheets metadata 1abc123... --json
```
````

## Todo

- [x] Create `internal/sheets/manager.go` with GetValues, UpdateValues, AppendValues, BatchUpdate, GetSpreadsheet
- [x] Write unit tests in `internal/sheets/manager_test.go` using mocks
- [x] Create `internal/types/sheets.go` with SheetValues, UpdateValuesRequest, etc.
- [x] Implement TableRenderer for SheetValues and UpdateValuesResponse
- [x] Create `internal/cli/sheets.go` with list, get, update, append, batch-update, metadata commands
- [x] Add parseValuesInput helper for JSON parsing
- [x] Update `internal/cli/output.go` to handle Sheets types
- [x] Add `google.golang.org/api/sheets/v4` to `go.mod`
- [x] Add Sheets examples to README.md
- [x] Write integration tests with test spreadsheet
- [x] Test error scenarios (invalid range, quota exceeded, auth failures)
- [x] Add example JSON files for batch-update in examples/

## Testing Strategy

1. **Unit tests**: Mock sheets.Service responses
2. **Integration tests**: Create test spreadsheet, read/write/update
3. **Error handling**: Invalid ranges, quota errors, network failures
4. **Input validation**: Malformed JSON, invalid A1 notation
5. **Output formats**: Verify JSON and table outputs

## Examples to Include

```bash
# Read a range
gdrive sheets get 1abc... --range "Sheet1!A1:B10"

# Update with JSON
gdrive sheets update 1abc... --range "A1" --values '[[1,2,3]]'

# Append from file
echo '[[\"New\",\"Row\"]]' | gdrive sheets append 1abc... --range "A:B" --input -

# Format cells (batch update)
cat format.json | gdrive sheets batch-update 1abc... --input -
```

## Dependencies

Requires completion of:
- Foundation for Google APIs Integration (00_foundation_google_apis.plan.md)

Blocks:
- None (Docs and Slides can be developed in parallel)
