# ADR-0002: PostgreSQL with pgx and sqlc

- Status: Accepted
- Date: 2026-07-26

## Decision

PostgreSQL is the transactional source of truth. Go accesses it with pgx and
sqlc-generated typed queries. Financial, capacity, reconciliation, and
reporting behavior uses explicit SQL and explicit transaction boundaries.

Money is always an integer minor-unit amount paired with an ISO 4217 currency.
Floating-point money is prohibited.
