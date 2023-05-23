import httpx
import os
import logging
import uuid

from typing import Optional, Dict, Any

from ..config import *
from ..errors import *


current_path = os.path.abspath(os.getcwd())


class API:

    def __init__(
        self,
        access_token: str = None,
        proxy: str = None,
        timeout=10,
        base_path=current_path,
        loglevel_stream=logging.INFO,
        host=Configs.YAY_PRODUCTION_HOST
    ):
        self.yaylib_version = Configs.YAYLIB_VERSION
        self.yay_api_version = Configs.YAY_API_VERSION
        self.yay_version_name = Configs.YAY_VERSION_NAME
        self.yay_api_key = Configs.YAY_API_KEY
        self.access_token = access_token
        self.proxy = proxy
        self.proxies = None
        self.timeout = timeout
        self.base_path = base_path
        self.host = host
        self.uuid = str(uuid.uuid4())
        self.session = httpx.Client(proxies=self.proxies, timeout=self.timeout)
        self.headers = {
            "Host": self.host,
            "X-App-Version": self.yay_api_version,
            "X-Device-Info": f"yay {self.yay_version_name} android 11 (3.5x 1440x2960 Galaxy S9)",
            "X-Device-Uuid": self.uuid,
            "X-Connection-Type": "wifi",
            "Accept-Language": "ja",
            "Content-Type": "application/json;charset=UTF-8"
        }
        self.logger = logging.getLogger(
            "yaylib version: " + self.yaylib_version)
        ch = logging.StreamHandler()
        ch.setLevel(loglevel_stream)
        ch.setFormatter(logging.Formatter(
            "%(asctime)s - %(levelname)s - %(message)s"))

        handler_existed = False
        for handler in self.logger.handlers:
            if isinstance(handler, logging.StreamHandler):
                handler_existed = True
                break
        if not handler_existed:
            self.logger.addHandler(ch)
        self.logger.setLevel(logging.DEBUG)

        if access_token:
            self.headers.setdefault("Authorization", f"Bearer {access_token}")

        self.logger.info("yaylib version: " + self.yaylib_version + " started")

    def request(self, method, endpoint, params=None, payload=None, user_auth=True, headers=None):
        if headers is None:
            headers = self.headers
        if not user_auth:
            headers["Authorization"] = None

        self.logger.debug(
            "Making API request:\n"
            f"{method}: {endpoint}\n"
            f"Parameters: {params}\n"
            f"Headers: {headers}\n"
            f"Body: {payload}"
        )

        response = self.session.request(
            method, endpoint, params=params, json=payload, headers=headers
        )
        self.logger.debug(
            "Received API response:\n"
            f"{response.status_code}\n"
            f"Headers: {response.headers}\n"
            f"Text: {response.text}\n"
        )
        self._handle_response(response)
        return response.json()

    def _make_request(
        self, method: str, endpoint: str, params: dict = None, payload: dict = None,
        data_type=None, user_auth=True, headers=None
    ):
        response = self.request(
            method, endpoint, params, payload, user_auth, headers
        )
        if data_type:
            return self._construct_response(response, data_type)
        return response

    def _construct_response(self, data, data_type):
        if data_type is not None:
            # TODO: 他のキーも追加する
            keys = [
                "user", "users", "post", "posts"
            ]
            for key in keys:
                if key in data:
                    data = data[key]
                    break

            if isinstance(data, list):
                data = [data_type(result) for result in data]
            elif data is not None:
                data = data_type(data)
        return data

    def _check_authorization(self) -> None:
        if self.headers.get("Authorization") is None:
            raise AuthenticationError(
                "Authorization is not present in the header.")

    def _handle_response(self, response):
        if response.status_code == 400:
            raise BadRequestError(response.text)
        if response.status_code == 401:
            raise AuthenticationError(response.text)
        if response.status_code == 403:
            raise ForbiddenError(response.text)
        if response.status_code == 404:
            raise NotFoundError(response.text)
        if response.status_code == 429:
            raise RateLimitError(response.text)
        if response.status_code == 500:
            raise YayServerError(response.text)
        if response.status_code and not 200 <= response.status_code < 300:
            raise HTTPError(response.text)
        return response
