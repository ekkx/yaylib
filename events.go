package yaylib

import (
	"encoding/json"
	"fmt"
)

// Event is a server-pushed event delivered through a Subscription.
// Concrete event types embed the wire-level fields; switch on the dynamic
// type to dispatch:
//
//	for ev := range sub.Events() {
//	    switch e := ev.(type) {
//	    case *yaylib.NewMessageEvent:
//	        log.Printf("chat %d: %s", e.GetRoomId(), e.GetText())
//	    case *yaylib.VideoProcessedEvent:
//	        log.Printf("video %d ready: %s", e.ID, e.VideoURL)
//	    }
//	}
type Event interface {
	isEvent()
}

// NewMessageEvent is delivered on MessagesChannel when a new chat message
// arrives. It carries the full chat message DTO.
type NewMessageEvent struct {
	RealmMessage
}

func (*NewMessageEvent) isEvent() {}

// VideoProcessedEvent is delivered on MessagesChannel after server-side
// video transcoding finishes for an attachment.
type VideoProcessedEvent struct {
	ID                   int64  `json:"id"`
	VideoProcessed       bool   `json:"video_processed"`
	VideoURL             string `json:"video_url,omitempty"`
	VideoThumbnailURL    string `json:"video_thumbnail_url,omitempty"`
	VideoThumbnailBigURL string `json:"video_thumbnail_big_url,omitempty"`
}

func (*VideoProcessedEvent) isEvent() {}

// ChatDeletedEvent is delivered on ChatRoomChannel when a chat room the
// user belongs to is deleted.
type ChatDeletedEvent struct {
	RoomID int64 `json:"room_id"`
}

func (*ChatDeletedEvent) isEvent() {}

// TotalChatRequestEvent reports the current count of pending chat
// requests addressed to the user.
type TotalChatRequestEvent struct {
	TotalCount int `json:"total_count"`
}

func (*TotalChatRequestEvent) isEvent() {}

// UnsubscribedEvent is delivered when the server forces the user out of a
// channel — typically because they were removed from the chat or group.
type UnsubscribedEvent struct {
	UserIDs []int64 `json:"user_ids"`
}

func (*UnsubscribedEvent) isEvent() {}

// GroupUpdatedEvent fires on GroupUpdatesChannel / GroupPostsChannel when
// a watched group has new content.
type GroupUpdatedEvent struct {
	GroupID int64 `json:"group_id"`
}

func (*GroupUpdatedEvent) isEvent() {}

// CallFinishedEvent fires when a voice call inside a chat room ends.
type CallFinishedEvent struct {
	ID int64 `json:"id"`
}

func (*CallFinishedEvent) isEvent() {}

// RawEvent wraps an event whose discriminator the SDK does not recognize.
// The Name is the server-supplied event tag; Data is the raw JSON payload
// so callers can decode it out-of-band.
type RawEvent struct {
	Name string
	Data json.RawMessage
}

func (*RawEvent) isEvent() {}

// decodeEvent maps a server-supplied event name + payload into a concrete
// Event type. Unknown event names fall back to *RawEvent so callers can
// still observe new server signals before the SDK is updated.
func decodeEvent(name string, data []byte) (Event, error) {
	switch name {
	case "new_message":
		var ev NewMessageEvent
		if err := json.Unmarshal(data, &ev); err != nil {
			return nil, fmt.Errorf("decode new_message: %w", err)
		}
		return &ev, nil
	case "video_processed":
		var ev VideoProcessedEvent
		if err := json.Unmarshal(data, &ev); err != nil {
			return nil, fmt.Errorf("decode video_processed: %w", err)
		}
		return &ev, nil
	case "chat_deleted":
		var ev ChatDeletedEvent
		if err := json.Unmarshal(data, &ev); err != nil {
			return nil, fmt.Errorf("decode chat_deleted: %w", err)
		}
		return &ev, nil
	case "total_chat_request":
		var ev TotalChatRequestEvent
		if err := json.Unmarshal(data, &ev); err != nil {
			return nil, fmt.Errorf("decode total_chat_request: %w", err)
		}
		return &ev, nil
	case "unsubscribed":
		var ev UnsubscribedEvent
		if err := json.Unmarshal(data, &ev); err != nil {
			return nil, fmt.Errorf("decode unsubscribed: %w", err)
		}
		return &ev, nil
	case "new_post":
		var ev GroupUpdatedEvent
		if err := json.Unmarshal(data, &ev); err != nil {
			return nil, fmt.Errorf("decode new_post: %w", err)
		}
		return &ev, nil
	case "conference_call_finished":
		var ev CallFinishedEvent
		if err := json.Unmarshal(data, &ev); err != nil {
			return nil, fmt.Errorf("decode conference_call_finished: %w", err)
		}
		return &ev, nil
	default:
		return &RawEvent{Name: name, Data: append(json.RawMessage(nil), data...)}, nil
	}
}
