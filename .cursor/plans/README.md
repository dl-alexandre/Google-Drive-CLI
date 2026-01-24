# Google Drive CLI - Implementation Plans

This directory contains detailed implementation plans for expanding the Google Drive CLI to support the full Google Workspace suite.

## Overview

The Google Drive CLI is being expanded from Drive-only operations to a comprehensive **Google Workspace CLI** supporting:
- **Google Drive** (existing)
- **Google Sheets** (planned)
- **Google Docs** (planned)
- **Google Slides** (planned)
- **Admin SDK Directory** (planned)

## Implementation Order

### Phase 1: Foundation (REQUIRED FIRST)
**Plan**: `00_foundation_google_apis.plan.md`

**Why first**: Establishes shared infrastructure that all other integrations depend on.

**Key deliverables**:
- Unified scope management for all Google APIs
- Service factory pattern
- Consistent error handling and retry logic with exponential backoff
- TableRenderer interface for output formatting
- Testing infrastructure with mocks

**Estimated complexity**: Medium
**Blocks**: All other plans

---

### Phase 2: API Integrations (PARALLEL)

These can be implemented in parallel after the foundation is complete. Recommended order based on user demand and complexity:

#### 2A. Sheets CLI Integration (RECOMMENDED FIRST)
**Plan**: `01_sheets_cli_integration.plan.md`

**Why**: Most commonly requested, enables data automation workflows.

**Key features**:
- Read/write/append to spreadsheet ranges
- Batch updates (formatting, formulas)
- List spreadsheets
- Extract data for AI/LLM processing

**Estimated complexity**: Medium
**Depends on**: Foundation
**Blocks**: None

---

#### 2B. Docs CLI Integration
**Plan**: `02_docs_cli_integration.plan.md`

**Why**: Enables document automation and text extraction.

**Key features**:
- Extract text content
- Batch document updates
- Create documents
- Template-based generation

**Estimated complexity**: Medium
**Depends on**: Foundation
**Blocks**: None

---

#### 2C. Slides CLI Integration
**Plan**: `03_slides_cli_integration.plan.md`

**Why**: Enables presentation automation and templating.

**Key features**:
- Extract text from presentations
- Template-based slide generation (replace placeholders)
- Batch slide updates
- Automated report generation

**Estimated complexity**: Medium
**Depends on**: Foundation
**Blocks**: None

---

#### 2D. Admin SDK Directory CLI Integration (DIFFERENT AUTH)
**Plan**: `04_admin_sdk_directory_cli.plan.md`

**Why**: Enterprise IT admin use cases, but requires different authentication.

**Key features**:
- User management (list, create, suspend, delete)
- Group management
- Group membership operations
- Bulk user operations

**Estimated complexity**: High
**Special requirements**: Service account with domain-wide delegation
**Depends on**: Foundation
**Blocks**: None

**Note**: This requires service account authentication, unlike the other APIs which can use OAuth. Implement this separately or last.

---

## Architecture Improvements

### Current Architecture
```
gdrive (root CLI)
├── files (Drive operations)
├── folders
├── permissions
├── drives (Shared Drives)
├── auth
└── config
```

### Target Architecture
```
gdrive (root CLI)
├── files (Drive operations)
├── folders
├── permissions
├── drives (Shared Drives)
├── sheets (NEW)
│   ├── list
│   ├── get
│   ├── update
│   ├── append
│   ├── batch-update
│   └── metadata
├── docs (NEW)
│   ├── list
│   ├── get
│   ├── read
│   ├── update
│   └── create
├── slides (NEW)
│   ├── list
│   ├── get
│   ├── read
│   ├── update
│   ├── create
│   └── replace (templating)
├── admin (NEW - requires service account)
│   ├── users
│   │   ├── list
│   │   ├── get
│   │   ├── create
│   │   ├── update
│   │   ├── suspend
│   │   └── delete
│   ├── groups
│   │   ├── list
│   │   ├── get
│   │   ├── create
│   │   └── delete
│   └── members
│       ├── list
│       ├── add
│       └── remove
├── auth
│   ├── login
│   ├── device
│   ├── service-account (NEW)
│   └── logout
└── config
```

## Key Technical Decisions

### 1. Scope Management Strategy
- Default OAuth scopes include Drive + Sheets + Docs + Slides for "workspace-basic" preset
- Admin SDK scopes separate (service account only)
- Auth presets: `workspace-basic`, `workspace-full`, `admin`
- Scope validation prevents incompatible combinations

### 2. Service Creation Pattern
- ServiceFactory abstracts service creation
- All services use shared HTTP client from auth manager
- Consistent error handling wrapper
- Retry logic with exponential backoff

### 3. Output Formatting
- TableRenderer interface for consistent table output
- All response types implement TableRenderer
- JSON output always available via `--json` flag
- Agent-friendly structured output

### 4. Error Handling
- Unified Google API error parser
- Rate limiting detection and retry
- Consistent error messages across APIs
- Exit codes aligned with existing patterns

### 5. Testing Strategy
- Unit tests with mocked Google API responses
- Integration tests with test Workspace account
- Fixtures for common API responses
- Error scenario coverage

## Dependencies

### Go Dependencies to Add
```
google.golang.org/api/sheets/v4
google.golang.org/api/docs/v1
google.golang.org/api/slides/v1
google.golang.org/api/admin/directory/v1
```

### Auth Requirements

**For Sheets, Docs, Slides**:
- OAuth2 user flow (existing)
- OR service account (existing)
- Scopes added to existing auth flow

**For Admin SDK**:
- Service account with domain-wide delegation (REQUIRED)
- Cannot use OAuth user flow
- Requires admin user impersonation

## Migration Strategy

### Backward Compatibility
All existing Drive commands remain unchanged. New commands are purely additive.

### Auth Migration
Users with existing OAuth tokens will need to re-authenticate to add Workspace scopes:
```bash
# Re-authenticate to add Workspace scopes
gdrive auth login --preset workspace-full

# Or keep existing Drive-only auth
gdrive auth login --scopes https://www.googleapis.com/auth/drive
```

### Documentation Updates
- Update README with all new commands
- Add setup guide for Admin SDK (domain-wide delegation)
- Include example workflows for each API
- Update API reference

## Testing Requirements

### Unit Tests
- All manager methods mocked
- Error scenarios covered
- Input validation tested

### Integration Tests
- Require test Google Workspace account
- Test each API operation end-to-end
- Verify output formats (JSON and table)
- Test pagination and limits

### End-to-End Workflows
- Sheet data extraction → processing → update
- Doc text extraction → AI processing
- Slide templating workflow
- User provisioning workflow

## Success Metrics

### Feature Completeness
- [ ] All planned commands implemented
- [ ] JSON and table output for all operations
- [ ] Pagination support where applicable
- [ ] Error handling covers all API error codes

### Quality
- [ ] >80% unit test coverage
- [ ] Integration tests passing
- [ ] Documentation complete
- [ ] Examples for all common use cases

### Performance
- [ ] Operations complete in <2s for simple operations
- [ ] Retry logic handles rate limits gracefully
- [ ] Pagination handles large result sets efficiently

## Risk Mitigation

### Auth Complexity
**Risk**: Multiple auth modes (OAuth vs service account) confusing users
**Mitigation**: Clear error messages, setup guide, validation

### API Rate Limits
**Risk**: Hitting Google API quotas
**Mitigation**: Exponential backoff, clear error messages, batch operations

### Scope Creep
**Risk**: Trying to support every API feature
**Mitigation**: Focus on core use cases first, document future enhancements

### Breaking Changes
**Risk**: Changes to existing commands
**Mitigation**: All changes are additive, maintain backward compatibility

## Future Enhancements

After initial implementation, consider:
- **Calendar API**: Event management
- **Gmail API**: Email operations
- **Meet API**: Meeting management
- **Tasks API**: Task management
- **Forms API**: Form response collection
- **Batch operations**: Multi-file/multi-user operations
- **Webhooks**: Change notifications
- **Export automation**: Scheduled exports

## Getting Started

1. **Read the foundation plan**: `00_foundation_google_apis.plan.md`
2. **Implement foundation**: This is required for everything else
3. **Choose an API**: Start with Sheets (recommended) or your preferred API
4. **Follow the plan**: Each plan has detailed implementation steps
5. **Test thoroughly**: Unit tests and integration tests
6. **Document**: Update README with examples

## Questions?

- **Architecture questions**: Review foundation plan
- **API-specific questions**: Review individual integration plans
- **Auth questions**: Check foundation plan auth section
- **Testing questions**: Review testing strategy in each plan

---

**Last Updated**: 2026-01-24
**Status**: ✅ **Completed** - All implementation plans have been successfully completed
**Next Step**: All planned features have been implemented. The Google Drive CLI now supports Drive, Sheets, Docs, Slides, and Admin SDK Directory operations.
