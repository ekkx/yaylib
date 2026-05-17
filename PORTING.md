# Porting yaylib to Other Languages

This document is the cross-language contract for yaylib. The Go
implementation is the reference; TypeScript and Python ports SHOULD
match the contract here, deviating only where language idiom requires.

The scope of "the contract" is the public surface a typical caller
touches: client construction, the request methods promoted from the
generated layer, the auth + session flow, the upload helpers, the
event stream, and the typed errors. Internal details (transports,
connection pools, mutex placement) are deliberately out of scope —
each language solves them in its own native way.

**Runtime scope**: server-side runtimes only. The Yay! API's CORS
configuration rejects requests from external browser origins, so the
SDK is not expected to run in a web page; the TypeScript port targets
Node 18+, Bun, Deno, and Cloudflare Workers (every JS runtime with
fetch and a Node-compatible crypto module). The Python port targets
the equivalent server-side runtimes. This lets every port reach for
its host's full crypto primitive set (MD5 etc.) without polyfilling.

---

## 1. Naming and case

| Concept | Go | TS | Python |
|---|---|---|---|
| Operation method | `client.LoginWithEmail(ctx)` | `client.loginWithEmail(...)` | `await client.login_with_email(...)` |
| Type / class | `LoginEmailUserResponse` | `LoginEmailUserResponse` | `LoginEmailUserResponse` |
| Constant | `MessageTypeEternalVideo` | `MessageType.EternalVideo` (or `"eternal_video"`) | `MessageType.ETERNAL_VIDEO` |
| Field on the request body | `body.SetEmail(v)` | `body.email = v` | `body.email = v` |

Operation IDs come from the spec and are globally unique. Each
language's port keeps the operation-ID identity (camelCase / snake_case
differs but the underlying name does not).

---

## 2. Client surface contract

Every operation MUST be reachable as a top-level method on the
client, with no service-name prefix. The Go SDK gets this for free
through embed promotion; TS / Python SHOULD attach all operations to
the `Client` instance directly so callers write
`client.getRecommendedTimeline(...)` without a `client.posts.` step.

An escape hatch to the raw generated layer MUST exist for cases where
the SDK-side wrapper is in the way (e.g. cache bypass, new-field
catch-up). In Go this is `client.<Service>APIService.<Op>(ctx)...`.
Each port chooses its own form (a `client.raw` namespace, an
`@unwrapped` decorator, etc.) but the access path MUST stay
documented and stable.

---

## 3. Construction and configuration

```go
// Go
client := yaylib.NewClient(
    yaylib.WithSessionStore(store),
    yaylib.WithDeviceUUID("..."),
)
defer client.Close()
```

```ts
// TS — config object, not functional options
const client = new Client({
  sessionStore: store,
  deviceUUID: "...",
});
await client.close();
```

```python
# Python — kwargs
client = Client(session_store=store, device_uuid="...")
async with client:
    ...
```

The configuration keys MUST be the same across languages (case
notwithstanding). Wire constants (APIKey, APIVersionKey,
APIVersionName, AppVersion) are baked in and overridable; clients
SHOULD start usable without any explicit option.

`Close` (Go), `close()` (TS), context manager (Python) all map to the
same lifecycle event: cancel background work owned by the Client (lazy
X-Client-IP fetch, every open EventStream).

### 3.1 Embedded wire constants

The Yay! server validates several constants on every request. The Go
reference exposes them as `Default*` values in `client.go`; every
port MUST publish constants with the same names and the same default
values so a caller constructing a client without options gets a
Yay!-compatible client out of the box.

Two composition rules are easy to miss when reading the Go code in
isolation, so they are pinned here:

- `User-Agent` = `<DeviceType> <DeviceOS> (<DeviceDensity>x <DeviceScreen> <DeviceModel>)` — e.g. `android 11 (3.5x 1440x2960 Galaxy S9)`.
- `X-Device-Info` = `yay <AppVersion> <User-Agent>` — e.g. `yay 4.26.1 android 11 (3.5x 1440x2960 Galaxy S9)`.

Two signing-only constants travel alongside the defaults and are
referenced from §7:

- `signed_info` MD5 suffix: `yayZ1` (wire pepper).
- `signed_version` platform identifier: `yay_android` (hard-coded; deviating risks per-platform allowlist rejection).

---

## 4. Tokens

Tokens are SDK-internal state. The contract is:

- One write-side method: `SetTokens(access, refresh)` — same name in
  every language, set both in one call.
- One read-side accessor that returns a snapshot value with `Access`
  and `Refresh` fields. In Go: `client.Tokens()`. In TS:
  `client.tokens`. In Python: `client.tokens` (property).
- The snapshot MUST be a copy / value type. Mutating the returned
  object MUST NOT change the Client.

The internal mutable storage MUST be hidden — no `client.tokens.access`
write-through.

Login flows (`LoginWithEmail`, `LoadSession`, the 401 auto-refresh
path) populate the tokens automatically.

---

## 5. Sessions

`SessionStore` is an interface with three methods: `Load`, `Save`,
`Delete`. The default file-backed implementation persists sessions as
JSON keyed by email.

```go
// Go
store, _ := yaylib.NewSessionStore("sessions.json")
client := yaylib.NewClient(yaylib.WithSessionStore(store))

// Restore explicitly:
sess, err := client.LoadSession("foo@example.com")

// LoginWithEmail uses the store transparently:
client.LoginWithEmail(ctx).Email(email).Password(pw).Execute()
// → cache hit returns synthesized response with httpResp == nil
// → cache miss issues HTTP login + persists on success
```

The contract:

- `Load(email)` returns the cached session or a "not found" sentinel
  (Go: `ErrNoSession`).
- `Save(session)` is upsert; the SDK calls it after fresh login and
  after each refresh.
- `Delete(email)` is idempotent; missing key is not an error.
- A persist failure during `LoginWithEmail` returns
  `ErrSessionSaveFailed` (Go) / `SessionSaveFailedError` (TS) /
  `SessionSaveFailed` exception (Python). The login response is still
  returned and the tokens are active in memory.
- A persist failure during the 401 auto-refresh path is non-fatal —
  the rotated tokens are still active. Surface the failure via the
  language's standard logging (Go: `log/slog` warn, TS: `console.warn`,
  Python: stdlib `logging.warning`) so it stays observable.

Sync vs async:

- Go: `Load` / `Save` / `Delete` are synchronous.
- TS: all three return `Promise<...>` — file I/O is async in Node.
- Python: all three are `async def`.

### 5.1 File-backed JSON schema

The default file-backed implementation persists every session under
one top-level object keyed by email. The exact shape (which every
port MUST honour for cross-language session-file compatibility):

```json
{
  "sessions": {
    "<email>": {
      "email":         "<string>",
      "user_id":       0,
      "access_token":  "<string>",
      "refresh_token": "<string>",
      "device_uuid":   "<string>",
      "updated_at":    "<RFC3339 timestamp>"
    }
  }
}
```

`device_uuid` is mandatory in records the SDK writes — restoring a
session without restoring its matching device UUID invalidates every
`signed_info` the client would later compute (§7). `user_id` is the
authenticated user id required by the upload helpers (§8). The outer
`sessions` map allows one file to hold multiple identities.

File permissions SHOULD be `0o600` on POSIX; the file contains
access tokens. Writes SHOULD be atomic (temp-file + rename) so a
crashed write never leaves a half-truncated session file behind.

---

## 6. LoginWithEmail is the only auth wrapper

LoginWithEmail has a fluent builder + transparent session caching.
This wrapper exists because LoginWithEmail is the cache boundary —
"log in or use the cached session" is one call:

```go
resp, _, err := client.LoginWithEmail(ctx).
    Email(email).Password(password).
    Execute()
// .NoCache() forces a fresh HTTP login.
```

Other auth-flow body endpoints (`CreateUser`, `EditUser`,
`LoginWithSns`, `OauthToken` / refresh) are intentionally NOT
wrapped. They sit elsewhere in the lifecycle and adding chain
wrappers to each would multiply API surface without changing
behavior. Every port MUST follow this rule:

- Wrap LoginWithEmail (cache transparent).
- Leave other body-builder auth endpoints as plain calls into the
  generated client.
- Document the asymmetry where the LoginWithEmail wrapper is defined.

### 6.1 401 auto-refresh contract

The transport refreshes the access token automatically on a 401 from
any non-OAuth endpoint. The Go reference is in `transport.go`; the
non-obvious behaviour the port MUST match:

1. **Refresh success**: re-issue the original request once with the new Bearer, return THAT response (drop the original 401 body).
2. **Refresh failure** (no refresh token / refresh itself returned non-2xx): return the original 401 with its body intact — the caller needs the server's "your credentials are bad" message to prompt re-auth.
3. **Refresh success + retry transport error**: return the retry transport error, NOT the stale 401. Restoring the 401 here would mislead the caller into prompting for re-auth when the real problem is the network.

Other invariants:

- The OAuth endpoint itself (`/api/v1/oauth/token` and subpaths) is excluded from the auto-refresh path to prevent infinite loops.
- The OAuth endpoint uses `Authorization: Basic <base64(APIKey)>` and `Content-Type: application/x-www-form-urlencoded`; every other endpoint uses `Authorization: Bearer <access>`.
- The server rotates BOTH `access_token` and `refresh_token` on each refresh — the SDK MUST persist the new refresh token.
- Concurrent 401s MUST collapse to a single refresh call (mutex / promise / asyncio lock); the others wait and retry with the new token. Parallel refreshes invalidate each other.

> **Re-issue with the UNWRAPPED HTTP client.** Both the 401-refresh
> retry (§6.1) and the transient-failure retry (§13) re-send the
> original request. That re-send MUST go through the raw / unwrapped
> HTTP layer (the bare `fetch` / `http.Client` / `aiohttp` call), NOT
> the middleware-or-interceptor-wrapped client that owns the retry &
> refresh logic. Re-entering the wrapped client recursively re-runs
> the same retry/refresh hook, so a permanently-failing endpoint
> recurses without bound and OOMs. (The TS port hit exactly this: the
> generated `ResponseContext.fetch` is the middleware-wrapped fetch,
> not the raw one.) On re-issue, also refresh per-request headers that
> go stale — at minimum `X-Timestamp`.

---

## 7. Signing helpers

Cross-language return shape: every `Generate*` helper returns either a
single bare value (when there's only one) or a typed struct (when
multiple correlated values must be returned together).

| Helper | Go return | TS return | Python return |
|---|---|---|---|
| `GenerateSignedInfo(ctx)` | `(SignedInfo, error)` | `Promise<SignedInfo>` | `await ... -> SignedInfo` |
| `GenerateSignedInfoAt(ts)` | `SignedInfo` | `SignedInfo` | `SignedInfo` |
| `GenerateSignedVersion()` | `string` | `string` | `str` |
| `GenerateXJwt()` | `string` | `string` | `str` |
| `GenerateCallActionSignature(ctx, ...)` | `(*SignaturePayload, error)` | `Promise<SignaturePayload>` | `await ... -> SignaturePayload` |

`SignedInfo` is `{ Timestamp int64, Value string }`; both fields are
needed at the call site (Timestamp goes in the request's `timestamp`
field, Value in `signed_info`).

`SignaturePayload` is the full payload returned by the call-action
signature endpoint and accepted by the validation endpoint.

Plain string returns (`SignedVersion`, `XJwt`) are not wrapped in a
struct — they're single values whose only use is to fill one request
header / field.

### 7.1 Cross-language test vectors

Formulas are implemented in `signing.go` and `jwt.go`; what the
contract pins is **byte-identical output across languages**. Every
port MUST add a golden-value test for each of the three vectors
below. If they pass, the implementation is correct; if any fails,
the formula has drifted (encoding, key choice, message body,
trailing whitespace).

**signed_info** with `APIKey` = default, `DeviceUUID` = `00000000-0000-0000-0000-000000000000`, `timestamp` = `1700000000`:

```
=> "253b7ca3e67590204034a73326d2be87"
```

**signed_version** with default `APIVersionKey` and `APIVersionName` = `4.26`:

```
=> "zyKEqdZqaEqJjSRItHvchU9HmgWHNLtn02sWXwzQ4iQ="
```

(Standard base64 with one trailing `=`; some language stdlib
encoders append `\n` and need explicit trimming.)

**X-Jwt** with default `APIVersionKey`, `iat` = `1700000000`:

```
=> "eyJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE3MDAwMDAwMDAsImV4cCI6MTcwMDAwMDAwNX0.-UYjGTM53-Uwe8tNGjBf4ZLx644_zUEK8fpW1w16KfA"
```

(Base64 URL-safe with `=` padding stripped — NOT standard base64.
Header JSON has no `typ` and no spaces; payload JSON has key order
`iat` then `exp`, also no spaces. The TTL is 5 seconds, so callers
MUST regenerate per request — a token prepared offline and sent
later fails.)

**Call action signature flow**: the signature is computed server-side
(client never sees the secret). `GenerateCallActionSignature` posts
to `/v1/calls/action_signature/generate` and returns the full
`SignaturePayload`; `ValidateCallActionSignature` replays it against
`/v2/calls/action_signature/validate` as six query parameters
(`call_id` / `sender_uuid` / `receiver_uuid` / `action` /
`timestamp` / `signature`). Every port wraps the validation call
to spread the payload across those six parameters so callers pass
the payload through as one value.

---

## 8. Uploads

Thirteen typed `Upload<Xxx>Image[s]` methods + `UploadVideo`. The
exhaustive category table (method ↔ bucket path ↔ thumbnail size ↔
file caps) is encoded in `upload.go`; every port MUST replicate the
same set with the same names and the same values. Naming convention:
self-user methods take no userID parameter (they read `Client.UserID`
and fail with a clear "not logged in" error when it is zero);
room/group methods take the corresponding ID as the first argument.

The pieces of the upload protocol that aren't visible from category
struct definitions alone — and that every port MUST honour:

**Filename format** (the server pattern-matches these strings; format
drift breaks uploads silently):

- Main: `<category-path>/<random16>_<unix_ms>_<index>_size_<W>x<H>.<ext>`
- Thumb: `<category-path>/thumb_<random16>_<unix_ms>_<index>_size_<W>x<H>.<ext>`
- `<random16>`: 16 chars from `[0-9A-Za-z]`, freshly generated per upload call.
- `<unix_ms>`: wall-clock millisecond at call time, shared across every file in a multi-image call.
- `<W>x<H>`: source image's native resolution. Main and thumb carry the SAME `WxH` token (the server enforces consistency).
- `<Y/M/D>` in date-bucketed categories: decimal, NO zero-padding — `2026/4/26`, not `2026/04/26`.
- Thumb extension MUST equal main extension.

**Server-canonical filename** (the most common porting bug):

The presigned-URL response returns a `filename` field with an `s3/`
prefix prepended (e.g. `s3/user/123/avatar/...`). The Upload method's
return value MUST be that canonical name, NOT the locally-generated
path the SDK PUT to. Callers pass it to
`EditUser.profile_icon_filename` / `CreatePost.attachment_filename` /
etc. Sending the locally-generated name (without `s3/`) makes the
server treat the upload as missing and silently drop the main image.

**Animated-GIF thumbnail rule**:

Animated GIFs MUST stay animated through the thumbnail step (decode
all frames → per-frame composite to full canvas → bilinear resize →
Floyd–Steinberg dither → re-encode with `DisposalNone`). Skipping
the thumb, encoding it as a static frame, or transcoding to JPEG
triggers server-side moderation and DELETES the main image too.

**Video uploads**:

- Endpoint: `GET /v1/users/presigned_url?video_file_name=<flat-name>` returns one `presigned_url` (no thumbnail companion).
- Flat filename: `<random16>_<unix_ms>.mp4` (no category path, no `_size_WxH` suffix).
- The server does NOT return a canonical filename for videos; `UploadVideo` returns the same flat name the SDK generated. Callers pass it to `CreatePost.video_file_name` / `SendChatMessage.video_file_name`.
- Content-Type: `video/mp4`.

**Content-Type mapping** (matching extension keeps S3 intermediaries
happy even though the presigned URL signs only the host):
`.jpg`/`.jpeg`/unknown → `image/jpeg`, `.png` → `image/png`,
`.gif` → `image/gif`, `.mp4` → `video/mp4`.

Image upload input shape: `{ filename, body }` per file. `filename`
is used to detect the extension and Content-Type. For multi-image
methods (`UploadPostImages`, `UploadChatMessageImages`,
`UploadReportImages`), arrays are passed; the per-method file cap is
9 except `UploadReportImages` which is 4.

Video upload input shape: a single byte stream (Go: `io.Reader`,
TS: `Blob | ReadableStream | Uint8Array`, Python: `AsyncIterator[bytes]
| bytes | BinaryIO`). The video upload endpoint does NOT return a
canonical filename, so the random filename the SDK generates is what
callers pass to CreatePost / SendChatMessage as `video_file_name`.

Note: the `filename` field common to image uploads has no role for
video. This asymmetry is intentional — the server's filename naming
rules differ between images (`<category-path>/<random>_<ts>_<index>_size_<W>x<H>.<ext>`)
and videos (flat `<random>_<ts>.mp4`), and exposing a `filename` knob
on UploadVideo would invite callers to fight that.

Streaming: image uploads always buffer (the thumbnail step needs the
full bytes anyway). Video uploads stream when the body implements the
language's seekable-stream interface (`io.ReadSeeker` in Go, fixed-size
`Blob` in TS, file objects with `tell()`/`seek()` in Python); otherwise
they fall back to a full in-memory buffer because the presigned PUT
needs a known Content-Length.

Return value: each Upload method returns the server-canonical filename
(with the `s3/` prefix) for the main image / video. Thumbnails live at
sibling paths the server resolves automatically. The canonical filename
is what callers pass back to `EditUser.profile_icon_filename`,
`CreatePost.attachment_filename`, etc.

---

## 9. Errors

Three layers:

1. **Wire layer** (`APIError` in Go = `gen.GenericOpenAPIError`):
   carries the raw HTTP body and status. Recover via
   `errors.As(err, &apiErr)` (Go), `instanceof APIError` (TS),
   `except APIError` (Python).

2. **Parsed payload** (`ErrorResponse`):
   the Yay! server's JSON body — `error_code`, `message`, optional
   `ban_until`, `retry_in`, `remaining_quota`. Parse with
   `yaylib.ErrorResponseOf(err)` (Go) /
   `yaylib.errorResponseOf(err)` (TS) /
   `yaylib.error_response_of(err)` (Python).

3. **Typed code** (`ErrorCode`): use this for `switch` / `match`
   dispatch. `yaylib.CodeOf(err)` (Go) /
   `yaylib.codeOf(err)` (TS) /
   `yaylib.code_of(err)` (Python).
   The value space is generated from the Yay! server's error
   catalog; the constants (`ErrCodeRequired2FA`,
   `ErrCodeUserTemporaryBanned`, ...) are the same names across
   languages.

The contract:

- `CodeOf(err)` MUST return `ErrCodeUnknown` (or the language's
  equivalent zero-ish code) when err is not an APIError or carries no
  recognizable code. Do NOT panic / throw.
- `ErrorResponseOf(err)` MUST return null / None / nil for
  unrecognized errors.
- An APIError stays usable for raw-body access (`apiErr.Body()` in
  Go) — do not collapse into `ErrorResponse` only.

---

## 10. Event stream

The event stream is a multiplexed real-time channel — one underlying
connection (a WebSocket against `wss://cable.yay.space`), N subscriptions
on top.

Common surface (same shape, language-idiomatic API):

| Concept | Go | TS | Python |
|---|---|---|---|
| Open the stream | `client.OpenEventStream(ctx, opts...)` | `await client.openEventStream(opts)` | `async with client.open_event_stream(opts) as stream` |
| Subscribe | `stream.Subscribe(ctx, ChatRoomChannel())` | `await stream.subscribe(ChatRoomChannel())` | `await stream.subscribe(ChatRoomChannel())` |
| Receive events | `for ev := range sub.Events() { ... }` | `sub.onNewMessage(ev => ...)` | `@sub.on_new_message` decorator |
| Unsubscribe | `sub.Unsubscribe(ctx)` | `await sub.unsubscribe()` | `await sub.unsubscribe()` |
| Close stream | `stream.Close()` | `await stream.close()` | exit `async with` |

The dispatch idiom differs by language but the underlying event
inventory is fixed:

| Wire `event` | SDK type name |
|---|---|
| `new_message` | `NewMessageEvent` |
| `video_processed` | `VideoProcessedEvent` |
| `chat_deleted` | `ChatDeletedEvent` |
| `total_chat_request` | `TotalChatRequestEvent` |
| `unsubscribed` | `UnsubscribedEvent` |
| `new_post` | `GroupUpdatedEvent` (note: the wire name is `new_post` but the payload is a group-id update) |
| `conference_call_finished` | `CallFinishedEvent` |

Channels:

- `ChatRoomChannel()` — global chat-room signals
- `MessagesChannel(roomID)` — messages + video_processed in one room
- `GroupUpdatesChannel()` — group state updates across all groups
- `GroupPostsChannel(groupID)` — new posts in one group's timeline

Reconnect policy is configurable (`MaxAttempts`, `InitialDelay`,
`MaxDelay`, `Disabled`). The SDK MUST treat a connection that stays
healthy past a fixed threshold (currently 30s) as evidence of
stability and reset the failure budget — a flap pattern (welcome
frame + immediate drop) MUST count toward MaxAttempts to avoid an
endless reconnect loop.

`OnDrop` callback (Go: `EventStreamOptions.OnDrop`, TS: option /
event listener, Python: callback kwarg) fires when a Subscription's
buffer is full and an event has to be dropped — this MUST stay
observable.

Authentication: every Open call fetches a short-lived token via
`GetWebSocketToken`. An unauthenticated Client (no access token) is
returned the underlying 401 as the Open error.

### 10.1 Wire constants

The handshake / subscribe / event-dispatch flow itself is in
`event_stream.go` (Rails-ActionCable-style JSON). What every port
MUST replicate exactly — these are pattern-matched server-side and
silent failures result from drift:

**Dial URL**: `wss://cable.yay.space?token=<jwt>&app_version=<APIVersionName>`.
The dial carries NO `Authorization` header (authentication is
exclusively the `token` query parameter). The token comes from
`GET /v1/users/ws_token` and has a 30-second TTL at handshake time
only (established connections continue past `exp`).

**Channel identifiers** (the server matches these strings literally
when routing later frames — the space after `,` is significant, and
the IDs are JSON STRINGS not numbers):

| Channel | Identifier (verbatim) |
|---|---|
| `ChatRoomChannel` | `{"channel":"ChatRoomChannel"}` |
| `MessagesChannel(roomID)` | `{"channel":"MessagesChannel", "chat_room_id":"<rid>"}` |
| `GroupUpdatesChannel` | `{"channel":"GroupUpdatesChannel"}` |
| `GroupPostsChannel(groupID)` | `{"channel":"GroupPostsChannel", "group_id":"<gid>"}` |

**Frame inventory**: client sends `{"command":"subscribe","identifier":...}`
and `{"command":"unsubscribe","identifier":...}`; server sends
`{"type":"welcome"}` once after dial, `{"type":"ping"}` periodically
(silently ignored — no pong), `{"type":"confirm_subscription","identifier":...}`,
`{"type":"disconnect", ...}`, and events as
`{"identifier":..., "message":{"event":..., "data":...}}` (some
events use `"message"` instead of `"data"` — accept both).

**Event-to-DTO inventory** — the wire `event` name does NOT always
match the SDK-side type name, and porters MUST preserve the SDK
names verbatim so user code is portable:

| Wire `event` | SDK type |
|---|---|
| `new_message` | `NewMessageEvent` |
| `video_processed` | `VideoProcessedEvent` |
| `chat_deleted` | `ChatDeletedEvent` |
| `total_chat_request` | `TotalChatRequestEvent` |
| `unsubscribed` | `UnsubscribedEvent` |
| `new_post` | `GroupUpdatedEvent` (payload is `{group_id}`, NOT a post) |
| `conference_call_finished` | `CallFinishedEvent` |

`video_thumbnail_big_url` is a WebSocket-only field on `new_message`
— it does NOT appear on the REST `Message` DTO. Treat it as nullable.

---

## 11. Resilience: server-as-source-of-truth

The Yay! server emits responses with shapes the spec doesn't always
predict — fields typed as String on the source side that arrive as
JSON numbers, enum values that aren't in the extracted vocabulary,
etc. The reference client absorbs these via lenient JSON
deserialization; strict decoders would otherwise reject every such
response and leave callers stuck.

Every port MUST follow these four rules so the SDK keeps working
when the wire format drifts:

1. **Non-path parameters are optional.** Path parameters stay
   required (the URL can't be assembled without them). Every other
   parameter — query, header, body field — MUST be emitted as
   optional regardless of how the spec or source side declares it.
   The server is the source of truth for which arguments are
   actually mandatory; SDK-side rejection just masks calls the
   server would have accepted.

2. **Enum unmarshaling is tolerant.** The generated typed
   constants are for IDE help and `IsValid()`-style introspection
   only. Receiving an unknown string MUST set the field to that
   string and continue, NOT raise an error. New server values then
   surface as data the caller can inspect, instead of breaking
   every call to that endpoint.

3. **Per-call raw escape hatch (`ExecuteRaw` / `executeRaw` /
   `execute_raw`).** Every request builder MUST expose a sibling
   that returns the raw response bytes alongside the HTTP response,
   bypassing the typed JSON decode entirely:

   ```go
   // Go
   body, httpResp, err := client.GetTimeline(ctx, "").ExecuteRaw()

   // TS
   const { body, httpResponse } = await client.getTimeline("").executeRaw();

   // Python
   body, http_response = await client.get_timeline("").execute_raw()
   ```

   Return contract: HTTP success returns `(body, response, nil)`;
   HTTP error (4xx/5xx) returns `(body, response, APIError)` so
   `CodeOf(err)` still works; transport failure returns
   `(nil, nil, err)`. Typed JSON-decode failures MUST NOT propagate
   through this path — that's the whole point of the escape hatch.

4. **Hand-curated overrides.** When the SDK ships a typed shape
   that's wrong, callers shouldn't have to live on `ExecuteRaw`
   forever. The build pipeline reads two YAML overlays applied to
   the spec before code generation:

   - **Manual enum overlay** — declares wire-only enums that have
     no source-side enum class to extract, plus per-case names
     that drive the generated constant identifiers.
   - **Manual DTO override overlay** — replaces a DTO field's
     extracted type with the wire reality. A `dto:` plus `field:`
     pair targets exactly one DTO; `field:` alone is a global
     rule applying to every DTO that declares that field name,
     and a per-DTO entry wins when both exist.

   Each port's build pipeline MUST consume the same overlays so
   the same wire patches reach all three languages from a single
   source.

The recommended workflow when the wire format diverges in
production: hit the mismatch, capture the body via `ExecuteRaw` (or
the `APIError.Body()` it left in the typed-path error), pick the
right overlay to fix it, regenerate, ship.

---

## 12. Required headers

The transport adds these on every API call:

- `User-Agent` (composed, see §3.1)
- `X-App-Version` (= `APIVersionName`)
- `X-Device-Info` (composed, see §3.1)
- `X-Device-UUID`
- `X-Client-IP` (fetched lazily from `GET /v2/users/timestamp.ip_address` on the first request, cached)
- `X-Connection-Type`
- `X-Connection-Speed` (literal `0 kbps`, with the space)
- `Accept-Language`
- `X-Timestamp` (fresh every request — MUST NOT be cached)
- `Authorization: Bearer <access>` for normal endpoints
- `Authorization: Basic <base64(APIKey)>` for `/api/v1/oauth/token*` only; base64 is standard (with padding), NOT URL-safe
- `Content-Type: application/json;charset=UTF-8` when JSON. The casing and the lack of space between `;` and `charset` are wire-significant; some intermediaries normalize this and the server has been observed to reject the normalized form.

Caller-provided header values win (set-if-absent) EXCEPT
`X-Timestamp` and `Content-Type` — those two are always normalized
by the transport so a test override of `X-Device-Info` / `User-Agent`
cannot accidentally pin a stale timestamp.

The Yay! server validates these. Missing or stale values produce
opaque 4xx responses that look like auth bugs. Every port MUST inject
the same set.

`X-Jwt` is required for some write endpoints (most visibly
`CreatePost`); generate it per-call via `GenerateXJwt`. The TTL is 5
seconds — a request prepared offline and sent later will fail.

Presigned PUT requests (image / video upload) MUST NOT carry the
`Authorization: Bearer` header — S3 signs via query parameters and a
spurious header confuses some intermediaries. Same rule for the
WebSocket dial: token goes in the query string, no `Authorization`
header.

### 12.1 Host routing

Most operations go to the primary API host, but a few are served from
a separate auxiliary host. A request for one of those sent to the
primary host comes back as a router-level 404 (`{"status":404,"error":
"Not Found"}`) — it is NOT removed, just elsewhere. The transport MUST
rewrite the request origin for those operations.

Invariants every port keeps:

- The operation→host map is **generated, not hand-written** — it ships
  as a data table in the generated layer (Go `genHostRoutes`, TS
  `gen/hostRoutes`, Python `_host_routes.HOST_ROUTES`), keyed by the
  upper-cased `"METHOD path"`. Never hard-code the endpoint list in the
  transport; consume the table.
- Host names are **symbolic** (`cassandra`), resolved to a base URL the
  client config exposes and lets the caller override
  (`WithCassandraBaseURL` / `cassandraBaseURL` / `cassandra_base_url`,
  default `https://cas.yay.space`). The override is what makes the
  routing testable offline and against the mock server.
- The rewrite runs **first**, before header injection / auth-refresh /
  retry, so every later stage observes the final URL (and the post-401
  replay re-sends to the same host).
- Matching is exact `"METHOD path"`. The currently-routed operations
  are all parameter-free, so the request path equals the OpenAPI path
  template; a port does NOT need a path-template matcher until a
  parameterized path joins the table (revisit then).
- An unrecognized symbolic host in the table resolves to "primary host"
  rather than an error — a spec change can't strand a request.

This is the same source-of-truth discipline as the rest of the
pipeline: the host assignment lives in the spec overlay and propagates
to every language through the generated table; a port that special-
cases it in transport code will drift. See §15 for the parity scenario.

### 12.2 Logging

Logging is **optional, injected, and silent by default**. The SDK is a
library: with no logger configured it MUST produce zero output (no
stderr, no stdout, no native-logger side effects). A logger is supplied
once at construction (`WithLogger` / `logger=` / the `logger` option)
and from then on the SDK emits structured records to it.

Logger shape (idiomatic per language, same four levels):

- Go: a `*slog.Logger`. Default is a logger over a discard handler.
- TypeScript: a `Logger` interface — `debug/info/warn/error(message,
  fields?)`. Default is a no-op object.
- Python: a stdlib `logging.Logger` (or anything duck-compatible).
  Default is a module logger with a `NullHandler` and propagation off.

Levels: `Debug`, `Info`, `Warn`, `Error`. `Info` is reserved (no call
sites yet). Construction MUST NOT do work when the level is disabled —
in particular the no-op default path must be allocation-cheap, so
`Debug` field assembly is only reached when a logger is present.

Every record carries a stable **`event`** key whose value is a
snake_case identifier. **The `event` value is the cross-language
contract** — the human-readable message text MAY differ slightly per
language, but the `event` value and the documented field keys MUST be
identical across ports. Adding or renaming an `event` is a contract
change: update this list in the same change.

Canonical events:

| level | `event` | fields | when |
|---|---|---|---|
| WARN  | `token_persist_fail` | `user_id` | refresh succeeded but persisting the rotated tokens to the session store failed (non-fatal; §6.1) |
| WARN  | `session_load_fail` | `err` | session store load raised on the login path; the SDK falls back to a network login (§5) |
| DEBUG | `http_request` | `method`, `path`, `host` | a request is about to be sent (after host-routing / header injection) |
| DEBUG | `http_retry` | `method`, `path`, `attempt`, `status`/`error`, `delay_ms` | a transient failure is being retried (§13) |
| DEBUG | `token_refresh` | `outcome` (`ok`/`failed`/`no_token`) | the 401 auto-refresh chain ran (§6.1) |
| DEBUG | `host_route_rewrite` | `op`, `host` | a request origin was rewritten to an auxiliary host (§12.1) |
| DEBUG | `ws_reconnect` | `attempt`, `delay_ms` | the event stream is reconnecting (§10) |
| ERROR | — | — | reserved; no call sites (every fatal currently returns to the caller) |

Redaction is a **hard contract, not a guideline**. The SDK MUST NEVER
pass to a logger: access or refresh tokens, password, API key,
`signed_info` / `signed_version`, the `X-Jwt` value, the
`Authorization` header, or any request/response body. `http_request`
logs method + path + host only — never the query string, headers, or
body. `email` and `user_id` MAY appear (they are useful and the caller
already holds them). A port that logs a banned value is a defect even
if the value is "only at debug level".

This mirrors the rest of the SDK's discipline: the contract lives here,
each language's `*_test.go` / `*.test.ts` / `test_*.py` translates the
same parity scenario (§15), and a port that diverges on event names or
redaction drifts silently. See §15.

---

## 13. Retry policy

`RetryPolicy` is a struct with `MaxAttempts`, `BaseDelay`, `MaxDelay`,
`RetryOnPOST`. Defaults: 3 attempts, 200ms base, 30s ceiling, no POST
retry. The zero value disables retries.

Method-by-method behavior:

- `429` retries on every method — the server explicitly asked.
- `5xx` and network errors retry only for safe methods (GET / HEAD /
  PUT / DELETE / OPTIONS) by default. POST / PATCH retry only when
  `RetryOnPOST=true` because the server has no idempotency-key
  concept.
- The body's `retry_in` field, when present, is honored over the
  computed backoff.
- Backoff is exponential with full jitter, clamped to `MaxDelay`.

`ReconnectPolicy` (the event-stream sibling) has the OPPOSITE zero
semantics: `MaxAttempts == 0` means "unlimited", because an unattended
event stream wants to keep recovering. Document the difference loudly.

---

## 14. UUID and identity

`NewUUIDv4` (Go) / `newUUIDv4()` (TS) / `new_uuid_v4()` (Python) all
produce a canonical 8-4-4-4-12 hex UUID v4 from the language's
crypto-grade randomness primitive. Failure of that primitive MUST
panic / throw — do NOT degrade to a weak fallback (collisions
guaranteed).

Device UUID is stable per process unless a session restore replaces
it. Persisted sessions carry the active DeviceUUID so a client that
loads a session keeps its device identity.

---

## 15. Test parity

Every port SHOULD include the following behavior tests (the Go
versions in `*_test.go` are reference fixtures — translate the
scenarios, not the Go API mechanics):

- Login fresh → tokens active + persisted to store.
- Login cached → no HTTP, tokens applied from store.
- Login store load error → return error without HTTP (rate-limit
  protection).
- Login save error → wrapped as `ErrSessionSaveFailed`, tokens still
  active.
- 2FA-required server response surfaces as `ErrCodeRequired2FA`.
- Required headers injected on every request; OAuth endpoints get
  Basic; others get Bearer.
- 401 → refresh → retry chain.
- 401 refresh failure surfaces the original 401 body.
- Refresh succeeds + retry network error surfaces the retry error,
  not a stale 401.
- Lazy X-Client-IP fetch populates the header on the second request.
- Retry honors `retry_in` from the body.
- Retry respects context cancellation.
- POST not retried by default; POST 429 retried.
- Upload happy paths for the major categories (avatar / post /
  chat-message / report) including correct filename shape and
  thumbnail upload.
- Upload presigned PUT carries no Bearer.
- Upload thumbnail extension matches main extension (incl. animated
  GIF preservation).
- Event stream subscribe + receive on each channel type.
- Event stream reconnect after server close, re-subscribes all subs.
- Multiple-sub reconnect: after reconnect both subs are resubscribed
  and both still receive events.
- Event stream subscribe rejection / timeout.
- WebSocket dial carries no Bearer.
- OnDrop fires when the per-sub buffer overflows.
- Stable connection (>= 30s) resets the failure budget.
- ExecuteRaw on a typed-decode-failing endpoint returns the raw body
  with no error (typed JSON-decode errors absorbed; HTTP errors
  still surface as APIError with body intact).
- A response whose JSON shape doesn't match the typed model: typed
  Execute returns the parse error with the raw body in
  APIError.Body(); ExecuteRaw on the same call returns
  (body, response, nil).
- Unknown enum value in a server response is accepted (typed field
  set to the unknown string, no error). IsValid() reports false on
  it; the typed constants stay usable.
- Host routing (§12.1): a host-routed operation goes to the
  configured auxiliary base URL; a non-routed operation stays on the
  primary base URL. Two stand-in servers, assert which one each call
  reached.
- Logging (§12.2): with no logger, an operation that hits the
  `token_persist_fail` path produces zero output. With a capturing
  logger injected, the same path emits one WARN record carrying
  `event=token_persist_fail`; assert no captured record (any level,
  any field) contains the access token, refresh token, password, or
  API key.

---

## 16. Versioning

`major` is locked across the three languages — a breaking change in
any language bumps all three. `minor` and `patch` drift freely. The
Yay! server's API version (`APIVersionName`) is exposed as a constant
and is NOT tied to the SDK's version number.

Module path note for Go: `github.com/ekkx/yaylib/v2` for v2.x;
v3 migration mechanically rewrites `/v2` → `/v3` in import paths.

---

## 17. What's deliberately NOT in this document

- Internal transport details (RoundTripper, connection pool, mutex
  placement). Pick the language-native equivalent.
- The OpenAPI spec and the generated DTOs. Each language bundles its
  own generated client; the SDK wrapper layer described here sits
  ON TOP of that.
- Code style. Follow each language's idiom — gofmt, prettier, ruff.
