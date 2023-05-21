import time
from typing import Dict, List

from ..config import *
from ..errors import *
from .api import (
    _check_authorization,
    _get,
    _post,
    _put,
    _delete,
    _handle_response,
)


def add_bookmark(self, user_id: int, post_id: int, headers: Dict[str, str | int] = None):
    headers = _check_authorization(self, headers)
    return _put(
        self=self,
        endpoint=f"https://{Endpoints.USER_V1}/{user_id}/bookmarks/{post_id}",
        headers=headers
    )


def add_group_highlight_post(self, group_id: int, post_id: int, headers: Dict[str, str | int] = None):
    headers = _check_authorization(self, headers)
    return _put(
        self=self,
        endpoint=f"https://{Endpoints.GROUPS_V1}/{group_id}/highlights/{post_id}",
        headers=headers
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
        attachment_9_filename: str = None,
        headers: Dict[str, str | int] = None):
    # TODO: @NotNull "uuid", "api_key", "timestamp", "signed_info"
    headers = _check_authorization(self, headers)
    params = {
        "text": text,
        "font_size": font_size,
        "color": color,
        "group_id": group_id,
        "call_type": call_type,
        "uuid": self.uuid,
        "api_key": self.yay_api_key,
        "timestamp": time.time(),
        "signed_info": self.signed_info,
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
    }
    return _post(
        self=self,
        endpoint=f"https://{Endpoints.POSTS_V2}/new_conference_call",
        params=params,
        headers=headers
    )


def create_group_pin_post(self, post_id: int, group_id: int, headers: Dict[str, str | int] = None):
    headers = _check_authorization(self, headers)
    params = {"post_id": post_id, "group_id": group_id}
    return _put(
        self=self,
        endpoint=f"https://{Endpoints.POSTS_V2}/group_pinned_post",
        params=params,
        headers=headers
    )


def create_pin_post(self, post_id: int, headers: Dict[str, str | int] = None):
    headers = _check_authorization(self, headers)
    params = {"id": post_id}
    return _post(
        self=self,
        endpoint=f"https://{Endpoints.PINNED_V1}/posts",
        params=params,
        headers=headers
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
        headers: Dict[str, str | int] = None):
    # TODO: @Header("X-Jwt") @NotNull String str,
    headers = _check_authorization(self, headers)
    params = {
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
    }
    return _post(
        self=self,
        endpoint=f"https://{Endpoints.POSTS_V3}/new",
        params=params,
        headers=headers
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
        headers: Dict[str, str | int] = None):
    # TODO: @Header("X-Jwt") @NotNull String str,
    headers = _check_authorization(self, headers)
    params = {
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
    }
    return _post(
        self=self,
        endpoint=f"https://{Endpoints.POSTS_V3}/repost",
        params=params,
        headers=headers
    )


def create_share_post(
        self,
        shareable_type: str,
        shareable_id: int,
        text: str = None,
        font_size: int = None,
        color: int = None,
        group_id: int = None,
        headers: Dict[str, str | int] = None):
    # TODO: @NotNull "uuid", "api_key", "timestamp", "signed_info"
    headers = _check_authorization(self, headers)
    params = {
        "shareable_type": shareable_type,
        "shareable_id": shareable_id,
        "text": text,
        "font_size": font_size,
        "color": color,
        "group_id": group_id,
        "uuid": self.uuid,
        "api_key": self.api_key,
        "timestamp": time.time(),
        "signed_info": self.signed_info,
    }
    return _post(
        self=self,
        endpoint=f"https://{Endpoints.POSTS_V2}/new_share_post",
        params=params,
        headers=headers
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
        headers: Dict[str, str | int] = None):
    # TODO: @Header("X-Jwt") @NotNull String str,
    headers = _check_authorization(self, headers)
    params = {
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
    }
    return _post(
        self=self,
        endpoint=f"https://{Endpoints.THREADS_V1}/posts",
        params=params,
        headers=headers
    )


def delete_all_post(self, headers: Dict[str, str | int] = None):
    headers = _check_authorization(self, headers)
    return _post(
        self=self,
        endpoint=f"https://{Endpoints.POSTS_V1}/delete_all_post",
        headers=headers
    )


def delete_group_pin_post(self, group_id: int, headers: Dict[str, str | int] = None):
    headers = _check_authorization(self, headers)
    params = {"group_id": group_id}
    return _delete(
        self=self,
        endpoint=f"https://{Endpoints.POSTS_V2}/group_pinned_post",
        params=params,
        headers=headers
    )


def delete_pin_post(self, post_id: int, headers: Dict[str, str | int] = None):
    headers = _check_authorization(self, headers)
    return _delete(
        self=self,
        endpoint=f"https://{Endpoints.PINNED_V1}/posts/{post_id}",
        headers=headers
    )


def get_bookmark(self, user_id: int, from_str: str = None, headers: Dict[str, str | int] = None):
    # TODO: @Nullable @Query("from") String str なんのfromか不明
    headers = _check_authorization(self, headers)
    params = {"from": from_str}
    return _get(
        self=self,
        endpoint=f"https://{Endpoints.USER_V1}/{user_id}/bookmarks",
        params=params,
        headers=headers
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
    headers: Dict[str, str | int] = None
):
    # TODO: not working {'next_page_value': None, 'result': 'success', 'posts': []}
    params = {
        "group_id": group_id,
        "from_timestamp": from_timestamp,
        "number": number,
        "category_id": category_id,
        "call_type": call_type,
        "include_circle_call": include_circle_call,
        "cross_generation": cross_generation,
        "exclude_recent_gomimushi": exclude_recent_gomimushi,
        "shared_interest_categories": shared_interest_categories,
    }
    return _get(
        self=self,
        endpoint=f"https://{Endpoints.POSTS_V2}/call_timeline",
        params=params,
        headers=headers
    )


def get_conversation(
    self,
    conversation_id: int,
    group_id: int = None,
    thread_id: int = None,
    from_post_id: int = None,
    number: int = 50,
    reverse: bool = True,
    headers: Dict[str, str | int] = None
):
    params = {
        "group_id": group_id,
        "thread_id": thread_id,
        "from_post_id": from_post_id,
        "number": number,
        "reverse": reverse,
    }
    return _get(
        self=self,
        endpoint=f"https://{Endpoints.CONVERSATIONS_V2}/{conversation_id}",
        params=params,
        headers=headers
    )


def get_conversation_root_posts(self):
    pass


def get_following_call_timeline(self):
    pass


def get_following_timeline(self):
    pass


def get_group_highlight_posts(self):
    pass


def get_group_timeline_by_keyword(self):
    pass


def get_group_timeline(self):
    pass


def get_timeline_by_hashtag(self):
    pass


def get_my_posts(self, from_post_id: int = None, number: int = 100, include_group_post: bool = False, headers: Dict[str, str | int] = None):
    # TODO: include_group_postはfalseだったらサークルの投稿は含まないはずなのにサークルの投稿しか出てこないしなんかおかしい
    headers = _check_authorization(self, headers)
    params = {
        "number": number,
        "include_group_post": include_group_post
    }
    if from_post_id:
        params["from_post_id"] = from_post_id
    return _get(
        self=self,
        endpoint=f"https://{Endpoints.POSTS_V2}/mine",
        params=params,
        headers=headers
    )


def get_post(self):
    pass


def get_oost_likers(self):
    pass


def get_reposts(self):
    pass


def get_posts(self):
    pass


def get_recommended_post_tags(self):
    pass


def get_recommended_posts(self):
    pass


def get_timeline_by_keyword(self):
    pass


def get_timeline(self):
    pass


def get_url_metadata(self):
    pass


def get_user_timeline(self):
    pass


def like_posts(self):
    pass


def remove_bookmark(self):
    pass


def remove_group_highlight_post(self):
    pass


def remove_posts(self):
    pass


def report_post(self):
    pass


def unlike_post(self):
    pass


def update_post(self):
    pass


def update_recommendation_feedback(self):
    pass


def validate_post(self):
    pass


def view_video(self):
    pass


def vote_survey(self):
    pass
