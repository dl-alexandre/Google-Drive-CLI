package executor

import (
	"context"
	"errors"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/dl-alexandre/gdrv/internal/files"
	"github.com/dl-alexandre/gdrv/internal/folders"
	"github.com/dl-alexandre/gdrv/internal/safety"
	"github.com/dl-alexandre/gdrv/internal/sync/diff"
	"github.com/dl-alexandre/gdrv/internal/sync/scanner"
	"github.com/dl-alexandre/gdrv/internal/types"
	"google.golang.org/api/drive/v3"
)

type Executor struct {
	files   *files.Manager
	folders *folders.Manager
}

type Options struct {
	Concurrency int
	DryRun      bool
	Force       bool
	Yes         bool
}

type State struct {
	LocalRoot    string
	RemoteRootID string
	LocalEntries map[string]scanner.LocalEntry
	RemoteEntries map[string]scanner.RemoteEntry
}

type Summary struct {
	Uploads   int
	Updates   int
	Downloads int
	Deletes   int
	Moves     int
	Mkdirs    int
}

func New(filesMgr *files.Manager, foldersMgr *folders.Manager) *Executor {
	return &Executor{
		files:   filesMgr,
		folders: foldersMgr,
	}
}

func (e *Executor) Apply(ctx context.Context, reqCtx *types.RequestContext, actions []diff.Action, state State, opts Options) (State, Summary, error) {
	summary := Summary{}
	if opts.DryRun {
		for _, action := range actions {
			summary = addSummary(summary, action.Type)
		}
		return state, summary, nil
	}

	remoteFolders := make(map[string]string)
	remoteFolders[""] = state.RemoteRootID
	for pathKey, entry := range state.RemoteEntries {
		if entry.IsDir {
			remoteFolders[pathKey] = entry.ID
		}
	}

	var mkdirRemote, mkdirLocal, moveRemote, moveLocal, uploads, updates, downloads, deleteRemote, deleteLocal []diff.Action
	for _, action := range actions {
		switch action.Type {
		case diff.ActionMkdirRemote:
			mkdirRemote = append(mkdirRemote, action)
		case diff.ActionMkdirLocal:
			mkdirLocal = append(mkdirLocal, action)
		case diff.ActionMoveRemote:
			moveRemote = append(moveRemote, action)
		case diff.ActionMoveLocal:
			moveLocal = append(moveLocal, action)
		case diff.ActionUpload:
			uploads = append(uploads, action)
		case diff.ActionUpdate:
			updates = append(updates, action)
		case diff.ActionDownload:
			downloads = append(downloads, action)
		case diff.ActionDeleteRemote:
			deleteRemote = append(deleteRemote, action)
		case diff.ActionDeleteLocal:
			deleteLocal = append(deleteLocal, action)
		}
	}

	sortByDepth(mkdirRemote, true)
	for _, action := range mkdirRemote {
		if _, err := e.ensureRemoteFolder(ctx, reqCtx, remoteFolders, action.Path); err != nil {
			return state, summary, err
		}
		summary = addSummary(summary, action.Type)
	}

	sortByDepth(mkdirLocal, true)
	for _, action := range mkdirLocal {
		if err := e.ensureLocalDir(state.LocalRoot, action.Path); err != nil {
			return state, summary, err
		}
		if entry, ok := state.RemoteEntries[action.Path]; ok {
			state.LocalEntries[action.Path] = scanner.LocalEntry{
				RelativePath: action.Path,
				AbsPath:      filepath.Join(state.LocalRoot, action.Path),
				IsDir:        true,
			}
			state.RemoteEntries[action.Path] = entry
		}
		summary = addSummary(summary, action.Type)
	}

	for _, action := range moveRemote {
		if err := e.applyMoveRemote(ctx, reqCtx, remoteFolders, state, action, opts); err != nil {
			return state, summary, err
		}
		summary = addSummary(summary, action.Type)
	}

	for _, action := range moveLocal {
		if err := e.applyMoveLocal(state.LocalRoot, state.LocalEntries, action); err != nil {
			return state, summary, err
		}
		summary = addSummary(summary, action.Type)
	}

	for _, action := range uploads {
		parentPath := path.Dir(action.Path)
		if parentPath == "." {
			parentPath = ""
		}
		if _, err := e.ensureRemoteFolder(ctx, reqCtx, remoteFolders, parentPath); err != nil {
			return state, summary, err
		}
	}

	for _, action := range updates {
		parentPath := path.Dir(action.Path)
		if parentPath == "." {
			parentPath = ""
		}
		if _, err := e.ensureRemoteFolder(ctx, reqCtx, remoteFolders, parentPath); err != nil {
			return state, summary, err
		}
	}

	transferMutex := &sync.Mutex{}

	if err := runConcurrent(ctx, uploads, opts.Concurrency, func(action diff.Action) error {
		localEntry := resolveLocalEntry(state.LocalEntries, action.Path, action.Local)
		if localEntry == nil {
			return nil
		}
		parentPath := path.Dir(action.Path)
		if parentPath == "." {
			parentPath = ""
		}
		parentID := remoteFolders[parentPath]
		result, err := e.files.Upload(ctx, reqCtx, localEntry.AbsPath, files.UploadOptions{
			ParentID: parentID,
			Name:     action.Name,
		})
		if err != nil {
			return err
		}
		transferMutex.Lock()
		state.RemoteEntries[action.Path] = scanner.RemoteEntry{
			RelativePath: action.Path,
			ID:           result.ID,
			ParentID:     parentID,
			IsDir:        result.MimeType == "application/vnd.google-apps.folder",
			Size:         result.Size,
			ModifiedTime: result.ModifiedTime,
			MD5Checksum:  result.MD5Checksum,
			MimeType:     result.MimeType,
		}
		transferMutex.Unlock()
		return nil
	}); err != nil {
		return state, summary, err
	}
	for range uploads {
		summary = addSummary(summary, diff.ActionUpload)
	}

	if err := runConcurrent(ctx, updates, opts.Concurrency, func(action diff.Action) error {
		localEntry := resolveLocalEntry(state.LocalEntries, action.Path, action.Local)
		remoteEntry := resolveRemoteEntry(state.RemoteEntries, action.Path, action.Remote)
		if localEntry == nil || remoteEntry == nil {
			return nil
		}
		result, err := e.files.UpdateContent(ctx, reqCtx, remoteEntry.ID, localEntry.AbsPath, files.UpdateContentOptions{})
		if err != nil {
			return err
		}
		transferMutex.Lock()
		state.RemoteEntries[action.Path] = scanner.RemoteEntry{
			RelativePath: action.Path,
			ID:           result.ID,
			ParentID:     remoteEntry.ParentID,
			IsDir:        result.MimeType == "application/vnd.google-apps.folder",
			Size:         result.Size,
			ModifiedTime: result.ModifiedTime,
			MD5Checksum:  result.MD5Checksum,
			MimeType:     result.MimeType,
		}
		transferMutex.Unlock()
		return nil
	}); err != nil {
		return state, summary, err
	}
	for range updates {
		summary = addSummary(summary, diff.ActionUpdate)
	}

	for _, action := range downloads {
		if err := e.ensureLocalDir(state.LocalRoot, path.Dir(action.Path)); err != nil {
			return state, summary, err
		}
	}

	if err := runConcurrent(ctx, downloads, opts.Concurrency, func(action diff.Action) error {
		remoteEntry := resolveRemoteEntry(state.RemoteEntries, action.Path, action.Remote)
		if remoteEntry == nil {
			return nil
		}
		absPath := filepath.Join(state.LocalRoot, action.Path)
		err := e.files.Download(ctx, reqCtx, remoteEntry.ID, files.DownloadOptions{
			OutputPath: absPath,
		})
		if err != nil {
			return err
		}
		transferMutex.Lock()
		info, statErr := os.Stat(absPath)
		if statErr == nil {
			modTime := info.ModTime().Unix()
			if remoteEntry.ModifiedTime != "" {
				if parsed, err := time.Parse(time.RFC3339, remoteEntry.ModifiedTime); err == nil {
					_ = os.Chtimes(absPath, parsed, parsed)
					modTime = parsed.Unix()
				}
			}
			state.LocalEntries[action.Path] = scanner.LocalEntry{
				RelativePath: action.Path,
				AbsPath:      absPath,
				IsDir:        false,
				Size:         info.Size(),
				ModTime:      modTime,
			}
		}
		transferMutex.Unlock()
		return nil
	}); err != nil {
		return state, summary, err
	}
	for range downloads {
		summary = addSummary(summary, diff.ActionDownload)
	}

	sortByDepth(deleteLocal, false)
	if len(deleteLocal) > 0 && !opts.DryRun {
		items := make([]string, 0, len(deleteLocal))
		for _, action := range deleteLocal {
			items = append(items, action.Path)
		}
		safetyOpts := safety.Default()
		safetyOpts.DryRun = opts.DryRun
		safetyOpts.Force = opts.Force
		safetyOpts.Yes = opts.Yes
		safetyOpts.Interactive = !opts.Force && !opts.Yes
		confirmed, err := safety.ConfirmDestructive(items, "delete local files", safetyOpts)
		if err != nil {
			return state, summary, err
		}
		if !confirmed {
			return state, summary, errors.New("operation cancelled by user")
		}
	}
	for _, action := range deleteLocal {
		if err := e.deleteLocal(state.LocalRoot, state.LocalEntries, action); err != nil {
			return state, summary, err
		}
		summary = addSummary(summary, action.Type)
	}

	sortByDepth(deleteRemote, false)
	for _, action := range deleteRemote {
		if err := e.deleteRemote(ctx, reqCtx, state.RemoteEntries, action, opts); err != nil {
			return state, summary, err
		}
		summary = addSummary(summary, action.Type)
	}

	return state, summary, nil
}

func runConcurrent(ctx context.Context, actions []diff.Action, concurrency int, handler func(diff.Action) error) error {
	if len(actions) == 0 {
		return nil
	}
	if concurrency <= 0 {
		concurrency = 1
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	jobs := make(chan diff.Action)
	errs := make(chan error, len(actions))
	var wg sync.WaitGroup

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for action := range jobs {
				if ctx.Err() != nil {
					continue
				}
				if err := handler(action); err != nil {
					errs <- err
					cancel()
				}
			}
		}()
	}

	for _, action := range actions {
		jobs <- action
	}
	close(jobs)
	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func addSummary(summary Summary, actionType diff.ActionType) Summary {
	switch actionType {
	case diff.ActionUpload:
		summary.Uploads++
	case diff.ActionUpdate:
		summary.Updates++
	case diff.ActionDownload:
		summary.Downloads++
	case diff.ActionDeleteLocal, diff.ActionDeleteRemote:
		summary.Deletes++
	case diff.ActionMoveLocal, diff.ActionMoveRemote:
		summary.Moves++
	case diff.ActionMkdirLocal, diff.ActionMkdirRemote:
		summary.Mkdirs++
	}
	return summary
}

func resolveLocalEntry(localEntries map[string]scanner.LocalEntry, path string, entry *scanner.LocalEntry) *scanner.LocalEntry {
	if entry != nil {
		return entry
	}
	if resolved, ok := localEntries[path]; ok {
		value := resolved
		return &value
	}
	return nil
}

func resolveRemoteEntry(remoteEntries map[string]scanner.RemoteEntry, path string, entry *scanner.RemoteEntry) *scanner.RemoteEntry {
	if entry != nil {
		return entry
	}
	if resolved, ok := remoteEntries[path]; ok {
		value := resolved
		return &value
	}
	return nil
}

func sortByDepth(actions []diff.Action, ascending bool) {
	sort.Slice(actions, func(i, j int) bool {
		di := depth(actions[i].Path)
		dj := depth(actions[j].Path)
		if ascending {
			return di < dj
		}
		return di > dj
	})
}

func depth(p string) int {
	if p == "" {
		return 0
	}
	return strings.Count(p, "/") + 1
}

func (e *Executor) ensureRemoteFolder(ctx context.Context, reqCtx *types.RequestContext, remoteFolders map[string]string, relPath string) (string, error) {
	if relPath == "" || relPath == "." {
		return remoteFolders[""], nil
	}
	if id, ok := remoteFolders[relPath]; ok {
		return id, nil
	}
	parent := path.Dir(relPath)
	if parent == "." {
		parent = ""
	}
	parentID, err := e.ensureRemoteFolder(ctx, reqCtx, remoteFolders, parent)
	if err != nil {
		return "", err
	}
	name := path.Base(relPath)
	result, err := e.folders.Create(ctx, reqCtx, name, parentID)
	if err != nil {
		return "", err
	}
	remoteFolders[relPath] = result.ID
	return result.ID, nil
}

func (e *Executor) ensureLocalDir(root, relPath string) error {
	if relPath == "" || relPath == "." {
		return nil
	}
	return os.MkdirAll(filepath.Join(root, relPath), 0700)
}

func (e *Executor) applyMoveLocal(root string, localEntries map[string]scanner.LocalEntry, action diff.Action) error {
	from := filepath.Join(root, action.FromPath)
	to := filepath.Join(root, action.ToPath)
	if err := os.MkdirAll(filepath.Dir(to), 0700); err != nil {
		return err
	}
	if err := os.Rename(from, to); err != nil {
		return err
	}
	fromPrefix := action.FromPath + "/"
	toPrefix := action.ToPath + "/"
	updates := make(map[string]scanner.LocalEntry)
	for key, entry := range localEntries {
		if key == action.FromPath {
			entry.RelativePath = action.ToPath
			entry.AbsPath = to
			updates[action.ToPath] = entry
			continue
		}
		if strings.HasPrefix(key, fromPrefix) {
			newKey := toPrefix + strings.TrimPrefix(key, fromPrefix)
			entry.RelativePath = newKey
			entry.AbsPath = filepath.Join(root, newKey)
			updates[newKey] = entry
			continue
		}
		updates[key] = entry
	}
	for key := range localEntries {
		delete(localEntries, key)
	}
	for key, entry := range updates {
		localEntries[key] = entry
	}
	return nil
}

func (e *Executor) applyMoveRemote(ctx context.Context, reqCtx *types.RequestContext, remoteFolders map[string]string, state State, action diff.Action, opts Options) error {
	prev := action.Prev
	fileID := ""
	currentParentID := ""
	if prev != nil && prev.DriveFileID != "" {
		fileID = prev.DriveFileID
		currentParentID = prev.DriveParentID
	} else if action.Remote != nil {
		fileID = action.Remote.ID
		currentParentID = action.Remote.ParentID
	}
	if fileID == "" {
		return nil
	}
	newParentPath := path.Dir(action.ToPath)
	if newParentPath == "." {
		newParentPath = ""
	}
	newParentID, err := e.ensureRemoteFolder(ctx, reqCtx, remoteFolders, newParentPath)
	if err != nil {
		return err
	}

	safetyOpts := safety.Default()
	safetyOpts.DryRun = opts.DryRun
	safetyOpts.Force = opts.Force
	safetyOpts.Yes = opts.Yes

	if currentParentID != newParentID {
		_, err := e.files.MoveWithSafety(ctx, reqCtx, fileID, newParentID, safetyOpts, nil)
		if err != nil {
			return err
		}
	}
	newName := path.Base(action.ToPath)
	if newName != "" {
		_, err := e.files.Update(ctx, reqCtx, fileID, &drive.File{Name: newName}, "")
		if err != nil {
			return err
		}
	}
	if entry, ok := state.RemoteEntries[action.FromPath]; ok {
		entry.RelativePath = action.ToPath
		entry.ParentID = newParentID
		delete(state.RemoteEntries, action.FromPath)
		state.RemoteEntries[action.ToPath] = entry
		if entry.IsDir {
			delete(remoteFolders, action.FromPath)
			remoteFolders[action.ToPath] = entry.ID
			fromPrefix := action.FromPath + "/"
			toPrefix := action.ToPath + "/"
			updates := make(map[string]scanner.RemoteEntry)
			var deletes []string
			for key, child := range state.RemoteEntries {
				if strings.HasPrefix(key, fromPrefix) {
					newKey := toPrefix + strings.TrimPrefix(key, fromPrefix)
					child.RelativePath = newKey
					updates[newKey] = child
					deletes = append(deletes, key)
				}
			}
			for _, key := range deletes {
				delete(state.RemoteEntries, key)
			}
			for key, child := range updates {
				state.RemoteEntries[key] = child
			}
		}
	}
	return nil
}

func (e *Executor) deleteLocal(root string, localEntries map[string]scanner.LocalEntry, action diff.Action) error {
	target := filepath.Join(root, action.Path)
	if entry, ok := localEntries[action.Path]; ok && entry.IsDir {
		if err := os.RemoveAll(target); err != nil {
			return err
		}
	} else {
		if err := os.Remove(target); err != nil && !os.IsNotExist(err) {
			return err
		}
	}
	delete(localEntries, action.Path)
	return nil
}

func (e *Executor) deleteRemote(ctx context.Context, reqCtx *types.RequestContext, remoteEntries map[string]scanner.RemoteEntry, action diff.Action, opts Options) error {
	entry, ok := remoteEntries[action.Path]
	if !ok && action.Remote != nil {
		entry = *action.Remote
	}
	if entry.ID == "" && action.Prev != nil {
		entry.ID = action.Prev.DriveFileID
		entry.IsDir = action.Prev.IsDir
	}
	if entry.ID == "" {
		return nil
	}
	safetyOpts := safety.Default()
	safetyOpts.DryRun = opts.DryRun
	safetyOpts.Force = opts.Force
	safetyOpts.Yes = opts.Yes

	if entry.IsDir {
		if err := e.folders.DeleteWithSafety(ctx, reqCtx, entry.ID, true, safetyOpts, nil); err != nil {
			return err
		}
	} else {
		if err := e.files.DeleteWithSafety(ctx, reqCtx, entry.ID, false, safetyOpts, nil); err != nil {
			return err
		}
	}
	delete(remoteEntries, action.Path)
	return nil
}
