import { describe, expect, it } from "vitest";

import { assertValidNavigation, navigation } from "./navigation";

describe("navigation", () => {
  it("contains only unique absolute routes", () => {
    expect(() => assertValidNavigation(navigation)).not.toThrow();
  });

  it("starts with the operational dashboard", () => {
    expect(navigation[0]?.items[0]).toEqual({
      label: "Dashboard",
      href: "/dashboard",
    });
  });
});
