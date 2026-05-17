// Chat-room members and sticker-pack stickers are JSON arrays on the
// wire. These bodies (shaped like the real server responses) must
// decode into typed arrays, including the nested collections.

import { ChatRoomsResponseFromJSON } from "../src/gen/models/ChatRoomsResponse";
import { StickerPacksResponseFromJSON } from "../src/gen/models/StickerPacksResponse";

let failed = 0;
function assert(name: string, cond: boolean, detail = ""): void {
  if (cond) console.log(`PASS ${name}`);
  else {
    failed++;
    console.log(`FAIL ${name}${detail ? `\n  ${detail}` : ""}`);
  }
}

const rooms = ChatRoomsResponseFromJSON({
  next_page_value: 1778180251,
  result: "success",
  chat_rooms: [{ id: 181237994, is_group: false, members: [{ id: 35152 }, { id: 7360590 }] }],
});
assert("chat_rooms decodes", rooms.chatRooms?.length === 1);
assert(
  "nested members is an array",
  Array.isArray(rooms.chatRooms?.[0]?.members) && rooms.chatRooms![0].members!.length === 2,
);

const packs = StickerPacksResponseFromJSON({
  sticker_packs: [
    { id: 25, name: "x", stickers: [{ id: 1261, extension: "png" }, { id: 1262, extension: "png" }] },
  ],
});
assert("sticker_packs decodes", packs.stickerPacks?.length === 1);
assert(
  "nested stickers is an array",
  Array.isArray(packs.stickerPacks?.[0]?.stickers) && packs.stickerPacks![0].stickers!.length === 2,
);

process.exit(failed === 0 ? 0 : 1);
