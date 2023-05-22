from typing import Dict, List

from ..config import *
from ..errors import *


def change_email(self):
    pass


def change_password(self):
    pass


def connect_account_with_sns(self):
    pass


def disconnect_account_with_sns(self):
    pass


def get_token(self, grant_type: str, refresh_token: str = None, email: str = None, password: str = None):
    return self._make_request(
        "POST", endpoint=f"https://{self.host}/api/v1/oauth/token",
        payload={
            "grant_type": grant_type,
            "email": email,
            "password": password,
            "refresh_token": refresh_token
        }
    )


def login_with_email(self, email: str, password: str):
    try:
        response = self._make_request(
            "POST", endpoint=f"https://{Endpoints.USER_V3}/login_with_email",
            payload={
                "api_key": self.yay_api_key,
                "email": email,
                "password": password,
                "uuid": self.uuid
            }
        )
        self.access_token = response["access_token"]
        self.refresh_token = response["refresh_token"]
        self.logged_in_as = response["user_id"]
        self.headers.setdefault('Authorization', f'Bearer {self.access_token}')

        self.logger.info(f'Successfully logged in as [{self.logged_in_as}]')
        return response

    except:
        self.logger.error('Login failed.')
        return None


def login_with_sns(self):
    pass


def logout(self):
    try:
        self._check_authorization()
        response = self._make_request(
            "POST", endpoint=f"https://{Endpoints.USER_V1}/logout",
            payload={"uuid": self.uuid}
        )
        self.headers.pop('Authorization', None)
        self.access_token = None
        self.refresh_token = None
        self.logged_in_as = None
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
    pass


def restore_user(self):
    pass


def revoke_tokens(self):
    pass


def save_account_with_email(self):
    pass
