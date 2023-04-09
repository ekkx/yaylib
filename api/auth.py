import requests
import re


from fake_useragent import UserAgent
from bs4 import BeautifulSoup


from endpoints import Endpoints as ep


class YayAuth(object):
    def __init__(self, proxy=None, timeout=10):
        self.timeout = timeout
        self.user_agent = str(UserAgent()['google chrome'])
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
        self.access_token = None
        self.api_key = None
        self.logged_in_as = None

    def login(self, email, password):
        response = requests.get(
            'https://yay.space/?modalMode=login',
            headers=self.headers,
            proxies={'http': self.proxy, 'https': self.proxy},
            timeout=self.timeout
        )

        soup = BeautifulSoup(response.content, 'html.parser')
        script = soup.find_all('script')[2].string
        self.api_key = re.search(r'gon\.API_KEY="(.+?)"', script).group(1)

        response = requests.post(
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

        if response.status_code == 429:
            print('Rate limit exceeded')
        elif response.status_code == 403:
            print('Invalid password or email')
        else:
            self.access_token = response.json()['access_token']
            self.logged_in_as = response.json()['user_id']
            self.headers.setdefault(
                'Authorization', f'Bearer {self.access_token}'
            )
            print('Login Successful.')

    def logout(self):
        if self.access_token:
            self.headers.pop('Authorization', None)
            self.access_token = None
            self.logged_in_as = None
        else:
            print("User is not logged in.")
