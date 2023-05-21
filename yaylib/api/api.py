import httpx

from typing import Dict

from ..config import *
from ..errors import *


def _check_authorization(self, headers) -> dict:
    if headers is None:
        headers = self.headers
    if headers.get("Authorization") is None:
        raise AuthenticationError(
            "Authorization is not present in the header.")
    return headers


def _get(self, endpoint: str, params: dict = None, headers: dict = None):
    if headers is None:
        headers = self.headers
    resp = httpx.get(endpoint, params=params, headers=self.headers,
                     proxies=self.proxies, timeout=self.timeout)
    _handle_response(resp)
    return resp.json()


def _post(self, endpoint: str, params: dict = None, headers: dict = None):
    if headers is None:
        headers = self.headers
    resp = httpx.post(endpoint, params=params, headers=self.headers,
                      proxies=self.proxies, timeout=self.timeout)
    _handle_response(resp)
    return resp.json()


def _put(self, endpoint: str, params: dict = None, headers: dict = None):
    if headers is None:
        headers = self.headers
    resp = httpx.put(endpoint, params=params, headers=self.headers,
                     proxies=self.proxies, timeout=self.timeout)
    _handle_response(resp)
    return resp.json()


def _delete(self, endpoint: str, params: dict = None, headers: dict = None):
    if headers is None:
        headers = self.headers
    resp = httpx.delete(endpoint, params=params, headers=self.headers,
                        proxies=self.proxies, timeout=self.timeout)
    _handle_response(resp)
    return resp.json()


def _handle_response(resp):
    if resp.status_code == 401:
        raise AuthenticationError("Failed to authenticate")
    if resp.status_code == 403:
        raise ForbiddenError("Forbidden")
    if resp.status_code == 429:
        raise RateLimitError("Rate limit exceeded")

    resp_json = resp.json()

    if "error_code" in resp_json:
        if resp_json["error_code"] == -343:
            raise ExceedCallQuotaError("Exceed call quota")
        if resp_json["error_code"] == -380:
            raise InvalidSignedInfoError("Invalid signed info")