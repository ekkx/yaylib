import requests
import re

from fake_useragent import UserAgent
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


class YayAuth(object):

    def __init__(self, proxy=None, timeout=10):
        self.timeout = timeout
        self.user_agent = UserAgent().chrome
        self.proxy = proxy
        self.proxies = None
        if proxy:
            self.proxies = {
                'http': f'http://{proxy}',
                'https': f'https://{proxy}'
            }
        self.headers = {
            'User-Agent': self.user_agent,
            'Accept': 'application/json, text/plain, */*',
            'Accept-Language': 'ja',
            'Referer': 'https://yay.space/',
            'Content-Type': 'application/json;charset=utf-8',
            'Agent': 'YayWeb',
            'X-Device-Info': f'Yay {self.user_agent}',
            'Origin': 'https://yay.space'
        }
        self.access_token = None
        self.refresh_token = None
        self.expires_in = None
        self.logged_in_as = None

    def login(self, email, password):
        resp = requests.get(
            'https://yay.space/?modalMode=login',
            headers=self.headers,
            proxies={'http': self.proxy, 'https': self.proxy},
            timeout=self.timeout
        )

        soup = BeautifulSoup(resp.content, 'html.parser')
        script = soup.find_all('script')[2].string
        self.api_key = re.search(r'gon\.API_KEY="(.+?)"', script).group(1)

        resp = requests.post(
            f'{ep.USER_v2}/login_with_email',
            params={
                'email': email,
                'password': password,
                'api_key': self.api_key,
                'uuid': ''
            },
            headers=self.headers,
            proxies=self.proxies,
            timeout=self.timeout
        )

        try:
            handle_response(resp)

            self.access_token = resp.json()['access_token']
            self.refresh_token = resp.json()['refresh_token']
            self.expires_in = resp.json()['expires_in']
            self.logged_in_as = resp.json()['user_id']

            self.headers.setdefault(
                'Authorization', f'Bearer {self.access_token}'
            )

            console_print(
                f'Successfully logged in as {self.logged_in_as}.', 'green')
            return True

        except ForbiddenError:
            console_print(
                'Login Failed...\n(Invalid email or password.)', 'red')
            return False

    def logout(self):
        if self.access_token:
            self.headers.pop('Authorization', None)
            self.access_token = None
            self.refresh_token = None
            self.expires_in = None
            self.logged_in_as = None
        else:
            console_print('User is not logged in.', 'red')
