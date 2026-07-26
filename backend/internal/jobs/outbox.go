package jobs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/fhmifarid/rehla/backend/internal/platform/logging"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
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
	db            *pgxpool.Pool
	logger        *slog.Logger
	pollInterval  time.Duration
	batchSize     int
	processed     metric.Int64Counter
	batchErrors   metric.Int64Counter
	batchDuration metric.Float64Histogram
}

func NewOutboxWorker(
	db *pgxpool.Pool,
	logger *slog.Logger,
	pollInterval time.Duration,
	batchSize int,
) (*OutboxWorker, error) {
	meter := otel.Meter("github.com/fhmifarid/rehla/backend/internal/jobs")
	processed, err := meter.Int64Counter(
		"rehla.outbox.events.processed",
		metric.WithDescription("Number of outbox events committed as processed."),
		metric.WithUnit("{event}"),
	)
	if err != nil {
		return nil, fmt.Errorf("create outbox processed counter: %w", err)
	}
	batchErrors, err := meter.Int64Counter(
		"rehla.outbox.batch.errors",
		metric.WithDescription("Number of failed outbox processing batches."),
		metric.WithUnit("{error}"),
	)
	if err != nil {
		return nil, fmt.Errorf("create outbox error counter: %w", err)
	}
	batchDuration, err := meter.Float64Histogram(
		"rehla.outbox.batch.duration",
		metric.WithDescription("Outbox batch processing duration."),
		metric.WithUnit("s"),
	)
	if err != nil {
		return nil, fmt.Errorf("create outbox duration histogram: %w", err)
	}

	return &OutboxWorker{
		db:            db,
		logger:        logger,
		pollInterval:  pollInterval,
		batchSize:     batchSize,
		processed:     processed,
		batchErrors:   batchErrors,
		batchDuration: batchDuration,
	}, nil
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

func (w *OutboxWorker) processBatch(ctx context.Context) (err error) {
	started := time.Now()
	ctx, span := otel.Tracer(
		"github.com/fhmifarid/rehla/backend/internal/jobs",
	).Start(ctx, "outbox.process_batch", trace.WithSpanKind(trace.SpanKindConsumer))
	defer func() {
		w.batchDuration.Record(ctx, time.Since(started).Seconds())
		if err != nil {
			w.batchErrors.Add(ctx, 1)
			span.RecordError(err)
			span.SetStatus(codes.Error, "outbox batch failed")
		}
		span.End()
	}()

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
	span.SetAttributes(attribute.Int("rehla.outbox.batch.message_count", len(messages)))

	for _, message := range messages {
		// Phase-one events are durably consumed and logged. Domain-specific
		// dispatchers are registered by later modules before emitting their events.
		logging.WithTraceContext(ctx, w.logger).InfoContext(ctx, "outbox event consumed",
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
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	w.processed.Add(ctx, int64(len(messages)))
	return nil
}
