package scanner

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"path"

	"github.com/dl-alexandre/gdrv/internal/sync/exclude"
	"github.com/dl-alexandre/gdrv/internal/sync/index"
)

func ScanLocal(ctx context.Context, root string, matcher *exclude.Matcher, prev map[string]index.SyncEntry) (map[string]LocalEntry, error) {
	entries := make(map[string]LocalEntry)

	err := filepath.WalkDir(root, func(current string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}

		if d.Type()&os.ModeSymlink != 0 {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		rel, err := filepath.Rel(root, current)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}
		rel = path.Clean(filepath.ToSlash(rel))

		if matcher != nil && matcher.IsExcluded(rel, d.IsDir()) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		if info.Mode().IsRegular() {
			hash := ""
			prevEntry, ok := prev[rel]
			if ok && prevEntry.LocalSize == info.Size() && prevEntry.LocalMTime == info.ModTime().Unix() && prevEntry.ContentHash != "" {
				hash = prevEntry.ContentHash
			} else {
				hash, err = hashFile(current)
				if err != nil {
					return err
				}
			}

			entries[rel] = LocalEntry{
				RelativePath: rel,
				AbsPath:      current,
				IsDir:        false,
				Size:         info.Size(),
				ModTime:      info.ModTime().Unix(),
				Hash:         hash,
			}
			return nil
		}

		if info.IsDir() {
			entries[rel] = LocalEntry{
				RelativePath: rel,
				AbsPath:      current,
				IsDir:        true,
				ModTime:      info.ModTime().Unix(),
			}
			return nil
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func hashFile(path string) (hash string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
