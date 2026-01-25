package cli

import (
	"github.com/dl-alexandre/gdrv/internal/types"
	"google.golang.org/api/drive/v3"
)

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
