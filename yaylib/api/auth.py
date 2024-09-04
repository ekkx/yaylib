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

from cryptography import fernet

from .. import config
from ..responses import LoginUpdateResponse, LoginUserResponse, Response, TokenResponse
from ..state import LocalUser
from ..utils import md5


class AuthApi:
    """認証 API

    Args:
        client (Client):
    """

    def __init__(self, client) -> None:
        # pylint: disable=import-outside-toplevel
        from ..client import Client

        self.__client: Client = client

    async def change_email(self, **params) -> LoginUpdateResponse:
        """メールアドレスを変更する

        Args:
            email (str):
            password (str):
            email_grant_token (str, optional):

        Returns:
            LoginUpdateResponse:
        """
        return await self.__client.request(
            "PUT",
            config.API_HOST + "/v1/users/change_email",
            json=params,
            return_type=LoginUpdateResponse,
        )

    async def change_password(self, **params) -> LoginUpdateResponse:
        """パスワードを変更する

        Args:
            current_password (str):
            new_password (str):

        Returns:
            LoginUpdateResponse:
        """
        return await self.__client.request(
            "PUT",
            config.API_HOST + "/v1/users/change_email",
            json=params,
            return_type=LoginUpdateResponse,
        )

    async def get_token(self, **params) -> TokenResponse:
        """認証トークンを取得する

        Args:
            grant_type (str):
            refresh_token (str, optional):
            email (str, optional):
            password (str, optional):

        Returns:
            TokenResponse:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/api/v1/oauth/token",
            json=params,
            return_type=TokenResponse,
        )

    async def login(
        self, email: str, password: str, two_fa_code: Optional[str] = None
    ) -> LoginUserResponse:
        """メールアドレスでログインする

        Args:
            email (str):
            password (str):
            two_fa_code (str, optional):

        Returns:
            LoginUserResponse:
        """
        if not self.__client.state.has_encryption_key():
            self.__client.state.set_encryption_key(password)

        user = self.__client.state.get_user_by_email(email)
        if user is not None:
            try:
                self.__client.state.set_user(self.__client.state.decrypt(user))
            except fernet.InvalidToken as exc:
                self.__client.state.destory(user.user_id)
                self.__client.logger.error(
                    # pylint: disable=line-too-long
                    "Failed to decrypt the credentials stored locally. This might be due to a recent password change. Please try logging in again."
                )
                raise exc

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

        payload = {
            "api_key": config.API_KEY,
            "email": email,
            "password": password,
            "uuid": self.__client.device_uuid,
        }
        if two_fa_code is not None:
            payload["two_fa_code"] = two_fa_code

        response: LoginUserResponse = await self.__client.request(
            "POST",
            config.API_HOST + "/v3/users/login_with_email",
            json=payload,
            return_type=LoginUserResponse,
        )

        self.__client.state.set_user(
            LocalUser(
                user_id=response.user_id,
                email=email,
                device_uuid=self.__client.device_uuid,
                access_token=response.access_token,
                refresh_token=response.refresh_token,
            )
        )
        self.__client.state.save()

        self.__client.logger.info(
            f"Authentication successful! - UID: {response.user_id}"
        )

        return response

    async def resend_confirm_email(self) -> Response:
        """確認メールを再送信する

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v2/users/resend_confirm_email",
            return_type=Response,
        )

    async def restore_user(self, **params) -> LoginUserResponse:
        """ユーザーを復元する

        Args:
            user_id (int):

        Returns:
            LoginUserResponse:
        """
        params.update(
            {
                "api_key": config.API_KEY,
                "uuid": self.__client.device_uuid,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__client.device_uuid,
                    int(datetime.now().timestamp()),
                    False,
                ),
            }
        )
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v2/users/restore",
            json=params,
        )

    async def save_account_with_email(self, **params) -> LoginUpdateResponse:
        """メールアドレスでアカウントを保存する

        Args:
            email (str):
            password (str, optional):
            current_password (str, optional):
            email_grant_token (str, optional):

        Returns:
            LoginUpdateResponse:
        """
        params.update({"api_key": config.API_KEY})
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v3/users/login_update",
            json=params,
            return_type=LoginUpdateResponse,
        )
