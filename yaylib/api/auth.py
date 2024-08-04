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

from datetime import datetime
from typing import Optional

from .. import config
from ..responses import (
    LoginUserResponse,
    LoginUpdateResponse,
    TokenResponse,
)
from ..utils import md5


class AuthApi:
    """認証 API"""

    def __init__(self, client) -> None:
        self.__client = client

    async def change_email(
        self, email: str, password: str, email_grant_token: Optional[str] = None
    ) -> LoginUpdateResponse:
        return await self.__client.request(
            "PUT",
            config.API_HOST + "/v1/users/change_email",
            json={
                "api_key": config.API_KEY,
                "email": email,
                "password": password,
                "email_grant_token": email_grant_token,
            },
            return_type=LoginUpdateResponse,
        )

    async def change_password(
        self, current_password: str, new_password: str
    ) -> LoginUpdateResponse:
        return await self.__client.request(
            "PUT",
            config.API_HOST + "/v1/users/change_email",
            json={
                "api_key": config.API_KEY,
                "current_password": current_password,
                "password": new_password,
            },
            return_type=LoginUpdateResponse,
        )

    async def get_token(
        self,
        grant_type: str,
        refresh_token: Optional[str] = None,
        email: Optional[str] = None,
        password: Optional[str] = None,
    ) -> TokenResponse:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/api/v1/oauth/token",
            json={
                "grant_type": grant_type,
                "email": email,
                "password": password,
                "refresh_token": refresh_token,
            },
            return_type=TokenResponse,
        )

    async def login(self, email: str, password: str) -> LoginUserResponse:
        user = self.__client.state.get_user_by_email(email=email)
        if user is not None:
            self.__client.logger.info(
                f"User found in local storage - UID: {user.user_id}"
            )
            return LoginUserResponse(
                {
                    "access_token": self.__client.access_token,
                    "refresh_token": self.__client.refresh_token,
                    "user_id": self.__client.user_id,
                }
            )

        response: LoginUserResponse = await self.__client.request(
            "POST",
            config.API_HOST + "/v3/users/login_with_email",
            json={
                "api_key": config.API_KEY,
                "email": email,
                "password": password,
                "uuid": self.__client.device_uuid,
            },
            return_type=LoginUserResponse,
        )

        self.__client.state.user_id = response.user_id
        self.__client.state.email = email
        self.__client.state.access_token = response.access_token
        self.__client.state.refresh_token = response.refresh_token
        self.__client.state.create()

        self.__client.logger.info(
            f"Authentication Successful! - UID: {response.user_id}"
        )

        return response

    async def logout(self):
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/users/logout",
            json={"uuid": self.__client.device_uuid},
        )

    async def resend_confirm_email(self):
        return await self.__client.request(
            "POST", config.API_HOST + "/v2/users/resend_confirm_email"
        )

    async def restore_user(self, user_id: int) -> LoginUserResponse:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v2/users/restore",
            json={
                "user_id": user_id,
                "api_key": config.API_KEY,
                "uuid": self.__client.device_uuid,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__client.device_uuid, int(datetime.now().timestamp()), False
                ),
            },
        )

    async def save_account_with_email(
        self,
        email: str,
        password: Optional[str] = None,
        current_password: Optional[str] = None,
        email_grant_token: Optional[str] = None,
    ) -> LoginUpdateResponse:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v3/users/login_update",
            json={
                "api_key": config.API_KEY,
                "email": email,
                "password": password,
                "current_password": current_password,
                "email_grant_token": email_grant_token,
            },
            return_type=LoginUpdateResponse,
        )
