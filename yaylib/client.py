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
import logging
import os
import random
import time
from datetime import datetime
from typing import Awaitable, Callable, Dict, Optional

import aiohttp

from . import __version__, config, utils
from .api.auth import AuthApi
from .api.call import CallApi
from .api.chat import ChatApi
from .api.group import GroupApi
from .api.misc import MiscApi
from .api.notification import NotificationApi
from .api.post import PostApi
from .api.review import ReviewApi
from .api.thread import ThreadApi
from .api.user import UserApi
from .device import Device
from .errors import (
    AccessTokenExpiredError,
    AuthenticationError,
    BadRequestError,
    ErrorCode,
    ForbiddenError,
    HTTPError,
    InternalServerError,
    NotFoundError,
    RateLimitError,
)
from .models import Model
from .responses import PostsResponse
from .state import LocalUser, State

__all__ = ["Client"]


class RateLimit:
    """レート制限を管理するクラス"""

    def __init__(self, wait_on_ratelimit: bool, max_retries: int) -> None:
        self.__wait_on_ratelimit = wait_on_ratelimit
        self.__max_retries = max_retries
        self.__retries_performed = 0
        self.__retry_after = 60 * 5

    @property
    def retries_performed(self) -> int:
        """レート制限によるリトライ回数

        Returns:
            int: リトライ回数
        """
        return self.__retries_performed

    @property
    def max_retries(self) -> int:
        """レート制限によるリトライ回数の上限

        Returns:
            int: リトライ回数の上限
        """
        return self.__max_retries

    def __max_retries_reached(self) -> bool:
        return not self.__wait_on_ratelimit or (
            self.__retries_performed >= self.__max_retries
        )

    def reset(self) -> None:
        """レート制限をリセットする"""
        self.__retries_performed = 0

    async def wait(self, exc: RateLimitError) -> None:
        """レート制限が解除されるまで待機する

        Raises:
            RateLimitError: レート制限エラー
        """
        if not self.__wait_on_ratelimit or self.__max_retries_reached():
            raise exc
        await asyncio.sleep(self.__retry_after)
        self.__retries_performed += 1


class HeaderManager:
    """HTTP ヘッダーのマネージャークラス"""

    def __init__(self, device: Device, state: State, locale="ja") -> None:
        self.__state = state
        self.__locale = locale
        self.__host = config.API_HOST
        self.__user_agent = device.generate_user_agent()
        self.__device_info = device.generate_device_info(config.VERSION_NAME)
        self.__app_version = config.API_VERSION_NAME
        self.__client_ip = ""
        self.__connection_speed = "0 kbps"
        self.__connection_type = "wifi"
        self.__content_type = "application/json;charset=UTF-8"

    @property
    def client_ip(self) -> str:
        return self.__client_ip

    @client_ip.setter
    def client_ip(self, value: str) -> None:
        self.__client_ip = value

    def generate(self, jwt_required=False) -> Dict[str, str]:
        """HTTPヘッダーを生成する"""
        headers = {
            "Host": self.__host,
            "User-Agent": self.__user_agent,
            "X-Timestamp": str(int(datetime.now().timestamp())),
            "X-App-Version": self.__app_version,
            "X-Device-Info": self.__device_info,
            "X-Device-UUID": self.__state.device_uuid,
            "X-Client-IP": self.__client_ip,
            "X-Connection-Type": self.__connection_type,
            "X-Connection-Speed": self.__connection_speed,
            "Accept-Language": self.__locale,
            "Content-Type": self.__content_type,
        }

        if jwt_required:
            headers.update({"X-Jwt": utils.generate_jwt()})

        if self.__client_ip != "":
            headers.update({"X-Client-IP": self.__client_ip})

        if self.__state.access_token != "":
            headers.update({"Authorization": "Bearer " + self.__state.access_token})

        return headers


current_path = os.path.abspath(os.getcwd())


class BaseClient:
    """yaylib クライアントの基底クラス"""

    def __init__(
        self,
        *,
        proxy_url: Optional[str] = None,
        timeout=60,
        base_path=current_path + "/.config/",
        loglevel=logging.INFO,
    ) -> None:
        self.__proxy_url = proxy_url
        self.__timeout = timeout
        self.__session = None

        self.logger = logging.getLogger("yaylib version: " + __version__)

        if not os.path.exists(base_path):
            os.makedirs(base_path)

        ch = logging.StreamHandler()
        ch.setLevel(loglevel)
        ch.setFormatter(utils.CustomFormatter())

        self.logger.addHandler(ch)
        self.logger.setLevel(logging.DEBUG)

        self.logger.info("yaylib version: %s started.", __version__)

    async def __is_ratelimit_error(self, response: aiohttp.ClientResponse) -> bool:
        if response.status == 429:
            return True
        if response.status == 400:
            response_json = await response.json(content_type=None)
            if response_json.get("error_code") == ErrorCode.QuotaLimitExceeded.value:
                return True
        return False

    async def __is_access_token_expired(self, response: aiohttp.ClientResponse) -> bool:
        response_json = await response.json(content_type=None)
        return response.status == 401 and (
            response_json.get("error_code") == ErrorCode.AccessTokenExpired.value
            or response_json.get("error_code") == ErrorCode.AccessTokenInvalid.value
        )

    async def base_request(
        self, method: str, url: str, **kwargs
    ) -> aiohttp.ClientResponse:
        self.logger.debug(
            "Making API request: [%s] %s\n\nParameters: %s\n\nHeaders: %s\n\nBody: %s\n",
            method,
            url,
            kwargs.get("params"),
            kwargs.get("headers"),
            kwargs.get("json"),
        )

        session = self.__session or aiohttp.ClientSession()
        async with session.request(
            method, url, proxy=self.__proxy_url, timeout=self.__timeout, **kwargs
        ) as response:
            await response.read()

        if self.__session is None:
            await session.close()

        self.logger.debug(
            "Received API response: [%s] %s\n\nHTTP Status: %s\n\nHeaders: %s\n\nResponse: %s\n",
            method,
            url,
            response.status,
            response.headers,
            await response.text(),
        )

        if await self.__is_ratelimit_error(response):
            raise RateLimitError(response)
        if await self.__is_access_token_expired(response):
            raise AccessTokenExpiredError(response)

        if response.status == 400:
            raise BadRequestError(response)
        if response.status == 401:
            raise AuthenticationError(response)
        if response.status == 403:
            raise ForbiddenError(response)
        if response.status == 404:
            raise NotFoundError(response)
        if response.status >= 500:
            raise InternalServerError(response)
        if response.status and not 200 <= response.status < 300:
            raise HTTPError(response)

        return response


class Client(
    BaseClient,
    # ws.WebSocketInteractor
):
    """yaylib のエントリーポイント"""

    def __init__(
        self,
        *,
        # intents: Optional[ws.Intents] = None,
        proxy_url: Optional[str] = None,
        timeout=30,
        max_retries=3,
        backoff_factor=1.5,
        wait_on_ratelimit=True,
        max_ratelimit_retries=15,
        min_delay=0.3,
        max_delay=1.2,
        base_path=current_path + "/.config/",
        state: Optional[State] = None,
        loglevel=logging.INFO,
    ) -> None:
        super().__init__(
            proxy_url=proxy_url, timeout=timeout, base_path=base_path, loglevel=loglevel
        )

        self.__min_delay = min_delay
        self.__max_delay = max_delay
        self.__last_request_ts = 0
        self.__max_retries = max_retries
        self.__backoff_factor = backoff_factor

        self.auth = AuthApi(self)
        self.call = CallApi(self)
        self.chat = ChatApi(self)
        self.group = GroupApi(self)
        self.misc = MiscApi(self)
        self.notification = NotificationApi(self)
        self.post = PostApi(self)
        self.review = ReviewApi(self)
        self.thread = ThreadApi(self)
        self.user = UserApi(self)

        self.__state = state or State(storage_path=base_path + "secret.db")
        self.__header_manager = HeaderManager(Device.instance(), self.__state)

        self.__ratelimit = RateLimit(wait_on_ratelimit, max_ratelimit_retries)

    @property
    def state(self) -> State:
        return self.__state

    @property
    def user_id(self) -> int:
        return self.__state.user_id

    @property
    def access_token(self) -> str:
        return self.__state.access_token

    @property
    def refresh_token(self) -> str:
        return self.__state.refresh_token

    @property
    def device_uuid(self) -> str:
        return self.__state.device_uuid

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
        *,
        params: Optional[dict] = None,
        json: Optional[dict] = None,
        headers: Optional[dict] = None,
        return_type: Optional[Model] = None,
    ) -> dict | Model:
        response = await self.base_request(
            method, url, params=params, json=json, headers=headers
        )
        response_json = await response.json(content_type=None)
        return self.__construct_response(response_json, return_type)

    async def __refresh_client_tokens(self) -> None:
        response = await self.auth.get_token(
            grant_type="refresh_token", refresh_token=self.__state.refresh_token
        )
        self.__state.set_user(
            LocalUser(
                user_id=self.__state.user_id,
                email=self.__state.email,
                device_uuid=self.__state.device_uuid,
                access_token=response.access_token,
                refresh_token=response.refresh_token,
            )
        )
        self.__state.update()

    async def request(
        self,
        method: str,
        url: str,
        *,
        params: Optional[dict] = None,
        json: Optional[dict] = None,
        headers: Optional[dict] = None,
        return_type: Optional[Model] = None,
        jwt_required=False,
    ) -> dict | Model:
        if not url.startswith("https://"):
            url = "https://" + url

        if not self.__header_manager.client_ip and "v2/users/timestamp" not in url:
            metadata = await self.user.get_timestamp()
            self.__header_manager.client_ip = metadata.ip_address

        backoff_duration = 0
        headers = headers or {}

        for i in range(max(1, self.__max_retries + 1)):
            try:
                await asyncio.sleep(backoff_duration)
                while True:
                    try:
                        headers.update(self.__header_manager.generate(jwt_required))
                        response = await self.__make_request(
                            method,
                            url,
                            params=utils.filter_dict(params),
                            json=utils.filter_dict(json),
                            headers=headers,
                            return_type=return_type,
                        )
                        break
                    except RateLimitError as exc:
                        self.logger.warning(
                            "Rate limit exceeded. Waiting... (%s/%s)",
                            self.__ratelimit.retries_performed,
                            self.__ratelimit.max_retries,
                        )
                        await self.__ratelimit.wait(exc)
                self.__ratelimit.reset()
                break
            except AccessTokenExpiredError as exc:
                if self.user_id == 0:
                    self.logger.error("Authentication required to perform the action.")
                    raise AuthenticationError(exc.response) from exc
                await self.__refresh_client_tokens()
            except AuthenticationError as exc:
                if "/api/v1/oauth/token" in url:
                    self.__state.destory(self.user_id)
                    self.logger.error(
                        "Failed to refresh credentials. Please try logging in again."
                    )
                    raise exc
            except InternalServerError as exc:
                self.logger.error(
                    "Request failed with status %s! Retrying...", exc.response.status
                )
                backoff_duration = self.__backoff_factor * (2**i)

        return response

    def insert_delay(self, callback: Callable):
        """リクエスト間の時間が1秒未満のときに遅延を挿入します"""
        if int(datetime.now().timestamp()) - self.__last_request_ts < 1:
            time.sleep(random.uniform(self.__min_delay, self.__max_delay))
        self.__last_request_ts = int(datetime.now().timestamp())
        return callback()

    def __sync_request(self, callback: Awaitable):
        """非同期リクエストを同期リクエストに変換します"""
        return self.insert_delay(lambda: asyncio.run(callback))

    def login(self, email: str, password: str, two_fa_code: str = None):
        """メールアドレスでログインする

        Args:
            email (str):
            password (str):
            two_fa_code (str):

        Returns:
            LoginUserResponse:
        """
        response = self.__sync_request(self.auth.login(email, password))

        if response.result == "error" and response.message == "Require 2FA":
            two_fa_code = str(
                input("Enter your 6-digit 2-step verification code or backup code\n-> ")
            )
            response = self.__sync_request(
                self.auth.login(email, password, two_fa_code)
            )

        if response.result == "error" and response.message == "2FA code invalid":
            raise ForbiddenError("Invalid 2-step verification code")

        if response.access_token is None:
            raise ForbiddenError("Invalid email or password.")

        return response

    def get_timeline(self, **params) -> PostsResponse:
        """タイムラインを取得する

        Args:
            number (int):
            number (int, optional):

        Returns:
            PostsResponse:
        """
        return self.__sync_request(self.post.get_timeline(**params))
