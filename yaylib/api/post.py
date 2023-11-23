"""
MIT License

Copyright (c) 2023-present qvco

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
"""

from __future__ import annotations

import json
from datetime import datetime

from ..config import Endpoints
from ..errors import ForbiddenError, NotFoundError
from ..models import ConferenceCall, Post, SharedUrl, Survey
from ..responses import (
    BookmarkPostResponse,
    CreatePostResponse,
    PostResponse,
    PostsResponse,
    PostLikersResponse,
    PostTagsResponse,
    LikePostsResponse,
    ValidationPostResponse,
)


def add_bookmark(
    self, user_id: int, post_id: int, access_token: str = None
) -> BookmarkPostResponse:
    response = self.request(
        "PUT",
        endpoint=f"{Endpoints.USERS_V1}/{user_id}/bookmarks/{post_id}",
        data_type=BookmarkPostResponse,
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Post has been added to the bookmarks.")
    return response


def add_group_highlight_post(
    self, group_id: int, post_id: int, access_token: str = None
):
    """

    投稿をグループのまとめに追加します

    """
    response = self.request(
        "PUT",
        endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/highlights/{post_id}",
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Post has been added to the group highlight.")
    return response


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
    message_tags: list = [],
    attachment_filename: str = None,
    attachment_2_filename: str = None,
    attachment_3_filename: str = None,
    attachment_4_filename: str = None,
    attachment_5_filename: str = None,
    attachment_6_filename: str = None,
    attachment_7_filename: str = None,
    attachment_8_filename: str = None,
    attachment_9_filename: str = None,
    access_token: str = None,
) -> ConferenceCall:
    if text is not None:
        if "@:start:" in text and ":end:" in text:
            text, message_tags = parse_mention_format(self, text)

    timestamp = int(datetime.now().timestamp())

    response = self.request(
        "POST",
        endpoint=f"{Endpoints.POSTS_V2}/new_conference_call",
        payload={
            "text": text,
            "font_size": font_size,
            "color": color,
            "group_id": group_id,
            "call_type": call_type,
            "uuid": self.uuid,
            "api_key": self.api_key,
            "timestamp": timestamp,
            "signed_info": self.generate_signed_info(self.device_uuid, timestamp),
            "category_id": category_id,
            "game_title": game_title,
            "joinable_by": joinable_by,
            "message_tags": str(message_tags),
            "attachment_filename": attachment_filename,
            "attachment_2_filename": attachment_2_filename,
            "attachment_3_filename": attachment_3_filename,
            "attachment_4_filename": attachment_4_filename,
            "attachment_5_filename": attachment_5_filename,
            "attachment_6_filename": attachment_6_filename,
            "attachment_7_filename": attachment_7_filename,
            "attachment_8_filename": attachment_8_filename,
            "attachment_9_filename": attachment_9_filename,
        },
        data_type=CreatePostResponse,
        auth_required=True,
        access_token=access_token,
    ).conference_call
    self.logger.info("Call post has been created.")
    return response


def create_group_pin_post(self, post_id: int, group_id: int, access_token: str = None):
    response = self.request(
        "PUT",
        endpoint=f"{Endpoints.POSTS_V2}/group_pinned_post",
        payload={"post_id": post_id, "group_id": group_id},
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Pinned the post in the group.")
    return response


def create_pin_post(self, post_id: int, access_token: str = None):
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.PINNED_V1}/posts",
        payload={"id": post_id},
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Pinned post.")
    return response


def mention(self, user_id: int) -> str:
    if not (isinstance(user_id, int) or str(user_id).isdigit()):
        raise ValueError(
            "The value of 'user_id' must be an integer or a string containing only digits."
        )
    return "@:start:" + str(user_id) + ":end:"


def convert_mention_format(self, text) -> tuple:
    # Do NOT write this function in client.py
    formatted_text = ""
    user_ids = []
    segments = text.split("@:start:")

    for i, segment in enumerate(segments):
        if i == 0:
            formatted_text += segment
            continue
        user_id, text = segment.split(":end:")
        username = self.get_user_without_leaving_footprint(user_id).user.nickname
        formatted_text += "@" + username + " " + text
        user_ids.append(user_id)

    return formatted_text, user_ids


def parse_mention_format(self, text) -> tuple:
    # Do NOT write this function in client.py
    user_ids = {}
    message_tags = []
    offset = 0

    text, user_ids = convert_mention_format(self, text)

    for user_id in user_ids:
        start = text.find("@", offset)
        if start == -1:
            break
        end = text.find(" ", start)
        if end == -1:
            end = len(text)
        username = text[start + 1 : end]
        message_tags.append(
            {
                "type": "user",
                "user_id": int(user_id),
                "offset": start,
                "length": len(username) + 1,
            }
        )
        offset = end

    return text, json.dumps(message_tags)


def create_post(
    self,
    text: str = None,
    font_size: int = 0,
    color: int = 0,
    in_reply_to: int = None,
    group_id: int = None,
    mention_ids: list[int] = None,
    choices: list[str] = None,
    shared_url: str = None,
    message_tags: list = [],
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
    access_token: str = None,
) -> Post:
    timestamp = int(datetime.now().timestamp())
    headers = self.session.headers.copy()
    headers["X-Jwt"] = self.generete_jwt(timestamp)
    
    if text is not None:
        if "@:start:" in text and ":end:" in text:
            text, message_tags = parse_mention_format(self, text)

    post_type = (
        "survey"
        if choices
        else "shareable_url"
        if shared_url
        else "video"
        if video_file_name
        else "image"
        if attachment_filename
        else "text"
    )

    if shared_url is not None:
        try:
            shared_url = self.get_url_metadata(
                url=shared_url, access_token=access_token
            ).data
        except ForbiddenError:
            self.logger.error("Unable to get the URL metadata.")
            shared_url = None

    response = self.request(
        "POST",
        endpoint=f"{Endpoints.POSTS_V3}/new",
        payload={
            "text": text,
            "font_size": font_size,
            "color": color,
            "in_reply_to": in_reply_to,
            "group_id": group_id,
            "post_type": post_type,
            "mention_ids": mention_ids,
            "choices": choices,
            "shared_url": shared_url,
            "message_tags": str(message_tags),
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
        },
        data_type=Post,
        headers=headers,
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Post has been created.")
    return response


def create_repost(
    self,
    post_id: int,
    text: str = None,
    font_size: int = None,
    color: int = None,
    in_reply_to: int = None,
    group_id: int = None,
    mention_ids: list[int] = None,
    choices: list[str] = None,
    shared_url: dict[str, str | int] = None,
    message_tags: list = [],
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
    access_token: str = None,
) -> Post:
    timestamp = int(datetime.now().timestamp())
    headers = self.session.headers.copy()
    headers["X-Jwt"] = self.generete_jwt(timestamp)

    if text is not None:
        if "@:start:" in text and ":end:" in text:
            text, message_tags = parse_mention_format(self, text)

    post_type = (
        "survey"
        if choices
        else "shareable_url"
        if shared_url
        else "video"
        if video_file_name
        else "image"
        if attachment_filename
        else "text"
    )

    if shared_url is not None:
        try:
            shared_url = self.get_url_metadata(
                url=shared_url, access_token=access_token
            ).data
        except ForbiddenError:
            self.logger.error("Unable to get the URL metadata.")
            shared_url = None

    response = self.request(
        "POST",
        endpoint=f"{Endpoints.POSTS_V3}/repost",
        payload={
            "post_id": post_id,
            "text": text,
            "font_size": font_size,
            "color": color,
            "in_reply_to": in_reply_to,
            "group_id": group_id,
            "post_type": post_type,
            "mention_ids": mention_ids,
            "choices": choices,
            "shared_url": shared_url,
            "message_tags": str(message_tags),
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
        },
        data_type=CreatePostResponse,
        headers=headers,
        auth_required=True,
        access_token=access_token,
    ).post
    self.logger.info("Repost has been created.")
    return response


def create_share_post(
    self,
    shareable_type: str,
    shareable_id: int,
    text: str = None,
    font_size: int = None,
    color: int = None,
    group_id: int = None,
    access_token: str = None,
) -> Post:
    timestamp = int(datetime.now().timestamp())

    response = self.request(
        "POST",
        endpoint=f"{Endpoints.POSTS_V2}/new_share_post",
        payload={
            "shareable_type": shareable_type,
            "shareable_id": shareable_id,
            "text": text,
            "font_size": font_size,
            "color": color,
            "group_id": group_id,
            "uuid": self.uuid,
            "api_key": self.api_key,
            "timestamp": timestamp,
            "signed_info": self.generate_signed_info(self.device_uuid, timestamp),
        },
        data_type=Post,
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Share post has been created.")
    return response


def create_thread_post(
    self,
    post_id: int,
    text: str = None,
    font_size: int = None,
    color: int = None,
    in_reply_to: int = None,
    group_id: int = None,
    mention_ids: list[int] = None,
    choices: list[str] = None,
    shared_url: dict[str, str | int] = None,
    message_tags: list = [],
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
    access_token: str = None,
) -> Post:
    timestamp = int(datetime.now().timestamp())
    headers = self.session.headers.copy()
    headers["X-Jwt"] = self.generete_jwt(timestamp)

    if text is not None:
        if "@:start:" in text and ":end:" in text:
            text, message_tags = parse_mention_format(self, text)

    post_type = (
        "survey"
        if choices
        else "shareable_url"
        if shared_url
        else "video"
        if video_file_name
        else "image"
        if attachment_filename
        else "text"
    )

    if shared_url is not None:
        try:
            shared_url = self.get_url_metadata(
                url=shared_url, access_token=access_token
            ).data
        except ForbiddenError:
            self.logger.error("Unable to get the URL metadata.")
            shared_url = None

    response = self.request(
        "POST",
        endpoint=f"{Endpoints.THREADS_V1}/{post_id}/posts",
        payload={
            "id": post_id,
            "text": text,
            "font_size": font_size,
            "color": color,
            "in_reply_to": in_reply_to,
            "group_id": group_id,
            "post_type": post_type,
            "mention_ids": mention_ids,
            "choices": choices,
            "shared_url": shared_url,
            "message_tags": str(message_tags),
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
        },
        data_type=Post,
        headers=headers,
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Thread post has been created.")
    return response


def delete_all_post(self, access_token: str = None):
    try:
        response = self.request(
            "POST",
            endpoint=f"{Endpoints.POSTS_V1}/delete_all_post",
            auth_required=True,
            access_token=access_token,
        )
        self.logger.info("Post deletion request has been sent.")
        return response
    except NotFoundError:
        self.logger.info("Post not found. Skipping...")


def delete_group_pin_post(self, group_id: int, access_token: str = None):
    response = self.request(
        "DELETE",
        endpoint=f"{Endpoints.POSTS_V2}/group_pinned_post",
        payload={"group_id": group_id},
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Unpinned post in the group.")
    return response


def delete_pin_post(self, post_id: int, access_token: str = None):
    response = self.request(
        "DELETE",
        endpoint=f"{Endpoints.PINNED_V1}/posts/{post_id}",
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Unpinned post.")
    return response


def get_bookmark(
    self, user_id: int, from_str: str = None, access_token: str = None
) -> PostsResponse:
    params = {}
    if from_str:
        params = {"from": from_str}
    return self.request(
        "GET",
        endpoint=f"{Endpoints.USERS_V1}/{user_id}/bookmarks",
        params=params,
        data_type=PostsResponse,
        auth_required=True,
        access_token=access_token,
    )


def get_timeline_calls(self, access_token: str = None, **params) -> PostsResponse:
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
    return self.request(
        "GET",
        endpoint=f"{Endpoints.POSTS_V2}/call_timeline",
        params=params,
        data_type=PostsResponse,
        access_token=access_token,
    )


def get_conversation(
    self, conversation_id: int, access_token: str = None, **params
) -> PostsResponse:
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
    return self.request(
        "GET",
        endpoint=f"{Endpoints.CONVERSATIONS_V2}/{conversation_id}",
        params=params,
        data_type=PostsResponse,
        access_token=access_token,
    )


def get_conversation_root_posts(
    self, post_ids: list[int], access_token: str = None
) -> PostsResponse:
    return self.request(
        "GET",
        endpoint=f"{Endpoints.CONVERSATIONS_V2}/root_posts",
        params={"ids[]": post_ids},
        data_type=PostsResponse,
        access_token=access_token,
    )


def get_following_call_timeline(
    self, access_token: str = None, **params
) -> PostsResponse:
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
    return self.request(
        "GET",
        endpoint=f"{Endpoints.POSTS_V2}/call_followers_timeline",
        params=params,
        data_type=PostsResponse,
        auth_required=True,
        access_token=access_token,
    )


def get_following_timeline(self, access_token: str = None, **params) -> PostsResponse:
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
    return self.request(
        "GET",
        endpoint=f"{Endpoints.POSTS_V2}/following_timeline",
        params=params,
        data_type=PostsResponse,
        auth_required=True,
        access_token=access_token,
    )


def get_group_highlight_posts(
    self, group_id: int, access_token: str = None, **params
) -> PostsResponse:
    """

    グループのまとめ投稿を取得します

    Parameters:
    ----------

        - group_id: int
        - from_post: int = None
        - number: int = None

    """
    return self.request(
        "GET",
        endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/highlights",
        params=params,
        data_type=PostsResponse,
        access_token=access_token,
    )


def get_group_timeline_by_keyword(
    self, group_id: int, keyword: str, access_token: str = None, **params
) -> PostsResponse:
    """

    Parameters:
    ----------

        - group_id: int
        - keyword: str
        - from_post_id: int = None
        - number: int = None
        - only_thread_posts: bool = False

    """
    params["keyword"] = keyword
    return self.request(
        "GET",
        endpoint=f"{Endpoints.GROUPS_V2}/{group_id}/posts/search",
        params=params,
        data_type=PostsResponse,
        access_token=access_token,
    )


def get_group_timeline(
    self, group_id: int, access_token: str = None, **params
) -> PostsResponse:
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
    params["group_id"] = group_id
    return self.request(
        "GET",
        endpoint=f"{Endpoints.POSTS_V2}/group_timeline",
        params=params,
        data_type=PostsResponse,
        access_token=access_token,
    )


def get_timeline_by_hashtag(
    self, hashtag: str, access_token: str = None, **params
) -> PostsResponse:
    """

    Parameters:
    ----------

        - hashtag: str - (required)
        - from_post_id: int - (optional)
        - number: int - (optional)

    """
    return self.request(
        "GET",
        endpoint=f"{Endpoints.POSTS_V2}/tags/{hashtag}",
        params=params,
        data_type=PostsResponse,
        access_token=access_token,
    )


def get_my_posts(self, access_token: str = None, **params) -> PostsResponse:
    """

    Parameters:
    ---------------

        - from_post_id: int - (optional)
        - number: int - (optional)
        - include_group_post: bool - (optional)

    """
    return self.request(
        "GET",
        endpoint=f"{Endpoints.POSTS_V2}/mine",
        params=params,
        data_type=PostsResponse,
        auth_required=True,
        access_token=access_token,
    )


def get_post(self, post_id: int, access_token: str = None) -> Post:
    return self.request(
        "GET",
        endpoint=f"{Endpoints.POSTS_V2}/{post_id}",
        data_type=PostResponse,
        access_token=access_token,
    ).post


def get_post_likers(
    self, post_id: int, access_token: str = None, **params
) -> PostLikersResponse:
    """

    Parameters:
    ---------------

        - from_id: int - (optional)
        - number: int - (optional)

    """
    return self.request(
        "GET",
        endpoint=f"{Endpoints.POSTS_V1}/{post_id}/likers",
        params=params,
        data_type=PostLikersResponse,
        access_token=access_token,
    )


def get_post_reposts(
    self, post_id: int, access_token: str = None, **params: int
) -> PostsResponse:
    """

    Parameters:
    ---------------

        - post_id: int - (required)
        - from_post_id: int - (optional)
        - number: int - (optional)

    """
    return self.request(
        "GET",
        endpoint=f"{Endpoints.POSTS_V2}/{post_id}/reposts",
        params=params,
        data_type=PostsResponse,
        access_token=access_token,
    )


def get_posts(self, post_ids: list[int], access_token: str = None) -> PostsResponse:
    return self.request(
        "GET",
        endpoint=f"{Endpoints.POSTS_V2}/multiple",
        params={"post_ids[]": post_ids},
        data_type=PostsResponse,
        access_token=access_token,
    )


def get_recommended_post_tags(
    self, tag: str = None, save_recent_search: bool = False, access_token: str = None
) -> PostTagsResponse:
    return self.request(
        "GET",
        endpoint=f"{Endpoints.POSTS_V1}/recommended_tag",
        payload={"tag": tag, "save_recent_search": save_recent_search},
        data_type=PostTagsResponse,
        access_token=access_token,
    )


def get_recommended_posts(self, access_token: str = None, **params) -> PostsResponse:
    """

    Parameters:
    ---------------

        - experiment_num: int - (Required)
        - variant_num: int - (Required)
        - number: int - (Optional)

    """
    return self.request(
        "GET",
        endpoint=f"{Endpoints.POSTS_V2}/recommended_timeline",
        params=params,
        data_type=PostsResponse,
        access_token=access_token,
    )


def get_timeline_by_keyword(
    self, keyword: str = None, access_token: str = None, **params
) -> PostsResponse:
    """

    Parameters:
    ---------------

        - keyword: str
        - from_post_id: int
        - number: int

    """
    params["keyword"] = keyword
    return self.request(
        "GET",
        endpoint=f"{Endpoints.POSTS_V2}/search",
        params=params,
        data_type=PostsResponse,
        access_token=access_token,
    )


def get_timeline(
    self, access_token: str = None, **params: int | str | bool
) -> PostsResponse:
    # - from: str - (optional)
    """

    Parameters:
    ---------------

        - noreply_mode: bool - (optional)
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
    endpoint = f"{Endpoints.POSTS_V2}/timeline"
    if "noreply_mode" in params and params["noreply_mode"] is True:
        endpoint = f"{Endpoints.POSTS_V2}/noreply_timeline"
    return self.request(
        "GET",
        endpoint=endpoint,
        params=params,
        data_type=PostsResponse,
        auth_required=True,
        access_token=access_token,
    )


def get_url_metadata(self, url: str, access_token: str = None) -> SharedUrl:
    return self.request(
        "GET",
        endpoint=f"{Endpoints.POSTS_V2}/url_metadata",
        params={"url": url},
        data_type=SharedUrl,
        access_token=access_token,
    )


def get_user_timeline(
    self, user_id: int, access_token: str = None, **params
) -> PostsResponse:
    """

    Parameters:
    ---------------

        - from_post_id: int - (optional)
        - number: int - (optional)
        - post_type: str - (optional)

    """
    params["user_id"] = user_id
    return self.request(
        "GET",
        endpoint=f"{Endpoints.POSTS_V2}/user_timeline",
        params=params,
        data_type=PostsResponse,
        access_token=access_token,
    )


def like_posts(
    self, post_ids: list[int], access_token: str = None
) -> LikePostsResponse:
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.POSTS_V2}/like",
        payload={"post_ids": post_ids},
        data_type=LikePostsResponse,
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Posts have been liked.")
    return response


def remove_bookmark(self, user_id: int, post_id: int, access_token: str = None):
    response = self.request(
        "DELETE",
        endpoint=f"{Endpoints.USERS_V1}/{user_id}/bookmarks/{post_id}",
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Bookmark has been removed.")
    return response


def remove_group_highlight_post(
    self, group_id: int, post_id: int, access_token: str = None
):
    response = self.request(
        "DELETE",
        endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/highlights/{post_id}",
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Group hightlight post removed.")
    return response


def remove_posts(self, post_ids: list[int], access_token: str = None):
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.POSTS_V2}/mass_destroy",
        payload={"posts_ids": post_ids},
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Posts have been removed.")
    return response


def report_post(
    self,
    post_id: int,
    opponent_id: int,
    category_id: int,
    reason: str = None,
    screenshot_filename: str = None,
    screenshot_2_filename: str = None,
    screenshot_3_filename: str = None,
    screenshot_4_filename: str = None,
    access_token: str = None,
):
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.POSTS_V3}/{post_id}/report",
        payload={
            "opponent_id": opponent_id,
            "category_id": category_id,
            "reason": reason,
            "screenshot_filename": screenshot_filename,
            "screenshot_2_filename": screenshot_2_filename,
            "screenshot_3_filename": screenshot_3_filename,
            "screenshot_4_filename": screenshot_4_filename,
        },
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Post has been reported.")
    return response


def unlike_post(self, post_id: int, access_token: str = None):
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.POSTS_V1}/{post_id}/unlike",
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Post has been unliked.")
    return response


def update_post(
    self,
    post_id: int,
    text: str = None,
    font_size: int = None,
    color: int = None,
    message_tags: list = [],
    access_token: str = None,
) -> Post:
    if "@:start:" in text and ":end:" in text:
        text, message_tags = parse_mention_format(self, text)

    timestamp = int(datetime.now().timestamp())

    response = self.request(
        "PUT",
        endpoint=f"{Endpoints.POSTS_V3}/{post_id}",
        payload={
            "text": text,
            "font_size": font_size,
            "color": color,
            "message_tags": str(message_tags),
            "api_key": self.api_key,
            "timestamp": timestamp,
            "signed_info": self.generate_signed_info(self.device_uuid, timestamp),
        },
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Post has been updated.")
    return response


def view_video(self, video_id: int, access_token: str = None):
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.POSTS_V1}/videos/{video_id}/view",
        access_token=access_token,
    )
    self.logger.info("Viewed the video.")
    return response


def vote_survey(
    self, survey_id: int, choice_id: int, access_token: str = None
) -> Survey:
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.SURVEYS_V2}/{survey_id}/vote",
        payload={"choice_id": choice_id},
        data_type=ValidationPostResponse,
        auth_required=True,
        access_token=access_token,
    ).survey
    self.logger.info("Voted.")
    return response
