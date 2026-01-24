---
name: Foundation for Google APIs Integration
overview: Establish shared infrastructure for Sheets, Docs, Slides, and Admin SDK integrations including unified scope management, service factory pattern, error handling, and output formatting.
todos: []
isProject: true
status: completed
---

> **Status**: ✅ **COMPLETED** - This foundation has been successfully implemented.

## Context

All four planned API integrations (Sheets, Docs, Slides, Admin SDK) share common needs that should be abstracted before implementing individual features. This foundation will:

- Prevent code duplication across integrations
- Ensure consistent error handling and rate limiting
- Provide a unified authentication strategy
- Enable easier testing and maintenance

Current state shows OAuth config is set per-command but scopes are hardcoded in auth defaults:

```go
// internal/cli/auth.go
authLoginCmd.Flags().StringSliceVar(&authScopes, "scopes", []string{utils.ScopeFile}, "OAuth scopes to request")
```

## Plan

### 1. Unified Scope Management

**Problem**: Each API needs specific scopes, but all must coexist. Admin SDK requires domain-wide delegation.

**Files to modify**:

- `internal/utils/constants.go`: Add all scopes for Sheets, Docs, Slides, Admin SDK
- `internal/auth/manager.go`: Add scope validation and compatibility checking

**Implementation**:

```go
// Add to constants.go
const (
    // Sheets scopes
    ScopeSheets         = "https://www.googleapis.com/auth/spreadsheets"
    ScopeSheetsReadonly = "https://www.googleapis.com/auth/spreadsheets.readonly"

    // Docs scopes
    ScopeDocs         = "https://www.googleapis.com/auth/documents"
    ScopeDocsReadonly = "https://www.googleapis.com/auth/documents.readonly"

    // Slides scopes
    ScopeSlides         = "https://www.googleapis.com/auth/presentations"
    ScopeSlidesReadonly = "https://www.googleapis.com/auth/presentations.readonly"

    // Admin SDK scopes (requires domain-wide delegation)
    ScopeAdminDirectoryUser         = "https://www.googleapis.com/auth/admin.directory.user"
    ScopeAdminDirectoryUserReadonly = "https://www.googleapis.com/auth/admin.directory.user.readonly"
    ScopeAdminDirectoryGroup        = "https://www.googleapis.com/auth/admin.directory.group"
    ScopeAdminDirectoryGroupReadonly = "https://www.googleapis.com/auth/admin.directory.group.readonly"
)

// Scope presets for common use cases
var (
    ScopesWorkspaceBasic    = []string{ScopeFile, ScopeSheets, ScopeDocs, ScopeSlides}
    ScopesWorkspaceFull     = []string{ScopeFull, ScopeSheets, ScopeDocs, ScopeSlides}
    ScopesAdmin             = []string{ScopeAdminDirectoryUser, ScopeAdminDirectoryGroup}
    ScopesWorkspaceWithAdmin = append(ScopesWorkspaceFull, ScopesAdmin...)
)
```

Add scope validation in `auth.Manager`:

```go
func (m *Manager) ValidateScopes(scopes []string) error {
    // Check for Admin SDK scopes with OAuth flow (not supported)
    hasAdmin := false
    hasOAuth := m.oauthConfig != nil

    for _, scope := range scopes {
        if strings.Contains(scope, "admin.directory") {
            hasAdmin = true
            break
        }
    }

    if hasAdmin && hasOAuth {
        return fmt.Errorf("Admin SDK scopes require service account auth with domain-wide delegation")
    }
    return nil
}
```

### 2. Service Factory Pattern

**Problem**: Each API integration duplicates service creation logic.

**Files to create/modify**:

- `internal/auth/service_factory.go` (new file)
- `internal/auth/manager.go`: Add generic service creation

**Implementation**:

```go
// service_factory.go
type ServiceType string

const (
    ServiceDrive     ServiceType = "drive"
    ServiceSheets    ServiceType = "sheets"
    ServiceDocs      ServiceType = "docs"
    ServiceSlides    ServiceType = "slides"
    ServiceAdminDir  ServiceType = "admin_directory"
)

type ServiceFactory struct {
    manager *Manager
}

func (f *ServiceFactory) CreateService(ctx context.Context, creds *types.Credentials, svcType ServiceType) (interface{}, error) {
    client := f.manager.GetHTTPClient(ctx, creds)

    switch svcType {
    case ServiceDrive:
        return drive.NewService(ctx, option.WithHTTPClient(client))
    case ServiceSheets:
        return sheets.NewService(ctx, option.WithHTTPClient(client))
    case ServiceDocs:
        return docs.NewService(ctx, option.WithHTTPClient(client))
    case ServiceSlides:
        return slides.NewService(ctx, option.WithHTTPClient(client))
    case ServiceAdminDir:
        return admin.NewService(ctx, option.WithHTTPClient(client))
    default:
        return nil, fmt.Errorf("unknown service type: %s", svcType)
    }
}
```

### 3. Unified Error Handling

**Problem**: Google API errors need consistent handling across all services.

**Files to create**:

- `internal/errors/google_api.go` (new file)
- `internal/errors/translator.go` (new file)

**Implementation**:

```go
// google_api.go
type APIError struct {
    Code       int
    Message    string
    Service    string
    Method     string
    Retryable  bool
    RateLimit  bool
}

func (e *APIError) Error() string {
    return fmt.Sprintf("%s API error (%d): %s", e.Service, e.Code, e.Message)
}

func ParseGoogleAPIError(err error, service string) *APIError {
    if err == nil {
        return nil
    }

    // Check for googleapi.Error
    if gErr, ok := err.(*googleapi.Error); ok {
        return &APIError{
            Code:      gErr.Code,
            Message:   gErr.Message,
            Service:   service,
            Retryable: isRetryableCode(gErr.Code),
            RateLimit: gErr.Code == 429,
        }
    }

    return &APIError{
        Code:    500,
        Message: err.Error(),
        Service: service,
    }
}

func isRetryableCode(code int) bool {
    return code == 429 || code == 500 || code == 502 || code == 503 || code == 504
}
```

### 4. Rate Limiting & Retry Logic

**Problem**: Google APIs have quotas that need exponential backoff.

**Files to create**:

- `internal/retry/backoff.go` (new file)

**Implementation**:

```go
type RetryConfig struct {
    MaxRetries     int
    InitialDelay   time.Duration
    MaxDelay       time.Duration
    Multiplier     float64
}

func DefaultRetryConfig() *RetryConfig {
    return &RetryConfig{
        MaxRetries:   5,
        InitialDelay: 1 * time.Second,
        MaxDelay:     32 * time.Second,
        Multiplier:   2.0,
    }
}

func WithRetry(ctx context.Context, config *RetryConfig, fn func() error) error {
    var lastErr error
    delay := config.InitialDelay

    for attempt := 0; attempt <= config.MaxRetries; attempt++ {
        if attempt > 0 {
            select {
            case <-ctx.Done():
                return ctx.Err()
            case <-time.After(delay):
            }
        }

        err := fn()
        if err == nil {
            return nil
        }

        // Check if retryable
        apiErr := ParseGoogleAPIError(err, "")
        if apiErr != nil && !apiErr.Retryable {
            return err
        }

        lastErr = err
        delay = time.Duration(float64(delay) * config.Multiplier)
        if delay > config.MaxDelay {
            delay = config.MaxDelay
        }
    }

    return fmt.Errorf("max retries exceeded: %w", lastErr)
}
```

### 5. Enhanced Output Formatting

**Problem**: Each API needs custom table/JSON output, but logic is duplicated.

**Files to modify**:

- `internal/cli/output.go`: Add interface-based formatting
- `internal/types/output.go` (new file): Define output interfaces

**Implementation**:

```go
// types/output.go
type TableRenderer interface {
    RenderTable(w io.Writer) error
    Headers() []string
    Rows() [][]string
}

type OutputData interface {
    AsJSON() (interface{}, error)
    AsTable() (TableRenderer, error)
}

// Modify output.go
func (w *OutputWriter) writeTable(data interface{}) error {
    // Try interface first
    if renderer, ok := data.(TableRenderer); ok {
        return w.renderTableFromInterface(renderer)
    }

    // Fall back to type switch for legacy types
    switch v := data.(type) {
    case []*types.DriveFile:
        return w.writeFileTable(v)
    case []*types.Permission:
        return w.writePermissionTable(v)
    default:
        return w.writeJSON(data)
    }
}
```

### 6. Testing Infrastructure

**Files to create**:

- `internal/testing/mocks/` (new directory)
- `internal/testing/mocks/google_api.go`: Mock Google API responses
- `internal/testing/fixtures/` (new directory): Test data fixtures

**Implementation**:

```go
// mocks/google_api.go
type MockGoogleService struct {
    Responses map[string]interface{}
    Errors    map[string]error
    CallCount map[string]int
}

func (m *MockGoogleService) Record(method string) {
    m.CallCount[method]++
}

func (m *MockGoogleService) SetResponse(method string, response interface{}) {
    m.Responses[method] = response
}

func (m *MockGoogleService) SetError(method string, err error) {
    m.Errors[method] = err
}
```

### 7. Update Auth CLI

**Files to modify**:

- `internal/cli/auth.go`: Add preset scope options

**Changes**:

```go
authLoginCmd.Flags().StringVar(&authPreset, "preset", "workspace-basic",
    "Scope preset: workspace-basic, workspace-full, admin")

func runAuthLogin(cmd *cobra.Command, args []string) error {
    var scopes []string

    switch authPreset {
    case "workspace-basic":
        scopes = utils.ScopesWorkspaceBasic
    case "workspace-full":
        scopes = utils.ScopesWorkspaceFull
    case "admin":
        scopes = utils.ScopesAdmin
    default:
        scopes = authScopes
    }

    if err := mgr.ValidateScopes(scopes); err != nil {
        return err
    }

    mgr.SetOAuthConfig(clientID, clientSecret, scopes)
    // ... rest
}
```

## Todo

- [x] Add all Google API scopes to constants.go
- [x] Implement scope validation in auth.Manager
- [x] Create service factory pattern
- [x] Implement unified error handling (google_api.go, translator.go)
- [x] Add retry logic with exponential backoff
- [x] Create TableRenderer interface and update output.go
- [x] Set up testing infrastructure with mocks
- [x] Update auth CLI with scope presets
- [x] Add integration tests for auth flow
- [x] Document scope requirements in README

## Dependencies

This plan must be completed before:

- Sheets CLI Integration
- Docs CLI Integration
- Slides CLI Integration
- Admin SDK Directory CLI Integration

## Testing Strategy

1. **Unit tests**: Mock Google API responses
2. **Integration tests**: Use test Google Workspace account
3. **Error scenarios**: Test rate limiting, auth failures, network errors
4. **Scope validation**: Test incompatible scope combinations