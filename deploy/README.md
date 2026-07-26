# Deployment

The initial production topology is a load balancer in front of replicated Go
API and Next.js instances, plus independently scaled workers. PostgreSQL,
object storage, and optional Redis should be managed services.

`compose.yaml` is for local development, not production. Production requires
TLS termination, a secret manager, private database networking, centralized
logs and traces, automated PostgreSQL backups with point-in-time recovery, and
tested restore procedures. Kubernetes is intentionally not assumed.
