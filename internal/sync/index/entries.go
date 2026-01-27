package index

import (
	"context"
)

func (d *DB) ListEntries(ctx context.Context, configID string) (entries []SyncEntry, err error) {
	rows, err := d.db.QueryContext(ctx, `
		SELECT config_id, relative_path, drive_file_id, drive_parent_id, is_dir, local_mtime, local_size, content_hash,
		       remote_mtime, remote_size, remote_md5, remote_mime_type, sync_state, last_sync
		FROM sync_entries WHERE config_id = ?
	`, configID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	for rows.Next() {
		entry, err := scanEntry(rows)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return entries, nil
}

func (d *DB) GetEntryByPath(ctx context.Context, configID, relPath string) (*SyncEntry, error) {
	row := d.db.QueryRowContext(ctx, `
		SELECT config_id, relative_path, drive_file_id, drive_parent_id, is_dir, local_mtime, local_size, content_hash,
		       remote_mtime, remote_size, remote_md5, remote_mime_type, sync_state, last_sync
		FROM sync_entries WHERE config_id = ? AND relative_path = ?
	`, configID, relPath)
	entry, err := scanEntry(row)
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (d *DB) GetEntryByFileID(ctx context.Context, configID, fileID string) (*SyncEntry, error) {
	row := d.db.QueryRowContext(ctx, `
		SELECT config_id, relative_path, drive_file_id, drive_parent_id, is_dir, local_mtime, local_size, content_hash,
		       remote_mtime, remote_size, remote_md5, remote_mime_type, sync_state, last_sync
		FROM sync_entries WHERE config_id = ? AND drive_file_id = ? LIMIT 1
	`, configID, fileID)
	entry, err := scanEntry(row)
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (d *DB) ListEntriesByHash(ctx context.Context, configID, hash string) (entries []SyncEntry, err error) {
	rows, err := d.db.QueryContext(ctx, `
		SELECT config_id, relative_path, drive_file_id, drive_parent_id, is_dir, local_mtime, local_size, content_hash,
		       remote_mtime, remote_size, remote_md5, remote_mime_type, sync_state, last_sync
		FROM sync_entries WHERE config_id = ? AND content_hash = ?
	`, configID, hash)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	for rows.Next() {
		entry, err := scanEntry(rows)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return entries, nil
}

func (d *DB) ReplaceEntries(ctx context.Context, configID string, entries []SyncEntry) (err error) {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, `DELETE FROM sync_entries WHERE config_id = ?`, configID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO sync_entries (
			config_id, relative_path, drive_file_id, drive_parent_id, is_dir, local_mtime, local_size, content_hash,
			remote_mtime, remote_size, remote_md5, remote_mime_type, sync_state, last_sync
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	for _, entry := range entries {
		_, err := stmt.ExecContext(ctx, entry.ConfigID, entry.RelativePath, entry.DriveFileID, entry.DriveParentID, boolToInt(entry.IsDir),
			entry.LocalMTime, entry.LocalSize, entry.ContentHash, entry.RemoteMTime, entry.RemoteSize, entry.RemoteMD5, entry.RemoteMimeType, entry.SyncState, entry.LastSync)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (d *DB) DeleteEntries(ctx context.Context, configID string) error {
	_, err := d.db.ExecContext(ctx, `DELETE FROM sync_entries WHERE config_id = ?`, configID)
	return err
}

func scanEntry(scanner interface {
	Scan(dest ...interface{}) error
}) (SyncEntry, error) {
	var entry SyncEntry
	var isDir int
	err := scanner.Scan(&entry.ConfigID, &entry.RelativePath, &entry.DriveFileID, &entry.DriveParentID, &isDir, &entry.LocalMTime, &entry.LocalSize, &entry.ContentHash,
		&entry.RemoteMTime, &entry.RemoteSize, &entry.RemoteMD5, &entry.RemoteMimeType, &entry.SyncState, &entry.LastSync)
	if err != nil {
		return SyncEntry{}, err
	}
	entry.IsDir = isDir != 0
	return entry, nil
}

func boolToInt(v bool) int {
	if v {
		return 1
	}
	return 0
}
