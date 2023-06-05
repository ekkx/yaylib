from datetime import datetime
from typing import Dict, List

from ..config import *
from ..errors import *
from ..models import *
from ..responses import *
from ..utils import *


def create_review(self, user_id: int, comment: str):
    self._check_authorization()
    timestamp = int(datetime.now().timestamp())
    return self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V2}/reviews/{user_id}",
        payload={
            "comment": comment,
            "uuid": self.uuid,
            "api_key": self.api_key,
            "timestamp": timestamp,
            "signed_info": signed_info_calculating(
                self.uuid, timestamp, shared_key=True
            ),
        },
    )


def create_reviews(self, user_ids: List[int], comment: str):
    self._check_authorization()
    timestamp = int(datetime.now().timestamp())
    return self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V1}/reviews",
        payload={
            "user_ids": user_ids,
            "comment": comment,
            "uuid": self.uuid,
            "api_key": self.api_key,
            "timestamp": timestamp,
            "signed_info": signed_info_calculating(
                self.uuid, timestamp, shared_key=True
            ),
        },
    )


def delete_reviews(self, review_ids: List[int]):
    self._check_authorization()
    return self._make_request(
        "DELETE", endpoint=f"{Endpoints.USERS_V1}/reviews",
        params={"review_ids[]": review_ids}
    )


def get_my_reviews(self, **params) -> ReviewsResponse:
    """

    Parameters
    ----------

        - from_id: int (optional)
        - number: int = (optional)

    """
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V1}/reviews/mine",
        params=params, data_type=ReviewsResponse
    )


def get_reviews(self, user_id: int, **params) -> ReviewsResponse:
    """

    Parameters
    ----------

        - user_id: int (required)
        - from_id: int = (optional)
        - number: int = (optional)

    """
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V1}/reviews/{user_id}",
        params=params, data_type=ReviewsResponse
    )


def pin_review(self, review_id: int):
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.PINNED_V1}/reviews",
        payload={"id": review_id}
    )


def unpin_review(self, review_id: int):
    self._check_authorization()
    return self._make_request(
        "DELETE", endpoint=f"{Endpoints.PINNED_V1}/reviews{review_id}"
    )
