import os
import asyncio
import httpx
import json
import logging
import tqdm

from fake_useragent import UserAgent
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

        self.access_token = access_token
        self.proxy = proxy
        self.proxies = None
        self.timeout = timeout
        self.base_path = base_path
        self.domain = domain
        self.user_agent = UserAgent().chrome
        self.headers = {
            
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
