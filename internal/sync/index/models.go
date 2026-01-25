package index

type SyncConfig struct {
	ID              string
	LocalRoot       string
	RemoteRootID    string
	ExcludePatterns []string
	ConflictPolicy  string
	Direction       string
	LastSyncTime    int64
	LastChangeToken string
}

type SyncEntry struct {
	ConfigID       string
	RelativePath   string
	DriveFileID    string
	DriveParentID  string
	IsDir          bool
	LocalMTime     int64
	LocalSize      int64
	ContentHash    string
	RemoteMTime    string
	RemoteSize     int64
	RemoteMD5      string
	RemoteMimeType string
	SyncState      string
	LastSync       int64
}
