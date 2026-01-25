package diff

import (
	"github.com/dl-alexandre/gdrv/internal/sync/index"
	"github.com/dl-alexandre/gdrv/internal/sync/scanner"
)

type ActionType string

const (
	ActionUpload      ActionType = "upload"
	ActionUpdate      ActionType = "update"
	ActionDownload    ActionType = "download"
	ActionDeleteLocal ActionType = "delete_local"
	ActionDeleteRemote ActionType = "delete_remote"
	ActionMoveLocal   ActionType = "move_local"
	ActionMoveRemote  ActionType = "move_remote"
	ActionMkdirLocal  ActionType = "mkdir_local"
	ActionMkdirRemote ActionType = "mkdir_remote"
)

type Action struct {
	Type     ActionType
	Path     string
	FromPath string
	ToPath   string
	Local    *scanner.LocalEntry
	Remote   *scanner.RemoteEntry
	Prev     *index.SyncEntry
	Name     string
}

type ConflictKind string

const (
	ConflictBothModified           ConflictKind = "both_modified"
	ConflictLocalDeletedRemoteModified ConflictKind = "local_deleted_remote_modified"
	ConflictRemoteDeletedLocalModified ConflictKind = "remote_deleted_local_modified"
	ConflictTypeMismatch            ConflictKind = "type_mismatch"
)

type Conflict struct {
	Path   string
	Kind   ConflictKind
	Local  *scanner.LocalEntry
	Remote *scanner.RemoteEntry
	Prev   *index.SyncEntry
}

type Mode string

const (
	ModePush         Mode = "push"
	ModePull         Mode = "pull"
	ModeBidirectional Mode = "bidirectional"
)
