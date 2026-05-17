// PORTING.md §2 / §3: Client is the single entry point. Every operation
// MUST be reachable as a top-level method on the client; the SDK uses
// composition over the generated tag-services (AuthApi / PostsApi / ...)
// to expose all operations without a `client.posts.` step.
//
// The full operation surface lives across 23 generated *Api classes. Each
// is instantiated against the same Configuration so they share the
// middleware chain (headers / auth / retry). A generated promotion layer
// will eventually emit the flat surface; for the spike we expose the
// per-service instances on the client so callers can reach operations as
// `client.usersAPI.getUser(...)` while we wire up the wrapper layer.

import {
  ActivitiesApi,
  AuthApi,
  AppsApi,
  BucketsApi,
  CallsApi,
  ChatRoomsApi,
  ConversationsApi,
  EmailVerificationUrlsApi,
  GamesApi,
  GenresApi,
  GiftsApi,
  GroupMuteApi,
  GroupsApi,
  HiddenApi,
  ModerationApi,
  NotificationSettingsApi,
  PinnedApi,
  PostsApi,
  ReceivedGiftsApi,
  StickerPacksApi,
  SurveysApi,
  ThreadsApi,
  UsersApi,
} from "./gen/apis";
import { Configuration, ResponseError, type ApiResponse } from "./gen/runtime";

import { APIError, asAPIError } from "./errors";
import {
  type Upload,
  type UploadCategory,
  type UploadDeps,
  type VideoBody,
  chatBackgroundUpload,
  chatMessageUpload,
  groupCoverUpload,
  groupCreationCoverUpload,
  groupCreationIconUpload,
  groupIconUpload,
  reportUpload,
  signupAvatarUpload,
  threadIconUpload,
  uploadImages,
  uploadSingleImage,
  uploadVideo,
  userAvatarUpload,
  userCoverUpload,
  userPostUpload,
  videoCallSnapshotUpload,
} from "./upload";
import { EventStream, type EventStreamOptions, openEventStream } from "./event_stream";

import {
  DEFAULT_ACCEPT_LANGUAGE,
  DEFAULT_API_KEY,
  DEFAULT_API_VERSION_KEY,
  DEFAULT_API_VERSION_NAME,
  DEFAULT_APP_VERSION,
  DEFAULT_BASE_URL,
  DEFAULT_CASSANDRA_BASE_URL,
  DEFAULT_CONNECTION_SPEED,
  DEFAULT_CONNECTION_TYPE,
  DEFAULT_DEVICE_DENSITY,
  DEFAULT_DEVICE_MODEL,
  DEFAULT_DEVICE_OS,
  DEFAULT_DEVICE_SCREEN,
  DEFAULT_DEVICE_TYPE,
  DEFAULT_EVENT_STREAM_URL,
  buildDeviceInfo,
  buildUserAgent,
} from "./config";
import { type Logger, noopLogger } from "./logger";
import { DEFAULT_RETRY_POLICY, type RetryPolicy, buildRetryMiddleware } from "./retry";
import type { Session, SessionStore } from "./session";
import type { Tokens } from "./tokens";
import { emptyTokens } from "./tokens";
import {
  buildAuthRefreshMiddleware,
  buildHeadersMiddleware,
  buildHostRoutingMiddleware,
} from "./transport";

export interface ClientOptions {
  baseURL?: string;
  // Auxiliary host for activity-feed operations (see gen/hostRoutes).
  // Defaults to https://cas.yay.space; point at a test server to
  // exercise host-routed endpoints offline.
  cassandraBaseURL?: string;
  eventStreamURL?: string;
  apiKey?: string;
  apiVersionKey?: string;
  apiVersionName?: string;
  appVersion?: string;
  userAgent?: string;
  deviceInfo?: string;
  deviceUUID?: string;
  connectionType?: string;
  connectionSpeed?: string;
  acceptLanguage?: string;
  sessionStore?: SessionStore;
  fetchApi?: typeof fetch;
  retryPolicy?: RetryPolicy;
  // Structured logger. Omitted = silent (the SDK is a library and
  // writes nothing by default). See PORTING.md §12.2.
  logger?: Logger;
  // Injectable WebSocket constructor for openEventStream. Defaults to the
  // platform global. Tests supply an in-memory pair so the event-stream
  // suite stays hermetic and dependency-free.
  webSocketFactory?: WebSocketFactory;
}

// WebSocketFactory builds a connection for the event stream. The shape is
// the standard WebSocket constructor; only the slice event_stream.ts uses
// is required (see WebSocketLike there).
export type WebSocketFactory = (url: string) => import("./event_stream").WebSocketLike;

// Stub UUIDv4 generator. PORTING.md §14 requires crypto-grade randomness
// and a failure-MUST-throw contract; a fuller implementation lands when
// the auth flow is wired up.
function newUUIDv4(): string {
  const c = globalThis.crypto;
  if (!c || typeof c.randomUUID !== "function") {
    throw new Error("yaylib: crypto.randomUUID unavailable; refusing to fabricate device UUID");
  }
  return c.randomUUID();
}

export class Client {
  // Configuration constants — public for inspection / override at runtime.
  readonly baseURL: string;
  readonly cassandraBaseURL: string;
  readonly eventStreamURL: string;
  readonly apiKey: string;
  readonly apiVersionKey: string;
  readonly apiVersionName: string;
  readonly appVersion: string;
  readonly userAgent: string;
  readonly deviceInfo: string;
  readonly connectionType: string;
  readonly connectionSpeed: string;
  readonly acceptLanguage: string;
  readonly retryPolicy: RetryPolicy;
  readonly logger: Logger;

  // Per-instance mutable state. tokens / userID / currentEmail / clientIP
  // are written by login / refresh / loadSession; everything else is
  // assigned once.
  private _deviceUUID: string;
  private _tokens: Tokens = emptyTokens();
  private _userID = 0;
  private _currentEmail = "";
  private _clientIP = "";
  // Single-flight refresh tracking — concurrent 401s collapse onto one
  // OAuth /token round-trip (PORTING.md §6.1).
  private _refreshInFlight: Promise<boolean> | null = null;
  private readonly _fetchApi?: typeof fetch;
  private readonly _webSocketFactory?: WebSocketFactory;
  private readonly _openStreams = new Set<EventStream>();
  readonly sessionStore?: SessionStore;

  // Per-tag API instances. The flat operation surface (PORTING.md §2)
  // builds on top of these via code-gen in a later step; for now callers
  // can drop down to `client.postsAPI.<op>` etc. when needed.
  readonly activitiesAPI: ActivitiesApi;
  readonly appsAPI: AppsApi;
  readonly authAPI: AuthApi;
  readonly bucketsAPI: BucketsApi;
  readonly callsAPI: CallsApi;
  readonly chatRoomsAPI: ChatRoomsApi;
  readonly conversationsAPI: ConversationsApi;
  readonly emailVerificationUrlsAPI: EmailVerificationUrlsApi;
  readonly gamesAPI: GamesApi;
  readonly genresAPI: GenresApi;
  readonly giftsAPI: GiftsApi;
  readonly groupMuteAPI: GroupMuteApi;
  readonly groupsAPI: GroupsApi;
  readonly hiddenAPI: HiddenApi;
  readonly moderationAPI: ModerationApi;
  readonly notificationSettingsAPI: NotificationSettingsApi;
  readonly pinnedAPI: PinnedApi;
  readonly postsAPI: PostsApi;
  readonly receivedGiftsAPI: ReceivedGiftsApi;
  readonly stickerPacksAPI: StickerPacksApi;
  readonly surveysAPI: SurveysApi;
  readonly threadsAPI: ThreadsApi;
  readonly usersAPI: UsersApi;

  constructor(opts: ClientOptions = {}) {
    this.baseURL = opts.baseURL ?? DEFAULT_BASE_URL;
    this.cassandraBaseURL = opts.cassandraBaseURL ?? DEFAULT_CASSANDRA_BASE_URL;
    this.eventStreamURL = opts.eventStreamURL ?? DEFAULT_EVENT_STREAM_URL;
    this.apiKey = opts.apiKey ?? DEFAULT_API_KEY;
    this.apiVersionKey = opts.apiVersionKey ?? DEFAULT_API_VERSION_KEY;
    this.apiVersionName = opts.apiVersionName ?? DEFAULT_API_VERSION_NAME;
    this.appVersion = opts.appVersion ?? DEFAULT_APP_VERSION;
    this.userAgent =
      opts.userAgent ??
      buildUserAgent({
        deviceType: DEFAULT_DEVICE_TYPE,
        deviceOS: DEFAULT_DEVICE_OS,
        deviceDensity: DEFAULT_DEVICE_DENSITY,
        deviceScreen: DEFAULT_DEVICE_SCREEN,
        deviceModel: DEFAULT_DEVICE_MODEL,
      });
    this.deviceInfo = opts.deviceInfo ?? buildDeviceInfo(this.appVersion, this.userAgent);
    this._deviceUUID = opts.deviceUUID ?? newUUIDv4();
    this.connectionType = opts.connectionType ?? DEFAULT_CONNECTION_TYPE;
    this.connectionSpeed = opts.connectionSpeed ?? DEFAULT_CONNECTION_SPEED;
    this.acceptLanguage = opts.acceptLanguage ?? DEFAULT_ACCEPT_LANGUAGE;
    this.sessionStore = opts.sessionStore;
    this._fetchApi = opts.fetchApi;
    this._webSocketFactory = opts.webSocketFactory;
    this.retryPolicy = opts.retryPolicy ?? DEFAULT_RETRY_POLICY;
    this.logger = opts.logger ?? noopLogger;

    // Middleware order is significant:
    //   1. host-routing — rewrites the origin for the few operations
    //      served from an auxiliary host, so every later middleware
    //      observes the final URL.
    //   2. headers — every other middleware sees the canonical request
    //      shape the server expects.
    //   3. auth-refresh — must run before retry so a 401 triggers a
    //      refresh + replay rather than counting toward the retry budget.
    //   4. retry — handles 429 / 5xx / transport errors on the response
    //      that auth-refresh either left intact or already retried.
    const config = new Configuration({
      basePath: this.baseURL,
      fetchApi: opts.fetchApi,
      middleware: [
        buildHostRoutingMiddleware({
          cassandraBaseURL: this.cassandraBaseURL,
          logger: this.logger,
        }),
        buildHeadersMiddleware({
          logger: this.logger,
          userAgent: this.userAgent,
          appVersion: this.apiVersionName,
          deviceInfo: this.deviceInfo,
          deviceUUID: this._deviceUUID,
          connectionType: this.connectionType,
          connectionSpeed: this.connectionSpeed,
          acceptLanguage: this.acceptLanguage,
          apiKey: this.apiKey,
          clientIP: () => this._clientIP,
          accessToken: () => this._tokens.access,
        }),
        buildAuthRefreshMiddleware({
          refresh: (stale) => this._tryRefresh(stale),
          accessToken: () => this._tokens.access,
          fetchApi: opts.fetchApi,
        }),
        buildRetryMiddleware({
          policy: this.retryPolicy,
          logger: this.logger,
          fetchApi: opts.fetchApi,
        }),
      ],
    });

    this.activitiesAPI = new ActivitiesApi(config);
    this.appsAPI = new AppsApi(config);
    this.authAPI = new AuthApi(config);
    this.bucketsAPI = new BucketsApi(config);
    this.callsAPI = new CallsApi(config);
    this.chatRoomsAPI = new ChatRoomsApi(config);
    this.conversationsAPI = new ConversationsApi(config);
    this.emailVerificationUrlsAPI = new EmailVerificationUrlsApi(config);
    this.gamesAPI = new GamesApi(config);
    this.genresAPI = new GenresApi(config);
    this.giftsAPI = new GiftsApi(config);
    this.groupMuteAPI = new GroupMuteApi(config);
    this.groupsAPI = new GroupsApi(config);
    this.hiddenAPI = new HiddenApi(config);
    this.moderationAPI = new ModerationApi(config);
    this.notificationSettingsAPI = new NotificationSettingsApi(config);
    this.pinnedAPI = new PinnedApi(config);
    this.postsAPI = new PostsApi(config);
    this.receivedGiftsAPI = new ReceivedGiftsApi(config);
    this.stickerPacksAPI = new StickerPacksApi(config);
    this.surveysAPI = new SurveysApi(config);
    this.threadsAPI = new ThreadsApi(config);
    this.usersAPI = new UsersApi(config);
  }

  // Tokens snapshot — mutating the returned value does not affect the
  // Client (PORTING.md §4).
  tokens(): Tokens {
    return { ...this._tokens };
  }

  setTokens(access: string, refresh: string): void {
    this._tokens = { access, refresh };
  }

  get deviceUUID(): string {
    return this._deviceUUID;
  }

  get userID(): number {
    return this._userID;
  }

  /** Internal — used by auth.ts / session.ts when activating a session. */
  setLoginIdentity(email: string, userID: number): void {
    this._currentEmail = email;
    this._userID = userID;
  }

  /** Internal — used by auth.ts when LoadSession restores a persisted UUID. */
  setDeviceUUID(uuid: string): void {
    this._deviceUUID = uuid;
  }

  get currentEmail(): string {
    return this._currentEmail;
  }

  get clientIP(): string {
    return this._clientIP;
  }

  // Background lifecycle hooks (close, lazy client-IP fetch, event-stream
  // tracking) land alongside the auth flow in a later commit. For now
  // close is a no-op so callers can wire it up unconditionally.
  async close(): Promise<void> {
    const streams = [...this._openStreams];
    this._openStreams.clear();
    await Promise.all(streams.map((s) => s.close().catch(() => undefined)));
  }

  /**
   * LoginWithEmail with transparent session caching (PORTING.md §6).
   *
   * If a SessionStore is configured, a hit returns the persisted session
   * synthesized into a LoginUserResponse without issuing any HTTP. A miss
   * (or `.noCache()`) issues the OAuth login, persists the session, and
   * activates tokens / userID / email on the client.
   *
   *   await client.loginWithEmail()
   *     .email("...").password("...").execute();
   *
   * For 2FA-required accounts, chain `.twoFACode("...")` on the retry.
   */
  loginWithEmail(): import("./auth").LoginWithEmailBuilder {
    // Lazy import to avoid a circular module load at construction time.
    // eslint-disable-next-line @typescript-eslint/no-var-requires
    const { loginWithEmail } = require("./auth") as typeof import("./auth");
    return loginWithEmail(this);
  }

  // rawFetch is the unwrapped fetch — no generated middleware, so no
  // Yay! headers and no `Authorization: Bearer …`. Used for the S3
  // presigned PUT (whose own signature would clash) and the WebSocket
  // handshake (which authenticates via the `token` query param).
  // PORTING.md §6.1: re-issuing through the middleware-wrapped fetch
  // recurses infinitely — always use this for out-of-band requests.
  private get _rawFetch(): typeof fetch {
    const f = this._fetchApi ?? globalThis.fetch;
    return (input: any, init?: any) => f(input, init);
  }

  private _uploadDeps(): UploadDeps {
    return {
      userID: () => this._userID,
      getPresignedURLs: async (fileNames: string[]) => {
        const resp = await this.bucketsAPI.getBucketPresignedUrls({ fileNames });
        return (resp.presignedUrls ?? []).map((p) => ({
          filename: p.filename ?? "",
          url: p.url ?? "",
        }));
      },
      getVideoPresignedURL: async (videoFileName: string) => {
        const resp = await this.usersAPI.getUserPresignedUrl({ videoFileName });
        return resp.presignedUrl ?? "";
      },
      rawFetch: (url, init) => this._rawFetch(url, init),
    };
  }

  // ---- Uploads (PORTING.md §8) ---------------------------------------
  // Self-user methods read this.userID and reject with a clear "not
  // logged in" error when it is zero; room/group methods take the ID as
  // the first argument. Each returns the server-canonical (s3/-prefixed)
  // filename callers pass back to editUser / createPost / etc.

  uploadAvatarImage(file: Upload): Promise<string> {
    return uploadSingleImage(this._uploadDeps(), userAvatarUpload(this._requireUserID()), file);
  }

  uploadCoverImage(file: Upload): Promise<string> {
    return uploadSingleImage(this._uploadDeps(), userCoverUpload(this._requireUserID()), file);
  }

  uploadPostImages(files: Upload[]): Promise<string[]> {
    return uploadImages(this._uploadDeps(), userPostUpload(this._requireUserID()), files);
  }

  uploadChatMessageImages(roomID: number, files: Upload[]): Promise<string[]> {
    return uploadImages(
      this._uploadDeps(),
      chatMessageUpload(roomID, this._requireUserID()),
      files,
    );
  }

  uploadChatBackgroundImage(roomID: number, file: Upload): Promise<string> {
    return uploadSingleImage(this._uploadDeps(), chatBackgroundUpload(roomID), file);
  }

  uploadGroupIconImage(groupID: number, file: Upload): Promise<string> {
    return uploadSingleImage(this._uploadDeps(), groupIconUpload(groupID), file);
  }

  uploadGroupCoverImage(groupID: number, file: Upload): Promise<string> {
    return uploadSingleImage(this._uploadDeps(), groupCoverUpload(groupID), file);
  }

  uploadThreadIconImage(groupID: number, file: Upload): Promise<string> {
    return uploadSingleImage(this._uploadDeps(), threadIconUpload(groupID), file);
  }

  uploadSignupAvatarImage(file: Upload): Promise<string> {
    return uploadSingleImage(this._uploadDeps(), signupAvatarUpload(), file);
  }

  uploadGroupCreationIconImage(file: Upload): Promise<string> {
    return uploadSingleImage(this._uploadDeps(), groupCreationIconUpload(), file);
  }

  uploadGroupCreationCoverImage(file: Upload): Promise<string> {
    return uploadSingleImage(this._uploadDeps(), groupCreationCoverUpload(), file);
  }

  uploadReportImages(files: Upload[]): Promise<string[]> {
    return uploadImages(this._uploadDeps(), reportUpload(), files);
  }

  uploadVideoCallSnapshotImage(file: Upload): Promise<string> {
    return uploadSingleImage(this._uploadDeps(), videoCallSnapshotUpload(), file);
  }

  uploadVideo(body: VideoBody): Promise<string> {
    return uploadVideo(this._uploadDeps(), body);
  }

  /** Lower-level upload entry point for callers that already hold a category. */
  uploadCategoryImages(category: UploadCategory, files: Upload[]): Promise<string[]> {
    return uploadImages(this._uploadDeps(), category, files);
  }

  private _requireUserID(): number {
    if (!this._userID) {
      throw new Error("yaylib: not logged in (call loginWithEmail before user-bound uploads)");
    }
    return this._userID;
  }

  // ---- Event stream (PORTING.md §10) ---------------------------------

  /**
   * Open a multiplexed real-time event channel. The per-connection token
   * is fetched via getWebSocketToken using the client's current
   * Authorization, so set the access token first (loginWithEmail /
   * setTokens / a restored session) — an unauthenticated client surfaces
   * the underlying 401 as the open error.
   */
  async openEventStream(opts: EventStreamOptions = {}): Promise<EventStream> {
    const stream = await openEventStream(
      {
        eventStreamURL: this.eventStreamURL,
        apiVersionName: this.apiVersionName,
        logger: this.logger,
        getWebSocketToken: async () => {
          const resp = await this.usersAPI.getWebSocketToken();
          return resp.token ?? "";
        },
        webSocketFactory: this._webSocketFactory,
      },
      opts,
    );
    this._openStreams.add(stream);
    void stream.done().finally(() => this._openStreams.delete(stream));
    return stream;
  }

  // ---- Raw escape hatch (PORTING.md §11.3) ----------------------------

  /**
   * executeRaw runs a generated `*Raw` operation and returns the raw
   * response bytes alongside the HTTP response, bypassing the typed JSON
   * decode entirely. Contract (PORTING.md §11.3):
   *
   *   - HTTP success      → { body, httpResponse, error: undefined }
   *   - HTTP error (4xx/5xx) → { body, httpResponse, error: APIError }
   *     so codeOf(error) still works
   *   - transport failure → { body: undefined, httpResponse: undefined,
   *     error: APIError }
   *
   * Typed JSON-decode failures never propagate — the raw bytes are read
   * straight off the response.
   *
   *   const { body, httpResponse } =
   *     await client.executeRaw(() => client.usersAPI.getUserTimestampRaw());
   */
  async executeRaw<T>(
    call: () => Promise<ApiResponse<T>>,
  ): Promise<{ body?: Uint8Array; httpResponse?: Response; error?: APIError }> {
    try {
      const api = await call();
      const body = new Uint8Array(await api.raw.arrayBuffer());
      return { body, httpResponse: api.raw, error: undefined };
    } catch (err) {
      if (err instanceof ResponseError) {
        const body = err.response
          ? new Uint8Array(await err.response.arrayBuffer())
          : new Uint8Array();
        const error = new APIError({
          message: err.message || err.response?.statusText || "API error",
          response: err.response,
          body,
          cause: err,
        });
        return { body, httpResponse: err.response, error };
      }
      // Transport failure (FetchError or anything non-HTTP).
      return { error: await asAPIError(err) };
    }
  }

  /** Internal — invoked by the auth-refresh middleware (transport.ts). */
  private _tryRefresh(staleAccess: string): Promise<boolean> {
    if (this._refreshInFlight) return this._refreshInFlight;
    const promise = this._doRefresh(staleAccess).finally(() => {
      this._refreshInFlight = null;
    });
    this._refreshInFlight = promise;
    return promise;
  }

  private async _doRefresh(staleAccess: string): Promise<boolean> {
    // Another concurrent caller already rotated the token while we were
    // queued. Skip the OAuth round-trip and let the caller retry with
    // the current Bearer.
    if (this._tokens.access !== staleAccess) return true;
    if (!this._tokens.refresh) {
      this.logger.debug("token refresh skipped", {
        event: "token_refresh",
        outcome: "no_token",
      });
      return false;
    }

    const form = new URLSearchParams({
      grant_type: "refresh_token",
      refresh_token: this._tokens.refresh,
    });
    const fetchImpl = this._fetchApi ?? globalThis.fetch;

    let res: Response;
    try {
      res = await fetchImpl(`${this.baseURL}/api/v1/oauth/token`, {
        method: "POST",
        headers: {
          Authorization: `Basic ${btoa(this.apiKey)}`,
          "Content-Type": "application/x-www-form-urlencoded",
        },
        body: form.toString(),
      });
    } catch {
      this.logger.debug("token refresh failed", {
        event: "token_refresh",
        outcome: "failed",
      });
      return false;
    }

    if (!res.ok) {
      try {
        await res.arrayBuffer();
      } catch {
        /* ignore */
      }
      this.logger.debug("token refresh failed", {
        event: "token_refresh",
        outcome: "failed",
      });
      return false;
    }

    let parsed: { access_token?: string; refresh_token?: string };
    try {
      parsed = (await res.json()) as {
        access_token?: string;
        refresh_token?: string;
      };
    } catch {
      this.logger.debug("token refresh failed", {
        event: "token_refresh",
        outcome: "failed",
      });
      return false;
    }
    if (!parsed.access_token || !parsed.refresh_token) {
      this.logger.debug("token refresh failed", {
        event: "token_refresh",
        outcome: "failed",
      });
      return false;
    }

    this._tokens = {
      access: parsed.access_token,
      refresh: parsed.refresh_token,
    };
    this.logger.debug("token refresh ok", {
      event: "token_refresh",
      outcome: "ok",
    });

    // Persist the rotated session so subsequent restarts don't re-login.
    // PORTING.md §5: persist failure during the 401 auto-refresh path is
    // non-fatal — log and continue.
    if (this.sessionStore && this._currentEmail) {
      const session: Session = {
        email: this._currentEmail,
        userID: this._userID,
        accessToken: this._tokens.access,
        refreshToken: this._tokens.refresh,
        deviceUUID: this._deviceUUID,
        updatedAt: new Date().toISOString(),
      };
      try {
        await this.sessionStore.save(session);
      } catch (err) {
        this.logger.warn("persist refreshed tokens failed", {
          event: "token_persist_fail",
          user_id: this._userID,
          err: err instanceof Error ? err.message : String(err),
        });
      }
    }

    return true;
  }
}
