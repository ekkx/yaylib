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

from .. import config
from ..responses import ReviewsResponse


class ReviewApi:
    """レビュー API"""

    def __init__(self, client) -> None:
        from ..client import Client  # pylint: disable=import-outside-toplevel

        self.__client: Client = client

    async def create_review(self, user_id: int, comment: str):
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/users/reviews/{user_id}",
            json={"comment": comment},
        )

    async def delete_reviews(self, review_ids: list[int]):
        return await self.__client.request(
            "DELETE",
            config.API_HOST + "/v1/users/reviews",
            params={"review_ids[]": review_ids},
        )

    async def get_my_reviews(self, **params) -> ReviewsResponse:
        """

        Parameters
        ----------

            - from_id: int (optional)
            - number: int = (optional)

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/users/reviews/mine",
            params=params,
            return_type=ReviewsResponse,
        )

    async def get_reviews(self, user_id: int, **params) -> ReviewsResponse:
        """

        Parameters
        ----------

            - user_id: int (required)
            - from_id: int = (optional)
            - number: int = (optional)

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v1/users/reviews/{user_id}",
            params=params,
            return_type=ReviewsResponse,
        )

    async def pin_review(self, review_id: int):
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/pinned/reviews",
            json={"id": review_id},
        )

    async def unpin_review(self, review_id: int):
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v1/pinned/reviews{review_id}",
        )
