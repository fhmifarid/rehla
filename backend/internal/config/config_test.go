package config

import "testing"

func TestLoadRequiresDatabaseURL(t *testing.T) {
	t.Setenv("REHLA_DATABASE_URL", "")
	if _, err := Load(); err == nil {
		t.Fatal("expected missing database URL to fail")
	}
}

func TestLoadParsesConfiguration(t *testing.T) {
	t.Setenv("REHLA_DATABASE_URL", "postgres://example")
	t.Setenv("REHLA_WORKER_BATCH_SIZE", "75")
	t.Setenv("REHLA_ALLOWED_ORIGINS", "https://admin.example, https://ops.example")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg.WorkerBatchSize != 75 {
		t.Fatalf("WorkerBatchSize = %d, want 75", cfg.WorkerBatchSize)
	}
	if len(cfg.AllowedOrigins) != 2 {
		t.Fatalf("AllowedOrigins length = %d, want 2", len(cfg.AllowedOrigins))
	}
}
