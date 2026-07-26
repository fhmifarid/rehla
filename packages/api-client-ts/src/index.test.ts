import { describe, expect, it, vi } from "vitest";

import { createRehlaClient } from "./index";

describe("createRehlaClient", () => {
  it("normalizes the base URL", async () => {
    const request = vi.fn<typeof fetch>().mockResolvedValue(
      new Response(JSON.stringify({
        name: "Rehla API",
        environment: "test",
        version: "0.1.0",
      }), { status: 200 }),
    );
    const client = createRehlaClient({
      baseURL: "https://api.example/",
      fetch: request,
    });

    await client.systemInfo();

    expect(request).toHaveBeenCalledWith(
      "https://api.example/v1/system/info",
      { signal: undefined },
    );
  });
});
