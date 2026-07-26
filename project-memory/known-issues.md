# Known issues

## Open

- The migration runner embeds a synchronized copy of upward migrations because
  Go embed patterns cannot traverse to `backend/migrations`. The check script
  prevents drift; consolidate this when tooling is formalized.
- OpenTelemetry exporters are not wired yet. Structured logs and request IDs
  exist, but traces and metrics remain part of foundation hardening.
- The admin shell exposes reserved navigation visually but has no
  authentication or permission filtering until Phases 2 and 3.
- The command palette trigger is visual only and inaccessible behavior is not
  shipped; interactive behavior belongs to Phase 3.
- Redis, MinIO, and Mailpit are development dependencies only. Their production
  providers and credentials are intentionally undecided.

## Verification

Populate with any failures from the current test and build pass. Do not close
Phase 1 until Go tests/vet, TypeScript tests/typecheck/lint, and Next build pass.
