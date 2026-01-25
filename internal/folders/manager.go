package folders

import (
	"context"
	"fmt"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/safety"
	"github.com/dl-alexandre/gdrv/internal/types"
	"github.com/dl-alexandre/gdrv/internal/utils"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

// Manager handles folder operations
type Manager struct {
	client *api.Client
	shaper *api.RequestShaper
}

// NewManager creates a new folder manager
func NewManager(client *api.Client) *Manager {
	return &Manager{
		client: client,
		shaper: api.NewRequestShaper(client),
	}
}

// Create creates a new folder
func (m *Manager) Create(ctx context.Context, reqCtx *types.RequestContext, name string, parentID string) (*types.DriveFile, error) {
	if parentID != "" {
		reqCtx.InvolvedParentIDs = append(reqCtx.InvolvedParentIDs, parentID)
	}

	metadata := &drive.File{
		Name:     name,
		MimeType: utils.MimeTypeFolder,
	}
	if parentID != "" {
		metadata.Parents = []string{parentID}
	}

	call := m.client.Service().Files.Create(metadata)
	call = m.shaper.ShapeFilesCreate(call, reqCtx)
	call = call.Fields("id,name,mimeType,size,createdTime,modifiedTime,parents,resourceKey,trashed,capabilities")

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.File, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	return convertDriveFile(result), nil
}

// List lists folder contents
func (m *Manager) List(ctx context.Context, reqCtx *types.RequestContext, folderID string, pageSize int, pageToken string) (*types.FileListResult, error) {
	reqCtx.InvolvedParentIDs = append(reqCtx.InvolvedParentIDs, folderID)

	query := fmt.Sprintf("'%s' in parents and trashed = false", folderID)

	call := m.client.Service().Files.List().Q(query)
	call = m.shaper.ShapeFilesList(call, reqCtx)
	call = call.Fields("nextPageToken,incompleteSearch,files(id,name,mimeType,size,createdTime,modifiedTime,parents,resourceKey,trashed,capabilities)")

	if pageSize > 0 {
		call = call.PageSize(int64(pageSize))
	}
	if pageToken != "" {
		call = call.PageToken(pageToken)
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
	}

	return &types.FileListResult{
		Files:            files,
		NextPageToken:    result.NextPageToken,
		IncompleteSearch: result.IncompleteSearch,
	}, nil
}

// Delete deletes a folder
func (m *Manager) Delete(ctx context.Context, reqCtx *types.RequestContext, folderID string, recursive bool) error {
	return m.DeleteWithSafety(ctx, reqCtx, folderID, recursive, safety.Default(), nil)
}

// DeleteWithSafety deletes a folder with safety controls.
// Supports dry-run mode, confirmation, and idempotency.
//
// Requirements:
//   - Requirement 13.1: Support --dry-run mode for destructive operations
//   - Requirement 13.2: Support --force flag to skip confirmations
//   - Requirement 13.3: Implement confirmation requirements for bulk operations
func (m *Manager) DeleteWithSafety(ctx context.Context, reqCtx *types.RequestContext, folderID string, recursive bool, opts safety.SafetyOptions, recorder safety.DryRunRecorder) error {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, folderID)

	// Get folder metadata for confirmation
	folder, err := m.Get(ctx, reqCtx, folderID, "id,name")
	if err != nil {
		return err
	}

	// If recursive, count contents for confirmation
	var contentCount int
	if recursive {
		contentCount, err = m.countContents(ctx, reqCtx, folderID)
		if err != nil {
			return err
		}
	}

	// Dry-run mode: record operation without executing
	if opts.DryRun && recorder != nil {
		safety.RecordDelete(recorder, folderID, folder.Name, true)
		if recursive {
			// In dry-run, we would recursively record all contents
			if err := m.deleteContentsWithSafety(ctx, reqCtx, folderID, opts, recorder); err != nil {
				return err
			}
		}
		return nil
	}

	// Confirmation for destructive operations
	if opts.ShouldConfirm() {
		if recursive && contentCount > 0 {
			confirmed, err := safety.Confirm(
				fmt.Sprintf("About to recursively delete folder '%s' containing %d items. Continue?", folder.Name, contentCount),
				opts,
			)
			if err != nil {
				return err
			}
			if !confirmed {
				return utils.NewAppError(utils.NewCLIError(utils.ErrCodeCancelled, "Operation cancelled by user").Build())
			}
		} else {
			confirmed, err := safety.Confirm(
				fmt.Sprintf("About to delete folder '%s'. Continue?", folder.Name),
				opts,
			)
			if err != nil {
				return err
			}
			if !confirmed {
				return utils.NewAppError(utils.NewCLIError(utils.ErrCodeCancelled, "Operation cancelled by user").Build())
			}
		}
	}

	if recursive {
		// List and delete all contents first
		if err := m.deleteContentsWithSafety(ctx, reqCtx, folderID, opts, recorder); err != nil {
			return err
		}
	}

	call := m.client.Service().Files.Delete(folderID)
	call = m.shaper.ShapeFilesDelete(call, reqCtx)

	_, err = api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (interface{}, error) {
		return nil, call.Do()
	})
	return err
}

func (m *Manager) deleteContentsWithSafety(ctx context.Context, reqCtx *types.RequestContext, folderID string, opts safety.SafetyOptions, recorder safety.DryRunRecorder) error {
	pageToken := ""
	for {
		result, err := m.List(ctx, reqCtx, folderID, 100, pageToken)
		if err != nil {
			return err
		}

		for _, file := range result.Files {
			if file.MimeType == utils.MimeTypeFolder {
				if err := m.deleteContentsWithSafety(ctx, reqCtx, file.ID, opts, recorder); err != nil {
					return err
				}
			}

			// Dry-run mode: record operation
			if opts.DryRun && recorder != nil {
				safety.RecordDelete(recorder, file.ID, file.Name, true)
				continue
			}

			// Add file ID to context for this deletion
			fileCtx := &types.RequestContext{
				Profile:           reqCtx.Profile,
				DriveID:           reqCtx.DriveID,
				InvolvedFileIDs:   []string{file.ID},
				InvolvedParentIDs: reqCtx.InvolvedParentIDs,
				RequestType:       reqCtx.RequestType,
				TraceID:           reqCtx.TraceID,
			}

			deleteCall := m.client.Service().Files.Delete(file.ID)
			deleteCall = m.shaper.ShapeFilesDelete(deleteCall, fileCtx)
			_, err := api.ExecuteWithRetry(ctx, m.client, fileCtx, func() (interface{}, error) {
				return nil, deleteCall.Do()
			})
			if err != nil {
				return err
			}
		}

		if result.NextPageToken == "" {
			break
		}
		pageToken = result.NextPageToken
	}
	return nil
}

func (m *Manager) countContents(ctx context.Context, reqCtx *types.RequestContext, folderID string) (int, error) {
	count := 0
	pageToken := ""
	for {
		result, err := m.List(ctx, reqCtx, folderID, 100, pageToken)
		if err != nil {
			return 0, err
		}

		for _, file := range result.Files {
			count++
			if file.MimeType == utils.MimeTypeFolder {
				subCount, err := m.countContents(ctx, reqCtx, file.ID)
				if err != nil {
					return 0, err
				}
				count += subCount
			}
		}

		if result.NextPageToken == "" {
			break
		}
		pageToken = result.NextPageToken
	}
	return count, nil
}

// Move moves a folder to a new parent
func (m *Manager) Move(ctx context.Context, reqCtx *types.RequestContext, folderID string, newParentID string) (*types.DriveFile, error) {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, folderID)
	reqCtx.InvolvedParentIDs = append(reqCtx.InvolvedParentIDs, newParentID)

	// Get current parents
	getCall := m.client.Service().Files.Get(folderID).Fields("parents")
	getCall = m.shaper.ShapeFilesGet(getCall, reqCtx)

	current, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.File, error) {
		return getCall.Do()
	})
	if err != nil {
		return nil, err
	}

	var removeParents string
	for i, p := range current.Parents {
		if i > 0 {
			removeParents += ","
		}
		removeParents += p
	}

	call := m.client.Service().Files.Update(folderID, &drive.File{})
	call = m.shaper.ShapeFilesUpdate(call, reqCtx)
	call = call.AddParents(newParentID)
	call = call.Fields("id,name,mimeType,size,createdTime,modifiedTime,parents,resourceKey,trashed,capabilities")
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

// Get retrieves folder metadata
func (m *Manager) Get(ctx context.Context, reqCtx *types.RequestContext, folderID string, fields string) (*types.DriveFile, error) {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, folderID)

	call := m.client.Service().Files.Get(folderID)
	call = m.shaper.ShapeFilesGet(call, reqCtx)
	if fields != "" {
		call = call.Fields(googleapi.Field(fields))
	} else {
		call = call.Fields("id,name,mimeType,size,createdTime,modifiedTime,parents,resourceKey,trashed,capabilities")
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.File, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	return convertDriveFile(result), nil
}

// Rename renames a folder
func (m *Manager) Rename(ctx context.Context, reqCtx *types.RequestContext, folderID string, newName string) (*types.DriveFile, error) {
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, folderID)

	call := m.client.Service().Files.Update(folderID, &drive.File{Name: newName})
	call = m.shaper.ShapeFilesUpdate(call, reqCtx)
	call = call.Fields("id,name,mimeType,size,createdTime,modifiedTime,parents,resourceKey,trashed,capabilities")

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.File, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	return convertDriveFile(result), nil
}

func convertDriveFile(f *drive.File) *types.DriveFile {
	file := &types.DriveFile{
		ID:           f.Id,
		Name:         f.Name,
		MimeType:     f.MimeType,
		Size:         f.Size,
		MD5Checksum:  f.Md5Checksum,
		CreatedTime:  f.CreatedTime,
		ModifiedTime: f.ModifiedTime,
		Parents:      f.Parents,
		ResourceKey:  f.ResourceKey,
		Trashed:      f.Trashed,
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
