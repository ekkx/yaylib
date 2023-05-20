import os
import asyncio
import httpx
import json
import logging
import tqdm
import aiohttp
from huepy import *

from . import models
from . import utils
from .version import version


class YayError(Exception):
    pass


class AuthenticationError(YayError):
    pass


class ForbiddenError(YayError):
    pass


class RateLimitError(YayError):
    pass


current_path = os.path.abspath(os.getcwd())


class Yay(object):

    def __init__(
            self,
            access_token: str = None,
            proxy: str = None,
            timeout=10,
            base_path=current_path,
            loglevel_stream=logging.INFO,
            domain="api.yay.space",
    ):
        """
        hey, this is yaylib.

        """
        self.yay_api_vision = "3.16"
        self.yay_vision_name = "3.16.1"
        self.access_token = access_token
        self.proxy = proxy
        self.proxies = None
        self.timeout = timeout
        self.base_path = base_path
        self.domain = domain
        self.user_agent = ""
        self.headers = {
            'Host': self.domain,
            'X-Timestamp': "",
            'X-App-Version': self.yay_api_vision,
            'X-Device-Info': f"yay {self.yay_vision_name} android 11 (3.5x 1440x2960 Galaxy S9)",
            'X-Device-Uuid': "",
            'X-Connection-Type': 'wifi',
            'Accept-Language': 'ja',
            'Content-Type': 'application/json;charset=UTF-8'
        }

        self.logger = logging.getLogger("yaylib version: " + version)
        ch = logging.StreamHandler()
        ch.setLevel(loglevel_stream)
        ch.setFormatter(logging.Formatter("%(asctime)s - %(levelname)s - %(message)s"))

        handler_existed = False
        for handler in self.logger.handlers:
            if isinstance(handler, logging.StreamHandler):
                handler_existed = True
                break
        if not handler_existed:
            self.logger.addHandler(ch)
        self.logger.setLevel(logging.DEBUG)

        if access_token:
            self.headers.setdefault("Authorization", f"Bearer {access_token}")

        self.logger.info("yaylib version: " + version + " started")
