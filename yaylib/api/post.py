"""
MIT License

Copyright (c) 2023 ekkx

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

from datetime import datetime

from .. import config
from ..errors import ClientError
from ..models import Post, SharedUrl
from ..responses import (
    BookmarkPostResponse,
    CreatePostResponse,
    LikePostsResponse,
    PostLikersResponse,
    PostResponse,
    PostsResponse,
    PostTagsResponse,
    Response,
    VoteSurveyResponse,
)
from ..utils import build_message_tags, get_post_type, md5


class PostApi:
    """投稿 API"""

    def __init__(self, client) -> None:
        from ..client import Client  # pylint: disable=import-outside-toplevel

        self.__client: Client = client

    async def add_bookmark(self, user_id: int, post_id: int) -> BookmarkPostResponse:
        return await self.__client.request(
            "PUT",
            config.API_HOST + f"/v1/users/{user_id}/bookmarks/{post_id}",
            return_type=BookmarkPostResponse,
        )

    async def add_group_highlight_post(self, group_id: int, post_id: int) -> Response:
        """

        投稿をグループのまとめに追加します

        """
        return await self.__client.request(
            "PUT",
            config.API_HOST + f"/v1/groups/{group_id}/highlights/{post_id}",
            return_type=Response,
        )

    async def create_call_post(
        self,
        text: str = None,
        font_size: int = None,
        color: int = None,
        group_id: int = None,
        call_type: str = None,
        category_id: int = None,
        game_title: str = None,
        joinable_by: str = None,
        message_tags: list = None,
        attachment_filename: str = None,
        attachment_2_filename: str = None,
        attachment_3_filename: str = None,
        attachment_4_filename: str = None,
        attachment_5_filename: str = None,
        attachment_6_filename: str = None,
        attachment_7_filename: str = None,
        attachment_8_filename: str = None,
        attachment_9_filename: str = None,
    ) -> CreatePostResponse:
        text, message_tags = build_message_tags(text)

        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/posts/new_conference_call",
            json={
                "text": text,
                "font_size": font_size,
                "color": color,
                "group_id": group_id,
                "call_type": call_type,
                "uuid": self.__client.device_uuid,
                "api_key": config.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__client.device_uuid, int(datetime.now().timestamp()), False
                ),
                "category_id": category_id,
                "game_title": game_title,
                "joinable_by": joinable_by,
                "message_tags": [] if message_tags is None else message_tags,
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
            return_type=CreatePostResponse,
        )

    async def create_group_pin_post(self, post_id: int, group_id: int) -> Response:
        return await self.__client.request(
            "PUT",
            config.API_HOST + "/v2/posts/group_pinned_post",
            json={"post_id": post_id, "group_id": group_id},
            return_type=Response,
        )

    async def create_pin_post(self, post_id: int) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/pinned/posts",
            json={"id": post_id},
            return_type=Response,
        )

    async def create_post(
        self,
        text: str = None,
        font_size: int = 0,
        color: int = 0,
        in_reply_to: int = None,
        group_id: int = None,
        mention_ids: list[int] = None,
        choices: list[str] = None,
        shared_url: str = None,
        message_tags: list = None,
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
        result = build_message_tags(text)
        if result is not None:
            text, message_tags = result

        post_type = get_post_type(
            choices=choices,
            shared_url=shared_url,
            video_file_name=video_file_name,
            attachment_filename=attachment_filename,
        )

        if shared_url is not None:
            try:
                shared_url = (await self.get_url_metadata(url=shared_url)).data
            except ClientError:
                self.__client.logger.error("Unable to get the URL metadata.")
                shared_url = None

        return await self.__client.request(
            "POST",
            config.API_HOST + "/v3/posts/new",
            json={
                "text": text,
                "font_size": font_size,
                "color": color,
                "in_reply_to": in_reply_to,
                "group_id": group_id,
                "post_type": post_type,
                "mention_ids": mention_ids,
                "choices": choices,
                "shared_url": shared_url,
                "message_tags": [] if message_tags is None else message_tags,
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
            return_type=Post,
            jwt_required=True,
        )

    async def create_repost(
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
        message_tags: list = None,
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
    ) -> CreatePostResponse:
        result = build_message_tags(text)
        if result is not None:
            text, message_tags = result

        post_type = get_post_type(
            choices=choices,
            shared_url=shared_url,
            video_file_name=video_file_name,
            attachment_filename=attachment_filename,
        )

        if shared_url is not None:
            try:
                shared_url = (await self.get_url_metadata(url=shared_url)).data
            except ClientError:
                self.__client.logger.error("Unable to get the URL metadata.")
                shared_url = None

        return await self.__client.request(
            "POST",
            config.API_HOST + "/v3/posts/repost",
            json={
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
                "message_tags": [] if message_tags is None else message_tags,
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
            return_type=CreatePostResponse,
            jwt_required=True,
        )

    async def create_share_post(
        self,
        shareable_type: str,
        shareable_id: int,
        text: str = None,
        font_size: int = None,
        color: int = None,
        group_id: int = None,
    ) -> Post:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v2/posts/new_share_post",
            json={
                "shareable_type": shareable_type,
                "shareable_id": shareable_id,
                "text": text,
                "font_size": font_size,
                "color": color,
                "group_id": group_id,
                "uuid": self.__client.device_uuid,
                "api_key": config.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__client.device_uuid, int(datetime.now().timestamp()), False
                ),
            },
            return_type=Post,
        )

    async def create_thread_post(
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
        message_tags: list = None,
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
        result = build_message_tags(text)
        if result is not None:
            text, message_tags = result

        post_type = get_post_type(
            choices=choices,
            shared_url=shared_url,
            video_file_name=video_file_name,
            attachment_filename=attachment_filename,
        )

        if shared_url is not None:
            try:
                shared_url = (await self.get_url_metadata(url=shared_url)).data
            except ClientError:
                self.__client.logger.error("Unable to get the URL metadata.")
                shared_url = None

        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/threads/{post_id}/posts",
            json={
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
                "message_tags": [] if message_tags is None else message_tags,
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
            return_type=Post,
            jwt_required=True,
        )

    async def delete_all_posts(self) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/posts/delete_all_post",
            return_type=Response,
        )

    async def unpin_group_post(self, group_id: int) -> Response:
        return await self.__client.request(
            "DELETE",
            config.API_HOST + "/v2/posts/group_pinned_post",
            json={"group_id": group_id},
            return_type=Response,
        )

    async def unpin_post(self, post_id: int) -> Response:
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v1/pinned/posts/{post_id}",
            return_type=Response,
        )

    async def get_bookmark(self, user_id: int, from_str: str = None) -> PostsResponse:
        params = {}
        if from_str:
            params = {"from": from_str}
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v1/users/{user_id}/bookmarks",
            params=params,
            return_type=PostsResponse,
        )

    async def get_timeline_calls(self, **params) -> PostsResponse:
        """

        Parameters:
        -----------

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
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/posts/call_timeline",
            params=params,
            return_type=PostsResponse,
        )

    async def get_conversation(self, conversation_id: int, **params) -> PostsResponse:
        """

        Parameters:
        -----------

            - conversation_id: int
            - group_id: int = None
            - thread_id: int = None
            - from_post_id: int = None
            - number: int = 50
            - reverse: bool = True

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v2/conversations/{conversation_id}",
            params=params,
            return_type=PostsResponse,
        )

    async def get_conversation_root_posts(self, post_ids: list[int]) -> PostsResponse:
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/conversations/root_posts",
            params={"ids[]": post_ids},
            return_type=PostsResponse,
        )

    async def get_following_call_timeline(self, **params) -> PostsResponse:
        """

        Parameters:
        -----------

            - from_timestamp: int = None
            - number: int = None
            - category_id: int = None
            - call_type: str = None
            - include_circle_call: bool = None
            - exclude_recent_gomimushi: bool = None

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/posts/call_followers_timeline",
            params=params,
            return_type=PostsResponse,
        )

    async def get_following_timeline(self, **params) -> PostsResponse:
        """

        Parameters:
        -----------

            - from_str: str = None
            - only_root: bool = None
            - order_by: str = None
            - number: int = None
            - mxn: int = None
            - reduce_selfie: bool = None
            - custom_generation_range: bool = None

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/posts/following_timeline",
            params=params,
            return_type=PostsResponse,
        )

    async def get_group_highlight_posts(self, group_id: int, **params) -> PostsResponse:
        """

        グループのまとめ投稿を取得します

        Parameters:
        -----------

            - group_id: int
            - from_post: int = None
            - number: int = None

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v1/groups/{group_id}/highlights",
            params=params,
            return_type=PostsResponse,
        )

    async def get_group_timeline_by_keyword(
        self, group_id: int, keyword: str, **params
    ) -> PostsResponse:
        """

        Parameters:
        -----------

            - group_id: int
            - keyword: str
            - from_post_id: int = None
            - number: int = None
            - only_thread_posts: bool = False

        """
        params["keyword"] = keyword
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v2/groups/{group_id}/posts/search",
            params=params,
            return_type=PostsResponse,
        )

    async def get_group_timeline(self, group_id: int, **params) -> PostsResponse:
        """

        Parameters:
        -----------

            - group_id: int
            - from_post_id: int
            - reverse: bool
            - post_type: str
            - number: int
            - only_root: bool

        """
        params["group_id"] = group_id
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/posts/group_timeline",
            params=params,
            return_type=PostsResponse,
        )

    async def get_timeline_by_hashtag(self, hashtag: str, **params) -> PostsResponse:
        """

        Parameters:
        -----------

            - hashtag: str - (required)
            - from_post_id: int - (optional)
            - number: int - (optional)

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v2/posts/tags/{hashtag}",
            params=params,
            return_type=PostsResponse,
        )

    async def get_my_posts(self, **params) -> PostsResponse:
        """

        Parameters:
        ---------------

            - from_post_id: int - (optional)
            - number: int - (optional)
            - include_group_post: bool - (optional)

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/posts/mine",
            params=params,
            return_type=PostsResponse,
        )

    async def get_post(self, post_id: int) -> PostResponse:
        return await self.__client.request(
            "GET", config.API_HOST + f"/v2/posts/{post_id}", return_type=PostResponse
        )

    async def get_post_likers(self, post_id: int, **params) -> PostLikersResponse:
        """

        Parameters:
        ---------------

            - from_id: int - (optional)
            - number: int - (optional)

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v1/posts/{post_id}/likers",
            params=params,
            return_type=PostLikersResponse,
        )

    async def get_post_reposts(self, post_id: int, **params: int) -> PostsResponse:
        """

        Parameters:
        ---------------

            - post_id: int - (required)
            - from_post_id: int - (optional)
            - number: int - (optional)

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v2/posts/{post_id}/reposts",
            params=params,
            return_type=PostsResponse,
        )

    async def get_posts(self, post_ids: list[int]) -> PostsResponse:
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/posts/multiple",
            params={"post_ids[]": post_ids},
            return_type=PostsResponse,
        )

    async def get_recommended_post_tags(
        self, tag: str = None, save_recent_search: bool = False
    ) -> PostTagsResponse:
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/posts/recommended_tag",
            json={"tag": tag, "save_recent_search": save_recent_search},
            return_type=PostTagsResponse,
        )

    async def get_recommended_posts(self, **params) -> PostsResponse:
        """

        Parameters:
        ---------------

            - experiment_num: int - (Required)
            - variant_num: int - (Required)
            - number: int - (Optional)

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/posts/recommended_timeline",
            params=params,
            return_type=PostsResponse,
        )

    async def get_timeline_by_keyword(
        self, keyword: str = None, **params
    ) -> PostsResponse:
        """

        Parameters:
        ---------------

            - keyword: str
            - from_post_id: int
            - number: int

        """
        params["keyword"] = keyword
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/posts/search",
            params=params,
            return_type=PostsResponse,
        )

    async def get_timeline(self, **params: int | str | bool) -> PostsResponse:
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
        endpoint = "/v2/posts/timeline"
        if "noreply_mode" in params and params["noreply_mode"] is True:
            endpoint = "/v2/posts/noreply_timeline"
        return await self.__client.request(
            "GET", config.API_HOST + endpoint, params=params, return_type=PostsResponse
        )

    async def get_url_metadata(self, url: str) -> SharedUrl:
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/posts/url_metadata",
            params={"url": url},
            return_type=SharedUrl,
        )

    async def get_user_timeline(self, user_id: int, **params) -> PostsResponse:
        """

        Parameters:
        ---------------

            - from_post_id: int - (optional)
            - number: int - (optional)
            - post_type: str - (optional)

        """
        params["user_id"] = user_id
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/posts/user_timeline",
            params=params,
            return_type=PostsResponse,
        )

    async def like(self, post_ids: list[int]) -> LikePostsResponse:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v2/posts/like",
            json={"post_ids": post_ids},
            return_type=LikePostsResponse,
        )

    async def remove_bookmark(self, user_id: int, post_id: int) -> Response:
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v1/users/{user_id}/bookmarks/{post_id}",
            return_type=Response,
        )

    async def remove_group_highlight_post(
        self, group_id: int, post_id: int
    ) -> Response:
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v1/groups/{group_id}/highlights/{post_id}",
            return_type=Response,
        )

    async def remove_posts(self, post_ids: list[int]) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v2/posts/mass_destroy",
            json={"posts_ids": post_ids},
            return_type=Response,
        )

    async def unlike(self, post_id: int) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/posts/{post_id}/unlike",
            return_type=Response,
        )

    async def update_post(
        self,
        post_id: int,
        text: str = None,
        font_size: int = None,
        color: int = None,
        message_tags: list = None,
    ) -> Post:
        result = build_message_tags(text)
        if result is not None:
            text, message_tags = result
        return await self.__client.request(
            "PUT",
            config.API_HOST + f"/v3/posts/{post_id}",
            json={
                "text": text,
                "font_size": font_size,
                "color": color,
                "message_tags": [] if message_tags is None else message_tags,
                "api_key": config.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__client.device_uuid, int(datetime.now().timestamp()), False
                ),
            },
        )

    async def view_video(self, video_id: int) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/posts/videos/{video_id}/view",
            return_type=Response,
        )

    async def vote_survey(self, survey_id: int, choice_id: int) -> VoteSurveyResponse:
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v2/surveys/{survey_id}/vote",
            json={"choice_id": choice_id},
            return_type=VoteSurveyResponse,
        )
