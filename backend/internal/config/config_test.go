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
	t.Setenv("REHLA_ALLOWED_ORIGINS", "https://admin.example, https://ops.example, https://admin.example")

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

func TestLoadRejectsInvalidOrigin(t *testing.T) {
	t.Setenv("REHLA_DATABASE_URL", "postgres://example")

	for _, origin := range []string{"https://admin.example/path", "https://admin.example/"} {
		t.Run(origin, func(t *testing.T) {
			t.Setenv("REHLA_ALLOWED_ORIGINS", origin)
			if _, err := Load(); err == nil {
				t.Fatalf("expected origin %q to fail", origin)
			}
		})
	}
}

func TestLoadRequiresHTTPSOriginsInProduction(t *testing.T) {
	t.Setenv("REHLA_ENV", "production")
	t.Setenv("REHLA_DATABASE_URL", "postgres://example")
	t.Setenv("REHLA_ALLOWED_ORIGINS", "http://admin.example")

	if _, err := Load(); err == nil {
		t.Fatal("expected an insecure production origin to fail")
	}
}
