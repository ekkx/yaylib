from datetime import datetime
from typing import Dict, List

from ..config import *
from ..errors import *
from ..models import *
from ..responses import *
from ..utils import *


def add_post_to_thread(self, post_id: int, thread_id: int) -> ThreadInfo:
    self._check_authorization()
    response = self._make_request(
        "PUT",
        endpoint=f"{Endpoints.POSTS_V3}/{post_id}/move_to_thread/{thread_id}",
        data_type=ThreadInfo,
    )
    self.logger.info(f"Post '{post_id}' added to the thread '{thread_id}'.")
    return response


def convert_post_to_thread(
    self, post_id: int, title: str = None, thread_icon_filename: str = None
) -> ThreadInfo:
    self._check_authorization()
    response = self._make_request(
        "POST",
        endpoint=f"{Endpoints.POSTS_V3}/{post_id}/move_to_thread",
        payload={"title": title, "thread_icon_filename": thread_icon_filename},
        data_type=ThreadInfo,
    )
    self.logger.info("Post has been converted to a thread.")
    return response


def create_thread(
    self, group_id: int, title: str, thread_icon_filename: str
) -> ThreadInfo:
    self._check_authorization()
    response = self._make_request(
        "POST",
        endpoint=f"{Endpoints.THREADS_V1}",
        payload={
            "group_id": group_id,
            "title": title,
            "thread_icon_filename": thread_icon_filename,
        },
        data_type=ThreadInfo,
    )
    self.logger.info("A new thread has been created.")
    return response


def get_group_thread_list(
    self, group_id: int, from_str: str = None, **params
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
    return self._make_request(
        "GET",
        endpoint=f"{Endpoints.THREADS_V1}",
        params=params,
        data_type=GroupThreadListResponse,
    )


def get_thread_joined_statuses(self, ids: List[int]) -> dict:
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.THREADS_V1}/joined_statuses", params={"ids[]": ids}
    )


def get_thread_posts(
    self, thread_id: int, from_str: str = None, **params
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
    return self._make_request(
        "GET",
        endpoint=f"{Endpoints.THREADS_V1}/{thread_id}/posts",
        params=params,
        data_type=PostsResponse,
    )


def join_thread(self, thread_id: int, user_id: int):
    self._check_authorization()
    response = self._make_request(
        "POST",
        endpoint=f"{Endpoints.THREADS_V1}/{thread_id}/members/{user_id}",
    )
    self.logger.info(f"Joined the thread '{thread_id}'.")
    return response


def leave_thread(self, thread_id: int, user_id: int):
    self._check_authorization()
    response = self._make_request(
        "DELETE",
        endpoint=f"{Endpoints.THREADS_V1}/{thread_id}/members/{user_id}",
    )
    self.logger.info("Left the thread.")
    return response


def remove_thread(self, thread_id: int):
    self._check_authorization()
    response = self._make_request(
        "DELETE",
        endpoint=f"{Endpoints.THREADS_V1}/{thread_id}",
    )
    self.logger.info(f"Thread '{thread_id}' has been removed.")
    return response


def update_thread(self, thread_id: int, title: str, thread_icon_filename: str):
    self._check_authorization()
    response = self._make_request(
        "PUT",
        endpoint=f"{Endpoints.THREADS_V1}/{thread_id}",
        payload={"title": title, "thread_icon_filename": thread_icon_filename},
    )
    self.logger.info(f"Thread '{thread_id}' has been updated.")
    return response
