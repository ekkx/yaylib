// Behavior-parity harness (port of parity_test.go).
//
// These helpers point the SDK at a shared behavior-contract server — the
// same instance the Go and Python suites drive — so the three languages
// verify identical retry / refresh / rate-limit / event-stream behavior
// instead of three divergent in-process fakes.
//
// The server is discovered from the environment. When the variables are
// unset the parity tests skip (see `skipUnlessMock`), so a standalone
// `npm test` still passes on the pure-function / typed-decode fixtures
// alone.
//
//   YAYLIB_MOCK_BASE_URL   http://host:port   (REST)
//   YAYLIB_MOCK_WS_URL     ws://host:port     (event stream)
//
// Behavior is selected per request with two headers the server reads:
// X-Yay-Mock-Scenario picks the failure/latency behavior and
// X-Yay-Mock-Session isolates each test's stateful counters. The server
// echoes diagnostics back on every response it owns: X-Mock-Branch (the
// decision taken) and X-Mock-Attempt (the per-session attempt ordinal
// for counter-based scenarios).

import { Client, type ClientOptions } from "../src/client";

export const ENV_MOCK_BASE_URL = "YAYLIB_MOCK_BASE_URL";
export const ENV_MOCK_WS_URL = "YAYLIB_MOCK_WS_URL";

export const HDR_SCENARIO = "X-Yay-Mock-Scenario";
export const HDR_SESSION = "X-Yay-Mock-Session";

// Diagnostic response headers (lower-cased for Headers.get, which is
// case-insensitive anyway — kept lower for clarity).
export const HDR_DIAG_BRANCH = "x-mock-branch";
export const HDR_DIAG_ATTEMPT = "x-mock-attempt";

// Real spec routes used to drive generic transport behavior. The happy
// fall-through needs a route the server knows so it answers 200; failure
// scenarios are path-independent. Neither is host-routed, so requests
// stay on the primary base URL. These match parity_test.go's
// parityGETPath / parityPOSTPath exactly.
export const PARITY_GET_PATH = "/v1/buckets/presigned_urls";
export const PARITY_POST_PATH = "/v1/calls/leave_agora_channel";

export function mockBaseURL(): string {
  return process.env[ENV_MOCK_BASE_URL] ?? "";
}

export function mockWSURL(): string {
  return process.env[ENV_MOCK_WS_URL] ?? "";
}

// newMockSession mints an opaque per-test key so the server's stateful
// scenario counters never collide across tests (or across the parallel
// three-language run).
export function newMockSession(): string {
  const b = new Uint8Array(16);
  globalThis.crypto.getRandomValues(b);
  return [...b].map((x) => x.toString(16).padStart(2, "0")).join("");
}

// scenarioFetch is the seam that mirrors Go's scenarioRT installed as the
// http.Client transport. The SDK threads `fetchApi` through every real
// network hop it owns — the generated runtime fetch, the retry replay
// (retry.ts deps.fetchApi), the 401-refresh replay (transport.ts
// ctx.fetchApi), the OAuth /token refresh POST (client._doRefresh) and
// the S3 presigned PUT (client._rawFetch). Wrapping it here therefore
// stamps the scenario + session headers on ALL of those, exactly as the
// Go transport hop does, so the headers ride refresh / retry replays and
// the presigned PUT just like the reference. It adds NO Authorization —
// it sits strictly outside the SDK's own auth, so a no-bearer presigned
// PUT stays no-bearer.
export function scenarioFetch(scenario: string, session: string): typeof fetch {
  return (input: Parameters<typeof fetch>[0], init: RequestInit = {}) => {
    const headers = new Headers(init.headers as HeadersInit | undefined);
    if (scenario) headers.set(HDR_SCENARIO, scenario);
    headers.set(HDR_SESSION, session);
    return fetch(input, { ...init, headers });
  };
}

// mockClient builds a Client pointed at the shared server with `scenario`
// active for this test's isolated session. Extra options are applied
// after the parity defaults so a test can override (e.g. retryPolicy).
export function mockClient(
  scenario: string,
  opts: Partial<ClientOptions> = {},
): Client {
  const base = mockBaseURL();
  const session = newMockSession();
  return new Client({
    baseURL: base,
    fetchApi: scenarioFetch(scenario, session),
    ...opts,
  });
}

// mockStreamClient builds a Client whose event stream targets the shared
// server with the given per-dial WS mode: "" confirms each subscribe and
// pushes the channel's representative event; "reject" rejects every
// subscribe; "no-confirm" stays silent (drives the subscribe timeout);
// "drop-after-confirm" confirms + pushes then closes (drives reconnect /
// re-subscribe). The mode rides the event-stream URL query, which the
// SDK preserves while appending the token and app_version. The REST
// ws-token fetch still rides the scenario seam (scenario "").
export function mockStreamClient(
  mode: string,
  opts: Partial<ClientOptions> = {},
): Client {
  const base = mockBaseURL();
  const session = newMockSession();
  let ws = mockWSURL();
  if (mode) ws += "/?mock=" + mode;
  return new Client({
    baseURL: base,
    eventStreamURL: ws,
    fetchApi: scenarioFetch("", session),
    ...opts,
  });
}

export interface MockOutcome {
  // HTTP status of the final response, or 0 if a transport error
  // surfaced before any response.
  status: number;
  // X-Mock-Branch / X-Mock-Attempt diagnostics off the final response.
  branch: string;
  attempt: string;
  // Final response body length (bytes).
  bodyLen: number;
  // true when the call completed with an HTTP response (2xx or not);
  // false when only a transport error surfaced (no httpResponse).
  hadResponse: boolean;
}

function outcomeOf(r: {
  body?: Uint8Array;
  httpResponse?: Response;
}): MockOutcome {
  const resp = r.httpResponse;
  return {
    status: resp?.status ?? 0,
    branch: resp?.headers.get(HDR_DIAG_BRANCH) ?? "",
    attempt: resp?.headers.get(HDR_DIAG_ATTEMPT) ?? "",
    bodyLen: r.body?.length ?? 0,
    hadResponse: resp !== undefined,
  };
}

// mockGet drives a GET against the shared server's generic primary route
// through the full SDK transport (headers / refresh / retry middleware),
// keeping the diagnostic headers for assertion. executeRaw absorbs typed
// decode and surfaces HTTP errors with the response intact, so it is the
// TS analogue of Go's c.httpClient.Get on the generic route.
export async function mockGet(
  c: Client,
  initOverrides?: RequestInit,
): Promise<MockOutcome> {
  const r = await c.executeRaw(() =>
    c.bucketsAPI.getBucketPresignedUrlsRaw({}, initOverrides),
  );
  return outcomeOf(r);
}

// mockPost drives a POST against the shared server's generic primary
// route through the full SDK transport.
export async function mockPost(
  c: Client,
  initOverrides?: RequestInit,
): Promise<MockOutcome> {
  const r = await c.executeRaw(() =>
    c.callsAPI.leaveAgoraChannelRaw({}, initOverrides),
  );
  return outcomeOf(r);
}

// resetMock clears every scenario counter on the shared server. Tests
// that mint a fresh session per case rarely need it; it exists for the
// few that reuse a session across phases.
export async function resetMock(): Promise<void> {
  const base = mockBaseURL();
  if (!base) return;
  await fetch(base + "/__mock/reset", { method: "POST" }).then(
    (r) => r.arrayBuffer(),
    () => undefined,
  );
}

// ---- standalone test scaffolding (mirrors the other tests/*.ts) -------

let failed = 0;

export function assert(name: string, cond: boolean, detail = ""): void {
  if (cond) {
    // eslint-disable-next-line no-console
    console.log(`PASS ${name}`);
  } else {
    failed++;
    // eslint-disable-next-line no-console
    console.log(`FAIL ${name}${detail ? `\n  ${detail}` : ""}`);
  }
}

export function assertEq<T>(name: string, got: T, want: T): void {
  assert(name, got === want, `got=${String(got)} want=${String(want)}`);
}

// finish exits the process with the standalone-suite convention: 0 when
// every assertion passed (or the suite skipped), non-zero otherwise.
export function finish(): never {
  process.exit(failed === 0 ? 0 : 1);
}

// skipUnlessMock prints a SKIP line and exits 0 when the shared server is
// not configured, so the standalone suite stays green. `needWS` also
// requires the WebSocket URL.
export function skipUnlessMock(suite: string, needWS = false): void {
  if (!mockBaseURL()) {
    // eslint-disable-next-line no-console
    console.log(`SKIP ${suite} (${ENV_MOCK_BASE_URL} unset)`);
    process.exit(0);
  }
  if (needWS && !mockWSURL()) {
    // eslint-disable-next-line no-console
    console.log(`SKIP ${suite} (${ENV_MOCK_WS_URL} unset)`);
    process.exit(0);
  }
}

// run wraps an async suite body with the crash handling the other suites
// use, then exits via finish().
export function run(body: () => Promise<void>): void {
  body()
    .then(finish)
    .catch((e) => {
      // eslint-disable-next-line no-console
      console.error("test crashed:", e);
      process.exit(2);
    });
}
