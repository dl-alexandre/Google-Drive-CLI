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
- **Authentication**: OAuth2 with device code fallback, multiple profiles, secure credential storage
- **Shared Drives Support**: Full support for Google Workspace Shared Drives
- **Advanced Safety Controls**: Dry-run mode, confirmation prompts, idempotent operations
- **Rich CLI Interface**: 25+ commands with help, examples, and multiple output formats (JSON, table)
- **Production Logging**: Structured logging with debug mode and trace correlation
- **Cross-Platform**: Works on macOS, Linux, and Windows

## Installation

### Install Script (Recommended)

```bash
# Install to ~/.local/bin (ensure it's on your PATH)
curl -fsSL https://raw.githubusercontent.com/dl-alexandre/Google-Drive-CLI/master/install.sh | bash
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

1. **Authenticate**:
   ```bash
   gdrive auth login --wide
   ```

2. **List files**:
   ```bash
   gdrive files list
   ```

3. **Upload a file**:
   ```bash
   gdrive files upload myfile.txt
   ```

4. **Download a file**:
   ```bash
   gdrive files download 1abc123... --output downloaded.txt
   ```

5. **Download a Google Doc as text**:
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

The CLI supports multiple authentication methods:

### OAuth2 Flow (Recommended)
```bash
gdrive auth login --wide
```
Opens a browser for authentication.

### Device Code Flow (Headless)
```bash
gdrive auth device
```
Displays a code to enter at https://www.google.com/device.

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

### Shared Drives
```bash
gdrive drives list                 # List Shared Drives
gdrive drives get <drive-id>       # Get drive details
```

### Configuration
```bash
gdrive config show                 # Show current config
gdrive config set <key> <value>    # Set config value
gdrive config reset                # Reset to defaults
```

### Other
```bash
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

# Environment variables
export GDRIVE_PROFILE=work
export GDRIVE_CONFIG_DIR=/path/to/config
```

## Troubleshooting

### Authentication Issues

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
