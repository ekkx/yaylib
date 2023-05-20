from typing import (Dict)

from ..config import (Endpoints, Configs)
from ..errors import (
    YayError,
    ForbiddenError,
    RateLimitError,
    AuthenticationError,
)
from .api import (
    _get,
    _post,
    _put,
    _delete,
    _check_authorization,
    _handle_response,
)


def add_bookmark(self, user_id: int, post_id: int, headers: Dict[str, str | int] = None):
    _check_authorization(self, headers)
    return _put(
        self=self,
        endpoint=f"https://{Endpoints.USER_V1}/{user_id}/bookmarks/{post_id}",
        headers=headers
    )


def add_group_highlight_post(self, group_id: int, post_id: int, headers: Dict[str, str | int] = None):
    _check_authorization(self, headers)
    return _put(
        self=self,
        endpoint=f"https://{Endpoints.GROUPS_V1}/{group_id}/highlights/{post_id}",
        headers=headers
    )


def create_group_call_post(self, group_id: int, post_id: int, headers: Dict[str, str | int] = None):
    _check_authorization(self, headers)
    return _put(
        self=self,
        endpoint=f"https://{Endpoints.GROUPS_V1}/{group_id}/highlights/{post_id}",
        headers=headers
    )


def get_my_posts(self, from_post_id: int = None, number: int = 100, include_group_post: bool = False, headers: Dict[str, str | int] = None):
    # include_group_postはfalseだったらサークルの投稿は含まないはずなのにサークルの投稿しか出てこないしなんかおかしい
    _check_authorization(self, headers)
    data = {
        "number": number,
        "include_group_post": include_group_post
    }
    if from_post_id:
        data["from_post_id"] = from_post_id
    return _get(
        self=self,
        endpoint=f"https://{Endpoints.POSTS_V2}/mine",
        data=data,
        headers=headers
    )