package telemetry

import (
	"context"
	"testing"

	"github.com/fhmifarid/rehla/backend/internal/config"
)

func TestSetupDisabled(t *testing.T) {
	shutdown, err := Setup(context.Background(), config.Config{})
	if err != nil {
		t.Fatalf("Setup() error = %v", err)
	}
	if err := shutdown(context.Background()); err != nil {
		t.Fatalf("shutdown() error = %v", err)
	}
}

func TestResourceContainsServiceIdentity(t *testing.T) {
	cfg := config.Config{
		ServiceName: "rehla-test",
		Environment: "test",
	}
	res, err := newResource(context.Background(), cfg)
	if err != nil {
		t.Fatalf("newResource() error = %v", err)
	}

	attributes := make(map[string]string)
	for _, item := range res.Attributes() {
		attributes[string(item.Key)] = item.Value.AsString()
	}
	for key, want := range map[string]string{
		"service.name":                "rehla-test",
		"service.version":             serviceVersion,
		"deployment.environment.name": "test",
	} {
		if got := attributes[key]; got != want {
			t.Errorf("%s = %q, want %q", key, got, want)
		}
	}
}
