package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/rehla-platform/rehla/backend/internal/config"
	"github.com/rehla-platform/rehla/backend/internal/database"
	"github.com/rehla-platform/rehla/backend/internal/jobs"
	"github.com/rehla-platform/rehla/backend/internal/platform/logging"
)

func main() {
	if err := run(); err != nil {
		slog.Error("worker stopped", "error", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	cfg.ServiceName = "rehla-worker"

	logger := logging.New(cfg.LogLevel, cfg.LogFormat)
	slog.SetDefault(logger)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	pool, err := database.Open(ctx, cfg.Database)
	if err != nil {
		return err
	}
	defer pool.Close()

	worker := jobs.NewOutboxWorker(pool, logger, cfg.WorkerPollInterval, cfg.WorkerBatchSize)
	logger.Info("worker started", "poll_interval", cfg.WorkerPollInterval, "batch_size", cfg.WorkerBatchSize)
	return worker.Run(ctx)
}
