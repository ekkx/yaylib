# Chat-room members and sticker-pack stickers are JSON arrays on the
# wire. These bodies (shaped like the real server responses) must
# decode into typed lists, including the nested collections.

from yaylib.models.chat_rooms_response import ChatRoomsResponse
from yaylib.models.sticker_packs_response import StickerPacksResponse


def test_chat_rooms_members_decode_as_list():
    r = ChatRoomsResponse.from_dict(
        {
            "next_page_value": 1778180251,
            "result": "success",
            "chat_rooms": [
                {"id": 181237994, "is_group": False, "members": [{"id": 35152}, {"id": 7360590}]}
            ],
        }
    )
    assert r.chat_rooms is not None and len(r.chat_rooms) == 1
    assert isinstance(r.chat_rooms[0].members, list)
    assert len(r.chat_rooms[0].members) == 2


def test_sticker_packs_stickers_decode_as_list():
    r = StickerPacksResponse.from_dict(
        {
            "sticker_packs": [
                {
                    "id": 25,
                    "name": "x",
                    "stickers": [
                        {"id": 1261, "extension": "png"},
                        {"id": 1262, "extension": "png"},
                    ],
                }
            ]
        }
    )
    assert r.sticker_packs is not None and len(r.sticker_packs) == 1
    assert isinstance(r.sticker_packs[0].stickers, list)
    assert len(r.sticker_packs[0].stickers) == 2
