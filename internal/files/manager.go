package files

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/safety"
	"github.com/dl-alexandre/gdrv/internal/types"
	"github.com/dl-alexandre/gdrv/internal/utils"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

// Manager handles file operations
type Manager struct {
	client *api.Client
	shaper *api.RequestShaper
}

// NewManager creates a new file manager
func NewManager(client *api.Client) *Manager {
	return &Manager{
		client: client,
		shaper: api.NewRequestShaper(client),
	}
}

// UploadOptions configures file upload
type UploadOptions struct {
	ParentID    string
	Name        string
	MimeType    string
	Convert     bool
	PinRevision bool
}

type UpdateContentOptions struct {
	Name     string
	MimeType string
	Fields   string
}

// DownloadOptions configures file download
type DownloadOptions struct {
	OutputPath   string
	MimeType     string
	Wait         bool
	Timeout      int // in seconds
	PollInterval int // in seconds
}

// ListOptions configures file listing
type ListOptions struct {
	ParentID       string
	Query          string
	PageSize       int
	PageToken      string
	OrderBy        string
	IncludeTrashed bool
	Fields         string
}

// Upload uploads a file to Drive
func (m *Manager) Upload(ctx context.Context, reqCtx *types.RequestContext, localPath string, opts UploadOptions) (*types.DriveFile, error) {
	file, err := os.Open(localPath)
	if err != nil {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidArgument,
			fmt.Sprintf("Failed to open file: %s", err)).Build())
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	name := opts.Name
	if name == "" {
		name = filepath.Base(localPath)
	}

	metadata := &drive.File{
		Name: name,
	}
	if opts.ParentID != "" {
		metadata.Parents = []string{opts.ParentID}
		reqCtx.InvolvedParentIDs = append(reqCtx.InvolvedParentIDs, opts.ParentID)
	}
	if opts.MimeType != "" {
		metadata.MimeType = opts.MimeType
	}

	// Select upload type based on file size
	uploadType := selectUploadType(stat.Size(), metadata)

	var result *drive.File

	switch uploadType {
	case "simple":
		result, err = m.simpleUpload(ctx, reqCtx, file, metadata, opts)
	case "multipart":
		result, err = m.multipartUpload(ctx, reqCtx, file, metadata, opts)
	case "resumable":
		result, err = m.resumableUpload(ctx, reqCtx, file, metadata, stat.Size(), opts)
	}

	if err != nil {
		return nil, err
	}

	// Update resource key cache
	if result.ResourceKey != "" {
		m.client.ResourceKeys().UpdateFromAPIResponse(result.Id, result.ResourceKey)
	}

	return convertDriveFile(result), nil
}

func (m *Manager) UpdateContent(ctx context.Context, reqCtx *types.RequestContext, fileID string, localPath string, opts UpdateContentOptions) (*types.DriveFile, error) {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, fileID)

	file, err := os.Open(localPath)
	if err != nil {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidArgument,
			fmt.Sprintf("Failed to open file: %s", err)).Build())
	}
	defer file.Close()

	metadata := &drive.File{}
	if opts.Name != "" {
		metadata.Name = opts.Name
	}
	if opts.MimeType != "" {
		metadata.MimeType = opts.MimeType
	}

	call := m.client.Service().Files.Update(fileID, metadata).Media(file)
	call = m.shaper.ShapeFilesUpdate(call, reqCtx)
	if opts.Fields != "" {
		call = call.Fields(googleapi.Field(opts.Fields))
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.File, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	if result.ResourceKey != "" {
		m.client.ResourceKeys().UpdateFromAPIResponse(result.Id, result.ResourceKey)
	}

	return convertDriveFile(result), nil
}

func selectUploadType(size int64, metadata *drive.File) string {
	// For files larger than 5MB, use resumable upload for better reliability
	if size > int64(utils.UploadSimpleMaxBytes) {
		return "resumable"
	}
	// If metadata is provided (name, mimeType, or parents), use multipart
	// This allows sending metadata and content in a single request
	if metadata.Name != "" || metadata.MimeType != "" || len(metadata.Parents) > 0 {
		return "multipart"
	}
	// For small files without metadata, simple upload is sufficient
	return "simple"
}

func (m *Manager) simpleUpload(ctx context.Context, reqCtx *types.RequestContext, reader io.Reader, metadata *drive.File, opts UploadOptions) (*drive.File, error) {
	call := m.client.Service().Files.Create(metadata).Media(reader)
	call = m.shaper.ShapeFilesCreate(call, reqCtx)

	return api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.File, error) {
		return call.Do()
	})
}

func (m *Manager) multipartUpload(ctx context.Context, reqCtx *types.RequestContext, reader io.Reader, metadata *drive.File, opts UploadOptions) (*drive.File, error) {
	call := m.client.Service().Files.Create(metadata).Media(reader)
	call = m.shaper.ShapeFilesCreate(call, reqCtx)

	return api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.File, error) {
		return call.Do()
	})
}

func (m *Manager) resumableUpload(ctx context.Context, reqCtx *types.RequestContext, reader io.Reader, metadata *drive.File, size int64, opts UploadOptions) (*drive.File, error) {
	call := m.client.Service().Files.Create(metadata).Media(reader)
	call = m.shaper.ShapeFilesCreate(call, reqCtx)
	call = call.ProgressUpdater(func(current, total int64) {
		// Progress callback - could be used to report upload progress
		// For now, this is a placeholder for future enhancement
	})

	return api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.File, error) {
		return call.Do()
	})
}

// Download downloads a file from Drive
func (m *Manager) Download(ctx context.Context, reqCtx *types.RequestContext, fileID string, opts DownloadOptions) error {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, fileID)

	// Get file metadata first with exportLinks included for Workspace files
	fields := "id,name,mimeType,size,capabilities,exportLinks"
	file, err := m.Get(ctx, reqCtx, fileID, fields)
	if err != nil {
		return err
	}

	// Check download capability
	if file.Capabilities != nil && !file.Capabilities.CanDownload {
		return utils.NewAppError(utils.NewCLIError(utils.ErrCodePermissionDenied,
			"File cannot be downloaded").
			WithContext("capability", "canDownload=false").
			Build())
	}

	outputPath := opts.OutputPath
	if outputPath == "" {
		if opts.MimeType == "text/plain" && utils.IsWorkspaceMimeType(file.MimeType) {
			outputPath = file.Name + ".txt"
		} else {
			outputPath = file.Name
		}
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidArgument,
			fmt.Sprintf("Failed to create output file: %s", err)).Build())
	}
	defer outFile.Close()

	// Check if it's a Workspace file that needs export
	if utils.IsWorkspaceMimeType(file.MimeType) {
		return m.exportFile(ctx, reqCtx, fileID, file, opts, outFile)
	}

	return m.downloadBlob(ctx, reqCtx, fileID, outFile)
}

func (m *Manager) downloadBlob(ctx context.Context, reqCtx *types.RequestContext, fileID string, writer io.Writer) error {
	call := m.client.Service().Files.Get(fileID)
	call = m.shaper.ShapeFilesGet(call, reqCtx)

	// Download content
	httpResp, err := call.Download()
	if err != nil {
		return utils.NewAppError(utils.NewCLIError(utils.ErrCodeNetworkError,
			fmt.Sprintf("Download failed: %s", err)).Build())
	}
	defer httpResp.Body.Close()

	_, err = io.Copy(writer, httpResp.Body)
	return err
}

func (m *Manager) exportFile(ctx context.Context, reqCtx *types.RequestContext, fileID string, file *types.DriveFile, opts DownloadOptions, writer io.Writer) error {
	mimeType := opts.MimeType
	if mimeType == "" {
		mimeType = "application/pdf" // Default export format
	}

	// Check if file size exceeds export limit (10MB)
	// Note: Google Workspace files don't have a size property, but exports may fail if too large
	if file.Size > int64(utils.ExportMaxBytes) {
		// Build error with exportLinks for manual download
		errBuilder := utils.NewCLIError(utils.ErrCodeExportSizeLimit,
			fmt.Sprintf("File exceeds 10MB export limit (size: %d bytes)", file.Size)).
			WithContext("fileId", fileID).
			WithContext("size", file.Size).
			WithContext("limit", utils.ExportMaxBytes)

		// Include exportLinks if available
		if len(file.ExportLinks) > 0 {
			errBuilder.WithContext("exportLinks", file.ExportLinks)
		}

		return utils.NewAppError(errBuilder.Build())
	}

	// Try direct export first
	call := m.client.Service().Files.Export(fileID, mimeType)
	header := m.client.ResourceKeys().BuildHeader(reqCtx.InvolvedFileIDs)
	if header != "" {
		call.Header().Set("X-Goog-Drive-Resource-Keys", header)
	}

	resp, err := call.Download()
	if err != nil {
		if apiErr, ok := err.(*googleapi.Error); ok {
			// Check if this is a size limit error
			if apiErr.Code == 403 {
				for _, e := range apiErr.Errors {
					if e.Reason == "exportSizeLimitExceeded" {
						errBuilder := utils.NewCLIError(utils.ErrCodeExportSizeLimit,
							"File exceeds 10MB export limit").
							WithHTTPStatus(403).
							WithDriveReason("exportSizeLimitExceeded").
							WithContext("fileId", fileID)

						if len(file.ExportLinks) > 0 {
							errBuilder.WithContext("exportLinks", file.ExportLinks).
								WithContext("suggestedAction", "use exportLinks to download via browser")
						}

						return utils.NewAppError(errBuilder.Build())
					}
				}
			}

			return utils.NewAppError(utils.NewCLIError(utils.ErrCodeNetworkError,
				fmt.Sprintf("Export failed: %s", apiErr.Message)).
				WithHTTPStatus(apiErr.Code).
				Build())
		}
		return utils.NewAppError(utils.NewCLIError(utils.ErrCodeNetworkError,
			fmt.Sprintf("Export failed: %s", err)).Build())
	}
	defer resp.Body.Close()

	// Check if response indicates long-running operation
	if resp.StatusCode == 202 {
		// Operation started, need to poll
		if !opts.Wait {
			return utils.NewAppError(utils.NewCLIError(utils.ErrCodeNetworkError,
				"Export initiated as long-running operation. Use --wait flag to poll for completion.").
				WithContext("statusCode", 202).
				Build())
		}

		return m.pollAndDownloadOperation(ctx, reqCtx, resp, opts, writer)
	}

	_, err = io.Copy(writer, resp.Body)
	return err
}

func (m *Manager) pollAndDownloadOperation(ctx context.Context, reqCtx *types.RequestContext, resp *http.Response, opts DownloadOptions, writer io.Writer) error {
	// Get operation name from response header
	operationName := resp.Header.Get("X-Goog-Upload-URL")
	if operationName == "" {
		operationName = resp.Header.Get("Location")
	}

	if operationName == "" {
		return utils.NewAppError(utils.NewCLIError(utils.ErrCodeNetworkError,
			"Long-running operation started but no operation URL provided").Build())
	}

	// Create poller with timeout and poll interval
	timeout := time.Duration(opts.Timeout) * time.Second
	if timeout == 0 {
		timeout = 5 * time.Minute // Default 5 minute timeout
	}

	pollInterval := time.Duration(opts.PollInterval) * time.Second
	if pollInterval == 0 {
		pollInterval = 5 * time.Second // Default 5 second poll interval
	}

	// Create HTTP client from Drive service
	httpClient := &http.Client{
		Timeout: timeout,
	}
	poller := api.NewOperationPoller(httpClient, pollInterval, timeout)

	// Poll until complete
	operation, err := poller.PollUntilComplete(ctx, operationName, reqCtx)
	if err != nil {
		return err
	}

	// Check for operation error
	if operation.Error != nil {
		return operation.Error
	}

	// Download from completed operation URI
	if operation.DownloadURI == "" {
		return utils.NewAppError(utils.NewCLIError(utils.ErrCodeNetworkError,
			"Operation completed but no download URI available").Build())
	}

	// Use the HTTP client to download
	req, err := http.NewRequestWithContext(ctx, "GET", operation.DownloadURI, nil)
	if err != nil {
		return err
	}

	downloadResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return utils.NewAppError(utils.NewCLIError(utils.ErrCodeNetworkError,
			fmt.Sprintf("Download from operation URI failed: %s", err)).Build())
	}
	defer downloadResp.Body.Close()

	if downloadResp.StatusCode != http.StatusOK {
		return utils.NewAppError(utils.NewCLIError(utils.ErrCodeNetworkError,
			fmt.Sprintf("Download failed with status %d", downloadResp.StatusCode)).Build())
	}

	_, err = io.Copy(writer, downloadResp.Body)
	return err
}

// Get retrieves file metadata
func (m *Manager) Get(ctx context.Context, reqCtx *types.RequestContext, fileID string, fields string) (*types.DriveFile, error) {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, fileID)

	call := m.client.Service().Files.Get(fileID)
	call = m.shaper.ShapeFilesGet(call, reqCtx)
	if fields != "" {
		call = call.Fields(googleapi.Field(fields))
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.File, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	// Update resource key cache
	if result.ResourceKey != "" {
		m.client.ResourceKeys().UpdateFromAPIResponse(result.Id, result.ResourceKey)
	}

	return convertDriveFile(result), nil
}

// Search searches for files using a query
func (m *Manager) Search(ctx context.Context, reqCtx *types.RequestContext, query string, opts ListOptions) (*types.FileListResult, error) {
	opts.Query = query
	return m.List(ctx, reqCtx, opts)
}

// List lists files
func (m *Manager) List(ctx context.Context, reqCtx *types.RequestContext, opts ListOptions) (*types.FileListResult, error) {
	call := m.client.Service().Files.List()
	call = m.shaper.ShapeFilesList(call, reqCtx)

	// Build query
	query := ""
	if opts.ParentID != "" {
		query = fmt.Sprintf("'%s' in parents", opts.ParentID)
		reqCtx.InvolvedParentIDs = append(reqCtx.InvolvedParentIDs, opts.ParentID)
	}
	if !opts.IncludeTrashed {
		if query != "" {
			query += " and "
		}
		query += "trashed = false"
	}
	if opts.Query != "" {
		if query != "" {
			query += " and "
		}
		query += opts.Query
	}
	if query != "" {
		call = call.Q(query)
	}

	if opts.PageSize > 0 {
		call = call.PageSize(int64(opts.PageSize))
	}
	if opts.PageToken != "" {
		call = call.PageToken(opts.PageToken)
	}
	if opts.OrderBy != "" {
		call = call.OrderBy(opts.OrderBy)
	}
	if opts.Fields != "" {
		call = call.Fields(googleapi.Field("nextPageToken,incompleteSearch,files(" + opts.Fields + ")"))
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.FileList, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	files := make([]*types.DriveFile, len(result.Files))
	for i, f := range result.Files {
		files[i] = convertDriveFile(f)
		if f.ResourceKey != "" {
			m.client.ResourceKeys().UpdateFromAPIResponse(f.Id, f.ResourceKey)
		}
	}

	return &types.FileListResult{
		Files:            files,
		NextPageToken:    result.NextPageToken,
		IncompleteSearch: result.IncompleteSearch,
	}, nil
}

// ListAll lists all files by following pagination
func (m *Manager) ListAll(ctx context.Context, reqCtx *types.RequestContext, opts ListOptions) ([]*types.DriveFile, error) {
	var allFiles []*types.DriveFile
	pageToken := opts.PageToken

	for {
		opts.PageToken = pageToken
		result, err := m.List(ctx, reqCtx, opts)
		if err != nil {
			return allFiles, err
		}

		allFiles = append(allFiles, result.Files...)

		if result.NextPageToken == "" {
			break
		}
		pageToken = result.NextPageToken
	}

	return allFiles, nil
}

// Delete deletes or trashes a file
func (m *Manager) Delete(ctx context.Context, reqCtx *types.RequestContext, fileID string, permanent bool) error {
	return m.DeleteWithSafety(ctx, reqCtx, fileID, permanent, safety.Default(), nil)
}

// DeleteWithSafety deletes or trashes a file with safety controls.
// Supports dry-run mode, confirmation, and idempotency.
//
// Requirements:
//   - Requirement 13.1: Support --dry-run mode for destructive operations
//   - Requirement 13.2: Support --force flag to skip confirmations
//   - Requirement 13.4: Add idempotent behavior for retry operations
func (m *Manager) DeleteWithSafety(ctx context.Context, reqCtx *types.RequestContext, fileID string, permanent bool, opts safety.SafetyOptions, recorder safety.DryRunRecorder) error {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, fileID)

	// Get file metadata for confirmation and dry-run display
	file, err := m.Get(ctx, reqCtx, fileID, "id,name")
	if err != nil {
		return err
	}

	// Dry-run mode: record operation without executing
	if opts.DryRun && recorder != nil {
		safety.RecordDelete(recorder, fileID, file.Name, permanent)
		return nil
	}

	// Confirmation for destructive operations
	if opts.ShouldConfirm() {
		operation := "trash"
		if permanent {
			operation = "permanently delete"
		}
		confirmed, err := safety.Confirm(
			fmt.Sprintf("About to %s '%s'. Continue?", operation, file.Name),
			opts,
		)
		if err != nil {
			return err
		}
		if !confirmed {
			return utils.NewAppError(utils.NewCLIError(utils.ErrCodeCancelled, "Operation cancelled by user").Build())
		}
	}

	if permanent {
		call := m.client.Service().Files.Delete(fileID)
		call = m.shaper.ShapeFilesDelete(call, reqCtx)

		_, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (interface{}, error) {
			return nil, call.Do()
		})
		return err
	}

	// Move to trash
	call := m.client.Service().Files.Update(fileID, &drive.File{Trashed: true})
	call = m.shaper.ShapeFilesUpdate(call, reqCtx)

	_, err = api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.File, error) {
		return call.Do()
	})
	return err
}

// Copy copies a file
func (m *Manager) Copy(ctx context.Context, reqCtx *types.RequestContext, fileID string, name string, parentID string) (*types.DriveFile, error) {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, fileID)
	if parentID != "" {
		reqCtx.InvolvedParentIDs = append(reqCtx.InvolvedParentIDs, parentID)
	}

	metadata := &drive.File{}
	if name != "" {
		metadata.Name = name
	}
	if parentID != "" {
		metadata.Parents = []string{parentID}
	}

	call := m.client.Service().Files.Copy(fileID, metadata)
	call = m.shaper.ShapeFilesCopy(call, reqCtx)

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.File, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	return convertDriveFile(result), nil
}

// Move moves a file to a new parent
func (m *Manager) Move(ctx context.Context, reqCtx *types.RequestContext, fileID string, newParentID string) (*types.DriveFile, error) {
	return m.MoveWithSafety(ctx, reqCtx, fileID, newParentID, safety.Default(), nil)
}

// MoveWithSafety moves a file to a new parent with safety controls.
// Supports dry-run mode and confirmation.
//
// Requirements:
//   - Requirement 13.1: Support --dry-run mode for destructive operations
func (m *Manager) MoveWithSafety(ctx context.Context, reqCtx *types.RequestContext, fileID string, newParentID string, opts safety.SafetyOptions, recorder safety.DryRunRecorder) (*types.DriveFile, error) {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, fileID)
	reqCtx.InvolvedParentIDs = append(reqCtx.InvolvedParentIDs, newParentID)

	// Get current file info
	file, err := m.Get(ctx, reqCtx, fileID, "parents,name")
	if err != nil {
		return nil, err
	}

	// Dry-run mode: record operation without executing
	if opts.DryRun && recorder != nil {
		safety.RecordMove(recorder, fileID, file.Name, newParentID, newParentID)
		return file, nil
	}

	var removeParents string
	if len(file.Parents) > 0 {
		for i, p := range file.Parents {
			if i > 0 {
				removeParents += ","
			}
			removeParents += p
		}
	}

	call := m.client.Service().Files.Update(fileID, &drive.File{})
	call = m.shaper.ShapeFilesUpdate(call, reqCtx)
	call = call.AddParents(newParentID)
	if removeParents != "" {
		call = call.RemoveParents(removeParents)
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.File, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	return convertDriveFile(result), nil
}

// Trash moves a file to trash
func (m *Manager) Trash(ctx context.Context, reqCtx *types.RequestContext, fileID string) (*types.DriveFile, error) {
	return m.updateTrashed(ctx, reqCtx, fileID, true)
}

// ListTrashed lists trashed files with pagination
func (m *Manager) ListTrashed(ctx context.Context, reqCtx *types.RequestContext, opts ListOptions) (*types.FileListResult, error) {
	// Force include trashed files
	opts.IncludeTrashed = true

	// Modify query to only show trashed files
	if opts.Query != "" {
		opts.Query = "trashed = true and (" + opts.Query + ")"
	} else {
		opts.Query = "trashed = true"
	}

	return m.List(ctx, reqCtx, opts)
}

// SearchTrashed searches for files in trash
func (m *Manager) SearchTrashed(ctx context.Context, reqCtx *types.RequestContext, query string, opts ListOptions) (*types.FileListResult, error) {
	opts.Query = query
	return m.ListTrashed(ctx, reqCtx, opts)
}

// Restore restores a file from trash
func (m *Manager) Restore(ctx context.Context, reqCtx *types.RequestContext, fileID string) (*types.DriveFile, error) {
	return m.updateTrashed(ctx, reqCtx, fileID, false)
}

func (m *Manager) updateTrashed(ctx context.Context, reqCtx *types.RequestContext, fileID string, trashed bool) (*types.DriveFile, error) {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, fileID)

	call := m.client.Service().Files.Update(fileID, &drive.File{Trashed: trashed})
	call = m.shaper.ShapeFilesUpdate(call, reqCtx)

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.File, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	return convertDriveFile(result), nil
}

// Update updates file metadata
func (m *Manager) Update(ctx context.Context, reqCtx *types.RequestContext, fileID string, metadata *drive.File, fields string) (*types.DriveFile, error) {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, fileID)

	call := m.client.Service().Files.Update(fileID, metadata)
	call = m.shaper.ShapeFilesUpdate(call, reqCtx)
	if fields != "" {
		call = call.Fields(googleapi.Field(fields))
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.File, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	// Update resource key cache
	if result.ResourceKey != "" {
		m.client.ResourceKeys().UpdateFromAPIResponse(result.Id, result.ResourceKey)
	}

	return convertDriveFile(result), nil
}

func convertDriveFile(f *drive.File) *types.DriveFile {
	file := &types.DriveFile{
		ID:             f.Id,
		Name:           f.Name,
		MimeType:       f.MimeType,
		Size:           f.Size,
		MD5Checksum:    f.Md5Checksum,
		CreatedTime:    f.CreatedTime,
		ModifiedTime:   f.ModifiedTime,
		Parents:        f.Parents,
		ResourceKey:    f.ResourceKey,
		ExportLinks:    f.ExportLinks,
		WebViewLink:    f.WebViewLink,
		WebContentLink: f.WebContentLink,
		Trashed:        f.Trashed,
	}

	if f.Capabilities != nil {
		file.Capabilities = &types.FileCapabilities{
			CanDownload:      f.Capabilities.CanDownload,
			CanEdit:          f.Capabilities.CanEdit,
			CanShare:         f.Capabilities.CanShare,
			CanDelete:        f.Capabilities.CanDelete,
			CanTrash:         f.Capabilities.CanTrash,
			CanReadRevisions: f.Capabilities.CanReadRevisions,
		}
	}

	return file
}
