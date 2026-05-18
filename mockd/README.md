# mockd

A behavior-contract HTTP server for the yaylib SDK test and
development suites. It answers every route the API exposes and, when a
request carries an `X-Yay-Mock-Scenario` header, produces a
deterministic failure / latency behavior so the Go, TypeScript and
Python suites can assert retry / refresh / rate-limit handling against
one shared server instead of three divergent in-process fakes.

This is a development and test tool. It is **not** part of any
published SDK and ships separately.

## Install

```bash
go install github.com/ekkx/yaylib/mockd@latest
```

## Run

```bash
mockd -addr :8080            # default :8080

# Ephemeral port + address discovery (handy for a parallel test harness):
mockd -addr 127.0.0.1:0 -ready-file /tmp/addr
# the resolved host:port is written to the ready-file once listening
```

Or from a checkout:

```bash
cd mockd
go run . -addr :8080
go test ./...                     # unit + integration tests
```

| Flag | Meaning |
|---|---|
| `-addr` | Listen address. Use `host:0` for an ephemeral port. Default `:8080`. |
| `-ready-file` | If set, the resolved `host:port` is written here once the server is listening (lets a harness use `:0` without a port race). |

## The scenario contract

> This grammar is a **cross-language contract**. The Go / TS / Python
> test suites emit these exact header strings. Do not change an
> accepted form without updating all suites.

Two request headers drive behavior:

| Header | Meaning |
|---|---|
| `X-Yay-Mock-Scenario` | Selects the behavior (see grammar below). Absent ⇒ happy path. |
| `X-Yay-Mock-Session` | Opaque per-test key (mint a UUID per test case). Stateful scenarios count requests **per `(session, scenario)`**, so parallel runs across all languages never share counters. |

### Grammar

| Scenario | Behavior |
|---|---|
| `fail-<status>-times-<k>` | The first **k** requests answer `<status>` with a `CommonErrorResponse` body, then subsequent requests fall through to the happy path. `<status>` 100–599, `k` ≥ 1. A `401` carries `error_code` = `AccessTokenExpired` (-3). |
| `delay-ms-<n>` | Sleep `n` ms, then fall through to the happy path. |
| `retry-after-<n>` | The **first** request answers `429` with `Retry-After: <n>` and body `retry_in=<n>`; subsequent requests fall through. |
| `expired-token` | Protected requests answer `401` (`error_code` -3) until the token endpoint is hit (a simulated refresh, which latches the session and returns a fresh `TokenResponse`); after that, requests fall through. |
| `ratelimit-quota-zero` | Every request answers `429` with `error_code` = `QuotaLimitExceeded` (-343) and `remaining_quota=0`. No fall-through. |
| `body-malformed` | `200` with a JSON **scalar** that cannot decode into any typed model. Exercises the typed-parse-fail / `ExecuteRaw` contract (typed `Execute` errors with the raw body in `APIError.Body()`; `ExecuteRaw` returns `(body, response, nil)`). |
| `body-unknown-enum` | `200` with `{"post_type":"__mock_unknown_enum__"}` — an unknown enum value the tolerant decoder must accept (field set to the unknown string, no error, `IsValid()` false). Point this at an endpoint whose typed model carries the `post_type` enum. |
| `ok-status-<2xx>` | A success answered with a non-200 2xx (e.g. `ok-status-201`, mirroring login's `201`) plus an empty-object body. The SDK must treat any `2xx` as success and decode it into the operation's all-optional success model, not key on the exact documented status. |

An unrecognized scenario value returns `400` with a
`CommonErrorResponse` — tests fail loud rather than silently get a
happy path.

### Composition note

The `(session, scenario)` counter is shared across a request chain, so
scenarios compose with the SDK's own 401→refresh→retry logic:

- `fail-401-times-1` ⇒ first protected request `401`, the SDK's
  refresh call succeeds (counter spent), the retried request
  succeeds. Exercises the refresh-then-retry path.
- `fail-401-times-2` ⇒ first protected request `401`, the refresh
  call also `401` (counter still within budget). Exercises the
  refresh-failure path.

### Diagnostic headers

Every HTTP response the mock owns carries observational headers so a
cross-language parity failure can be diagnosed from the response
alone (SDKs ignore them):

| Header | Meaning |
|---|---|
| `X-Mock-Scenario` | The scenario header the mock saw (omitted when none). |
| `X-Mock-Attempt` | The `(session, scenario)` attempt ordinal — present for the counter-based scenarios (`fail-…-times-…`, `retry-after-…`). |
| `X-Mock-Branch` | The decision branch taken: `happy`, `fail`, `fail-exhausted-happy`, `delay-happy`, `retry-after-429`, `retry-after-happy`, `expired-token-401`, `expired-token-refresh`, `expired-token-happy`, `ratelimit-429`, `body-malformed`, `body-unknown-enum`, `ok-status`, `bad-scenario`, `s3-put`, `s3-put-rejected`, `admin-*`. |

(The WS upgrade is not annotated — these cover the HTTP
request/response path.)

### Happy path

- `POST /api/v1/oauth/token` always returns a deterministic
  `TokenResponse` (login and refresh issue distinct access tokens),
  so auth/login parity works with no scenario header.
- Every other known route returns `200 {}`. Bodies are intentionally
  empty: behavioral fidelity is the scenario engine's job, and each
  SDK decodes its own typed models. Unknown routes return `404`
  (path/method routing is real).

### Admin routes

| Route | Effect |
|---|---|
| `GET /__mock/health` | `200 {"status":"ok"}` |
| `POST /__mock/reset` | Clears all scenario state (`204`). Lets a suite isolate cases without minting new session ids. |
| `POST /__mock/ws/close` | Closes every live event-stream connection (`204`). Drives reconnect / multi-subscription parity. |

## Upload contract

The presigned-URL endpoints are owned so upload parity has real
bodies, and the presigned URLs point back at this server so the PUT
stays in-process:

| Endpoint | Behavior |
|---|---|
| `GET /v1/buckets/presigned_urls?file_names[]=…` | One `{filename, url}` per name. `filename` is the **server-canonical** value (`s3/` prefix prepended) — the SDK threads this into the create/edit call. `url` targets the in-process S3 receiver. |
| `GET /v1/users/presigned_url?video_file_name=…` | Single `{presigned_url}` for video (flat name, no canonical prefix). |
| `PUT /__mock/s3/<name>` | The S3 receiver. `200` on accept; **`403` if the request carries `Authorization`** — a presigned PUT must not carry a bearer, enforced server-side so all languages get the same check. |

These GETs are normal requests, so scenario headers still apply
(e.g. `fail-503-times-1` on the presigned GET).

## Event-stream contract

The server speaks the ActionCable-style JSON protocol the SDK
expects. A WS dial is detected by the upgrade headers on any path
(the event-stream base URL is configurable), authenticates via the
`token` query, and **must not carry `Authorization`** (the upgrade is
refused if it does).

Frames:

```
server -> {"type":"welcome","sid":"mock-sid"}
client -> {"command":"subscribe","identifier":"<json>"}
server -> {"identifier":"<json>","type":"confirm_subscription"}
server -> {"identifier":"<json>","message":{"event":"<n>","data":{…}}}
client -> {"command":"unsubscribe","identifier":"<json>"}
```

After a successful subscribe the server pushes one representative
event chosen by the channel named in the identifier:
`ChatRoomChannel`→`chat_deleted`, `MessagesChannel`→`new_message`,
`GroupUpdatesChannel`/`GroupPostsChannel`→`new_post`.

Per-dial behavior is selected with the `mock` query parameter (the
portable control channel — WS dials carry no custom headers):

| `mock=` | Behavior |
|---|---|
| _(absent)_ | Confirm every subscribe, then push the representative event. |
| `reject` | Reject every subscribe (`reject_subscription`). |
| `no-confirm` | Ignore subscribes — exercises the client's confirm timeout. |
| `drop-after-confirm` | Confirm + push, then close. The SDK reconnects to the same URL and the next connection confirms + pushes again, so reconnect / re-subscribe is observable without out-of-band orchestration. |

## Layout

```
mockd/
├── main.go                 server entrypoint (flags, graceful shutdown)
├── internal/scenario/      header grammar parser + session state + tests
├── internal/mock/          scenario middleware, error bodies, router wiring + tests
└── internal/oas/           typed router (machine-generated; do not hand-edit)
```
