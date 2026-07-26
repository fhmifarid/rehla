import Link from "next/link";
import type { ReactNode } from "react";

import { navigation } from "@/lib/navigation";

type AdminShellProps = {
  children: ReactNode;
};

export function AdminShell({ children }: AdminShellProps) {
  return (
    <div className="admin-shell">
      <aside className="sidebar" aria-label="Primary navigation">
        <div className="brand">
          <span className="brand-mark" aria-hidden="true">
            R
          </span>
          <span>
            <strong>Rehla</strong>
            <small>Operations</small>
          </span>
        </div>
        <nav className="navigation">
          {navigation.map((group) => (
            <section key={group.label} className="nav-group">
              <h2>{group.label}</h2>
              {group.items.map((item) => (
                <Link
                  className={item.href === "/dashboard" ? "nav-link active" : "nav-link"}
                  href={item.href as never}
                  key={item.href}
                >
                  <span className="nav-dot" aria-hidden="true" />
                  {item.label}
                </Link>
              ))}
            </section>
          ))}
        </nav>
        <div className="sidebar-footer">
          <span className="environment-dot" aria-hidden="true" />
          Local environment
        </div>
      </aside>

      <div className="workspace">
        <header className="topbar">
          <div>
            <p className="eyebrow">Sunday, 26 July</p>
            <p className="welcome">Welcome back, Operations</p>
          </div>
          <div className="topbar-actions">
            <button className="command-trigger" type="button" aria-label="Open command palette">
              Search anything
              <kbd>⌘ K</kbd>
            </button>
            <button className="icon-button" type="button" aria-label="View notifications">
              <span aria-hidden="true">◌</span>
            </button>
            <div className="avatar" aria-label="Signed in as Rehla administrator">
              RA
            </div>
          </div>
        </header>
        <main className="content">{children}</main>
      </div>
    </div>
  );
}
