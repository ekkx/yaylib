import asyncio

from .errors import RateLimitError


class RateLimit:
    def __init__(self, wait_on_ratelimit: bool, max_retries: int) -> None:
        self.__wait_on_ratelimit = wait_on_ratelimit
        self.__max_retries = max_retries
        self.__retries_performed = 0
        self.__retry_after = 60 * 5  # retry after 5 minutes

    @property
    def max_retries_reached(self) -> bool:
        return not self.__wait_on_ratelimit or (
            self.__retries_performed >= self.__max_retries
        )

    async def wait(self):
        if not self.__wait_on_ratelimit:
            raise RateLimitError()

        await asyncio.sleep(self.__retry_after)
        self.__retries_performed += 1
