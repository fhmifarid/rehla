# Known issues

## Open

- The migration runner embeds a synchronized copy of upward migrations because
  Go embed patterns cannot traverse to `backend/migrations`. CI now verifies all
  upward files; consolidate this when migration tooling is formalized.
- OpenTelemetry exporters are not wired yet. Structured logs and request IDs
  exist, but traces and metrics remain part of foundation hardening.
- The admin shell exposes reserved navigation visually but has no
  authentication or permission filtering until Phases 2 and 3.
- The command palette trigger is visual only and inaccessible behavior is not
  shipped; interactive behavior belongs to Phase 3.
- Redis, MinIO, and Mailpit are development dependencies only. Their production
  providers and credentials are intentionally undecided.
- npm installation reports dependency advisories that still require explicit
  audit triage. Do not run automatic forced upgrades; review affected runtime
  paths and patched versions first.

## Verification

Go tests/vet, TypeScript tests/typecheck/lint, the Next standalone build, and
both container builds pass locally. The remaining Phase 1 gates are
dependency-advisory triage, OpenTelemetry wiring, and the remote CI pass.
