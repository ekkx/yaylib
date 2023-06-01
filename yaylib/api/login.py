from datetime import datetime
from typing import Dict, List

from ..config import *
from ..errors import *
from ..models import *
from ..responses import *
from ..utils import *


def change_email(
        self,
        email: str,
        password: str,
        email_grant_token: str = None
) -> LoginUpdateResponse:
    return self._make_request(
        "PUT", endpoint=f"{Endpoints.USERS_V1}/change_email",
        payload={
            "api_key": self.api_key,
            "email": email,
            "password": password,
            "email_grant_token": email_grant_token
        }, data_type=LoginUpdateResponse
    )


def change_password(
        self,
        current_password: str,
        new_password: str
) -> LoginUpdateResponse:
    return self._make_request(
        "PUT", endpoint=f"{Endpoints.USERS_V1}/change_email",
        payload={
            "api_key": self.api_key,
            "current_password": current_password,
            "password": new_password
        }, data_type=LoginUpdateResponse
    )


def connect_account_with_sns(self):
    pass


def disconnect_account_with_sns(self):
    pass


def get_token(
        self,
        grant_type: str,
        refresh_token: str = None,
        email: str = None,
        password: str = None
) -> TokenResponse:
    return self._make_request(
        "POST", endpoint=f"{self.host}/api/v1/oauth/token",
        payload={
            "grant_type": grant_type,
            "email": email,
            "password": password,
            "refresh_token": refresh_token
        }, data_type=TokenResponse
    )


def login_with_email(self, email: str, password: str) -> LoginUserResponse:
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V3}/login_with_email",
        payload={
            "api_key": self.api_key,
            "email": email,
            "password": password,
            "uuid": self.uuid
        }, data_type=LoginUserResponse
    )
    message = "Invalid email or password."
    if response.access_token is None:
        raise ForbiddenError(message)

    self.login_data = response
    self.session.headers.setdefault(
        "Authorization", f"Bearer {self.login_data.access_token}"
    )

    self.logger.info(f'Successfully logged in as [{self.login_data.user_id}]')
    return response


def login_with_sns(self):
    pass


def logout(self):
    try:
        self._check_authorization()
        response = self._make_request(
            "POST", endpoint=f"{Endpoints.USERS_V1}/logout",
            payload={"uuid": self.uuid}
        )
        self.session.headers.pop('Authorization', None)
        self.login_data = None
        self.logger.info('User has Logged out.')
        return response

    except:
        self.logger.error(f'User is not logged in.')
        return None


def migrate_token(self):
    pass


def register_device_token(self):
    pass


def resend_confirm_email(self):
    return self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V2}/resend_confirm_email"
    )


def restore_user(self, user_id: int) -> LoginUserResponse:
    return self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V2}/restore",
        payload={
            "user_id": user_id,
            "api_key": self.api_key,
            "uuid": self.uuid,
            "timestamp": int(datetime.now().timestamp()),
            "signed_info": signed_info_calculating(
                self.api_key, self.device_uuid,
                int(datetime.now().timestamp())
            ),
        }
    )


def revoke_tokens(self):
    return self._make_request(
        "DELETE", endpoint=f"{Endpoints.USERS_V1}/device_tokens"
    )


def save_account_with_email(
        self,
        email: str,
        password: str = None,
        current_password: str = None,
        email_grant_token: str = None
) -> LoginUpdateResponse:
    return self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V3}/login_update",
        payload={
            "api_key": self.api_key,
            "email": email,
            "password": password,
            "current_password": current_password,
            "email_grant_token": email_grant_token
        }, data_type=LoginUpdateResponse
    )
