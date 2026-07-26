# Rehla

Rehla is a travel commerce and operations platform. This repository is a
greenfield modular monolith: a Go API and worker own all business rules, a
PostgreSQL database is the transactional source of truth, and a Next.js admin
application consumes the public API contract.

## Prerequisites

- Go 1.26
- Node.js 24 LTS and npm 11
- Docker with Compose v2

## Quick start

```bash
cp .env.example .env
docker compose up -d postgres redis minio mailpit
make migrate-up
make api
```

The Makefile loads the local `.env` file when it exists. Keep real production
secrets outside the repository and inject them through the deployment secret
manager.

In another terminal:

```bash
npm ci
npm run dev --workspace=@rehla/admin
```

The API listens on `http://localhost:58080`, the native admin dev server on
`http://localhost:3000`, PostgreSQL on `localhost:55432`, MinIO on
`http://localhost:9001`, and Mailpit on `http://localhost:8025`. The full
Compose stack exposes the admin container on `http://localhost:53000`.

Useful endpoints:

- `GET /healthz` — process liveness
- `GET /readyz` — dependency readiness
- `GET /v1/system/info` — safe public runtime metadata
- `GET /openapi.yaml` — API contract

See [docs/development.md](docs/development.md) for the complete workflow and
[project-memory/README.md](project-memory/README.md) for implementation state.
