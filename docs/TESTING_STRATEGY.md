# Testing Strategy & Coverage Improvement Plan

## Current Coverage Status

| Package | Coverage | Priority | Difficulty |
|---------|----------|----------|------------|
| internal/utils | 100.0% | ✅ Complete | - |
| internal/export | 97.1% | ✅ Complete | - |
| internal/safety | 77.4% | 🟢 Good | Low |
| internal/api | 72.7% | 🟢 Good | Medium |
| internal/logging | 67.1% | 🟢 Good | Low |
| internal/config | 63.6% | 🟢 Good | Low |
| internal/slides | 52.1% | 🟡 Medium | Medium |
| cmd/gdrv | 50.0% | 🟡 Medium | Low |
| internal/docs | 49.0% | 🟡 Medium | Medium |
| internal/types | 36.5% | 🟡 Medium | Low |
| internal/sheets | 35.7% | 🟡 Medium | Medium |
| internal/resolver | 30.7% | 🔴 Needs Work | High |
| internal/errors | 27.5% | 🔴 Needs Work | Medium |
| internal/auth | 24.8% | 🔴 Critical | High |
| internal/drives | 23.3% | 🔴 Needs Work | Medium |
| internal/cli | 16.6% | 🔴 Critical | Medium |
| internal/admin | 13.7% | 🔴 Needs Work | Medium |
| **internal/files** | **3.1%** | 🔴 **Critical** | **High** |
| **internal/folders** | **3.2%** | 🔴 **Critical** | **High** |
| **internal/permissions** | **0.9%** | 🔴 **Critical** | **High** |
| **internal/revisions** | **0.9%** | 🔴 Needs Work | Medium |
| internal/sync/* | 0.0% | 🔴 No Tests | High |
| pkg/version | 0.0% | 🟢 Trivial | Low |

## Why Some Packages Have Low Coverage

### API-Dependent Packages (files, folders, permissions, drives)

These packages are challenging to test because they:
1. Depend heavily on Google Drive API calls via `api.Client`
2. Use `api.ExecuteWithRetry()` which requires a fully initialized Drive service
3. Need complex mocking infrastructure

**Current Test Structure:**
- Property-based tests (excellent coverage of edge cases)
- Unit tests for helper functions and type conversions
- **Missing:** Integration-style tests for Manager methods

**What's Needed:**
- Mock `api.Client` or create test doubles
- Mock `drive.Service` responses
- Test happy paths, error paths, and retry logic

### Authentication Package (auth - 24.8%)

**Missing Coverage:**
- OAuth flow integration
- Service account loading
- Keyring storage operations (platform-specific)
- Token refresh logic
- Profile management

**Recommendation:** Create table-driven tests for core logic, mock keyring operations

### CLI Package (cli - 16.6%)

**Missing Coverage:**
- Command execution paths
- Flag parsing and validation
- Output formatting
- Error handling in commands

**Recommendation:** Create test cases that exercise command logic without full API calls

### Sync Package (0.0%)

Appears to be a newer or incomplete feature. All subpackages (conflict, diff, exclude, executor, index, scanner) have no tests.

## Recommended Testing Approach

### Phase 1: Quick Wins (Low-hanging fruit)

1. **pkg/version** (0% → 100%)
   - Trivial: just test the version string functions
   ```go
   func TestVersion(t *testing.T) {
       if Version == "" { t.Error("Version should not be empty") }
   }
   ```

2. **cmd/gdrv main_test.go** (50% → 80%)
   - Test the `run()` function with different scenarios
   - Mock `cli.Execute()` responses

3. **internal/types** (36.5% → 70%)
   - Add tests for all struct marshaling/unmarshaling
   - Test type conversions and validations

### Phase 2: Core Logic Testing (Medium effort)

4. **internal/errors** (27.5% → 70%)
   - Test error classification logic
   - Test error message formatting
   - Test exit code mapping

5. **internal/resolver** (30.7% → 60%)
   - Test path parsing logic
   - Test cache behavior
   - Mock API calls for path resolution

6. **internal/cli** (16.6% → 50%)
   - Test command registration
   - Test flag validation
   - Test output formatting
   - Create helper to mock global state

### Phase 3: API-Dependent Packages (High effort)

7. **internal/files, folders, permissions** (3% → 40%+)

   **Strategy A: Create Mock Infrastructure**
   ```go
   // Create internal/testing/mocks package
   type MockDriveService struct {
       FilesService   *MockFilesService
       PermissionsService *MockPermissionsService
   }

   type MockFilesService struct {
       GetFunc    func(fileID string) *drive.File
       ListFunc   func() *drive.FileList
       CreateFunc func(file *drive.File) *drive.File
       // ... etc
   }
   ```

   **Strategy B: Interface-Based Testing**
   - Extract interfaces for testability
   - Use dependency injection
   - Create test implementations

   **Strategy C: Integration Tests**
   - Use real credentials in CI (encrypted secrets)
   - Tag as `// +build integration`
   - Run separately from unit tests

8. **internal/auth** (24.8% → 60%)
   - Mock keyring operations
   - Test OAuth config generation
   - Test token refresh logic
   - Test profile switching

### Phase 4: New Features

9. **internal/sync** (0% → 60%)
   - Start with TDD for new code
   - Test each subpackage independently
   - Create integration tests for full sync workflow

## Testing Best Practices for This Codebase

### 1. Use Table-Driven Tests

```go
func TestSomething(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    string
        wantErr bool
    }{
        {"happy path", "input", "output", false},
        {"error case", "bad", "", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Something(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
            }
            if got != tt.want {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}
```

### 2. Property-Based Tests (Already Good!)

The codebase already has excellent property tests (e.g., `*_property_test.go`). Continue this pattern!

### 3. Test Helpers

Create test helpers in `internal/testing/`:
- `testhelpers.go`: Common test utilities
- `mocks/`: Mock implementations
- `fixtures/`: Test data

### 4. Integration Tests

```bash
# Run only unit tests
go test ./...

# Run integration tests (requires credentials)
go test -tags=integration ./test/integration/...
```

### 5. Coverage Targets

- **Critical paths** (auth, files, folders, permissions): 60%+ coverage
- **Core logic** (API, CLI, errors): 70%+ coverage
- **Utilities** (utils, export): 90%+ coverage
- **Overall project**: 50%+ coverage

## How to Run Tests

```bash
# Run all tests with coverage
make test-coverage

# Run specific package
go test -v -cover ./internal/files/

# Run with race detection
go test -v -race ./...

# Generate HTML coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run only fast tests (exclude slow integration tests)
go test -short ./...
```

## Next Steps

1. ✅ **Done**: Added comprehensive tests for files package helper functions
2. **Next**: Create mock infrastructure for API-dependent packages
3. **Then**: Increase CLI package coverage
4. **Then**: Add auth package tests with mocked keyring
5. **Finally**: Create integration test suite for CI/CD

## Contributing

When adding new features:
1. Write tests first (TDD)
2. Aim for 70%+ coverage on new code
3. Use table-driven tests for multiple scenarios
4. Add property tests for invariants
5. Document complex test setups
