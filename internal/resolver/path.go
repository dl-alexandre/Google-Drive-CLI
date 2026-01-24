package resolver

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/dl-alexandre/gdrive/internal/api"
	"github.com/dl-alexandre/gdrive/internal/types"
	"github.com/dl-alexandre/gdrive/internal/utils"
	"google.golang.org/api/drive/v3"
)

// PathResolver resolves human-readable paths to Drive file IDs
type PathResolver struct {
	client   *api.Client
	shaper   *api.RequestShaper
	cache    *pathCache
	cacheTTL time.Duration
}

type pathCache struct {
	mu      sync.RWMutex
	entries map[string]cacheEntry
}

type cacheEntry struct {
	fileID    string
	timestamp time.Time
}

// NewPathResolver creates a new path resolver
func NewPathResolver(client *api.Client, cacheTTL time.Duration) *PathResolver {
	return &PathResolver{
		client:   client,
		shaper:   api.NewRequestShaper(client),
		cacheTTL: cacheTTL,
		cache: &pathCache{
			entries: make(map[string]cacheEntry),
		},
	}
}

// SearchDomain represents the scope of path resolution
type SearchDomain string

const (
	SearchDomainMyDrive      SearchDomain = "my-drive"
	SearchDomainSharedDrive  SearchDomain = "shared-drive"
	SearchDomainSharedWithMe SearchDomain = "shared-with-me"
	SearchDomainAllDrives    SearchDomain = "all-drives"
	SearchDomainDomain       SearchDomain = "domain" // Workspace domain-wide
)

// ResolveOptions configures path resolution
type ResolveOptions struct {
	DriveID             string
	SearchDomain        SearchDomain
	IncludeSharedWithMe bool
	UseCache            bool
	StrictMode          bool
	MaxAncestorDepth    int // For shared-with-me ancestor walk validation (default: 10)
}

// ResolveResult contains path resolution results
type ResolveResult struct {
	FileID       string
	File         *types.DriveFile
	Ambiguous    bool
	Cached       bool
	Matches      []*types.DriveFile
	SearchDomain SearchDomain // Which domain the result came from
}

// Resolve resolves a path to a file ID
func (r *PathResolver) Resolve(ctx context.Context, reqCtx *types.RequestContext, path string, opts ResolveOptions) (*ResolveResult, error) {
	// Normalize path
	path = normalizePath(path)

	// Set default max ancestor depth if not specified
	if opts.MaxAncestorDepth == 0 {
		opts.MaxAncestorDepth = 10
	}

	// Determine search domain if not explicitly set
	if opts.SearchDomain == "" {
		opts.SearchDomain = r.determineSearchDomain(opts)
	}

	// Check cache first
	if opts.UseCache {
		cacheKey := r.makeCacheKey(path, opts)
		if cached, ok := r.checkCacheByKey(cacheKey); ok {
			return &ResolveResult{
				FileID:       cached,
				Cached:       true,
				SearchDomain: opts.SearchDomain,
			}, nil
		}
	}

	// Split path into segments
	segments := strings.Split(path, "/")
	if len(segments) == 0 {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidPath, "Empty path").Build())
	}

	// Handle shared-with-me paths specially
	if opts.SearchDomain == SearchDomainSharedWithMe {
		return r.resolveSharedWithMePath(ctx, reqCtx, segments, path, opts)
	}

	// Standard parent-based resolution for My Drive and Shared Drives
	return r.resolveParentBasedPath(ctx, reqCtx, segments, path, opts)
}

// determineSearchDomain determines the appropriate search domain based on options
func (r *PathResolver) determineSearchDomain(opts ResolveOptions) SearchDomain {
	if opts.DriveID != "" {
		return SearchDomainSharedDrive
	}
	if opts.IncludeSharedWithMe {
		return SearchDomainAllDrives // Search across all accessible content
	}
	return SearchDomainMyDrive
}

// resolveParentBasedPath resolves a path using standard parent-child traversal
func (r *PathResolver) resolveParentBasedPath(ctx context.Context, reqCtx *types.RequestContext, segments []string, path string, opts ResolveOptions) (*ResolveResult, error) {
	// Start from root
	currentID := "root"
	if opts.DriveID != "" {
		currentID = opts.DriveID
		reqCtx.DriveID = opts.DriveID
	}

	// Walk path segments
	for i, segment := range segments {
		if segment == "" {
			continue
		}

		matches, err := r.findByName(ctx, reqCtx, currentID, segment, opts)
		if err != nil {
			return nil, err
		}

		if len(matches) == 0 {
			return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeFileNotFound,
				fmt.Sprintf("Path segment not found: %s (at %s)", segment, strings.Join(segments[:i+1], "/"))).
				WithContext("path", path).
				WithContext("segment", segment).
				WithContext("searchDomain", string(opts.SearchDomain)).
				Build())
		}

		// Apply disambiguation
		if len(matches) > 1 {
			if opts.StrictMode {
				return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeAmbiguousPath,
					fmt.Sprintf("Ambiguous path: multiple matches for '%s'", segment)).
					WithContext("matches", len(matches)).
					WithContext("matchCount", len(matches)).
					Build())
			}
			// Use deterministic ordering with domain preference
			matches = r.sortMatchesWithDomainPreference(matches, opts.SearchDomain)
		}

		currentID = matches[0].ID

		// If this is the last segment, return full result
		if i == len(segments)-1 {
			result := &ResolveResult{
				FileID:       currentID,
				File:         matches[0],
				Ambiguous:    len(matches) > 1,
				Matches:      matches,
				SearchDomain: opts.SearchDomain,
			}

			// Update cache
			if opts.UseCache {
				cacheKey := r.makeCacheKey(path, opts)
				r.updateCacheByKey(cacheKey, currentID)
			}

			return result, nil
		}
	}

	return &ResolveResult{
		FileID:       currentID,
		SearchDomain: opts.SearchDomain,
	}, nil
}

// resolveSharedWithMePath handles resolution for shared-with-me items
func (r *PathResolver) resolveSharedWithMePath(ctx context.Context, reqCtx *types.RequestContext, segments []string, path string, opts ResolveOptions) (*ResolveResult, error) {
	// For single-segment paths, use simple name lookup
	if len(segments) == 1 && segments[0] != "" {
		return r.resolveSharedWithMeSingleSegment(ctx, reqCtx, segments[0], path, opts)
	}

	// For multi-level paths, enforce strict mode in production or warn
	if opts.StrictMode {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidPath,
			"Multi-level shared-with-me paths require file IDs in strict mode").
			WithContext("path", path).
			WithContext("segments", len(segments)).
			WithContext("reason", "shared-with-me items may lack traversable folder hierarchy").
			WithContext("suggestedAction", "use file ID directly or single-segment query").
			Build())
	}

	// In non-strict mode, attempt ancestor-walk verification
	return r.resolveSharedWithMeMultiLevel(ctx, reqCtx, segments, path, opts)
}

// resolveSharedWithMeSingleSegment resolves a single-segment shared-with-me path
func (r *PathResolver) resolveSharedWithMeSingleSegment(ctx context.Context, reqCtx *types.RequestContext, name string, path string, opts ResolveOptions) (*ResolveResult, error) {
	// Search for items shared with me by name
	matches, err := r.findSharedWithMeByName(ctx, reqCtx, name, opts)
	if err != nil {
		return nil, err
	}

	if len(matches) == 0 {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeFileNotFound,
			fmt.Sprintf("No shared-with-me items found with name: %s", name)).
			WithContext("path", path).
			WithContext("searchDomain", "shared-with-me").
			Build())
	}

	// Apply disambiguation
	if len(matches) > 1 {
		if opts.StrictMode {
			return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeAmbiguousPath,
				fmt.Sprintf("Ambiguous path: multiple shared-with-me items match '%s'", name)).
				WithContext("matches", len(matches)).
				WithContext("matchCount", len(matches)).
				Build())
		}
		matches = r.sortMatchesWithDomainPreference(matches, SearchDomainSharedWithMe)
	}

	result := &ResolveResult{
		FileID:       matches[0].ID,
		File:         matches[0],
		Ambiguous:    len(matches) > 1,
		Matches:      matches,
		SearchDomain: SearchDomainSharedWithMe,
	}

	// Update cache
	if opts.UseCache {
		cacheKey := r.makeCacheKey(path, opts)
		r.updateCacheByKey(cacheKey, matches[0].ID)
	}

	return result, nil
}

// resolveSharedWithMeMultiLevel attempts to resolve multi-level shared-with-me paths with ancestor verification
func (r *PathResolver) resolveSharedWithMeMultiLevel(ctx context.Context, reqCtx *types.RequestContext, segments []string, path string, opts ResolveOptions) (*ResolveResult, error) {
	// Start by finding candidates for the first segment
	firstSegment := segments[0]
	if firstSegment == "" {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidPath, "Empty first segment in path").Build())
	}

	candidates, err := r.findSharedWithMeByName(ctx, reqCtx, firstSegment, opts)
	if err != nil {
		return nil, err
	}

	if len(candidates) == 0 {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeFileNotFound,
			fmt.Sprintf("No shared-with-me items found with name: %s", firstSegment)).
			WithContext("path", path).
			Build())
	}

	// For each candidate, try to traverse the remaining path segments
	var validMatches []*types.DriveFile
	for _, candidate := range candidates {
		// Try to traverse from this candidate
		currentID := candidate.ID
		valid := true

		for i := 1; i < len(segments); i++ {
			segment := segments[i]
			if segment == "" {
				continue
			}

			// Find children of current node
			childMatches, err := r.findByNameWithParent(ctx, reqCtx, currentID, segment, opts)
			if err != nil {
				// Permission error or network error
				if isPermissionError(err) {
					return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodePermissionDenied,
						"Insufficient permissions to verify shared-with-me path ancestors").
						WithContext("path", path).
						WithContext("failedAt", strings.Join(segments[:i+1], "/")).
						WithContext("suggestedAction", "use file ID or verify access permissions").
						Build())
				}
				return nil, err
			}

			if len(childMatches) == 0 {
				valid = false
				break
			}

			// Take the first match (or apply disambiguation if needed)
			if len(childMatches) > 1 {
				childMatches = r.sortMatchesWithDomainPreference(childMatches, SearchDomainSharedWithMe)
			}
			currentID = childMatches[0].ID

			// Check depth bounds
			if i >= opts.MaxAncestorDepth {
				return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeInvalidPath,
					fmt.Sprintf("Path depth exceeds maximum ancestor walk depth (%d)", opts.MaxAncestorDepth)).
					WithContext("path", path).
					WithContext("maxDepth", opts.MaxAncestorDepth).
					Build())
			}

			// If this is the last segment, we found a valid match
			if i == len(segments)-1 {
				// Fetch full file details for the final match
				finalFile, err := r.getFileByID(ctx, reqCtx, currentID)
				if err != nil {
					return nil, err
				}
				validMatches = append(validMatches, finalFile)
			}
		}

		if !valid {
			continue
		}
	}

	if len(validMatches) == 0 {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeFileNotFound,
			"No valid shared-with-me paths found matching all segments").
			WithContext("path", path).
			WithContext("candidatesChecked", len(candidates)).
			Build())
	}

	// Apply disambiguation
	if len(validMatches) > 1 {
		if opts.StrictMode {
			return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeAmbiguousPath,
				"Multiple shared-with-me paths match the given segments").
				WithContext("matches", len(validMatches)).
				Build())
		}
		validMatches = r.sortMatchesWithDomainPreference(validMatches, SearchDomainSharedWithMe)
	}

	result := &ResolveResult{
		FileID:       validMatches[0].ID,
		File:         validMatches[0],
		Ambiguous:    len(validMatches) > 1,
		Matches:      validMatches,
		SearchDomain: SearchDomainSharedWithMe,
	}

	// Update cache
	if opts.UseCache {
		cacheKey := r.makeCacheKey(path, opts)
		r.updateCacheByKey(cacheKey, validMatches[0].ID)
	}

	return result, nil
}

func (r *PathResolver) findByName(ctx context.Context, reqCtx *types.RequestContext, parentID, name string, opts ResolveOptions) ([]*types.DriveFile, error) {
	return r.findByNameWithParent(ctx, reqCtx, parentID, name, opts)
}

// findByNameWithParent searches for files with a specific name under a parent
func (r *PathResolver) findByNameWithParent(ctx context.Context, reqCtx *types.RequestContext, parentID, name string, opts ResolveOptions) ([]*types.DriveFile, error) {
	// Escape single quotes in name
	escapedName := escapeQueryString(name)

	query := fmt.Sprintf("'%s' in parents and name = '%s' and trashed = false", parentID, escapedName)

	call := r.client.Service().Files.List().Q(query)
	call = r.shaper.ShapeFilesList(call, reqCtx)
	call = call.Fields("files(id,name,mimeType,parents,resourceKey,shortcutDetails,owners,driveId)")

	result, err := api.ExecuteWithRetry(ctx, r.client, reqCtx, func() (*drive.FileList, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	matches := make([]*types.DriveFile, len(result.Files))
	for i, f := range result.Files {
		matches[i] = &types.DriveFile{
			ID:          f.Id,
			Name:        f.Name,
			MimeType:    f.MimeType,
			Parents:     f.Parents,
			ResourceKey: f.ResourceKey,
		}

		// Update resource key cache
		if f.ResourceKey != "" {
			r.client.ResourceKeys().UpdateFromAPIResponse(f.Id, f.ResourceKey)
		}
	}

	return matches, nil
}

// findSharedWithMeByName searches for files shared with me by name
func (r *PathResolver) findSharedWithMeByName(ctx context.Context, reqCtx *types.RequestContext, name string, opts ResolveOptions) ([]*types.DriveFile, error) {
	// Escape single quotes in name
	escapedName := escapeQueryString(name)

	// Query for shared-with-me items with the given name
	query := fmt.Sprintf("sharedWithMe = true and name = '%s' and trashed = false", escapedName)

	// Create a new request context for shared-with-me search
	sharedReqCtx := &types.RequestContext{
		Profile:           reqCtx.Profile,
		DriveID:           "", // No specific drive for shared-with-me
		InvolvedFileIDs:   []string{},
		InvolvedParentIDs: []string{},
		RequestType:       reqCtx.RequestType,
		TraceID:           reqCtx.TraceID,
	}

	call := r.client.Service().Files.List().Q(query)
	call = r.shaper.ShapeFilesList(call, sharedReqCtx)
	call = call.Fields("files(id,name,mimeType,parents,resourceKey,shortcutDetails,owners,driveId,shared)")

	result, err := api.ExecuteWithRetry(ctx, r.client, sharedReqCtx, func() (*drive.FileList, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	matches := make([]*types.DriveFile, len(result.Files))
	for i, f := range result.Files {
		matches[i] = &types.DriveFile{
			ID:          f.Id,
			Name:        f.Name,
			MimeType:    f.MimeType,
			Parents:     f.Parents,
			ResourceKey: f.ResourceKey,
		}

		// Update resource key cache
		if f.ResourceKey != "" {
			r.client.ResourceKeys().UpdateFromAPIResponse(f.Id, f.ResourceKey)
		}
	}

	return matches, nil
}

// getFileByID retrieves a file by its ID
func (r *PathResolver) getFileByID(ctx context.Context, reqCtx *types.RequestContext, fileID string) (*types.DriveFile, error) {
	call := r.client.Service().Files.Get(fileID)
	call = r.shaper.ShapeFilesGet(call, reqCtx)
	call = call.Fields("id,name,mimeType,parents,resourceKey,shortcutDetails,owners,driveId")

	result, err := api.ExecuteWithRetry(ctx, r.client, reqCtx, func() (*drive.File, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	file := &types.DriveFile{
		ID:          result.Id,
		Name:        result.Name,
		MimeType:    result.MimeType,
		Parents:     result.Parents,
		ResourceKey: result.ResourceKey,
	}

	// Update resource key cache
	if result.ResourceKey != "" {
		r.client.ResourceKeys().UpdateFromAPIResponse(result.Id, result.ResourceKey)
	}

	return file, nil
}

// isPermissionError checks if an error is a permission-related error
func isPermissionError(err error) bool {
	if appErr, ok := err.(*utils.AppError); ok {
		return appErr.CLIError.Code == utils.ErrCodePermissionDenied
	}
	return false
}

func (r *PathResolver) checkCache(path, driveID string) (string, bool) {
	r.cache.mu.RLock()
	defer r.cache.mu.RUnlock()

	key := cacheKey(path, driveID)
	entry, ok := r.cache.entries[key]
	if !ok {
		return "", false
	}

	if time.Since(entry.timestamp) > r.cacheTTL {
		return "", false
	}

	return entry.fileID, true
}

func (r *PathResolver) checkCacheByKey(key string) (string, bool) {
	r.cache.mu.RLock()
	defer r.cache.mu.RUnlock()

	entry, ok := r.cache.entries[key]
	if !ok {
		return "", false
	}

	if time.Since(entry.timestamp) > r.cacheTTL {
		return "", false
	}

	return entry.fileID, true
}

func (r *PathResolver) updateCache(path, driveID, fileID string) {
	r.cache.mu.Lock()
	defer r.cache.mu.Unlock()

	key := cacheKey(path, driveID)
	r.cache.entries[key] = cacheEntry{
		fileID:    fileID,
		timestamp: time.Now(),
	}
}

func (r *PathResolver) updateCacheByKey(key, fileID string) {
	r.cache.mu.Lock()
	defer r.cache.mu.Unlock()

	r.cache.entries[key] = cacheEntry{
		fileID:    fileID,
		timestamp: time.Now(),
	}
}

// InvalidateCache removes a path from the cache
func (r *PathResolver) InvalidateCache(path, driveID string) {
	r.cache.mu.Lock()
	defer r.cache.mu.Unlock()

	key := cacheKey(path, driveID)
	delete(r.cache.entries, key)
}

// InvalidateCacheByKey removes a cache entry by key
func (r *PathResolver) InvalidateCacheByKey(key string) {
	r.cache.mu.Lock()
	defer r.cache.mu.Unlock()

	delete(r.cache.entries, key)
}

// makeCacheKey creates a cache key that includes search domain
func (r *PathResolver) makeCacheKey(path string, opts ResolveOptions) string {
	// Include search domain in cache key to differentiate between different resolution contexts
	return fmt.Sprintf("%s:%s:%s", opts.DriveID, string(opts.SearchDomain), path)
}

// ClearCache removes all cached entries
func (r *PathResolver) ClearCache() {
	r.cache.mu.Lock()
	defer r.cache.mu.Unlock()

	r.cache.entries = make(map[string]cacheEntry)
}

func cacheKey(path, driveID string) string {
	return driveID + ":" + path
}

func normalizePath(path string) string {
	path = strings.TrimPrefix(path, "/")
	path = strings.TrimSuffix(path, "/")
	return path
}

func escapeQueryString(s string) string {
	// Escape backslashes first, then single quotes
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "'", "\\'")
	return s
}

// sortMatchesWithDomainPreference sorts matches with domain preference
// Priority: My Drive > Shared Drives > shared-with-me
func (r *PathResolver) sortMatchesWithDomainPreference(matches []*types.DriveFile, currentDomain SearchDomain) []*types.DriveFile {
	// Use bubble sort for simplicity (adequate for small match sets)
	for i := 0; i < len(matches)-1; i++ {
		for j := i + 1; j < len(matches); j++ {
			if r.shouldSwapWithDomain(matches[i], matches[j], currentDomain) {
				matches[i], matches[j] = matches[j], matches[i]
			}
		}
	}
	return matches
}

// shouldSwapWithDomain determines if two files should be swapped based on disambiguation rules
func (r *PathResolver) shouldSwapWithDomain(a, b *types.DriveFile, currentDomain SearchDomain) bool {
	// Prefer non-shortcut globally
	aIsShortcut := a.MimeType == utils.MimeTypeShortcut
	bIsShortcut := b.MimeType == utils.MimeTypeShortcut
	if aIsShortcut != bIsShortcut {
		return aIsShortcut // Swap if 'a' is shortcut and 'b' is not
	}

	// Determine domain priority for each file
	aDomain := r.inferFileDomain(a)
	bDomain := r.inferFileDomain(b)

	// Apply domain preference: My Drive (1) > Shared Drive (2) > shared-with-me (3)
	aPriority := r.getDomainPriority(aDomain)
	bPriority := r.getDomainPriority(bDomain)

	if aPriority != bPriority {
		return aPriority > bPriority // Swap if 'a' has lower priority (higher number)
	}

	// Then by name (lexicographic)
	if a.Name != b.Name {
		return a.Name > b.Name // Swap if 'a' comes after 'b'
	}

	// Finally by ID for stability
	return a.ID > b.ID
}

// inferFileDomain infers which domain a file belongs to based on its properties
func (r *PathResolver) inferFileDomain(file *types.DriveFile) SearchDomain {
	// If file has no parents or empty parents, it's likely shared-with-me
	if len(file.Parents) == 0 {
		return SearchDomainSharedWithMe
	}

	// Files in My Drive have parent "root" or descend from root
	// This is a heuristic - in practice, we'd need to check the parent chain
	// For now, assume files with parents are in My Drive or Shared Drive
	// We can't easily distinguish without additional API calls

	// Default to My Drive for files with parents
	return SearchDomainMyDrive
}

// getDomainPriority returns the priority for a search domain (lower is better)
func (r *PathResolver) getDomainPriority(domain SearchDomain) int {
	switch domain {
	case SearchDomainMyDrive:
		return 1
	case SearchDomainSharedDrive:
		return 2
	case SearchDomainSharedWithMe:
		return 3
	default:
		return 4
	}
}

