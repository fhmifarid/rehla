package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Database struct {
	URL            string
	MaxConnections int32
	MinConnections int32
	ConnectTimeout time.Duration
}

type Config struct {
	Environment        string
	ServiceName        string
	HTTPAddr           string
	Database           Database
	LogLevel           string
	LogFormat          string
	ShutdownTimeout    time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	IdleTimeout        time.Duration
	WorkerPollInterval time.Duration
	WorkerBatchSize    int
	AllowedOrigins     []string
}

func Load() (Config, error) {
	cfg := Config{
		Environment: env("REHLA_ENV", "local"),
		ServiceName: env("REHLA_SERVICE_NAME", "rehla-api"),
		HTTPAddr:    env("REHLA_HTTP_ADDR", ":8080"),
		Database: Database{
			URL: env("REHLA_DATABASE_URL", ""),
		},
		LogLevel:  env("REHLA_LOG_LEVEL", "info"),
		LogFormat: env("REHLA_LOG_FORMAT", "json"),
	}

	var err error
	if cfg.Database.MaxConnections, err = int32Env("REHLA_DATABASE_MAX_CONNS", 20); err != nil {
		return Config{}, err
	}
	if cfg.Database.MinConnections, err = int32Env("REHLA_DATABASE_MIN_CONNS", 2); err != nil {
		return Config{}, err
	}
	if cfg.WorkerBatchSize, err = intEnv("REHLA_WORKER_BATCH_SIZE", 50); err != nil {
		return Config{}, err
	}
	if cfg.Database.ConnectTimeout, err = durationEnv("REHLA_DATABASE_CONNECT_TIMEOUT", 5*time.Second); err != nil {
		return Config{}, err
	}
	if cfg.ShutdownTimeout, err = durationEnv("REHLA_SHUTDOWN_TIMEOUT", 15*time.Second); err != nil {
		return Config{}, err
	}
	if cfg.ReadTimeout, err = durationEnv("REHLA_READ_TIMEOUT", 10*time.Second); err != nil {
		return Config{}, err
	}
	if cfg.WriteTimeout, err = durationEnv("REHLA_WRITE_TIMEOUT", 30*time.Second); err != nil {
		return Config{}, err
	}
	if cfg.IdleTimeout, err = durationEnv("REHLA_IDLE_TIMEOUT", 60*time.Second); err != nil {
		return Config{}, err
	}
	if cfg.WorkerPollInterval, err = durationEnv("REHLA_WORKER_POLL_INTERVAL", 2*time.Second); err != nil {
		return Config{}, err
	}

	origins := env("REHLA_ALLOWED_ORIGINS", "http://localhost:3000")
	for _, origin := range strings.Split(origins, ",") {
		if trimmed := strings.TrimSpace(origin); trimmed != "" {
			cfg.AllowedOrigins = append(cfg.AllowedOrigins, trimmed)
		}
	}

	if err := cfg.validate(); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (c Config) validate() error {
	switch c.Environment {
	case "local", "test", "staging", "production":
	default:
		return fmt.Errorf("REHLA_ENV must be local, test, staging, or production")
	}
	if c.Database.URL == "" {
		return fmt.Errorf("REHLA_DATABASE_URL is required")
	}
	if c.Database.MaxConnections < 1 || c.Database.MinConnections < 0 ||
		c.Database.MinConnections > c.Database.MaxConnections {
		return fmt.Errorf("database connection bounds are invalid")
	}
	if c.WorkerBatchSize < 1 || c.WorkerBatchSize > 1000 {
		return fmt.Errorf("REHLA_WORKER_BATCH_SIZE must be between 1 and 1000")
	}
	if c.LogFormat != "json" && c.LogFormat != "text" {
		return fmt.Errorf("REHLA_LOG_FORMAT must be json or text")
	}
	return nil
}

func env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func durationEnv(key string, fallback time.Duration) (time.Duration, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback, nil
	}
	parsed, err := time.ParseDuration(value)
	if err != nil || parsed <= 0 {
		return 0, fmt.Errorf("%s must be a positive duration", key)
	}
	return parsed, nil
}

func intEnv(key string, fallback int) (int, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback, nil
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be an integer", key)
	}
	return parsed, nil
}

func int32Env(key string, fallback int32) (int32, error) {
	parsed, err := intEnv(key, int(fallback))
	return int32(parsed), err
}
