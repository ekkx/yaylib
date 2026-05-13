// PORTING.md §2 / §3: Client is the single entry point. Every operation
// MUST be reachable as a top-level method on the client; the SDK uses
// composition over the generated tag-services (AuthApi / PostsApi / ...)
// to expose all operations without a `client.posts.` step.
//
// The full operation surface lives across 23 generated *Api classes. Each
// is instantiated against the same Configuration so they share the
// middleware chain (headers / auth / retry). A code-generation pass
// equivalent to Go's `cmd/gen-aliases` will eventually emit the flat
// promotion layer; for the spike we expose the per-service instances on
// the client so callers can reach operations as
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
import { Configuration } from "./gen/runtime";

import {
  DEFAULT_ACCEPT_LANGUAGE,
  DEFAULT_API_KEY,
  DEFAULT_API_VERSION_KEY,
  DEFAULT_API_VERSION_NAME,
  DEFAULT_APP_VERSION,
  DEFAULT_BASE_URL,
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
import type { SessionStore } from "./session";
import type { Tokens } from "./tokens";
import { emptyTokens } from "./tokens";
import { buildHeadersMiddleware } from "./transport";

export interface ClientOptions {
  baseURL?: string;
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
}

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

  // Per-instance mutable state. tokens / userID / currentEmail / clientIP
  // are written by login / refresh / loadSession; everything else is
  // assigned once.
  private _deviceUUID: string;
  private _tokens: Tokens = emptyTokens();
  private _userID = 0;
  private _currentEmail = "";
  private _clientIP = "";
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

    const config = new Configuration({
      basePath: this.baseURL,
      fetchApi: opts.fetchApi,
      middleware: [
        buildHeadersMiddleware({
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
    /* TODO: cancel lazy fetches and open event streams. */
  }
}
