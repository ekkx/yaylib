import requests
import re

from fake_useragent import UserAgent
from bs4 import BeautifulSoup

from endpoints import Endpoints as ep


class YayAuth(object):

    def __init__(self, proxy=None, timeout=10):
        self.timeout = timeout
        self.user_agent = UserAgent().chrome
        self.proxy = f'http://{proxy}' if proxy else None
        self.headers = {
            'User-Agent': self.user_agent,
            'Accept': 'application/json, text/plain, */*',
            'Accept-Language': 'ja',
            'Referer': 'https://yay.space/',
            'Content-Type': 'application/json;charset=utf-8',
            'Agent': 'YayWeb 3.12.1',
            'X-Device-Info': f'Yay 3.12.1 Web ({self.user_agent})',
            'Origin': 'https://yay.space'
        }
        self.api_key = None
        self.access_token = None
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
            proxies={'http': self.proxy, 'https': self.proxy},
            timeout=self.timeout
        )
        resp.raise_for_status()

        self.access_token = resp.json()['access_token']
        self.logged_in_as = resp.json()['user_id']
        if self.access_token:
            self.headers.setdefault(
                'Authorization', f'Bearer {self.access_token}'
            )
            print('Login Successful.')
        else:
            print("Login Failed.")

    def logout(self):
        if not self.access_token:
            print("User is not logged in.")
            return

        self.headers.pop('Authorization', None)
        self.access_token = None
        self.logged_in_as = None
        print('Logout Successful.')
