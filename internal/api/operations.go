package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dl-alexandre/gdrv/internal/types"
	"github.com/dl-alexandre/gdrv/internal/utils"
)

// OperationStatus represents the status of a long-running operation
type OperationStatus string

const (
	OperationStatusPending  OperationStatus = "pending"
	OperationStatusRunning  OperationStatus = "running"
	OperationStatusComplete OperationStatus = "complete"
	OperationStatusFailed   OperationStatus = "failed"
	OperationStatusExpired  OperationStatus = "expired"
)

// Operation represents a long-running operation
type Operation struct {
	Name        string
	Done        bool
	DownloadURI string
	Error       error
	Metadata    map[string]interface{}
}

// OperationPoller polls for operation completion
type OperationPoller struct {
	client       *http.Client
	pollInterval time.Duration
	timeout      time.Duration
}

// NewOperationPoller creates a new operation poller
func NewOperationPoller(client *http.Client, pollInterval, timeout time.Duration) *OperationPoller {
	return &OperationPoller{
		client:       client,
		pollInterval: pollInterval,
		timeout:      timeout,
	}
}

// PollUntilComplete polls an operation until it completes or times out
func (p *OperationPoller) PollUntilComplete(ctx context.Context, operationName string, reqCtx *types.RequestContext) (*Operation, error) {
	deadline := time.Now().Add(p.timeout)

	for time.Now().Before(deadline) {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		op, err := p.checkOperation(ctx, operationName)
		if err != nil {
			// Check if operation expired
			if isOperationExpired(err) {
				return &Operation{
					Name: operationName,
					Done: false,
					Error: utils.NewAppError(utils.NewCLIError(utils.ErrCodeOperationExpired,
						"Operation expired. Re-issue the download request.").
						WithContext("operationName", operationName).
						Build()),
				}, nil
			}
			return nil, err
		}

		if op.Done {
			return op, nil
		}

		time.Sleep(p.pollInterval)
	}

	return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeTimeout,
		"Operation polling timed out").
		WithContext("operationName", operationName).
		WithContext("timeout", p.timeout.String()).
		Build())
}

func (p *OperationPoller) checkOperation(ctx context.Context, name string) (*Operation, error) {
	// In real implementation, this would call operations.get API
	// For now, return a placeholder
	return &Operation{
		Name: name,
		Done: false,
	}, nil
}

func isOperationExpired(err error) bool {
	// Check for 404 Not Found or specific expired indicators
	if appErr, ok := err.(*utils.AppError); ok {
		return appErr.CLIError.Code == utils.ErrCodeFileNotFound
	}
	return false
}

// ClassifyOperationError classifies operation errors
func ClassifyOperationError(err error) string {
	if appErr, ok := err.(*utils.AppError); ok {
		switch appErr.CLIError.Code {
		case utils.ErrCodeFileNotFound:
			return "expired"
		case utils.ErrCodeNetworkError, utils.ErrCodeRateLimited:
			return "retryable"
		default:
			return "fatal"
		}
	}
	return "unknown"
}

// DownloadFromURI downloads content from a URI
func DownloadFromURI(ctx context.Context, client *http.Client, uri string, writer io.Writer) error {
	req, err := http.NewRequestWithContext(ctx, "GET", uri, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed with status %d", resp.StatusCode)
	}

	_, err = io.Copy(writer, resp.Body)
	return err
}
