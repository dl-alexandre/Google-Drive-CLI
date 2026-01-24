# Google Sheets Workflows

This guide covers working with Google Sheets using the Google Drive CLI.

## Table of Contents

- [Reading Spreadsheet Data](#reading-spreadsheet-data)
- [Writing/Updating Values](#writingupdating-values)
- [Creating Spreadsheets](#creating-spreadsheets)
- [Batch Operations](#batch-operations)
- [Common Automation Patterns](#common-automation-patterns)

## Reading Spreadsheet Data

### Get Values from a Range

Read values from a specific range using A1 notation:

```bash
# Read a single cell
gdrive sheets values get <spreadsheet-id> "Sheet1!A1" --json

# Read a range
gdrive sheets values get <spreadsheet-id> "Sheet1!A1:C10" --json

# Read entire column
gdrive sheets values get <spreadsheet-id> "Sheet1!A:A" --json

# Read entire row
gdrive sheets values get <spreadsheet-id> "Sheet1!1:1" --json
```

**Example Output:**
```json
{
  "range": "Sheet1!A1:C3",
  "majorDimension": "ROWS",
  "values": [
    ["Name", "Age", "City"],
    ["Alice", "30", "New York"],
    ["Bob", "25", "San Francisco"]
  ]
}
```

### Get Spreadsheet Metadata

Retrieve information about the spreadsheet structure:

```bash
gdrive sheets get <spreadsheet-id> --json
```

This returns details about sheets, locale, timezone, and other metadata.

### List Spreadsheets

Find spreadsheets in your Drive:

```bash
# List all spreadsheets
gdrive sheets list --json

# Filter by parent folder
gdrive sheets list --parent <folder-id> --json

# Search with query
gdrive sheets list --query "name contains 'Report'" --json

# Paginate through results
gdrive sheets list --paginate --json
```

## Writing/Updating Values

### Update Values

Replace values in a range:

```bash
# Update with inline JSON
gdrive sheets values update <spreadsheet-id> "Sheet1!A1" \
  --values '[[\"Name\",\"Age\"],[\"Alice\",30]]' \
  --value-input-option USER_ENTERED

# Update from file
gdrive sheets values update <spreadsheet-id> "Sheet1!A1" \
  --values-file data.json \
  --value-input-option USER_ENTERED
```

**Value Input Options:**
- `RAW`: Values are stored exactly as entered (default)
- `USER_ENTERED`: Values are parsed as if typed by a user (formulas, dates, etc.)

**Example `data.json`:**
```json
[
  ["Product", "Price", "Quantity"],
  ["Widget A", 19.99, 100],
  ["Widget B", 29.99, 50]
]
```

### Append Values

Add new rows or columns to existing data:

```bash
# Append rows
gdrive sheets values append <spreadsheet-id> "Sheet1!A:B" \
  --values '[[\"Widget C\",39.99]]' \
  --value-input-option USER_ENTERED

# Append from file
gdrive sheets values append <spreadsheet-id> "Sheet1!A:B" \
  --values-file new-rows.json
```

**Note:** The range should specify the columns you're appending to (e.g., `A:B` for columns A and B).

### Clear Values

Remove values from a range without deleting cells:

```bash
gdrive sheets values clear <spreadsheet-id> "Sheet1!A2:C10" --json
```

## Creating Spreadsheets

### Create a New Spreadsheet

```bash
# Create in root
gdrive sheets create "My Spreadsheet" --json

# Create in specific folder
gdrive sheets create "Q1 Report" --parent <folder-id> --json
```

The command returns the spreadsheet ID which you can use for subsequent operations.

## Batch Operations

Batch operations allow you to perform multiple updates in a single API call, including formatting, formulas, and structural changes.

### Batch Update Format

Create a JSON file with batch update requests:

**Example `batch-update.json`:**
```json
[
  {
    "updateSpreadsheetProperties": {
      "properties": {
        "title": "Updated Title"
      },
      "fields": "title"
    }
  },
  {
    "updateCells": {
      "range": {
        "sheetId": 0,
        "startRowIndex": 0,
        "endRowIndex": 1,
        "startColumnIndex": 0,
        "endColumnIndex": 3
      },
      "rows": [
        {
          "values": [
            {
              "userEnteredFormat": {
                "backgroundColor": {
                  "red": 0.2,
                  "green": 0.4,
                  "blue": 0.8
                },
                "textFormat": {
                  "bold": true,
                  "foregroundColor": {
                    "red": 1.0,
                    "green": 1.0,
                    "blue": 1.0
                  }
                }
              }
            },
            {},
            {}
          ]
        }
      ],
      "fields": "userEnteredFormat"
    }
  },
  {
    "setDataValidation": {
      "range": {
        "sheetId": 0,
        "startRowIndex": 1,
        "endRowIndex": 100,
        "startColumnIndex": 2,
        "endColumnIndex": 3
      },
      "rule": {
        "condition": {
          "type": "NUMBER_GREATER",
          "values": [
            {
              "userEnteredValue": "0"
            }
          ]
        },
        "showCustomUi": true
      }
    }
  }
]
```

### Execute Batch Update

```bash
# From file
gdrive sheets batch-update <spreadsheet-id> --requests-file batch-update.json --json

# From stdin
cat batch-update.json | gdrive sheets batch-update <spreadsheet-id> --requests-file - --json
```

### Common Batch Operations

**Add a new sheet:**
```json
[
  {
    "addSheet": {
      "properties": {
        "title": "New Sheet"
      }
    }
  }
]
```

**Delete a sheet:**
```json
[
  {
    "deleteSheet": {
      "sheetId": 123456789
    }
  }
]
```

**Insert rows:**
```json
[
  {
    "insertDimension": {
      "range": {
        "sheetId": 0,
        "dimension": "ROWS",
        "startIndex": 5,
        "endIndex": 7
      }
    }
  }
]
```

**Apply formula:**
```json
[
  {
    "updateCells": {
      "range": {
        "sheetId": 0,
        "startRowIndex": 1,
        "endRowIndex": 2,
        "startColumnIndex": 3,
        "endColumnIndex": 4
      },
      "rows": [
        {
          "values": [
            {
              "userEnteredValue": {
                "formulaValue": "=SUM(A1:B1)"
              }
            }
          ]
        }
      ],
      "fields": "userEnteredValue"
    }
  }
]
```

## Common Automation Patterns

### Data Export Script

Export spreadsheet data to CSV:

```bash
#!/bin/bash
SPREADSHEET_ID="your-spreadsheet-id"
RANGE="Sheet1!A1:Z1000"

# Get data as JSON
gdrive sheets values get "$SPREADSHEET_ID" "$RANGE" --json | \
  jq -r '.values[] | @csv' > export.csv
```

### Automated Data Entry

Add data from a CSV file:

```bash
#!/bin/bash
SPREADSHEET_ID="your-spreadsheet-id"

# Convert CSV to JSON array
cat data.csv | jq -R -s -c 'split("\n") | map(split(","))' | \
  gdrive sheets values append "$SPREADSHEET_ID" "Sheet1!A:B" --values-file -
```

### Daily Report Generation

Create a daily report with formatted data:

```bash
#!/bin/bash
SPREADSHEET_ID="your-spreadsheet-id"
DATE=$(date +%Y-%m-%d)

# Prepare data
cat > daily-data.json <<EOF
[
  ["Date", "Sales", "Orders"],
  ["$DATE", "$(get_sales)", "$(get_orders)"]
]
EOF

# Append to spreadsheet
gdrive sheets values append "$SPREADSHEET_ID" "Sheet1!A:C" \
  --values-file daily-data.json \
  --value-input-option USER_ENTERED
```

### Template-Based Spreadsheet Creation

Create spreadsheets from templates:

```bash
#!/bin/bash
# 1. Create template spreadsheet
TEMPLATE_ID=$(gdrive sheets create "Template" --json | jq -r '.id')

# 2. Set up template structure
gdrive sheets values update "$TEMPLATE_ID" "Sheet1!A1" \
  --values '[[\"Name\",\"Email\",\"Department\"]]' \
  --value-input-option USER_ENTERED

# 3. Copy template for new use case
NEW_ID=$(gdrive files copy "$TEMPLATE_ID" "New Report" --json | jq -r '.id')

# 4. Populate with data
gdrive sheets values update "$NEW_ID" "Sheet1!A2" \
  --values-file data.json \
  --value-input-option USER_ENTERED
```

### Conditional Formatting

Apply conditional formatting via batch update:

```json
[
  {
    "addConditionalFormatRule": {
      "rule": {
        "ranges": [
          {
            "sheetId": 0,
            "startRowIndex": 1,
            "endRowIndex": 100,
            "startColumnIndex": 2,
            "endColumnIndex": 3
          }
        ],
        "booleanRule": {
          "condition": {
            "type": "NUMBER_GREATER",
            "values": [
              {
                "userEnteredValue": "100"
              }
            ]
          },
          "format": {
            "backgroundColor": {
              "red": 0.8,
              "green": 0.9,
              "blue": 0.8
            }
          }
        }
      },
      "index": 0
    }
  }
]
```

### Error Handling

Handle errors gracefully in scripts:

```bash
#!/bin/bash
SPREADSHEET_ID="your-spreadsheet-id"

if ! OUTPUT=$(gdrive sheets values get "$SPREADSHEET_ID" "A1" --json 2>&1); then
  ERROR=$(echo "$OUTPUT" | jq -r '.error.message // "Unknown error"')
  echo "Error: $ERROR" >&2
  exit 1
fi

# Process successful output
echo "$OUTPUT" | jq '.values'
```

## Tips and Best Practices

1. **Use USER_ENTERED for formulas**: When inserting formulas or dates, use `--value-input-option USER_ENTERED`

2. **Batch operations for efficiency**: Group multiple updates into a single batch update request

3. **Handle large ranges carefully**: For very large ranges, consider pagination or chunking

4. **Store spreadsheet IDs**: Save spreadsheet IDs for reuse in scripts

5. **Use JSON output**: Always use `--json` for programmatic access

6. **Validate ranges**: Ensure A1 notation is correct before running updates

7. **Test with small ranges**: Test your operations on small ranges before applying to large datasets

## See Also

- [Google Sheets API Documentation](https://developers.google.com/sheets/api)
- [A1 Notation Guide](https://developers.google.com/sheets/api/guides/concepts#a1_notation)
- Main CLI [README](../../README.md)
