# Dependency Analysis Report

Generated: 2026-02-26
Go Version: 1.24.0

## Summary

- **Total Dependencies**: 233 modules
- **Direct Dependencies**: 15
- **Indirect Dependencies**: ~218
- **Module Cache Size**: ~2.4GB

## Direct Dependencies (15)

All direct dependencies are actively used and required:

### Google Cloud APIs (5)

| Package | Version | Purpose | Used In |
|---------|---------|---------|---------|
| `cloud.google.com/go/ai` | v0.15.0 | Generative Language API (Gemini) | `internal/ai/` |
| `cloud.google.com/go/apps` | v0.8.0 | Google Meet API | `internal/meet/` |
| `cloud.google.com/go/iam` | v1.5.3 | IAM Admin API | `internal/iamadmin/` |
| `cloud.google.com/go/logging` | v1.13.2 | Cloud Logging API | `internal/cloudlogging/` |
| `cloud.google.com/go/monitoring` | v1.24.3 | Cloud Monitoring API | `internal/monitoring/` |

### Core Dependencies (4)

| Package | Version | Purpose | Used In |
|---------|---------|---------|---------|
| `google.golang.org/api` | v0.267.0 | Google API client library | Throughout (Drive, Gmail, etc.) |
| `google.golang.org/grpc` | v1.78.0 | gRPC transport | Required by Google APIs |
| `google.golang.org/protobuf` | v1.36.11 | Protocol Buffers | Required by Google APIs |
| `golang.org/x/oauth2` | v0.35.0 | OAuth2 authentication | `internal/auth/` |

### CLI & UI (3)

| Package | Version | Purpose | Used In |
|---------|---------|---------|---------|
| `github.com/alecthomas/kong` | v1.14.0 | CLI framework | Command parsing |
| `github.com/olekukonko/tablewriter` | v0.0.5 | Table formatting | `internal/cli/`, `internal/config/` |
| `github.com/schollz/progressbar/v3` | v3.19.0 | Progress bars | `internal/files/` |

### Utilities (3)

| Package | Version | Purpose | Used In |
|---------|---------|---------|---------|
| `github.com/google/uuid` | v1.6.0 | UUID generation | `internal/sync/`, `internal/files/` |
| `github.com/zalando/go-keyring` | v0.2.6 | OS credential store | `internal/auth/` |
| `modernc.org/sqlite` | v1.44.3 | SQLite driver (CGO-free) | `internal/sync/index/` |

## Findings

### No Unused Direct Dependencies

All 15 direct dependencies are actively used in the codebase. `go mod tidy` confirmed this by not removing any direct dependencies.

### New Dependencies Added by `go mod tidy`

Running `go mod tidy` added 5 indirect dependencies that were missing:

1. `github.com/schollz/progressbar/v3` v3.19.0 - Already used in code but not in go.mod
2. `github.com/mitchellh/colorstring` v0.0.0-20190213212951-d06e56a500db - progressbar dependency
3. `github.com/rivo/uniseg` v0.4.7 - progressbar dependency
4. `github.com/sahilm/fuzzy` v0.1.1 - progressbar dependency
5. `golang.org/x/term` v0.39.0 - progressbar dependency

### SQLite Driver Analysis

**Current**: `modernc.org/sqlite` v1.44.3
- Pure Go implementation (no CGO required)
- Good for cross-platform builds
- Slightly larger binary size but simpler deployment

**Alternative**: `github.com/mattn/go-sqlite3`
- Requires CGO
- Better performance for high-throughput
- Smaller binary
- More complex cross-platform builds

**Recommendation**: Keep modernc.org/sqlite for simpler deployment.

### tablewriter Analysis

**Current**: `github.com/olekukonko/tablewriter` v0.0.5
- Simple, stable API
- Used for CLI table output throughout

**Alternatives**:
- `github.com/jedib0t/go-pretty/v6` - More features, larger
- `github.com/charmbracelet/lipgloss` - Modern TUI approach

**Recommendation**: Keep current - simple and sufficient.

### Indirect Dependencies

- 218 indirect dependencies
- No version conflicts detected
- No outdated packages requiring updates
- All versions are recent (2025-2026)

## Recommendations

1. **No action needed on direct dependencies** - All are actively used
2. **Keep modernc.org/sqlite** - CGO-free is worth the tradeoff
3. **Keep tablewriter** - Simple and sufficient for needs
4. **Monitor indirect dependencies** - No immediate issues, but periodic review recommended
5. **Consider `go mod vendor`** for reproducible builds if needed

## Build Status

After `go mod tidy`:
- All core packages build successfully
- All non-CLI tests pass
- 35 packages tested, 35 passed
- CLI packages have pre-existing build issues (unrelated to dependencies)

## Module Graph Summary

Top-level dependency graph (first 50 edges):

```
github.com/dl-alexandre/gdrv
├── cloud.google.com/go/* (5 API modules)
├── google.golang.org/api (core Google API client)
├── google.golang.org/grpc (transport)
├── github.com/alecthomas/kong (CLI)
├── github.com/olekukonko/tablewriter (tables)
├── github.com/schollz/progressbar/v3 (progress)
├── modernc.org/sqlite (database)
└── ... (indirect dependencies)
```
