import requests
import re

from bs4 import BeautifulSoup

from ..config import Endpoints as ep
from ..exceptions import (
    YayError,
    AuthenticationError,
    ForbiddenError,
    RateLimitError,
    ExceedCallQuotaError,
    UnknownError
)
from ..utils import handle_response, console_print


def login(self, email, password):
    try:
        resp = self._post(
            f'{ep.USER_v2}/login_with_email',
            data={
                'email': email,
                'password': password,
                'api_key': self.api_key,
                'uuid': ''
            }
        )

        self.access_token = resp['access_token']
        self.refresh_token = resp['refresh_token']
        self.expires_in = resp['expires_in']
        self.logged_in_as = resp['user_id']

        self.headers.setdefault(
            'Authorization', f'Bearer {self.access_token}')

        self.logger.info(
            'Successfully logged in as [{}]'.format(self.logged_in_as))
        return True

    except ForbiddenError:
        self.logger.error('Login Failed')
        return False


def logout(self):
    if self.access_token:
        self.headers.pop('Authorization', None)
        self.access_token = None
        self.refresh_token = None
        self.expires_in = None
        self.logged_in_as = None
        self.logger.info('User has Logged out')
    else:
        console_print('User is not logged in', 'red')


def get_api_key(self):
    resp = requests.get(
        'https://yay.space/?modalMode=login',
        headers=self.headers, proxies=self.proxies, timeout=self.timeout
    )
    soup = BeautifulSoup(resp.content, 'html.parser')
    script = soup.find_all('script')[2].string
    return re.search(r'gon\.API_KEY="(.+?)"', script).group(1)
