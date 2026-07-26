# Admin design system

The admin interface is an operational tool: dense enough for expert workflows,
but calm, legible, and keyboard accessible.

## Principles

- Server Components by default; client components only for browser interaction.
- Every control has an accessible name and visible focus state.
- Status is never communicated with color alone.
- List state belongs in URL parameters where practical.
- English LTR and Arabic RTL are first-class layout directions.
- Destructive, reversible, and approval-gated actions are visually distinct.

Shared patterns will include application shell, page heading, metric card,
data table, filters, saved views, form sections, confirmation dialogs, empty
states, error states, and skeleton loading.
