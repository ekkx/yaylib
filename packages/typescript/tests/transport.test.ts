// Transport behavior parity (PORTING.md §15 #6a / #6b / #10 / #9).
//
// These exercise the request-shaping the SDK guarantees on every
// outbound call against an in-process HTTP server stub:
//
//   #6a  the required headers (User-Agent / X-App-Version / X-Device-*
//        / X-Connection-* / Accept-Language / X-Timestamp) ride every
//        request.
//   #6b  the OAuth token endpoint gets `Authorization: Basic`, while a
//        normal endpoint gets `Authorization: Bearer` once tokens are
//        set.
//   #10  the caller's outbound IP is learned lazily: the first request
//        carries no X-Client-IP, and a later request does once the SDK
//        has observed it from a server response.
//   #9   a 401 that refreshes successfully but whose retried original
//        then hits a transport error surfaces that transport error —
//        never a stale 401.

import { createServer, type Server } from "node:http";
import type { AddressInfo } from "node:net";

import { Client } from "../src/client";

let failed = 0;
function assert(name: string, cond: boolean, detail = ""): void {
  if (cond) console.log(`PASS ${name}`);
  else {
    failed++;
    console.log(`FAIL ${name}${detail ? `\n  ${detail}` : ""}`);
  }
}

interface Captured {
  method: string;
  path: string;
  headers: Record<string, string>;
}

type Reply = { status: number; body: string; headers?: Record<string, string> };
type HandlerFn = (req: Captured, body: string) => Reply | "destroy";

// withServer runs an in-process server that records every request's
// method / path / headers, then defers the response to the handler. A
// handler returning "destroy" kills the socket so the client sees a
// transport-level failure (the TS analogue of the Go test hijacking and
// closing the TCP connection on the post-refresh retry).
async function withServer(
  handler: HandlerFn,
  fn: (baseURL: string, captured: Captured[]) => Promise<void>,
): Promise<void> {
  const captured: Captured[] = [];
  const server: Server = createServer((req, res) => {
    const chunks: Buffer[] = [];
    req.on("data", (c) => chunks.push(c));
    req.on("end", () => {
      const headers: Record<string, string> = {};
      for (const [k, v] of Object.entries(req.headers)) {
        if (typeof v === "string") headers[k.toLowerCase()] = v;
        else if (Array.isArray(v)) headers[k.toLowerCase()] = v.join(", ");
      }
      const rec: Captured = {
        method: req.method ?? "GET",
        path: (req.url ?? "").split("?")[0],
        headers,
      };
      captured.push(rec);
      const out = handler(rec, Buffer.concat(chunks).toString("utf8"));
      if (out === "destroy") {
        req.socket.destroy();
        return;
      }
      res.writeHead(out.status, {
        "Content-Type": "application/json",
        ...(out.headers ?? {}),
      });
      res.end(out.body);
    });
  });
  await new Promise<void>((r) => server.listen(0, r));
  const baseURL = `http://127.0.0.1:${(server.address() as AddressInfo).port}`;
  try {
    await fn(baseURL, captured);
  } finally {
    server.closeAllConnections?.();
    await new Promise<void>((r) => server.close(() => r()));
  }
}

const NO_RETRY = { maxAttempts: 0, baseDelay: 1, maxDelay: 1, retryOnPOST: false };

function get(headers: Record<string, string>, key: string): string {
  return headers[key.toLowerCase()] ?? "";
}

// #6a — every request carries the required headers; with tokens set the
// Authorization scheme is Bearer on a normal endpoint.
async function requiredHeadersInjected(): Promise<void> {
  await withServer(
    () => ({ status: 200, body: "{}" }),
    async (baseURL, captured) => {
      const c = new Client({ baseURL, retryPolicy: NO_RETRY });
      c.setTokens("ACC", "REF");
      await c.bucketsAPI.getBucketPresignedUrls({ fileNames: [] });

      const h = captured[0].headers;
      for (const name of [
        "User-Agent",
        "X-App-Version",
        "X-Device-Info",
        "X-Device-UUID",
        "X-Connection-Type",
        "X-Connection-Speed",
        "Accept-Language",
        "X-Timestamp",
      ]) {
        assert(`required header present: ${name}`, get(h, name) !== "", `${name} empty`);
      }
      assert(
        "Authorization is Bearer on a normal endpoint",
        get(h, "Authorization") === "Bearer ACC",
        get(h, "Authorization"),
      );
    },
  );
}

// #6b — the OAuth token endpoint gets Basic; a normal endpoint gets
// Bearer. Asserted on the same client so the scheme selection is the
// transport's, not the caller's.
async function oauthEndpointUsesBasicAuth(): Promise<void> {
  await withServer(
    (req) => {
      if (req.path === "/api/v1/oauth/token") {
        return {
          status: 200,
          body: JSON.stringify({ access_token: "A", refresh_token: "R" }),
        };
      }
      return { status: 200, body: "{}" };
    },
    async (baseURL, captured) => {
      const c = new Client({ baseURL, retryPolicy: NO_RETRY });
      c.setTokens("ACC", "REF");

      await c.authAPI.oauthToken({ grantType: "refresh_token", refreshToken: "R" });
      await c.bucketsAPI.getBucketPresignedUrls({ fileNames: [] });

      const oauthReq = captured.find((r) => r.path === "/api/v1/oauth/token");
      const normalReq = captured.find((r) => r.path === "/v1/buckets/presigned_urls");
      assert("oauth request captured", oauthReq !== undefined);
      assert("normal request captured", normalReq !== undefined);

      const oauthAuth = get(oauthReq!.headers, "Authorization");
      const expectedBasic = `Basic ${Buffer.from(c.apiKey).toString("base64")}`;
      assert(
        "oauth endpoint Authorization is Basic <base64(apiKey)>",
        oauthAuth === expectedBasic,
        `got=${oauthAuth} want=${expectedBasic}`,
      );
      assert(
        "normal endpoint Authorization is Bearer after setTokens",
        get(normalReq!.headers, "Authorization") === "Bearer ACC",
        get(normalReq!.headers, "Authorization"),
      );
    },
  );
}

// #10 — lazy X-Client-IP. The first request carries none; the SDK learns
// the IP from the timestamp lookup response; a later request carries it.
async function lazyClientIPPopulatesHeader(): Promise<void> {
  const FETCHED_IP = "203.0.113.7";
  await withServer(
    (req) => {
      if (req.path.endsWith("/v2/users/timestamp")) {
        return {
          status: 200,
          body: JSON.stringify({ time: 0, ip_address: FETCHED_IP }),
        };
      }
      return { status: 200, body: "{}" };
    },
    async (baseURL, captured) => {
      const c = new Client({ baseURL, retryPolicy: NO_RETRY });

      // First request kicks off the background IP lookup; it goes out
      // without X-Client-IP.
      await c.bucketsAPI.getBucketPresignedUrls({ fileNames: [] });

      // Wait until the SDK has observed the IP (or time out).
      const deadline = Date.now() + 2000;
      while (Date.now() < deadline && c.clientIP !== FETCHED_IP) {
        await new Promise((r) => setTimeout(r, 20));
      }
      assert(
        "client IP learned from lookup response",
        c.clientIP === FETCHED_IP,
        `clientIP=${c.clientIP}`,
      );

      // Second request now carries the learned X-Client-IP.
      await c.bucketsAPI.getBucketPresignedUrls({ fileNames: [] });

      const bucketReqs = captured.filter(
        (r) => r.path === "/v1/buckets/presigned_urls",
      );
      const first = bucketReqs[0];
      const second = bucketReqs[bucketReqs.length - 1];
      assert(
        "first request carries no X-Client-IP (lookup is async)",
        get(first.headers, "X-Client-IP") === "",
        get(first.headers, "X-Client-IP"),
      );
      assert(
        "second request carries the learned X-Client-IP",
        get(second.headers, "X-Client-IP") === FETCHED_IP,
        get(second.headers, "X-Client-IP"),
      );
    },
  );
}

// #9 — 401 → refresh succeeds → retried original hits a transport error.
// The surfaced error must be the retry/network error, not a stale 401.
async function refreshOkButRetryNetworkErrorSurfaces(): Promise<void> {
  let dataAttempt = 0;
  await withServer(
    (req): Reply | "destroy" => {
      if (req.path === "/api/v1/oauth/token") {
        return {
          status: 200,
          body: JSON.stringify({
            access_token: "NEW_ACC",
            refresh_token: "NEW_REF",
          }),
        };
      }
      dataAttempt++;
      if (dataAttempt === 1) {
        return { status: 401, body: '{"error_code":-3,"message":"expired"}' };
      }
      // Post-refresh retry: kill the socket so fetch rejects with a
      // transport error.
      return "destroy";
    },
    async (baseURL) => {
      const c = new Client({ baseURL, retryPolicy: NO_RETRY });
      c.setTokens("STALE", "REF");

      let threw = false;
      let status401 = false;
      try {
        await c.bucketsAPI.getBucketPresignedUrls({ fileNames: [] });
      } catch (err) {
        threw = true;
        // A surfaced ResponseError/APIError carrying a 401 would mean
        // the stale 401 leaked — assert it did not.
        const anyErr = err as { response?: { status?: number } };
        status401 = anyErr.response?.status === 401;
      }
      assert("broken retry surfaces an error", threw);
      assert(
        "surfaced error is the transport failure, not a stale 401",
        threw && !status401,
      );
      assert(
        "tokens stayed rotated after the successful refresh",
        c.tokens().access === "NEW_ACC",
        `access=${c.tokens().access}`,
      );
    },
  );
}

(async () => {
  await requiredHeadersInjected();
  await oauthEndpointUsesBasicAuth();
  await lazyClientIPPopulatesHeader();
  await refreshOkButRetryNetworkErrorSurfaces();
  process.exit(failed === 0 ? 0 : 1);
})().catch((e) => {
  console.error("test crashed:", e);
  process.exit(2);
});
