// PORTING.md §6.1 contract verification with an in-memory fetch stub.
// We exercise the three documented outcomes plus the concurrency
// collapse, mirroring Go's transport_test.go scenarios.
//
// Run with `npm test -- auth_refresh` or directly via tsx.

import { createServer, type Server } from "node:http";
import type { AddressInfo } from "node:net";

import { Client } from "../src/client";

interface ServerHandler {
  (path: string, method: string, body: string): {
    status: number;
    body: string;
    headers?: Record<string, string>;
  };
}

async function withServer<T>(
  handler: ServerHandler,
  fn: (baseURL: string) => Promise<T>,
): Promise<T> {
  const server: Server = createServer((req, res) => {
    const chunks: Buffer[] = [];
    req.on("data", (c) => chunks.push(c));
    req.on("end", () => {
      const body = Buffer.concat(chunks).toString("utf8");
      const out = handler(req.url ?? "", req.method ?? "GET", body);
      res.writeHead(out.status, {
        "Content-Type": "application/json",
        ...(out.headers ?? {}),
      });
      res.end(out.body);
    });
  });
  await new Promise<void>((resolve) => server.listen(0, resolve));
  const addr = server.address() as AddressInfo;
  const baseURL = `http://127.0.0.1:${addr.port}`;
  try {
    return await fn(baseURL);
  } finally {
    await new Promise<void>((resolve) => server.close(() => resolve()));
  }
}

let failed = 0;
function assertEq<T>(name: string, got: T, want: T): void {
  if (got !== want) {
    failed++;
    // eslint-disable-next-line no-console
    console.log(`FAIL ${name}\n  got=${String(got)}\n  want=${String(want)}`);
  } else {
    // eslint-disable-next-line no-console
    console.log(`PASS ${name}`);
  }
}

// Scenario 1: 401 → refresh success → retry succeeds.
// PORTING:S7
async function scenarioRefreshSuccess(): Promise<void> {
  let oauthHits = 0;
  let timelineHits = 0;
  let lastAuth = "";
  await withServer(
    (path, method, body) => {
      if (path === "/api/v1/oauth/token") {
        oauthHits++;
        // Accept the form-encoded refresh request.
        if (!body.includes("grant_type=refresh_token")) {
          return { status: 400, body: '{"error":"bad_grant"}' };
        }
        return {
          status: 200,
          body: JSON.stringify({
            access_token: "fresh-access",
            refresh_token: "fresh-refresh",
            user_id: 7,
          }),
        };
      }
      // Mock GET /v2/users/timestamp — return 401 the first time, 200 after.
      if (path.includes("/v2/users/timestamp")) {
        timelineHits++;
        // Just look at request method/body — we'll inspect the Authorization
        // via the body callback shape isn't easy, so instead infer state.
        if (timelineHits === 1) return { status: 401, body: '{"error_code":-1}' };
        return {
          status: 200,
          body: JSON.stringify({ time: 1700000000, ip_address: "203.0.113.5" }),
        };
      }
      return { status: 404, body: "" };
    },
    async (baseURL) => {
      const client = new Client({ baseURL });
      client.setTokens("stale-access", "stale-refresh");
      // GetUserTimestamp lives on UsersApi in the generated layer.
      const res = await client.usersAPI.getUserTimestamp();
      assertEq(
        "scenario 1: timeline returns refreshed payload",
        res.time,
        1700000000,
      );
      assertEq("scenario 1: oauth hit exactly once", oauthHits, 1);
      assertEq(
        "scenario 1: timeline hit twice (stale + retry)",
        timelineHits,
        2,
      );
      assertEq("scenario 1: access token rotated", client.tokens().access, "fresh-access");
      assertEq(
        "scenario 1: refresh token rotated",
        client.tokens().refresh,
        "fresh-refresh",
      );
      void lastAuth;
    },
  );
}

// Scenario 2: 401 → refresh fails (server returns 401 on /token) → original
// 401 surfaces to the caller.
// PORTING:S8
async function scenarioRefreshFails(): Promise<void> {
  await withServer(
    (path, _method, _body) => {
      if (path === "/api/v1/oauth/token") {
        return {
          status: 401,
          body: '{"error_code":-100,"message":"refresh expired"}',
        };
      }
      if (path.includes("/v2/users/timestamp")) {
        return {
          status: 401,
          body: '{"error_code":-1,"message":"unauthorized"}',
        };
      }
      return { status: 404, body: "" };
    },
    async (baseURL) => {
      const client = new Client({ baseURL });
      client.setTokens("stale-access", "stale-refresh");
      let thrownStatus = -1;
      try {
        await client.usersAPI.getUserTimestamp();
      } catch (err) {
        // typescript-fetch throws ResponseError for non-2xx.
        const r = (err as { response?: Response }).response;
        if (r) thrownStatus = r.status;
      }
      assertEq("scenario 2: original 401 surfaced", thrownStatus, 401);
      assertEq(
        "scenario 2: tokens unchanged after failed refresh",
        client.tokens().access,
        "stale-access",
      );
    },
  );
}

// Scenario 3: no refresh token (client has only an access token) — refresh
// is skipped entirely and the original 401 propagates.
// PORTING:S8
async function scenarioNoRefreshToken(): Promise<void> {
  let oauthHits = 0;
  await withServer(
    (path) => {
      if (path === "/api/v1/oauth/token") {
        oauthHits++;
        return { status: 200, body: "{}" };
      }
      if (path.includes("/v2/users/timestamp")) {
        return { status: 401, body: '{"error_code":-1}' };
      }
      return { status: 404, body: "" };
    },
    async (baseURL) => {
      const client = new Client({ baseURL });
      client.setTokens("only-access", "");
      try {
        await client.usersAPI.getUserTimestamp();
        failed++;
        // eslint-disable-next-line no-console
        console.log("FAIL scenario 3: expected throw");
      } catch (err) {
        const r = (err as { response?: Response }).response;
        assertEq("scenario 3: original 401 surfaced", r?.status ?? -1, 401);
      }
      assertEq("scenario 3: oauth never called", oauthHits, 0);
    },
  );
}

// Scenario 4: concurrent 401s collapse to one refresh.
// PORTING:S11
async function scenarioConcurrentRefresh(): Promise<void> {
  let oauthHits = 0;
  let timelineHits = 0;
  await withServer(
    (path, _method, body) => {
      if (path === "/api/v1/oauth/token") {
        oauthHits++;
        if (!body.includes("grant_type=refresh_token")) {
          return { status: 400, body: '{"error":"bad_grant"}' };
        }
        return {
          status: 200,
          body: JSON.stringify({
            access_token: "fresh",
            refresh_token: "fresh-r",
            user_id: 7,
          }),
        };
      }
      if (path.includes("/v2/users/timestamp")) {
        timelineHits++;
        if (timelineHits <= 3) return { status: 401, body: '{"error_code":-1}' };
        return {
          status: 200,
          body: JSON.stringify({ time: 1700000001, ip_address: "203.0.113.5" }),
        };
      }
      return { status: 404, body: "" };
    },
    async (baseURL) => {
      const client = new Client({ baseURL });
      client.setTokens("stale", "stale-r");
      const results = await Promise.all([
        client.usersAPI.getUserTimestamp(),
        client.usersAPI.getUserTimestamp(),
        client.usersAPI.getUserTimestamp(),
      ]);
      assertEq(
        "scenario 4: oauth called exactly once across 3 concurrent 401s",
        oauthHits,
        1,
      );
      assertEq(
        "scenario 4: all three retries succeed",
        results.every((r) => r.time === 1700000001),
        true,
      );
    },
  );
}

(async () => {
  await scenarioRefreshSuccess();
  await scenarioRefreshFails();
  await scenarioNoRefreshToken();
  await scenarioConcurrentRefresh();
  process.exit(failed === 0 ? 0 : 1);
})().catch((e) => {
  // eslint-disable-next-line no-console
  console.error("test crashed:", e);
  process.exit(2);
});
