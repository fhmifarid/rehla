package logging

import (
	"context"
	"log/slog"
	"os"
	"strings"

	"go.opentelemetry.io/otel/trace"
)

func New(level, format string) *slog.Logger {
	var parsed slog.Level
	switch strings.ToLower(level) {
	case "debug":
		parsed = slog.LevelDebug
	case "warn":
		parsed = slog.LevelWarn
	case "error":
		parsed = slog.LevelError
	default:
		parsed = slog.LevelInfo
	}

	options := &slog.HandlerOptions{Level: parsed}
	if format == "text" {
		return slog.New(slog.NewTextHandler(os.Stdout, options))
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, options))
}

func WithTraceContext(ctx context.Context, logger *slog.Logger) *slog.Logger {
	spanContext := trace.SpanContextFromContext(ctx)
	if !spanContext.IsValid() {
		return logger
	}
	return logger.With(
		"trace_id", spanContext.TraceID().String(),
		"span_id", spanContext.SpanID().String(),
		"trace_sampled", spanContext.IsSampled(),
	)
}
