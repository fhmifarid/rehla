# Module status

| Module | Status | Notes |
|---|---|---|
| Platform configuration | Implemented | Startup validation and safe defaults |
| Structured logging | Implemented | JSON/text slog and request correlation |
| HTTP API foundation | Implemented | Routing, recovery, errors, security headers |
| PostgreSQL foundation | Implemented | pgx pool, concurrent migrations, generated sqlc access |
| Worker/outbox | Implemented | Durable polling and locking foundation |
| OpenAPI | Implemented | Canonical 3.1 contract embedded directly |
| Next.js admin | Foundation implemented | Strict shell and dashboard; not authenticated |
| TypeScript API client | Foundation implemented | System-info call and error type |
| Identity and security | Not started | Current next module |
| All product domains | Not started | Follow phase order in specification |

“Implemented” here means the Phase 1 acceptance boundary, not final platform
completion.
