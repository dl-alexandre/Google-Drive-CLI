package diff

import (
	"github.com/dl-alexandre/gdrv/internal/sync/index"
)

func ApplyRenames(result Result, snapshot Snapshot) Result {
	result = applyLocalRenames(result, snapshot)
	result = applyRemoteRenames(result, snapshot)
	return result
}

func applyLocalRenames(result Result, snapshot Snapshot) Result {
	deletedByHash := make(map[string][]index.SyncEntry)
	for path, prev := range snapshot.Prev {
		if prev.IsDir {
			continue
		}
		if _, ok := snapshot.Local[path]; ok {
			continue
		}
		if prev.ContentHash == "" {
			continue
		}
		deletedByHash[prev.ContentHash] = append(deletedByHash[prev.ContentHash], prev)
	}

	used := make(map[string]bool)
	for path, local := range snapshot.Local {
		if local.IsDir {
			continue
		}
		if _, ok := snapshot.Prev[path]; ok {
			continue
		}
		if local.Hash == "" {
			continue
		}
		candidates := deletedByHash[local.Hash]
		if len(candidates) != 1 {
			continue
		}
		prev := candidates[0]
		if used[prev.RelativePath] {
			continue
		}
		used[prev.RelativePath] = true

		result.Actions = removeAction(result.Actions, ActionDeleteRemote, prev.RelativePath)
		result.Actions = removeAction(result.Actions, ActionUpload, path)
		result.Actions = append(result.Actions, Action{
			Type:     ActionMoveRemote,
			FromPath: prev.RelativePath,
			ToPath:   path,
			Path:     path,
			Local:    &local,
			Prev:     &prev,
		})
	}

	return result
}

func applyRemoteRenames(result Result, snapshot Snapshot) Result {
	prevByID := make(map[string]index.SyncEntry)
	for _, prev := range snapshot.Prev {
		if prev.DriveFileID != "" {
			prevByID[prev.DriveFileID] = prev
		}
	}

	for path, remote := range snapshot.Remote {
		prev, ok := prevByID[remote.ID]
		if !ok {
			continue
		}
		if prev.RelativePath == path {
			continue
		}
		result.Actions = removeAction(result.Actions, ActionDeleteLocal, prev.RelativePath)
		result.Actions = removeAction(result.Actions, ActionDownload, path)
		result.Actions = removeAction(result.Actions, ActionMkdirLocal, path)
		result.Actions = append(result.Actions, Action{
			Type:     ActionMoveLocal,
			FromPath: prev.RelativePath,
			ToPath:   path,
			Path:     path,
			Remote:   &remote,
			Prev:     &prev,
		})
	}

	return result
}

func removeAction(actions []Action, actionType ActionType, path string) []Action {
	filtered := actions[:0]
	for _, action := range actions {
		if action.Type == actionType && action.Path == path {
			continue
		}
		filtered = append(filtered, action)
	}
	return filtered
}
