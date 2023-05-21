from typing import Dict, List

from ..config import *
from ..errors import *
from .api import (
    _check_authorization,
    _get,
    _post,
    _put,
    _delete,
    _handle_response,
)


def change_email(self):
    pass


def change_password(self):
    pass


def connect_account_with_sns(self):
    pass


def disconnect_account_with_sns(self):
    pass


def get_token(self, grant_type: str, email: str = None, password: str = None, refresh_token: str = None):
    headers = _check_authorization(self, headers)
    params = {
        "grant_type": grant_type,
        "email": email,
        "password": password,
        "refresh_token": refresh_token
    }
    return _post(
        self=self,
        endpoint=f"https://{self.domain}/api/v1/oauth/token",
        params=params,
        headers=headers
    )


def login_with_email(self, email: str, password: str, headers: Dict[str, str | int] = None):
    try:
        resp = _post(
            self=self,
            endpoint=f"https://{Endpoints.USER_V3}/login_with_email",
            params={
                "api_key": self.yay_api_key,
                "email": email,
                "password": password,
                "uuid": self.uuid
            },
            headers=headers
        )
        self.access_token = resp["access_token"]
        self.refresh_token = resp["refresh_token"]
        self.logged_in_as = resp["user_id"]
        self.headers.setdefault('Authorization', f'Bearer {self.access_token}')

        self.logger.info(f'Successfully logged in as [{self.logged_in_as}]')
        return resp

    except:
        self.logger.error('Login failed.')
        return None


def login_with_sns(self):
    pass


def logout(self, headers: Dict[str, str | int] = None):
    try:
        headers = _check_authorization(self, headers)
        resp = _post(
            self=self,
            endpoint=f"https://{Endpoints.USER_V1}/logout",
            params={"uuid": self.uuid},
            headers=headers
        )
        self.headers.pop('Authorization', None)
        self.access_token = None
        self.refresh_token = None
        self.logged_in_as = None
        self.logger.info('User has Logged out.')
        return resp

    except:
        self.logger.error(f'User is not logged in.')
        return None


def migrate_token(self):
    pass


def register_device_token(self):
    pass


def resend_confirm_email(self):
    pass


def restore_user(self):
    pass


def revoke_tokens(self):
    pass


def save_account_with_email(self):
    pass
