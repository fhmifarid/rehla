package migrations

import (
	"context"
	"embed"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed sql/*.up.sql
var files embed.FS

type Runner struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func New(db *pgxpool.Pool, logger *slog.Logger) *Runner {
	return &Runner{db: db, logger: logger}
}

func (r *Runner) Up(ctx context.Context) error {
	if _, err := r.db.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version BIGINT PRIMARY KEY,
			name TEXT NOT NULL,
			applied_at TIMESTAMPTZ NOT NULL DEFAULT now()
		)`); err != nil {
		return fmt.Errorf("create migration table: %w", err)
	}

	entries, err := files.ReadDir("sql")
	if err != nil {
		return err
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })
	for _, entry := range entries {
		version, err := migrationVersion(entry.Name())
		if err != nil {
			return err
		}
		var applied bool
		if err := r.db.QueryRow(ctx,
			`SELECT EXISTS (SELECT 1 FROM schema_migrations WHERE version = $1)`,
			version,
		).Scan(&applied); err != nil {
			return err
		}
		if applied {
			continue
		}
		body, err := files.ReadFile(filepath.Join("sql", entry.Name()))
		if err != nil {
			return err
		}
		if err := r.apply(ctx, version, entry.Name(), string(body)); err != nil {
			return err
		}
		r.logger.Info("migration applied", "version", version, "name", entry.Name())
	}
	return nil
}

func (r *Runner) apply(ctx context.Context, version int64, name, body string) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if _, err := tx.Exec(ctx, `SELECT pg_advisory_xact_lock(731942001)`); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx, body); err != nil {
		return fmt.Errorf("apply migration %s: %w", name, err)
	}
	if _, err := tx.Exec(ctx,
		`INSERT INTO schema_migrations (version, name) VALUES ($1, $2)`,
		version, name,
	); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (r *Runner) Status(ctx context.Context, output io.Writer) error {
	rows, err := r.db.Query(ctx, `
		SELECT version, name, applied_at
		FROM schema_migrations
		ORDER BY version`)
	if err != nil {
		return fmt.Errorf("read migration status (run migrate up first): %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var version int64
		var name string
		var appliedAt any
		if err := rows.Scan(&version, &name, &appliedAt); err != nil {
			return err
		}
		if _, err := fmt.Fprintf(output, "%06d  %s  %v\n", version, name, appliedAt); err != nil {
			return err
		}
	}
	return rows.Err()
}

func migrationVersion(name string) (int64, error) {
	prefix, _, found := strings.Cut(name, "_")
	if !found {
		return 0, fmt.Errorf("invalid migration filename %q", name)
	}
	version, err := strconv.ParseInt(prefix, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid migration version in %q", name)
	}
	return version, nil
}
