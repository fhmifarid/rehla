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
	t.Setenv("REHLA_OTEL_ENABLED", "true")
	t.Setenv("REHLA_OTEL_EXPORTER_OTLP_ENDPOINT", "http://collector:4318")
	t.Setenv("REHLA_OTEL_TRACE_SAMPLE_RATIO", "0.25")

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
	if !cfg.Telemetry.Enabled {
		t.Fatal("Telemetry.Enabled = false, want true")
	}
	if cfg.Telemetry.TraceSampleRatio != 0.25 {
		t.Fatalf("Telemetry.TraceSampleRatio = %f, want 0.25", cfg.Telemetry.TraceSampleRatio)
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

func TestLoadRejectsInvalidTelemetryConfiguration(t *testing.T) {
	tests := map[string]struct {
		endpoint string
		ratio    string
	}{
		"missing endpoint": {ratio: "0.1"},
		"invalid endpoint": {endpoint: "collector:4318", ratio: "0.1"},
		"invalid ratio":    {endpoint: "https://collector.example", ratio: "1.1"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Setenv("REHLA_DATABASE_URL", "postgres://example")
			t.Setenv("REHLA_OTEL_ENABLED", "true")
			t.Setenv("REHLA_OTEL_EXPORTER_OTLP_ENDPOINT", test.endpoint)
			t.Setenv("REHLA_OTEL_TRACE_SAMPLE_RATIO", test.ratio)

			if _, err := Load(); err == nil {
				t.Fatal("expected invalid telemetry configuration to fail")
			}
		})
	}
}

func TestLoadRequiresHTTPSForProductionTelemetry(t *testing.T) {
	t.Setenv("REHLA_ENV", "production")
	t.Setenv("REHLA_DATABASE_URL", "postgres://example")
	t.Setenv("REHLA_ALLOWED_ORIGINS", "https://admin.example")
	t.Setenv("REHLA_OTEL_ENABLED", "true")
	t.Setenv("REHLA_OTEL_EXPORTER_OTLP_ENDPOINT", "http://collector.example:4318")

	if _, err := Load(); err == nil {
		t.Fatal("expected an insecure production telemetry endpoint to fail")
	}
}
