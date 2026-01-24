package api

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

// ResourceKeyManager manages resource keys for link-shared files
type ResourceKeyManager struct {
	mu    sync.RWMutex
	cache map[string]resourceKeyEntry
	path  string
}

type resourceKeyEntry struct {
	ResourceKey string `json:"resourceKey"`
	Timestamp   int64  `json:"timestamp"`
	Source      string `json:"source"` // url, api, shortcut
}

// NewResourceKeyManager creates a new resource key manager
func NewResourceKeyManager() *ResourceKeyManager {
	mgr := &ResourceKeyManager{
		cache: make(map[string]resourceKeyEntry),
	}
	return mgr
}

// SetCachePath sets the path for persisting resource keys
func (m *ResourceKeyManager) SetCachePath(path string) error {
	m.path = path
	return m.load()
}

// AddKey adds a resource key to the cache
func (m *ResourceKeyManager) AddKey(fileID, resourceKey, source string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache[fileID] = resourceKeyEntry{
		ResourceKey: resourceKey,
		Timestamp:   timeNow().Unix(),
		Source:      source,
	}
	if err := m.save(); err != nil {
		return
	}
}

// GetKey retrieves a resource key from the cache
func (m *ResourceKeyManager) GetKey(fileID string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	entry, ok := m.cache[fileID]
	if !ok {
		return "", false
	}
	return entry.ResourceKey, true
}

// BuildHeader builds the X-Goog-Drive-Resource-Keys header
func (m *ResourceKeyManager) BuildHeader(fileIDs []string) string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var pairs []string
	for _, id := range fileIDs {
		if entry, ok := m.cache[id]; ok {
			pairs = append(pairs, id+"/"+entry.ResourceKey)
		}
	}

	if len(pairs) == 0 {
		return ""
	}
	return strings.Join(pairs, ",")
}

// ParseFromURL extracts fileID and resourceKey from a Drive sharing URL
func (m *ResourceKeyManager) ParseFromURL(url string) (string, string, bool) {
	// Only accept HTTPS URLs from drive.google.com
	if !strings.HasPrefix(url, "https://drive.google.com/") {
		return "", "", false
	}

	// Match patterns like:
	// https://drive.google.com/file/d/FILE_ID/view?resourcekey=KEY
	// https://drive.google.com/open?id=FILE_ID&resourcekey=KEY
	// https://drive.google.com/drive/folders/FILE_ID

	fileIDPattern := regexp.MustCompile(`/d/([a-zA-Z0-9_-]+)`)
	folderIDPattern := regexp.MustCompile(`/folders/([a-zA-Z0-9_-]+)`)
	resourceKeyPattern := regexp.MustCompile(`resourcekey=([a-zA-Z0-9_-]+)`)

	fileIDMatch := fileIDPattern.FindStringSubmatch(url)
	if fileIDMatch == nil {
		// Try folders pattern
		fileIDMatch = folderIDPattern.FindStringSubmatch(url)
		if fileIDMatch == nil {
			// Try open?id= pattern
			idPattern := regexp.MustCompile(`[?&]id=([a-zA-Z0-9_-]+)`)
			fileIDMatch = idPattern.FindStringSubmatch(url)
			if fileIDMatch == nil {
				return "", "", false
			}
		}
	}

	resourceKeyMatch := resourceKeyPattern.FindStringSubmatch(url)
	if resourceKeyMatch == nil {
		return fileIDMatch[1], "", true // Valid URL but no resource key
	}

	return fileIDMatch[1], resourceKeyMatch[1], true
}

// Invalidate removes a resource key from the cache
func (m *ResourceKeyManager) Invalidate(fileID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.cache, fileID)
	if err := m.save(); err != nil {
		return
	}
}

// Clear removes all cached resource keys
func (m *ResourceKeyManager) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache = make(map[string]resourceKeyEntry)
	if err := m.save(); err != nil {
		return
	}
}

// UpdateFromAPIResponse updates cache from file metadata
func (m *ResourceKeyManager) UpdateFromAPIResponse(fileID, resourceKey string) {
	if resourceKey != "" {
		m.AddKey(fileID, resourceKey, "api")
	}
}

func (m *ResourceKeyManager) load() error {
	if m.path == "" {
		return nil
	}

	data, err := os.ReadFile(m.path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	return json.Unmarshal(data, &m.cache)
}

func (m *ResourceKeyManager) save() error {
	if m.path == "" {
		return nil
	}

	data, err := json.MarshalIndent(m.cache, "", "  ")
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(m.path), 0700); err != nil {
		return err
	}
	return os.WriteFile(m.path, data, 0600)
}

// timeProvider is an interface for getting time
type timeProvider interface {
	Unix() int64
}

// realTimeProvider implements timeProvider using time.Now()
type realTimeProvider struct{}

func (t *realTimeProvider) Unix() int64 {
	return time.Now().Unix()
}

// timeNow is a variable for testing
var timeNow = func() timeProvider {
	return &realTimeProvider{}
}
