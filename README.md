# Google Drive CLI

[![Go Report Card](https://goreportcard.com/badge/github.com/dl-alexandre/gdrv)](https://goreportcard.com/report/github.com/dl-alexandre/gdrv)

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

- **Drive Activity API (v2)** - Audit file and folder activity
- **Drive Labels API (v2)** - Structured metadata and custom taxonomy
- **Drive Changes API (v3)** - Real-time change tracking and sync
- **Permission Auditing** - Security analysis and bulk permission management

See [docs/API-GUIDE.md](docs/API-GUIDE.md) for detailed API documentation.

## Quick Installation

```bash
# Install script (recommended)
curl -fsSL https://raw.githubusercontent.com/dl-alexandre/Google-Drive-CLI/master/install.sh | bash

# Homebrew
brew tap dl-alexandre/tap
brew install gdrv

# Build from source
git clone https://github.com/dl-alexandre/Google-Drive-CLI.git
cd Google-Drive-CLI
go build -o gdrv ./cmd/gdrv
```

See [docs/INSTALLATION.md](docs/INSTALLATION.md) for all installation methods.


## Shell Completion

Generate and install shell completion scripts for bash, zsh, fish, and PowerShell.

### Bash

```bash
# Load completion for current session
source <(gdrv completion bash)

# Install permanently (Linux)
gdrv completion bash > /etc/bash_completion.d/gdrv

# Install permanently (macOS with Homebrew bash-completion)
gdrv completion bash > /usr/local/etc/bash_completion.d/gdrv
```

### Zsh

```bash
# Load completion for current session
source <(gdrv completion zsh)

# Install permanently
mkdir -p ~/.config/zsh/completions
gdrv completion zsh > ~/.config/zsh/completions/_gdrv
# Add to ~/.zshrc:
echo 'fpath+=(~/.config/zsh/completions)' >> ~/.zshrc
```

### Fish

```bash
# Load completion for current session
gdrv completion fish | source

# Install permanently
gdrv completion fish > ~/.config/fish/completions/gdrv.fish
```

### PowerShell

```powershell
# Load completion for current session
gdrv completion powershell | Out-String | Invoke-Expression

# Install permanently (add to your PowerShell profile)
gdrv completion powershell > $PROFILE.CurrentUserCurrentHost
```

## Quick Start

1. **Authenticate**:
   ```bash
   gdrv auth login
   ```
   Default is full read/write access. Use `--preset workspace-basic` for read-only.

2. **List files**:
   ```bash
   gdrv files list
   ```

3. **Upload a file**:
   ```bash
   gdrv files upload myfile.txt
   ```

4. **Download a file**:
   ```bash
   gdrv files download 1abc123... --output downloaded.txt
   ```

5. **Download a Google Doc as text**:
   ```bash
   gdrv files download 1abc123... --doc
   ```

## Authentication Basics

```bash
# OAuth2 (opens browser)
gdrv auth login

# Device code (headless)
gdrv auth device

# Service account
gdrv auth service-account --key-file ./service-account.json

# Multiple profiles
gdrv auth login --profile work
gdrv --profile work files list
```

**Scope Presets:**
- `workspace-full` - Full read/write access (default)
- `workspace-basic` - Read-only access
- `admin` - Admin SDK for user/group management
- `workspace-complete` - All APIs including Activity, Labels, and Changes

See [docs/AUTHENTICATION.md](docs/AUTHENTICATION.md) for complete authentication documentation.

## AI Agent Quickstart

```bash
# Always use --json for machine-readable output
gdrv files list --json

# Auto-pagination to get all results
gdrv files list --paginate --json

# Preview destructive operations
gdrv files delete 1abc123... --dry-run

# Sort and filter
gdrv files list --query "mimeType = 'application/pdf'" --order-by "modifiedTime desc" --json
```

See [docs/AGENTS.md](docs/AGENTS.md) for complete AI agent best practices.

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | Authentication required |
| 3 | Invalid argument |
| 4 | Resource not found |
| 5 | Permission denied |
| 6 | Rate limited |

## Documentation

- **[Installation Guide](docs/INSTALLATION.md)** - All installation methods
- **[Authentication Guide](docs/AUTHENTICATION.md)** - OAuth flows, scope presets, service accounts
- **[API Guide](docs/API-GUIDE.md)** - Complete command reference
- **[AI Agent Best Practices](docs/AGENTS.md)** - JSON output, pagination, exit codes
- **[Troubleshooting](docs/TROUBLESHOOTING.md)** - Common issues and solutions
- **[Changelog](docs/CHANGELOG.md)** - Release history

## Configuration

Config file locations:
- macOS: `~/Library/Application Support/gdrv/config.json`
- Linux: `~/.config/gdrv/config.json`
- Windows: `%APPDATA%\gdrv\config.json`

Environment variables:
```bash
export GDRV_PROFILE=work
export GDRV_CONFIG_DIR=/path/to/config
export GDRV_REQUIRE_CUSTOM_OAUTH=1
```

## Privacy & Terms

- **[Privacy Policy](https://github.com/dl-alexandre/Google-Drive-CLI/blob/master/docs/PRIVACY.md)** - How we handle your data (we don't collect any)
- **[Terms of Service](https://github.com/dl-alexandre/Google-Drive-CLI/blob/master/docs/TERMS.md)** - License, warranty, and usage terms

**TL;DR:** gdrv stores credentials locally encrypted, contacts only Google APIs, collects zero data, and is fully open source.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make changes with tests
4. Run `go test ./...`
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
# CI trigger

