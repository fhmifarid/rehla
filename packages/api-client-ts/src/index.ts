export type APIErrorPayload = {
  error: {
    code: string;
    message: string;
    request_id?: string;
    retryable: boolean;
    details?: Array<{ field: string; message: string }>;
  };
};

export class RehlaAPIError extends Error {
  readonly code: string;
  readonly requestId?: string;
  readonly retryable: boolean;

  constructor(payload: APIErrorPayload["error"]) {
    super(payload.message);
    this.name = "RehlaAPIError";
    this.code = payload.code;
    this.requestId = payload.request_id;
    this.retryable = payload.retryable;
  }
}

export type RehlaClientOptions = {
  baseURL: string;
  fetch?: typeof globalThis.fetch;
};

export function createRehlaClient(options: RehlaClientOptions) {
  const request = options.fetch ?? globalThis.fetch;
  const baseURL = options.baseURL.replace(/\/$/, "");

  return {
    async systemInfo(signal?: AbortSignal) {
      const response = await request(`${baseURL}/v1/system/info`, { signal });
      if (!response.ok) {
        throw new RehlaAPIError((await response.json() as APIErrorPayload).error);
      }
      return response.json() as Promise<{
        name: string;
        environment: "local" | "test" | "staging" | "production";
        version: string;
      }>;
    },
  };
}
