package sync

import (
	"context"
	"errors"
	"path/filepath"
	"strings"
	"time"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/files"
	"github.com/dl-alexandre/gdrv/internal/folders"
	"github.com/dl-alexandre/gdrv/internal/sync/conflict"
	"github.com/dl-alexandre/gdrv/internal/sync/diff"
	"github.com/dl-alexandre/gdrv/internal/sync/exclude"
	"github.com/dl-alexandre/gdrv/internal/sync/executor"
	"github.com/dl-alexandre/gdrv/internal/sync/index"
	"github.com/dl-alexandre/gdrv/internal/sync/scanner"
	"github.com/dl-alexandre/gdrv/internal/types"
)

type Engine struct {
	client        *api.Client
	filesMgr      *files.Manager
	foldersMgr    *folders.Manager
	indexDB       *index.DB
	remoteScanner *scanner.RemoteScanner
}

type Options struct {
	Mode           diff.Mode
	ConflictPolicy conflict.Policy
	Delete         bool
	DryRun         bool
	Force          bool
	Yes            bool
	Concurrency    int
	UseChanges     bool
}

type Plan struct {
	Actions     []diff.Action
	Conflicts   []diff.Conflict
	Local       map[string]scanner.LocalEntry
	Remote      map[string]scanner.RemoteEntry
	Prev        map[string]index.SyncEntry
	ChangeToken string
}

type Result struct {
	Plan    Plan
	Summary executor.Summary
}

func NewEngine(client *api.Client, db *index.DB) *Engine {
	return &Engine{
		client:        client,
		filesMgr:      files.NewManager(client),
		foldersMgr:    folders.NewManager(client),
		indexDB:       db,
		remoteScanner: scanner.NewRemoteScanner(client),
	}
}

func (e *Engine) Close() error {
	if e == nil || e.indexDB == nil {
		return nil
	}
	return e.indexDB.Close()
}

func (e *Engine) Plan(ctx context.Context, cfg index.SyncConfig, opts Options, reqCtx *types.RequestContext) (Plan, error) {
	prevList, err := e.indexDB.ListEntries(ctx, cfg.ID)
	if err != nil {
		return Plan{}, err
	}
	prevMap := make(map[string]index.SyncEntry)
	for _, entry := range prevList {
		prevMap[entry.RelativePath] = entry
	}

	localRoot := cfg.LocalRoot
	if !filepath.IsAbs(localRoot) {
		localRoot, err = filepath.Abs(localRoot)
		if err != nil {
			return Plan{}, err
		}
	}

	matcher := exclude.New(cfg.ExcludePatterns)
	localEntries, err := scanner.ScanLocal(ctx, localRoot, matcher, prevMap)
	if err != nil {
		return Plan{}, err
	}

	var remoteEntries map[string]scanner.RemoteEntry
	changeToken := cfg.LastChangeToken

	if opts.UseChanges && changeToken != "" {
		newToken := ""
		fullScan := false
		remoteEntries, newToken, fullScan, err = e.remoteScanner.ListTreeWithChanges(ctx, reqCtx, cfg.RemoteRootID, changeToken, prevList)
		if err != nil {
			return Plan{}, err
		}
		if fullScan {
			changeToken, err = e.remoteScanner.GetStartPageToken(ctx, reqCtx)
			if err != nil {
				return Plan{}, err
			}
		} else if newToken != "" {
			changeToken = newToken
		}
	} else {
		remoteEntries, err = e.remoteScanner.ListTree(ctx, reqCtx, cfg.RemoteRootID)
		if err != nil {
			return Plan{}, err
		}
		if opts.UseChanges {
			changeToken, err = e.remoteScanner.GetStartPageToken(ctx, reqCtx)
			if err != nil {
				return Plan{}, err
			}
		}
	}

	mode := opts.Mode
	if mode == "" {
		mode = parseMode(cfg.Direction)
	}

	result := diff.Compute(diff.Snapshot{
		Local:  localEntries,
		Remote: remoteEntries,
		Prev:   prevMap,
	}, mode, opts.Delete)
	result = diff.ApplyRenames(result, diff.Snapshot{
		Local:  localEntries,
		Remote: remoteEntries,
		Prev:   prevMap,
	})

	policy := opts.ConflictPolicy
	if policy == "" {
		policy = parsePolicy(cfg.ConflictPolicy)
	}

	resolved, remaining := conflict.Resolve(result.Conflicts, policy)
	result.Actions = append(result.Actions, resolved...)
	result.Conflicts = remaining

	return Plan{
		Actions:     result.Actions,
		Conflicts:   result.Conflicts,
		Local:       localEntries,
		Remote:      remoteEntries,
		Prev:        prevMap,
		ChangeToken: changeToken,
	}, nil
}

func (e *Engine) Apply(ctx context.Context, cfg index.SyncConfig, plan Plan, opts Options, reqCtx *types.RequestContext) (Result, error) {
	exec := executor.New(e.filesMgr, e.foldersMgr)
	state := executor.State{
		LocalRoot:    cfg.LocalRoot,
		RemoteRootID: cfg.RemoteRootID,
		LocalEntries: plan.Local,
		RemoteEntries: plan.Remote,
	}
	state, summary, err := exec.Apply(ctx, reqCtx, plan.Actions, state, executor.Options{
		Concurrency: opts.Concurrency,
		DryRun:      opts.DryRun,
		Force:       opts.Force,
		Yes:         opts.Yes,
	})
	if err != nil {
		return Result{}, err
	}

	if !opts.DryRun {
		newEntries := buildIndexEntries(cfg.ID, state.LocalEntries, state.RemoteEntries)
		if err := e.indexDB.ReplaceEntries(ctx, cfg.ID, newEntries); err != nil {
			return Result{}, err
		}
		cfg.LastSyncTime = time.Now().Unix()
		if plan.ChangeToken != "" {
			cfg.LastChangeToken = plan.ChangeToken
		}
		if err := e.indexDB.UpsertConfig(ctx, cfg); err != nil {
			return Result{}, err
		}
	}

	return Result{
		Plan:    plan,
		Summary: summary,
	}, nil
}

func buildIndexEntries(configID string, local map[string]scanner.LocalEntry, remote map[string]scanner.RemoteEntry) []index.SyncEntry {
	paths := make(map[string]struct{})
	for p := range local {
		paths[p] = struct{}{}
	}
	for p := range remote {
		paths[p] = struct{}{}
	}
	var entries []index.SyncEntry
	for p := range paths {
		localEntry, localOK := local[p]
		remoteEntry, remoteOK := remote[p]
		entry := index.SyncEntry{
			ConfigID:     configID,
			RelativePath: p,
			IsDir:        localEntry.IsDir || remoteEntry.IsDir,
			LocalMTime:   localEntry.ModTime,
			LocalSize:    localEntry.Size,
			ContentHash:  localEntry.Hash,
			RemoteMTime:  remoteEntry.ModifiedTime,
			RemoteSize:   remoteEntry.Size,
			RemoteMD5:    remoteEntry.MD5Checksum,
			RemoteMimeType: remoteEntry.MimeType,
			DriveFileID:  remoteEntry.ID,
			DriveParentID: remoteEntry.ParentID,
		}
		if !localOK {
			entry.LocalMTime = 0
			entry.LocalSize = 0
			entry.ContentHash = ""
		}
		if !remoteOK {
			entry.RemoteMTime = ""
			entry.RemoteSize = 0
			entry.RemoteMD5 = ""
			entry.RemoteMimeType = ""
			entry.DriveFileID = ""
			entry.DriveParentID = ""
		}
		entries = append(entries, entry)
	}
	return entries
}

func parseMode(value string) diff.Mode {
	switch strings.ToLower(value) {
	case "push":
		return diff.ModePush
	case "pull":
		return diff.ModePull
	default:
		return diff.ModeBidirectional
	}
}

func parsePolicy(value string) conflict.Policy {
	switch strings.ToLower(value) {
	case "local-wins":
		return conflict.PolicyLocalWins
	case "remote-wins":
		return conflict.PolicyRemoteWins
	default:
		return conflict.PolicyRenameBoth
	}
}

func EnsureConfig(cfg *index.SyncConfig) error {
	if cfg == nil {
		return errors.New("config is nil")
	}
	if cfg.ID == "" || cfg.LocalRoot == "" || cfg.RemoteRootID == "" {
		return errors.New("config missing required fields")
	}
	if cfg.ConflictPolicy == "" {
		cfg.ConflictPolicy = string(conflict.PolicyRenameBoth)
	}
	if cfg.Direction == "" {
		cfg.Direction = string(diff.ModeBidirectional)
	}
	return nil
}
