# Implementation progress

## 2026-07-26 — Phase 1 foundation

Implemented:

- monorepo directories and root workflows;
- Go API, worker, and migration commands;
- validated environment configuration;
- structured logging and request IDs;
- stable API error envelope and recovery middleware;
- PostgreSQL pool, first migration, and sqlc configuration;
- health, readiness, system-info, and OpenAPI routes;
- transactional outbox table and locking worker;
- Next.js App Router admin shell and foundation dashboard;
- TypeScript API-client boundary;
- Docker development stack for PostgreSQL, Redis, MinIO, and Mailpit;
- container builds and CI quality/security gates;
- architecture ADRs, UI foundation docs, and project memory.

Verification results are recorded in `known-issues.md` until the final Phase 1
test pass completes.

## 2026-07-26 — Phase 1 foundation hardening

Implemented:

- reproducible npm lockfile and a verified Next.js standalone build;
- self-hosted Geist fonts so production builds do not call Google Fonts;
- explicit, credential-aware CORS with validated exact origins;
- removal of untrusted forwarded-IP handling until trusted proxies are configured;
- concurrency-safe and idempotent migration application;
- generated sqlc access committed and checked for drift in CI;
- one canonical OpenAPI document embedded directly in the API;
- PostgreSQL-backed CI migration smoke tests and concurrent migration coverage;
- a collision-resistant local PostgreSQL host port and Makefile `.env` loading.

Verified locally:

- Go tests and `go vet`;
- npm lint, typecheck, unit tests, lockfile dry-run, and production build;
- backend and admin container builds;
- migration application, repeat application, status, and two concurrent runners
  against isolated Compose databases;
- full Compose smoke tests for API liveness/readiness, canonical OpenAPI,
  allowed and denied CORS, and the containerized admin dashboard.
