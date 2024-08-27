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

from typing import List

from .. import config
from ..responses import (
    BgmsResponse,
    CallActionSignatureResponse,
    CallStatusResponse,
    ConferenceCallResponse,
    GamesResponse,
    GenresResponse,
    PostResponse,
    PostsResponse,
    Response,
    UsersByTimestampResponse,
)


class CallApi:
    """通話 API"""

    def __init__(self, client) -> None:
        # pylint: disable=import-outside-toplevel
        from ..client import Client

        self.__client: Client = client

    async def get_user_active_call(self, user_id: int) -> PostResponse:
        """ユーザーが参加中の通話を取得する

        Args:
            user_id (int):

        Returns:
            PostResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/posts/active_call",
            params={"user_id": user_id},
            return_type=PostResponse,
        )

    async def get_bgms(self) -> BgmsResponse:
        """通話のBGMを取得する

        Returns:
            BgmsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/calls/bgm",
            return_type=BgmsResponse,
        )

    async def get_call(self, call_id: int) -> ConferenceCallResponse:
        """通話を取得する

        Args:
            call_id (int):

        Returns:
            ConferenceCallResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v1/calls/conferences/{call_id}",
            return_type=ConferenceCallResponse,
        )

    async def get_call_invitable_users(
        self, call_id: int, **params
    ) -> UsersByTimestampResponse:
        """通話に招待可能なユーザーを取得する

        Args:
            call_id (int):
            from_timestamp (int, optional):

        Returns:
            UsersByTimestampResponse:
        """
        # @Nullable @Query("user[nickname]")
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v1/calls/{call_id}/users/invitable",
            params=params,
            return_type=UsersByTimestampResponse,
        )

    async def get_call_status(self, opponent_id: int) -> CallStatusResponse:
        """通話の状態を取得します

        Args:
            opponent_id (int):

        Returns:
            CallStatusResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v1/calls/phone_status/{opponent_id}",
            return_type=CallStatusResponse,
        )

    async def get_games(self, **params) -> GamesResponse:
        """通話に設定可能なゲームを取得する

        Args:
            number (int, optional):
            ids (List[int], optional):
            from_id (int, optional):

        Returns:
            GamesResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/games/apps",
            params=params,
            return_type=GamesResponse,
        )

    async def get_genres(self, **params) -> GenresResponse:
        """通話のジャンルを取得する

        Args:
            number (int, optional):
            from (int, optional):

        Returns:
            GenresResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/genres",
            params=params,
            return_type=GenresResponse,
        )

    async def get_group_calls(self, **params) -> PostsResponse:
        """サークル内の通話を取得する

        Args:
            number (int, optional):
            group_category_id (int, optional):
            from_timestamp (int, optional):
            scope (str, optional):

        Returns:
            PostsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/posts/group_calls",
            params=params,
            return_type=PostsResponse,
        )

    async def invite_online_followings_to_call(
        self, call_id: int, **params
    ) -> Response:
        """オンラインの友達をまとめて通話に招待します

        Args:
            call_id (int, optional):
            group_id (str, optional):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/calls/{call_id}/bulk_invite",
            params=params,
            return_type=Response,
        )

    async def invite_users_to_call(self, call_id: int, user_ids: List[int]) -> Response:
        """通話に複数のユーザーを招待する

        Args:
            call_id (int):
            user_ids (List[int]):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/calls/conference_calls/{call_id}/invite",
            payload={"call_id": call_id, "user_ids": user_ids},
            return_type=Response,
        )

    async def invite_users_to_chat_call(self, **params) -> Response:
        """ユーザーをチャット通話に招待する

        Args:
            chat_room_id (int):
            room_id (int):
            room_url (str):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v2/calls/invite",
            payload=params,
            return_type=Response,
        )

    async def kick_user_from_call(
        self, call_id: int, **params
    ) -> CallActionSignatureResponse:
        """ユーザーを通話からキックする

        Args:
            call_id (int):
            uuid (int):
            ban (bool):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v3/calls/conference_calls/{call_id}/kick",
            payload=params,
            return_type=CallActionSignatureResponse,
        )

    async def start_call(self, call_id: int, **params) -> Response:
        """通話を開始する

        Args:
            call_id (int):
            joinable_by (str):
            game_title (str, optional):
            category_id (str, optional):

        Returns:
            Response:
        """
        return await self.__client.request(
            "PUT",
            config.API_HOST + f"/v1/calls/{call_id}",
            payload=params,
            return_type=Response,
        )

    async def set_user_role(self, call_id: int, user_id: int, role: str) -> Response:
        """通話の参加者に役職を付与する

        Args:
            call_id (int):
            user_id (int):
            role (str):

        Returns:
            Response:
        """
        return await self.__client.request(
            "PUT",
            config.API_HOST + f"/v1/calls/{call_id}/users/{user_id}",
            payload={"role": role},
            return_type=Response,
        )

    async def join_call(self, **params) -> ConferenceCallResponse:
        """通話に参加する

        Args:
            conference_id (int):
            call_sid (str, optional):

        Returns:
            ConferenceCallResponse:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/calls/start_conference_call",
            payload=params,
            return_type=ConferenceCallResponse,
        )

    async def leave_call(self, **params) -> Response:
        """通話から退出する

        Args:
            conference_id (int):
            call_sid (str, optional):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/calls/leave_conference_call",
            payload=params,
            return_type=Response,
        )

    async def join_call_as_anonymous(self, **params) -> ConferenceCallResponse:
        """匿名で通話に参加する

        Args:
            conference_id (int):
            agora_uid (str):

        Returns:
            ConferenceCallResponse:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/anonymous_calls/start_conference_call",
            payload=params,
            return_type=ConferenceCallResponse,
        )

    async def leave_call_as_anonymous(self, **params) -> Response:
        """匿名で参加した通話を退出する

        Args:
            conference_id (int):
            agora_uid (str, optional):

        Returns:
            Response: _description_
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/anonymous_calls/leave_conference_call",
            payload=params,
            return_type=Response,
        )
