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

from .. import config
from ..responses import (
    ChatRoomResponse,
    ChatRoomsResponse,
    CreateChatRoomResponse,
    FollowUsersResponse,
    GifsDataResponse,
    MessageResponse,
    MessagesResponse,
    Response,
    StickerPacksResponse,
    TotalChatRequestResponse,
    UnreadStatusResponse,
)


class ChatApi:
    """チャット API"""

    def __init__(self, client) -> None:
        # pylint: disable=import-outside-toplevel
        from ..client import Client

        self.__client: Client = client

    async def accept_chat_requests(self, **params) -> Response:
        """チャットリクエストを承認する

        Args:
            chat_room_ids (List[int]):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/chat_rooms/accept_chat_request",
            json=params,
            return_type=Response,
        )

    async def check_unread_status(self, **params) -> UnreadStatusResponse:
        """チャットの未読ステータスを確認する

        Args:
            from_time (int):

        Returns:
            UnreadStatusResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/chat_rooms/unread_status",
            params=params,
            return_type=UnreadStatusResponse,
        )

    async def create_group_chat(self, **params) -> CreateChatRoomResponse:
        """グループチャットを作成する

        Args:
            name (str):
            with_user_ids (List[int]):
            icon_filename (str, optional):
            background_filename (str, optional):

        Returns:
            CreateChatRoomResponse:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v3/chat_rooms/new",
            json=params,
            return_type=CreateChatRoomResponse,
        )

    async def create_private_chat(self, **params) -> CreateChatRoomResponse:
        """個人チャットを作成する

        Args:
            with_user_id (int):
            matching_id (int, optional):
            hima_chat (bool, optional):

        Returns:
            CreateChatRoomResponse:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/chat_rooms/new",
            json=params,
            return_type=CreateChatRoomResponse,
        )

    async def delete_chat_background(self, room_id: int) -> Response:
        """チャットの背景を削除する

        Args:
            room_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v2/chat_rooms/{room_id}/background",
            return_type=Response,
        )

    async def delete_message(self, room_id: int, message_id: int) -> Response:
        """チャットメッセージを削除する

        Args:
            room_id (int):
            message_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v1/chat_rooms/{room_id}/messages/{message_id}/delete",
            return_type=Response,
        )

    async def edit_chat_room(self, chat_room_id: int, **params) -> Response:
        """チャットルームを編集する

        Args:
            chat_room_id (int):
            name (str):
            icon_filename (str, optional):
            background_filename (str, optional):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/chat_rooms/{chat_room_id}/edit",
            json=params,
            return_type=Response,
        )

    async def get_chatable_users(
        self,
        # @Body @Nullable SearchUsersRequest searchUsersRequest
        **params,
    ) -> FollowUsersResponse:
        """チャット可能なユーザーを取得する

        Args:
            from_follow_id (int, optional):
            from_timestamp (int, optional):
            order_by (str, optional):

        Returns:
            FollowUsersResponse:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/users/followings/chatable",
            json=params,
            return_type=FollowUsersResponse,
        )

    async def get_gifs_data(self) -> GifsDataResponse:
        """チャット用 GIF データを取得する

        Returns:
            GifsDataResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/hidden/chats",
            return_type=GifsDataResponse,
        )

    async def get_hidden_chat_rooms(self, **params) -> ChatRoomsResponse:
        """非表示に設定したチャットルームを取得する

        Args:
            from_timestamp (int, optional):
            number (int, optional)

        Returns:
            ChatRoomsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/hidden/chats",
            params=params,
            return_type=ChatRoomsResponse,
        )

    async def get_main_chat_rooms(self, **params) -> ChatRoomsResponse:
        """メインのチャットルームを取得する

        Args:
            from_timestamp (int, optional):

        Returns:
            ChatRoomsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/chat_rooms/main_list",
            params=params,
            return_type=ChatRoomsResponse,
        )

    async def get_messages(self, chat_room_id: int, **params) -> MessagesResponse:
        """メッセージを取得する

        Args:
            from_message_id (int, optional):
            to_message_id (int, optional):

        Returns:
            MessagesResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v2/chat_rooms/{chat_room_id}/messages",
            params=params,
            return_type=MessagesResponse,
        )

    async def get_chat_requests(self, **params) -> ChatRoomsResponse:
        """チャットリクエストを取得する

        Args:
            number (int, optional):
            from_timestamp (int, optional):

        Returns:
            ChatRoomsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/chat_rooms/request_list",
            params=params,
            return_type=ChatRoomsResponse,
        )

    async def get_chat_room(self, chat_room_id: int) -> ChatRoomResponse:
        """チャットルームを取得する

        Args:
            chat_room_id (int):

        Returns:
            ChatRoomResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v2/chat_rooms/{chat_room_id}",
            return_type=ChatRoomResponse,
        )

    async def get_sticker_packs(self) -> StickerPacksResponse:
        """チャット用のスタンプを取得する

        Returns:
            StickerPacksResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/sticker_packs",
            return_type=StickerPacksResponse,
        )

    async def get_total_chat_requests(self) -> TotalChatRequestResponse:
        """チャットリクエストの総数を取得する

        Returns:
            TotalChatRequestResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/chat_rooms/total_chat_request",
            return_type=TotalChatRequestResponse,
        )

    async def hide_chat(self, chat_room_id: int) -> Response:
        """チャットルームを非表示にする

        Args:
            chat_room_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/hidden/chats",
            json={"chat_room_id": chat_room_id},
            return_type=Response,
        )

    async def invite_to_chat(self, chat_room_id: int, **params) -> Response:
        """チャットルームにユーザーを招待する

        Args:
            chat_room_id (int):
            with_user_ids (List[int]):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v2/chat_rooms/{chat_room_id}/invite",
            json=params,
            return_type=Response,
        )

    async def kick_users_from_chat(self, chat_room_id: int, **params) -> Response:
        """チャットルームからユーザーを追放する

        Args:
            chat_room_id (int):
            with_user_ids (List[int]):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v2/chat_rooms/{chat_room_id}/kick",
            json=params,
            return_type=Response,
        )

    async def pin_chat(self, room_id: int) -> Response:
        """チャットルームをピン留めする

        Args:
            room_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/chat_rooms/{room_id}/pinned",
            return_type=Response,
        )

    async def read_message(self, chat_room_id: int, message_id: int) -> Response:
        """メッセージを既読にする

        Args:
            chat_room_id (int):
            message_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST
            + f"/v2/chat_rooms/{chat_room_id}/messages/{message_id}/read",
            return_type=Response,
        )

    async def refresh_chat_rooms(self, **params) -> ChatRoomsResponse:
        """チャットルームを更新する

        Args:
            from_time (int, optional):

        Returns:
            ChatRoomsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/chat_rooms/update",
            params=params,
            return_type=ChatRoomsResponse,
        )

    async def delete_chat_rooms(self, **params) -> Response:
        """チャットルームを削除する

        Args:
            chat_room_ids (List[int]):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/chat_rooms/mass_destroy",
            json=params,
            return_type=Response,
        )

    async def send_message(self, chat_room_id: int, **params) -> MessageResponse:
        """チャットを送信する

        Args:
            chat_room_id (int):

        Returns:
            MessageResponse:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v3/chat_rooms/{chat_room_id}/messages/new",
            json=params,
            return_type=MessageResponse,
        )

    async def unhide_chat(self, **params) -> Response:
        """非表示に設定したチャットルームを表示する

        Args:
            chat_room_ids (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "DELETE",
            config.API_HOST + "/v1/hidden/chats",
            params=params,
            return_type=Response,
        )

    async def unpin_chat(self, chat_room_id: int) -> Response:
        """チャットのピン留めを解除する

        Args:
            chat_room_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v1/chat_rooms/{chat_room_id}/pinned",
            return_type=Response,
        )
