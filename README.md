# Google Drive CLI

A **fast**, **lightweight**, and **AI-agent friendly** CLI for Google Drive. Manage files with zero friction.

## Why gdrv?

| Problem | Solution |
|---------|----------|
| Manual Google Drive work | Automate everything from CLI |
| Slow, heavy tooling | Go binary, fast startup |
| Not AI-agent friendly | JSON output, explicit flags, clean exit codes |

## Features

- **Complete Google Drive Integration**: Upload, download, list, search, and manage files and folders
- **Google Workspace Integration**: Full support for Google Sheets, Docs, and Slides with read/write operations
- **Admin SDK Support**: Manage Google Workspace users and groups via Admin SDK Directory API
- **Authentication**: OAuth2 with device code fallback, multiple profiles, secure credential storage, service account support
- **Shared Drives Support**: Full support for Google Workspace Shared Drives
- **Advanced Safety Controls**: Dry-run mode, confirmation prompts, idempotent operations
- **Rich CLI Interface**: 50+ commands with help, examples, and multiple output formats (JSON, table)
- **Production Logging**: Structured logging with debug mode and trace correlation
- **Cross-Platform**: Works on macOS, Linux, and Windows

### Advanced APIs

The following Google Drive APIs have been implemented to extend `gdrv` with advanced auditing, metadata management, and synchronization capabilities.

#### Drive Activity API (v2)

Monitor and audit file and folder activity across your Google Drive with detailed activity streams.

**Use Cases:**
- Audit trails for compliance and security monitoring
- Track who accessed, modified, or shared files
- Monitor file lifecycle events (creation, modification, deletion, permission changes)
- Generate activity reports for specific files, folders, or time ranges
- Detect suspicious or unauthorized access patterns

**API Documentation:** [Drive Activity API v2](https://developers.google.com/drive/activity/v2)

**Required OAuth Scopes:**
- `https://www.googleapis.com/auth/drive.activity` - Full activity access
- `https://www.googleapis.com/auth/drive.activity.readonly` - Read-only activity access

**Example Commands:**

```bash
# Query recent activity for all accessible files
gdrv activity query --json

# Query activity for a specific file
gdrv activity query --file-id 1abc123... --json

# Query activity within a time range
gdrv activity query --start-time "2026-01-01T00:00:00Z" --end-time "2026-01-31T23:59:59Z" --json

# Query activity for a folder (including descendants)
gdrv activity query --folder-id 0ABC123... --ancestor-name "folders/0ABC123..." --json

# Filter by activity types (edit, comment, share, permission_change, etc.)
gdrv activity query --action-types "edit,share,permission_change" --json

# Get activity for a specific user
gdrv activity query --user user@example.com --json

# Paginate through activity results
gdrv activity query --limit 100 --page-token "TOKEN" --json
```

**Command Flags:**
- `--file-id`: Filter by specific file ID
- `--folder-id`: Filter by folder ID (includes descendants)
- `--ancestor-name`: Filter by ancestor folder (e.g., "folders/123")
- `--start-time`: Start of time range (RFC3339 format)
- `--end-time`: End of time range (RFC3339 format)
- `--action-types`: Comma-separated action types (edit, comment, share, permission_change, move, delete, restore, etc.)
- `--user`: Filter by user email
- `--limit`: Maximum results per page
- `--page-token`: Pagination token
- `--json`: JSON output

#### Drive Labels API (v2)

Apply custom metadata taxonomy and structured labeling to files and folders for advanced organization and workflows.

**Use Cases:**
- Custom metadata classification systems (e.g., document types, project codes, retention policies)
- Enterprise content management and governance
- Automated workflows based on label values
- Searchable structured metadata beyond file properties
- Compliance and records management tagging

**API Documentation:** [Drive Labels API v2](https://developers.google.com/drive/labels/overview)

**Required OAuth Scopes:**
- `https://www.googleapis.com/auth/drive.labels` - Full labels access
- `https://www.googleapis.com/auth/drive.labels.readonly` - Read-only labels access
- `https://www.googleapis.com/auth/drive.admin.labels` - Admin label management (requires Admin SDK)
- `https://www.googleapis.com/auth/drive.admin.labels.readonly` - Read-only admin labels

**Example Commands:**

```bash
# List available labels
gdrv labels list --json

# Get label schema
gdrv labels get <label-id> --json

# List labels applied to a file
gdrv labels file list <file-id> --json

# Apply a label to a file
gdrv labels file apply <file-id> <label-id> --fields '{"field1":"value1","field2":"value2"}' --json

# Update label fields on a file
gdrv labels file update <file-id> <label-id> --fields '{"field1":"new_value"}' --json

# Remove a label from a file
gdrv labels file remove <file-id> <label-id>

# Search files by label
gdrv files list --query "labels/<label-id> exists" --json

# Create a label (admin only)
gdrv labels create "Document Type" --fields "Type:choice:Contract,Invoice,Report" --json

# Publish a label (admin only)
gdrv labels publish <label-id>

# Disable a label (admin only)
gdrv labels disable <label-id>
```

**Command Flags:**
- `--fields`: JSON object of field values (key-value pairs)
- `--label-id`: Label ID to apply/modify
- `--view`: Label view mode (LABEL_VIEW_FULL, LABEL_VIEW_BASIC)
- `--customer`: Customer ID for admin operations
- `--json`: JSON output

#### Drive Changes API (v3)

Track changes to files and folders for real-time synchronization and automation.

**Use Cases:**
- Build synchronization tools (like Dropbox sync)
- Real-time monitoring of Drive changes
- Incremental backup systems
- Webhook-triggered automation workflows
- Change notification systems

**API Documentation:** [Drive Changes API v3](https://developers.google.com/drive/api/v3/reference/changes)

**Required OAuth Scopes:**
- `https://www.googleapis.com/auth/drive` - Full Drive access (includes changes)
- `https://www.googleapis.com/auth/drive.readonly` - Read-only changes access
- `https://www.googleapis.com/auth/drive.file` - Changes for files created by the app

**Example Commands:**

```bash
# Get the starting page token (start of change log)
gdrv changes start-page-token --json

# List changes since a page token
gdrv changes list --page-token "12345" --json

# List changes with auto-pagination
gdrv changes list --page-token "12345" --paginate --json

# List changes for a specific Shared Drive
gdrv changes list --page-token "12345" --drive-id <drive-id> --json

# Watch for changes (webhook setup)
gdrv changes watch --page-token "12345" --webhook-url "https://example.com/webhook" --json

# Stop watching for changes
gdrv changes stop <channel-id> <resource-id>

# List changes with specific fields
gdrv changes list --page-token "12345" --fields "nextPageToken,newStartPageToken,changes(fileId,time,removed,file(name,mimeType))" --json

# List changes including removed files
gdrv changes list --page-token "12345" --include-removed true --json

# List changes for specific change types
gdrv changes list --page-token "12345" --restrict-to-my-drive false --json
```

**Command Flags:**
- `--page-token`: Page token to list changes from (required for list)
- `--drive-id`: Shared Drive ID to monitor
- `--include-corpus-removals`: Include changes outside the target corpus
- `--include-items-from-all-drives`: Include items from all drives
- `--include-permissions-for-view`: Include permissions with published view
- `--include-removed`: Include removed items
- `--restrict-to-my-drive`: Restrict changes to My Drive only
- `--supports-all-drives`: Support all drives (Shared Drives)
- `--webhook-url`: Webhook URL for change notifications
- `--expiration`: Webhook expiration time
- `--limit`: Maximum results per page
- `--paginate`: Auto-paginate through all changes
- `--json`: JSON output

#### Permissions Enhancements

Enhanced permission auditing and access analysis tools built on top of the existing Drive Permissions API.

**Use Cases:**
- Security audits and compliance reporting
- Detect oversharing and public access risks
- Analyze permission inheritance and effective access
- Bulk permission cleanup and remediation
- Generate access reports for specific users or groups

**Example Commands:**

```bash
# Audit all files with public access
gdrv permissions audit public --json

# Audit all files shared with external domains
gdrv permissions audit external --json

# Audit permissions for a specific user
gdrv permissions audit user user@example.com --json

# Find files with "anyone with link" access
gdrv permissions audit anyone-with-link --json

# Analyze permission inheritance for a folder
gdrv permissions analyze <folder-id> --recursive --json

# Generate permission report for a file/folder
gdrv permissions report <file-id> --json

# Bulk remove public access
gdrv permissions bulk-remove public --folder-id <folder-id> --dry-run

# Bulk change role (e.g., downgrade all "writer" to "reader")
gdrv permissions bulk-update <folder-id> --from-role writer --to-role reader --dry-run

# Find files accessible by a specific email
gdrv permissions search --email user@example.com --json

# List all files with "commenter" access
gdrv permissions search --role commenter --json
```

**Command Flags:**
- `--recursive`: Include descendants (for folders)
- `--dry-run`: Preview changes without executing
- `--from-role`: Source role for bulk operations
- `--to-role`: Target role for bulk operations
- `--email`: Filter by email address
- `--role`: Filter by permission role
- `--type`: Filter by permission type (user, group, domain, anyone)
- `--json`: JSON output

---

**Implementation Status:** ✅ All APIs Fully Implemented

All four advanced APIs have been successfully implemented and are ready for use. The implementation includes:

- ✅ **Drive Activity API (v2)** - Query file and folder activity with comprehensive filtering
- ✅ **Drive Labels API (v2)** - Manage labels and apply structured metadata to files
- ✅ **Drive Changes API (v3)** - Track file changes for sync and automation workflows
- ✅ **Permissions Enhancements** - Audit, analyze, and bulk-manage permissions

**Available Scope Presets:**

All new scope presets are now available for authentication:

- `workspace-activity`: Workspace + Activity API (read-only)
- `workspace-labels`: Workspace + Labels API
- `workspace-sync`: Workspace + Changes API
- `workspace-complete`: All Workspace APIs + Activity + Labels + Changes

Use these presets with any authentication command:

```bash
gdrv auth login --preset workspace-activity
gdrv auth login --preset workspace-complete
gdrv auth service-account --key-file ./key.json --preset workspace-complete
```

## Installation

### Install Script (Recommended)

```bash
# Install to ~/.local/bin (ensure it's on your PATH)
curl -fsSL https://raw.githubusercontent.com/dl-alexandre/Google-Drive-CLI/master/install.sh | bash
```

### Homebrew (Tap)

```bash
brew tap dl-alexandre/tap
brew install gdrv
```

### Download Binary

Download the latest release from the [releases page](https://github.com/dl-alexandre/Google-Drive-CLI/releases).

```bash
# Make executable and move to PATH
chmod +x gdrv
sudo mv gdrv /usr/local/bin/
```

### Build from Source

```bash
git clone https://github.com/dl-alexandre/Google-Drive-CLI.git
cd Google-Drive-CLI
go build -o gdrv ./cmd/gdrv
```

## Quick Start

1. **(Optional) Configure a custom OAuth client**:
   ```bash
   # By default, gdrv uses the bundled OAuth client.
   # To use your own client, set GDRV_CLIENT_ID
   # and GDRV_CLIENT_SECRET only if your client type requires it.
   export GDRV_CLIENT_ID="your-client-id"
   export GDRV_CLIENT_SECRET="your-client-secret"
   ```

2. **Authenticate**:
   ```bash
   gdrv auth login --preset workspace-basic
   ```

3. **List files**:
   ```bash
   gdrv files list
   ```

4. **Upload a file**:
   ```bash
   gdrv files upload myfile.txt
   ```

5. **Download a file**:
   ```bash
   gdrv files download 1abc123... --output downloaded.txt
   ```

6. **Download a Google Doc as text**:
   ```bash
   gdrv files download 1abc123... --doc
   ```

## Agent Quickstart

This CLI is designed to be used by AI agents and automation scripts. Key features for agent usage:

### JSON Output

Always use `--json` for machine-readable output:

```bash
# List files as JSON
gdrv files list --json

# Get file metadata
gdrv files get 1abc123... --json

# Upload returns the created file object
gdrv files upload report.pdf --json
```

### Pagination Control

Use `--paginate` to automatically fetch all pages:

```bash
# Get ALL files (auto-pagination)
gdrv files list --paginate --json

# Get all trashed files
gdrv files list-trashed --paginate --json

# Get all Shared Drives
gdrv drives list --paginate --json
```

Or control pagination manually:

```bash
# Get first page
gdrv files list --limit 100 --json

# Use nextPageToken from response for next page
gdrv files list --limit 100 --page-token "TOKEN_FROM_PREVIOUS" --json
```

### Sorting and Filtering

```bash
# Sort by modified time (newest first)
gdrv files list --order-by "modifiedTime desc" --json

# Search by name
gdrv files list --query "name contains 'report'" --json

# Combined: recent PDFs
gdrv files list --query "mimeType = 'application/pdf'" --order-by "modifiedTime desc" --json
```

### Non-Interactive Mode

Destructive commands run without prompts by default. Use `--dry-run` to preview:

```bash
# Preview what would be deleted
gdrv files delete 1abc123... --dry-run

# Actually delete (no prompt)
gdrv files delete 1abc123...

# Permanently delete (bypasses trash)
gdrv files delete 1abc123... --permanent
```

### Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | Authentication required |
| 3 | Invalid argument |
| 4 | Resource not found |
| 5 | Permission denied |
| 6 | Rate limited |

### Agent Best Practices

1. **Always use `--json`** - Parse structured output, not human-readable tables
2. **Use `--paginate`** - Don't miss items due to pagination limits
3. **Check exit codes** - Handle errors programmatically
4. **Use file IDs** - More reliable than paths for Shared Drives
5. **Use `--dry-run`** - Preview destructive operations before executing

## Authentication

The CLI supports multiple authentication methods and scope presets. OAuth uses Authorization Code + PKCE for installed/desktop apps (public client). Any bundled client secret is treated as public and not relied on for security.

### OAuth Client Sources and Precedence

Credentials are resolved in this order:
1. CLI flags (`--client-id`, `--client-secret`)
2. Environment variables (`GDRV_CLIENT_ID`, `GDRV_CLIENT_SECRET`)
3. Config file (`oauthClientId`, `oauthClientSecret`)
4. Bundled OAuth client (release builds)

No partial overrides: if any OAuth client variable is set, all required OAuth client fields must be set (client ID always; secret only if your client type requires it).

**Config file path (defaults):**
- macOS: `~/Library/Application Support/gdrv/config.json`
- Linux: `~/.config/gdrv/config.json`
- Windows: `%APPDATA%\\gdrv\\config.json`
- Override with `GDRV_CONFIG_DIR`

**Contributor/CI policy:** set `GDRV_REQUIRE_CUSTOM_OAUTH=1` to refuse bundled credentials.

Bundled client credentials may rotate between releases. If you see `invalid_client` errors with the bundled client, upgrade or configure a custom client.

### Shared OAuth Client Notes

- The bundled client secret is public and not relied on for security (PKCE is used).
- The shared client is hosted in a dedicated Google Cloud project with quota monitoring and a rotation plan.
- If the shared client is disabled or rotated, the CLI will instruct you to upgrade or configure a custom client.

### Token Storage

- Preferred: system keyring (Keychain / Secret Service / Credential Manager).
- Fallback: encrypted file storage at `.../credentials/<profile>.enc` with `0600` permissions and a local key file at `.../.keyfile`.
- Plain file storage is development-only and must be explicitly forced.
- `gdrv auth logout` removes local credentials only (does not revoke remote consent).

### Custom OAuth Client Prerequisites

If you want to use your own OAuth client:
1. Create a project in Google Cloud Console
2. Enable the Google Drive API
3. Create OAuth 2.0 credentials (Desktop application)
4. Set credentials via environment variables or command flags:

```bash
export GDRV_CLIENT_ID="your-client-id"
export GDRV_CLIENT_SECRET="your-client-secret" # only if required by your client type

gdrv auth login --client-id "your-client-id" --client-secret "your-client-secret"
```

### OAuth2 Flow (Recommended)
```bash
gdrv auth login
```
Opens a browser for authentication using a loopback redirect on `127.0.0.1` with an ephemeral port.
If the browser cannot be opened, gdrv will fall back to manual code entry.

**Manual fallback (headless):**
```bash
gdrv auth login --no-browser
```
Prompts for a manual code paste after you approve access in a browser. This fallback is used for headless environments, browser launch failures, or loopback binding issues. You can force it with `--no-browser` or `GDRV_NO_BROWSER=1`.

### Device Code Flow (Headless)
```bash
gdrv auth device
```
Displays a code to enter at https://www.google.com/device.

### Service Account Authentication
```bash
gdrv auth service-account --key-file ./service-account-key.json --preset workspace-basic
```
Loads credentials from a service account JSON key file. Use `--impersonate-user` for Admin SDK scopes.

### Scope Presets

| Preset | Description | Use Case |
|--------|-------------|----------|
| `workspace-basic` | Read-only Drive, Sheets, Docs, Slides, Labels | Viewing and downloading |
| `workspace-full` | Full Drive, Sheets, Docs, Slides, Labels | Editing and management |
| `admin` | Admin Directory users and groups, Admin Labels | User/group/label administration |
| `workspace-with-admin` | Workspace full + Admin Directory + Admin Labels | Full workspace + admin |
| `workspace-activity` | Workspace basic + Activity API | Read-only with activity auditing |
| `workspace-labels` | Workspace full + Labels API | Full access with label management |
| `workspace-sync` | Workspace full + Changes API | Full access with change tracking |
| `workspace-complete` | All Workspace APIs + Activity + Labels + Changes | Complete API access |

Use `workspace-basic` for least-privilege read-only access; use `workspace-full` only when write access is required. Use the specialized presets (`workspace-activity`, `workspace-labels`, `workspace-sync`, `workspace-complete`) when you need the advanced APIs.

```bash
# Basic presets
gdrv auth login --preset workspace-basic
gdrv auth login --preset workspace-full
gdrv auth login --preset admin
gdrv auth login --preset workspace-with-admin

# Advanced API presets
gdrv auth login --preset workspace-activity
gdrv auth login --preset workspace-labels
gdrv auth login --preset workspace-sync
gdrv auth login --preset workspace-complete

# Device code flow
gdrv auth device --preset workspace-basic

# Service account
gdrv auth service-account --key-file ./key.json --preset workspace-complete
```

### Custom Scopes

```bash
gdrv auth login --scopes "https://www.googleapis.com/auth/drive.file,https://www.googleapis.com/auth/spreadsheets.readonly"
gdrv auth service-account --key-file ./key.json --scopes "https://www.googleapis.com/auth/drive.file"
```

Available scopes:

**Drive Scopes:**
- `https://www.googleapis.com/auth/drive` - Full Drive access
- `https://www.googleapis.com/auth/drive.file` - Per-file access
- `https://www.googleapis.com/auth/drive.readonly` - Read-only Drive access
- `https://www.googleapis.com/auth/drive.metadata.readonly` - Read-only metadata

**Workspace Scopes:**
- `https://www.googleapis.com/auth/spreadsheets` - Full Sheets access
- `https://www.googleapis.com/auth/spreadsheets.readonly` - Read-only Sheets
- `https://www.googleapis.com/auth/documents` - Full Docs access
- `https://www.googleapis.com/auth/documents.readonly` - Read-only Docs
- `https://www.googleapis.com/auth/presentations` - Full Slides access
- `https://www.googleapis.com/auth/presentations.readonly` - Read-only Slides

**Admin SDK Scopes:**
- `https://www.googleapis.com/auth/admin.directory.user` - User management
- `https://www.googleapis.com/auth/admin.directory.user.readonly` - Read-only users
- `https://www.googleapis.com/auth/admin.directory.group` - Group management
- `https://www.googleapis.com/auth/admin.directory.group.readonly` - Read-only groups

**Advanced API Scopes:**
- `https://www.googleapis.com/auth/drive.activity` - Full Activity API access
- `https://www.googleapis.com/auth/drive.activity.readonly` - Read-only Activity
- `https://www.googleapis.com/auth/drive.labels` - Full Labels access
- `https://www.googleapis.com/auth/drive.labels.readonly` - Read-only Labels
- `https://www.googleapis.com/auth/drive.admin.labels` - Admin label management
- `https://www.googleapis.com/auth/drive.admin.labels.readonly` - Read-only admin labels

### Multiple Profiles
```bash
# Create and switch profiles
gdrv auth login --profile work
gdrv auth login --profile personal

# Use specific profile
gdrv --profile work files list
```

### OAuth Testing-Mode Limits
If your OAuth consent screen is in testing mode, refresh tokens expire after 7 days and Google enforces a 100 refresh-token issuance cap per client. If you see repeated `invalid_grant` errors, re-authenticate and revoke unused tokens in Google Cloud Console or move the app to production to avoid the testing-mode limits.

## Commands

### File Operations
```bash
gdrv files upload <file>          # Upload file
gdrv files download <file-id>     # Download file
gdrv files list                   # List files
gdrv files delete <file-id>       # Delete file
gdrv files trash <file-id>        # Move to trash
gdrv files restore <file-id>      # Restore from trash
gdrv files revisions <file-id>    # List revisions
```

### Folder Operations
```bash
gdrv folders create <name>        # Create folder
gdrv folders list <folder-id>     # List contents
gdrv folders delete <folder-id>   # Delete folder
gdrv folders move <id> <parent>   # Move folder
```

### Permission Management
```bash
gdrv permissions list <file-id>           # List permissions
gdrv permissions create <file-id> --type user --email user@example.com --role reader
gdrv permissions update <file-id> <perm-id> --role writer
gdrv permissions delete <file-id> <perm-id>
gdrv permissions public <file-id>         # Create public link
```

### Google Sheets Operations

Manage Google Sheets spreadsheets with full read and write capabilities.

**Required OAuth Scopes:**
- Read operations: `https://www.googleapis.com/auth/spreadsheets.readonly` or `https://www.googleapis.com/auth/spreadsheets`
- Write operations: `https://www.googleapis.com/auth/spreadsheets`
- Use preset: `workspace-basic` (read-only) or `workspace-full` (read/write)

**API Documentation:** [Google Sheets API v4](https://developers.google.com/sheets/api)

```bash
# List spreadsheets
gdrv sheets list                                # List all spreadsheets
gdrv sheets list --parent <folder-id>          # List spreadsheets in a folder
gdrv sheets list --query "name contains 'Report'" --json
gdrv sheets list --paginate --json            # Get all spreadsheets

# Create a spreadsheet
gdrv sheets create "My Spreadsheet"           # Create a new spreadsheet
gdrv sheets create "Budget 2026" --parent <folder-id> --json

# Get spreadsheet metadata
gdrv sheets get <spreadsheet-id>                # Get spreadsheet details
gdrv sheets get 1abc123... --json

# Read and write values
gdrv sheets values get <spreadsheet-id> <range> # Get values from a range
gdrv sheets values get 1abc123... "Sheet1!A1:B10" --json
gdrv sheets values update <spreadsheet-id> <range> # Update values in a range
gdrv sheets values update 1abc123... "Sheet1!A1:B2" --values '[[1,2],[3,4]]'
gdrv sheets values update 1abc123... "Sheet1!A1:B2" --values-file data.json
gdrv sheets values append <spreadsheet-id> <range> # Append values to a range
gdrv sheets values append 1abc123... "Sheet1!A1" --values '[[5,6]]' --value-input-option RAW
gdrv sheets values clear <spreadsheet-id> <range> # Clear values from a range

# Batch update spreadsheet
gdrv sheets batch-update <spreadsheet-id>       # Batch update spreadsheet
gdrv sheets batch-update 1abc123... --requests-file examples/sheets/batch-update.json
```

**Command Flags:**

- **List flags:** `--parent`, `--query`, `--limit`, `--page-token`, `--order-by`, `--fields`, `--paginate`
- **Create flags:** `--parent`
- **Batch-update flags:** `--requests`, `--requests-file`
- **Values update/append flags:** `--value-input-option` (RAW or USER_ENTERED), `--values`, `--values-file`

**Examples:**

```bash
# List all spreadsheets with pagination
gdrv sheets list --paginate --json

# Create a spreadsheet in a specific folder
gdrv sheets create "Q1 Report" --parent 0ABC123... --json

# Read a range of values
gdrv sheets values get 1abc123... "Sheet1!A1:C10" --json

# Update values from a JSON file
gdrv sheets values update 1abc123... "Sheet1!A1" \
  --values-file data.json \
  --value-input-option USER_ENTERED

# Batch update with multiple operations
gdrv sheets batch-update 1abc123... \
  --requests-file examples/sheets/batch-update.json
```

### Google Docs Operations

Manage Google Docs documents with content reading and batch update capabilities.

**Required OAuth Scopes:**
- Read operations: `https://www.googleapis.com/auth/documents.readonly` or `https://www.googleapis.com/auth/documents`
- Write operations: `https://www.googleapis.com/auth/documents`
- Use preset: `workspace-basic` (read-only) or `workspace-full` (read/write)

**API Documentation:** [Google Docs API v1](https://developers.google.com/docs/api)

```bash
# List documents
gdrv docs list                                # List all documents
gdrv docs list --parent <folder-id>          # List documents in a folder
gdrv docs list --query "name contains 'Report'" --json
gdrv docs list --paginate --json            # Get all documents

# Create a document
gdrv docs create "My Document"               # Create a new document
gdrv docs create "Meeting Notes" --parent <folder-id> --json

# Get document metadata
gdrv docs get <document-id>                   # Get document details
gdrv docs get 1abc123... --json

# Read document content
gdrv docs read <document-id>                  # Extract plain text from document
gdrv docs read 1abc123...                     # Print plain text
gdrv docs read 1abc123... --json             # Get structured content

# Batch update document
gdrv docs update <document-id>                # Batch update document
gdrv docs update 1abc123... --requests-file updates.json
gdrv docs update 1abc123... --requests-file examples/docs/batch-update.json
```

**Command Flags:**

- **List flags:** `--parent`, `--query`, `--limit`, `--page-token`, `--order-by`, `--fields`, `--paginate`
- **Create flags:** `--parent`
- **Update flags:** `--requests`, `--requests-file`

**Examples:**

```bash
# List all documents with pagination
gdrv docs list --paginate --json

# Create a document in a specific folder
gdrv docs create "Project Plan" --parent 0ABC123... --json

# Read document content as plain text
gdrv docs read 1abc123...

# Read document content as structured JSON
gdrv docs read 1abc123... --json

# Update document with batch requests
gdrv docs update 1abc123... \
  --requests-file examples/docs/batch-update.json
```

### Google Slides Operations

Manage Google Slides presentations with content reading, batch updates, and text replacement (templating) capabilities.

**Required OAuth Scopes:**
- Read operations: `https://www.googleapis.com/auth/presentations.readonly` or `https://www.googleapis.com/auth/presentations`
- Write operations: `https://www.googleapis.com/auth/presentations`
- Use preset: `workspace-basic` (read-only) or `workspace-full` (read/write)

**API Documentation:** [Google Slides API v1](https://developers.google.com/slides/api)

```bash
# List presentations
gdrv slides list                                # List all presentations
gdrv slides list --parent <folder-id>          # List presentations in a folder
gdrv slides list --query "name contains 'Deck'" --json
gdrv slides list --paginate --json            # Get all presentations

# Create a presentation
gdrv slides create "My Presentation"           # Create a new presentation
gdrv slides create "Q1 Review" --parent <folder-id> --json

# Get presentation metadata
gdrv slides get <presentation-id>               # Get presentation details
gdrv slides get 1abc123... --json

# Read presentation content
gdrv slides read <presentation-id>              # Extract text from all slides
gdrv slides read 1abc123...                     # Print text from all slides
gdrv slides read 1abc123... --json             # Get structured content

# Batch update presentation
gdrv slides update <presentation-id>            # Batch update presentation
gdrv slides update 1abc123... --requests-file updates.json
gdrv slides update 1abc123... --requests-file examples/slides/batch-update.json

# Replace text placeholders (templating)
gdrv slides replace <presentation-id>           # Replace text placeholders
gdrv slides replace 1abc123... --data '{"{{NAME}}":"Alice","{{DATE}}":"2026-01-24"}'
gdrv slides replace 1abc123... --file examples/slides/replacements.json
```

**Command Flags:**

- **List flags:** `--parent`, `--query`, `--limit`, `--page-token`, `--order-by`, `--fields`, `--paginate`
- **Create flags:** `--parent`
- **Update flags:** `--requests`, `--requests-file`
- **Replace flags:** `--data` (JSON string), `--file` (JSON file path)

**Examples:**

```bash
# List all presentations with pagination
gdrv slides list --paginate --json

# Create a presentation in a specific folder
gdrv slides create "Team Meeting" --parent 0ABC123... --json

# Read text from all slides
gdrv slides read 1abc123...

# Read structured presentation content
gdrv slides read 1abc123... --json

# Replace placeholders using inline JSON
gdrv slides replace 1abc123... \
  --data '{"{{NAME}}":"Alice","{{DATE}}":"2026-01-24","{{TITLE}}":"Manager"}'

# Replace placeholders using a JSON file
gdrv slides replace 1abc123... \
  --file examples/slides/replacements.json

# Batch update with multiple operations
gdrv slides update 1abc123... \
  --requests-file examples/slides/batch-update.json
```

### Shared Drives
```bash
gdrv drives list                 # List Shared Drives
gdrv drives get <drive-id>       # Get drive details
```

### Admin SDK Operations

Manage Google Workspace users and groups through the Admin SDK Directory API.

**⚠️ Important Authentication Requirements:**

Admin SDK operations **require service account authentication** with domain-wide delegation enabled. This is different from regular OAuth authentication.

**Required Setup:**
1. Create a service account in Google Cloud Console
2. Enable domain-wide delegation for the service account
3. Authorize the required scopes in Google Workspace Admin Console
4. Download the service account JSON key file
5. Authenticate using the service account with user impersonation

**Required OAuth Scopes:**
- User management: `https://www.googleapis.com/auth/admin.directory.user`
- Group management: `https://www.googleapis.com/auth/admin.directory.group`
- Use preset: `admin` or `workspace-with-admin`

**API Documentation:** 
- [Admin SDK Directory API - Users](https://developers.google.com/admin-sdk/directory/reference/rest/v1/users)
- [Admin SDK Directory API - Groups](https://developers.google.com/admin-sdk/directory/reference/rest/v1/groups)

**Authentication Example:**

```bash
# Authenticate with service account and impersonate admin user
gdrv auth service-account \
  --key-file ./service-account-key.json \
  --impersonate-user admin@example.com \
  --preset admin
```

#### User Management

```bash
# List users
gdrv admin users list --domain example.com
gdrv admin users list --domain example.com --json
gdrv admin users list --domain example.com --paginate --json
gdrv admin users list --domain example.com --query "name:John" --json

# Get user details
gdrv admin users get user@example.com
gdrv admin users get user@example.com --fields "id,name,email" --json

# Create a user
gdrv admin users create newuser@example.com \
  --given-name "John" \
  --family-name "Doe" \
  --password "TempPass123!"

# Update a user
gdrv admin users update user@example.com --given-name "Jane" --family-name "Smith"
gdrv admin users update user@example.com --suspended true
gdrv admin users update user@example.com --org-unit-path "/Engineering/Developers"

# Suspend/unsuspend a user
gdrv admin users suspend user@example.com
gdrv admin users unsuspend user@example.com

# Delete a user
gdrv admin users delete user@example.com
```

**User Command Flags:**

- **List flags:** `--domain` or `--customer`, `--query`, `--limit`, `--page-token`, `--order-by`, `--fields`, `--paginate`
- **Get flags:** `--fields`
- **Create flags:** `--given-name` (required), `--family-name` (required), `--password` (required)
- **Update flags:** `--given-name`, `--family-name`, `--suspended` (true/false), `--org-unit-path`

#### Group Management

```bash
# List groups
gdrv admin groups list --domain example.com
gdrv admin groups list --domain example.com --json
gdrv admin groups list --domain example.com --paginate --json
gdrv admin groups list --domain example.com --query "name:Team" --json

# Get group details
gdrv admin groups get group@example.com
gdrv admin groups get group@example.com --fields "id,name,email" --json

# Create a group
gdrv admin groups create group@example.com "Team Group" \
  --description "Team access group"

# Update a group
gdrv admin groups update group@example.com --name "New Name"
gdrv admin groups update group@example.com --description "Updated description"

# Delete a group
gdrv admin groups delete group@example.com
```

**Group Command Flags:**

- **List flags:** `--domain` or `--customer`, `--query`, `--limit`, `--page-token`, `--order-by`, `--fields`, `--paginate`
- **Get flags:** `--fields`
- **Create flags:** `--description`
- **Update flags:** `--name`, `--description`

#### Group Membership Management

```bash
# List group members
gdrv admin groups members list team@example.com
gdrv admin groups members list team@example.com --json
gdrv admin groups members list team@example.com --roles MANAGER --json
gdrv admin groups members list team@example.com --paginate --json

# Add member to group
gdrv admin groups members add team@example.com user@example.com --role MEMBER
gdrv admin groups members add team@example.com user@example.com --role MANAGER
gdrv admin groups members add team@example.com user@example.com --role OWNER

# Remove member from group
gdrv admin groups members remove team@example.com user@example.com
```

**Group Members Command Flags:**

- **List flags:** `--limit`, `--page-token`, `--roles` (OWNER, MANAGER, MEMBER), `--fields`, `--paginate`
- **Add flags:** `--role` (OWNER, MANAGER, or MEMBER, default: MEMBER)

**Examples:**

```bash
# List all users in a domain
gdrv admin users list --domain example.com --paginate --json

# Create a new user
gdrv admin users create john.doe@example.com \
  --given-name "John" \
  --family-name "Doe" \
  --password "SecurePass123!" \
  --json

# Suspend a user
gdrv admin users suspend john.doe@example.com

# List all groups
gdrv admin groups list --domain example.com --paginate --json

# Create a group and add members
gdrv admin groups create team@example.com "Engineering Team" \
  --description "Engineering team members"
gdrv admin groups members add team@example.com john.doe@example.com --role MEMBER
```

### Drive Activity API Operations

Query and monitor file and folder activity across Google Drive with detailed activity streams.

**Required OAuth Scopes:**
- Read-only: `https://www.googleapis.com/auth/drive.activity.readonly`
- Full access: `https://www.googleapis.com/auth/drive.activity`
- Use preset: `workspace-activity` (read-only) or `workspace-complete`

**API Documentation:** [Drive Activity API v2](https://developers.google.com/drive/activity/v2)

```bash
# Query recent activity for all accessible files
gdrv activity query --json

# Query activity for a specific file
gdrv activity query --file-id 1abc123... --json

# Query activity within a time range
gdrv activity query --start-time "2026-01-01T00:00:00Z" --end-time "2026-01-31T23:59:59Z" --json

# Query activity for a folder (including descendants)
gdrv activity query --folder-id 0ABC123... --json

# Filter by activity types
gdrv activity query --action-types "edit,share,permission_change" --json

# Get activity for a specific user
gdrv activity query --user user@example.com --json

# Paginate through activity results
gdrv activity query --limit 100 --page-token "TOKEN" --json
```

### Drive Labels API Operations

Manage Drive labels and apply structured metadata to files and folders.

**Required OAuth Scopes:**
- Read-only: `https://www.googleapis.com/auth/drive.labels.readonly`
- Full access: `https://www.googleapis.com/auth/drive.labels`
- Admin: `https://www.googleapis.com/auth/drive.admin.labels`
- Use preset: `workspace-labels` or `workspace-complete`

**API Documentation:** [Drive Labels API v2](https://developers.google.com/drive/labels/overview)

```bash
# List available labels
gdrv labels list --json

# Get label schema
gdrv labels get <label-id> --json

# List labels applied to a file
gdrv labels file list <file-id> --json

# Apply a label to a file
gdrv labels file apply <file-id> <label-id> \
  --fields '{"field1":"value1","field2":"value2"}' --json

# Update label fields on a file
gdrv labels file update <file-id> <label-id> \
  --fields '{"field1":"new_value"}' --json

# Remove a label from a file
gdrv labels file remove <file-id> <label-id>

# Create a label (admin only)
gdrv labels create "Document Type" --json

# Publish a label (admin only)
gdrv labels publish <label-id>

# Disable a label (admin only)
gdrv labels disable <label-id>
```

### Drive Changes API Operations

Track changes to files and folders for real-time synchronization and automation.

**Required OAuth Scopes:**
- Uses standard Drive scopes (no additional scopes needed)
- Use preset: `workspace-sync` or `workspace-complete`

**API Documentation:** [Drive Changes API v3](https://developers.google.com/drive/api/v3/reference/changes)

```bash
# Get the starting page token
gdrv changes start-page-token --json

# List changes since a page token
gdrv changes list --page-token "12345" --json

# List changes with auto-pagination
gdrv changes list --page-token "12345" --paginate --json

# List changes for a specific Shared Drive
gdrv changes list --page-token "12345" --drive-id <drive-id> --json

# Watch for changes (webhook setup)
gdrv changes watch --page-token "12345" \
  --webhook-url "https://example.com/webhook" --json

# Stop watching for changes
gdrv changes stop <channel-id> <resource-id>

# List changes including removed files
gdrv changes list --page-token "12345" --include-removed --json
```

### Permission Auditing and Analysis

Enhanced permission auditing and access analysis tools for security and compliance.

**Required OAuth Scopes:**
- Uses standard Drive scopes (no additional scopes needed)

```bash
# Audit all files with public access
gdrv permissions audit public --json

# Audit all files shared with external domains
gdrv permissions audit external --internal-domain example.com --json

# Audit permissions for a specific user
gdrv permissions audit user user@example.com --json

# Find files with "anyone with link" access
gdrv permissions audit anyone-with-link --json

# Analyze permission inheritance for a folder
gdrv permissions analyze <folder-id> --recursive --json

# Generate permission report for a file/folder
gdrv permissions report <file-id> --internal-domain example.com --json

# Bulk remove public access (dry-run first)
gdrv permissions bulk remove-public --folder-id <folder-id> --dry-run --json

# Bulk change role (e.g., downgrade all "writer" to "reader")
gdrv permissions bulk update-role --folder-id <folder-id> \
  --from-role writer --to-role reader --dry-run --json

# Find files accessible by a specific email
gdrv permissions search --email user@example.com --json

# List all files with "commenter" access
gdrv permissions search --role commenter --json
```

### Configuration
```bash
gdrv config show                 # Show current config
gdrv config set <key> <value>    # Set config value
gdrv config reset                # Reset to defaults
```

Config file defaults:
- macOS: `~/Library/Application Support/gdrv/config.json`
- Linux: `~/.config/gdrv/config.json`
- Windows: `%APPDATA%\\gdrv\\config.json`
- Override with `GDRV_CONFIG_DIR`

OAuth client fields in config (optional):
- `oauthClientId`
- `oauthClientSecret` (only if required by your client type)

### Other
```bash
gdrv auth login [--preset <preset>] [--wide] [--scopes <scopes>] [--no-browser] [--client-id <id>] [--client-secret <secret>] [--profile <name>]
gdrv auth device [--preset <preset>] [--wide] [--client-id <id>] [--client-secret <secret>] [--profile <name>]
gdrv auth service-account --key-file <file> [--preset <preset>] [--scopes <scopes>] [--impersonate-user <email>] [--profile <name>]
gdrv auth status                 # Show auth status
gdrv auth profiles               # Manage profiles
gdrv auth logout                 # Clear credentials
gdrv about                       # Show API capabilities
```

## Output Formats

### Table Format (Default)
```bash
gdrv files list
```

### JSON Format
```bash
gdrv files list --json
```

### Quiet Mode
```bash
gdrv files upload file.txt --quiet
```

## Safety Controls

### Dry Run (Preview)
Preview what would happen without executing:
```bash
gdrv files delete 123 --dry-run
```

### Default Behavior (Non-Interactive)
By default, commands execute without prompts for agent-friendliness:
```bash
# Executes immediately (no confirmation prompt)
gdrv files delete 123
```

### Interactive Mode
For interactive use, the CLI will prompt for confirmation when safety checks require it.

## Configuration

Configure behavior via config commands or environment variables:

```bash
# Set default output format
gdrv config set output_format json

# Set cache TTL
gdrv config set cache_ttl 300

# OAuth credentials
export GDRV_CLIENT_ID="your-client-id"
export GDRV_CLIENT_SECRET="your-client-secret" # only if required by your client type

# Environment variables
export GDRV_PROFILE=work
export GDRV_CONFIG_DIR=/path/to/config
export GDRV_REQUIRE_CUSTOM_OAUTH=1
```

## Troubleshooting

### Authentication Issues

**"OAuth client credentials missing"**
```bash
# Use bundled client (release builds) or set a custom client:
export GDRV_CLIENT_ID="your-client-id"
export GDRV_CLIENT_SECRET="your-client-secret" # only if required by your client type
gdrv auth login
```

**"Browser not opening"**
Use manual fallback or device code flow:
```bash
gdrv auth login --no-browser
gdrv auth device
```

**"Invalid credentials"**
Re-authenticate:
```bash
gdrv auth logout
gdrv auth login
```

**"Missing required scope"**
```bash
gdrv auth status
gdrv auth login --preset workspace-full
```

### Permission Errors

**"Insufficient permissions"**
Check your OAuth scopes and Shared Drive access:
```bash
gdrv auth status
```

### Path Resolution

**"File not found"**
Use file IDs for Shared Drives:
```bash
gdrv files list --drive-id <drive-id>
```

### Performance Issues

**Slow uploads/downloads**
Check your network and use resumable uploads for large files.

**Rate limiting**
The CLI automatically handles rate limits with exponential backoff.

## Development

### Running Tests
```bash
# Unit tests
go test ./...

# Integration tests (requires credentials)
go test -tags=integration ./test/integration/...
```

### Building
```bash
go build -o gdrv cmd/gdrv/main.go
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make changes with tests
4. Run `go test ./...`
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
