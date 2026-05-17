package yaylib

import (
	"encoding/json"
	"testing"

	"github.com/ekkx/yaylib/v2/gen"
)

// Chat-room members and sticker-pack stickers are JSON arrays on the
// wire. These bodies (shaped like the real server responses) must
// decode into typed slices, including the nested collections.

func TestChatRoomsResponseDecodesMembersArray(t *testing.T) {
	const body = `{"next_page_value":1778180251,"result":"success",
		"chat_rooms":[{"id":181237994,"is_group":false,
		"members":[{"id":35152},{"id":7360590}]}]}`
	var r gen.ChatRoomsResponse
	if err := json.Unmarshal([]byte(body), &r); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	rooms := r.GetChatRooms()
	if len(rooms) != 1 {
		t.Fatalf("chat_rooms len = %d, want 1", len(rooms))
	}
	if got := rooms[0].GetMembers(); len(got) != 2 {
		t.Fatalf("members len = %d, want 2", len(got))
	}
}

func TestChatRoomResponseDecodesNestedRoom(t *testing.T) {
	const body = `{"chat":{"id":181237994,"is_group":false,
		"members":[{"id":1}]}}`
	var r gen.ChatRoomResponse
	if err := json.Unmarshal([]byte(body), &r); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	chat := r.GetChat()
	if chat.GetId() != 181237994 {
		t.Fatalf("chat.id = %d, want 181237994", chat.GetId())
	}
	if len(chat.GetMembers()) != 1 {
		t.Fatalf("chat.members len = %d, want 1", len(chat.GetMembers()))
	}
}

func TestStickerPacksResponseDecodesStickersArray(t *testing.T) {
	const body = `{"sticker_packs":[{"id":25,"name":"x",
		"stickers":[{"id":1261,"extension":"png"},{"id":1262,"extension":"png"}]}]}`
	var r gen.StickerPacksResponse
	if err := json.Unmarshal([]byte(body), &r); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	packs := r.GetStickerPacks()
	if len(packs) != 1 {
		t.Fatalf("sticker_packs len = %d, want 1", len(packs))
	}
	if got := packs[0].GetStickers(); len(got) != 2 {
		t.Fatalf("stickers len = %d, want 2", len(got))
	}
}
