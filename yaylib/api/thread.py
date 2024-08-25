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
from ..models import ThreadInfo
from ..responses import GroupThreadListResponse, PostsResponse, Response


class ThreadApi:
    """スレッド API"""

    def __init__(self, client) -> None:
        # pylint: disable=import-outside-toplevel
        from ..client import Client

        self.__client: Client = client

    async def add_post_to_thread(self, post_id: int, thread_id: int) -> ThreadInfo:
        """投稿をスレッドに追加する

        Args:
            post_id (int):
            thread_id (int):

        Returns:
            ThreadInfo:
        """
        return await self.__client.request(
            "PUT",
            config.API_HOST + f"/v3/posts/{post_id}/move_to_thread/{thread_id}",
            return_type=ThreadInfo,
        )

    async def convert_post_to_thread(self, post_id: int, **params) -> ThreadInfo:
        """投稿をスレッドに変換する

        Args:
            post_id (int):
            title (str, optional):
            thread_icon_filename (str, optional):

        Returns:
            ThreadInfo:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v3/posts/{post_id}/move_to_thread",
            json=params,
            return_type=ThreadInfo,
        )

    async def create_thread(
        self, group_id: int, title: str, thread_icon_filename: str
    ) -> ThreadInfo:
        """スレッドを作成する

        Args:
            group_id (int):
            title (str):
            thread_icon_filename (str):

        Returns:
            ThreadInfo:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/threads/",
            json={
                "group_id": group_id,
                "title": title,
                "thread_icon_filename": thread_icon_filename,
            },
            return_type=ThreadInfo,
        )

    async def get_group_thread_list(self, **params) -> GroupThreadListResponse:
        """サークルのスレッド一覧を取得する

        Args:
            group_id (int):
            from_str (str, optional):
            join_status (str, optional):

        Returns:
            GroupThreadListResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/threads/",
            params=params,
            return_type=GroupThreadListResponse,
        )

    async def get_thread_joined_statuses(self, ids: List[int]) -> Response:
        """スレッド参加ステータスを取得する

        Args:
            ids (List[int]):

        Returns:
            Response:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/threads/joined_statuses",
            params={"ids[]": ids},
            return_type=Response,
        )

    async def get_thread_posts(self, thread_id: int, **params) -> PostsResponse:
        """スレッド内のタイムラインを取得する

        Args:
            thread_id (int):
            from_str (str, optional):
            number (str, optional):

        Returns:
            PostsResponse:
        """
        params["from"] = params.get("from_str")
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v1/threads/{thread_id}/posts",
            params=params,
            return_type=PostsResponse,
        )

    async def join_thread(self, thread_id: int, user_id: int) -> Response:
        """スレッドに参加する

        Args:
            thread_id (int):
            user_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/threads/{thread_id}/members/{user_id}",
            return_type=Response,
        )

    async def leave_thread(self, thread_id: int, user_id: int) -> Response:
        """スレッドから脱退する

        Args:
            thread_id (int):
            user_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v1/threads/{thread_id}/members/{user_id}",
            return_type=Response,
        )

    async def delete_thread(
        self,
        thread_id: int,
    ) -> Response:
        """スレッドを削除する

        Args:
            thread_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v1/threads/{thread_id}",
            return_type=Response,
        )

    async def update_thread(self, thread_id: int, **params) -> Response:
        """スレッドをアップデートする

        Args:
            thread_id (int):
            title (str):
            thread_icon_filename (str):

        Returns:
            Response:
        """
        return await self.__client.request(
            "PUT",
            config.API_HOST + f"/v1/threads/{thread_id}",
            json=params,
            return_type=Response,
        )
