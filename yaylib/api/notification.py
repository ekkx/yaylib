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
from ..config import Configs
from ..responses import ActivitiesResponse


class NotificationAPI(object):
    def __init__(self, base: client.BaseClient) -> None:
        self.__base = base

    def get_user_activities(self, **params) -> ActivitiesResponse:
        """

        Parameters
        ----------
            - important: bool - (required)
            - from_timestamp: int - (optional)
            - number: int - (optional)

        """
        return self.__base._request(
            "GET",
            base_url=Configs.STAGING_HOST_2,
            route="/api/user_activities",
            params=params,
            data_type=ActivitiesResponse,
        )

    def get_user_merged_activities(self, **params) -> ActivitiesResponse:
        """
        Parameters
        ----------

            - from_timestamp: int - (optional)
            - number: int - (optional)

        """
        return self.__base._request(
            "GET",
            base_url=Configs.STAGING_HOST_2,
            route="/api/v2/user_activities",
            params=params,
            data_type=ActivitiesResponse,
        )

    def received_notification(self, pid: str, type: str, opened_at: int = None):
        # TODO: opened_atはnullalbeか確認する
        return self.__base._request(
            "POST",
            route="/api/received_push_notifications",
            payload={"pid": pid, "type": type, "opened_at": opened_at},
        )
