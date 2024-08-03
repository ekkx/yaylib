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

import asyncio

from .errors import RateLimitError


class RateLimit:
    """レート制限を管理するクラス"""

    def __init__(self, wait_on_ratelimit: bool, max_retries: int) -> None:
        self.__wait_on_ratelimit = wait_on_ratelimit
        self.__max_retries = max_retries
        self.__retries_performed = 0
        self.__retry_after = 60 * 5  # retry after 5 minutes

    def __max_retries_reached(self) -> bool:
        return not self.__wait_on_ratelimit or (
            self.__retries_performed >= self.__max_retries
        )

    async def wait(self):
        """レート制限が解除されるまで待機する

        Raises:
            RateLimitError: レート制限エラー
        """
        if not self.__wait_on_ratelimit or self.__max_retries_reached():
            raise RateLimitError("Maximum retries reached.")
        await asyncio.sleep(self.__retry_after)
        self.__retries_performed += 1
