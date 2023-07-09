from datetime import datetime
from typing import List

from ..config import Endpoints
from ..responses import ReviewsResponse
from ..utils import signed_info_calculating


def create_review(self, user_id: int, comment: str):
    self._check_authorization()
    response = self._make_request(
        "POST",
        endpoint=f"{Endpoints.USERS_V1}/reviews/{user_id}",
        payload={"comment": comment},
    )
    self.logger.info(f"Review has been sent to {user_id}.")
    return response


def create_reviews(self, user_ids: List[int], comment: str):
    self._check_authorization()
    timestamp = int(datetime.now().timestamp())
    response = self._make_request(
        "POST",
        endpoint=f"{Endpoints.USERS_V1}/reviews",
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
    self.logger.info("Reviews have been sent to multiple users.")
    return response


def delete_reviews(self, review_ids: List[int]):
    self._check_authorization()
    response = self._make_request(
        "DELETE",
        endpoint=f"{Endpoints.USERS_V1}/reviews",
        params={"review_ids[]": review_ids},
    )
    self.logger.info("Reviews have been deleted.")
    return response


def get_my_reviews(self, **params) -> ReviewsResponse:
    """

    Parameters
    ----------

        - from_id: int (optional)
        - number: int = (optional)

    """
    self._check_authorization()
    return self._make_request(
        "GET",
        endpoint=f"{Endpoints.USERS_V1}/reviews/mine",
        params=params,
        data_type=ReviewsResponse,
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
        "GET",
        endpoint=f"{Endpoints.USERS_V1}/reviews/{user_id}",
        params=params,
        data_type=ReviewsResponse,
    )


def pin_review(self, review_id: int):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.PINNED_V1}/reviews", payload={"id": review_id}
    )
    self.logger.info("Pinned the review.")
    return response


def unpin_review(self, review_id: int):
    self._check_authorization()
    response = self._make_request(
        "DELETE", endpoint=f"{Endpoints.PINNED_V1}/reviews{review_id}"
    )
    self.logger.info("Unpinned the review.")
    return response
