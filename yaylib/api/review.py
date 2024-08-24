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
from ..responses import Response, ReviewsResponse


class ReviewApi:
    """レビュー API"""

    def __init__(self, client) -> None:
        # pylint: disable=import-outside-toplevel
        from ..client import Client

        self.__client: Client = client

    async def create_review(self, user_id: int, comment: str) -> Response:
        """レターを送信する

        Args:
            user_id (int):
            comment (str):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/users/reviews/{user_id}",
            json={"comment": comment},
            return_type=Response,
        )

    async def delete_reviews(self, review_ids: List[int]) -> Response:
        """レターを削除する

        Args:
            review_ids (List[int]):

        Returns:
            Response:
        """
        return await self.__client.request(
            "DELETE",
            config.API_HOST + "/v1/users/reviews",
            params={"review_ids[]": review_ids},
            return_type=Response,
        )

    async def get_my_reviews(self, **params) -> ReviewsResponse:
        """送信したレターを取得する

        Args:
            from_id (int, optional):
            number (int, optional):

        Returns:
            ReviewsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/users/reviews/mine",
            params=params,
            return_type=ReviewsResponse,
        )

    async def get_reviews(self, user_id: int, **params) -> ReviewsResponse:
        """ユーザーが受け取ったレターを取得する

        Args:
            user_id (int):
            from_id (int, optional):
            number (int, optional):

        Returns:
            ReviewsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v1/users/reviews/{user_id}",
            params=params,
            return_type=ReviewsResponse,
        )

    async def pin_review(self, review_id: int) -> Response:
        """レターをピン留めする

        Args:
            review_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/pinned/reviews",
            json={"id": review_id},
            return_type=Response,
        )

    async def unpin_review(self, review_id: int) -> Response:
        """レターのピン留めを解除する

        Args:
            review_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v1/pinned/reviews{review_id}",
            return_type=Response,
        )
