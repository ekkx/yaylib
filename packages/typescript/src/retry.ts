// PORTING.md §13: retry policy for transient failures (429, 5xx,
// transport errors). The wire contract:
//
//   - 429 retries on every method — the server explicitly asked.
//   - 5xx / transport errors retry only on safe methods (GET / HEAD /
//     PUT / DELETE / OPTIONS) by default. POST / PATCH retry only when
//     `retryOnPOST` is true because the Yay! server has no
//     idempotency-key concept.
//   - The response body's `retry_in` field, when present, is honored
//     over the computed exponential backoff (seconds).
//   - Backoff is exponential with full jitter, clamped to `maxDelay`.
//   - The zero value (`{ maxAttempts: 0 }`) disables retry entirely.

import type { Middleware, ResponseContext } from "./gen/runtime.js";
import { type Logger, noopLogger } from "./logger.js";

export interface RetryPolicy {
  // Total attempt cap (initial attempt + retries). 0 disables retry.
  // Default: 3.
  maxAttempts: number;
  // Base delay before the first retry. Default: 200ms.
  baseDelay: number;
  // Hard ceiling on any single sleep. Default: 30_000ms.
  maxDelay: number;
  // Allow retrying POST / PATCH responses on 5xx / transport errors.
  // Default: false. 429 retries on every method regardless.
  retryOnPOST: boolean;
}

export const DEFAULT_RETRY_POLICY: RetryPolicy = {
  maxAttempts: 3,
  baseDelay: 200,
  maxDelay: 30_000,
  retryOnPOST: false,
};

const SAFE_METHODS = new Set([
  "GET",
  "HEAD",
  "PUT",
  "DELETE",
  "OPTIONS",
]);

function normalizeMethod(m: string | undefined): string {
  return (m ?? "GET").toUpperCase();
}

// methodAllowsRetry returns true if the method may retry on 5xx /
// transport errors under the given policy. 429 is handled separately
// (it retries on every method).
function methodAllowsRetry(method: string, policy: RetryPolicy): boolean {
  if (SAFE_METHODS.has(method)) return true;
  if (method === "POST" || method === "PATCH") return policy.retryOnPOST;
  return false;
}

// shouldRetryStatus returns true for an HTTP response code that warrants
// a retry. 429 retries on every method; 5xx retries when the method
// allows it (per methodAllowsRetry).
function shouldRetryStatus(
  status: number,
  method: string,
  policy: RetryPolicy,
): boolean {
  if (status === 429) return true;
  if (status >= 500 && status <= 599) return methodAllowsRetry(method, policy);
  return false;
}

// fullJitterDelay returns a random duration in [0, min(maxDelay,
// baseDelay * 2^(attempt-1))]. `attempt` is the upcoming retry's
// 1-indexed attempt number (1 = first retry).
function fullJitterDelay(
  attempt: number,
  policy: RetryPolicy,
  random: () => number = Math.random,
): number {
  const exp = Math.min(
    policy.maxDelay,
    policy.baseDelay * Math.pow(2, Math.max(0, attempt - 1)),
  );
  return Math.floor(random() * exp);
}

// retryAfterFromBody returns the server-supplied wait duration parsed
// out of the response body, in milliseconds. The Yay! server uses the
// `retry_in` field on CommonErrorResponse in seconds. Returns null when
// not present / unparseable; the caller falls back to the computed
// backoff.
async function retryAfterFromBody(
  response: Response,
): Promise<number | null> {
  // Clone so the caller's response stays readable.
  try {
    const cloned = response.clone();
    const text = await cloned.text();
    if (!text) return null;
    const parsed = JSON.parse(text) as Record<string, unknown>;
    const v = parsed["retry_in"];
    if (typeof v === "number" && Number.isFinite(v) && v >= 0) {
      return Math.floor(v * 1000);
    }
    return null;
  } catch {
    return null;
  }
}

// sleep resolves after `ms` milliseconds. Exposed for tests that want
// to substitute a deterministic clock.
export function sleep(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

export interface RetryDeps {
  policy: RetryPolicy;
  logger?: Logger;
  random?: () => number;
  sleepFn?: (ms: number) => Promise<void>;
  // Raw fetch used for retry attempts. MUST be the unwrapped fetch
  // (configuration.fetchApi or globalThis.fetch), NOT the
  // middleware-wrapped ResponseContext.fetch — re-entering the wrapped
  // fetch would recursively re-run this very middleware and loop
  // forever on a permanently-failing endpoint.
  fetchApi?: typeof fetch;
}

// buildRetryMiddleware wires the retry loop into typescript-fetch.
// Two hooks contribute:
//   - post: when fetch returned a response, we may retry on 429 / 5xx.
//   - onError: when fetch threw a transport error, we may retry on
//     methods that allow retry. The recovered response then re-enters
//     the post hook chain, so a network failure followed by a 5xx will
//     trigger both loops naturally.
//
// Each loop owns its own attempt counter; in the worst case a request
// can spend up to (maxAttempts) attempts on network errors AND
// (maxAttempts) more on 429/5xx. This matches the Go reference, which
// also separates the two loops.
export function buildRetryMiddleware(deps: RetryDeps): Middleware {
  const policy = deps.policy;
  const logger = deps.logger ?? noopLogger;
  const random = deps.random ?? Math.random;
  const sleepFn = deps.sleepFn ?? sleep;

  const pathOf = (u: string): string => {
    try {
      return new URL(u).pathname;
    } catch {
      return u;
    }
  };
  const rawFetch = deps.fetchApi ?? globalThis.fetch;

  // Re-issue the original request with a freshened X-Timestamp (the
  // value baked in by the headers middleware is now stale). Uses the
  // unwrapped fetch so the retry does not recurse through the middleware
  // chain.
  const refetch = (url: string, init: RequestInit): Promise<Response> => {
    const headers: Record<string, string> = {
      ...((init.headers as Record<string, string>) ?? {}),
    };
    for (const k of Object.keys(headers)) {
      if (k.toLowerCase() === "x-timestamp") delete headers[k];
    }
    headers["X-Timestamp"] = String(Math.floor(Date.now() / 1000));
    return rawFetch(url, { ...init, headers });
  };

  if (policy.maxAttempts <= 1) {
    // Zero-value semantics: middleware is a no-op so callers explicitly
    // opt out of retries with `new Client({ retryPolicy: { maxAttempts: 0 } })`.
    return {};
  }

  return {
    post: async (rc: ResponseContext): Promise<Response | void> => {
      const method = normalizeMethod(rc.init.method);
      let response = rc.response;
      if (!shouldRetryStatus(response.status, method, policy)) return undefined;

      // attempt = 1 corresponds to the very first retry (the initial
      // attempt that produced `rc.response` doesn't count toward the
      // budget here).
      let attempt = 1;
      while (attempt < policy.maxAttempts) {
        const bodyWait = await retryAfterFromBody(response);
        // Body's retry_in is a server directive — honored as-is. The
        // computed backoff is already clamped inside fullJitterDelay,
        // so we only clamp again on that path.
        const delay =
          bodyWait !== null
            ? bodyWait
            : fullJitterDelay(attempt, policy, random);
        logger.debug("retrying request", {
          event: "http_retry",
          method,
          path: pathOf(rc.url),
          attempt,
          status: response.status,
          delay_ms: delay,
        });
        await sleepFn(delay);

        let next: Response;
        try {
          next = await refetch(rc.url, rc.init);
        } catch {
          // Transport failure mid-retry — surface the latest non-2xx
          // response we already have. The caller's typed-fetch wrapper
          // will turn this into a ResponseError on its own.
          return response;
        }
        response = next;
        attempt++;
        if (!shouldRetryStatus(response.status, method, policy)) {
          return response;
        }
      }
      return response;
    },
    onError: async (ec): Promise<Response | void> => {
      const method = normalizeMethod(ec.init.method);
      if (!methodAllowsRetry(method, policy)) return undefined;

      let attempt = 1;
      while (attempt < policy.maxAttempts) {
        const delay = Math.min(fullJitterDelay(attempt, policy, random), policy.maxDelay);
        logger.debug("retrying request", {
          event: "http_retry",
          method,
          path: pathOf(ec.url),
          attempt,
          error: "transport",
          delay_ms: delay,
        });
        await sleepFn(delay);
        attempt++;
        try {
          // Returning a Response (any status, including 5xx) hands
          // control back to the post-middleware chain — the post hook
          // then takes over for 429 / 5xx retries.
          return await refetch(ec.url, ec.init);
        } catch {
          if (attempt >= policy.maxAttempts) {
            // Out of budget — let the original transport error
            // propagate so the caller sees the wire failure.
            return undefined;
          }
        }
      }
      return undefined;
    },
  };
}
