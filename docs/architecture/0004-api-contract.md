# ADR-0004: Contract-first HTTP API

- Status: Accepted
- Date: 2026-07-26

## Decision

Use an OpenAPI 3.1 contract with versioned `/v1` resource routes. All failures
use a stable error envelope with a machine code, safe message, request ID,
retryability, and optional field violations.

Go owns domain rules. Next.js and future Flutter clients never reproduce
authoritative pricing, authorization, inventory, or financial decisions.
