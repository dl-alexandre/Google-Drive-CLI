package scanner

import (
	"context"
	"path"
	"strings"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/sync/index"
	"github.com/dl-alexandre/gdrv/internal/types"
	"github.com/dl-alexandre/gdrv/internal/utils"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

type RemoteScanner struct {
	client *api.Client
	shaper *api.RequestShaper
}

func NewRemoteScanner(client *api.Client) *RemoteScanner {
	return &RemoteScanner{
		client: client,
		shaper: api.NewRequestShaper(client),
	}
}

func (s *RemoteScanner) ListTree(ctx context.Context, reqCtx *types.RequestContext, rootID string) (map[string]RemoteEntry, error) {
	entries := make(map[string]RemoteEntry)
	queue := []remoteNode{{ID: rootID, Path: ""}}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		children, err := s.listChildren(ctx, reqCtx, node.ID)
		if err != nil {
			return nil, err
		}

		for _, child := range children {
			rel := child.Name
			if node.Path != "" {
				rel = path.Join(node.Path, child.Name)
			}
			entry := RemoteEntry{
				RelativePath: rel,
				ID:           child.ID,
				ParentID:     node.ID,
				IsDir:        child.MimeType == utils.MimeTypeFolder,
				Size:         child.Size,
				ModifiedTime: child.ModifiedTime,
				MD5Checksum:  child.MD5Checksum,
				MimeType:     child.MimeType,
			}
			entries[rel] = entry
			if entry.IsDir {
				queue = append(queue, remoteNode{ID: child.ID, Path: rel})
			}
		}
	}

	return entries, nil
}

func (s *RemoteScanner) ListTreeWithChanges(ctx context.Context, reqCtx *types.RequestContext, rootID, changeToken string, prevEntries []index.SyncEntry) (map[string]RemoteEntry, string, bool, error) {
	entries := make(map[string]RemoteEntry)
	fileIDToPath := make(map[string]string)
	infoCache := make(map[string]parentInfo)

	for _, e := range prevEntries {
		remote := RemoteEntry{
			RelativePath: e.RelativePath,
			ID:           e.DriveFileID,
			ParentID:     e.DriveParentID,
			IsDir:        e.IsDir,
			Size:         e.RemoteSize,
			ModifiedTime: e.RemoteMTime,
			MD5Checksum:  e.RemoteMD5,
			MimeType:     e.RemoteMimeType,
		}
		entries[e.RelativePath] = remote
		if e.DriveFileID != "" {
			fileIDToPath[e.DriveFileID] = e.RelativePath
		}
	}

	needsFullScan := false
	newToken := ""
	pageToken := changeToken

	for {
		call := s.client.Service().Changes.List(pageToken)
		call = call.SupportsAllDrives(true).IncludeItemsFromAllDrives(true)
		if reqCtx.DriveID != "" {
			call = call.DriveId(reqCtx.DriveID)
		}
		call = call.Fields("nextPageToken,newStartPageToken,changes(fileId,removed,file(id,name,mimeType,size,modifiedTime,md5Checksum,parents,resourceKey))")

		result, err := api.ExecuteWithRetry(ctx, s.client, reqCtx, func() (*drive.ChangeList, error) {
			return call.Do()
		})
		if err != nil {
			return nil, "", false, err
		}

		for _, change := range result.Changes {
			if change.Removed || change.File == nil {
				if pathValue, ok := fileIDToPath[change.FileId]; ok {
					removeDescendants(entries, fileIDToPath, pathValue)
				}
				continue
			}

			parentID := ""
			if len(change.File.Parents) > 0 {
				parentID = change.File.Parents[0]
			}
			if parentID == "" {
				if pathValue, ok := fileIDToPath[change.File.Id]; ok {
					removeDescendants(entries, fileIDToPath, pathValue)
				}
				continue
			}

			parentPath, underRoot, err := s.resolveParentPath(ctx, reqCtx, rootID, parentID, fileIDToPath, infoCache, 0)
			if err != nil {
				needsFullScan = true
				break
			}
			if !underRoot {
				if pathValue, ok := fileIDToPath[change.File.Id]; ok {
					removeDescendants(entries, fileIDToPath, pathValue)
				}
				continue
			}

			rel := change.File.Name
			if parentPath != "" {
				rel = path.Join(parentPath, change.File.Name)
			}
			if parentPath != "" {
				fileIDToPath[parentID] = parentPath
			}

			if oldPath, ok := fileIDToPath[change.File.Id]; ok && oldPath != rel {
				if change.File.MimeType == utils.MimeTypeFolder {
					rekeyDescendants(entries, fileIDToPath, oldPath, rel)
					delete(entries, oldPath)
				} else {
					delete(entries, oldPath)
				}
			}

			entry := RemoteEntry{
				RelativePath: rel,
				ID:           change.File.Id,
				ParentID:     parentID,
				IsDir:        change.File.MimeType == utils.MimeTypeFolder,
				Size:         change.File.Size,
				ModifiedTime: change.File.ModifiedTime,
				MD5Checksum:  change.File.Md5Checksum,
				MimeType:     change.File.MimeType,
			}

			entries[rel] = entry
			fileIDToPath[change.File.Id] = rel
			if change.File.ResourceKey != "" {
				s.client.ResourceKeys().UpdateFromAPIResponse(change.File.Id, change.File.ResourceKey)
			}
		}

		if needsFullScan {
			break
		}

		if result.NewStartPageToken != "" {
			newToken = result.NewStartPageToken
		}
		if result.NextPageToken == "" {
			break
		}
		pageToken = result.NextPageToken
	}

	if needsFullScan {
		fullEntries, err := s.ListTree(ctx, reqCtx, rootID)
		if err != nil {
			return nil, "", false, err
		}
		return fullEntries, newToken, true, nil
	}

	return entries, newToken, false, nil
}

func (s *RemoteScanner) GetStartPageToken(ctx context.Context, reqCtx *types.RequestContext) (string, error) {
	call := s.client.Service().Changes.GetStartPageToken()
	call = call.SupportsAllDrives(true)
	if reqCtx.DriveID != "" {
		call = call.DriveId(reqCtx.DriveID)
	}

	result, err := api.ExecuteWithRetry(ctx, s.client, reqCtx, func() (*drive.StartPageToken, error) {
		return call.Do()
	})
	if err != nil {
		return "", err
	}
	return result.StartPageToken, nil
}

type remoteNode struct {
	ID   string
	Path string
}

func (s *RemoteScanner) listChildren(ctx context.Context, reqCtx *types.RequestContext, parentID string) ([]*types.DriveFile, error) {
	query := "'" + parentID + "' in parents and trashed = false"
	call := s.client.Service().Files.List().Q(query)
	call = s.shaper.ShapeFilesList(call, reqCtx)
	call = call.Fields("nextPageToken,files(id,name,mimeType,size,modifiedTime,md5Checksum,parents,resourceKey)")

	var results []*types.DriveFile
	for {
		list, err := api.ExecuteWithRetry(ctx, s.client, reqCtx, func() (*drive.FileList, error) {
			return call.Do()
		})
		if err != nil {
			return nil, err
		}
		for _, f := range list.Files {
			file := &types.DriveFile{
				ID:          f.Id,
				Name:        f.Name,
				MimeType:    f.MimeType,
				Size:        f.Size,
				MD5Checksum: f.Md5Checksum,
				Parents:     f.Parents,
				ResourceKey: f.ResourceKey,
				ModifiedTime: f.ModifiedTime,
			}
			results = append(results, file)
			if f.ResourceKey != "" {
				s.client.ResourceKeys().UpdateFromAPIResponse(f.Id, f.ResourceKey)
			}
		}
		if list.NextPageToken == "" {
			break
		}
		call = call.PageToken(list.NextPageToken)
	}
	return results, nil
}

type parentInfo struct {
	Name     string
	Parents  []string
	MimeType string
}

func (s *RemoteScanner) resolveParentPath(ctx context.Context, reqCtx *types.RequestContext, rootID, parentID string, fileIDToPath map[string]string, cache map[string]parentInfo, depth int) (string, bool, error) {
	if parentID == rootID {
		return "", true, nil
	}
	if parentID == "" {
		return "", false, nil
	}
	if existing, ok := fileIDToPath[parentID]; ok {
		return existing, true, nil
	}
	if depth > 50 {
		return "", false, nil
	}
	info, err := s.getParentInfo(ctx, reqCtx, parentID, cache)
	if err != nil {
		return "", false, err
	}
	if len(info.Parents) == 0 {
		return "", false, nil
	}
	parentPath, ok, err := s.resolveParentPath(ctx, reqCtx, rootID, info.Parents[0], fileIDToPath, cache, depth+1)
	if err != nil || !ok {
		return "", ok, err
	}
	rel := info.Name
	if parentPath != "" {
		rel = path.Join(parentPath, info.Name)
	}
	fileIDToPath[parentID] = rel
	return rel, true, nil
}

func (s *RemoteScanner) getParentInfo(ctx context.Context, reqCtx *types.RequestContext, fileID string, cache map[string]parentInfo) (parentInfo, error) {
	if info, ok := cache[fileID]; ok {
		return info, nil
	}
	reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, fileID)
	call := s.client.Service().Files.Get(fileID)
	call = s.shaper.ShapeFilesGet(call, reqCtx)
	call = call.Fields(googleapi.Field("id,name,mimeType,parents,resourceKey"))
	result, err := api.ExecuteWithRetry(ctx, s.client, reqCtx, func() (*drive.File, error) {
		return call.Do()
	})
	if err != nil {
		return parentInfo{}, err
	}
	if result.ResourceKey != "" {
		s.client.ResourceKeys().UpdateFromAPIResponse(result.Id, result.ResourceKey)
	}
	info := parentInfo{
		Name:     result.Name,
		Parents:  result.Parents,
		MimeType: result.MimeType,
	}
	cache[fileID] = info
	return info, nil
}

func removeDescendants(entries map[string]RemoteEntry, fileIDToPath map[string]string, basePath string) {
	if basePath == "" {
		return
	}
	prefix := basePath + "/"
	for id, value := range fileIDToPath {
		if value == basePath || strings.HasPrefix(value, prefix) {
			delete(fileIDToPath, id)
		}
	}
	for key := range entries {
		if key == basePath || strings.HasPrefix(key, prefix) {
			delete(entries, key)
		}
	}
}

func rekeyDescendants(entries map[string]RemoteEntry, fileIDToPath map[string]string, oldPath, newPath string) {
	oldPrefix := oldPath + "/"
	newPrefix := newPath + "/"
	updates := make(map[string]RemoteEntry)
	for key, entry := range entries {
		if strings.HasPrefix(key, oldPrefix) {
			newKey := newPrefix + strings.TrimPrefix(key, oldPrefix)
			entry.RelativePath = newKey
			updates[newKey] = entry
			delete(entries, key)
		}
	}
	for key, entry := range updates {
		entries[key] = entry
	}
	for id, value := range fileIDToPath {
		if strings.HasPrefix(value, oldPrefix) {
			fileIDToPath[id] = newPrefix + strings.TrimPrefix(value, oldPrefix)
		}
	}
}
