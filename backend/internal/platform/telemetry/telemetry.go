package telemetry

import (
	"context"
	"errors"
	"fmt"

	"github.com/fhmifarid/rehla/backend/internal/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const (
	serviceVersion       = "0.1.0"
	maxExportRequestSize = 4 << 20
)

type Shutdown func(context.Context) error

func Setup(ctx context.Context, cfg config.Config) (Shutdown, error) {
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	if !cfg.Telemetry.Enabled {
		return func(context.Context) error { return nil }, nil
	}

	res, err := newResource(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("build telemetry resource: %w", err)
	}

	traceExporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpointURL(cfg.Telemetry.Endpoint),
		otlptracehttp.WithTimeout(cfg.Telemetry.ExportTimeout),
		otlptracehttp.WithCompression(otlptracehttp.GzipCompression),
		otlptracehttp.WithMaxRequestSize(maxExportRequestSize),
	)
	if err != nil {
		return nil, fmt.Errorf("create OTLP trace exporter: %w", err)
	}

	metricExporter, err := otlpmetrichttp.New(ctx,
		otlpmetrichttp.WithEndpointURL(cfg.Telemetry.Endpoint),
		otlpmetrichttp.WithTimeout(cfg.Telemetry.ExportTimeout),
		otlpmetrichttp.WithCompression(otlpmetrichttp.GzipCompression),
		otlpmetrichttp.WithMaxRequestSize(maxExportRequestSize),
	)
	if err != nil {
		_ = traceExporter.Shutdown(ctx)
		return nil, fmt.Errorf("create OTLP metric exporter: %w", err)
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.ParentBased(
			sdktrace.TraceIDRatioBased(cfg.Telemetry.TraceSampleRatio),
		)),
		sdktrace.WithBatcher(traceExporter),
	)
	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(
			metricExporter,
			sdkmetric.WithInterval(cfg.Telemetry.MetricExportInterval),
			sdkmetric.WithTimeout(cfg.Telemetry.ExportTimeout),
		)),
	)

	otel.SetTracerProvider(tracerProvider)
	otel.SetMeterProvider(meterProvider)

	return func(shutdownCtx context.Context) error {
		return errors.Join(
			meterProvider.Shutdown(shutdownCtx),
			tracerProvider.Shutdown(shutdownCtx),
		)
	}, nil
}

func newResource(ctx context.Context, cfg config.Config) (*resource.Resource, error) {
	return resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithTelemetrySDK(),
		resource.WithOS(),
		resource.WithContainer(),
		resource.WithHost(),
		resource.WithProcessPID(),
		resource.WithProcessExecutableName(),
		resource.WithProcessRuntimeName(),
		resource.WithProcessRuntimeVersion(),
		resource.WithAttributes(
			attribute.String("service.name", cfg.ServiceName),
			attribute.String("service.version", serviceVersion),
			attribute.String("deployment.environment.name", cfg.Environment),
		),
	)
}
