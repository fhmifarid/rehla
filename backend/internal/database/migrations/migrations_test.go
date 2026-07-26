package migrations

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestMigrationVersion(t *testing.T) {
	t.Parallel()

	version, err := migrationVersion("000123_add_orders.up.sql")
	if err != nil {
		t.Fatalf("migrationVersion() error = %v", err)
	}
	if version != 123 {
		t.Fatalf("migrationVersion() = %d, want 123", version)
	}
}

func TestMigrationVersionRejectsInvalidNames(t *testing.T) {
	t.Parallel()

	for _, name := range []string{
		"missing-separator.sql",
		"invalid_name.up.sql",
		"000000_zero.up.sql",
		"-1_negative.up.sql",
	} {
		name := name
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if _, err := migrationVersion(name); err == nil {
				t.Fatalf("migrationVersion(%q) unexpectedly succeeded", name)
			}
		})
	}
}

func TestEmbeddedMigrationsMatchCanonical(t *testing.T) {
	t.Parallel()

	canonicalFiles, err := filepath.Glob(filepath.Join("..", "..", "..", "migrations", "*.up.sql"))
	if err != nil {
		t.Fatalf("find canonical migrations: %v", err)
	}
	embeddedEntries, err := files.ReadDir("sql")
	if err != nil {
		t.Fatalf("read embedded migrations: %v", err)
	}
	if len(canonicalFiles) != len(embeddedEntries) {
		t.Fatalf(
			"canonical migration count = %d, embedded count = %d",
			len(canonicalFiles),
			len(embeddedEntries),
		)
	}

	for _, canonicalFile := range canonicalFiles {
		name := filepath.Base(canonicalFile)
		canonicalBody, err := os.ReadFile(canonicalFile)
		if err != nil {
			t.Fatalf("read canonical migration %s: %v", name, err)
		}
		embeddedBody, err := files.ReadFile(path.Join("sql", name))
		if err != nil {
			t.Fatalf("read embedded migration %s: %v", name, err)
		}
		if !bytes.Equal(canonicalBody, embeddedBody) {
			t.Fatalf("embedded migration differs from canonical migration: %s", name)
		}
	}
}

func TestRunnerUpIsConcurrentAndIdempotent(t *testing.T) {
	databaseURL := os.Getenv("REHLA_TEST_DATABASE_URL")
	if databaseURL == "" {
		t.Skip("REHLA_TEST_DATABASE_URL is not configured")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	adminPool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		t.Fatalf("create admin pool: %v", err)
	}
	t.Cleanup(adminPool.Close)

	schema := fmt.Sprintf("migration_test_%d", time.Now().UnixNano())
	identifier := pgx.Identifier{schema}.Sanitize()
	if _, err := adminPool.Exec(ctx, "CREATE SCHEMA "+identifier); err != nil {
		t.Fatalf("create test schema: %v", err)
	}
	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cleanupCancel()
		if _, err := adminPool.Exec(cleanupCtx, "DROP SCHEMA "+identifier+" CASCADE"); err != nil {
			t.Errorf("drop test schema: %v", err)
		}
	})

	poolConfig, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		t.Fatalf("parse pool configuration: %v", err)
	}
	poolConfig.ConnConfig.RuntimeParams["search_path"] = schema
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		t.Fatalf("create migration pool: %v", err)
	}
	t.Cleanup(pool.Close)

	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	results := make(chan error, 2)
	for range 2 {
		go func() {
			results <- New(pool, logger).Up(ctx)
		}()
	}
	for range 2 {
		if err := <-results; err != nil {
			t.Fatalf("concurrent Up() error = %v", err)
		}
	}

	var migrationCount int
	if err := pool.QueryRow(ctx, "SELECT count(*) FROM schema_migrations").Scan(&migrationCount); err != nil {
		t.Fatalf("count migrations: %v", err)
	}
	if migrationCount != 1 {
		t.Fatalf("migration count = %d, want 1", migrationCount)
	}

	var outboxTable string
	if err := pool.QueryRow(ctx, "SELECT to_regclass('outbox_events')::text").Scan(&outboxTable); err != nil {
		t.Fatalf("find outbox table: %v", err)
	}
	if outboxTable != "outbox_events" {
		t.Fatalf("outbox table = %q, want outbox_events", outboxTable)
	}
}
