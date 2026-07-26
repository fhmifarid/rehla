package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/rehla-platform/rehla/backend/internal/config"
	"github.com/rehla-platform/rehla/backend/internal/database"
	"github.com/rehla-platform/rehla/backend/internal/database/migrations"
	"github.com/rehla-platform/rehla/backend/internal/platform/logging"
)

func main() {
	if err := run(); err != nil {
		slog.Error("migration command failed", "error", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("usage: migrate [up|status]")
	}

	cfg, err := config.Load()
	if err != nil {
		return err
	}
	logger := logging.New(cfg.LogLevel, cfg.LogFormat)
	slog.SetDefault(logger)

	ctx := context.Background()
	pool, err := database.Open(ctx, cfg.Database)
	if err != nil {
		return err
	}
	defer pool.Close()

	runner := migrations.New(pool, logger)
	switch os.Args[1] {
	case "up":
		return runner.Up(ctx)
	case "status":
		return runner.Status(ctx, os.Stdout)
	default:
		return fmt.Errorf("unknown migration command %q", os.Args[1])
	}
}
