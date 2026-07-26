# ADR-0001: Modular monolith

- Status: Accepted
- Date: 2026-07-26

## Context

Rehla has many closely related transactional domains. Payments, capacity,
orders, wallet movements, and journals frequently need one PostgreSQL
transaction and one audit boundary.

## Decision

Use a modular monolith. The API and worker are separate Go processes built from
one module and backed by one PostgreSQL database. Domain packages may depend on
shared platform contracts, but not on another module's persistence details.

## Consequences

Cross-domain invariants can be enforced atomically. Deployments stay simple.
Modules must maintain explicit boundaries to avoid a distributed monolith
inside one process. A service may only be extracted after operational evidence
shows an independent scaling or isolation need.
