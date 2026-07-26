import type { Metadata } from "next";

import { AdminShell } from "@/components/admin-shell";

export const metadata: Metadata = {
  title: "Dashboard",
};

const metrics = [
  { label: "Gross sales", value: "EGP 0", delta: "Awaiting commerce data", tone: "emerald" },
  { label: "Active orders", value: "0", delta: "Order module is next", tone: "blue" },
  { label: "Pending reviews", value: "0", delta: "No payment proofs", tone: "amber" },
  { label: "Upcoming departures", value: "0", delta: "No departures scheduled", tone: "violet" },
];

export default function DashboardPage() {
  return (
    <AdminShell>
      <section className="page-heading">
        <div>
          <p className="eyebrow">Operational overview</p>
          <h1>Dashboard</h1>
          <p className="subtitle">
            A unified view of commerce, finance, and travel operations.
          </p>
        </div>
        <div className="heading-actions">
          <button className="button secondary" type="button">
            Last 30 days
          </button>
          <button className="button primary" type="button">
            Create order
          </button>
        </div>
      </section>

      <section className="metrics-grid" aria-label="Key metrics">
        {metrics.map((metric) => (
          <article className="metric-card" key={metric.label}>
            <div className={`metric-icon ${metric.tone}`} aria-hidden="true" />
            <p>{metric.label}</p>
            <strong>{metric.value}</strong>
            <small>{metric.delta}</small>
          </article>
        ))}
      </section>

      <section className="dashboard-grid">
        <article className="panel revenue-panel">
          <div className="panel-heading">
            <div>
              <p className="panel-kicker">Revenue</p>
              <h2>Sales performance</h2>
            </div>
            <span className="status-pill">Foundation mode</span>
          </div>
          <div className="chart-empty" role="img" aria-label="Revenue chart awaiting order data">
            <div className="chart-grid" />
            <div className="chart-line" />
            <p>Revenue data will appear after the order module is enabled.</p>
          </div>
        </article>

        <article className="panel">
          <div className="panel-heading">
            <div>
              <p className="panel-kicker">System</p>
              <h2>Foundation health</h2>
            </div>
          </div>
          <ul className="health-list">
            <li>
              <span className="health-check">✓</span>
              <span><strong>Go API</strong><small>Health and readiness routes</small></span>
              <span className="status-pill success">Ready</span>
            </li>
            <li>
              <span className="health-check">✓</span>
              <span><strong>PostgreSQL</strong><small>Migration and sqlc foundation</small></span>
              <span className="status-pill success">Ready</span>
            </li>
            <li>
              <span className="health-check">✓</span>
              <span><strong>Worker</strong><small>Transactional outbox consumer</small></span>
              <span className="status-pill success">Ready</span>
            </li>
          </ul>
        </article>
      </section>

      <section className="panel roadmap-panel">
        <div className="panel-heading">
          <div>
            <p className="panel-kicker">Delivery</p>
            <h2>Platform rollout</h2>
          </div>
          <span className="phase-label">Phase 1 of 18</span>
        </div>
        <div className="progress-track" aria-label="Phase 1 complete">
          <span />
        </div>
        <div className="milestones">
          <span><i className="done" />Foundation</span>
          <span><i />Identity & Security</span>
          <span><i />Admin Foundation</span>
          <span><i />Catalog</span>
          <span><i />Travel Operations</span>
        </div>
      </section>
    </AdminShell>
  );
}
