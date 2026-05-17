// Embedded wire constants. PORTING.md §3.1 fixes the defaults so a
// freshly-constructed Client speaks to the Yay! servers without any
// explicit option. Each constant is overridable via the corresponding
// ClientOptions field.

export const DEFAULT_API_KEY =
  "ccd59ee269c01511ba763467045c115779fcae3050238a252f1bd1a4b65cfec6";

// HS256 key used by signed_version and X-Jwt.
export const DEFAULT_API_VERSION_KEY = "34c8c1cdf29b46a492e8ec58e6db37ec";

// Value sent as X-App-Version and embedded in signed_version.
export const DEFAULT_API_VERSION_NAME = "4.26";

// Build version embedded in X-Device-Info.
export const DEFAULT_APP_VERSION = "4.26.1";

export const DEFAULT_DEVICE_TYPE = "android";
export const DEFAULT_DEVICE_OS = "11";
export const DEFAULT_DEVICE_DENSITY = "3.5";
export const DEFAULT_DEVICE_SCREEN = "1440x2960";
export const DEFAULT_DEVICE_MODEL = "Galaxy S9";

export const DEFAULT_CONNECTION_TYPE = "wifi";
export const DEFAULT_CONNECTION_SPEED = "0 kbps";
export const DEFAULT_ACCEPT_LANGUAGE = "ja";

export const DEFAULT_BASE_URL = "https://api.yay.space";

// Auxiliary host a few activity-feed operations are served from
// (see gen/hostRoutes). Requests for those operations are routed here
// instead of DEFAULT_BASE_URL; override via ClientOptions.cassandraBaseURL
// (e.g. to point at a test server).
export const DEFAULT_CASSANDRA_BASE_URL = "https://cas.yay.space";

export const DEFAULT_EVENT_STREAM_URL = "wss://cable.yay.space";
export const MEDIA_CDN_BASE = "https://cdn.yay.space/uploads";

// Wire pepper concatenated after APIKey + DeviceUUID + timestamp before
// MD5-hashing for signed_info.
export const SIGNED_INFO_SHARED_KEY = "yayZ1";

// Platform identifier prepended to APIVersionName inside the
// signed_version HMAC payload. Hard-coded; deviating risks per-platform
// allowlist rejection.
export const SIGNED_VERSION_PLATFORM = "yay_android";

export function buildUserAgent(opts: {
  deviceType: string;
  deviceOS: string;
  deviceDensity: string;
  deviceScreen: string;
  deviceModel: string;
}): string {
  return `${opts.deviceType} ${opts.deviceOS} (${opts.deviceDensity}x ${opts.deviceScreen} ${opts.deviceModel})`;
}

export function buildDeviceInfo(appVersion: string, userAgent: string): string {
  return `yay ${appVersion} ${userAgent}`;
}

// Returns the public CDN URL for a filename returned by any Upload* method.
// Use this rather than reading profile_icon / cover_image / attachment_url
// from API responses, which can carry a legacy prefix that no longer
// resolves.
export function mediaURL(filename: string): string {
  if (!filename) return "";
  return `${MEDIA_CDN_BASE}/${filename.replace(/^\/+/, "")}`;
}
