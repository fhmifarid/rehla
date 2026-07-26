# ADR-0003: Transactional outbox

- Status: Accepted
- Date: 2026-07-26

## Decision

Domain changes and their events are committed in the same PostgreSQL
transaction. Workers claim pending events with `FOR UPDATE SKIP LOCKED`.
Consumers must become idempotent before external side effects are introduced.

The outbox is not an audit log. Audit records are immutable actor-facing
evidence; outbox records are delivery infrastructure.
