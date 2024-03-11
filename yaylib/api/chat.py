"""
MIT License

Copyright (c) 2023 ekkx

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
"""

from __future__ import annotations

from .. import client
from ..responses import (
    ChatRoomResponse,
    ChatRoomsResponse,
    TotalChatRequestResponse,
    CreateChatRoomResponse,
    FollowUsersResponse,
    GifsDataResponse,
    MessageResponse,
    MessagesResponse,
    StickerPacksResponse,
    UnreadStatusResponse,
)


class ChatAPI(object):
    def __init__(self, base: client.BaseClient) -> None:
        self.__base = base

    def accept_chat_requests(self, chat_room_ids: list[int]):
        return self.__base._request(
            "POST",
            route=f"/v1/chat_rooms/accept_chat_request",
            payload={"chat_room_ids": chat_room_ids},
            bypass_delay=True,
        )

    def check_unread_status(self, from_time: int) -> UnreadStatusResponse:
        return self.__base._request(
            "GET",
            route=f"/v1/chat_rooms/unread_status",
            params={"from_time": from_time},
            data_type=UnreadStatusResponse,
            bypass_delay=True,
        )

    def create_group_chat(
        self,
        name: str,
        with_user_ids: list[int],
        icon_filename: str = None,
        background_filename: str = None,
    ) -> CreateChatRoomResponse:
        return self.__base._request(
            "POST",
            route=f"/v3/chat_rooms/new",
            payload={
                "name": name,
                "with_user_ids": with_user_ids,
                "icon_filename": icon_filename,
                "background_filename": background_filename,
            },
            data_type=CreateChatRoomResponse,
        )

    def create_private_chat(
        self, with_user_id: int, matching_id: int = None, hima_chat: bool = False
    ) -> CreateChatRoomResponse:
        return self.__base._request(
            "POST",
            route=f"/v1/chat_rooms/new",
            payload={
                "with_user_id": with_user_id,
                "matching_id": matching_id,
                "hima_chat": hima_chat,
            },
            data_type=CreateChatRoomResponse,
        )

    def delete_background(self, room_id: int):
        return self.__base._request(
            "DELETE",
            route=f"/v2/chat_rooms/{room_id}/background",
            bypass_delay=True,
        )

    def delete_message(self, room_id: int, message_id: int):
        return self.__base._request(
            "DELETE",
            route=f"/v1/chat_rooms/{room_id}/messages/{message_id}/delete",
            bypass_delay=True,
        )

    def edit_chat_room(
        self,
        chat_room_id: int,
        name: str,
        icon_filename: str = None,
        background_filename: str = None,
    ):
        return self.__base._request(
            "POST",
            route=f"/v1/chat_rooms/{chat_room_id}/edit",
            payload={
                "name": name,
                "icon_filename": icon_filename,
                "background_filename": background_filename,
            },
            bypass_delay=True,
        )

    def get_chatable_users(
        self,
        # @Body @Nullable SearchUsersRequest searchUsersRequest
        from_follow_id: int = None,
        from_timestamp: int = None,
        order_by: str = None,
    ) -> FollowUsersResponse:
        return self.__base._request(
            "POST",
            route=f"/v1/users/followings/chatable",
            payload={
                "from_follow_id": from_follow_id,
                "from_timestamp": from_timestamp,
                "order_by": order_by,
            },
            data_type=FollowUsersResponse,
        )

    def get_gifs_data(self) -> GifsDataResponse:
        return self.__base._request(
            "GET",
            route=f"/v1/hidden/chats",
            data_type=GifsDataResponse,
            bypass_delay=True,
        )

    def get_hidden_chat_rooms(self, **params) -> ChatRoomsResponse:
        """

        Parameters:
        ---------------

            - from_timestamp: int - (optional)
            - number: int - (optional)

        """
        return self.__base._request(
            "GET",
            route=f"/v1/hidden/chats",
            params=params,
            data_type=ChatRoomsResponse,
            bypass_delay=True,
        )

    def get_main_chat_rooms(self, from_timestamp: int = None) -> ChatRoomsResponse:
        params = {}
        if from_timestamp:
            params["from_timestamp"] = from_timestamp
        return self.__base._request(
            "GET",
            route=f"/v1/chat_rooms/main_list",
            params=params,
            data_type=ChatRoomsResponse,
            bypass_delay=True,
        )

    def get_messages(self, chat_room_id: int, **params) -> MessagesResponse:
        """

        Parameters:
        ---------------
            - from_message_id: int - (optional)
            - to_message_id: int - (optional)

        """
        return self.__base._request(
            "GET",
            route=f"/v2/chat_rooms/{chat_room_id}/messages",
            params=params,
            data_type=MessagesResponse,
            bypass_delay=True,
        )

    def get_request_chat_rooms(self, **params) -> ChatRoomsResponse:
        """

        Parameters:
        -----------

            - number: int (optional)
            - from_timestamp: int (optional)

        """
        return self.__base._request(
            "GET",
            route=f"/v1/chat_rooms/request_list",
            params=params,
            data_type=ChatRoomsResponse,
            bypass_delay=True,
        )

    def get_chat_room(self, chat_room_id: int) -> ChatRoomResponse:
        return self.__base._request(
            "GET",
            route=f"/v2/chat_rooms/{chat_room_id}",
            data_type=ChatRoomResponse,
            bypass_delay=True,
        )

    def get_sticker_packs(self) -> StickerPacksResponse:
        return self.__base._request(
            "GET",
            route="/v2/sticker_packs",
            data_type=StickerPacksResponse,
            bypass_delay=True,
        )

    def get_total_chat_requests(self) -> TotalChatRequestResponse:
        return self.__base._request(
            "GET",
            route=f"/v1/chat_rooms/total_chat_request",
            data_type=TotalChatRequestResponse,
        )

    def hide_chat(self, chat_room_id: int):
        return self.__base._request(
            "POST",
            route=f"/v1/hidden/chats",
            payload={"chat_room_id": chat_room_id},
        )

    def invite_to_chat(self, chat_room_id: int, user_ids: list[int]):
        return self.__base._request(
            "POST",
            route=f"/v2/chat_rooms/{chat_room_id}/invite",
            payload={"with_user_ids": user_ids},
        )

    def kick_users_from_chat(self, chat_room_id: int, user_ids: list[int]):
        return self.__base._request(
            "POST",
            route=f"/v2/chat_rooms/{chat_room_id}/kick",
            payload={"with_user_ids": user_ids},
        )

    def pin_chat(self, room_id: int):
        return self.__base._request("POST", route=f"/v1/chat_rooms/{room_id}/pinned")

    def read_message(self, chat_room_id: int, message_id: int):
        return self.__base._request(
            "POST",
            route=f"/v2/chat_rooms/{chat_room_id}/messages/{message_id}/read",
        )

    def refresh_chat_rooms(self, from_time: int = None) -> ChatRoomsResponse:
        params = {}
        if from_time:
            params["from_time"] = from_time
        return self.__base._request(
            "GET",
            route=f"/v2/chat_rooms/update",
            params=params,
            data_type=ChatRoomsResponse,
        )

    def remove_chat_rooms(self, chat_room_ids: list[int]):
        chat_room_ids = (
            [chat_room_ids] if isinstance(chat_room_ids, int) else chat_room_ids
        )
        return self.__base._request(
            "POST",
            route=f"/v1/chat_rooms/mass_destroy",
            payload={"chat_room_ids": chat_room_ids},
        )

    def report_chat_room(
        self,
        chat_room_id: int,
        opponent_id: int,
        category_id: int,
        reason: str = None,
        screenshot_filename: str = None,
        screenshot_2_filename: str = None,
        screenshot_3_filename: str = None,
        screenshot_4_filename: str = None,
    ):
        return self.__base._request(
            "POST",
            route=f"/v3/chat_rooms/{chat_room_id}/report",
            payload={
                "chat_room_id": chat_room_id,
                "opponent_id": opponent_id,
                "category_id": category_id,
                "reason": reason,
                "screenshot_filename": screenshot_filename,
                "screenshot_2_filename": screenshot_2_filename,
                "screenshot_3_filename": screenshot_3_filename,
                "screenshot_4_filename": screenshot_4_filename,
            },
        )

    def send_message(self, chat_room_id: int, **params) -> MessageResponse:
        return self.__base._request(
            "POST",
            route=f"/v3/chat_rooms/{chat_room_id}/messages/new",
            payload=params,
            data_type=MessageResponse,
            bypass_delay=True,
        )

    def unhide_chat(self, chat_room_ids: int):
        return self.__base._request(
            "DELETE",
            route=f"/v1/hidden/chats",
            params={"chat_room_ids": chat_room_ids},
        )

    def unpin_chat(self, chat_room_id: int):
        return self.__base._request(
            "DELETE", route=f"/v1/chat_rooms/{chat_room_id}/pinned"
        )
