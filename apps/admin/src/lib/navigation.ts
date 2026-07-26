export type NavigationItem = {
  label: string;
  href: string;
};

export type NavigationGroup = {
  label: string;
  items: NavigationItem[];
};

export const navigation: NavigationGroup[] = [
  {
    label: "Home",
    items: [
      { label: "Dashboard", href: "/dashboard" },
      { label: "Activity", href: "/activity" },
      { label: "My Tasks", href: "/tasks" },
    ],
  },
  {
    label: "Commerce",
    items: [
      { label: "Orders", href: "/orders" },
      { label: "Customers", href: "/customers" },
      { label: "Packages", href: "/packages" },
      { label: "Products & Add-ons", href: "/products" },
      { label: "Pricing", href: "/pricing" },
      { label: "Promotions", href: "/promotions" },
    ],
  },
  {
    label: "Travel Operations",
    items: [
      { label: "Departures", href: "/departures" },
      { label: "Availability", href: "/availability" },
      { label: "Reservations", href: "/reservations" },
      { label: "Travelers", href: "/travelers" },
      { label: "Manifests", href: "/manifests" },
    ],
  },
  {
    label: "Finance",
    items: [
      { label: "Overview", href: "/finance" },
      { label: "Payments", href: "/payments" },
      { label: "Reconciliation", href: "/reconciliation" },
      { label: "Wallets", href: "/wallets" },
      { label: "Refunds", href: "/refunds" },
      { label: "Financial Ledger", href: "/ledger" },
    ],
  },
  {
    label: "Customer Experience",
    items: [
      { label: "Support Inbox", href: "/support" },
      { label: "Tickets", href: "/tickets" },
      { label: "Notifications", href: "/notifications" },
    ],
  },
  {
    label: "Platform",
    items: [
      { label: "Analytics", href: "/analytics" },
      { label: "Automation", href: "/automation" },
      { label: "Security", href: "/security" },
      { label: "System", href: "/system" },
      { label: "Settings", href: "/settings" },
    ],
  },
];

export function assertValidNavigation(groups: NavigationGroup[]): void {
  const paths = new Set<string>();
  for (const group of groups) {
    if (!group.label || group.items.length === 0) {
      throw new Error("Every navigation group requires a label and items.");
    }
    for (const item of group.items) {
      if (!item.href.startsWith("/")) {
        throw new Error(`Navigation path must be absolute: ${item.href}`);
      }
      if (paths.has(item.href)) {
        throw new Error(`Duplicate navigation path: ${item.href}`);
      }
      paths.add(item.href);
    }
  }
}
