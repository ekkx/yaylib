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
    BgmsResponse,
    CallStatusResponse,
    ConferenceCallResponse,
    GamesResponse,
    GenresResponse,
    PostResponse,
    PostsResponse,
    UsersByTimestampResponse,
)


class CallAPI(object):
    def __init__(self, base: client.BaseClient) -> None:
        self.__base = base

    def bump_call(self, call_id: int, participant_limit: int = None):
        params = {}
        if participant_limit:
            params["participant_limit"] = participant_limit
        return self.__base._request(
            "POST", route=f"/v1/calls/{call_id}/bump", params=params
        )

    def get_user_active_call(self, user_id: int) -> PostResponse:
        return self.__base._request(
            "GET",
            route=f"/v1/posts/active_call",
            params={"user_id": user_id},
            data_type=PostResponse,
        )

    def get_bgms(self) -> BgmsResponse:
        return self.__base._request(
            "GET", route="/v1/calls/bgm", data_type=BgmsResponse
        )

    def get_call(self, call_id: int) -> ConferenceCallResponse:
        return self.__base._request(
            "GET",
            route=f"/v1/calls/conferences/{call_id}",
            data_type=ConferenceCallResponse,
        )

    def get_call_invitable_users(
        self, call_id: int, from_timestamp: int = None
    ) -> UsersByTimestampResponse:
        # @Nullable @Query("user[nickname]")
        params = {}
        if from_timestamp:
            params["from_timestamp"] = from_timestamp
        return self.__base._request(
            "GET",
            route=f"/v1/calls/{call_id}/users/invitable",
            params=params,
            data_type=UsersByTimestampResponse,
        )

    def get_call_status(self, opponent_id: int) -> CallStatusResponse:
        return self.__base._request(
            "GET",
            route=f"/v1/calls/phone_status/{opponent_id}",
            data_type=CallStatusResponse,
        )

    def get_games(self, **params) -> GamesResponse:
        """

        Parameters
        ----------
            - number: int - (optional)
            - ids: list[int] - (optional)
            - from_id: int - (optional)

        """
        return self.__base._request(
            "GET",
            route=f"/v1/games/apps",
            params=params,
            data_type=GamesResponse,
        )

    def get_genres(self, **params) -> GenresResponse:
        """

        Parameters
        ----------
            - number: int - (optional)
            - from: int - (optional)

        """
        return self.__base._request(
            "GET",
            route=f"/v1/genres",
            params=params,
            data_type=GenresResponse,
        )

    def get_group_calls(self, **params) -> PostsResponse:
        """

        Parameters
        ----------
            - number: int - (optional)
            - group_category_id: int - (optional)
            - from_timestamp: int - (optional)
            - scope: str - (optional)

        """
        return self.__base._request(
            "GET",
            route="/v1/posts/group_calls",
            params=params,
            data_type=PostsResponse,
        )

    def invite_to_call_bulk(self, call_id: int, group_id: int = None):
        """

        Parameters
        ----------
            - call_id: int - (required)
            - group_id: int - (optional)

        """
        params = {}
        if group_id:
            params["group_id"] = group_id
        return self.__base._request(
            "POST",
            route=f"/v1/calls/{call_id}/bulk_invite",
            params=params,
        )

    def invite_users_to_call(self, call_id: int, user_ids: list[int]):
        """

        Parameters
        ----------
            - call_id: int - (required)
            - user_ids: list[int] - (required)

        """
        return self.__base._request(
            "POST",
            route=f"/v1/calls/conference_calls/{call_id}/invite",
            payload={"call_id": call_id, "user_ids": user_ids},
        )

    def invite_users_to_chat_call(self, chat_room_id: int, room_id: int, room_url: str):
        return self.__base._request(
            "POST",
            route="/v2/calls/invite",
            payload={
                "chat_room_id": chat_room_id,
                "room_id": room_id,
                "room_url": room_url,
            },
        )

    def kick_and_ban_from_call(self, call_id: int, user_id: int):
        return self.__base._request(
            "POST",
            route=f"/v1/calls/conference_calls/{call_id}/kick",
            payload={"user_id": user_id},
        )

    def set_call(
        self,
        call_id: int,
        joinable_by: str,
        game_title: str = None,
        category_id: str = None,
    ):
        return self.__base._request(
            "PUT",
            route=f"/v1/calls/{call_id}",
            payload={
                "joinable_by": joinable_by,
                "game_title": game_title,
                "category_id": category_id,
            },
        )

    def set_user_role(self, call_id: int, user_id: int, role: str):
        return self.__base._request(
            "PUT",
            route=f"/v1/calls/{call_id}/users/{user_id}",
            payload={"role": role},
        )

    def start_call(
        self, conference_id: int, call_sid: str = None
    ) -> ConferenceCallResponse:
        return self.__base._request(
            "POST",
            route="/v1/calls/start_conference_call",
            payload={"conference_id": conference_id, "call_sid": call_sid},
            data_type=ConferenceCallResponse,
        )

    def start_anonymous_call(
        self, conference_id: int, agora_uid: str
    ) -> ConferenceCallResponse:
        return self.__base._request(
            "POST",
            route="/v1/anonymous_calls/start_conference_call",
            payload={"conference_id": conference_id, "agora_uid": agora_uid},
            data_type=ConferenceCallResponse,
        )

    def stop_anonymous_call(self, conference_id: int, agora_uid: str = None):
        return self.__base._request(
            "POST",
            route="/v1/anonymous_calls/leave_conference_call",
            payload={"conference_id": conference_id, "agora_uid": agora_uid},
        )

    def stop_call(self, conference_id: int, call_sid: str = None):
        return self.__base._request(
            "POST",
            route="/v1/calls/leave_conference_call",
            payload={"conference_id": conference_id, "call_sid": call_sid},
        )
