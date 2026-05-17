// Retry-policy behavior parity (port of retry_test.go). Driven against
// the shared behavior server: the scenario engine enforces the failure
// pattern and the test asserts the observable outcome (final status, the
// echoed X-Mock-Attempt ordinal, elapsed time, error surfaced) rather
// than a local request counter. Skips when the server is unconfigured.

import { Client } from "../src/client";
import type { RetryPolicy } from "../src/retry";
import {
  assert,
  assertEq,
  mockBaseURL,
  mockClient,
  mockGet,
  mockPost,
  run,
  scenarioFetch,
  newMockSession,
  skipUnlessMock,
} from "./parity_harness";

const policy = (p: Partial<RetryPolicy>): RetryPolicy => ({
  maxAttempts: 3,
  baseDelay: 1,
  maxDelay: 30,
  retryOnPOST: false,
  ...p,
});

// fail-503-times-2: two 503s then happy; a GET retries past both.
async function retriesOn5xx(): Promise<void> {
  const c = mockClient("fail-503-times-2", {
    retryPolicy: policy({ maxAttempts: 5, baseDelay: 1, maxDelay: 10 }),
  });
  const r = await mockGet(c);
  assertEq("retriesOn5xx: status 200 (retried past two 503s)", r.status, 200);
  assertEq("retriesOn5xx: attempt = 3 (two failures + success)", r.attempt, "3");
  assertEq(
    "retriesOn5xx: branch = fail-exhausted-happy",
    r.branch,
    "fail-exhausted-happy",
  );
}

// RetryOnPOST=false (default): POST 5xx is not retried, but a 429 is
// retried regardless — the server explicitly asked and there is no
// duplicate-creation risk.
// PORTING:S14
async function retriesPOSTOn429(): Promise<void> {
  const c = mockClient("fail-429-times-1", {
    retryPolicy: policy({ maxAttempts: 3, baseDelay: 1 }),
  });
  const r = await mockPost(c);
  assertEq("retriesPOSTOn429: status 200", r.status, 200);
  assertEq("retriesPOSTOn429: attempt = 2 (POST 429 retried once)", r.attempt, "2");
}

// fail-503-times-1 would succeed on a retry, but POST 5xx is not retried
// by default, so the single 503 surfaces.
// PORTING:S14
async function noRetryPOSTByDefault(): Promise<void> {
  const c = mockClient("fail-503-times-1", {
    retryPolicy: policy({ maxAttempts: 5, baseDelay: 1 }),
  });
  const r = await mockPost(c);
  assertEq("noRetryPOSTByDefault: status 503 (POST not retried)", r.status, 503);
  assertEq("noRetryPOSTByDefault: attempt = 1 (no retry)", r.attempt, "1");
}

async function retryPOSTWhenEnabled(): Promise<void> {
  const c = mockClient("fail-503-times-1", {
    retryPolicy: policy({ maxAttempts: 3, baseDelay: 1, retryOnPOST: true }),
  });
  const r = await mockPost(c);
  assertEq("retryPOSTWhenEnabled: status 200", r.status, 200);
  assertEq(
    "retryPOSTWhenEnabled: attempt = 2 (POST retried when enabled)",
    r.attempt,
    "2",
  );
}

async function respectsMaxAttempts(): Promise<void> {
  const c = mockClient("fail-503-times-99", {
    retryPolicy: policy({ maxAttempts: 4, baseDelay: 1, maxDelay: 5 }),
  });
  const r = await mockGet(c);
  assertEq(
    "respectsMaxAttempts: status 503 (final response surfaced)",
    r.status,
    503,
  );
  assertEq(
    "respectsMaxAttempts: attempt = 4 (stopped at maxAttempts)",
    r.attempt,
    "4",
  );
}

async function disabledByZeroPolicy(): Promise<void> {
  // maxAttempts 0 ⇒ retry middleware is a no-op.
  const c = mockClient("fail-503-times-99", {
    retryPolicy: policy({ maxAttempts: 0, baseDelay: 1, maxDelay: 1 }),
  });
  const r = await mockGet(c);
  assertEq(
    "disabledByZeroPolicy: attempt = 1 (zero policy disables retry)",
    r.attempt,
    "1",
  );
}

// retry-after-1: initial 429 carrying retry_in: 1 (seconds), then happy.
// The middleware must honor the body directive over its computed backoff.
// PORTING:S12
async function honorsRetryInBody(): Promise<void> {
  const c = mockClient("retry-after-1", {
    retryPolicy: policy({ maxAttempts: 2, baseDelay: 1, maxDelay: 5000 }),
  });
  const start = Date.now();
  const r = await mockGet(c);
  const elapsed = Date.now() - start;
  assertEq("honorsRetryInBody: status 200 (retried after the 429)", r.status, 200);
  assert(
    "honorsRetryInBody: elapsed >= 900ms (server said retry_in: 1)",
    elapsed >= 900,
    `elapsed=${elapsed}ms`,
  );
}

// Cancellation must surface an error rather than a stale success. Go uses
// a 100ms context against 500ms base backoff; here the timings are shrunk
// (TS sleeps aren't abort-aware) but the asserted contract is identical
// to Go's (`if err == nil { fatal }`): an aborted request never returns
// a clean success. The raw operation is awaited directly — executeRaw
// would try to drain an aborted-and-thus-unusable body — so we just
// assert the call rejected, which is the behavioral contract.
// PORTING:S13
async function retryRespectsCancellation(): Promise<void> {
  const session = newMockSession();
  const c = new Client({
    baseURL: mockBaseURL(),
    fetchApi: scenarioFetch("fail-503-times-99", session),
    retryPolicy: policy({ maxAttempts: 4, baseDelay: 20, maxDelay: 50 }),
  });
  const ac = new AbortController();
  const timer = setTimeout(() => ac.abort(), 50);
  let threw = false;
  try {
    await c.bucketsAPI.getBucketPresignedUrlsRaw({}, { signal: ac.signal });
  } catch {
    threw = true;
  } finally {
    clearTimeout(timer);
  }
  assert(
    "retryRespectsCancellation: aborted request surfaced an error, not a clean success",
    threw,
  );
}

skipUnlessMock("retry parity");
run(async () => {
  await retriesOn5xx();
  await retriesPOSTOn429();
  await noRetryPOSTByDefault();
  await retryPOSTWhenEnabled();
  await respectsMaxAttempts();
  await disabledByZeroPolicy();
  await honorsRetryInBody();
  await retryRespectsCancellation();
});
