package scanner

type LocalEntry struct {
	RelativePath string
	AbsPath      string
	IsDir        bool
	Size         int64
	ModTime      int64
	Hash         string
}

type RemoteEntry struct {
	RelativePath string
	ID           string
	ParentID     string
	IsDir        bool
	Size         int64
	ModifiedTime string
	MD5Checksum  string
	MimeType     string
}
