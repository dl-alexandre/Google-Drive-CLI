package errors

import (
	"github.com/dl-alexandre/gdrive/internal/logging"
	"github.com/dl-alexandre/gdrive/internal/types"
	"github.com/dl-alexandre/gdrive/internal/utils"
	"google.golang.org/api/googleapi"
)

func ClassifyGoogleAPIError(service string, err error, reqCtx *types.RequestContext, logger logging.Logger) error {
	apiErr, ok := err.(*googleapi.Error)
	if !ok {
		logger.Error("Non-API error",
			logging.F("error", err.Error()),
			logging.F("traceId", reqCtx.TraceID),
		)
		return utils.NewAppError(utils.NewCLIError(utils.ErrCodeNetworkError, err.Error()).
			WithRetryable(true).
			WithContext("traceId", reqCtx.TraceID).
			WithContext("service", service).
			Build())
	}

	var code string
	var retryable bool

	switch apiErr.Code {
	case 400:
		code = utils.ErrCodeInvalidArgument
		for _, e := range apiErr.Errors {
			switch e.Reason {
			case "invalidSharingRequest":
				code = utils.ErrCodeSharingRestricted
			case "teamDriveFileLimitExceeded":
				code = utils.ErrCodeQuotaExceeded
			}
		}
	case 401:
		code = utils.ErrCodeAuthExpired
	case 403:
		code = utils.ErrCodePermissionDenied
		for _, e := range apiErr.Errors {
			switch e.Reason {
			case "storageQuotaExceeded":
				code = utils.ErrCodeQuotaExceeded
			case "sharingRateLimitExceeded", "userRateLimitExceeded", "rateLimitExceeded":
				code = utils.ErrCodeRateLimited
				retryable = true
			case "dailyLimitExceeded":
				code = utils.ErrCodeRateLimited
			case "domainPolicy":
				code = utils.ErrCodePolicyViolation
			}
		}
	case 404:
		code = utils.ErrCodeFileNotFound
	case 409:
		code = utils.ErrCodeInvalidArgument
	case 429:
		code = utils.ErrCodeRateLimited
		retryable = true
	case 500, 502, 503, 504:
		code = utils.ErrCodeNetworkError
		retryable = true
	default:
		code = utils.ErrCodeUnknown
		retryable = apiErr.Code >= 500
	}

	logger.Error("API error classified",
		logging.F("httpStatus", apiErr.Code),
		logging.F("errorCode", code),
		logging.F("retryable", retryable),
		logging.F("message", apiErr.Message),
		logging.F("traceId", reqCtx.TraceID),
		logging.F("service", service),
	)

	builder := utils.NewCLIError(code, apiErr.Message).
		WithHTTPStatus(apiErr.Code).
		WithRetryable(retryable).
		WithContext("traceId", reqCtx.TraceID).
		WithContext("requestType", string(reqCtx.RequestType)).
		WithContext("service", service)

	if len(apiErr.Errors) > 0 {
		if service == "drive" {
			builder.WithDriveReason(apiErr.Errors[0].Reason)
		}
		switch apiErr.Errors[0].Reason {
		case "storageQuotaExceeded":
			builder.WithContext("suggestedAction", "free up space in Google Drive or upgrade storage")
		case "sharingRateLimitExceeded", "userRateLimitExceeded", "rateLimitExceeded":
			builder.WithContext("suggestedAction", "wait before retrying")
		case "dailyLimitExceeded":
			builder.WithContext("suggestedAction", "quota will reset in 24 hours")
		case "appNotAuthorizedToFile":
			builder.WithContext("suggestedAction", "file may require access via web interface first")
		case "insufficientFilePermissions":
			builder.WithContext("capability", "write_access_required")
		case "domainPolicy":
			builder.WithContext("suggestedAction", "contact domain administrator")
		}
	}

	switch code {
	case utils.ErrCodeAuthExpired:
		builder.WithContext("suggestedAction", "run 'gdrive auth login' to re-authenticate")
	case utils.ErrCodeFileNotFound:
		if reqCtx.DriveID != "" {
			builder.WithContext("searchDomain", "sharedDrive").
				WithContext("driveId", reqCtx.DriveID)
		}
		builder.WithContext("suggestedAction", "verify file ID or path is correct and accessible")
	case utils.ErrCodeRateLimited:
		builder.WithContext("suggestedAction", "rate limit exceeded, retrying with backoff")
	}

	if apiErr.Code == 409 {
		builder.WithContext("conflict", true)
	}

	if apiErr.Code >= 500 && apiErr.Code <= 504 {
		builder.WithContext("serverError", true).
			WithContext("suggestedAction", "temporary server error, retrying")
	}

	return utils.NewAppError(builder.Build())
}
