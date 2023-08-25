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

import base64
import datetime
import hashlib
import hmac
import httpx
import json
import logging
import os
import random
import time
import uuid

from .login import get_token
from .user import get_timestamp

from ..config import Configs, ErrorType, ErrorMessage
from ..errors import (
    HTTPError,
    BadRequestError,
    AuthenticationError,
    ForbiddenError,
    NotFoundError,
    RateLimitError,
    YayServerError,
)
from ..utils import Colors

try:
    from json.decoder import JSONDecodeError
except ImportError:
    JSONDecodeError = ValueError


current_path = os.path.abspath(os.getcwd())


class API:
    def __init__(
        self,
        access_token: str = None,
        proxy: str = None,
        max_retries=3,
        backoff_factor=1.5,
        wait_on_rate_limit=True,
        min_delay=0.3,
        max_delay=1.2,
        timeout=30,
        err_lang="ja",
        base_path=current_path + "/config/",
        save_cookie_file=True,
        encrypt_cookie=False,
        cookie_filename="cookies",
        loglevel=logging.INFO,
    ):
        self.yaylib_version = Configs.YAYLIB_VERSION
        self.api_version = Configs.YAY_API_VERSION
        self.api_key = Configs.YAY_API_KEY
        self.fernet = None
        self._secret_key = None

        self.proxy = {}
        if proxy is not None:
            self.proxy["http://"] = "http://" + proxy
            self.proxy["https://"] = "http://" + proxy

        self.current_ts = None
        self.last_req_ts = None
        self.max_retries = max_retries
        self.retry_statuses = [500, 502, 503, 504]
        self.backoff_factor = backoff_factor
        self.wait_on_rate_limit = wait_on_rate_limit
        self.min_delay = min_delay
        self.max_delay = max_delay
        self.timeout = timeout
        self.err_lang = err_lang
        self.base_path = base_path
        self.save_cookie_file = save_cookie_file
        self.encrypt_cookie = encrypt_cookie
        self.cookie_filename = cookie_filename
        self._cookies = {}

        self._generate_all_uuids()
        self.session = httpx.Client(proxies=self.proxy, timeout=self.timeout)
        self.session.headers.update(Configs.REQUEST_HEADERS)
        self.session.headers.update({"X-Device-UUID": self.device_uuid})
        if access_token is not None:
            self.session.headers.setdefault("Authorization", "Bearer " + access_token)

        self.logger = logging.getLogger("yaylib version: " + self.yaylib_version)

        if not os.path.exists(base_path):
            os.makedirs(base_path)

        ch = logging.StreamHandler()
        ch.setLevel(loglevel)
        ch.setFormatter(logging.Formatter("%(asctime)s - %(levelname)s - %(message)s"))

        self.logger.addHandler(ch)
        self.logger.setLevel(logging.DEBUG)

        self._set_client_ip()

        self.logger.info("yaylib version: " + self.yaylib_version + " started.")

    def _generate_all_uuids(self):
        self.device_uuid = self.generate_uuid(True)
        self.uuid = self.generate_uuid(True)

    def _set_client_ip(self):
        response = get_timestamp(self)
        self.session.headers.update({"X-Client-IP": response.ip_address})

    def _request(
        self,
        method,
        endpoint,
        params=None,
        payload=None,
        user_auth=True,
        headers=None,
        auth_required=False,
        bypass_delay=False,
        access_token=None,
    ):
        headers = headers or self.session.headers.copy()
        headers = self._prepare_auth(headers, access_token, user_auth, auth_required)

        response, backoff_duration = None, 0
        max_rate_limit_retries = 15  # roughly equivalent to 60 mins, plus extra 15 mins
        auth_retry_count, max_auth_retries = 0, 2

        # retry the request based on max_retries
        for i in range(self.max_retries):
            time.sleep(backoff_duration)
            headers = self._prepare_headers(headers)

            self._log_request_info(method, endpoint, params, headers, payload)

            response = self.session.request(
                method, endpoint, params=params, json=payload, headers=headers
            )

            if self.last_req_ts is not None and self.current_ts - self.last_req_ts < 1:
                # insert delays if interval between last request
                # and current request is less than a sec
                self._delay(self.min_delay, self.max_delay)

            self._log_response_info(response)

            if self._is_rate_limit(response) and self.wait_on_rate_limit:
                # continue attempting request until successful
                # or maximum number of retries is reached
                rate_limit_retry_count = 0
                while rate_limit_retry_count < max_rate_limit_retries - 1:
                    retry_after = 60 * 5
                    self.logger.info(
                        f"Rate limit exceeded. Waiting for {retry_after} seconds..."
                    )
                    time.sleep(retry_after + 1)  # sleep for extra sec

                    headers = self._prepare_headers(headers)

                    self._log_request_info(method, endpoint, params, headers, payload)

                    response = self.session.request(
                        method, endpoint, params=params, json=payload, headers=headers
                    )

                    self._log_response_info(response)

                    if not self._is_rate_limit(response):
                        break

                    rate_limit_retry_count += 1

                if rate_limit_retry_count >= max_rate_limit_retries:
                    raise RateLimitError("Maximum rate limit retries exceeded.")

            if response.status_code == 401 and self.save_cookie_file:
                # remove the cookie file and stop the proccessing if refresh token has expired
                if "/api/v1/oauth/token" in endpoint:
                    os.remove(self.base_path + self.cookie_filename + ".json")
                    raise AuthenticationError(
                        "Refresh token expired. Try logging in again."
                    )

                auth_retry_count += 1

                if auth_retry_count < max_auth_retries:
                    headers = self._refresh_access_token(headers)
                    continue

                else:
                    os.remove(self.base_path + self.cookie_filename + ".json")
                    raise AuthenticationError(
                        "Maximum authentication retries exceeded. Try logging in again."
                    )

            if response.status_code not in self.retry_statuses:
                break

            if response is not None:
                self.logger.error(
                    f"Request failed with status code {response.status_code}. Retrying...",
                    exc_info=True,
                )
            else:
                self.logger.error("Request failed. Retrying...")

            backoff_duration = self.backoff_factor * (2**i)

        self._update_last_req_ts(bypass_delay)
        formatted_response = self._format_response(response)

        return self._handle_response(response, formatted_response)

    def _prepare_auth(self, headers, access_token, user_auth, auth_required):
        if access_token is not None:
            headers["Authorization"] = "Bearer " + access_token

        if not user_auth and "Authorization" in headers:
            del headers["Authorization"]

        if auth_required and "Authorization" not in headers:
            raise AuthenticationError("Access Denied - Authentication Required!")

        return headers

    def _prepare_headers(self, headers):
        self._update_current_ts()
        headers.update({"X-Timestamp": str(self.current_ts)})
        return headers

    def _log_request_info(self, method, endpoint, params, headers, payload):
        request_info = (
            f"Making API request:\n\n"
            f"{Colors.HEADER}{method}: {endpoint}{Colors.RESET}\n\n"
            f"Parameters: {params}\n\n"
            f"Headers: {headers}\n\n"
            f"Body: {payload}\n"
        )
        self.logger.debug(request_info)

    def _log_response_info(self, response: httpx.Response):
        response_info = (
            f"Received API response:\n\n"
            f"Status code: {response.status_code}\n\n"
            f"Headers: {response.headers}\n\n"
            f"Response: {response.text}\n"
        )
        self.logger.debug(response_info)

    def _refresh_access_token(self, headers):
        self.logger.debug("Access token expired. Refreshing tokens...")

        response = get_token(
            self, grant_type="refresh_token", refresh_token=self.refresh_token
        )

        self.cookies.update(
            {
                "authentication": {
                    "access_token": response.access_token,
                    "refresh_token": response.refresh_token,
                }
            }
        )

        # copy the cookies to ensure its value remains unchanged during encryption
        cookies = self.cookies.copy()
        self.save_cookies(cookies)

        # only for the next retry
        headers["Authorization"] = "Bearer " + response.access_token

        self.session.headers["Authorization"] = "Bearer " + response.access_token

        return headers

    def _update_current_ts(self):
        self.current_ts = int(datetime.datetime.now().timestamp())

    def _update_last_req_ts(self, bypass_delay):
        self.last_req_ts = None
        if not bypass_delay:
            self.last_req_ts = int(datetime.datetime.now().timestamp())

    def _format_response(self, response: httpx.Response):
        try:
            formatted_response = response.json()
        except JSONDecodeError:
            formatted_response = response.text
        return formatted_response

    def _handle_response(self, response: httpx.Response, formatted_response):
        if isinstance(formatted_response, dict):
            formatted_response = self._translate_error_message(formatted_response)

        if response.status_code == 400:
            raise BadRequestError(formatted_response)
        if response.status_code == 401:
            raise AuthenticationError(formatted_response)
        if response.status_code == 403:
            raise ForbiddenError(formatted_response)
        if response.status_code == 404:
            raise NotFoundError(formatted_response)
        if response.status_code == 429:
            raise RateLimitError(formatted_response)
        if response.status_code == 500:
            raise YayServerError(formatted_response)
        if response.status_code and not 200 <= response.status_code < 300:
            raise HTTPError(formatted_response)

        return formatted_response

    def _translate_error_message(self, response):
        if self.err_lang == "ja":
            try:
                error_code = response.get("error_code", None)
                if error_code is not None:
                    error_type = ErrorType(error_code)
                    if error_type.name in ErrorMessage.__members__:
                        error_message = ErrorMessage[error_type.name].value
                        response["message"] = error_message
                return response
            except ValueError:
                return response
        else:
            return response

    def _make_request(
        self,
        method: str,
        endpoint: str,
        params: dict = None,
        payload: dict = None,
        data_type=None,
        user_auth=True,
        headers=None,
        auth_required=False,
        bypass_delay=False,
        access_token=None,
    ):
        response = self._request(
            method,
            endpoint,
            params,
            payload,
            user_auth,
            headers,
            auth_required,
            bypass_delay,
            access_token,
        )
        if data_type:
            return self._construct_response(response, data_type)
        return response

    @property
    def cookies(self):
        return self._cookies

    @cookies.setter
    def cookies(self, cookies):
        self._cookies = cookies

    @property
    def access_token(self):
        auth_header = self.session.headers.get("Authorization")
        if auth_header is not None:
            return auth_header.replace("Bearer ", "")
        return self.cookies.get("authentication", {}).get("access_token")

    @property
    def refresh_token(self):
        return self.cookies.get("authentication", {}).get("refresh_token")

    @property
    def user_id(self):
        return self.cookies.get("user", {}).get("user_id")

    @property
    def email(self):
        return self.cookies.get("user", {}).get("email")

    @email.setter
    def email(self, email):
        self._cookies["user"] = self._cookies.get("user", {})
        self._cookies["user"]["email"] = email

    @property
    def device_uuid(self):
        return self.cookies.get("device", {}).get("device_uuid")

    @device_uuid.setter
    def device_uuid(self, device_uuid):
        self._cookies["device"] = self._cookies.get("device", {})
        self._cookies["device"]["device_uuid"] = device_uuid

    @property
    def secret_key(self):
        return self._secret_key

    @secret_key.setter
    def secret_key(self, value):
        self._secret_key = value

    @staticmethod
    def encrypt_cookies(fernet, cookies):
        access_token = cookies.get("authentication", {}).get("access_token")
        refresh_token = cookies.get("authentication", {}).get("refresh_token")
        device_uuid = cookies.get("device", {}).get("device_uuid")
        cookies.update(
            {
                "authentication": {
                    "access_token": fernet.encrypt(access_token.encode()).decode(),
                    "refresh_token": fernet.encrypt(refresh_token.encode()).decode(),
                },
                "device": {
                    "device_uuid": fernet.encrypt(device_uuid.encode()).decode()
                },
            }
        )
        return cookies

    @staticmethod
    def decrypt_cookies(fernet, cookies):
        access_token = cookies.get("authentication", {}).get("access_token")
        refresh_token = cookies.get("authentication", {}).get("refresh_token")
        device_uuid = cookies.get("device", {}).get("device_uuid")
        cookies.update(
            {
                "authentication": {
                    "access_token": fernet.decrypt(access_token).decode(),
                    "refresh_token": fernet.decrypt(refresh_token).decode(),
                },
                "device": {"device_uuid": fernet.decrypt(device_uuid).decode()},
            }
        )
        return cookies

    @staticmethod
    def generate_uuid(uuid_type=True):
        generated_uuid = str(uuid.uuid4())
        if uuid_type:
            return generated_uuid
        else:
            return generated_uuid.replace("-", "")

    @staticmethod
    def generate_signed_info(uuid, timestamp, shared_key=False):
        shared_key = Configs.YAY_SHARED_KEY if shared_key is True else ""
        return hashlib.md5(
            (Configs.YAY_API_KEY + uuid + str(timestamp) + shared_key).encode()
        ).hexdigest()

    @staticmethod
    def generate_signed_version():
        return base64.b64encode(
            hmac.new(
                Configs.YAY_API_VERSION_KEY.encode(),
                "yay_android/{}".format(Configs.YAY_API_VERSION).encode(),
                hashlib.sha256,
            ).digest()
        ).decode("utf-8")

    @staticmethod
    def _delay(min_delay, max_delay):
        sleep_time = random.uniform(min_delay, max_delay)
        time.sleep(sleep_time)

    @staticmethod
    def _is_rate_limit(response: httpx.Response):
        if response.status_code == 429:
            return True
        if response.status_code == 400:
            try:
                json_response = response.json()
                if json_response.get("error_code") == -343:
                    return True
            except JSONDecodeError:
                return False
        return False

    @staticmethod
    def _construct_response(data, data_type):
        if data_type is not None:
            if isinstance(data, list):
                data = [data_type(result) for result in data]
            elif data is not None:
                data = data_type(data)
        return data

    def load_cookies(self):
        if not os.path.exists(self.base_path + self.cookie_filename + ".json"):
            return None

        with open(self.base_path + self.cookie_filename + ".json", "r") as f:
            loaded_cookies = json.load(f)

        if self.encrypt_cookie and self.fernet is not None:
            loaded_cookies = self.decrypt_cookies(self.fernet, loaded_cookies)

        return loaded_cookies

    def save_cookies(self, cookies):
        email = cookies.get("user", {}).get("email")
        if email is not None:
            cookies["user"]["email"] = hashlib.sha256(email.encode()).hexdigest()
        else:
            cookies["user"]["email"] = self.load_cookies().get("user", {}).get("email")

        if self.encrypt_cookie and self.fernet is not None:
            cookies = self.encrypt_cookies(self.fernet, cookies)

        with open(self.base_path + self.cookie_filename + ".json", "w") as f:
            json.dump(cookies, f, indent=4)
