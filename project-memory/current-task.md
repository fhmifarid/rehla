# Current task

## Phase 2 — Identity and security

Before implementation, write the module contract covering purpose, actors,
stories, admin screens, forms, validation, states, transitions, permissions,
tables, indexes, constraints, endpoints, commands, queries, events, audit,
errors, tests, and acceptance criteria.

The first coherent slice should establish:

1. staff and customer account identities;
2. normalized verified email ownership;
3. password hashing policy and breached-password-safe controls;
4. opaque, rotating sessions and refresh-token reuse detection;
5. email verification and password reset token families;
6. roles, permissions, teams, assignments, and server-side authorization;
7. immutable audit records;
8. authentication rate limiting and security telemetry.

Google authentication and MFA follow after the email/session core is compiling,
migrated, and integration-tested.
