import httpx
import os
import logging
import uuid

from typing import Optional, Dict, Any

from .config import *
from .version import version
from .api.login import *
from .api.post import *
from .api.users import *


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
        self.yay_api_key = Configs.YAY_API_KEY
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

    # LOGIN

    def login_with_email(self, email: str, password: str, headers: Dict[str, str | int] = None):
        return login_with_email(self, email, password, headers)

    # POST

    def create_post(
            self,
            text: str = None,
            font_size: int = 0,
            color: int = 0,
            in_reply_to: int = None,
            group_id: int = None,
            post_type: str = None,
            mention_ids: List[int] = None,
            choices: List[str] = None,
            shared_url: Dict[str, str | int] = None,
            message_tags: str = "[]",
            attachment_filename: str = None,
            attachment_2_filename: str = None,
            attachment_3_filename: str = None,
            attachment_4_filename: str = None,
            attachment_5_filename: str = None,
            attachment_6_filename: str = None,
            attachment_7_filename: str = None,
            attachment_8_filename: str = None,
            attachment_9_filename: str = None,
            video_file_name: str = None,
            headers: Dict[str, str | int] = None,
    ):
        return create_post(
            self,
            text,
            font_size,
            color,
            in_reply_to,
            group_id,
            post_type,
            mention_ids,
            choices,
            shared_url,
            message_tags,
            attachment_filename,
            attachment_2_filename,
            attachment_3_filename,
            attachment_4_filename,
            attachment_5_filename,
            attachment_6_filename,
            attachment_7_filename,
            attachment_8_filename,
            attachment_9_filename,
            video_file_name,
            headers,
        )

    def get_my_posts(
            self,
            from_post_id: int = None,
            number: int = 100,
            include_group_post: bool = False,
            headers: Dict[str, str | int] = None
    ):
        return get_my_posts(self, from_post_id, number, include_group_post, headers)

    def get_timeline_calls(
        self,
        group_id: int = None,
        from_timestamp: int = None,
        number: int = None,
        category_id: int = None,
        call_type: str = "voice",
        include_circle_call: bool = None,
        cross_generation: bool = None,
        exclude_recent_gomimushi: bool = None,
        shared_interest_categories: bool = None,
        headers: Dict[str, str | int] = None,
    ):
        return get_timeline_calls(
            self,
            group_id,
            from_timestamp,
            number,
            category_id,
            call_type,
            include_circle_call,
            cross_generation,
            exclude_recent_gomimushi,
            shared_interest_categories,
            headers,
        )

    def get_conversation(
        self,
        conversation_id: int,
        group_id: int = None,
        thread_id: int = None,
        from_post_id: int = None,
        number=50,
        reverse: bool = True,
        headers: Dict[str, str | int] = None,
    ):
        return get_conversation(
            self,
            conversation_id,
            group_id,
            thread_id,
            from_post_id,
            number,
            reverse,
            headers,
        )

    # USERS

    def contact_friends(self, headers: Dict[str, str | int] = None):
        return contact_friends(self, headers)
