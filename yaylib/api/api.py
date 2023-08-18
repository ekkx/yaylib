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
import hashlib
import hmac
import httpx
import json
import logging
import os
import time
import uuid

from .login import get_token
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
        backoff_factor=1.0,
        wait_on_rate_limit=True,
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

        self.max_retries = max_retries
        self.retry_statuses = [500, 502, 503, 504]
        self.backoff_factor = backoff_factor
        self.wait_on_rate_limit = wait_on_rate_limit
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
        self.session.headers.update({"X-Device-Uuid": self.device_uuid})
        if access_token is not None:
            self.session.headers.setdefault("Authorization", f"Bearer {access_token}")

        self.logger = logging.getLogger("yaylib version: " + self.yaylib_version)

        if not os.path.exists(base_path):
            os.makedirs(base_path)

        ch = logging.StreamHandler()
        ch.setLevel(loglevel)
        ch.setFormatter(logging.Formatter("%(asctime)s - %(levelname)s - %(message)s"))

        self.logger.addHandler(ch)
        self.logger.setLevel(logging.DEBUG)

        self.logger.info("yaylib version: " + self.yaylib_version + " started")

    def _request(
        self,
        method,
        endpoint,
        params=None,
        payload=None,
        user_auth=True,
        headers=None,
        auth_required=False,
        access_token=None,
    ):
        headers = headers or self.session.headers.copy()

        if access_token is not None:
            headers["Authorization"] = f"Bearer {access_token}"

        if not user_auth and "Authorization" in headers:
            del headers["Authorization"]

        if auth_required and "Authorization" not in headers:
            raise AuthenticationError("Access Denied – Authentication Required!")

        response = None
        backoff_duration = 0
        rate_limit_retry_count = 0
        max_rate_limit_retries = 15  # roughly equivalent to 60 mins, plus extra 15 mins
        auth_retry_count = 0
        max_auth_retries = 2

        # retry the request based on max_retries
        for i in range(self.max_retries):
            time.sleep(backoff_duration)

            self.logger.debug(
                "Making API request:\n\n"
                f"{Colors.HEADER}{method}: {endpoint}{Colors.RESET}\n\n"
                f"Parameters: {params}\n\n"
                f"Headers: {headers}\n\n"
                f"Body: {payload}\n"
            )

            response = self.session.request(
                method, endpoint, params=params, json=payload, headers=headers
            )

            self.logger.debug(
                "Received API response:\n\n"
                f"Status Code: {response.status_code}\n\n"
                f"Headers: {response.headers}\n\n"
                f"Response: {response.text}\n"
            )

            if response.status_code == 429 and self.wait_on_rate_limit:
                # continue attempting request until successful
                # or maximum number of retries is reached
                rate_limit_retry_count += 1
                if rate_limit_retry_count < max_rate_limit_retries:
                    retry_after = 60 * 5
                    self.logger.info(
                        f"Rate limit exceeded. Waiting for {retry_after} seconds..."
                    )
                    time.sleep(retry_after + 1)  # sleep for extra sec
                    continue

                else:
                    raise RateLimitError("Maximum rate limit retries exceeded.")

            if response.status_code == 401 and self.save_cookie_file is True:
                # remove the cookie file and stop the proccessing if refresh token has expired
                if "/api/v1/oauth/token" in endpoint:
                    os.remove(self.base_path + self.cookie_filename + ".json")
                    raise AuthenticationError(
                        "Refresh token expired. Try logging in again."
                    )

                auth_retry_count += 1

                if auth_retry_count < max_auth_retries:
                    # refresh access token using the stored refresh token
                    self.logger.debug("Access token expired. Refreshing tokens...")

                    response = get_token(
                        self,
                        grant_type="refresh_token",
                        refresh_token=self.refresh_token,
                    )

                    self.cookies = {
                        "access_token": response.access_token,
                        "refresh_token": response.refresh_token,
                        "user_id": response.user_id,
                        "email": self.email,
                    }

                    # copy the cookies to ensure its value remains unchanged during encryption
                    cookies = self.cookies.copy()
                    self.save_cookies(cookies)

                    # only for the next retry
                    headers["Authorization"] = f"Bearer {response.access_token}"

                    self.session.headers[
                        "Authorization"
                    ] = f"Bearer {response.access_token}"

                    # continue to the next retry iteration
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

        if response is None:
            return None

        try:
            formatted_response = response.json()
        except JSONDecodeError:
            formatted_response = response.text

        return self._handle_response(response, formatted_response)

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
        result = all(key in cookies for key in Configs.COOKIE_PROPERTIES)
        if result is False:
            raise ValueError("Invalid cookie properties.")
        self._cookies = cookies

    @property
    def access_token(self):
        auth_header = self.session.headers.get("Authorization")
        if auth_header is not None:
            return self.cookies.get("access_token") or auth_header.replace(
                "Bearer ", ""
            )
        return self.cookies.get("access_token")

    @property
    def refresh_token(self):
        return self.cookies.get("refresh_token")

    @property
    def user_id(self):
        return self.cookies.get("user_id")

    @property
    def email(self):
        return self.cookies.get("email")

    @email.setter
    def email(self, email):
        self._cookies["email"] = email

    @property
    def secret_key(self):
        return self._secret_key

    @secret_key.setter
    def secret_key(self, value):
        self._secret_key = value

    @staticmethod
    def encrypt_cookies(fernet, cookies):
        access_token = cookies.get("access_token")
        refresh_token = cookies.get("refresh_token")
        cookies.update(
            {
                "access_token": fernet.encrypt(access_token.encode()).decode(),
                "refresh_token": fernet.encrypt(refresh_token.encode()).decode(),
            }
        )
        return cookies

    @staticmethod
    def decrypt_cookies(fernet, cookies):
        access_token = cookies.get("access_token")
        refresh_token = cookies.get("refresh_token")
        cookies.update(
            {
                "access_token": fernet.decrypt(access_token).decode(),
                "refresh_token": fernet.decrypt(refresh_token).decode(),
            }
        )
        return cookies

    @staticmethod
    def _construct_response(data, data_type):
        if data_type is not None:
            if isinstance(data, list):
                data = [data_type(result) for result in data]
            elif data is not None:
                data = data_type(data)
        return data

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

    def load_cookies(self):
        if not os.path.exists(self.base_path + self.cookie_filename + ".json"):
            return None

        with open(self.base_path + self.cookie_filename + ".json", "r") as f:
            loaded_cookies = json.load(f)

        result = all(key in loaded_cookies for key in Configs.COOKIE_PROPERTIES)
        if result is False:
            os.remove(self.base_path + self.cookie_filename + ".json")
            raise ValueError("Invalid cookie properties.")

        if self.encrypt_cookie and self.fernet is not None:
            loaded_cookies = self.decrypt_cookies(self.fernet, loaded_cookies)

        return loaded_cookies

    def save_cookies(self, cookies):
        email = cookies.get("email")
        if email is not None:
            cookies["email"] = hashlib.sha256(email.encode()).hexdigest()
        else:
            cookies["email"] = self.load_cookies().get("email")

        result = all(key in cookies for key in Configs.COOKIE_PROPERTIES)
        if result is False:
            raise ValueError("Invalid cookie properties.")

        if self.encrypt_cookie and self.fernet is not None:
            cookies = self.encrypt_cookies(self.fernet, cookies)

        with open(self.base_path + self.cookie_filename + ".json", "w") as f:
            json.dump(cookies, f, indent=4)

    def _handle_response(self, response, formatted_response):
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

    def _generate_all_uuids(self):
        self.device_uuid = self.generate_uuid(True)
        self.uuid = self.generate_uuid(True)
