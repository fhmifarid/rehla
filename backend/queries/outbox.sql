-- name: EnqueueOutboxEvent :one
INSERT INTO outbox_events (
    event_type,
    aggregate_type,
    aggregate_id,
    payload,
    available_at
) VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetOutboxEvent :one
SELECT *
FROM outbox_events
WHERE id = $1;
