package revisions

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/types"
	"github.com/dl-alexandre/gdrv/internal/utils"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

// Manager handles file revision operations
type Manager struct {
	client *api.Client
	shaper *api.RequestShaper
}

// NewManager creates a new revision manager
func NewManager(client *api.Client) *Manager {
	return &Manager{
		client: client,
		shaper: api.NewRequestShaper(client),
	}
}

// ListOptions configures revision listing
type ListOptions struct {
	PageSize  int
	PageToken string
	Fields    string
}

// ListResult represents paginated revision list response
type ListResult struct {
	Revisions     []*types.Revision
	NextPageToken string
}

// List retrieves file revisions with pagination
func (m *Manager) List(ctx context.Context, reqCtx *types.RequestContext, fileID string, opts ListOptions) (*ListResult, error) {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, fileID)

	// Get file metadata first to check capabilities
	fileCall := m.client.Service().Files.Get(fileID)
	fileCall = m.shaper.ShapeFilesGet(fileCall, reqCtx)
	fileCall = fileCall.Fields("id,capabilities")

	file, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.File, error) {
		return fileCall.Do()
	})
	if err != nil {
		return nil, err
	}

	// Check canReadRevisions capability
	if file.Capabilities != nil && !file.Capabilities.CanReadRevisions {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodePermissionDenied,
			"Cannot read revisions for this file").
			WithContext("capability", "canReadRevisions=false").
			WithContext("fileId", fileID).
			Build())
	}

	// List revisions
	call := m.client.Service().Revisions.List(fileID)
	call = m.shaper.ShapeRevisionsList(call, reqCtx)

	if opts.PageSize > 0 {
		call = call.PageSize(int64(opts.PageSize))
	}
	if opts.PageToken != "" {
		call = call.PageToken(opts.PageToken)
	}
	if opts.Fields != "" {
		call = call.Fields(googleapi.Field("nextPageToken,revisions(" + opts.Fields + ")"))
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.RevisionList, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	revisions := make([]*types.Revision, len(result.Revisions))
	for i, r := range result.Revisions {
		revisions[i] = convertRevision(r)
	}

	return &ListResult{
		Revisions:     revisions,
		NextPageToken: result.NextPageToken,
	}, nil
}

// Get retrieves a specific revision
func (m *Manager) Get(ctx context.Context, reqCtx *types.RequestContext, fileID string, revisionID string) (*types.Revision, error) {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, fileID)

	call := m.client.Service().Revisions.Get(fileID, revisionID)
	call = m.shaper.ShapeRevisionsGet(call, reqCtx)

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.Revision, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	return convertRevision(result), nil
}

// DownloadOptions configures revision download
type DownloadOptions struct {
	OutputPath       string
	AcknowledgeAbuse bool
}

// Download downloads a specific revision
func (m *Manager) Download(ctx context.Context, reqCtx *types.RequestContext, fileID string, revisionID string, opts DownloadOptions) error {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, fileID)

	// Get revision metadata first
	revision, err := m.Get(ctx, reqCtx, fileID, revisionID)
	if err != nil {
		return err
	}

	// Check if revision is marked keepForever for blob files
	// For blob files (non-Workspace), revisions must be kept forever to be downloadable
	if !revision.KeepForever {
		return utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidArgument,
			"Revision must be marked keepForever=true before downloading").
			WithContext("revisionId", revisionID).
			WithContext("fileId", fileID).
			WithContext("suggestedAction", "use update command to set keepForever=true first").
			Build())
	}

	outputPath := opts.OutputPath
	if outputPath == "" {
		if revision.OriginalFilename != "" {
			outputPath = fmt.Sprintf("%s.rev%s", revision.OriginalFilename, revisionID)
		} else {
			outputPath = fmt.Sprintf("file.rev%s", revisionID)
		}
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidArgument,
			fmt.Sprintf("Failed to create output file: %s", err)).Build())
	}
	defer outFile.Close()

	// Download revision content
	call := m.client.Service().Revisions.Get(fileID, revisionID)
	call = m.shaper.ShapeRevisionsGet(call, reqCtx)
	if opts.AcknowledgeAbuse {
		call = call.AcknowledgeAbuse(true)
	}

	resp, err := call.Download()
	if err != nil {
		if apiErr, ok := err.(*googleapi.Error); ok {
			return utils.NewAppError(utils.NewCLIError(utils.ErrCodeNetworkError,
				fmt.Sprintf("Download failed: %s", apiErr.Message)).
				WithHTTPStatus(apiErr.Code).
				Build())
		}
		return utils.NewAppError(utils.NewCLIError(utils.ErrCodeNetworkError,
			fmt.Sprintf("Download failed: %s", err)).Build())
	}
	defer resp.Body.Close()

	_, err = io.Copy(outFile, resp.Body)
	return err
}

// UpdateOptions configures revision update
type UpdateOptions struct {
	KeepForever bool
}

// Update updates revision metadata
func (m *Manager) Update(ctx context.Context, reqCtx *types.RequestContext, fileID string, revisionID string, opts UpdateOptions) (*types.Revision, error) {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, fileID)

	metadata := &drive.Revision{
		KeepForever: opts.KeepForever,
	}

	call := m.client.Service().Revisions.Update(fileID, revisionID, metadata)
	call = m.shaper.ShapeRevisionsUpdate(call, reqCtx)

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.Revision, error) {
		return call.Do()
	})
	if err != nil {
		// Check for keepForever limit error
		if apiErr, ok := err.(*googleapi.Error); ok {
			if apiErr.Code == 403 {
				for _, e := range apiErr.Errors {
					if e.Reason == "revisionLimitExceeded" {
						return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeResourceLimit,
							"Cannot mark revision as keepForever: 200 revision limit reached").
							WithHTTPStatus(403).
							WithDriveReason("revisionLimitExceeded").
							WithContext("fileId", fileID).
							WithContext("revisionId", revisionID).
							WithContext("limit", 200).
							Build())
					}
				}
			}
		}
		return nil, err
	}

	return convertRevision(result), nil
}

// Restore restores a file to a specific revision
func (m *Manager) Restore(ctx context.Context, reqCtx *types.RequestContext, fileID string, revisionID string) (*types.DriveFile, error) {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, fileID)

	// For blob files, we need to download the revision and upload as new head
	// Get revision metadata
	revision, err := m.Get(ctx, reqCtx, fileID, revisionID)
	if err != nil {
		return nil, err
	}

	// Ensure revision is kept forever before downloading
	if !revision.KeepForever {
		// Set keepForever first
		_, err = m.Update(ctx, reqCtx, fileID, revisionID, UpdateOptions{KeepForever: true})
		if err != nil {
			return nil, fmt.Errorf("failed to mark revision as keepForever: %w", err)
		}
	}

	// Create temporary file for download
	tmpFile, err := os.CreateTemp("", "gdrv-restore-*")
	if err != nil {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeInternalError,
			fmt.Sprintf("Failed to create temporary file: %s", err)).Build())
	}
	tmpPath := tmpFile.Name()
	tmpFile.Close()
	defer os.Remove(tmpPath)

	// Download revision
	err = m.Download(ctx, reqCtx, fileID, revisionID, DownloadOptions{
		OutputPath: tmpPath,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to download revision: %w", err)
	}

	// Upload revision content as new head revision
	contentFile, err := os.Open(tmpPath)
	if err != nil {
		return nil, err
	}
	defer contentFile.Close()

	metadata := &drive.File{}
	call := m.client.Service().Files.Update(fileID, metadata).Media(contentFile)
	call = m.shaper.ShapeFilesUpdate(call, reqCtx)

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.File, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	return convertDriveFile(result), nil
}

func convertRevision(r *drive.Revision) *types.Revision {
	return &types.Revision{
		ID:               r.Id,
		ModifiedTime:     r.ModifiedTime,
		KeepForever:      r.KeepForever,
		Size:             r.Size,
		MimeType:         r.MimeType,
		OriginalFilename: r.OriginalFilename,
	}
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
