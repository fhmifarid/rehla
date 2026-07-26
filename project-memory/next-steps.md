# Next steps

1. Triage npm dependency advisories and complete the container/CI verification
   pass.
2. Wire the OpenTelemetry tracing and metrics baseline.
3. Write `docs/modules/identity-security.md` using the module implementation
   contract from the specification.
4. Design identity, session, RBAC, audit, and security-event tables with threat
   modeling and PostgreSQL constraints.
5. Implement the email/password and rotating-session backend slice with
   integration tests.
6. Add staff login and MFA enrollment/verification UI only after the backend
   contract is complete.
7. Update every project-memory map after that milestone.

Do not begin catalog or commerce modules before identity, RBAC, audit, and the
admin foundation are complete.
