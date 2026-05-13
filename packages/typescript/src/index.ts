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

export {
  APIError,
  type ErrorResponse,
  errorResponseOf,
  codeOf,
  asAPIError,
} from "./errors";

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
