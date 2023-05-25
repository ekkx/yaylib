from datetime import datetime
from typing import Dict, List

from ..config import *
from ..errors import *
from ..models import *
from ..utils import *


def add_bookmark(self, user_id: int, post_id: int) -> BookmarkPostResponse:
    self._check_authorization()
    return self._make_request(
        "PUT", endpoint=f"{Endpoints.USER_V1}/{user_id}/bookmarks/{post_id}",
        data_type=BookmarkPostResponse
    )


def add_group_highlight_post(self, group_id: int, post_id: int):
    self._check_authorization()
    return self._make_request(
        "PUT", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/highlights/{post_id}",
    )


def create_call_post(
        self,
        text: str = None,
        font_size: int = None,
        color: int = None,
        group_id: int = None,
        call_type: str = None,
        category_id: int = None,
        game_title: str = None,
        joinable_by: str = None,
        message_tags: str = "[]",
        attachment_filename: str = None,
        attachment_2_filename: str = None,
        attachment_3_filename: str = None,
        attachment_4_filename: str = None,
        attachment_5_filename: str = None,
        attachment_6_filename: str = None,
        attachment_7_filename: str = None,
        attachment_8_filename: str = None,
        attachment_9_filename: str = None
) -> ConferenceCall:
    # TODO: @NotNull "uuid", "api_key", "timestamp", "signed_info"
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.POSTS_V2}/new_conference_call",
        payload={
            "text": text,
            "font_size": font_size,
            "color": color,
            "group_id": group_id,
            "call_type": call_type,
            "uuid": self.uuid,
            "api_key": self.api_key,
            "timestamp": int(datetime.now().timestamp()),
            "signed_info": signed_info_calculating(self.api_key, self.device_uuid, int(datetime.now().timestamp())),
            "category_id": category_id,
            "game_title": game_title,
            "joinable_by": joinable_by,
            "message_tags": message_tags,
            "attachment_filename": attachment_filename,
            "attachment_2_filename": attachment_2_filename,
            "attachment_3_filename": attachment_3_filename,
            "attachment_4_filename": attachment_4_filename,
            "attachment_5_filename": attachment_5_filename,
            "attachment_6_filename": attachment_6_filename,
            "attachment_7_filename": attachment_7_filename,
            "attachment_8_filename": attachment_8_filename,
            "attachment_9_filename": attachment_9_filename,
        }, data_type=CreatePostResponse
    )
    return response.conference_call


def create_group_pin_post(self, post_id: int, group_id: int):
    self._check_authorization()
    return self._make_request(
        "PUT", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/highlights/{post_id}",
        payload={"post_id": post_id, "group_id": group_id}
    )


def create_pin_post(self, post_id: int):
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.PINNED_V1}/posts",
        payload={"id": post_id}
    )


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
) -> Post:
    self._check_authorization()
    headers = self.session.headers
    headers["X-Jwt"] = self.get_web_socket_token()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.POSTS_V3}/new",
        payload={
            "text": text,
            "font_size": font_size,
            "color": color,
            "in_reply_to": in_reply_to,
            "group_id": group_id,
            "post_type": post_type,
            "mention_ids[]": mention_ids,
            "choices[]": choices,
            "shared_url": shared_url,
            "message_tags": message_tags,
            "attachment_filename": attachment_filename,
            "attachment_2_filename": attachment_2_filename,
            "attachment_3_filename": attachment_3_filename,
            "attachment_4_filename": attachment_4_filename,
            "attachment_5_filename": attachment_5_filename,
            "attachment_6_filename": attachment_6_filename,
            "attachment_7_filename": attachment_7_filename,
            "attachment_8_filename": attachment_8_filename,
            "attachment_9_filename": attachment_9_filename,
            "video_file_name": video_file_name,
        }, data_type=Post, headers=headers
    )


def create_repost(
        self,
        post_id: int,
        text: str = None,
        font_size: int = None,
        color: int = None,
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
) -> Post:
    self._check_authorization()
    headers = self.session.headers
    headers["X-Jwt"] = self.get_web_socket_token()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.POSTS_V3}/repost",
        payload={
            "post_id": post_id,
            "text": text,
            "font_size": font_size,
            "color": color,
            "in_reply_to": in_reply_to,
            "group_id": group_id,
            "post_type": post_type,
            "mention_ids[]": mention_ids,
            "choices[]": choices,
            "shared_url": shared_url,
            "message_tags": message_tags,
            "attachment_filename": attachment_filename,
            "attachment_2_filename": attachment_2_filename,
            "attachment_3_filename": attachment_3_filename,
            "attachment_4_filename": attachment_4_filename,
            "attachment_5_filename": attachment_5_filename,
            "attachment_6_filename": attachment_6_filename,
            "attachment_7_filename": attachment_7_filename,
            "attachment_8_filename": attachment_8_filename,
            "attachment_9_filename": attachment_9_filename,
            "video_file_name": video_file_name,
        }, data_type=CreatePostResponse, headers=headers
    )
    return response.post


def create_share_post(
        self,
        shareable_type: str,
        shareable_id: int,
        text: str = None,
        font_size: int = None,
        color: int = None,
        group_id: int = None,
) -> Post:
    # TODO: @NotNull "uuid", "api_key", "timestamp", "signed_info"
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.POSTS_V2}/new_share_post",
        payload={
            "shareable_type": shareable_type,
            "shareable_id": shareable_id,
            "text": text,
            "font_size": font_size,
            "color": color,
            "group_id": group_id,
            "uuid": self.uuid,
            "api_key": self.api_key,
            "timestamp": int(datetime.now().timestamp()),
            "signed_info": signed_info_calculating(self.api_key, self.device_uuid, int(datetime.now().timestamp())),
        }, data_type=Post
    )


def create_thread_post(
        self,
        post_id: int,
        text: str = None,
        font_size: int = None,
        color: int = None,
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
) -> Post:
    self._check_authorization()
    headers = self.session.headers
    headers["X-Jwt"] = self.get_web_socket_token()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.THREADS_V1}/posts",
        payload={
            "id": post_id,
            "text": text,
            "font_size": font_size,
            "color": color,
            "in_reply_to": in_reply_to,
            "group_id": group_id,
            "post_type": post_type,
            "mention_ids[]": mention_ids,
            "choices[]": choices,
            "shared_url": shared_url,
            "message_tags": message_tags,
            "attachment_filename": attachment_filename,
            "attachment_2_filename": attachment_2_filename,
            "attachment_3_filename": attachment_3_filename,
            "attachment_4_filename": attachment_4_filename,
            "attachment_5_filename": attachment_5_filename,
            "attachment_6_filename": attachment_6_filename,
            "attachment_7_filename": attachment_7_filename,
            "attachment_8_filename": attachment_8_filename,
            "attachment_9_filename": attachment_9_filename,
            "video_file_name": video_file_name,
        }, data_type=Post, headers=headers
    )


def delete_all_post(self):
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.POSTS_V1}/delete_all_post",
    )


def delete_group_pin_post(self, group_id: int):
    self._check_authorization()
    return self._make_request(
        "DELETE", endpoint=f"{Endpoints.POSTS_V2}/group_pinned_post",
        payload={"group_id": group_id}
    )


def delete_pin_post(self, post_id: int):
    self._check_authorization()
    return self._make_request(
        "DELETE", endpoint=f"{Endpoints.PINNED_V1}/posts/{post_id}"
    )


def get_bookmark(self, user_id: int, from_str: str = None) -> PostsResponse:
    # TODO: @Nullable @Query("from") String str なんのfromか不明
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USER_V1}/{user_id}/bookmarks",
        params={"from": from_str}, data_type=PostsResponse
    )


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
) -> PostsResponse:
    # TODO: not working {'next_page_value': None, 'result': 'success', 'posts': []}
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.POSTS_V2}/call_timeline",
        params={
            "group_id": group_id,
            "from_timestamp": from_timestamp,
            "number": number,
            "category_id": category_id,
            "call_type": call_type,
            "include_circle_call": include_circle_call,
            "cross_generation": cross_generation,
            "exclude_recent_gomimushi": exclude_recent_gomimushi,
            "shared_interest_categories": shared_interest_categories,
        }, data_type=PostsResponse
    )


def get_conversation(
        self,
        conversation_id: int,
        group_id: int = None,
        thread_id: int = None,
        from_post_id: int = None,
        number: int = 50,
        reverse: bool = True,
) -> PostsResponse:
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.CONVERSATIONS_V2}/{conversation_id}",
        params={
            "group_id": group_id,
            "thread_id": thread_id,
            "from_post_id": from_post_id,
            "number": number,
            "reverse": reverse,
        }, data_type=PostsResponse
    )


def get_conversation_root_posts(self, post_ids: List[int]) -> PostsResponse:
    pass


def get_following_call_timeline(
        self,
        from_timestamp: int = None,
        number: int = None,
        category_id: int = None,
        call_type: str = None,
        include_circle_call: bool = None,
        exclude_recent_gomimushi: bool = None
) -> PostsResponse:
    pass


def get_following_timeline(
        self,
        from_str: str = None,
        only_root: bool = None,
        order_by: str = None,
        number: int = None,
        mxn: int = None,
        reduce_selfie: bool = None,
        custom_generation_range: bool = None
) -> PostsResponse:
    pass


def get_group_highlight_posts(
        self,
        group_id: int,
        from_post: int = None,
        number: int = None
) -> PostsResponse:
    pass


def get_group_timeline_by_keyword(
        self,
        group_id: int,
        keyword: str,
        from_post_id: int = None,
        number: int = None,
        only_thread_posts: bool = False
) -> PostsResponse:
    pass


def get_group_timeline(
        self,
        group_id: int,
        from_post_id: int = None,
        reverse: bool = None,
        post_type: str = None,
        number: int = None,
        only_root: bool = None
) -> PostsResponse:
    pass


def get_timeline_by_hashtag(
        self,
        hashtag: str,
        from_post_id: int = None,
        number: int = None
) -> PostsResponse:
    pass


def get_my_posts(self, **kwargs: int | bool) -> PostsResponse:
    """

    Parameters:
    ---------------

        - from_post_id: int - (optional)
        - number: int - (optional)
        - include_group_post: bool - (optional)

    """
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.POSTS_V2}/mine",
        params=kwargs, data_type=PostsResponse
    )


def get_post(self, post_id: int) -> Post:
    # TODO: @Header("Cache-Control") @Nullable String str);
    response = self._make_request(
        "GET", endpoint=f"{Endpoints.POSTS_V2}/{post_id}",
        data_type=PostResponse
    )
    return response.post


def get_post_likers(
        self,
        post_id: int,
        from_id: int = None,
        number: int = None
) -> PostLikersResponse:
    params = {}
    if from_id is not None:
        params["from_id"] = from_id
    if number is not None:
        params["number"] = max(1, min(number, 100))

    return self._make_request(
        "GET", endpoint=f"{Endpoints.POSTS_V1}/{post_id}/likers",
        params=params, data_type=PostLikersResponse
    )


def get_post_reposts(self, post_id: int, **kwargs: int) -> PostsResponse:
    """

    Parameters:
    ---------------

        - post_id: int - (required)
        - from_post_id: int - (optional)
        - number: int - (optional)

    """
    return self._make_request(
        "GET", endpoint=f"{Endpoints.POSTS_V2}/{post_id}/reposts",
        params=kwargs, data_type=PostsResponse
    )


def get_posts(self, post_ids: List[int]) -> PostsResponse:
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.POSTS_V2}/multiple",
        params={"post_ids[]": post_ids}, data_type=PostsResponse
    )


def get_recommended_post_tags(
        self,
        tag: str = None,
        save_recent_search: bool = False
) -> PostTagsResponse:
    pass


def get_recommended_posts(
        self,
        experiment_num: int,
        variant_num: int,
        number: int
) -> PostsResponse:
    pass


def get_timeline_by_keyword(
        self,
        keyword: str = None,
        from_post_id: int = None,
        number: int = None
) -> PostsResponse:
    pass


def get_timeline(self, **kwargs: int | str | bool) -> PostsResponse:
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
    return self._make_request(
        # @GET("v2/posts/{noreply_mode}timeline")
        "GET", endpoint=f"{Endpoints.POSTS_V2}/timeline",
        params=kwargs, data_type=PostsResponse
    )


def get_url_metadata(self, url: str) -> SharedUrl:
    pass


def get_user_timeline(
        self,
        user_id: int,
        from_post_id: int = None,
        post_type: str = None,
        number: int = None
) -> PostsResponse:
    pass


def like_posts(self, post_ids: List[int]) -> LikePostsResponse:
    pass


def remove_bookmark(self, user_id: int, post_id: int):
    pass


def remove_group_highlight_post(self, group_id: int, post_id: int):
    pass


def remove_posts(self, post_ids: List[int]):
    pass


def report_post(
        self,
        post_id: int,
        opponent_id: int,
        category_id: int,
        reason: str = None,
        screenshot_filename: str = None,
        screenshot_2_filename: str = None,
        screenshot_3_filename: str = None,
        screenshot_4_filename: str = None
):
    pass


def unlike_post(self, post_id: int):
    pass


def update_post(
        self,
        post_id: int,
        text: str = None,
        font_size: int = None,
        color: int = None,
        message_tags: str = "[]",
) -> Post:
    # @NotNull @Part("uuid") String str2,
    # @NotNull @Part("api_key") String str3,
    # @Part("timestamp") long j3,
    # @NotNull @Part("signed_info") String str4);
    pass


def update_recommendation_feedback(
        self,
        post_id: int,
        experiment_num: int,
        variant_num: int,
        feedback_result: str
):
    pass


def validate_post(
        self,
        text: str,
        group_id: int = None,
        thread_id: int = None
) -> bool:
    # ValidationPostResponse
    pass


def view_video(self, video_id: int):
    pass


def vote_survey(self, survey_id: int, choice_id: int) -> Survey:
    # VoteSurveyResponse
    pass
