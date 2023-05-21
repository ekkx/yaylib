import httpx
import os
import logging
import uuid

from http import HTTPStatus
from typing import Optional, Dict, Any

from .config import *
from .version import version
from .api.post import (
    get_my_posts,
)
from .api.users import (
    contact_friends,
)


current_path = os.path.abspath(os.getcwd())


class Client(object):

    def __init__(
            self,
            access_token: str = None,
            proxy: str = None,
            timeout=10,
            base_path=current_path,
            loglevel_stream=logging.INFO,
            domain=Configs.YAY_PRODUCTION_HOST,
    ):
        """
        hey, this is yaylib.

        """
        self.yay_api_version = Configs.YAY_API_VERSION
        self.yay_version_name = Configs.YAY_VERSION_NAME
        self.access_token = access_token
        self.proxy = proxy
        self.proxies = None
        self.timeout = timeout
        self.base_path = base_path
        self.domain = domain
        self.uuid = str(uuid.uuid4())
        self.headers = {
            "Host": Configs.YAY_PRODUCTION_HOST,
            "X-App-Version": self.yay_api_version,
            "X-Device-Info": f"yay {self.yay_version_name} android 11 (3.5x 1440x2960 Galaxy S9)",
            "X-Device-Uuid": self.uuid,
            "X-Connection-Type": "wifi",
            "Accept-Language": "ja",
            "Content-Type": "application/json;charset=UTF-8"
        }
        self.logger = logging.getLogger("yaylib version: " + version)
        ch = logging.StreamHandler()
        ch.setLevel(loglevel_stream)
        ch.setFormatter(logging.Formatter(
            "%(asctime)s - %(levelname)s - %(message)s"))

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

    def get_my_posts(self, from_post_id: int = None, number: int = 100, include_group_post: bool = False, headers: Dict[str, str | int] = None):
        return get_my_posts(self, from_post_id, number, include_group_post, headers)

    def contact_friends(self, headers: Dict[str, str | int] = None):
        return contact_friends(self, headers)
