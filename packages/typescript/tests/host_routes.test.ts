// Host-routing parity (mirrors the Go host_routes_test.go). A few
// operations are served from an auxiliary host; the transport must
// send those — and only those — there, leaving every other request on
// the primary host. Two stand-in servers pin the routing offline.

import { createServer, type Server } from "node:http";
import type { AddressInfo } from "node:net";

import { Client } from "../src/client";

async function startServer(hits: string[]): Promise<{ url: string; close: () => Promise<void> }> {
  const server: Server = createServer((req, res) => {
    hits.push(`${req.method ?? "GET"} ${(req.url ?? "").split("?")[0]}`);
    let body = "";
    req.on("data", () => {});
    req.on("end", () => {
      res.writeHead(200, { "Content-Type": "application/json" });
      res.end("{}");
    });
    void body;
  });
  await new Promise<void>((resolve) => server.listen(0, resolve));
  const addr = server.address() as AddressInfo;
  return {
    url: `http://127.0.0.1:${addr.port}`,
    close: async () => {
      server.closeAllConnections?.();
      await new Promise<void>((resolve) => server.close(() => resolve()));
    },
  };
}

let failed = 0;
function assert(name: string, cond: boolean, detail = ""): void {
  if (cond) console.log(`PASS ${name}`);
  else {
    failed++;
    console.log(`FAIL ${name}${detail ? `\n  ${detail}` : ""}`);
  }
}

async function main(): Promise<void> {
  const primaryHits: string[] = [];
  const auxHits: string[] = [];
  const primary = await startServer(primaryHits);
  const aux = await startServer(auxHits);

  try {
    const client = new Client({
      baseURL: primary.url,
      cassandraBaseURL: aux.url,
    });
    client.setTokens("ACC", "REF");

    // Host-routed: GET /api/v2/user_activities must hit the aux host.
    await client.activitiesAPI.getUserActivities({});
    assert(
      "activity feed routed to auxiliary host",
      auxHits.includes("GET /api/v2/user_activities"),
      `aux=${JSON.stringify(auxHits)}`,
    );
    assert(
      "primary host did not receive the host-routed op",
      !primaryHits.includes("GET /api/v2/user_activities"),
      `primary=${JSON.stringify(primaryHits)}`,
    );

    // Unrouted: GET /v2/users/timestamp must stay on the primary host.
    await client.usersAPI.getUserTimestamp();
    assert(
      "unrouted request stayed on primary host",
      primaryHits.some((h) => h.endsWith("/v2/users/timestamp")),
      `primary=${JSON.stringify(primaryHits)}`,
    );
    assert(
      "auxiliary host received no unrouted request",
      !auxHits.some((h) => h.endsWith("/v2/users/timestamp")),
      `aux=${JSON.stringify(auxHits)}`,
    );
  } finally {
    await primary.close();
    await aux.close();
  }

  process.exit(failed === 0 ? 0 : 1);
}

void main();
