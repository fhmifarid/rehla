package jobs

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type outboxMessage struct {
	ID            string
	EventType     string
	AggregateType string
	AggregateID   string
	Payload       json.RawMessage
	Attempts      int
}

type OutboxWorker struct {
	db           *pgxpool.Pool
	logger       *slog.Logger
	pollInterval time.Duration
	batchSize    int
}

func NewOutboxWorker(db *pgxpool.Pool, logger *slog.Logger, pollInterval time.Duration, batchSize int) *OutboxWorker {
	return &OutboxWorker{db: db, logger: logger, pollInterval: pollInterval, batchSize: batchSize}
}

func (w *OutboxWorker) Run(ctx context.Context) error {
	ticker := time.NewTicker(w.pollInterval)
	defer ticker.Stop()

	for {
		if err := w.processBatch(ctx); err != nil && !errors.Is(err, context.Canceled) {
			w.logger.Error("outbox batch failed", "error", err)
		}
		select {
		case <-ctx.Done():
			w.logger.Info("worker shutdown complete")
			return nil
		case <-ticker.C:
		}
	}
}

func (w *OutboxWorker) processBatch(ctx context.Context) error {
	tx, err := w.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	rows, err := tx.Query(ctx, `
		SELECT id::text, event_type, aggregate_type, aggregate_id, payload, attempts
		FROM outbox_events
		WHERE processed_at IS NULL AND available_at <= now()
		ORDER BY created_at
		FOR UPDATE SKIP LOCKED
		LIMIT $1`, w.batchSize)
	if err != nil {
		return err
	}

	var messages []outboxMessage
	for rows.Next() {
		var message outboxMessage
		if err := rows.Scan(&message.ID, &message.EventType, &message.AggregateType,
			&message.AggregateID, &message.Payload, &message.Attempts); err != nil {
			rows.Close()
			return err
		}
		messages = append(messages, message)
	}
	if err := rows.Err(); err != nil {
		rows.Close()
		return err
	}
	rows.Close()

	for _, message := range messages {
		// Phase-one events are durably consumed and logged. Domain-specific
		// dispatchers are registered by later modules before emitting their events.
		w.logger.Info("outbox event consumed",
			"event_id", message.ID,
			"event_type", message.EventType,
			"aggregate_type", message.AggregateType,
			"aggregate_id", message.AggregateID,
		)
		if _, err := tx.Exec(ctx, `
			UPDATE outbox_events
			SET processed_at = now(), attempts = attempts + 1, updated_at = now()
			WHERE id = $1`, message.ID); err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
}
