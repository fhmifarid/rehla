# Architecture decisions

| ADR | Decision | Status |
|---|---|---|
| 0001 | Modular monolith with API and worker processes | Accepted |
| 0002 | PostgreSQL, pgx, sqlc, explicit SQL | Accepted |
| 0003 | PostgreSQL transactional outbox | Accepted |
| 0004 | Contract-first OpenAPI 3.1 HTTP API | Accepted |

Full records are in `docs/architecture/`.

## Foundation baselines

- Go 1.26.x
- PostgreSQL 18.x
- Node.js 24 LTS for builds and production
- Next.js 16.2 patched release line, App Router, strict TypeScript
- Structured `slog` JSON in deployed environments
- Environment-only bootstrap configuration; a secret manager is mandatory in
  production
- No Kubernetes assumption for the initial deployment
