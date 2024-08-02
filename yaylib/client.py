import asyncio
import logging
import os
from typing import Optional, Awaitable

import aiohttp

from .api.post import PostApi

from . import __version__
from . import ws
from . import config
from .errors import (
    HTTPError,
    BadRequestError,
    AuthenticationError,
    ForbiddenError,
    NotFoundError,
    RateLimitError,
    YayServerError,
    ErrorCode,
)
from .models import (
    Model,
)
from .ratelimit import RateLimit
from .responses import (
    PostsResponse,
)
from .storage import Storage


__all__ = ["Client"]


class HeaderInterceptor:
    def __init__(self, storage: Storage, locale="ja") -> None:
        self.__locale = locale
        self.__host = config.API_HOST
        self.__user_agent = config.USER_AGENT
        self.__device_info = config.DEVICE_INFO
        self.__app_version = config.API_VERSION_NAME
        self.__storage = storage
        self.__client_ip = ""
        self.__connection_speed = ""
        self.__connection_type = "wifi"
        self.__content_type = "application/json;charset=UTF-8"

    def intercept(self, jwt_required=False) -> dict:
        pass


class BaseClient:
    def __init__(self, proxy_url: Optional[str] = None, timeout: int = 60) -> None:
        self.__proxy_url = proxy_url
        self.__timeout = timeout
        self.__session = aiohttp.ClientSession()

    async def __is_ratelimit_error(self, response: aiohttp.ClientResponse) -> bool:
        if response.status == 429:
            return True
        if response.status == 400:
            response_json = await response.json()
            if response_json.get("error_code") == ErrorCode.QuotaLimitExceeded.value:
                return True
        return False

    def __construct_response(
        self, response: dict, data_type: Optional[Model] = None
    ) -> dict | Model:
        if data_type is not None:
            if isinstance(response, list):
                response = [data_type(result) for result in response]
            elif response is not None:
                response = data_type(response)
        return response

    async def __make_request(
        self,
        method: str,
        url: str,
        params: Optional[dict] = None,
        json: Optional[dict] = None,
        headers: Optional[dict] = None,
    ) -> aiohttp.ClientResponse:
        async with self.__session.request(
            method,
            url,
            params=params,
            json=json,
            headers=headers,
            proxy=self.__proxy_url,
            timeout=self.__timeout,
        ) as response:
            await response.read()

        await self.__session.close()

        is_ratelimited = await self.__is_ratelimit_error(response)
        if is_ratelimited:
            raise RateLimitError(response)

        if response.status == 400:
            raise BadRequestError(response)
        if response.status == 401:
            raise AuthenticationError(response)
        if response.status == 403:
            raise ForbiddenError(response)
        if response.status == 404:
            raise NotFoundError(response)
        if response.status >= 500:
            raise YayServerError(response)
        if response.status and not 200 <= response.status < 300:
            raise HTTPError(response)

        return response

    async def _request(
        self,
        method: str,
        url: str,
        params: Optional[dict] = None,
        json: Optional[dict] = None,
        headers: Optional[dict] = None,
        return_type: Optional[Model] = None,
    ) -> dict | object:
        response = await self.__make_request(method, url, params, json, headers)

        response_json = await response.json()

        return self.__construct_response(response_json, return_type)


current_path = os.path.abspath(os.getcwd())


class Client(
    BaseClient,
    # ws.WebSocketInteractor
):
    def __init__(
        self,
        *,
        intents: Optional[ws.Intents] = None,
        proxy_url: Optional[str] = None,
        timeout: int = 30,
        max_retries: int = 3,
        backoff_factor: float = 1.5,
        wait_on_ratelimit: bool = True,
        max_ratelimit_retries: int = 15,
        min_delay: float = 0.3,
        max_delay: float = 1.2,
        err_lang: str = "ja",
        base_path: str = current_path + "/.config/",
        use_storage: bool = True,
        storage_password: Optional[str] = None,
        storage_filename: str = "cookies",
        loglevel: int = logging.INFO,
    ) -> None:
        super().__init__(proxy_url=proxy_url, timeout=timeout)

        self.__ratelimit = RateLimit(wait_on_ratelimit, max_ratelimit_retries)
        self.__max_retries = max_retries
        self.__backoff_factor = backoff_factor
        self.__min_delay = min_delay
        self.__max_delay = max_delay
        self.__err_lang = err_lang

        # TODO: init storage

        self.post = PostApi(self)

        self.logger: logging.Logger = logging.getLogger(
            "yaylib version: " + __version__
        )

        if not os.path.exists(base_path):
            os.makedirs(base_path)

        ch: logging.StreamHandler = logging.StreamHandler()
        ch.setLevel(loglevel)
        ch.setFormatter(logging.Formatter("%(asctime)s - %(levelname)s - %(message)s"))

        self.logger.addHandler(ch)
        self.logger.setLevel(logging.DEBUG)

        self.logger.info("yaylib version: " + __version__ + " started.")

    async def request(
        self,
        method: str,
        url: str,
        params: Optional[dict] = None,
        json: Optional[dict] = None,
        headers: Optional[dict] = None,
        return_type: Optional[Model] = None,
    ) -> dict | Model:
        # interact with headers, strage and make logs

        # TODO: filter params and json
        # TODO: set client ip address to request header if not exists

        # while True:
        #     try:
        #         response = await self._make_request(
        #             method, url, params, json, headers, return_type
        #         )
        #         break
        #     except RateLimitError:
        #         if self.__ratelimit.max_retries_reached:
        #             raise RateLimitError("Maximum retries reached.")
        #         else:
        #             await self.__ratelimit.wait()
        #             # TODO: update headers

        # TODO: handle access token expiries

        return await self._request(method, url, params, json, headers, return_type)

    def __perfirm_sync(self, callback: Awaitable):
        loop = asyncio.get_event_loop()
        return loop.run_until_complete(callback)

    def __sync_request(self, callback: Awaitable):
        # TODO: create delay

        response = self.__perfirm_sync(callback)

        # TODO: use backoff factor

        return response

    def get_timeline(self, **params) -> PostsResponse:
        return self.__sync_request(self.post.get_timeline(**params))
