# Financial rules

These invariants apply before financial modules exist:

1. Money is `amount_minor BIGINT` plus ISO 4217 `currency CHAR(3)`; floating
   point is forbidden.
2. Posted financial records are immutable. Corrections use linked reversals.
3. Every journal must balance by currency before commit.
4. Wallet balance is derived from an immutable ledger, never directly set.
5. Available, reserved, pending, and promotional wallet amounts are distinct.
6. Payment and refund commands require idempotency keys and transactional
   uniqueness.
7. Payment allocation, reconciliation, ledger posting, audit, and outbox
   emission share a transaction where the domain invariant requires it.
8. Row locks or compare-and-swap constraints prevent double spending,
   double allocation, and overselling.
9. Maker-checker operations cannot be approved by their maker.
10. Hard deletion of financial evidence is prohibited.

No financial behavior has been implemented in Phase 1.
