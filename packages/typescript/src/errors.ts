// PORTING.md §9: three layers of error access — the wire-level APIError,
// the parsed ErrorResponse body, and the typed ErrorCode for switch /
// dispatch.

import { ResponseError, FetchError } from "./gen/runtime";

// APIError carries the raw HTTP body and status. typescript-fetch already
// throws `ResponseError` (a non-2xx HTTP response) and `FetchError` (a
// transport failure); we wrap both so callers have one stable name to
// match against.
export class APIError extends Error {
  readonly response?: Response;
  readonly body: Uint8Array;
  readonly cause?: unknown;

  constructor(opts: {
    message: string;
    response?: Response;
    body?: Uint8Array;
    cause?: unknown;
  }) {
    super(opts.message);
    this.name = "APIError";
    this.response = opts.response;
    this.body = opts.body ?? new Uint8Array();
    this.cause = opts.cause;
  }

  bodyText(): string {
    return new TextDecoder().decode(this.body);
  }
}

// ErrorResponse is the parsed shape of the Yay! server's JSON error body.
// All fields are optional because the server may emit a subset depending
// on the error code.
export interface ErrorResponse {
  errorCode?: number;
  message?: string;
  banUntil?: number;
  retryIn?: number;
  remainingQuota?: number;
}

// errorResponseOf returns the parsed payload when err is an APIError with
// a JSON body, otherwise null.
export function errorResponseOf(err: unknown): ErrorResponse | null {
  if (!(err instanceof APIError) || err.body.length === 0) return null;
  try {
    const parsed = JSON.parse(err.bodyText()) as Record<string, unknown>;
    const out: ErrorResponse = {};
    if (typeof parsed.error_code === "number") out.errorCode = parsed.error_code;
    if (typeof parsed.message === "string") out.message = parsed.message;
    if (typeof parsed.ban_until === "number") out.banUntil = parsed.ban_until;
    if (typeof parsed.retry_in === "number") out.retryIn = parsed.retry_in;
    if (typeof parsed.remaining_quota === "number")
      out.remainingQuota = parsed.remaining_quota;
    return out;
  } catch {
    return null;
  }
}

// codeOf surfaces the server's error_code for switch dispatch. Returns 0
// (ErrCodeUnknown placeholder) when the error is not an APIError or carries
// no recognizable code. The typed ErrorCode constants are generated —
// until that lands the port returns numeric codes only.
export function codeOf(err: unknown): number {
  const r = errorResponseOf(err);
  return r?.errorCode ?? 0;
}

// asAPIError converts a typescript-fetch thrown error (ResponseError or
// FetchError) into our stable APIError shape. Async because reading the
// response body is async.
export async function asAPIError(err: unknown): Promise<APIError> {
  if (err instanceof APIError) return err;
  if (err instanceof ResponseError) {
    const body = err.response ? new Uint8Array(await err.response.arrayBuffer()) : new Uint8Array();
    return new APIError({
      message: err.message || err.response?.statusText || "API error",
      response: err.response,
      body,
      cause: err,
    });
  }
  if (err instanceof FetchError) {
    return new APIError({
      message: err.message,
      cause: err.cause,
    });
  }
  return new APIError({
    message: err instanceof Error ? err.message : String(err),
    cause: err,
  });
}
