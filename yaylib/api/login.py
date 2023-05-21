from typing import Dict, List

from ..config import *
from ..errors import *
from .api import (
    _get,
    _post,
    _put,
    _delete,
    _check_authorization,
    _handle_response,
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


def logout(self, headers: Dict[str, str | int] = None):
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