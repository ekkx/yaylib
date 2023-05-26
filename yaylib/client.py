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

    #### Yay! API v3 Client

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

    def add_bookmark(self, user_id: int, post_id: int) -> BookmarkPostResponse:
        return add_bookmark(self, user_id, post_id)

    def add_group_highlight_post(self, group_id: int, post_id: int):
        return add_group_highlight_post(self, group_id, post_id)

    def create_call_post(
            self, *, text: str = None, font_size: int = None, color: int = None, group_id: int = None,
            call_type: str = None, category_id: int = None, game_title: str = None,
            joinable_by: str = None, message_tags: str = "[]", attachment_filename: str = None,
            attachment_2_filename: str = None, attachment_3_filename: str = None,
            attachment_4_filename: str = None, attachment_5_filename: str = None,
            attachment_6_filename: str = None, attachment_7_filename: str = None,
            attachment_8_filename: str = None, attachment_9_filename: str = None,
    ) -> ConferenceCall:
        return create_call_post(
            self, text, font_size, color, group_id, call_type, category_id,
            game_title, joinable_by, message_tags, attachment_filename,
            attachment_2_filename, attachment_3_filename, attachment_4_filename,
            attachment_5_filename, attachment_6_filename, attachment_7_filename,
            attachment_8_filename, attachment_9_filename
        )

    def create_group_pin_post(self, post_id: int, group_id: int):
        return create_group_pin_post(self, post_id, group_id)

    def create_pin_post(self, post_id: int):
        return create_pin_post(self, post_id)

    def create_post(
            self, *, text: str = None, font_size: int = 0, color: int = 0, in_reply_to: int = None,
            group_id: int = None, post_type: str = None, mention_ids: List[int] = None,
            choices: List[str] = None, shared_url: Dict[str, str | int] = None, message_tags: str = "[]",
            attachment_filename: str = None, attachment_2_filename: str = None,
            attachment_3_filename: str = None, attachment_4_filename: str = None,
            attachment_5_filename: str = None, attachment_6_filename: str = None,
            attachment_7_filename: str = None, attachment_8_filename: str = None,
            attachment_9_filename: str = None, video_file_name: str = None,
    ) -> Post:
        return create_post(
            self, text, font_size, color, in_reply_to, group_id, post_type,
            mention_ids, choices, shared_url, message_tags, attachment_filename,
            attachment_2_filename, attachment_3_filename, attachment_4_filename,
            attachment_5_filename, attachment_6_filename, attachment_7_filename,
            attachment_8_filename, attachment_9_filename, video_file_name
        )

    def create_repost(
            self, post_id: int, *, text: str = None, font_size: int = None, color: int = None,
            in_reply_to: int = None, group_id: int = None, post_type: str = None, mention_ids: List[int] = None,
            choices: List[str] = None, shared_url: Dict[str, Union[str, int]] = None, message_tags: str = "[]",
            attachment_filename: str = None, attachment_2_filename: str = None,
            attachment_3_filename: str = None, attachment_4_filename: str = None,
            attachment_5_filename: str = None, attachment_6_filename: str = None,
            attachment_7_filename: str = None, attachment_8_filename: str = None,
            attachment_9_filename: str = None, video_file_name: str = None,
    ) -> Post:
        return create_repost(
            self, post_id, text, font_size, color, in_reply_to, group_id, post_type,
            mention_ids, choices, shared_url, message_tags, attachment_filename,
            attachment_2_filename, attachment_3_filename, attachment_4_filename,
            attachment_5_filename, attachment_6_filename, attachment_7_filename,
            attachment_8_filename, attachment_9_filename, video_file_name,
        )

    def create_share_post(
            self, shareable_type: str, shareable_id: int, *, text: str = None,
            font_size: int = None, color: int = None, group_id: int = None,
    ) -> Post:
        return create_share_post(
            self, shareable_type, text, font_size, color, group_id
        )

    def create_thread_post(
        self, post_id: int, *, text: str = None, font_size: int = None, color: int = None,
        in_reply_to: int = None, group_id: int = None, post_type: str = None, mention_ids: List[int] = None,
        choices: List[str] = None, shared_url: Dict[str, Union[str, int]] = None, message_tags: str = "[]",
        attachment_filename: str = None, attachment_2_filename: str = None,
        attachment_3_filename: str = None, attachment_4_filename: str = None,
        attachment_5_filename: str = None, attachment_6_filename: str = None,
        attachment_7_filename: str = None, attachment_8_filename: str = None,
        attachment_9_filename: str = None, video_file_name: str = None,
    ) -> Post:
        return create_thread_post(
            self, post_id, text, font_size, color, in_reply_to, group_id, post_type,
            mention_ids, choices, shared_url, message_tags, attachment_filename,
            attachment_2_filename, attachment_3_filename, attachment_4_filename,
            attachment_5_filename, attachment_6_filename, attachment_7_filename,
            attachment_8_filename, attachment_9_filename, video_file_name,
        )

    def delete_all_post(self):
        return delete_all_post(self)

    def delete_group_pin_post(self, group_id: int):
        return delete_group_pin_post(self, group_id)

    def delete_pin_post(self, post_id: int):
        return delete_pin_post(self, post_id)

    def get_bookmark(self, user_id: int, *, from_str: str = None) -> PostsResponse:
        return get_bookmark(self, user_id, from_str)

    def get_timeline_calls(self, **params) -> PostsResponse:
        """

        Parameters:
        ----------

            - group_id: int = None
            - from_timestamp: int = None
            - number: int = None
            - category_id: int = None
            - call_type: str = "voice"
            - include_circle_call: bool = None
            - cross_generation: bool = None
            - exclude_recent_gomimushi: bool = None
            - shared_interest_categories: bool = None

        """
        return get_timeline_calls(self, **params)

    def get_conversation(self, conversation_id: int, **params) -> PostsResponse:
        """

        Parameters:
        ----------

            - conversation_id: int
            - group_id: int = None
            - thread_id: int = None
            - from_post_id: int = None
            - number: int = 50
            - reverse: bool = True

        """
        return get_conversation(self, conversation_id, **params)

    def get_conversation_root_posts(self, post_ids: List[int]) -> PostsResponse:
        return get_conversation_root_posts(self, post_ids)

    def get_following_call_timeline(self, **params) -> PostsResponse:
        """

        Parameters:
        ----------

            - from_timestamp: int = None
            - number: int = None
            - category_id: int = None
            - call_type: str = None
            - include_circle_call: bool = None
            - exclude_recent_gomimushi: bool = None

        """
        return get_following_call_timeline(self, **params)

    def get_following_timeline(self, **params) -> PostsResponse:
        """

        Parameters:
        ----------

            - from_str: str = None
            - only_root: bool = None
            - order_by: str = None
            - number: int = None
            - mxn: int = None
            - reduce_selfie: bool = None
            - custom_generation_range: bool = None

        """
        return get_following_timeline(self, **params)

    def get_group_highlight_posts(self, group_id: int, **params) -> PostsResponse:
        """

        Parameters:
        ----------

            - group_id: int
            - from_post: int = None
            - number: int = None

        """
        return get_group_highlight_posts(self, group_id, **params)

    def get_group_timeline_by_keyword(self, group_id: int, keyword: str, **params) -> PostsResponse:
        """

        Parameters:
        ----------

            - group_id: int
            - keyword: str
            - from_post_id: int = None
            - number: int = None
            - only_thread_posts: bool = False

        """
        return get_group_timeline_by_keyword(self, group_id, keyword, **params)

    def get_group_timeline(self, group_id: int, **params) -> PostsResponse:
        """

        Parameters:
        ----------

            - group_id: int
            - from_post_id: int
            - reverse: bool
            - post_type: str
            - number: int
            - only_root: bool

        """
        return get_group_timeline(self, group_id, **params)

    def get_timeline_by_hashtag(self, hashtag: str, **params) -> PostsResponse:
        """

        Parameters:
        ----------

            - hashtag: str - (required)
            - from_post_id: int - (optional)
            - number: int - (optional)

        """
        return get_timeline_by_hashtag(self, hashtag, **params)

    def get_my_posts(self, **params) -> PostsResponse:
        """

        Parameters:
        ---------------

            - from_post_id: int - (optional)
            - number: int - (optional)
            - include_group_post: bool - (optional)

        """
        return get_my_posts(self, **params)

    def get_post(self, post_id: int) -> Post:
        return get_post(self, post_id)

    def get_post_likers(self, post_id: int, **params) -> PostLikersResponse:
        """

        Parameters:
        ---------------

            - from_id: int - (optional)
            - number: int - (optional)

        """
        return get_post_likers(self, post_id, **params)

    def get_post_reposts(self, post_id: int, **params: int) -> PostsResponse:
        """

        Parameters:
        ---------------

            - post_id: int - (required)
            - from_post_id: int - (optional)
            - number: int - (optional)

        """
        return get_post_reposts(self, post_id, **params)

    def get_posts(self, post_ids: List[int]) -> PostsResponse:
        return get_posts(self, post_ids)

    def get_recommended_post_tags(
        self, tag: str = None, save_recent_search: bool = False
    ) -> PostTagsResponse:
        return get_recommended_post_tags(self, tag, save_recent_search)

    def get_recommended_posts(self, **params) -> PostsResponse:
        """

        Parameters:
        ---------------

            - experiment_num: int
            - variant_num: int
            - number: int

        """
        return get_recommended_posts(self, **params)

    def get_timeline_by_keyword(self, keyword: str = None, **params) -> PostsResponse:
        """

        Parameters:
        ---------------

            - keyword: str
            - from_post_id: int
            - number: int

        """
        return get_timeline_by_keyword(self, keyword, **params)

    def get_timeline(self, **params: int | str | bool) -> PostsResponse:
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
        return get_timeline(self, **params)

    def get_url_metadata(self, url: str) -> SharedUrl:
        return get_url_metadata(self, url)

    def get_user_timeline(self, user_id: int, **params) -> PostsResponse:
        """

        Parameters:
        ---------------

            - from_post_id: int - (optional)
            - number: int - (optional)
            - post_type: str - (optional)

        """
        return get_user_timeline(self, user_id, **params)

    def like_posts(self, post_ids: List[int]) -> LikePostsResponse:
        return like_posts(self, post_ids)

    def remove_bookmark(self, user_id: int, post_id: int):
        return remove_bookmark(self, user_id, post_id)

    def remove_group_highlight_post(self, group_id: int, post_id: int):
        return remove_group_highlight_post(self, group_id, post_id)

    def remove_posts(self, post_ids: List[int]):
        return remove_posts(self, post_ids)

    def report_post(
        self, post_id: int, opponent_id: int, category_id: int, *, reason: str = None,
        screenshot_filename: str = None, screenshot_2_filename: str = None,
        screenshot_3_filename: str = None, screenshot_4_filename: str = None
    ):
        return report_post(
            self, post_id, opponent_id, category_id, reason, screenshot_filename,
            screenshot_2_filename, screenshot_3_filename, screenshot_4_filename
        )

    def unlike_post(self, post_id: int):
        return unlike_post(self, post_id)

    def update_post(
        self, post_id: int, text: str = None, font_size: int = None,
        color: int = None, message_tags: str = "[]",
    ) -> Post:
        return update_post(
            self, post_id, text, font_size, color, message_tags
        )

    def update_recommendation_feedback(
        self, post_id: int, feedback_result: str, *,
        experiment_num: int, variant_num: int,
    ):
        return update_recommendation_feedback(
            self, post_id, feedback_result, experiment_num, variant_num
        )

    def validate_post(self, text: str, *, group_id: int = None, thread_id: int = None) -> ValidationPostResponse:
        return validate_post(self, text, group_id, thread_id)

    def view_video(self, video_id: int):
        return view_video(self, video_id)

    def vote_survey(self, survey_id: int, choice_id: int) -> Survey:
        return vote_survey(self, survey_id, choice_id)

    # REVIEW

    # THREAD

    # USER

    def get_user(self, user_id: int) -> User:
        return get_user(self, user_id)
