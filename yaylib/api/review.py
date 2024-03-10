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
from ..responses import ReviewsResponse


class ReviewAPI(object):
    def __init__(self, base: client.BaseClient) -> None:
        self.__base = base

    def create_review(self, user_id: int, comment: str):
        return self.__base._request(
            "POST",
            route=f"/v1/users/reviews/{user_id}",
            payload={"comment": comment},
        )

    def delete_reviews(self, review_ids: list[int]):
        return self.__base._request(
            "DELETE",
            route=f"/v1/users/reviews",
            params={"review_ids[]": review_ids},
        )

    def get_my_reviews(self, **params) -> ReviewsResponse:
        """

        Parameters
        ----------

            - from_id: int (optional)
            - number: int = (optional)

        """
        return self.__base._request(
            "GET",
            route=f"/v1/users/reviews/mine",
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
        return self.__base._request(
            "GET",
            route=f"/v1/users/reviews/{user_id}",
            params=params,
            data_type=ReviewsResponse,
        )

    def pin_review(self, review_id: int):
        return self.__base._request(
            "POST", route=f"/v1/pinned/reviews", payload={"id": review_id}
        )

    def unpin_review(self, review_id: int):
        return self.__base._request("DELETE", route=f"/v1/pinned/reviews{review_id}")
