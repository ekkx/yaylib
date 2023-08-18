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

from datetime import datetime
from ..config import Endpoints
from ..errors import ForbiddenError
from ..responses import LoginUserResponse, LoginUpdateResponse, TokenResponse


def change_email(
    self,
    email: str,
    password: str,
    email_grant_token: str = None,
    access_token: str = None,
) -> LoginUpdateResponse:
    response = self._make_request(
        "PUT",
        endpoint=f"{Endpoints.USERS_V1}/change_email",
        payload={
            "api_key": self.api_key,
            "email": email,
            "password": password,
            "email_grant_token": email_grant_token,
        },
        data_type=LoginUpdateResponse,
        access_token=access_token,
    )
    self.logger.info("Your email has been changed.")
    return response


def change_password(
    self, current_password: str, new_password: str, access_token: str = None
) -> LoginUpdateResponse:
    response = self._make_request(
        "PUT",
        endpoint=f"{Endpoints.USERS_V1}/change_email",
        payload={
            "api_key": self.api_key,
            "current_password": current_password,
            "password": new_password,
        },
        data_type=LoginUpdateResponse,
        access_token=access_token,
    )
    self.logger.info("Your password has been changed..")
    return response


def get_token(
    self,
    grant_type: str,
    refresh_token: str = None,
    email: str = None,
    password: str = None,
    access_token: str = None,
) -> TokenResponse:
    return self._make_request(
        "POST",
        endpoint=f"{Endpoints.BASE_API_URL}/api/v1/oauth/token",
        payload={
            "grant_type": grant_type,
            "email": email,
            "password": password,
            "refresh_token": refresh_token,
        },
        data_type=TokenResponse,
        access_token=access_token,
    )


def login_with_email(self, email: str, password: str) -> LoginUserResponse:
    """

    メールアドレスでログインします

    """
    response = self._make_request(
        "POST",
        endpoint=f"{Endpoints.USERS_V3}/login_with_email",
        payload={
            "api_key": self.api_key,
            "email": email,
            "password": password,
            "uuid": self.uuid,
        },
        data_type=LoginUserResponse,
    )
    if response.access_token is None:
        raise ForbiddenError("Invalid email or password.")
    self.logger.info(f"Successfully logged in as '{response.user_id}'")
    return response


def logout(self, access_token: str = None):
    try:
        self._check_authorization(access_token)
        response = self._make_request(
            "POST",
            endpoint=f"{Endpoints.USERS_V1}/logout",
            payload={"uuid": self.uuid},
            access_token=access_token,
        )
        self._cookies = {}
        self.session.headers.pop("Authorization", None)
        self.logger.info("User has logged out.")
        return response

    except:
        self.logger.error("User is not logged in.")
        return None


def resend_confirm_email(self, access_token: str = None):
    return self._make_request(
        "POST",
        endpoint=f"{Endpoints.USERS_V2}/resend_confirm_email",
        access_token=access_token,
    )


def restore_user(self, user_id: int, access_token: str = None) -> LoginUserResponse:
    timestamp = int(datetime.now().timestamp())
    response = self._make_request(
        "POST",
        endpoint=f"{Endpoints.USERS_V2}/restore",
        payload={
            "user_id": user_id,
            "api_key": self.api_key,
            "uuid": self.uuid,
            "timestamp": timestamp,
            "signed_info": self.generate_signed_info(self.device_uuid, timestamp),
        },
        access_token=access_token,
    )
    self.logger.info("User has been restored.")
    return response


def revoke_tokens(self, access_token: str = None):
    response = self._make_request(
        "DELETE",
        endpoint=f"{Endpoints.USERS_V1}/device_tokens",
        access_token=access_token,
    )
    self.logger.info("Token has been revoked.")
    return response


def save_account_with_email(
    self,
    email: str,
    password: str = None,
    current_password: str = None,
    email_grant_token: str = None,
    access_token: str = None,
) -> LoginUpdateResponse:
    response = self._make_request(
        "POST",
        endpoint=f"{Endpoints.USERS_V3}/login_update",
        payload={
            "api_key": self.api_key,
            "email": email,
            "password": password,
            "current_password": current_password,
            "email_grant_token": email_grant_token,
        },
        data_type=LoginUpdateResponse,
        access_token=access_token,
    )
    self.logger.info("Account has been save with email.")
    return response
