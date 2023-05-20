import httpx
from typing import (Dict)

from ..config import (Endpoints, Configs)
from ..errors import (
    YayError,
    ForbiddenError,
    RateLimitError,
    AuthenticationError,
)


def _get(self, endpoint: str, data: dict = None, headers: dict = None):
    if headers is None:
        headers = self.headers
    resp = httpx.get(endpoint, params=data, headers=self.headers, proxies=self.proxies, timeout=self.timeout)
    _handle_response(resp)
    return resp.json()


def _post(self, endpoint: str, data: dict = None, headers: dict = None):
    if headers is None:
        headers = self.headers
    resp = httpx.post(endpoint, params=data, headers=self.headers, proxies=self.proxies, timeout=self.timeout)
    _handle_response(resp)
    return resp.json()


def _put(self, endpoint: str, data: dict = None, headers: dict = None):
    if headers is None:
        headers = self.headers
    resp = httpx.put(endpoint, params=data, headers=self.headers, proxies=self.proxies, timeout=self.timeout)
    _handle_response(resp)
    return resp.json()


def _delete(self, endpoint: str, data: dict = None, headers: dict = None):
    if headers is None:
        headers = self.headers
    resp = httpx.delete(endpoint, params=data, headers=self.headers, proxies=self.proxies, timeout=self.timeout)
    _handle_response(resp)
    return resp.json()


def _check_authorization(self, headers):        
    if self.headers.get("Authorization") is None:
        raise AuthenticationError("Authorization is not present in the header.")


def _handle_response(resp):
    return resp