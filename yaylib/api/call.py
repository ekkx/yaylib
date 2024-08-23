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

from typing import Optional

from .. import config
from ..responses import (
    BgmsResponse,
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
        from ..client import Client  # pylint: disable=import-outside-toplevel

        self.__client: Client = client

    async def bump_call(
        self, call_id: int, participant_limit: Optional[int] = None
    ) -> Response:
        params = {}
        if participant_limit:
            params["participant_limit"] = participant_limit
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/calls/{call_id}/bump",
            params=params,
            return_type=Response,
        )

    async def get_user_active_call(self, user_id: int) -> PostResponse:
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/posts/active_call",
            params={"user_id": user_id},
            return_type=PostResponse,
        )

    async def get_bgms(self) -> BgmsResponse:
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/calls/bgm",
            return_type=BgmsResponse,
        )

    async def get_call(self, call_id: int) -> ConferenceCallResponse:
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v1/calls/conferences/{call_id}",
            return_type=ConferenceCallResponse,
        )

    async def get_call_invitable_users(
        self, call_id: int, from_timestamp: Optional[int] = None
    ) -> UsersByTimestampResponse:
        # @Nullable @Query("user[nickname]")
        params = {}
        if from_timestamp:
            params["from_timestamp"] = from_timestamp
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v1/calls/{call_id}/users/invitable",
            params=params,
            return_type=UsersByTimestampResponse,
        )

    async def get_call_status(self, opponent_id: int) -> CallStatusResponse:
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v1/calls/phone_status/{opponent_id}",
            return_type=CallStatusResponse,
        )

    async def get_games(self, **params) -> GamesResponse:
        """

        Parameters
        ----------
            - number: int - (optional)
            - ids: list[int] - (optional)
            - from_id: int - (optional)

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
            number (Optional[int]):
            from (Optional[int]):

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
            number (Optional[int]):
            group_category_id (Optional[int]):
            from_timestamp (Optional[int]):
            scope (Optional[str]):

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
        self, call_id: int, group_id: Optional[int] = None
    ) -> Response:
        """通話に一括で招待します

        Args:
            call_id (Optional[int]):
            group_id (Optional[str]):
        """
        params = {}
        if group_id:
            params["group_id"] = group_id
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/calls/{call_id}/bulk_invite",
            params=params,
            return_type=Response,
        )

    async def invite_users_to_call(self, call_id: int, user_ids: list[int]) -> Response:
        """通話に複数のユーザーを招待します

        Args:
            call_id (int):
            user_ids (List[int]):
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/calls/conference_calls/{call_id}/invite",
            payload={"call_id": call_id, "user_ids": user_ids},
            return_type=Response,
        )

    async def invite_users_to_chat_call(
        self, chat_room_id: int, room_id: int, room_url: str
    ) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v2/calls/invite",
            payload={
                "chat_room_id": chat_room_id,
                "room_id": room_id,
                "room_url": room_url,
            },
            return_type=Response,
        )

    async def kick_user_from_call(self, call_id: int, user_id: int) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/calls/conference_calls/{call_id}/kick",
            payload={"user_id": user_id},
            return_type=Response,
        )

    async def set_call(
        self,
        call_id: int,
        joinable_by: str,
        game_title: Optional[str] = None,
        category_id: Optional[str] = None,
    ) -> Response:
        return await self.__client.request(
            "PUT",
            config.API_HOST + f"/v1/calls/{call_id}",
            payload={
                "joinable_by": joinable_by,
                "game_title": game_title,
                "category_id": category_id,
            },
            return_type=Response,
        )

    async def set_user_role(self, call_id: int, user_id: int, role: str) -> Response:
        return await self.__client.request(
            "PUT",
            config.API_HOST + f"/v1/calls/{call_id}/users/{user_id}",
            payload={"role": role},
            return_type=Response,
        )

    async def join_call(
        self, conference_id: int, call_sid: Optional[str] = None
    ) -> ConferenceCallResponse:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/calls/start_conference_call",
            payload={"conference_id": conference_id, "call_sid": call_sid},
            return_type=ConferenceCallResponse,
        )

    async def leave_call(
        self, conference_id: int, call_sid: Optional[str] = None
    ) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/calls/leave_conference_call",
            payload={"conference_id": conference_id, "call_sid": call_sid},
            return_type=Response,
        )

    async def join_call_as_anonymous(
        self, conference_id: int, agora_uid: str
    ) -> ConferenceCallResponse:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/anonymous_calls/start_conference_call",
            payload={"conference_id": conference_id, "agora_uid": agora_uid},
            return_type=ConferenceCallResponse,
        )

    async def leave_call_as_anonymous(
        self, conference_id: int, agora_uid: Optional[str] = None
    ) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/anonymous_calls/leave_conference_call",
            payload={"conference_id": conference_id, "agora_uid": agora_uid},
            return_type=Response,
        )
