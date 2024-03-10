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

from datetime import datetime

from .. import client
from ..config import Configs
from ..responses import (
    LoginUserResponse,
    LoginUpdateResponse,
    RegisterDeviceTokenResponse,
    TokenResponse,
)
from ..utils import md5


class AuthAPI(object):
    def __init__(self, base: client.BaseClient) -> None:
        self.__base = base

    def change_email(
        self, email: str, password: str, email_grant_token: str = None
    ) -> LoginUpdateResponse:
        return self.__base._request(
            "PUT",
            route="/v1/users/change_email",
            payload={
                "api_key": Configs.API_KEY,
                "email": email,
                "password": password,
                "email_grant_token": email_grant_token,
            },
            data_type=LoginUpdateResponse,
        )

    def change_password(
        self, current_password: str, new_password: str
    ) -> LoginUpdateResponse:
        return self.__base._request(
            "PUT",
            route="/v1/users/change_email",
            payload={
                "api_key": Configs.API_KEY,
                "current_password": current_password,
                "password": new_password,
            },
            data_type=LoginUpdateResponse,
        )

    def get_token(
        self,
        grant_type: str,
        refresh_token: str = None,
        email: str = None,
        password: str = None,
    ) -> TokenResponse:
        return self.__base._request(
            "POST",
            route="/api/v1/oauth/token",
            payload={
                "grant_type": grant_type,
                "email": email,
                "password": password,
                "refresh_token": refresh_token,
            },
            data_type=TokenResponse,
            bypass_delay=True,
        )

    def login_with_email(self, email: str, password: str) -> LoginUserResponse:
        """

        メールアドレスでログインします

        """
        return self.__base._request(
            "POST",
            route="/v3/users/login_with_email",
            payload={
                "api_key": Configs.API_KEY,
                "email": email,
                "password": password,
                "uuid": self.__base.uuid,
            },
            data_type=LoginUserResponse,
        )

    def logout(self):
        return self.__base._request(
            "POST",
            route="/v1/users/logout",
            payload={"uuid": self.__base.uuid},
        )

    def resend_confirm_email(self):
        return self.__base._request("POST", route="/v2/users/resend_confirm_email")

    def restore_user(self, user_id: int) -> LoginUserResponse:
        return self.__base._request(
            "POST",
            route="/v2/users/restore",
            payload={
                "user_id": user_id,
                "api_key": Configs.API_KEY,
                "uuid": self.__base.uuid,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": self.__signed_info,
            },
        )

    def register_device_token(
        self,
        device_token: str,
        device_type: str,
        os_version: str,
        app_version: str,
        screen_resolution: str,
        screen_density: str,
        device_model: str,
        appsflyer_id: str,
        advertising_id: str = None,
    ) -> RegisterDeviceTokenResponse:
        return self.__base._request(
            "POST",
            route="/v2/users/device_tokens/new",
            payload={
                "device_token": device_token,
                "device_type": device_type,
                "uuid": self.__base.uuid,
                "os_version": os_version,
                "app_version": app_version,
                "screen_resolution": screen_resolution,
                "screen_density": screen_density,
                "device_model": device_model,
                "appsflyer_id": appsflyer_id,
                "advertising_id": advertising_id,
            },
            data_type=RegisterDeviceTokenResponse,
        )

    def revoke_tokens(self):
        return self.__base._request(
            "DELETE",
            route="/v1/users/device_tokens",
        )

    def save_account_with_email(
        self,
        email: str,
        password: str = None,
        current_password: str = None,
        email_grant_token: str = None,
    ) -> LoginUpdateResponse:
        return self.__base._request(
            "POST",
            route="/v3/users/login_update",
            payload={
                "api_key": Configs.API_KEY,
                "email": email,
                "password": password,
                "current_password": current_password,
                "email_grant_token": email_grant_token,
            },
            data_type=LoginUpdateResponse,
        )

    @property
    def __signed_info(self) -> str:
        return md5(self.__base.device_uuid, int(datetime.now().timestamp()), False)
