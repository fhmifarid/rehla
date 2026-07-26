# Database map

## Applied in Phase 1

### `schema_migrations`

Owned by migration infrastructure. Records an append-only migration version,
name, and application timestamp.

### `outbox_events`

Durable domain-event delivery queue.

Important columns: UUID `id`, `event_type`, aggregate identity, JSONB payload,
occurrence/availability/processing timestamps, attempt count, and last error.
Pending events have a partial index on `(available_at, created_at)`.

Workers claim rows using `FOR UPDATE SKIP LOCKED`. Domain modules must insert an
event in the same transaction as the corresponding state change.

## Extensions

- `pgcrypto` for database-side UUID generation
- `pg_trgm` for the planned global-search baseline

No customer, identity, financial, or inventory tables exist yet.
