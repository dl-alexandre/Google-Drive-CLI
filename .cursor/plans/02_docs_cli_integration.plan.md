---
name: Docs CLI Integration
overview: Add Google Docs API support to read document structure, extract text content, and modify documents using batchUpdate with consistent CLI patterns.
todos: []
isProject: false
status: completed
dependencies:
  - 00_foundation_google_apis.plan.md
---

> **Status**: ✅ **COMPLETED** - Docs CLI integration has been successfully implemented.

## Context

Foundation provides: scope management, service factory, error handling, retry logic, TableRenderer interface.

Use cases:
- Extract text from Docs for AI/LLM processing
- Template-based document generation
- Automated report creation
- Content migration

Existing patterns from Sheets integration can be adapted for Docs structure.

## Plan

### 1. Docs Manager Layer

**Files to create**:
- `internal/docs/manager.go`: Core Docs operations
- `internal/docs/manager_test.go`: Unit tests
- `internal/types/docs.go`: Docs-specific types

**Key operations**:
```go
type Manager struct {
    service *docs.Service
    retryConfig *retry.RetryConfig
}

// GetDocument fetches full document structure
func (m *Manager) GetDocument(ctx context.Context, documentID string) (*types.Document, error)

// ExtractText extracts plain text content
func (m *Manager) ExtractText(ctx context.Context, documentID string) (*types.DocumentText, error)

// BatchUpdate performs batch updates (insert text, format, etc.)
func (m *Manager) BatchUpdate(ctx context.Context, documentID string, requests []*docs.Request) (*types.DocsBatchUpdateResponse, error)

// CreateDocument creates a new document
func (m *Manager) CreateDocument(ctx context.Context, title string) (*types.Document, error)
```

### 2. Docs Types

**Files to create**:
- `internal/types/docs.go`: Document, DocumentText, DocsBatchUpdateRequest, etc.

**Key types**:
```go
type Document struct {
    ID           string
    Title        string
    RevisionID   string
    SuggestionsViewMode string
    DocumentStyle DocumentStyle
    Body         DocumentBody
}

type DocumentText struct {
    DocumentID string
    Title      string
    Text       string      // Plain text extracted
    WordCount  int
    CharCount  int
}

// Implement TableRenderer for text output
func (t *DocumentText) Headers() []string {
    return []string{"Document ID", "Title", "Words", "Characters"}
}

func (t *DocumentText) Rows() [][]string {
    return [][]string{{
        t.DocumentID,
        t.Title,
        fmt.Sprintf("%d", t.WordCount),
        fmt.Sprintf("%d", t.CharCount),
    }}
}

type DocsBatchUpdateResponse struct {
    DocumentID   string
    RevisionID   string
    RepliesCount int
}

// Helper to extract text from document structure
func extractTextFromBody(body *docs.Body) string {
    var text strings.Builder
    for _, element := range body.Content {
        if element.Paragraph != nil {
            for _, elem := range element.Paragraph.Elements {
                if elem.TextRun != nil {
                    text.WriteString(elem.TextRun.Content)
                }
            }
        }
    }
    return text.String()
}
```

### 3. CLI Commands

**Files to create**:
- `internal/cli/docs.go`: Docs command implementation

**Commands**:
```go
var docsCmd = &cobra.Command{
    Use:   "docs",
    Short: "Google Docs operations",
    Long:  "Read and modify Google Docs",
}

// docs list - list document files via Drive API
// docs get <doc-id> - get document structure
// docs read <doc-id> - extract plain text
// docs update <doc-id> - batch update from JSON
// docs create <title> - create new document
```

**Example implementations**:
```go
// Extract text command
var docsReadCmd = &cobra.Command{
    Use:   "read <document-id>",
    Short: "Extract plain text from document",
    Args:  cobra.ExactArgs(1),
    RunE:  runDocsRead,
}

func runDocsRead(cmd *cobra.Command, args []string) error {
    ctx := context.Background()
    documentID := args[0]

    mgr, creds, err := getAuthAndCreds()
    if err != nil {
        return err
    }

    svc, err := createDocsService(ctx, mgr, creds)
    if err != nil {
        return err
    }

    docsMgr := docs.NewManager(svc)
    text, err := docsMgr.ExtractText(ctx, documentID)
    if err != nil {
        return err
    }

    // If --json, output structured data
    // Otherwise, just print text
    if outputFormat == "json" {
        return outputWriter.Write(text)
    }

    fmt.Println(text.Text)
    return nil
}

// Batch update command
var docsUpdateCmd = &cobra.Command{
    Use:   "update <document-id>",
    Short: "Batch update document",
    Args:  cobra.ExactArgs(1),
    RunE:  runDocsUpdate,
}

var (
    docsUpdateInput string
)

func init() {
    docsUpdateCmd.Flags().StringVar(&docsUpdateInput, "input", "", "JSON file with requests (- for stdin)")
    docsUpdateCmd.MarkFlagRequired("input")
}

func runDocsUpdate(cmd *cobra.Command, args []string) error {
    ctx := context.Background()
    documentID := args[0]

    var reader io.Reader
    if docsUpdateInput == "-" {
        reader = os.Stdin
    } else {
        f, err := os.Open(docsUpdateInput)
        if err != nil {
            return fmt.Errorf("open input: %w", err)
        }
        defer f.Close()
        reader = f
    }

    var requests []*docs.Request
    if err := json.NewDecoder(reader).Decode(&requests); err != nil {
        return fmt.Errorf("parse requests: %w", err)
    }

    mgr, creds, err := getAuthAndCreds()
    if err != nil {
        return err
    }

    svc, err := createDocsService(ctx, mgr, creds)
    if err != nil {
        return err
    }

    docsMgr := docs.NewManager(svc)
    resp, err := docsMgr.BatchUpdate(ctx, documentID, requests)
    if err != nil {
        return err
    }

    return outputWriter.Write(resp)
}

// Create document command
var docsCreateCmd = &cobra.Command{
    Use:   "create <title>",
    Short: "Create a new document",
    Args:  cobra.ExactArgs(1),
    RunE:  runDocsCreate,
}

func runDocsCreate(cmd *cobra.Command, args []string) error {
    ctx := context.Background()
    title := args[0]

    mgr, creds, err := getAuthAndCreds()
    if err != nil {
        return err
    }

    svc, err := createDocsService(ctx, mgr, creds)
    if err != nil {
        return err
    }

    docsMgr := docs.NewManager(svc)
    doc, err := docsMgr.CreateDocument(ctx, title)
    if err != nil {
        return err
    }

    return outputWriter.Write(doc)
}
```

### 4. Update Output Writer

**Files to modify**:
- `internal/cli/output.go`: Add Docs types

```go
case *types.Document:
    return w.writeDocumentTable(v)
case *types.DocumentText:
    return w.renderTableFromInterface(v)
case *types.DocsBatchUpdateResponse:
    return w.renderTableFromInterface(v)
```

### 5. Documentation

**Files to modify**:
- `README.md`: Add Docs section
- `go.mod`: Add `google.golang.org/api/docs/v1`

**README additions**:
````markdown
### Docs Operations

```bash
# List all documents
gdrive docs list --json

# Get document structure
gdrive docs get 1abc123... --json

# Extract plain text
gdrive docs read 1abc123...
gdrive docs read 1abc123... --json  # Structured output

# Create document
gdrive docs create "My Report" --json

# Update document (insert text, format)
cat updates.json | gdrive docs update 1abc123... --input -
```

**Example updates.json**:
```json
[
  {
    "insertText": {
      "location": {"index": 1},
      "text": "Hello World\n"
    }
  },
  {
    "updateTextStyle": {
      "range": {"startIndex": 1, "endIndex": 11},
      "textStyle": {"bold": true},
      "fields": "bold"
    }
  }
]
```
````

## Todo

- [x] Create `internal/docs/manager.go` with GetDocument, ExtractText, BatchUpdate, CreateDocument
- [x] Write unit tests in `internal/docs/manager_test.go`
- [x] Create `internal/types/docs.go` with Document, DocumentText, DocsBatchUpdateResponse
- [x] Implement TableRenderer for DocumentText
- [x] Create `internal/cli/docs.go` with list, get, read, update, create commands
- [x] Add text extraction helper (traverse document structure)
- [x] Update `internal/cli/output.go` for Docs types
- [x] Add `google.golang.org/api/docs/v1` to `go.mod`
- [x] Add Docs examples to README.md
- [x] Create example JSON files for batch updates
- [x] Write integration tests with test document
- [x] Test text extraction edge cases (tables, lists, headers)

## Testing Strategy

1. **Unit tests**: Mock docs.Service responses
2. **Integration tests**: Create test doc, read, update, extract text
3. **Text extraction**: Test with various document structures (tables, lists, headers, footers)
4. **Batch updates**: Test insertText, deleteText, formatting
5. **Error handling**: Invalid document ID, permission errors

## Use Case Examples

**AI/LLM Text Extraction**:
```bash
# Extract text for processing
gdrive docs read 1abc123... > content.txt
cat content.txt | llm "Summarize this"
```

**Template Generation**:
```bash
# Create doc from template
TEMPLATE_ID=1xyz789...
gdrive docs get $TEMPLATE_ID --json | \
  jq '.body' | \
  # Modify structure
  gdrive docs create "New Report" --template -
```

**Automated Reporting**:
```bash
# Insert data into doc
cat <<EOF | gdrive docs update 1abc123... --input -
[
  {"insertText": {"location": {"index": 1}, "text": "Report Date: $(date)\n"}},
  {"insertText": {"location": {"index": 20}, "text": "Total Sales: \$10,000\n"}}
]
EOF
```

## Dependencies

Requires:
- Foundation for Google APIs Integration (00_foundation_google_apis.plan.md)

Blocks:
- None (parallel with Sheets and Slides)
