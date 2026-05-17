// Public surface of yaylib. Re-exports the generated DTO / API layer
// from ./gen and the hand-written wrappers from the wrapper modules.

export * from "./gen";

export {
  Client,
  type ClientOptions,
} from "./client";

export {
  type Tokens,
  emptyTokens,
} from "./tokens";

export {
  type Session,
  type SessionStore,
  MemorySessionStore,
  NoSessionError,
  SessionSaveFailedError,
  newMemoryStore,
} from "./session";

export { FileSessionStore, newSessionStore } from "./session_file";

export {
  APIError,
  type ErrorResponse,
  errorResponseOf,
  codeOf,
  asAPIError,
} from "./errors";

export * from "./error_codes";

export {
  DEFAULT_API_KEY,
  DEFAULT_API_VERSION_KEY,
  DEFAULT_API_VERSION_NAME,
  DEFAULT_APP_VERSION,
  DEFAULT_BASE_URL,
  DEFAULT_EVENT_STREAM_URL,
  MEDIA_CDN_BASE,
  mediaURL,
} from "./config";

export {
  type SignedInfo,
  generateSignedInfoAt,
  generateSignedVersion,
  generateXJwt,
} from "./signing";

export { type LoginWithEmailBuilder, loginWithEmail } from "./auth";

export {
  type RetryPolicy,
  DEFAULT_RETRY_POLICY,
} from "./retry";

export {
  type Upload,
  type UploadCategory,
  type VideoBody,
  MAX_IMAGES_PER_UPLOAD,
  MAX_REPORT_IMAGES_PER_UPLOAD,
} from "./upload";

export {
  type Channel,
  chatRoomChannel,
  messagesChannel,
  groupUpdatesChannel,
  groupPostsChannel,
} from "./channels";

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
} from "./events";

export {
  EventStream,
  Subscription,
  type EventStreamOptions,
  type ReconnectPolicy,
  type WebSocketLike,
} from "./event_stream";

export { type WebSocketFactory } from "./client";
export { type Logger, type LogFields, noopLogger } from "./logger";
