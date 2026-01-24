---
name: Slides CLI Integration
overview: Add Google Slides API support for reading presentation structure, extracting content, and performing template-based updates for automated slide generation.
todos: []
isProject: false
status: completed
dependencies:
  - 00_foundation_google_apis.plan.md
---

> **Status**: ✅ **COMPLETED** - Slides CLI integration has been successfully implemented.

## Context

Foundation provides: scope management, service factory, error handling, retry logic, TableRenderer interface.

Use cases:
- Template-based slide generation (merge data into placeholders)
- Automated report presentations
- Bulk slide updates
- Content extraction for analysis

Similar patterns to Sheets/Docs but with presentation-specific operations (slides, shapes, placeholders).

## Plan

### 1. Slides Manager Layer

**Files to create**:
- `internal/slides/manager.go`: Core Slides operations
- `internal/slides/manager_test.go`: Unit tests
- `internal/types/slides.go`: Slides-specific types

**Key operations**:
```go
type Manager struct {
    service *slides.Service
    retryConfig *retry.RetryConfig
}

// GetPresentation fetches full presentation structure
func (m *Manager) GetPresentation(ctx context.Context, presentationID string) (*types.Presentation, error)

// ListSlides lists all slides in presentation
func (m *Manager) ListSlides(ctx context.Context, presentationID string) ([]*types.Slide, error)

// ExtractText extracts text from all slides
func (m *Manager) ExtractText(ctx context.Context, presentationID string) (*types.PresentationText, error)

// BatchUpdate performs batch updates (replace placeholders, add slides, etc.)
func (m *Manager) BatchUpdate(ctx context.Context, presentationID string, requests []*slides.Request) (*types.SlidesBatchUpdateResponse, error)

// CreatePresentation creates new presentation
func (m *Manager) CreatePresentation(ctx context.Context, title string) (*types.Presentation, error)

// ReplaceAllText replaces text globally (templating)
func (m *Manager) ReplaceAllText(ctx context.Context, presentationID string, replacements map[string]string) (*types.SlidesBatchUpdateResponse, error)
```

### 2. Slides Types

**Files to create**:
- `internal/types/slides.go`: Presentation, Slide, PresentationText, etc.

**Key types**:
```go
type Presentation struct {
    ID         string
    Title      string
    PageSize   PageSize
    SlideCount int
    Slides     []Slide
    Masters    []Master
    Layouts    []Layout
}

type Slide struct {
    ObjectID          string
    Index             int
    SlideProperties   SlideProperties
    PageElements      []PageElement
    NotesPage         NotesPage
}

type PageElement struct {
    ObjectID string
    Type     string // SHAPE, IMAGE, TABLE, etc.
    Title    string
    Shape    *Shape
}

type Shape struct {
    ShapeType string
    Text      string
    Placeholder *Placeholder
}

type Placeholder struct {
    Type  string // TITLE, BODY, CENTERED_TITLE, etc.
    Index int
}

type PresentationText struct {
    PresentationID string
    Title          string
    SlideCount     int
    TextBySlide    []SlideText
}

type SlideText struct {
    SlideIndex int
    ObjectID   string
    Text       string
}

// Implement TableRenderer
func (p *PresentationText) Headers() []string {
    return []string{"Slide", "Object ID", "Text"}
}

func (p *PresentationText) Rows() [][]string {
    rows := make([][]string, len(p.TextBySlide))
    for i, st := range p.TextBySlide {
        rows[i] = []string{
            fmt.Sprintf("%d", st.SlideIndex),
            st.ObjectID,
            truncate(st.Text, 50),
        }
    }
    return rows
}

type SlidesBatchUpdateResponse struct {
    PresentationID string
    RepliesCount   int
    WriteControl   WriteControl
}

type PageSize struct {
    Width  int64
    Height int64
    Unit   string
}

// Helper to extract text from presentation
func extractTextFromPresentation(pres *slides.Presentation) *PresentationText {
    result := &PresentationText{
        PresentationID: pres.PresentationId,
        Title:          pres.Title,
        SlideCount:     len(pres.Slides),
        TextBySlide:    []SlideText{},
    }

    for i, slide := range pres.Slides {
        for _, element := range slide.PageElements {
            if element.Shape != nil && element.Shape.Text != nil {
                text := extractTextFromShape(element.Shape)
                if text != "" {
                    result.TextBySlide = append(result.TextBySlide, SlideText{
                        SlideIndex: i,
                        ObjectID:   element.ObjectId,
                        Text:       text,
                    })
                }
            }
        }
    }

    return result
}
```

### 3. CLI Commands

**Files to create**:
- `internal/cli/slides.go`: Slides command implementation

**Commands**:
```go
var slidesCmd = &cobra.Command{
    Use:   "slides",
    Short: "Google Slides operations",
    Long:  "Read and modify Google Slides presentations",
}

// slides list - list presentation files via Drive API
// slides get <pres-id> - get presentation structure
// slides read <pres-id> - extract text from all slides
// slides update <pres-id> - batch update from JSON
// slides create <title> - create new presentation
// slides replace <pres-id> - replace placeholders (templating)
```

**Template replacement command** (most useful):
```go
var slidesReplaceCmd = &cobra.Command{
    Use:   "replace <presentation-id>",
    Short: "Replace text placeholders (templating)",
    Long:  "Replace {{placeholders}} with values for template-based generation",
    Args:  cobra.ExactArgs(1),
    RunE:  runSlidesReplace,
}

var (
    slidesReplaceFile string
    slidesReplaceData string
)

func init() {
    slidesReplaceCmd.Flags().StringVar(&slidesReplaceFile, "file", "", "JSON file with replacements")
    slidesReplaceCmd.Flags().StringVar(&slidesReplaceData, "data", "", "JSON string with replacements")
}

func runSlidesReplace(cmd *cobra.Command, args []string) error {
    ctx := context.Background()
    presentationID := args[0]

    // Parse replacements: {"{{NAME}}": "Alice", "{{DATE}}": "2026-01-24"}
    var replacements map[string]string
    if slidesReplaceFile != "" {
        data, err := os.ReadFile(slidesReplaceFile)
        if err != nil {
            return fmt.Errorf("read file: %w", err)
        }
        if err := json.Unmarshal(data, &replacements); err != nil {
            return fmt.Errorf("parse JSON: %w", err)
        }
    } else if slidesReplaceData != "" {
        if err := json.Unmarshal([]byte(slidesReplaceData), &replacements); err != nil {
            return fmt.Errorf("parse JSON: %w", err)
        }
    } else {
        return fmt.Errorf("either --file or --data required")
    }

    mgr, creds, err := getAuthAndCreds()
    if err != nil {
        return err
    }

    svc, err := createSlidesService(ctx, mgr, creds)
    if err != nil {
        return err
    }

    slidesMgr := slides.NewManager(svc)
    resp, err := slidesMgr.ReplaceAllText(ctx, presentationID, replacements)
    if err != nil {
        return err
    }

    return outputWriter.Write(resp)
}
```

### 4. Templating Implementation

**Key feature: ReplaceAllText in manager**:
```go
func (m *Manager) ReplaceAllText(ctx context.Context, presentationID string, replacements map[string]string) (*types.SlidesBatchUpdateResponse, error) {
    requests := make([]*slides.Request, 0, len(replacements))

    for find, replace := range replacements {
        requests = append(requests, &slides.Request{
            ReplaceAllText: &slides.ReplaceAllTextRequest{
                ContainsText: &slides.SubstringMatchCriteria{
                    Text:      find,
                    MatchCase: true,
                },
                ReplaceText: replace,
            },
        })
    }

    return m.BatchUpdate(ctx, presentationID, requests)
}
```

### 5. Documentation

**Files to modify**:
- `README.md`: Add Slides section
- `go.mod`: Add `google.golang.org/api/slides/v1`

**README additions**:
````markdown
### Slides Operations

```bash
# List all presentations
gdrive slides list --json

# Get presentation structure
gdrive slides get 1abc123... --json

# Extract text from all slides
gdrive slides read 1abc123...
gdrive slides read 1abc123... --json

# Create presentation
gdrive slides create "Q1 Report" --json

# Replace placeholders (templating)
gdrive slides replace 1abc123... --data '{"{{NAME}}":"Alice","{{Q1_SALES}}":"$100K"}'

# Batch update
cat updates.json | gdrive slides update 1abc123... --input -
```

**Template Example**:

1. Create template presentation with placeholders: `{{NAME}}`, `{{DATE}}`, `{{TOTAL_SALES}}`
2. Use CLI to generate reports:

```bash
# Generate monthly report
gdrive slides replace 1template123... \
  --data '{"{{NAME}}":"January Report","{{DATE}}":"2026-01","{{TOTAL_SALES}}":"$50K"}' \
  --json
```
````

## Todo

- [x] Create `internal/slides/manager.go` with GetPresentation, ListSlides, ExtractText, BatchUpdate, CreatePresentation, ReplaceAllText
- [x] Write unit tests in `internal/slides/manager_test.go`
- [x] Create `internal/types/slides.go` with Presentation, Slide, PresentationText, SlidesBatchUpdateResponse
- [x] Implement text extraction helper (traverse page elements)
- [x] Create `internal/cli/slides.go` with list, get, read, update, create, replace commands
- [x] Implement ReplaceAllText for templating use case
- [x] Update `internal/cli/output.go` for Slides types
- [x] Add `google.golang.org/api/slides/v1` to `go.mod`
- [x] Add Slides examples to README.md
- [x] Create example template presentation
- [x] Create example replacement JSON files
- [x] Write integration tests with test presentation
- [x] Test templating with various placeholder patterns

## Testing Strategy

1. **Unit tests**: Mock slides.Service responses
2. **Integration tests**: Create test presentation, read, update, replace
3. **Text extraction**: Test with various slide layouts (title, body, tables)
4. **Templating**: Test ReplaceAllText with multiple patterns
5. **Batch updates**: Test createSlide, deleteSlide, updateShapeProperties

## Use Case Examples

**Monthly Report Generation**:
```bash
# Template with {{MONTH}}, {{REVENUE}}, {{EXPENSES}}
TEMPLATE_ID=1template123...

for month in Jan Feb Mar; do
  gdrive slides replace $TEMPLATE_ID \
    --data "{\"{{MONTH}}\":\"$month\",\"{{REVENUE}}\":\"$((RANDOM % 100))K\",\"{{EXPENSES}}\":\"$((RANDOM % 50))K\"}" \
    --json > ${month}_report.json
done
```

**Presentation Analysis**:
```bash
# Extract all text for analysis
gdrive slides read 1abc123... --json | \
  jq -r '.textBySlide[].text' | \
  wc -w  # Word count
```

**Bulk Slide Updates**:
```bash
# Add footer to all slides
cat <<EOF | gdrive slides update 1abc123... --input -
[
  {
    "createShape": {
      "objectId": "footer",
      "shapeType": "TEXT_BOX",
      "elementProperties": {
        "pageObjectId": "slide1",
        "size": {"width": {"magnitude": 100, "unit": "PT"}, "height": {"magnitude": 20, "unit": "PT"}},
        "transform": {"translateX": 50, "translateY": 500, "unit": "PT"}
      }
    }
  },
  {
    "insertText": {
      "objectId": "footer",
      "text": "Confidential - 2026"
    }
  }
]
EOF
```

## Dependencies

Requires:
- Foundation for Google APIs Integration (00_foundation_google_apis.plan.md)

Blocks:
- None (parallel with Sheets and Docs)
