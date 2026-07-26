# API map

## System

| Method | Path | Authentication | Purpose |
|---|---|---|---|
| GET | `/healthz` | None | Process liveness |
| GET | `/readyz` | None | PostgreSQL readiness |
| GET | `/openapi.yaml` | None | Embedded foundation contract |
| GET | `/v1/system/info` | None | Safe runtime metadata |

## Error contract

Errors contain `error.code`, a safe `message`, `request_id`, `retryable`, and
optional field `details`. Internal causes are logged and are never serialized.

The canonical contract is `backend/openapi/openapi.yaml`. It is embedded
directly by the `backend/openapi` package, so the runtime and repository cannot
drift through a duplicated contract file.
