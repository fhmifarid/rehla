# Architecture

## Runtime

```text
Next.js admin ─┐
Flutter app ───┼──> Go HTTP API ───> PostgreSQL
Integrations ──┘          │
                          └── transactional outbox ──> Go worker

Optional supporting services: Redis, S3-compatible object storage, SMTP.
```

## Style

Rehla is a modular monolith. Each domain will contain explicit domain types,
policies, state transitions, repositories, services, HTTP adapters, events,
and tests where those concepts are meaningful. Platform packages provide
configuration, logging, HTTP contracts, database mechanics, telemetry, and
shared primitives; they do not absorb domain behavior.

Commands that alter sensitive state must authenticate, authorize, validate,
enforce policies and transitions, commit domain/audit/outbox changes in one
database transaction, then allow workers to perform idempotent side effects.

## Source of truth

- Go: business rules and authorization.
- PostgreSQL: transactional state.
- OpenAPI: external HTTP contract.
- `Plan.md`: product requirements.
- `project-memory`: delivery truth and decisions.
