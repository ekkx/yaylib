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

from .. import config
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
    """チャット API"""

    def __init__(self, client) -> None:
        self.__client = client

    async def accept_chat_requests(self, chat_room_ids: list[int]):
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/chat_rooms/accept_chat_request",
            json={"chat_room_ids": chat_room_ids},
        )

    async def check_unread_status(self, from_time: int) -> UnreadStatusResponse:
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/chat_rooms/unread_status",
            params={"from_time": from_time},
            return_type=UnreadStatusResponse,
        )

    async def create_group_chat(
        self,
        name: str,
        with_user_ids: list[int],
        icon_filename: str = None,
        background_filename: str = None,
    ) -> CreateChatRoomResponse:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v3/chat_rooms/new",
            json={
                "name": name,
                "with_user_ids": with_user_ids,
                "icon_filename": icon_filename,
                "background_filename": background_filename,
            },
            return_type=CreateChatRoomResponse,
        )

    async def create_private_chat(
        self, with_user_id: int, matching_id: int = None, hima_chat: bool = False
    ) -> CreateChatRoomResponse:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/chat_rooms/new",
            json={
                "with_user_id": with_user_id,
                "matching_id": matching_id,
                "hima_chat": hima_chat,
            },
            return_type=CreateChatRoomResponse,
        )

    async def delete_background(self, room_id: int):
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v2/chat_rooms/{room_id}/background",
        )

    async def delete_message(self, room_id: int, message_id: int):
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v1/chat_rooms/{room_id}/messages/{message_id}/delete",
        )

    async def edit_chat_room(
        self,
        chat_room_id: int,
        name: str,
        icon_filename: str = None,
        background_filename: str = None,
    ):
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/chat_rooms/{chat_room_id}/edit",
            json={
                "name": name,
                "icon_filename": icon_filename,
                "background_filename": background_filename,
            },
        )

    async def get_chatable_users(
        self,
        # @Body @Nullable SearchUsersRequest searchUsersRequest
        from_follow_id: int = None,
        from_timestamp: int = None,
        order_by: str = None,
    ) -> FollowUsersResponse:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/users/followings/chatable",
            json={
                "from_follow_id": from_follow_id,
                "from_timestamp": from_timestamp,
                "order_by": order_by,
            },
            return_type=FollowUsersResponse,
        )

    async def get_gifs_data(self) -> GifsDataResponse:
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/hidden/chats",
            return_type=GifsDataResponse,
        )

    async def get_hidden_chat_rooms(self, **params) -> ChatRoomsResponse:
        """

        Parameters:
        ---------------

            - from_timestamp: int - (optional)
            - number: int - (optional)

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/hidden/chats",
            params=params,
            return_type=ChatRoomsResponse,
        )

    async def get_main_chat_rooms(
        self, from_timestamp: int = None
    ) -> ChatRoomsResponse:
        params = {}
        if from_timestamp:
            params["from_timestamp"] = from_timestamp
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/chat_rooms/main_list",
            params=params,
            return_type=ChatRoomsResponse,
        )

    async def get_messages(self, chat_room_id: int, **params) -> MessagesResponse:
        """

        Parameters:
        ---------------
            - from_message_id: int - (optional)
            - to_message_id: int - (optional)

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v2/chat_rooms/{chat_room_id}/messages",
            params=params,
            return_type=MessagesResponse,
        )

    async def get_request_chat_rooms(self, **params) -> ChatRoomsResponse:
        """

        Parameters:
        -----------

            - number: int (optional)
            - from_timestamp: int (optional)

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/chat_rooms/request_list",
            params=params,
            return_type=ChatRoomsResponse,
        )

    async def get_chat_room(self, chat_room_id: int) -> ChatRoomResponse:
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v2/chat_rooms/{chat_room_id}",
            return_type=ChatRoomResponse,
        )

    async def get_sticker_packs(self) -> StickerPacksResponse:
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/sticker_packs",
            return_type=StickerPacksResponse,
        )

    async def get_total_chat_requests(self) -> TotalChatRequestResponse:
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/chat_rooms/total_chat_request",
            return_type=TotalChatRequestResponse,
        )

    async def hide_chat(self, chat_room_id: int):
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/hidden/chats",
            json={"chat_room_id": chat_room_id},
        )

    async def invite_to_chat(self, chat_room_id: int, user_ids: list[int]):
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v2/chat_rooms/{chat_room_id}/invite",
            json={"with_user_ids": user_ids},
        )

    async def kick_users_from_chat(self, chat_room_id: int, user_ids: list[int]):
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v2/chat_rooms/{chat_room_id}/kick",
            json={"with_user_ids": user_ids},
        )

    async def pin_chat(self, room_id: int):
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/chat_rooms/{room_id}/pinned",
        )

    async def read_message(self, chat_room_id: int, message_id: int):
        return await self.__client.request(
            "POST",
            config.API_HOST
            + f"/v2/chat_rooms/{chat_room_id}/messages/{message_id}/read",
        )

    async def refresh_chat_rooms(self, from_time: int = None) -> ChatRoomsResponse:
        params = {}
        if from_time:
            params["from_time"] = from_time
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/chat_rooms/update",
            params=params,
            return_type=ChatRoomsResponse,
        )

    async def remove_chat_rooms(self, chat_room_ids: list[int]):
        chat_room_ids = (
            [chat_room_ids] if isinstance(chat_room_ids, int) else chat_room_ids
        )
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/chat_rooms/mass_destroy",
            json={"chat_room_ids": chat_room_ids},
        )

    async def report_chat_room(
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
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v3/chat_rooms/{chat_room_id}/report",
            json={
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

    async def send_message(self, chat_room_id: int, **params) -> MessageResponse:
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v3/chat_rooms/{chat_room_id}/messages/new",
            json=params,
            return_type=MessageResponse,
        )

    async def unhide_chat(self, chat_room_ids: int):
        return await self.__client.request(
            "DELETE",
            config.API_HOST + "/v1/hidden/chats",
            params={"chat_room_ids": chat_room_ids},
        )

    async def unpin_chat(self, chat_room_id: int):
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v1/chat_rooms/{chat_room_id}/pinned",
        )
