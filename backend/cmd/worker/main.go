package main

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/fhmifarid/rehla/backend/internal/config"
	"github.com/fhmifarid/rehla/backend/internal/database"
	"github.com/fhmifarid/rehla/backend/internal/jobs"
	"github.com/fhmifarid/rehla/backend/internal/platform/logging"
	"github.com/fhmifarid/rehla/backend/internal/platform/telemetry"
)

func main() {
	if err := run(); err != nil {
		slog.Error("worker stopped", "error", err)
		os.Exit(1)
	}
}

func run() (runErr error) {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	cfg.ServiceName = "rehla-worker"

	logger := logging.New(cfg.LogLevel, cfg.LogFormat)
	slog.SetDefault(logger)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	shutdownTelemetry, err := telemetry.Setup(ctx, cfg)
	if err != nil {
		return err
	}
	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()
		runErr = errors.Join(runErr, shutdownTelemetry(shutdownCtx))
	}()
	logger.Info("telemetry configured", "enabled", cfg.Telemetry.Enabled)

	pool, err := database.Open(ctx, cfg.Database)
	if err != nil {
		return err
	}
	defer pool.Close()

	worker, err := jobs.NewOutboxWorker(pool, logger, cfg.WorkerPollInterval, cfg.WorkerBatchSize)
	if err != nil {
		return err
	}
	logger.Info("worker started", "poll_interval", cfg.WorkerPollInterval, "batch_size", cfg.WorkerBatchSize)
	return worker.Run(ctx)
}
