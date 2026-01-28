package changes

import (
	"context"
	"fmt"
	"time"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/types"
	"github.com/dl-alexandre/gdrv/internal/utils"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

type Manager struct {
	client *api.Client
	shaper *api.RequestShaper
}

func NewManager(client *api.Client) *Manager {
	return &Manager{
		client: client,
		shaper: api.NewRequestShaper(client),
	}
}

func (m *Manager) GetStartPageToken(ctx context.Context, reqCtx *types.RequestContext, driveID string) (string, error) {
	call := m.client.Service().Changes.GetStartPageToken()

	if driveID != "" {
		call = call.DriveId(driveID)
		call = call.SupportsAllDrives(true)
		reqCtx.DriveID = driveID
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.StartPageToken, error) {
		return call.Do()
	})
	if err != nil {
		return "", err
	}

	return result.StartPageToken, nil
}

func (m *Manager) List(ctx context.Context, reqCtx *types.RequestContext, opts types.ListOptions) (*types.ChangeList, error) {
	if opts.PageToken == "" {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidArgument,
			"page-token is required for listing changes").Build())
	}

	call := m.client.Service().Changes.List(opts.PageToken)

	if opts.DriveID != "" {
		call = call.DriveId(opts.DriveID)
		reqCtx.DriveID = opts.DriveID
	}

	if opts.IncludeCorpusRemovals {
		call = call.IncludeCorpusRemovals(true)
	}

	if opts.IncludeItemsFromAllDrives {
		call = call.IncludeItemsFromAllDrives(true)
	}

	if opts.IncludePermissionsForView != "" {
		call = call.IncludePermissionsForView(opts.IncludePermissionsForView)
	}

	if opts.IncludeRemoved {
		call = call.IncludeRemoved(true)
	}

	if opts.RestrictToMyDrive {
		call = call.RestrictToMyDrive(true)
	}

	if opts.SupportsAllDrives {
		call = call.SupportsAllDrives(true)
	}

	if opts.Limit > 0 {
		call = call.PageSize(int64(opts.Limit))
	}

	if opts.Fields != "" {
		call = call.Fields(googleapi.Field(opts.Fields))
	}

	if opts.Spaces != "" {
		call = call.Spaces(opts.Spaces)
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.ChangeList, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	return convertChangeList(result), nil
}

func (m *Manager) Watch(ctx context.Context, reqCtx *types.RequestContext, pageToken string, webhookURL string, opts types.WatchOptions) (*types.Channel, error) {
	if pageToken == "" {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidArgument,
			"page-token is required for watching changes").Build())
	}

	if webhookURL == "" {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidArgument,
			"webhook-url is required for watching changes").Build())
	}

	channel := &drive.Channel{
		Id:      generateChannelID(),
		Type:    "web_hook",
		Address: webhookURL,
	}

	if opts.Token != "" {
		channel.Token = opts.Token
	}

	if opts.Expiration > 0 {
		channel.Expiration = opts.Expiration
	}

	call := m.client.Service().Changes.Watch(pageToken, channel)

	if opts.DriveID != "" {
		call = call.DriveId(opts.DriveID)
		reqCtx.DriveID = opts.DriveID
	}

	if opts.IncludeCorpusRemovals {
		call = call.IncludeCorpusRemovals(true)
	}

	if opts.IncludeItemsFromAllDrives {
		call = call.IncludeItemsFromAllDrives(true)
	}

	if opts.IncludePermissionsForView != "" {
		call = call.IncludePermissionsForView(opts.IncludePermissionsForView)
	}

	if opts.IncludeRemoved {
		call = call.IncludeRemoved(true)
	}

	if opts.RestrictToMyDrive {
		call = call.RestrictToMyDrive(true)
	}

	if opts.SupportsAllDrives {
		call = call.SupportsAllDrives(true)
	}

	if opts.Spaces != "" {
		call = call.Spaces(opts.Spaces)
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.Channel, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	return convertChannel(result), nil
}

func (m *Manager) Stop(ctx context.Context, reqCtx *types.RequestContext, channelID string, resourceID string) error {
	if channelID == "" {
		return utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidArgument,
			"channel-id is required for stopping a watch").Build())
	}

	if resourceID == "" {
		return utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidArgument,
			"resource-id is required for stopping a watch").Build())
	}

	channel := &drive.Channel{
		Id:         channelID,
		ResourceId: resourceID,
	}

	_, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*struct{}, error) {
		err := m.client.Service().Channels.Stop(channel).Do()
		return &struct{}{}, err
	})

	return err
}

func convertChangeList(apiList *drive.ChangeList) *types.ChangeList {
	changes := make([]types.Change, 0, len(apiList.Changes))
	for _, c := range apiList.Changes {
		changes = append(changes, convertChange(c))
	}

	return &types.ChangeList{
		Changes:           changes,
		NextPageToken:     apiList.NextPageToken,
		NewStartPageToken: apiList.NewStartPageToken,
	}
}

func convertChange(apiChange *drive.Change) types.Change {
	change := types.Change{
		ChangeType: apiChange.Type,
		FileID:     apiChange.FileId,
		Removed:    apiChange.Removed,
		DriveID:    apiChange.DriveId,
	}

	if apiChange.Time != "" {
		if t, err := parseTime(apiChange.Time); err == nil {
			change.Time = t
		}
	}

	if apiChange.File != nil {
		change.File = convertDriveFile(apiChange.File)
	}

	if apiChange.Drive != nil {
		change.Drive = convertSharedDrive(apiChange.Drive)
	}

	return change
}

func convertDriveFile(apiFile *drive.File) *types.DriveFile {
	file := &types.DriveFile{
		ID:             apiFile.Id,
		Name:           apiFile.Name,
		MimeType:       apiFile.MimeType,
		Trashed:        apiFile.Trashed,
		WebViewLink:    apiFile.WebViewLink,
		WebContentLink: apiFile.WebContentLink,
		Parents:        apiFile.Parents,
		MD5Checksum:    apiFile.Md5Checksum,
		ResourceKey:    apiFile.ResourceKey,
	}

	if apiFile.Size > 0 {
		file.Size = apiFile.Size
	}

	if apiFile.CreatedTime != "" {
		file.CreatedTime = apiFile.CreatedTime
	}

	if apiFile.ModifiedTime != "" {
		file.ModifiedTime = apiFile.ModifiedTime
	}

	if apiFile.ExportLinks != nil {
		file.ExportLinks = apiFile.ExportLinks
	}

	if apiFile.Capabilities != nil {
		file.Capabilities = &types.FileCapabilities{
			CanDownload:      apiFile.Capabilities.CanDownload,
			CanEdit:          apiFile.Capabilities.CanEdit,
			CanShare:         apiFile.Capabilities.CanShare,
			CanDelete:        apiFile.Capabilities.CanDelete,
			CanTrash:         apiFile.Capabilities.CanTrash,
			CanReadRevisions: apiFile.Capabilities.CanReadRevisions,
		}
	}

	return file
}

func convertSharedDrive(apiDrive *drive.Drive) *types.SharedDrive {
	drive := &types.SharedDrive{
		ID:                  apiDrive.Id,
		Name:                apiDrive.Name,
		ColorRgb:            apiDrive.ColorRgb,
		BackgroundImageLink: apiDrive.BackgroundImageLink,
		Hidden:              apiDrive.Hidden,
		ThemeID:             apiDrive.ThemeId,
	}

	if apiDrive.CreatedTime != "" {
		if t, err := parseTime(apiDrive.CreatedTime); err == nil {
			drive.CreatedTime = t
		}
	}

	if apiDrive.Capabilities != nil {
		drive.Capabilities = &types.DriveCapabilities{
			CanAddChildren:   apiDrive.Capabilities.CanAddChildren,
			CanComment:       apiDrive.Capabilities.CanComment,
			CanCopy:          apiDrive.Capabilities.CanCopy,
			CanDeleteDrive:   apiDrive.Capabilities.CanDeleteDrive,
			CanDownload:      apiDrive.Capabilities.CanDownload,
			CanEdit:          apiDrive.Capabilities.CanEdit,
			CanListChildren:  apiDrive.Capabilities.CanListChildren,
			CanManageMembers: apiDrive.Capabilities.CanManageMembers,
			CanReadRevisions: apiDrive.Capabilities.CanReadRevisions,
			CanRename:        apiDrive.Capabilities.CanRename,
			CanRenameDrive:   apiDrive.Capabilities.CanRenameDrive,
			CanShare:         apiDrive.Capabilities.CanShare,
			CanTrashChildren: apiDrive.Capabilities.CanTrashChildren,
		}
	}

	return drive
}

func convertChannel(apiChannel *drive.Channel) *types.Channel {
	channel := &types.Channel{
		ID:          apiChannel.Id,
		ResourceID:  apiChannel.ResourceId,
		ResourceURI: apiChannel.ResourceUri,
		Token:       apiChannel.Token,
		Type:        apiChannel.Type,
		Address:     apiChannel.Address,
	}

	if apiChannel.Expiration > 0 {
		channel.Expiration = apiChannel.Expiration
	}

	if apiChannel.Params != nil {
		channel.Params = apiChannel.Params
	}

	return channel
}

func parseTime(timeStr string) (time.Time, error) {
	return time.Parse(time.RFC3339, timeStr)
}

func generateChannelID() string {
	return fmt.Sprintf("gdrv-changes-%d", time.Now().UnixNano())
}
