# Google Docs Workflows

This guide covers working with Google Docs using the Google Drive CLI.

## Table of Contents

- [Reading Document Content](#reading-document-content)
- [Creating Documents](#creating-documents)
- [Updating Documents](#updating-documents)
- [Text Extraction](#text-extraction)

## Reading Document Content

### Get Document Structure

Retrieve the full document structure including formatting, styles, and content:

```bash
gdrive docs get <document-id> --json
```

This returns detailed metadata about the document, including:
- Document ID and title
- Revision ID
- Document style properties
- Body structure with all elements

**Example Output:**
```json
{
  "documentId": "1abc123...",
  "title": "My Document",
  "revisionId": "rev123...",
  "body": {
    "content": [
      {
        "paragraph": {
          "elements": [
            {
              "textRun": {
                "content": "Hello, World!\n",
                "textStyle": {}
              }
            }
          ],
          "paragraphStyle": {}
        }
      }
    ]
  }
}
```

### List Documents

Find documents in your Drive:

```bash
# List all documents
gdrive docs list --json

# Filter by parent folder
gdrive docs list --parent <folder-id> --json

# Search with query
gdrive docs list --query "name contains 'Report'" --json

# Paginate through results
gdrive docs list --paginate --json
```

## Creating Documents

### Create a New Document

Create a blank document:

```bash
# Create in root
gdrive docs create "My Document" --json

# Create in specific folder
gdrive docs create "Q1 Report" --parent <folder-id> --json
```

The command returns the document ID which you can use for subsequent operations.

### Create Document with Initial Content

Create a document and immediately add content:

```bash
#!/bin/bash
DOC_ID=$(gdrive docs create "New Document" --json | jq -r '.id')

# Add initial content
cat > initial-content.json <<EOF
[
  {
    "insertText": {
      "location": {
        "index": 1
      },
      "text": "Welcome to My Document\n\nThis is the first paragraph.\n"
    }
  }
]
EOF

gdrive docs update "$DOC_ID" --requests-file initial-content.json --json
```

## Updating Documents

### Batch Update Format

Documents are updated using batch update requests. Each request performs a specific operation.

**Example `batch-update.json`:**
```json
[
  {
    "insertText": {
      "location": {
        "index": 1
      },
      "text": "Hello World\n"
    }
  },
  {
    "updateTextStyle": {
      "range": {
        "startIndex": 1,
        "endIndex": 12
      },
      "textStyle": {
        "bold": true
      },
      "fields": "bold"
    }
  }
]
```

### Execute Batch Update

```bash
# From file
gdrive docs update <document-id> --requests-file batch-update.json --json

# From stdin
cat batch-update.json | gdrive docs update <document-id> --requests-file - --json
```

### Common Update Operations

#### Insert Text

```json
[
  {
    "insertText": {
      "location": {
        "index": 1
      },
      "text": "New text content\n"
    }
  }
]
```

**Note:** The `index` represents the position in the document. Index 1 is the beginning of the document.

#### Delete Text

```json
[
  {
    "deleteContentRange": {
      "range": {
        "startIndex": 1,
        "endIndex": 10
      }
    }
  }
]
```

#### Format Text

**Bold text:**
```json
[
  {
    "updateTextStyle": {
      "range": {
        "startIndex": 1,
        "endIndex": 10
      },
      "textStyle": {
        "bold": true
      },
      "fields": "bold"
    }
  }
]
```

**Italic text:**
```json
[
  {
    "updateTextStyle": {
      "range": {
        "startIndex": 1,
        "endIndex": 10
      },
      "textStyle": {
        "italic": true
      },
      "fields": "italic"
    }
  }
]
```

**Change font size:**
```json
[
  {
    "updateTextStyle": {
      "range": {
        "startIndex": 1,
        "endIndex": 10
      },
      "textStyle": {
        "fontSize": {
          "magnitude": 18,
          "unit": "PT"
        }
      },
      "fields": "fontSize"
    }
  }
]
```

**Change font color:**
```json
[
  {
    "updateTextStyle": {
      "range": {
        "startIndex": 1,
        "endIndex": 10
      },
      "textStyle": {
        "foregroundColor": {
          "color": {
            "rgbColor": {
              "red": 1.0,
              "green": 0.0,
              "blue": 0.0
            }
          }
        }
      },
      "fields": "foregroundColor"
    }
  }
]
```

#### Insert Table

```json
[
  {
    "insertTable": {
      "location": {
        "index": 1
      },
      "rows": 3,
      "columns": 2
    }
  }
]
```

#### Insert Page Break

```json
[
  {
    "insertPageBreak": {
      "location": {
        "index": 100
      }
    }
  }
]
```

#### Insert Horizontal Rule

```json
[
  {
    "insertHorizontalRule": {
      "location": {
        "index": 50
      }
    }
  }
]
```

## Text Extraction

### Read Document Content

Extract plain text from a document:

```bash
# Plain text output
gdrive docs read <document-id>

# Structured JSON output
gdrive docs read <document-id> --json
```

**Example JSON Output:**
```json
{
  "documentId": "1abc123...",
  "title": "My Document",
  "text": "Hello, World!\n\nThis is a paragraph.\n\nAnother paragraph here.",
  "wordCount": 10,
  "characterCount": 65
}
```

### Extract Text for Processing

Use text extraction for AI/LLM processing or analysis:

```bash
#!/bin/bash
DOC_ID="your-document-id"

# Extract text
TEXT=$(gdrive docs read "$DOC_ID" --json | jq -r '.text')

# Process with external tool
echo "$TEXT" | llm "Summarize this document"

# Save to file
gdrive docs read "$DOC_ID" --json | jq -r '.text' > document.txt
```

### Extract Specific Sections

If you need to extract specific parts, first get the document structure:

```bash
# Get full structure
gdrive docs get <document-id> --json | jq '.body.content[]'

# Extract only paragraphs
gdrive docs get <document-id> --json | \
  jq -r '.body.content[] | select(.paragraph != null) | 
         .paragraph.elements[] | select(.textRun != null) | 
         .textRun.content' | \
  tr -d '\n' | fold -w 80
```

## Common Automation Patterns

### Template-Based Document Generation

Create documents from templates:

```bash
#!/bin/bash
TEMPLATE_ID="template-document-id"
CUSTOMER_NAME="Acme Corp"
DATE=$(date +%Y-%m-%d)

# Create new document from template (copy)
NEW_DOC_ID=$(gdrive files copy "$TEMPLATE_ID" "Report for $CUSTOMER_NAME" --json | jq -r '.id')

# Replace placeholders (requires custom script or Slides API for better templating)
# For Docs, you'll need to use find/replace operations
cat > replacements.json <<EOF
[
  {
    "replaceAllText": {
      "containsText": {
        "text": "{{CUSTOMER_NAME}}",
        "matchCase": false
      },
      "replaceText": "$CUSTOMER_NAME"
    }
  },
  {
    "replaceAllText": {
      "containsText": {
        "text": "{{DATE}}",
        "matchCase": false
      },
      "replaceText": "$DATE"
    }
  }
]
EOF

gdrive docs update "$NEW_DOC_ID" --requests-file replacements.json --json
```

### Automated Report Generation

Generate reports with dynamic content:

```bash
#!/bin/bash
DOC_ID="your-document-id"

# Prepare report data
SALES=$(get_sales_total)
ORDERS=$(get_order_count)
DATE=$(date +"%B %d, %Y")

# Create update requests
cat > report-update.json <<EOF
[
  {
    "insertText": {
      "location": {
        "index": 1
      },
      "text": "Daily Sales Report\n\n"
    }
  },
  {
    "updateTextStyle": {
      "range": {
        "startIndex": 1,
        "endIndex": 20
      },
      "textStyle": {
        "bold": true,
        "fontSize": {
          "magnitude": 24,
          "unit": "PT"
        }
      },
      "fields": "bold,fontSize"
    }
  },
  {
    "insertText": {
      "location": {
        "index": 21
      },
      "text": "Date: $DATE\n\nTotal Sales: \$$SALES\nTotal Orders: $ORDERS\n"
    }
  }
]
EOF

gdrive docs update "$DOC_ID" --requests-file report-update.json --json
```

### Document Merge

Combine content from multiple documents:

```bash
#!/bin/bash
# Extract text from source documents
DOC1_TEXT=$(gdrive docs read "doc1-id" --json | jq -r '.text')
DOC2_TEXT=$(gdrive docs read "doc2-id" --json | jq -r '.text')

# Create merged document
MERGED_ID=$(gdrive docs create "Merged Document" --json | jq -r '.id')

# Insert content from both documents
cat > merge-content.json <<EOF
[
  {
    "insertText": {
      "location": {
        "index": 1
      },
      "text": "$DOC1_TEXT\n\n---\n\n$DOC2_TEXT\n"
    }
  }
]
EOF

gdrive docs update "$MERGED_ID" --requests-file merge-content.json --json
```

### Bulk Document Processing

Process multiple documents:

```bash
#!/bin/bash
# List all documents
DOC_IDS=$(gdrive docs list --json | jq -r '.[].id')

# Process each document
for DOC_ID in $DOC_IDS; do
  echo "Processing $DOC_ID"
  
  # Extract text
  TEXT=$(gdrive docs read "$DOC_ID" --json | jq -r '.text')
  
  # Do something with the text
  WORD_COUNT=$(echo "$TEXT" | wc -w)
  echo "Document $DOC_ID has $WORD_COUNT words"
done
```

### Document Formatting Automation

Apply consistent formatting:

```bash
#!/bin/bash
DOC_ID="your-document-id"

# Format all headings
cat > format-headings.json <<EOF
[
  {
    "updateParagraphStyle": {
      "range": {
        "startIndex": 1,
        "endIndex": 1000
      },
      "paragraphStyle": {
        "namedStyleType": "HEADING_1"
      },
      "fields": "namedStyleType"
    }
  }
]
EOF

gdrive docs update "$DOC_ID" --requests-file format-headings.json --json
```

## Tips and Best Practices

1. **Index positions**: Remember that index 1 is the beginning of the document. Use the document structure to find correct indices.

2. **Batch operations**: Group multiple updates into a single batch update request for efficiency.

3. **Text extraction**: Use `docs read` for simple text extraction, `docs get` for full structure analysis.

4. **Error handling**: Always check document IDs exist before performing updates.

5. **Use JSON output**: Always use `--json` for programmatic access.

6. **Handle special characters**: Escape special characters in JSON when inserting text.

7. **Test with small updates**: Test your update operations on small documents first.

8. **Document structure**: Understand the document structure (paragraphs, tables, etc.) before complex updates.

## See Also

- [Google Docs API Documentation](https://developers.google.com/docs/api)
- [Batch Update Reference](https://developers.google.com/docs/api/reference/rest/v1/documents/batchUpdate)
- Main CLI [README](../../README.md)
