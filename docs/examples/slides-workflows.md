# Google Slides Workflows

This guide covers working with Google Slides using the Google Drive CLI.

## Table of Contents

- [Reading Presentations](#reading-presentations)
- [Creating Presentations](#creating-presentations)
- [Template Replacement](#template-replacement-placeholder-substitution)
- [Automated Report Generation](#automated-report-generation)

## Reading Presentations

### Get Presentation Structure

Retrieve the full presentation structure including slides, layouts, and content:

```bash
gdrive slides get <presentation-id> --json
```

This returns detailed metadata about the presentation, including:
- Presentation ID and title
- Page size and dimensions
- List of slides with their properties
- Master slides and layouts
- Page elements (shapes, images, tables)

**Example Output:**
```json
{
  "presentationId": "1abc123...",
  "title": "My Presentation",
  "pageSize": {
    "width": {
      "magnitude": 720,
      "unit": "PT"
    },
    "height": {
      "magnitude": 405,
      "unit": "PT"
    }
  },
  "slides": [
    {
      "objectId": "slide1",
      "pageElements": [
        {
          "objectId": "element1",
          "shape": {
            "shapeType": "TEXT_BOX",
            "text": {
              "textElements": [
                {
                  "textRun": {
                    "content": "Title Slide\n"
                  }
                }
              ]
            }
          }
        }
      ]
    }
  ]
}
```

### Read Presentation Content

Extract text content from all slides:

```bash
# Plain text output
gdrive slides read <presentation-id>

# Structured JSON output
gdrive slides read <presentation-id> --json
```

**Example JSON Output:**
```json
{
  "presentationId": "1abc123...",
  "title": "My Presentation",
  "slideCount": 3,
  "textBySlide": [
    {
      "slideIndex": 0,
      "objectId": "element1",
      "text": "Title Slide\n"
    },
    {
      "slideIndex": 1,
      "objectId": "element2",
      "text": "Content Slide\n\nThis is the body text."
    }
  ]
}
```

### List Presentations

Find presentations in your Drive:

```bash
# List all presentations
gdrive slides list --json

# Filter by parent folder
gdrive slides list --parent <folder-id> --json

# Search with query
gdrive slides list --query "name contains 'Report'" --json

# Paginate through results
gdrive slides list --paginate --json
```

## Creating Presentations

### Create a New Presentation

Create a blank presentation:

```bash
# Create in root
gdrive slides create "My Presentation" --json

# Create in specific folder
gdrive slides create "Q1 Report" --parent <folder-id> --json
```

The command returns the presentation ID which you can use for subsequent operations.

### Create Presentation with Initial Content

Create a presentation and immediately add slides:

```bash
#!/bin/bash
PRES_ID=$(gdrive slides create "New Presentation" --json | jq -r '.id')

# Add a title slide
cat > initial-slide.json <<EOF
[
  {
    "createSlide": {
      "insertionIndex": 0,
      "slideLayoutReference": {
        "predefinedLayout": "TITLE"
      }
    }
  },
  {
    "insertText": {
      "objectId": "title",
      "text": "Welcome"
    }
  }
]
EOF

gdrive slides update "$PRES_ID" --requests-file initial-slide.json --json
```

## Template Replacement (Placeholder Substitution)

The Slides API supports powerful template replacement functionality, allowing you to replace text placeholders throughout a presentation.

### Basic Template Replacement

Replace placeholders using a simple key-value map:

```bash
# Inline JSON
gdrive slides replace <presentation-id> \
  --data '{"{{NAME}}":"Alice","{{DATE}}":"2026-01-24","{{TOTAL}}":"$100K"}' \
  --json

# From file
gdrive slides replace <presentation-id> \
  --file replacements.json \
  --json
```

**Example `replacements.json`:**
```json
{
  "{{NAME}}": "Alice",
  "{{DATE}}": "2026-01-24",
  "{{TOTAL}}": "$100K",
  "{{COMPANY}}": "Acme Corp"
}
```

### How Template Replacement Works

1. **Create a template presentation** with placeholders like `{{NAME}}`, `{{DATE}}`, etc.
2. **Use the replace command** to substitute all occurrences of placeholders with actual values
3. **The replacement is case-sensitive** by default

**Example Template Slide:**
```
Title: Monthly Report for {{NAME}}
Date: {{DATE}}
Total Sales: {{TOTAL}}
```

After replacement:
```
Title: Monthly Report for Alice
Date: 2026-01-24
Total Sales: $100K
```

### Advanced Template Patterns

**Multiple replacements in one call:**
```json
{
  "{{CUSTOMER_NAME}}": "Acme Corporation",
  "{{QUARTER}}": "Q1 2026",
  "{{REVENUE}}": "$500,000",
  "{{EXPENSES}}": "$300,000",
  "{{PROFIT}}": "$200,000",
  "{{EMPLOYEE_COUNT}}": "150"
}
```

**Nested placeholders** (replace in order):
```json
{
  "{{FIRST_NAME}}": "John",
  "{{LAST_NAME}}": "Doe",
  "{{FULL_NAME}}": "John Doe",
  "{{EMAIL}}": "john.doe@example.com"
}
```

## Automated Report Generation

### Monthly Report Generation

Generate monthly reports from a template:

```bash
#!/bin/bash
TEMPLATE_ID="template-presentation-id"

# Generate report for each month
for MONTH in Jan Feb Mar Apr May Jun; do
  REPORT_NAME="${MONTH} 2026 Report"
  
  # Copy template
  NEW_ID=$(gdrive files copy "$TEMPLATE_ID" "$REPORT_NAME" --json | jq -r '.id')
  
  # Prepare replacements
  cat > replacements.json <<EOF
{
  "{{MONTH}}": "$MONTH",
  "{{YEAR}}": "2026",
  "{{REVENUE}}": "$(get_revenue $MONTH)",
  "{{EXPENSES}}": "$(get_expenses $MONTH)",
  "{{PROFIT}}": "$(get_profit $MONTH)"
}
EOF
  
  # Replace placeholders
  gdrive slides replace "$NEW_ID" --file replacements.json --json
  
  echo "Generated: $REPORT_NAME"
done
```

### Customer-Specific Presentations

Generate personalized presentations for each customer:

```bash
#!/bin/bash
TEMPLATE_ID="customer-template-id"

# Read customer data (CSV format: name,email,revenue)
while IFS=, read -r NAME EMAIL REVENUE; do
  PRES_NAME="Report for $NAME"
  
  # Copy template
  NEW_ID=$(gdrive files copy "$TEMPLATE_ID" "$PRES_NAME" --json | jq -r '.id')
  
  # Replace customer-specific data
  cat > replacements.json <<EOF
{
  "{{CUSTOMER_NAME}}": "$NAME",
  "{{CUSTOMER_EMAIL}}": "$EMAIL",
  "{{CUSTOMER_REVENUE}}": "$REVENUE"
}
EOF
  
  gdrive slides replace "$NEW_ID" --file replacements.json --json
  
  # Share with customer (optional)
  # gdrive permissions create "$NEW_ID" --role reader --type user --email "$EMAIL"
  
done < customers.csv
```

### Batch Update Operations

For more complex updates beyond text replacement:

**Example `batch-update.json`:**
```json
[
  {
    "createSlide": {
      "insertionIndex": 1,
      "slideLayoutReference": {
        "predefinedLayout": "TITLE_AND_BODY"
      }
    }
  },
  {
    "insertText": {
      "objectId": "newSlideTitle",
      "text": "New Slide Title"
    }
  },
  {
    "updateSlidesPosition": {
      "slideObjectIds": ["slide2"],
      "insertionIndex": 0
    }
  },
  {
    "deleteObject": {
      "objectId": "oldSlideId"
    }
  }
]
```

Execute batch update:
```bash
gdrive slides update <presentation-id> --requests-file batch-update.json --json
```

### Common Batch Operations

#### Add a New Slide

```json
[
  {
    "createSlide": {
      "insertionIndex": 1,
      "slideLayoutReference": {
        "predefinedLayout": "TITLE_AND_BODY"
      }
    }
  }
]
```

#### Delete a Slide

```json
[
  {
    "deleteObject": {
      "objectId": "slide1"
    }
  }
]
```

#### Insert Text into Shape

```json
[
  {
    "insertText": {
      "objectId": "shapeId",
      "text": "New text content"
    }
  }
]
```

#### Update Shape Properties

```json
[
  {
    "updateShapeProperties": {
      "objectId": "shapeId",
      "shapeProperties": {
        "contentAlignment": "MIDDLE"
      },
      "fields": "contentAlignment"
    }
  }
]
```

#### Duplicate Slide

```json
[
  {
    "duplicateObject": {
      "objectId": "slide1",
      "objectIds": {
        "slide1": "slide1_copy"
      }
    }
  }
]
```

## Common Automation Patterns

### Weekly Status Report

Generate weekly status reports automatically:

```bash
#!/bin/bash
TEMPLATE_ID="weekly-status-template"
WEEK=$(date +%U)
YEAR=$(date +%Y)

# Get week data
COMPLETED=$(get_completed_tasks)
IN_PROGRESS=$(get_in_progress_tasks)
BLOCKED=$(get_blocked_tasks)

# Create report
REPORT_ID=$(gdrive files copy "$TEMPLATE_ID" "Week $WEEK Status Report" --json | jq -r '.id')

cat > replacements.json <<EOF
{
  "{{WEEK}}": "$WEEK",
  "{{YEAR}}": "$YEAR",
  "{{COMPLETED}}": "$COMPLETED",
  "{{IN_PROGRESS}}": "$IN_PROGRESS",
  "{{BLOCKED}}": "$BLOCKED"
}
EOF

gdrive slides replace "$REPORT_ID" --file replacements.json --json
```

### Presentation Analysis

Analyze presentation content:

```bash
#!/bin/bash
PRES_ID="presentation-id"

# Extract all text
TEXT=$(gdrive slides read "$PRES_ID" --json | jq -r '.textBySlide[].text' | tr '\n' ' ')

# Word count
WORD_COUNT=$(echo "$TEXT" | wc -w)
echo "Total words: $WORD_COUNT"

# Slide count
SLIDE_COUNT=$(gdrive slides get "$PRES_ID" --json | jq '.slides | length')
echo "Total slides: $SLIDE_COUNT"

# Average words per slide
AVG_WORDS=$((WORD_COUNT / SLIDE_COUNT))
echo "Average words per slide: $AVG_WORDS"
```

### Bulk Template Processing

Process multiple templates:

```bash
#!/bin/bash
# List all template presentations
TEMPLATES=$(gdrive slides list --query "name contains 'Template'" --json | jq -r '.[].id')

# Process each template
for TEMPLATE_ID in $TEMPLATES; do
  echo "Processing template: $TEMPLATE_ID"
  
  # Generate report from template
  REPORT_ID=$(gdrive files copy "$TEMPLATE_ID" "Generated Report" --json | jq -r '.id')
  
  # Apply replacements
  gdrive slides replace "$REPORT_ID" --file replacements.json --json
  
  echo "Generated report: $REPORT_ID"
done
```

### Dynamic Chart Data

While Slides API doesn't directly update charts, you can use this pattern:

```bash
#!/bin/bash
# 1. Generate data in Sheets
SHEET_ID=$(gdrive sheets create "Chart Data" --json | jq -r '.id')
gdrive sheets values update "$SHEET_ID" "A1:B5" \
  --values '[[\"Month\",\"Sales\"],[\"Jan\",100],[\"Feb\",150],[\"Mar\",200]]' \
  --value-input-option USER_ENTERED

# 2. Create presentation with linked chart
# (Manual step: Create slide with chart linked to the sheet)

# 3. Update sheet data to refresh chart
gdrive sheets values update "$SHEET_ID" "B2" \
  --values '[[120]]' \
  --value-input-option USER_ENTERED
```

## Tips and Best Practices

1. **Template design**: Use clear, unique placeholders like `{{CUSTOMER_NAME}}` that won't conflict with regular text.

2. **Case sensitivity**: Template replacement is case-sensitive. Use consistent casing in your templates.

3. **Placeholder format**: Use double curly braces `{{PLACEHOLDER}}` to avoid conflicts with single braces used in formatting.

4. **Batch operations**: Use batch updates for complex structural changes, use replace for simple text substitution.

5. **Slide IDs**: When working with specific slides, use object IDs from the presentation structure.

6. **Test replacements**: Test your replacement files on a copy before applying to important presentations.

7. **Use JSON output**: Always use `--json` for programmatic access.

8. **Handle errors**: Check that presentation IDs exist and you have write access before updates.

9. **Template versioning**: Keep templates in a dedicated folder and version them.

10. **Performance**: For large presentations, consider breaking updates into smaller batches.

## See Also

- [Google Slides API Documentation](https://developers.google.com/slides/api)
- [Batch Update Reference](https://developers.google.com/slides/api/reference/rest/v1/presentations/batchUpdate)
- Main CLI [README](../../README.md)
