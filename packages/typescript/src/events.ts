// PORTING.md §10: the fixed event inventory. The wire `event` name does
// NOT always match the SDK-side type name (e.g. `new_post` carries a
// group-id update, not a post); the SDK names are preserved verbatim so
// user code stays portable across the Go / TS / Python ports.

export interface NewMessageEvent {
  readonly type: "NewMessageEvent";
  // The full chat-message DTO as delivered on the wire. video_thumbnail
  // _big_url is a WebSocket-only field (absent on the REST Message DTO);
  // treat it as nullable.
  readonly message: Record<string, unknown>;
}

export interface VideoProcessedEvent {
  readonly type: "VideoProcessedEvent";
  readonly id: number;
  readonly videoProcessed: boolean;
  readonly videoUrl?: string;
  readonly videoThumbnailUrl?: string;
  readonly videoThumbnailBigUrl?: string;
}

export interface ChatDeletedEvent {
  readonly type: "ChatDeletedEvent";
  readonly roomId: number;
}

export interface TotalChatRequestEvent {
  readonly type: "TotalChatRequestEvent";
  readonly totalCount: number;
}

export interface UnsubscribedEvent {
  readonly type: "UnsubscribedEvent";
  readonly userIds: number[];
}

// GroupUpdatedEvent fires on GroupUpdatesChannel / GroupPostsChannel. Note
// the wire name for the posts case is `new_post` but the payload is a
// group-id update, not a post.
export interface GroupUpdatedEvent {
  readonly type: "GroupUpdatedEvent";
  readonly groupId: number;
}

export interface CallFinishedEvent {
  readonly type: "CallFinishedEvent";
  readonly id: number;
}

// RawEvent wraps an event whose discriminator the SDK does not recognize,
// so callers can still observe new server signals before the SDK is
// updated.
export interface RawEvent {
  readonly type: "RawEvent";
  readonly name: string;
  readonly data: unknown;
}

export type Event =
  | NewMessageEvent
  | VideoProcessedEvent
  | ChatDeletedEvent
  | TotalChatRequestEvent
  | UnsubscribedEvent
  | GroupUpdatedEvent
  | CallFinishedEvent
  | RawEvent;

function num(v: unknown): number {
  return typeof v === "number" ? v : typeof v === "string" ? Number(v) || 0 : 0;
}

function str(v: unknown): string | undefined {
  return typeof v === "string" ? v : undefined;
}

// decodeEvent maps a server-supplied event name + payload into a concrete
// Event. Unknown names fall back to RawEvent. Decoding is tolerant
// (PORTING.md §11): a malformed field degrades to a zero value rather
// than throwing, so one bad payload never kills the read pump.
export function decodeEvent(name: string, data: unknown): Event {
  const d = (data && typeof data === "object" ? data : {}) as Record<string, unknown>;
  switch (name) {
    case "new_message":
      return { type: "NewMessageEvent", message: d };
    case "video_processed":
      return {
        type: "VideoProcessedEvent",
        id: num(d.id),
        videoProcessed: d.video_processed === true,
        videoUrl: str(d.video_url),
        videoThumbnailUrl: str(d.video_thumbnail_url),
        videoThumbnailBigUrl: str(d.video_thumbnail_big_url),
      };
    case "chat_deleted":
      return { type: "ChatDeletedEvent", roomId: num(d.room_id) };
    case "total_chat_request":
      return { type: "TotalChatRequestEvent", totalCount: num(d.total_count) };
    case "unsubscribed":
      return {
        type: "UnsubscribedEvent",
        userIds: Array.isArray(d.user_ids) ? (d.user_ids as unknown[]).map(num) : [],
      };
    case "new_post":
      return { type: "GroupUpdatedEvent", groupId: num(d.group_id) };
    case "conference_call_finished":
      return { type: "CallFinishedEvent", id: num(d.id) };
    default:
      return { type: "RawEvent", name, data };
  }
}
