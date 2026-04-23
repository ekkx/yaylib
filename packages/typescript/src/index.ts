/**
 * Base URL of the Yay! REST API.
 */
export const API_BASE_URL = "https://api.yay.space";

/**
 * WebSocket endpoint for Yay! real-time events (ActionCable).
 */
export const WS_URL = "wss://cable.yay.space";

/**
 * Generate a v4 UUID, suitable for use as the `X-Device-UUID` header on
 * every Yay! API request.
 */
export function createDeviceUuid(): string {
  return (globalThis as { crypto: { randomUUID: () => string } }).crypto.randomUUID();
}
