# Google Drive CLI Documentation

Welcome to the Google Drive CLI documentation. This directory contains comprehensive guides and examples for using the Google Workspace features of the CLI.

## Overview

The Google Drive CLI provides powerful command-line tools for interacting with Google Workspace services:

- **Google Sheets**: Read, write, and manage spreadsheet data
- **Google Docs**: Extract content, create documents, and perform batch updates
- **Google Slides**: Generate presentations, replace placeholders, and automate slide creation
- **Admin SDK**: Manage users, groups, and group memberships in Google Workspace

## Documentation Index

### [Sheets Workflows](./examples/sheets-workflows.md)
Complete guide to working with Google Sheets:
- Reading spreadsheet data
- Writing and updating values
- Creating spreadsheets
- Batch operations
- Common automation patterns

### [Docs Workflows](./examples/docs-workflows.md)
Guide to Google Docs operations:
- Reading document content
- Creating documents
- Updating documents
- Text extraction

### [Slides Workflows](./examples/slides-workflows.md)
Automated presentation generation:
- Reading presentations
- Creating presentations
- Template replacement (placeholder substitution)
- Automated report generation

### [Admin Workflows](./examples/admin-workflows.md)
Google Workspace administration:
- User provisioning
- Group management
- Bulk operations
- Service account setup guide

## Getting Started

Before using these features, ensure you have:

1. **Authenticated** with the CLI:
   ```bash
   gdrive auth login --preset workspace-basic
   ```

2. **Required scopes** for your use case:
   - Sheets: `https://www.googleapis.com/auth/spreadsheets`
   - Docs: `https://www.googleapis.com/auth/documents`
   - Slides: `https://www.googleapis.com/auth/presentations`
   - Admin SDK: Service account with domain-wide delegation

## Common Patterns

### JSON Output

All commands support `--json` for machine-readable output:

```bash
gdrive sheets values get <spreadsheet-id> "Sheet1!A1:C10" --json
```

### File Input

Many commands accept input from files or stdin:

```bash
# From file
gdrive sheets values update <id> "A1" --values-file data.json

# From stdin
echo '[[1,2,3]]' | gdrive sheets values append <id> "A:B" --values-file -
```

### Error Handling

The CLI provides structured error output:

```bash
gdrive sheets values get <id> "A1" --json
# On error, returns JSON with error details
```

## Examples Directory

See the [`examples/`](../examples/) directory for sample JSON files:
- `examples/sheets/batch-update.json` - Sheets batch update examples
- `examples/docs/batch-update.json` - Docs batch update examples
- `examples/slides/replacements.json` - Slides placeholder replacements

## Need Help?

- Check the main [README](../README.md) for installation and basic usage
- Review the specific workflow guides linked above
- Use `gdrive <command> --help` for command-specific help
