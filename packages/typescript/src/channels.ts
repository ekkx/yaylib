// PORTING.md §10.1: channel identifier strings. The server matches these
// strings literally when routing later frames — the space after `,` is
// significant, and the IDs are JSON STRINGS, not numbers. Do not
// reformat.

export interface Channel {
  readonly identifier: string;
}

// ChatRoomChannel — global chat-room signals (new chat requests, chat
// deletions, forced unsubscribe notices).
export function chatRoomChannel(): Channel {
  return { identifier: `{"channel":"ChatRoomChannel"}` };
}

// MessagesChannel — messages + video_processed events for one chat room.
export function messagesChannel(roomID: number | string): Channel {
  return { identifier: `{"channel":"MessagesChannel", "chat_room_id":"${roomID}"}` };
}

// GroupUpdatesChannel — group state updates across all groups.
export function groupUpdatesChannel(): Channel {
  return { identifier: `{"channel":"GroupUpdatesChannel"}` };
}

// GroupPostsChannel — new posts within one group's timeline.
export function groupPostsChannel(groupID: number | string): Channel {
  return { identifier: `{"channel":"GroupPostsChannel", "group_id":"${groupID}"}` };
}
