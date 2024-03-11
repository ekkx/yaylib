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
from ..models import ThreadInfo
from ..responses import GroupThreadListResponse, PostsResponse


class ThreadAPI(object):
    def __init__(self, base: client.BaseClient) -> None:
        self.__base = base

    def add_post_to_thread(self, post_id: int, thread_id: int) -> ThreadInfo:
        return self.__base._request(
            "PUT",
            route=f"/v3/posts/{post_id}/move_to_thread/{thread_id}",
            data_type=ThreadInfo,
        )

    def convert_post_to_thread(
        self, post_id: int, title: str = None, thread_icon_filename: str = None
    ) -> ThreadInfo:
        return self.__base._request(
            "POST",
            route=f"/v3/posts/{post_id}/move_to_thread",
            payload={"title": title, "thread_icon_filename": thread_icon_filename},
            data_type=ThreadInfo,
        )

    def create_thread(
        self, group_id: int, title: str, thread_icon_filename: str
    ) -> ThreadInfo:
        return self.__base._request(
            "POST",
            route=f"/v1/threads/",
            payload={
                "group_id": group_id,
                "title": title,
                "thread_icon_filename": thread_icon_filename,
            },
            data_type=ThreadInfo,
        )

    def get_group_thread_list(
        self, group_id: int, from_str: str = None, **params
    ) -> GroupThreadListResponse:
        """

        Parameters:
        -----------

            - group_id: int
            - from_str: str = None
            - join_status: str = None

        """
        params["group_id"] = group_id
        if from_str:
            params["from"] = from_str
        return self.__base._request(
            "GET",
            route=f"/v1/threads/",
            params=params,
            data_type=GroupThreadListResponse,
        )

    def get_thread_joined_statuses(self, ids: list[int]) -> dict:
        return self.__base._request(
            "GET",
            route=f"/v1/threads/joined_statuses",
            params={"ids[]": ids},
        )

    def get_thread_posts(
        self, thread_id: int, from_str: str = None, **params
    ) -> PostsResponse:
        """

        Parameters:
        -----------

            - post_type: str
            - number: int = None
            - from_str: str = None

        """
        if from_str:
            params["from"] = from_str
        return self.__base._request(
            "GET",
            route=f"/v1/threads/{thread_id}/posts",
            params=params,
            data_type=PostsResponse,
        )

    def join_thread(self, thread_id: int, user_id: int):
        return self.__base._request(
            "POST", route=f"/v1/threads/{thread_id}/members/{user_id}"
        )

    def leave_thread(self, thread_id: int, user_id: int):
        return self.__base._request(
            "DELETE",
            route=f"/v1/threads/{thread_id}/members/{user_id}",
        )

    def remove_thread(
        self,
        thread_id: int,
    ):
        return self.__base._request(
            "DELETE",
            route=f"/v1/threads/{thread_id}",
        )

    def update_thread(self, thread_id: int, title: str, thread_icon_filename: str):
        return self.__base._request(
            "PUT",
            route=f"/v1/threads/{thread_id}",
            payload={"title": title, "thread_icon_filename": thread_icon_filename},
        )
