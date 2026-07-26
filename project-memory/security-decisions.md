# Security decisions

- Deny by default. Navigation visibility never substitutes for API
  authorization.
- Authentication, authorization, validation, domain policy, persistence,
  audit, and outbox boundaries are server-side.
- Request IDs accept only a limited safe character set and length.
- Panic details and internal errors are logged, never returned to clients.
- HTTP responses set content-type, framing, referrer, and content-sniffing
  protections. CSP will be specialized for authenticated admin pages.
- Secrets never use `NEXT_PUBLIC_*`, never enter logs, and never live in the
  database as plaintext; database records may hold secret-manager references.
- Production database connections require TLS and private networking.
- Session tokens will be opaque, stored hashed, rotated, revocable, and
  scoped to a client/device. Refresh reuse revokes the token family.
- Sensitive actions require immutable audit context, including actor, reason,
  before/after state where safe, request ID, and time.
- Dependency, secret, static-analysis, test, and container gates run in CI.
