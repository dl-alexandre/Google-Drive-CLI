package utils

import (
	"fmt"

	"github.com/dl-alexandre/gdrv/internal/types"
)

// Exit codes
const (
	ExitSuccess = 0
	// Auth errors (10-19)
	ExitAuthRequired      = 10
	ExitAuthExpired       = 11
	ExitAuthInvalid       = 12
	ExitScopeInsufficient = 13
	// File operation errors (20-29)
	ExitFileNotFound             = 20
	ExitPermissionDenied         = 21
	ExitQuotaExceeded            = 22
	ExitExportSizeLimit          = 23
	ExitRevisionNotDownloadable  = 24
	ExitRevisionKeepForeverLimit = 25
	// Network errors (30-39)
	ExitNetworkError     = 30
	ExitTimeout          = 31
	ExitRateLimited      = 32
	ExitOperationExpired = 33
	// Validation errors (40-49)
	ExitInvalidArgument = 40
	ExitInvalidPath     = 41
	ExitAmbiguousPath   = 42
	ExitInvalidMimeType = 43
	// Policy errors (50-59)
	ExitPolicyViolation   = 50
	ExitSharingRestricted = 51
	// Batch errors
	ExitBatchPartialFailure = 60
	// Unknown
	ExitUnknown = 99
)

// Error codes (tool-owned, stable)
const (
	ErrCodeAuthRequired             = "AUTH_REQUIRED"
	ErrCodeAuthExpired              = "AUTH_EXPIRED"
	ErrCodeAuthClientMissing        = "AUTH_CLIENT_MISSING"
	ErrCodeAuthClientInvalid        = "AUTH_CLIENT_INVALID"
	ErrCodeAuthClientPartial        = "AUTH_CLIENT_PARTIAL"
	ErrCodeScopeInsufficient        = "SCOPE_INSUFFICIENT"
	ErrCodeFileNotFound             = "FILE_NOT_FOUND"
	ErrCodePermissionDenied         = "PERMISSION_DENIED"
	ErrCodeQuotaExceeded            = "QUOTA_EXCEEDED"
	ErrCodeExportSizeLimit          = "EXPORT_SIZE_LIMIT"
	ErrCodeRevisionNotDownloadable  = "REVISION_NOT_DOWNLOADABLE"
	ErrCodeRevisionKeepForeverLimit = "REVISION_KEEP_FOREVER_LIMIT"
	ErrCodeNetworkError             = "NETWORK_ERROR"
	ErrCodeTimeout                  = "TIMEOUT"
	ErrCodeRateLimited              = "RATE_LIMITED"
	ErrCodeOperationExpired         = "OPERATION_EXPIRED"
	ErrCodeInvalidArgument          = "INVALID_ARGUMENT"
	ErrCodeInvalidPath              = "INVALID_PATH"
	ErrCodeAmbiguousPath            = "AMBIGUOUS_PATH"
	ErrCodeInvalidMimeType          = "INVALID_MIME_TYPE"
	ErrCodePolicyViolation          = "POLICY_VIOLATION"
	ErrCodeSharingRestricted        = "SHARING_RESTRICTED"
	ErrCodeBatchPartialFailure      = "BATCH_PARTIAL_FAILURE"
	ErrCodeCancelled                = "CANCELLED"
	ErrCodeResourceLimit            = "RESOURCE_LIMIT"
	ErrCodeInternalError            = "INTERNAL_ERROR"
	ErrCodeUnknown                  = "UNKNOWN"
)

// CLIErrorBuilder helps construct CLIError instances
type CLIErrorBuilder struct {
	err types.CLIError
}

// NewCLIError creates a new error builder
func NewCLIError(code, message string) *CLIErrorBuilder {
	return &CLIErrorBuilder{
		err: types.CLIError{
			Code:    code,
			Message: message,
		},
	}
}

func (b *CLIErrorBuilder) WithHTTPStatus(status int) *CLIErrorBuilder {
	b.err.HTTPStatus = status
	return b
}

func (b *CLIErrorBuilder) WithDriveReason(reason string) *CLIErrorBuilder {
	b.err.DriveReason = reason
	return b
}

func (b *CLIErrorBuilder) WithRetryable(retryable bool) *CLIErrorBuilder {
	b.err.Retryable = retryable
	return b
}

func (b *CLIErrorBuilder) WithContext(key string, value interface{}) *CLIErrorBuilder {
	if b.err.Context == nil {
		b.err.Context = make(map[string]interface{})
	}
	b.err.Context[key] = value
	return b
}

func (b *CLIErrorBuilder) Build() types.CLIError {
	return b.err
}

// GetExitCode returns the exit code for an error code
func GetExitCode(errorCode string) int {
	mapping := map[string]int{
		ErrCodeAuthRequired:             ExitAuthRequired,
		ErrCodeAuthExpired:              ExitAuthExpired,
		ErrCodeAuthClientMissing:        ExitAuthRequired,
		ErrCodeAuthClientInvalid:        ExitAuthRequired,
		ErrCodeAuthClientPartial:        ExitAuthRequired,
		ErrCodeScopeInsufficient:        ExitScopeInsufficient,
		ErrCodeFileNotFound:             ExitFileNotFound,
		ErrCodePermissionDenied:         ExitPermissionDenied,
		ErrCodeQuotaExceeded:            ExitQuotaExceeded,
		ErrCodeExportSizeLimit:          ExitExportSizeLimit,
		ErrCodeRevisionNotDownloadable:  ExitRevisionNotDownloadable,
		ErrCodeRevisionKeepForeverLimit: ExitRevisionKeepForeverLimit,
		ErrCodeNetworkError:             ExitNetworkError,
		ErrCodeTimeout:                  ExitTimeout,
		ErrCodeRateLimited:              ExitRateLimited,
		ErrCodeOperationExpired:         ExitOperationExpired,
		ErrCodeInvalidArgument:          ExitInvalidArgument,
		ErrCodeInvalidPath:              ExitInvalidPath,
		ErrCodeAmbiguousPath:            ExitAmbiguousPath,
		ErrCodeInvalidMimeType:          ExitInvalidMimeType,
		ErrCodePolicyViolation:          ExitPolicyViolation,
		ErrCodeSharingRestricted:        ExitSharingRestricted,
		ErrCodeBatchPartialFailure:      ExitBatchPartialFailure,
	}
	if code, ok := mapping[errorCode]; ok {
		return code
	}
	return ExitUnknown
}

// AppError is a custom error type that carries CLI error info
type AppError struct {
	CLIError types.CLIError
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.CLIError.Code, e.CLIError.Message)
}

// NewAppError creates an AppError from a CLIError
func NewAppError(cliErr types.CLIError) *AppError {
	return &AppError{CLIError: cliErr}
}
