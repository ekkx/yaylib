import httpx
import os
import logging

from typing import Optional, Dict, Any

from ..config import *
from ..errors import *
from ..utils import *


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
        self.api_key = Configs.YAY_API_KEY
        self.access_token = access_token
        self.proxy = proxy
        self.proxies = None
        self.timeout = timeout
        self.base_path = base_path
        self.host = "https://" + host

        if self.proxy:
            self.proxies = {f"http://{self.proxy}", f"https://{self.proxy}"}

        self.generate_all_uuids()
        self.session = httpx.Client(proxies=self.proxies, timeout=self.timeout)
        self.session.headers.update(Configs.REQUEST_HEADERS)
        self.session.headers.update({"X-Device-Uuid": self.device_uuid})
        if access_token:
            self.session.headers.setdefault(
                "Authorization", f"Bearer {access_token}"
            )

        self.logger = logging.getLogger(
            "yaylib version: " + self.yaylib_version
        )
        ch = logging.StreamHandler()
        ch.setLevel(loglevel_stream)
        ch.setFormatter(logging.Formatter(
            "%(asctime)s - %(levelname)s - %(message)s"
        ))

        handler_existed = False
        for handler in self.logger.handlers:
            if isinstance(handler, logging.StreamHandler):
                handler_existed = True
                break
        if not handler_existed:
            self.logger.addHandler(ch)
        self.logger.setLevel(logging.DEBUG)

        self.logger.info("yaylib version: " + self.yaylib_version + " started")

    def request(self, method, endpoint, params=None, payload=None, user_auth=True, headers=None):
        if headers is None:
            headers = self.session.headers
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
            self, method: str, endpoint: str, params: dict = None,
            payload: dict = None, data_type=None, user_auth=True, headers=None
    ):
        response = self.request(
            method, endpoint, params, payload, user_auth, headers
        )
        if data_type:
            return self._construct_response(response, data_type)
        return response

    def _construct_response(self, data, data_type):
        if data_type is not None:
            if isinstance(data, list):
                data = [data_type(result) for result in data]
            elif data is not None:
                data = data_type(data)
        return data

    def _check_authorization(self) -> None:
        if self.session.headers.get("Authorization") is None:
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

    def generate_all_uuids(self):
        self.device_uuid = generate_uuid(uuid_type=True)
        self.uuid = generate_uuid(uuid_type=True)
