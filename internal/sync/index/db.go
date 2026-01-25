package index

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

type DB struct {
	db *sql.DB
}

func Open(path string) (*DB, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	instance := &DB{db: db}
	if err := instance.Migrate(context.Background()); err != nil {
		_ = db.Close()
		return nil, err
	}

	return instance, nil
}

func (d *DB) Close() error {
	if d == nil || d.db == nil {
		return nil
	}
	return d.db.Close()
}

func (d *DB) Migrate(ctx context.Context) error {
	_, err := d.db.ExecContext(ctx, schemaSQL)
	return err
}

const schemaSQL = `
CREATE TABLE IF NOT EXISTS sync_configs (
	id TEXT PRIMARY KEY,
	local_root TEXT NOT NULL,
	remote_root_id TEXT NOT NULL,
	exclude_patterns TEXT,
	conflict_policy TEXT NOT NULL,
	direction TEXT NOT NULL,
	last_sync_time INTEGER,
	last_change_token TEXT
);

CREATE TABLE IF NOT EXISTS sync_entries (
	config_id TEXT NOT NULL,
	relative_path TEXT NOT NULL,
	drive_file_id TEXT,
	drive_parent_id TEXT,
	is_dir INTEGER NOT NULL DEFAULT 0,
	local_mtime INTEGER,
	local_size INTEGER,
	content_hash TEXT,
	remote_mtime TEXT,
	remote_size INTEGER,
	remote_md5 TEXT,
	remote_mime_type TEXT,
	sync_state TEXT,
	last_sync INTEGER,
	PRIMARY KEY (config_id, relative_path),
	FOREIGN KEY (config_id) REFERENCES sync_configs(id)
);

CREATE INDEX IF NOT EXISTS idx_content_hash ON sync_entries(config_id, content_hash);
CREATE INDEX IF NOT EXISTS idx_drive_file_id ON sync_entries(drive_file_id);
`
