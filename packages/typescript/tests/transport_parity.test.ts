// Transport 401-chain behavior parity (port of the parity-eligible
// transport_test.go scenarios). The in-process header-injection /
// OAuth-Basic / lazy-client-IP / refresh-then-network-error cases stay
// local fixtures by design (see auth_refresh.test.ts etc.); only the two
// shared-server scenarios are translated here.
//
// §15 says translate the scenario, not the Go API mechanics: the
// Bearer-header progression is a transport internal, so the behavioral
// assertion is final success + the seeded tokens rotated to the
// server-issued ones (expired-token), and the original 401 body
// preserved when refresh itself fails (fail-401-times-2).

import { assert, assertEq, mockClient, mockGet, run, skipUnlessMock } from "./parity_harness";

const noRetry = { maxAttempts: 0, baseDelay: 1, maxDelay: 1, retryOnPOST: false };

// expired-token: protected requests 401 (error_code -3) until the token
// endpoint is hit (a simulated refresh), then the happy path. This is the
// 401 → refresh → retry chain.
// PORTING:S7
async function refreshAndRetriesOn401(): Promise<void> {
  const c = mockClient("expired-token", { retryPolicy: noRetry });
  c.setTokens("STALE", "REF");

  const r = await mockGet(c);
  assertEq(
    "401 refresh+retry: final status 200 (refresh + retry succeeded)",
    r.status,
    200,
  );
  const tok = c.tokens();
  assert(
    "401 refresh+retry: access token rotated to server-issued value",
    tok.access !== "STALE" && tok.access !== "",
    `access=${tok.access}`,
  );
  assert(
    "401 refresh+retry: refresh token rotated to server-issued value",
    tok.refresh !== "REF" && tok.refresh !== "",
    `refresh=${tok.refresh}`,
  );
}

// fail-401-times-2: the data request 401s and the refresh call (same
// session+scenario counter) also 401s, so refresh fails and the original
// 401 surfaces with its body intact.
// PORTING:S8
async function refreshFailureSurfaces401(): Promise<void> {
  const c = mockClient("fail-401-times-2", { retryPolicy: noRetry });
  c.setTokens("STALE", "REF");

  const r = await mockGet(c);
  assertEq(
    "401 refresh-fail: status 401 (refresh failed, original 401 surfaced)",
    r.status,
    401,
  );
  assert(
    "401 refresh-fail: original 401 error body preserved",
    r.bodyLen > 0,
    `bodyLen=${r.bodyLen}`,
  );
}

skipUnlessMock("transport 401 parity");
run(async () => {
  await refreshAndRetriesOn401();
  await refreshFailureSurfaces401();
});
