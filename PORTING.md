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

---

## 8. Uploads

Thirteen typed `Upload<Xxx>Image[s]` methods + `UploadVideo`. Each
maps to a specific bucket-side category path and thumbnail policy
(see the Go source for the exhaustive table; the table is part of
the contract — every port MUST replicate the categories with the
same paths and same thumbnail dimensions).

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

- `User-Agent`
- `X-App-Version`
- `X-Device-Info`
- `X-Device-UUID`
- `X-Client-IP` — fetched lazily after the first request, cached
- `X-Connection-Type`
- `X-Connection-Speed`
- `Accept-Language`
- `X-Timestamp` — fresh every request
- `Authorization: Bearer <access>` — for normal endpoints
- `Authorization: Basic base64(api_key)` — for `/api/v1/oauth/token*`
- `Content-Type: application/json;charset=UTF-8` — when JSON

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
