package diff

import (
	"sort"

	"github.com/dl-alexandre/gdrv/internal/sync/index"
	"github.com/dl-alexandre/gdrv/internal/sync/scanner"
)

type Snapshot struct {
	Local  map[string]scanner.LocalEntry
	Remote map[string]scanner.RemoteEntry
	Prev   map[string]index.SyncEntry
}

type Result struct {
	Actions   []Action
	Conflicts []Conflict
}

func Compute(snapshot Snapshot, mode Mode, includeDeletes bool) Result {
	paths := make(map[string]struct{})
	for p := range snapshot.Local {
		paths[p] = struct{}{}
	}
	for p := range snapshot.Remote {
		paths[p] = struct{}{}
	}
	for p := range snapshot.Prev {
		paths[p] = struct{}{}
	}

	var allPaths []string
	for p := range paths {
		allPaths = append(allPaths, p)
	}
	sort.Strings(allPaths)

	var actions []Action
	var conflicts []Conflict

	for _, path := range allPaths {
		localEntry, localOK := snapshot.Local[path]
		remoteEntry, remoteOK := snapshot.Remote[path]
		prevEntry, prevOK := snapshot.Prev[path]

		localPtr := entryPtr(localOK, localEntry)
		remotePtr := remotePtr(remoteOK, remoteEntry)
		prevPtr := prevPtr(prevOK, prevEntry)

		if localOK && remoteOK && localEntry.IsDir != remoteEntry.IsDir {
			conflicts = append(conflicts, Conflict{
				Path:   path,
				Kind:   ConflictTypeMismatch,
				Local:  localPtr,
				Remote: remotePtr,
				Prev:   prevPtr,
			})
			continue
		}

		localChanged := localOK && isLocalModified(localEntry, prevPtr)
		remoteChanged := remoteOK && isRemoteModified(remoteEntry, prevPtr)

		localDeleted := !localOK && prevOK
		remoteDeleted := !remoteOK && prevOK

		switch {
		case localOK && remoteOK:
			if localChanged && remoteChanged {
				conflicts = append(conflicts, Conflict{
					Path:   path,
					Kind:   ConflictBothModified,
					Local:  localPtr,
					Remote: remotePtr,
					Prev:   prevPtr,
				})
				continue
			}
			if localChanged {
				actions = append(actions, Action{
					Type:   ActionUpdate,
					Path:   path,
					Local:  localPtr,
					Remote: remotePtr,
					Prev:   prevPtr,
				})
				continue
			}
			if remoteChanged {
				actions = append(actions, Action{
					Type:   ActionDownload,
					Path:   path,
					Local:  localPtr,
					Remote: remotePtr,
					Prev:   prevPtr,
				})
				continue
			}
		case localOK && !remoteOK:
			if localEntry.IsDir {
				if prevOK && prevEntry.IsDir {
					if includeDeletes && remoteDeleted {
						actions = append(actions, Action{
							Type:   ActionDeleteLocal,
							Path:   path,
							Local:  localPtr,
							Prev:   prevPtr,
						})
					}
					continue
				}
				actions = append(actions, Action{
					Type:  ActionMkdirRemote,
					Path:  path,
					Local: localPtr,
					Prev:  prevPtr,
				})
				continue
			}
			if prevOK && prevEntry.DriveFileID != "" {
				if localChanged && !remoteOK {
					conflicts = append(conflicts, Conflict{
						Path:   path,
						Kind:   ConflictRemoteDeletedLocalModified,
						Local:  localPtr,
						Remote: nil,
						Prev:   prevPtr,
					})
					continue
				}
				if includeDeletes {
					actions = append(actions, Action{
						Type:  ActionDeleteLocal,
						Path:  path,
						Local: localPtr,
						Prev:  prevPtr,
					})
				}
				continue
			}
			actions = append(actions, Action{
				Type:  ActionUpload,
				Path:  path,
				Local: localPtr,
				Prev:  prevPtr,
			})
		case !localOK && remoteOK:
			if remoteEntry.IsDir {
				if prevOK && prevEntry.IsDir {
					if includeDeletes && localDeleted {
						actions = append(actions, Action{
							Type:   ActionDeleteRemote,
							Path:   path,
							Remote: remotePtr,
							Prev:   prevPtr,
						})
					}
					continue
				}
				actions = append(actions, Action{
					Type:   ActionMkdirLocal,
					Path:   path,
					Remote: remotePtr,
					Prev:   prevPtr,
				})
				continue
			}
			if prevOK && prevEntry.LocalMTime > 0 {
				if remoteChanged && !localOK {
					conflicts = append(conflicts, Conflict{
						Path:   path,
						Kind:   ConflictLocalDeletedRemoteModified,
						Local:  nil,
						Remote: remotePtr,
						Prev:   prevPtr,
					})
					continue
				}
				if includeDeletes {
					actions = append(actions, Action{
						Type:   ActionDeleteRemote,
						Path:   path,
						Remote: remotePtr,
						Prev:   prevPtr,
					})
				}
				continue
			}
			actions = append(actions, Action{
				Type:   ActionDownload,
				Path:   path,
				Remote: remotePtr,
				Prev:   prevPtr,
			})
		default:
			if localDeleted && remoteDeleted {
				continue
			}
		}
	}

	actions = filterActionsForMode(actions, mode)

	return Result{
		Actions:   actions,
		Conflicts: conflicts,
	}
}

func filterActionsForMode(actions []Action, mode Mode) []Action {
	if mode == ModeBidirectional {
		return actions
	}
	filtered := make([]Action, 0, len(actions))
	for _, action := range actions {
		switch mode {
		case ModePush:
			if action.Type == ActionUpload || action.Type == ActionUpdate || action.Type == ActionDeleteRemote || action.Type == ActionMkdirRemote || action.Type == ActionMoveRemote {
				filtered = append(filtered, action)
			}
		case ModePull:
			if action.Type == ActionDownload || action.Type == ActionDeleteLocal || action.Type == ActionMkdirLocal || action.Type == ActionMoveLocal {
				filtered = append(filtered, action)
			}
		}
	}
	return filtered
}

func entryPtr(ok bool, entry scanner.LocalEntry) *scanner.LocalEntry {
	if !ok {
		return nil
	}
	e := entry
	return &e
}

func remotePtr(ok bool, entry scanner.RemoteEntry) *scanner.RemoteEntry {
	if !ok {
		return nil
	}
	e := entry
	return &e
}

func prevPtr(ok bool, entry index.SyncEntry) *index.SyncEntry {
	if !ok {
		return nil
	}
	e := entry
	return &e
}

func isLocalModified(local scanner.LocalEntry, prev *index.SyncEntry) bool {
	if prev == nil {
		return true
	}
	if local.IsDir != prev.IsDir {
		return true
	}
	if local.IsDir {
		return false
	}
	if local.Size != prev.LocalSize || local.ModTime != prev.LocalMTime {
		if prev.ContentHash != "" && local.Hash != "" {
			return prev.ContentHash != local.Hash
		}
		return true
	}
	return false
}

func isRemoteModified(remote scanner.RemoteEntry, prev *index.SyncEntry) bool {
	if prev == nil {
		return true
	}
	if remote.IsDir != prev.IsDir {
		return true
	}
	if remote.IsDir {
		return false
	}
	if remote.MD5Checksum != "" && prev.RemoteMD5 != "" {
		return remote.MD5Checksum != prev.RemoteMD5
	}
	if remote.ModifiedTime != "" && prev.RemoteMTime != "" {
		return remote.ModifiedTime != prev.RemoteMTime
	}
	if remote.Size != prev.RemoteSize {
		return true
	}
	return false
}
