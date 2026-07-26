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

In another terminal:

```bash
npm install
npm run dev --workspace=@rehla/admin
```

The API listens on `http://localhost:8080`, the admin on
`http://localhost:3000`, MinIO on `http://localhost:9001`, and Mailpit on
`http://localhost:8025`.

Useful endpoints:

- `GET /healthz` — process liveness
- `GET /readyz` — dependency readiness
- `GET /v1/system/info` — safe public runtime metadata
- `GET /openapi.yaml` — API contract

See [docs/development.md](docs/development.md) for the complete workflow and
[project-memory/README.md](project-memory/README.md) for implementation state.
