"""
MIT License

Copyright (c) 2023-present Qvco, Konn

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

from datetime import datetime
from typing import List

from ..config import Endpoints
from ..responses import ReviewsResponse


def create_review(self, user_id: int, comment: str, access_token: str = None):
    self._check_authorization(access_token)
    response = self._make_request(
        "POST",
        endpoint=f"{Endpoints.USERS_V1}/reviews/{user_id}",
        payload={"comment": comment},
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info(f"Review has been sent to {user_id}.")
    return response


def create_reviews(self, user_ids: List[int], comment: str, access_token: str = None):
    self._check_authorization(access_token)
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
            "signed_info": self.generate_signed_info(
                self.uuid, timestamp, shared_key=True
            ),
        },
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Reviews have been sent to multiple users.")
    return response


def delete_reviews(self, review_ids: List[int], access_token: str = None):
    self._check_authorization(access_token)
    response = self._make_request(
        "DELETE",
        endpoint=f"{Endpoints.USERS_V1}/reviews",
        params={"review_ids[]": review_ids},
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Reviews have been deleted.")
    return response


def get_my_reviews(self, access_token: str = None, **params) -> ReviewsResponse:
    """

    Parameters
    ----------

        - from_id: int (optional)
        - number: int = (optional)

    """
    self._check_authorization(access_token)
    return self._make_request(
        "GET",
        endpoint=f"{Endpoints.USERS_V1}/reviews/mine",
        params=params,
        data_type=ReviewsResponse,
        auth_required=True,
        access_token=access_token,
    )


def get_reviews(
    self, user_id: int, access_token: str = None, **params
) -> ReviewsResponse:
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
        access_token=access_token,
    )


def pin_review(self, review_id: int, access_token: str = None):
    self._check_authorization(access_token)
    response = self._make_request(
        "POST",
        endpoint=f"{Endpoints.PINNED_V1}/reviews",
        payload={"id": review_id},
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Pinned the review.")
    return response


def unpin_review(self, review_id: int, access_token: str = None):
    self._check_authorization(access_token)
    response = self._make_request(
        "DELETE",
        endpoint=f"{Endpoints.PINNED_V1}/reviews{review_id}",
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Unpinned the review.")
    return response
