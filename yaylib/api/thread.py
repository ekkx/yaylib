"""
MIT License

Copyright (c) 2023-present qvco

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

from ..config import Endpoints
from ..models import ThreadInfo
from ..responses import GroupThreadListResponse, PostsResponse


def add_post_to_thread(
    self, post_id: int, thread_id: int, access_token: str = None
) -> ThreadInfo:
    response = self.request(
        "PUT",
        endpoint=f"{Endpoints.POSTS_V3}/{post_id}/move_to_thread/{thread_id}",
        data_type=ThreadInfo,
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info(f"Post '{post_id}' added to the thread '{thread_id}'.")
    return response


def convert_post_to_thread(
    self,
    post_id: int,
    title: str = None,
    thread_icon_filename: str = None,
    access_token: str = None,
) -> ThreadInfo:
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.POSTS_V3}/{post_id}/move_to_thread",
        payload={"title": title, "thread_icon_filename": thread_icon_filename},
        data_type=ThreadInfo,
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Post has been converted to a thread.")
    return response


def create_thread(
    self, group_id: int, title: str, thread_icon_filename: str, access_token: str = None
) -> ThreadInfo:
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.THREADS_V1}",
        payload={
            "group_id": group_id,
            "title": title,
            "thread_icon_filename": thread_icon_filename,
        },
        data_type=ThreadInfo,
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("A new thread has been created.")
    return response


def get_group_thread_list(
    self, group_id: int, from_str: str = None, access_token: str = None, **params
) -> GroupThreadListResponse:
    """

    Parameters:
    ----------

        - group_id: int
        - from_str: str = None
        - join_status: str = None

    """
    params["group_id"] = group_id
    if from_str:
        params["from"] = from_str
    return self.request(
        "GET",
        endpoint=f"{Endpoints.THREADS_V1}",
        params=params,
        data_type=GroupThreadListResponse,
        access_token=access_token,
    )


def get_thread_joined_statuses(self, ids: list[int], access_token: str = None) -> dict:
    return self.request(
        "GET",
        endpoint=f"{Endpoints.THREADS_V1}/joined_statuses",
        params={"ids[]": ids},
        auth_required=True,
        access_token=access_token,
    )


def get_thread_posts(
    self, thread_id: int, from_str: str = None, access_token: str = None, **params
) -> PostsResponse:
    """

    Parameters:
    ----------

        - post_type: str
        - number: int = None
        - from_str: str = None

    """
    if from_str:
        params["from"] = from_str
    return self.request(
        "GET",
        endpoint=f"{Endpoints.THREADS_V1}/{thread_id}/posts",
        params=params,
        data_type=PostsResponse,
        access_token=access_token,
    )


def join_thread(self, thread_id: int, user_id: int, access_token: str = None):
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.THREADS_V1}/{thread_id}/members/{user_id}",
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info(f"Joined the thread '{thread_id}'.")
    return response


def leave_thread(self, thread_id: int, user_id: int, access_token: str = None):
    response = self.request(
        "DELETE",
        endpoint=f"{Endpoints.THREADS_V1}/{thread_id}/members/{user_id}",
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Left the thread.")
    return response


def remove_thread(self, thread_id: int, access_token: str = None):
    response = self.request(
        "DELETE",
        endpoint=f"{Endpoints.THREADS_V1}/{thread_id}",
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info(f"Thread '{thread_id}' has been removed.")
    return response


def update_thread(
    self,
    thread_id: int,
    title: str,
    thread_icon_filename: str,
    access_token: str = None,
):
    response = self.request(
        "PUT",
        endpoint=f"{Endpoints.THREADS_V1}/{thread_id}",
        payload={"title": title, "thread_icon_filename": thread_icon_filename},
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info(f"Thread '{thread_id}' has been updated.")
    return response
