// executeRaw parity tests (PORTING.md §11.3). The raw escape hatch must
// return the bytes + HTTP response while bypassing the typed JSON decode:
//   - HTTP success      → { body, httpResponse, error: undefined }
//   - HTTP 4xx/5xx      → { body, httpResponse, error: APIError } and
//                          codeOf(error) still resolves the server code
//   - bad JSON on 2xx   → still returns raw bytes, never throws
//   - transport failure → { error } only

import { createServer, type Server } from "node:http";
import type { AddressInfo } from "node:net";

import { Client } from "../src/client";
import { codeOf, APIError } from "../src/errors";

let failed = 0;
function assert(name: string, cond: boolean, detail = ""): void {
  if (cond) console.log(`PASS ${name}`);
  else {
    failed++;
    console.log(`FAIL ${name}${detail ? `\n  ${detail}` : ""}`);
  }
}

async function withServer(
  reply: { status: number; body: string },
  fn: (baseURL: string) => Promise<void>,
): Promise<void> {
  const server: Server = createServer((_req, res) => {
    res.writeHead(reply.status, { "Content-Type": "application/json" }).end(reply.body);
  });
  await new Promise<void>((r) => server.listen(0, r));
  const baseURL = `http://127.0.0.1:${(server.address() as AddressInfo).port}`;
  try {
    await fn(baseURL);
  } finally {
    server.closeAllConnections?.();
    await new Promise<void>((r) => server.close(() => r()));
  }
}

const NO_RETRY = { maxAttempts: 0, baseDelay: 1, maxDelay: 1, retryOnPOST: false };
const dec = (b?: Uint8Array) => (b ? new TextDecoder().decode(b) : "");

async function successReturnsRawBytes(): Promise<void> {
  const payload = JSON.stringify({ time: 1700000000, ip_address: "1.2.3.4" });
  await withServer({ status: 200, body: payload }, async (baseURL) => {
    const c = new Client({ baseURL, retryPolicy: NO_RETRY });
    const r = await c.executeRaw(() => c.usersAPI.getUserTimestampRaw());
    assert("success: error undefined", r.error === undefined);
    assert("success: httpResponse present + ok", r.httpResponse?.ok === true);
    assert("success: raw bytes are the verbatim body", dec(r.body) === payload, dec(r.body));
  });
}

async function badJsonOnSuccessDoesNotThrow(): Promise<void> {
  await withServer({ status: 200, body: "this is <not> json" }, async (baseURL) => {
    const c = new Client({ baseURL, retryPolicy: NO_RETRY });
    let crashed = false;
    let body = "";
    try {
      const r = await c.executeRaw(() => c.usersAPI.getUserTimestampRaw());
      body = dec(r.body);
    } catch {
      crashed = true;
    }
    assert("bad json on 2xx: never throws (decode bypassed)", !crashed);
    assert("bad json on 2xx: raw bytes preserved", body === "this is <not> json", body);
  });
}

async function httpErrorReturnsAPIErrorWithCode(): Promise<void> {
  const errBody = JSON.stringify({ error_code: 7, message: "forbidden" });
  await withServer({ status: 403, body: errBody }, async (baseURL) => {
    const c = new Client({ baseURL, retryPolicy: NO_RETRY });
    const r = await c.executeRaw(() => c.usersAPI.getUserTimestampRaw());
    assert("http error: error is APIError", r.error instanceof APIError);
    assert("http error: httpResponse present", r.httpResponse?.status === 403, String(r.httpResponse?.status));
    assert("http error: raw error body returned", dec(r.body) === errBody, dec(r.body));
    assert("http error: codeOf(error) === 7", codeOf(r.error) === 7, String(codeOf(r.error)));
  });
}

async function transportFailureReturnsErrorOnly(): Promise<void> {
  // Unroutable port → fetch transport failure.
  const c = new Client({ baseURL: "http://127.0.0.1:1", retryPolicy: NO_RETRY });
  const r = await c.executeRaw(() => c.usersAPI.getUserTimestampRaw());
  assert("transport failure: error present", r.error instanceof APIError);
  assert("transport failure: no httpResponse", r.httpResponse === undefined);
  assert("transport failure: no body", r.body === undefined);
}

(async () => {
  await successReturnsRawBytes();
  await badJsonOnSuccessDoesNotThrow();
  await httpErrorReturnsAPIErrorWithCode();
  await transportFailureReturnsErrorOnly();
  process.exit(failed === 0 ? 0 : 1);
})().catch((e) => {
  console.error("test crashed:", e);
  process.exit(2);
});
