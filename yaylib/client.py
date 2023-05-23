from typing import Optional, Dict, List, Any

from .api import API
from .api.call import *
from .api.chat import *
from .api.group import *
from .api.login import *
from .api.post import *
from .api.review import *
from .api.thread import *
from .api.user import *

from .config import *
from .utils import *


class Client(API):
    """
    Client( \
        access_token=None, proxy=None, timeout=10, \
        base_path=current_path, loglevel_stream=logging.INFO, \
        host=Configs.YAY_PRODUCTION_HOST \
    )

    ---

    Yay! API v3 Client

    """

    # CALL

    # CHAT

    # GROUP

    # LOGIN

    def get_token(self, grant_type="refresh_token", refresh_token: str = None, email: str = None, password: str = None):
        return get_token(self, grant_type, refresh_token, email, password)

    def login_with_email(self, email: str, password: str):
        return login_with_email(self, email, password)

    def logout(self):
        return logout(self)

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
            video_file_name: str = None
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
        )

    def get_my_posts(
            self,
            from_post_id: int = None,
            number: int = 100,
            include_group_post: bool = False,
            headers: Dict[str, str | int] = None
    ):
        return get_my_posts(self, from_post_id, number, include_group_post, headers)

    def get_post(self, post_id: int) -> Post:
        return get_post(self, post_id)

    def get_posts(self, post_ids: List[int]) -> List[Post]:
        return get_posts(self, post_ids)

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
    
    # REVIEW

    # THREAD

    # USER

    def contact_friends(self):
        return contact_friends(self)

    def get_user(self, user_id: int) -> User:
        return get_user(self, user_id)
