package api

import (
	"testing"

	"github.com/dl-alexandre/gdrive/internal/utils"
)

type testError struct {
	msg string
}

func (e *testError) Error() string {
	return e.msg
}

func TestIsOperationExpired(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			"file not found",
			utils.NewAppError(utils.NewCLIError(utils.ErrCodeFileNotFound, "not found").Build()),
			true,
		},
		{
			"other app error",
			utils.NewAppError(utils.NewCLIError(utils.ErrCodeNetworkError, "network error").Build()),
			false,
		},
		{
			"non app error",
			&testError{msg: "some error"},
			false,
		},
		{
			"nil error",
			nil,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isOperationExpired(tt.err)
			if got != tt.want {
				t.Fatalf("isOperationExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClassifyOperationError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			"expired",
			utils.NewAppError(utils.NewCLIError(utils.ErrCodeFileNotFound, "not found").Build()),
			"expired",
		},
		{
			"retryable network",
			utils.NewAppError(utils.NewCLIError(utils.ErrCodeNetworkError, "network error").Build()),
			"retryable",
		},
		{
			"retryable rate limited",
			utils.NewAppError(utils.NewCLIError(utils.ErrCodeRateLimited, "rate limited").Build()),
			"retryable",
		},
		{
			"fatal",
			utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidArgument, "invalid").Build()),
			"fatal",
		},
		{
			"unknown",
			&testError{msg: "some error"},
			"unknown",
		},
		{
			"nil",
			nil,
			"unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClassifyOperationError(tt.err)
			if got != tt.want {
				t.Fatalf("ClassifyOperationError() = %q, want %q", got, tt.want)
			}
		})
	}
}
