package index

import (
	"context"
	"database/sql"
	"encoding/json"
)

func (d *DB) UpsertConfig(ctx context.Context, cfg SyncConfig) error {
	patterns, err := json.Marshal(cfg.ExcludePatterns)
	if err != nil {
		return err
	}

	_, err = d.db.ExecContext(ctx, `
		INSERT INTO sync_configs (
			id, local_root, remote_root_id, exclude_patterns, conflict_policy, direction, last_sync_time, last_change_token
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			local_root=excluded.local_root,
			remote_root_id=excluded.remote_root_id,
			exclude_patterns=excluded.exclude_patterns,
			conflict_policy=excluded.conflict_policy,
			direction=excluded.direction,
			last_sync_time=excluded.last_sync_time,
			last_change_token=excluded.last_change_token
	`, cfg.ID, cfg.LocalRoot, cfg.RemoteRootID, string(patterns), cfg.ConflictPolicy, cfg.Direction, cfg.LastSyncTime, cfg.LastChangeToken)
	return err
}

func (d *DB) GetConfig(ctx context.Context, id string) (*SyncConfig, error) {
	row := d.db.QueryRowContext(ctx, `
		SELECT id, local_root, remote_root_id, exclude_patterns, conflict_policy, direction, last_sync_time, last_change_token
		FROM sync_configs WHERE id = ?
	`, id)

	var cfg SyncConfig
	var patterns string
	err := row.Scan(&cfg.ID, &cfg.LocalRoot, &cfg.RemoteRootID, &patterns, &cfg.ConflictPolicy, &cfg.Direction, &cfg.LastSyncTime, &cfg.LastChangeToken)
	if err != nil {
		return nil, err
	}
	if patterns != "" {
		_ = json.Unmarshal([]byte(patterns), &cfg.ExcludePatterns)
	}
	return &cfg, nil
}

func (d *DB) ListConfigs(ctx context.Context) ([]SyncConfig, error) {
	rows, err := d.db.QueryContext(ctx, `
		SELECT id, local_root, remote_root_id, exclude_patterns, conflict_policy, direction, last_sync_time, last_change_token
		FROM sync_configs ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var configs []SyncConfig
	for rows.Next() {
		var cfg SyncConfig
		var patterns string
		if err := rows.Scan(&cfg.ID, &cfg.LocalRoot, &cfg.RemoteRootID, &patterns, &cfg.ConflictPolicy, &cfg.Direction, &cfg.LastSyncTime, &cfg.LastChangeToken); err != nil {
			return nil, err
		}
		if patterns != "" {
			_ = json.Unmarshal([]byte(patterns), &cfg.ExcludePatterns)
		}
		configs = append(configs, cfg)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return configs, nil
}

func (d *DB) DeleteConfig(ctx context.Context, id string) error {
	_, err := d.db.ExecContext(ctx, `DELETE FROM sync_configs WHERE id = ?`, id)
	return err
}

func (d *DB) ConfigExists(ctx context.Context, id string) (bool, error) {
	row := d.db.QueryRowContext(ctx, `SELECT 1 FROM sync_configs WHERE id = ? LIMIT 1`, id)
	var v int
	if err := row.Scan(&v); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
