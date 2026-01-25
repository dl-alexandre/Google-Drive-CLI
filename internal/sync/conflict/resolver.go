package conflict

import (
	"path"
	"strings"

	"github.com/dl-alexandre/gdrv/internal/sync/diff"
)

type Policy string

const (
	PolicyLocalWins  Policy = "local-wins"
	PolicyRemoteWins Policy = "remote-wins"
	PolicyRenameBoth Policy = "rename-both"
)

func Resolve(conflicts []diff.Conflict, policy Policy) ([]diff.Action, []diff.Conflict) {
	var actions []diff.Action
	var remaining []diff.Conflict

	for _, conflict := range conflicts {
		switch policy {
		case PolicyLocalWins:
			actions = append(actions, resolveLocalWins(conflict)...)
		case PolicyRemoteWins:
			actions = append(actions, resolveRemoteWins(conflict)...)
		case PolicyRenameBoth:
			resolved, ok := resolveRenameBoth(conflict)
			if ok {
				actions = append(actions, resolved...)
			} else {
				actions = append(actions, resolveLocalWins(conflict)...)
			}
		default:
			remaining = append(remaining, conflict)
		}
	}

	return actions, remaining
}

func resolveLocalWins(conflict diff.Conflict) []diff.Action {
	switch conflict.Kind {
	case diff.ConflictBothModified:
		return []diff.Action{{
			Type:   diff.ActionUpdate,
			Path:   conflict.Path,
			Local:  conflict.Local,
			Remote: conflict.Remote,
			Prev:   conflict.Prev,
		}}
	case diff.ConflictLocalDeletedRemoteModified:
		return []diff.Action{{
			Type:   diff.ActionDeleteRemote,
			Path:   conflict.Path,
			Remote: conflict.Remote,
			Prev:   conflict.Prev,
		}}
	case diff.ConflictRemoteDeletedLocalModified:
		return []diff.Action{{
			Type:  diff.ActionUpload,
			Path:  conflict.Path,
			Local: conflict.Local,
			Prev:  conflict.Prev,
		}}
	case diff.ConflictTypeMismatch:
		return []diff.Action{
			{
				Type:   diff.ActionDeleteRemote,
				Path:   conflict.Path,
				Remote: conflict.Remote,
				Prev:   conflict.Prev,
			},
			{
				Type:  diff.ActionUpload,
				Path:  conflict.Path,
				Local: conflict.Local,
				Prev:  conflict.Prev,
			},
		}
	}
	return nil
}

func resolveRemoteWins(conflict diff.Conflict) []diff.Action {
	switch conflict.Kind {
	case diff.ConflictBothModified:
		return []diff.Action{{
			Type:   diff.ActionDownload,
			Path:   conflict.Path,
			Local:  conflict.Local,
			Remote: conflict.Remote,
			Prev:   conflict.Prev,
		}}
	case diff.ConflictLocalDeletedRemoteModified:
		return []diff.Action{{
			Type:   diff.ActionDownload,
			Path:   conflict.Path,
			Remote: conflict.Remote,
			Prev:   conflict.Prev,
		}}
	case diff.ConflictRemoteDeletedLocalModified:
		return []diff.Action{{
			Type:  diff.ActionDeleteLocal,
			Path:  conflict.Path,
			Local: conflict.Local,
			Prev:  conflict.Prev,
		}}
	case diff.ConflictTypeMismatch:
		return []diff.Action{
			{
				Type:  diff.ActionDeleteLocal,
				Path:  conflict.Path,
				Local: conflict.Local,
				Prev:  conflict.Prev,
			},
			{
				Type:   diff.ActionDownload,
				Path:   conflict.Path,
				Remote: conflict.Remote,
				Prev:   conflict.Prev,
			},
		}
	}
	return nil
}

func resolveRenameBoth(conflict diff.Conflict) ([]diff.Action, bool) {
	if conflict.Kind != diff.ConflictBothModified || conflict.Local == nil || conflict.Remote == nil {
		return nil, false
	}
	localPath := addSuffix(conflict.Path, ".local")
	remotePath := addSuffix(conflict.Path, ".remote")
	actions := []diff.Action{
		{
			Type:     diff.ActionMoveLocal,
			FromPath: conflict.Path,
			ToPath:   localPath,
			Path:     localPath,
			Local:    conflict.Local,
			Prev:     conflict.Prev,
		},
		{
			Type:     diff.ActionMoveRemote,
			FromPath: conflict.Path,
			ToPath:   remotePath,
			Path:     remotePath,
			Remote:   conflict.Remote,
			Prev:     conflict.Prev,
		},
		{
			Type: diff.ActionUpload,
			Path: localPath,
			Prev: conflict.Prev,
		},
		{
			Type: diff.ActionDownload,
			Path: remotePath,
			Prev: conflict.Prev,
		},
	}
	return actions, true
}

func addSuffix(p string, suffix string) string {
	ext := path.Ext(p)
	base := strings.TrimSuffix(p, ext)
	return base + suffix + ext
}
