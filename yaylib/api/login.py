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
from cryptography.fernet import Fernet

from ..config import Configs, Endpoints
from ..errors import AuthenticationError, ForbiddenError
from ..responses import LoginUserResponse, LoginUpdateResponse, TokenResponse
from ..utils import (
    Colors,
    console_print,
    load_credentials,
    save_credentials,
    decrypt,
    signed_info_calculating,
)


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
    self.logger.info(self, fname="Your email has been changed.")
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
    self.logger.info(self, fname="Your password has been changed..")
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
        endpoint=f"{Configs.YAY_API_URL}/api/v1/oauth/token",
        payload={
            "grant_type": grant_type,
            "email": email,
            "password": password,
            "refresh_token": refresh_token,
        },
        data_type=TokenResponse,
        access_token=access_token,
    )


def is_valid_token(self, access_token: str):
    headers = self.session.headers
    headers.setdefault("Authorization", f"Bearer {access_token}")
    try:
        self.get_web_socket_token(headers)
        return True
    except AuthenticationError:
        return False


def login_with_email(
    self, email: str, password: str, secret_key: str = None
) -> LoginUserResponse:
    credentials = load_credentials(base_path=self.base_path, check_email=email)
    if credentials is not None:
        if secret_key is not None:
            self.secret_key = secret_key
            self.fernet = Fernet(secret_key)
            credentials = decrypt(fernet=self.fernet, credentials=credentials)
            self.session.headers.setdefault(
                "Authorization", f"Bearer {credentials['access_token']}"
            )
            self.logger.info(f"Successfully logged in as '{credentials['user_id']}'")
            return credentials
        else:
            message = f"{Colors.WARNING}Credential file found. The 'secret_key' must be provided to decrypt the credentials.{Colors.RESET}"
            console_print(message)

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

    message = "Invalid email or password."
    if response.access_token is None:
        raise ForbiddenError(message)

    self.session.headers.setdefault("Authorization", f"Bearer {response.access_token}")
    self.logger.info(f"Successfully logged in as '{response.user_id}'")

    secret_key = Fernet.generate_key()
    self.secret_key = secret_key
    self.fernet = Fernet(secret_key)

    console_print(
        f"Your 'secret_key' for {Colors.BOLD + email + Colors.RESET} is: {Colors.OKGREEN + secret_key.decode() + Colors.RESET}",
        "Please copy and securely store this key in a safe location.",
        "For more information, visit: https://github.com/qvco/yaylib/blob/master/docs/API-Reference/login/login.md",
    )

    save_credentials(
        base_path=self.base_path,
        fernet=self.fernet,
        access_token=response.access_token,
        refresh_token=response.refresh_token,
        user_id=response.user_id,
        email=email,
    )
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
            "signed_info": signed_info_calculating(self.device_uuid, timestamp),
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
