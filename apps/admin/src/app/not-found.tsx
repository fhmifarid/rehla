import Link from "next/link";

export default function NotFound() {
  return (
    <main className="not-found">
      <p className="eyebrow">404</p>
      <h1>This module has not been activated yet.</h1>
      <p>The route is reserved in the platform map and will arrive in its planned phase.</p>
      <Link className="button primary" href="/dashboard">
        Return to dashboard
      </Link>
    </main>
  );
}
