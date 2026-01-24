# Google Drive CLI

A **fast**, **lightweight**, and **AI-agent friendly** CLI for Google Drive. Manage files with zero friction.

## Why gdrive?

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

## Installation

### Install Script (Recommended)

```bash
# Install to ~/.local/bin (ensure it's on your PATH)
curl -fsSL https://raw.githubusercontent.com/dl-alexandre/Google-Drive-CLI/master/install.sh | bash
```

### Homebrew (Tap)

```bash
brew tap dl-alexandre/tap
brew install gdrive
```

### Download Binary

Download the latest release from the [releases page](https://github.com/dl-alexandre/Google-Drive-CLI/releases).

```bash
# Make executable and move to PATH
chmod +x gdrive
sudo mv gdrive /usr/local/bin/
```

### Build from Source

```bash
git clone https://github.com/dl-alexandre/Google-Drive-CLI.git
cd gdrive
go build -o gdrive ./cmd/gdrive
```

## Quick Start

1. **Set OAuth credentials**:
   ```bash
   export GDRIVE_CLIENT_ID="your-client-id"
   export GDRIVE_CLIENT_SECRET="your-client-secret"
   ```

2. **Authenticate**:
   ```bash
   gdrive auth login --preset workspace-basic
   ```

3. **List files**:
   ```bash
   gdrive files list
   ```

4. **Upload a file**:
   ```bash
   gdrive files upload myfile.txt
   ```

5. **Download a file**:
   ```bash
   gdrive files download 1abc123... --output downloaded.txt
   ```

6. **Download a Google Doc as text**:
   ```bash
   gdrive files download 1abc123... --doc
   ```

## Agent Quickstart

This CLI is designed to be used by AI agents and automation scripts. Key features for agent usage:

### JSON Output

Always use `--json` for machine-readable output:

```bash
# List files as JSON
gdrive files list --json

# Get file metadata
gdrive files get 1abc123... --json

# Upload returns the created file object
gdrive files upload report.pdf --json
```

### Pagination Control

Use `--paginate` to automatically fetch all pages:

```bash
# Get ALL files (auto-pagination)
gdrive files list --paginate --json

# Get all trashed files
gdrive files list-trashed --paginate --json

# Get all Shared Drives
gdrive drives list --paginate --json
```

Or control pagination manually:

```bash
# Get first page
gdrive files list --limit 100 --json

# Use nextPageToken from response for next page
gdrive files list --limit 100 --page-token "TOKEN_FROM_PREVIOUS" --json
```

### Sorting and Filtering

```bash
# Sort by modified time (newest first)
gdrive files list --order-by "modifiedTime desc" --json

# Search by name
gdrive files list --query "name contains 'report'" --json

# Combined: recent PDFs
gdrive files list --query "mimeType = 'application/pdf'" --order-by "modifiedTime desc" --json
```

### Non-Interactive Mode

Destructive commands run without prompts by default. Use `--dry-run` to preview:

```bash
# Preview what would be deleted
gdrive files delete 1abc123... --dry-run

# Actually delete (no prompt)
gdrive files delete 1abc123...

# Permanently delete (bypasses trash)
gdrive files delete 1abc123... --permanent
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

The CLI supports multiple authentication methods and scope presets. OAuth client credentials are required.

### Prerequisites

1. Create a project in Google Cloud Console
2. Enable the Google Drive API
3. Create OAuth 2.0 credentials (Desktop application)
4. Set credentials via environment variables or command flags:

```bash
export GDRIVE_CLIENT_ID="your-client-id"
export GDRIVE_CLIENT_SECRET="your-client-secret"

gdrive auth login --client-id "your-client-id" --client-secret "your-client-secret"
```

### OAuth2 Flow (Recommended)
```bash
gdrive auth login
```
Opens a browser for authentication.

### Device Code Flow (Headless)
```bash
gdrive auth device
```
Displays a code to enter at https://www.google.com/device.

### Service Account Authentication
```bash
gdrive auth service-account --key-file ./service-account-key.json --preset workspace-basic
```
Loads credentials from a service account JSON key file. Use `--impersonate-user` for Admin SDK scopes.

### Scope Presets

| Preset | Description | Use Case |
|--------|-------------|----------|
| `workspace-basic` | Read-only Drive, Sheets, Docs, Slides | Viewing and downloading |
| `workspace-full` | Full Drive, Sheets, Docs, Slides | Editing and management |
| `admin` | Admin Directory users and groups | User/group administration |
| `workspace-with-admin` | Workspace full + Admin Directory | Full workspace + admin |

```bash
gdrive auth login --preset workspace-basic
gdrive auth login --preset workspace-full
gdrive auth login --preset admin
gdrive auth login --preset workspace-with-admin
gdrive auth device --preset workspace-basic
gdrive auth service-account --key-file ./key.json --preset workspace-basic
```

### Custom Scopes

```bash
gdrive auth login --scopes "https://www.googleapis.com/auth/drive.file,https://www.googleapis.com/auth/spreadsheets.readonly"
gdrive auth service-account --key-file ./key.json --scopes "https://www.googleapis.com/auth/drive.file"
```

Available scopes:
- `https://www.googleapis.com/auth/drive`
- `https://www.googleapis.com/auth/drive.file`
- `https://www.googleapis.com/auth/drive.readonly`
- `https://www.googleapis.com/auth/drive.metadata.readonly`
- `https://www.googleapis.com/auth/spreadsheets`
- `https://www.googleapis.com/auth/spreadsheets.readonly`
- `https://www.googleapis.com/auth/documents`
- `https://www.googleapis.com/auth/documents.readonly`
- `https://www.googleapis.com/auth/presentations`
- `https://www.googleapis.com/auth/presentations.readonly`
- `https://www.googleapis.com/auth/admin.directory.user`
- `https://www.googleapis.com/auth/admin.directory.user.readonly`
- `https://www.googleapis.com/auth/admin.directory.group`
- `https://www.googleapis.com/auth/admin.directory.group.readonly`

### Multiple Profiles
```bash
# Create and switch profiles
gdrive auth login --profile work
gdrive auth login --profile personal

# Use specific profile
gdrive --profile work files list
```

## Commands

### File Operations
```bash
gdrive files upload <file>          # Upload file
gdrive files download <file-id>     # Download file
gdrive files list                   # List files
gdrive files delete <file-id>       # Delete file
gdrive files trash <file-id>        # Move to trash
gdrive files restore <file-id>      # Restore from trash
gdrive files revisions <file-id>    # List revisions
```

### Folder Operations
```bash
gdrive folders create <name>        # Create folder
gdrive folders list <folder-id>     # List contents
gdrive folders delete <folder-id>   # Delete folder
gdrive folders move <id> <parent>   # Move folder
```

### Permission Management
```bash
gdrive permissions list <file-id>           # List permissions
gdrive permissions create <file-id> --type user --email user@example.com --role reader
gdrive permissions update <file-id> <perm-id> --role writer
gdrive permissions delete <file-id> <perm-id>
gdrive permissions public <file-id>         # Create public link
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
gdrive sheets list                                # List all spreadsheets
gdrive sheets list --parent <folder-id>          # List spreadsheets in a folder
gdrive sheets list --query "name contains 'Report'" --json
gdrive sheets list --paginate --json            # Get all spreadsheets

# Create a spreadsheet
gdrive sheets create "My Spreadsheet"           # Create a new spreadsheet
gdrive sheets create "Budget 2026" --parent <folder-id> --json

# Get spreadsheet metadata
gdrive sheets get <spreadsheet-id>                # Get spreadsheet details
gdrive sheets get 1abc123... --json

# Read and write values
gdrive sheets values get <spreadsheet-id> <range> # Get values from a range
gdrive sheets values get 1abc123... "Sheet1!A1:B10" --json
gdrive sheets values update <spreadsheet-id> <range> # Update values in a range
gdrive sheets values update 1abc123... "Sheet1!A1:B2" --values '[[1,2],[3,4]]'
gdrive sheets values update 1abc123... "Sheet1!A1:B2" --values-file data.json
gdrive sheets values append <spreadsheet-id> <range> # Append values to a range
gdrive sheets values append 1abc123... "Sheet1!A1" --values '[[5,6]]' --value-input-option RAW
gdrive sheets values clear <spreadsheet-id> <range> # Clear values from a range

# Batch update spreadsheet
gdrive sheets batch-update <spreadsheet-id>       # Batch update spreadsheet
gdrive sheets batch-update 1abc123... --requests-file examples/sheets/batch-update.json
```

**Command Flags:**

- **List flags:** `--parent`, `--query`, `--limit`, `--page-token`, `--order-by`, `--fields`, `--paginate`
- **Create flags:** `--parent`
- **Batch-update flags:** `--requests`, `--requests-file`
- **Values update/append flags:** `--value-input-option` (RAW or USER_ENTERED), `--values`, `--values-file`

**Examples:**

```bash
# List all spreadsheets with pagination
gdrive sheets list --paginate --json

# Create a spreadsheet in a specific folder
gdrive sheets create "Q1 Report" --parent 0ABC123... --json

# Read a range of values
gdrive sheets values get 1abc123... "Sheet1!A1:C10" --json

# Update values from a JSON file
gdrive sheets values update 1abc123... "Sheet1!A1" \
  --values-file data.json \
  --value-input-option USER_ENTERED

# Batch update with multiple operations
gdrive sheets batch-update 1abc123... \
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
gdrive docs list                                # List all documents
gdrive docs list --parent <folder-id>          # List documents in a folder
gdrive docs list --query "name contains 'Report'" --json
gdrive docs list --paginate --json            # Get all documents

# Create a document
gdrive docs create "My Document"               # Create a new document
gdrive docs create "Meeting Notes" --parent <folder-id> --json

# Get document metadata
gdrive docs get <document-id>                   # Get document details
gdrive docs get 1abc123... --json

# Read document content
gdrive docs read <document-id>                  # Extract plain text from document
gdrive docs read 1abc123...                     # Print plain text
gdrive docs read 1abc123... --json             # Get structured content

# Batch update document
gdrive docs update <document-id>                # Batch update document
gdrive docs update 1abc123... --requests-file updates.json
gdrive docs update 1abc123... --requests-file examples/docs/batch-update.json
```

**Command Flags:**

- **List flags:** `--parent`, `--query`, `--limit`, `--page-token`, `--order-by`, `--fields`, `--paginate`
- **Create flags:** `--parent`
- **Update flags:** `--requests`, `--requests-file`

**Examples:**

```bash
# List all documents with pagination
gdrive docs list --paginate --json

# Create a document in a specific folder
gdrive docs create "Project Plan" --parent 0ABC123... --json

# Read document content as plain text
gdrive docs read 1abc123...

# Read document content as structured JSON
gdrive docs read 1abc123... --json

# Update document with batch requests
gdrive docs update 1abc123... \
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
gdrive slides list                                # List all presentations
gdrive slides list --parent <folder-id>          # List presentations in a folder
gdrive slides list --query "name contains 'Deck'" --json
gdrive slides list --paginate --json            # Get all presentations

# Create a presentation
gdrive slides create "My Presentation"           # Create a new presentation
gdrive slides create "Q1 Review" --parent <folder-id> --json

# Get presentation metadata
gdrive slides get <presentation-id>               # Get presentation details
gdrive slides get 1abc123... --json

# Read presentation content
gdrive slides read <presentation-id>              # Extract text from all slides
gdrive slides read 1abc123...                     # Print text from all slides
gdrive slides read 1abc123... --json             # Get structured content

# Batch update presentation
gdrive slides update <presentation-id>            # Batch update presentation
gdrive slides update 1abc123... --requests-file updates.json
gdrive slides update 1abc123... --requests-file examples/slides/batch-update.json

# Replace text placeholders (templating)
gdrive slides replace <presentation-id>           # Replace text placeholders
gdrive slides replace 1abc123... --data '{"{{NAME}}":"Alice","{{DATE}}":"2026-01-24"}'
gdrive slides replace 1abc123... --file examples/slides/replacements.json
```

**Command Flags:**

- **List flags:** `--parent`, `--query`, `--limit`, `--page-token`, `--order-by`, `--fields`, `--paginate`
- **Create flags:** `--parent`
- **Update flags:** `--requests`, `--requests-file`
- **Replace flags:** `--data` (JSON string), `--file` (JSON file path)

**Examples:**

```bash
# List all presentations with pagination
gdrive slides list --paginate --json

# Create a presentation in a specific folder
gdrive slides create "Team Meeting" --parent 0ABC123... --json

# Read text from all slides
gdrive slides read 1abc123...

# Read structured presentation content
gdrive slides read 1abc123... --json

# Replace placeholders using inline JSON
gdrive slides replace 1abc123... \
  --data '{"{{NAME}}":"Alice","{{DATE}}":"2026-01-24","{{TITLE}}":"Manager"}'

# Replace placeholders using a JSON file
gdrive slides replace 1abc123... \
  --file examples/slides/replacements.json

# Batch update with multiple operations
gdrive slides update 1abc123... \
  --requests-file examples/slides/batch-update.json
```

### Shared Drives
```bash
gdrive drives list                 # List Shared Drives
gdrive drives get <drive-id>       # Get drive details
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
gdrive auth service-account \
  --key-file ./service-account-key.json \
  --impersonate-user admin@example.com \
  --preset admin
```

#### User Management

```bash
# List users
gdrive admin users list --domain example.com
gdrive admin users list --domain example.com --json
gdrive admin users list --domain example.com --paginate --json
gdrive admin users list --domain example.com --query "name:John" --json

# Get user details
gdrive admin users get user@example.com
gdrive admin users get user@example.com --fields "id,name,email" --json

# Create a user
gdrive admin users create newuser@example.com \
  --given-name "John" \
  --family-name "Doe" \
  --password "TempPass123!"

# Update a user
gdrive admin users update user@example.com --given-name "Jane" --family-name "Smith"
gdrive admin users update user@example.com --suspended true
gdrive admin users update user@example.com --org-unit-path "/Engineering/Developers"

# Suspend/unsuspend a user
gdrive admin users suspend user@example.com
gdrive admin users unsuspend user@example.com

# Delete a user
gdrive admin users delete user@example.com
```

**User Command Flags:**

- **List flags:** `--domain` or `--customer`, `--query`, `--limit`, `--page-token`, `--order-by`, `--fields`, `--paginate`
- **Get flags:** `--fields`
- **Create flags:** `--given-name` (required), `--family-name` (required), `--password` (required)
- **Update flags:** `--given-name`, `--family-name`, `--suspended` (true/false), `--org-unit-path`

#### Group Management

```bash
# List groups
gdrive admin groups list --domain example.com
gdrive admin groups list --domain example.com --json
gdrive admin groups list --domain example.com --paginate --json
gdrive admin groups list --domain example.com --query "name:Team" --json

# Get group details
gdrive admin groups get group@example.com
gdrive admin groups get group@example.com --fields "id,name,email" --json

# Create a group
gdrive admin groups create group@example.com "Team Group" \
  --description "Team access group"

# Update a group
gdrive admin groups update group@example.com --name "New Name"
gdrive admin groups update group@example.com --description "Updated description"

# Delete a group
gdrive admin groups delete group@example.com
```

**Group Command Flags:**

- **List flags:** `--domain` or `--customer`, `--query`, `--limit`, `--page-token`, `--order-by`, `--fields`, `--paginate`
- **Get flags:** `--fields`
- **Create flags:** `--description`
- **Update flags:** `--name`, `--description`

#### Group Membership Management

```bash
# List group members
gdrive admin groups members list team@example.com
gdrive admin groups members list team@example.com --json
gdrive admin groups members list team@example.com --roles MANAGER --json
gdrive admin groups members list team@example.com --paginate --json

# Add member to group
gdrive admin groups members add team@example.com user@example.com --role MEMBER
gdrive admin groups members add team@example.com user@example.com --role MANAGER
gdrive admin groups members add team@example.com user@example.com --role OWNER

# Remove member from group
gdrive admin groups members remove team@example.com user@example.com
```

**Group Members Command Flags:**

- **List flags:** `--limit`, `--page-token`, `--roles` (OWNER, MANAGER, MEMBER), `--fields`, `--paginate`
- **Add flags:** `--role` (OWNER, MANAGER, or MEMBER, default: MEMBER)

**Examples:**

```bash
# List all users in a domain
gdrive admin users list --domain example.com --paginate --json

# Create a new user
gdrive admin users create john.doe@example.com \
  --given-name "John" \
  --family-name "Doe" \
  --password "SecurePass123!" \
  --json

# Suspend a user
gdrive admin users suspend john.doe@example.com

# List all groups
gdrive admin groups list --domain example.com --paginate --json

# Create a group and add members
gdrive admin groups create team@example.com "Engineering Team" \
  --description "Engineering team members"
gdrive admin groups members add team@example.com john.doe@example.com --role MEMBER
```

### Configuration
```bash
gdrive config show                 # Show current config
gdrive config set <key> <value>    # Set config value
gdrive config reset                # Reset to defaults
```

### Other
```bash
gdrive auth login [--preset <preset>] [--wide] [--scopes <scopes>] [--client-id <id>] [--client-secret <secret>] [--profile <name>]
gdrive auth device [--preset <preset>] [--wide] [--client-id <id>] [--client-secret <secret>] [--profile <name>]
gdrive auth service-account --key-file <file> [--preset <preset>] [--scopes <scopes>] [--impersonate-user <email>] [--profile <name>]
gdrive auth status                 # Show auth status
gdrive auth profiles               # Manage profiles
gdrive auth logout                 # Clear credentials
gdrive about                       # Show API capabilities
```

## Output Formats

### Table Format (Default)
```bash
gdrive files list
```

### JSON Format
```bash
gdrive files list --json
```

### Quiet Mode
```bash
gdrive files upload file.txt --quiet
```

## Safety Controls

### Dry Run (Preview)
Preview what would happen without executing:
```bash
gdrive files delete 123 --dry-run
```

### Default Behavior (Non-Interactive)
By default, commands execute without prompts for agent-friendliness:
```bash
# Executes immediately (no confirmation prompt)
gdrive files delete 123
```

### Interactive Mode
For interactive use, the CLI will prompt for confirmation when safety checks require it.

## Configuration

Configure behavior via config commands or environment variables:

```bash
# Set default output format
gdrive config set output_format json

# Set cache TTL
gdrive config set cache_ttl 300

# OAuth credentials
export GDRIVE_CLIENT_ID="your-client-id"
export GDRIVE_CLIENT_SECRET="your-client-secret"

# Environment variables
export GDRIVE_PROFILE=work
export GDRIVE_CONFIG_DIR=/path/to/config
```

## Troubleshooting

### Authentication Issues

**"OAuth client ID and secret required"**
```bash
export GDRIVE_CLIENT_ID="your-client-id"
export GDRIVE_CLIENT_SECRET="your-client-secret"
gdrive auth login
```

**"Browser not opening"**
Use device code flow:
```bash
gdrive auth device
```

**"Invalid credentials"**
Re-authenticate:
```bash
gdrive auth logout
gdrive auth login
```

**"Missing required scope"**
```bash
gdrive auth status
gdrive auth login --preset workspace-full
```

### Permission Errors

**"Insufficient permissions"**
Check your OAuth scopes and Shared Drive access:
```bash
gdrive auth status
```

### Path Resolution

**"File not found"**
Use file IDs for Shared Drives:
```bash
gdrive files list --drive-id <drive-id>
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
go build -o gdrive cmd/gdrive/main.go
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make changes with tests
4. Run `go test ./...`
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
