// Public surface of yaylib. Re-exports the generated DTO / API layer
// from ./gen and the hand-written wrappers from the wrapper modules.

export * from "./gen/index.js";

export {
  Client,
  type ClientOptions,
} from "./client.js";

export {
  type Tokens,
  emptyTokens,
} from "./tokens.js";

export {
  type Session,
  type SessionStore,
  MemorySessionStore,
  NoSessionError,
  SessionSaveFailedError,
  newMemoryStore,
} from "./session.js";

export { FileSessionStore, newSessionStore } from "./session_file.js";

export {
  APIError,
  type ErrorResponse,
  errorResponseOf,
  codeOf,
  asAPIError,
} from "./errors.js";

export * from "./error_codes.js";

export {
  DEFAULT_API_KEY,
  DEFAULT_API_VERSION_KEY,
  DEFAULT_API_VERSION_NAME,
  DEFAULT_APP_VERSION,
  DEFAULT_BASE_URL,
  DEFAULT_EVENT_STREAM_URL,
  MEDIA_CDN_BASE,
  mediaURL,
} from "./config.js";

export {
  type SignedInfo,
  generateSignedInfoAt,
  generateSignedVersion,
  generateXJwt,
} from "./signing.js";

export { type LoginWithEmailParams, loginWithEmail } from "./auth.js";

export {
  type RetryPolicy,
  DEFAULT_RETRY_POLICY,
} from "./retry.js";

export {
  type Upload,
  type UploadCategory,
  type VideoBody,
  MAX_IMAGES_PER_UPLOAD,
  MAX_REPORT_IMAGES_PER_UPLOAD,
} from "./upload.js";

export {
  type Channel,
  chatRoomChannel,
  messagesChannel,
  groupUpdatesChannel,
  groupPostsChannel,
} from "./channels.js";

export type {
  Event,
  NewMessageEvent,
  VideoProcessedEvent,
  ChatDeletedEvent,
  TotalChatRequestEvent,
  UnsubscribedEvent,
  GroupUpdatedEvent,
  CallFinishedEvent,
  RawEvent,
} from "./events.js";

export {
  EventStream,
  Subscription,
  type EventStreamOptions,
  type ReconnectPolicy,
  type WebSocketLike,
} from "./event_stream.js";

export { type WebSocketFactory } from "./client.js";
export type { SignaturePayload } from "./gen/models/index.js";
export { type Logger, type LogFields, noopLogger } from "./logger.js";
