from typing import Union, Dict, List

from .api import API
from .api.call import *
from .api.chat import *
from .api.group import *
from .api.login import *
from .api.misc import *
from .api.post import *
from .api.review import *
from .api.thread import *
from .api.user import *

from .config import *
from .errors import *
from .models import *
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

    # MISC

    def get_web_socket_token(self) -> str:
        return get_web_socket_token(self)

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
            shared_url: Dict[str, Union[str, int]] = None,
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

    def get_my_posts(self, **kwargs: Union[int, bool]) -> PostsResponse:
        """

        Parameters:
        ---------------

            - from_post_id: int - (optional)
            - number: int - (optional)
            - include_group_post: bool - (optional)

        """
        return get_my_posts(self, **kwargs)

    def get_post(self, post_id: int) -> Post:
        return get_post(self, post_id)

    def get_post_likers(
            self,
            post_id: int,
            from_id: int = None,
            number: int = None
    ) -> PostLikersResponse:
        return get_post_likers(self, post_id, from_id, number)

    def get_post_reposts(self, post_id: int, **kwargs: int) -> PostsResponse:
        """

        Parameters:
        ---------------

            - post_id: int - (required)
            - from_post_id: int - (optional)
            - number: int - (optional)

        """
        return get_post_reposts(self, post_id, **kwargs)

    def get_posts(self, post_ids: List[int]) -> List[Post]:
        return get_posts(self, post_ids)

    def get_timeline(self, **kwargs: Union[int, str, bool]) -> PostsResponse:
        # noreply_mode: str = None
        # - from: str - (optional)
        """

        Parameters:
        ---------------

            - from_post_id: int - (optional)
            - number: int - (optional)
            - order_by: str - (optional)
            - experiment_older_age_rules: bool - (optional)
            - shared_interest_categories: bool - (optional)
            - mxn: int - (optional)
            - en: int - (optional)
            - vn: int - (optional)
            - reduce_selfie: bool - (optional)
            - custom_generation_range: bool - (optional)

        """
        return get_timeline(self, **kwargs)

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
        headers: Dict[str, Union[str, int]] = None,
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
        headers: Dict[str, Union[str, int]] = None,
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

    def get_user(self, user_id: int) -> User:
        return get_user(self, user_id)
