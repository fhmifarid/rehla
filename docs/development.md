# Development workflow

## Local services

Copy `.env.example` to `.env`. The example contains development-only
credentials and no production secrets.

```bash
docker compose up -d postgres redis minio mailpit
make migrate-up
```

Run `make api`, `make worker`, and `npm run dev` in separate terminals. Run the
entire container stack with `docker compose up --build`.

## Quality checks

```bash
make test
make lint
make build
```

Generate typed database access after changing queries or schema:

```bash
make sqlc
```

Migration files are append-only after deployment. The embedded copy used by
the migration binary lives in `backend/internal/database/migrations/sql`; CI
must verify it matches `backend/migrations` until the runner is extracted into
a dedicated build-time package.

## Configuration

All backend variables use the `REHLA_` prefix. Configuration is read at process
startup and validated before network listeners or workers start. Production
secrets must come from the deployment secret manager, never from committed
files or `NEXT_PUBLIC_*` variables.

## API changes

The canonical contract is `backend/openapi/openapi.yaml`. API changes require:

1. updating the contract;
2. implementing server behavior;
3. updating generated clients;
4. adding contract and behavior tests;
5. updating `project-memory/api-map.md`.
