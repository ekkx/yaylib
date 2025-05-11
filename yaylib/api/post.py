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

from datetime import datetime
from typing import List

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


# pylint: disable=too-many-public-methods
class PostApi:
    """投稿 API"""

    def __init__(self, client) -> None:
        # pylint: disable=import-outside-toplevel
        from ..client import Client

        self.__client: Client = client

    async def add_bookmark(self, user_id: int, post_id: int) -> BookmarkPostResponse:
        """ブックマークに追加する

        Args:
            user_id (int):
            post_id (int):

        Returns:
            BookmarkPostResponse:
        """
        return await self.__client.request(
            "PUT",
            config.API_HOST + f"/v1/users/{user_id}/bookmarks/{post_id}",
            return_type=BookmarkPostResponse,
        )

    async def add_group_highlight_post(self, group_id: int, post_id: int) -> Response:
        """投稿をグループのまとめに追加する

        Args:
            group_id (int):
            post_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "PUT",
            config.API_HOST + f"/v1/groups/{group_id}/highlights/{post_id}",
            return_type=Response,
        )

    async def create_call_post(self, text: str = None, **params) -> CreatePostResponse:
        """通話の投稿を作成する

        Args:
            text (str, optional):
            font_size (int, optional):
            color (int, optional):
            group_id (int, optional):
            call_type (str, optional):
            category_id (int, optional):
            game_title (str, optional):
            joinable_by (str, optional):
            message_tags (List, optional):
            attachment_filename (str, optional):
            attachment_2_filename (str, optional):
            attachment_3_filename (str, optional):
            attachment_4_filename (str, optional):
            attachment_5_filename (str, optional):
            attachment_6_filename (str, optional):
            attachment_7_filename (str, optional):
            attachment_8_filename (str, optional):
            attachment_9_filename (str, optional):

        Returns:
            CreatePostResponse:
        """
        result = build_message_tags(text)
        if result is not None:
            text, message_tags = result

        params.update(
            {
                "text": text,
                "uuid": self.__client.device_uuid,
                "api_key": config.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__client.device_uuid, int(datetime.now().timestamp()), False
                ),
                "message_tags": [] if result is None else message_tags,
            }
        )

        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/posts/new_conference_call",
            json=params,
            return_type=CreatePostResponse,
        )

    async def pin_group_post(self, post_id: int, group_id: int) -> Response:
        """サークルの投稿をピン留めする

        Args:
            post_id (int):
            group_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "PUT",
            config.API_HOST + "/v2/posts/group_pinned_post",
            json={"post_id": post_id, "group_id": group_id},
            return_type=Response,
        )

    async def pin_post(self, post_id: int) -> Response:
        """プロフィールに投稿をピン留めする

        Args:
            post_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/pinned/posts",
            json={"id": post_id},
            return_type=Response,
        )

    async def create_post(self, text: str = None, **params) -> Post:
        """投稿を作成する

        Args:
            text (str, optional):
            font_size (int, optional): . Defaults to 0.
            color (int, optional): . Defaults to 0.
            in_reply_to (int, optional):
            group_id (int, optional):
            mention_ids (List[int], optional):
            choices (List[str], optional):
            shared_url (str, optional):
            message_tags (List, optional):
            attachment_filename (str, optional):
            attachment_2_filename (str, optional):
            attachment_3_filename (str, optional):
            attachment_4_filename (str, optional):
            attachment_5_filename (str, optional):
            attachment_6_filename (str, optional):
            attachment_7_filename (str, optional):
            attachment_8_filename (str, optional):
            attachment_9_filename (str, optional):
            video_file_name (str, optional):

        Returns:
            Post:
        """
        result = build_message_tags(text)
        if result is not None:
            text, message_tags = result

        post_type = get_post_type(
            choices=params.get("choices"),
            shared_url=params.get("shared_url"),
            video_file_name=params.get("video_file_name"),
            attachment_filename=params.get("attachment_filename"),
        )

        shared_url = params.get("shared_url")
        if shared_url is not None:
            try:
                shared_url = (
                    await self.get_url_metadata(url=params.get("shared_url"))
                ).data
            except ClientError:
                self.__client.logger.error("Unable to get the URL metadata.")
                shared_url = None

        params.update(
            {
                "text": text,
                "post_type": post_type,
                "shared_url": shared_url,
                "message_tags": [] if result is None else message_tags,
            }
        )

        return await self.__client.request(
            "POST",
            config.API_HOST + "/v3/posts/new",
            json=params,
            return_type=Post,
            jwt_required=True,
        )

    async def create_repost(
        self, post_id: int, text: str = None, **params
    ) -> CreatePostResponse:
        """投稿を(´∀｀∩)↑age↑する

        Args:
            post_id (int):
            text (str, optional):
            font_size (int, optional):
            color (int, optional):
            in_reply_to (int, optional):
            group_id (int, optional):
            mention_ids (List[int], optional):
            choices (List[str], optional):
            shared_url (Dict[str, str  |  int], optional):
            message_tags (List, optional):
            attachment_filename (str, optional):
            attachment_2_filename (str, optional):
            attachment_3_filename (str, optional):
            attachment_4_filename (str, optional):
            attachment_5_filename (str, optional):
            attachment_6_filename (str, optional):
            attachment_7_filename (str, optional):
            attachment_8_filename (str, optional):
            attachment_9_filename (str, optional):
            video_file_name (str, optional):

        Returns:
            CreatePostResponse:
        """
        result = build_message_tags(text)
        if result is not None:
            text, message_tags = result

        post_type = get_post_type(
            choices=params.get("choices"),
            shared_url=params.get("shared_url"),
            video_file_name=params.get("video_file_name"),
            attachment_filename=params.get("attachment_filename"),
        )

        if params.get("shared_url") is not None:
            try:
                shared_url = (
                    await self.get_url_metadata(url=params.get("shared_url"))
                ).data
            except ClientError:
                self.__client.logger.error("Unable to get the URL metadata.")
                shared_url = None

        params.update(
            {
                "post_id": post_id,
                "text": text,
                "post_type": post_type,
                "shared_url": shared_url,
                "message_tags": [] if result is None else message_tags,
            }
        )

        return await self.__client.request(
            "POST",
            config.API_HOST + "/v3/posts/repost",
            json=params,
            return_type=CreatePostResponse,
            jwt_required=True,
        )

    async def create_share_post(
        self, shareable_type: str, shareable_id: int, text: str = None, **params
    ) -> Post:
        """シェア投稿を作成する

        Args:
            shareable_type (str):
            shareable_id (int):
            text (str, optional):
            font_size (int, optional):
            color (int, optional):
            group_id (int, optional):

        Returns:
            Post:
        """
        params.update(
            {
                "shareable_type": shareable_type,
                "shareable_id": shareable_id,
                "text": text,
                "uuid": self.__client.device_uuid,
                "api_key": config.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__client.device_uuid, int(datetime.now().timestamp()), False
                ),
            }
        )
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v2/posts/new_share_post",
            json=params,
            return_type=Post,
        )

    async def create_thread_post(
        self, thread_id: int, text: str = None, **params
    ) -> Post:
        """スレッドの投稿を作成する

        Args:
            thread_id (int):
            text (str, optional):
            font_size (int, optional):
            color (int, optional):
            in_reply_to (int, optional):
            group_id (int, optional):
            mention_ids (List[int], optional):
            choices (list[str], optional):
            shared_url (Dict[str, str  |  int], optional):
            message_tags (List, optional):
            attachment_filename (str, optional):
            attachment_2_filename (str, optional):
            attachment_3_filename (str, optional):
            attachment_4_filename (str, optional):
            attachment_5_filename (str, optional):
            attachment_6_filename (str, optional):
            attachment_7_filename (str, optional):
            attachment_8_filename (str, optional):
            attachment_9_filename (str, optional):
            video_file_name (str, optional):

        Returns:
            Post:
        """
        result = build_message_tags(text)
        if result is not None:
            text, message_tags = result

        post_type = get_post_type(
            choices=params.get("choices"),
            shared_url=params.get("shared_url"),
            video_file_name=params.get("video_file_name"),
            attachment_filename=params.get("attachment_filename"),
        )

        shared_url = params.get("shared_url")
        if shared_url is not None:
            try:
                shared_url = (
                    await self.get_url_metadata(url=params.get("shared_url"))
                ).data
            except ClientError:
                self.__client.logger.error("Unable to get the URL metadata.")
                shared_url = None

        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/threads/{thread_id}/posts",
            json={
                "text": text,
                "post_type": post_type,
                "shared_url": shared_url,
                "message_tags": [] if result is None else message_tags,
            },
            return_type=Post,
            jwt_required=True,
        )

    async def delete_all_posts(self) -> Response:
        """すべての自分の投稿を削除する

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/posts/delete_all_post",
            return_type=Response,
        )

    async def unpin_group_post(self, group_id: int) -> Response:
        """グループのピン投稿を解除する

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "DELETE",
            config.API_HOST + "/v2/posts/group_pinned_post",
            json={"group_id": group_id},
            return_type=Response,
        )

    async def unpin_post(self, post_id: int) -> Response:
        """プロフィール投稿のピンを解除する

        Args:
            post_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v1/pinned/posts/{post_id}",
            return_type=Response,
        )

    async def get_bookmark(self, user_id: int, **params) -> PostsResponse:
        """ブックマークを取得する

        Args:
            user_id (int):
            from_str (str, optional):

        Returns:
            PostsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v1/users/{user_id}/bookmarks",
            params=params,
            return_type=PostsResponse,
        )

    async def get_timeline_calls(self, **params) -> PostsResponse:
        """誰でも通話を取得する

        Args:
            group_id (int, optional):
            from_timestamp (int, optional):
            number (int, optional):
            category_id (int, optional):
            call_type (str, optional): Defaults to "voice".
            include_circle_call (bool, optional):
            cross_generation (bool, optional):
            exclude_recent_gomimushi (bool, optional):
            shared_interest_categories (bool, optional):

        Returns:
            PostsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/posts/call_timeline",
            params=params,
            return_type=PostsResponse,
        )

    async def get_conversation(self, conversation_id: int, **params) -> PostsResponse:
        """リプライを含める投稿の会話を取得する

        Args:
            conversation_id (int):
            group_id (int, optional):
            thread_id (int, optional):
            from_post_id (int, optional):
            number (int, optional):
            reverse (bool, optional):

        Returns:
            PostsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v2/conversations/{conversation_id}",
            params=params,
            return_type=PostsResponse,
        )

    async def get_conversation_root_posts(self, post_ids: List[int]) -> PostsResponse:
        """会話の原点の投稿を取得する

        Args:
            post_ids (List[int]):

        Returns:
            PostsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/conversations/root_posts",
            params={"ids[]": post_ids},
            return_type=PostsResponse,
        )

    async def get_following_call_timeline(self, **params) -> PostsResponse:
        """フォロー中の通話を取得する

        Args:
            from_timestamp (int, optional):
            number (int, optional):
            category_id (int, optional):
            include_circle_call (bool, optional):
            exclude_recent_gomimushi (bool, optional):

        Returns:
            PostsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/posts/call_followers_timeline",
            params=params,
            return_type=PostsResponse,
        )

    async def get_following_timeline(self, **params) -> PostsResponse:
        """フォロー中のタイムラインを取得する

        Args:
            from_str (str, optional):
            only_root (bool, optional):
            order_by (str, optional):
            number (int, optional):
            mxn (int, optional):
            reduce_selfie (bool, optional):
            custom_generation_range (bool, optional):

        Returns:
            PostsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/posts/following_timeline",
            params=params,
            return_type=PostsResponse,
        )

    async def get_group_highlight_posts(self, group_id: int, **params) -> PostsResponse:
        """グループのまとめ投稿を取得する

        Args:
            group_id (int):
            from_post (int, optional):
            number (int, optional):

        Returns:
            PostsResponse:
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
        """グループの投稿をキーワードで検索する

        Args:
            group_id (int):
            keyword (str):
            from_post_id (int, optional):
            number (int, optional):
            only_thread_posts (bool, optional):

        Returns:
            PostsResponse:
        """
        params["keyword"] = keyword
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v2/groups/{group_id}/posts/search",
            params=params,
            return_type=PostsResponse,
        )

    async def get_group_timeline(self, group_id: int, **params) -> PostsResponse:
        """グループのタイムラインを取得する

        Args:
            group_id (int):
            from_post_id (int, optional):
            reverse (bool, optional):
            post_type (str, optional):
            number (int, optional):
            only_root (bool, optional):

        Returns:
            PostsResponse:
        """
        params["group_id"] = group_id
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/posts/group_timeline",
            params=params,
            return_type=PostsResponse,
        )

    async def get_timeline_by_hashtag(self, hashtag: str, **params) -> PostsResponse:
        """ハッシュタグでタイムラインを検索する

        Args:
            hashtag (str):
            from_post_id (int, optional):
            number (int, optional):

        Returns:
            PostsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v2/posts/tags/{hashtag}",
            params=params,
            return_type=PostsResponse,
        )

    async def get_my_posts(self, **params) -> PostsResponse:
        """自分の投稿を取得する

        Args:
            from_post_id (int, optional):
            number (int, optional):
            include_group_post (bool, optional):

        Returns:
            PostsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/posts/mine",
            params=params,
            return_type=PostsResponse,
        )

    async def get_post(self, post_id: int) -> PostResponse:
        """投稿の詳細を取得する

        Args:
            post_id (int):

        Returns:
            PostResponse:
        """
        return await self.__client.request(
            "GET", config.API_HOST + f"/v2/posts/{post_id}", return_type=PostResponse
        )

    async def get_post_likers(self, post_id: int, **params) -> PostLikersResponse:
        """投稿にいいねしたユーザーを取得する

        Args:
            post_id (int):
            from_id (int, optional):
            number (int, optional):

        Returns:
            PostLikersResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v1/posts/{post_id}/likers",
            params=params,
            return_type=PostLikersResponse,
        )

    async def get_reposts(self, post_id: int, **params: int) -> PostsResponse:
        """投稿の(´∀｀∩)↑age↑を取得する

        Args:
            post_id (int):
            from_post_id (int, optional):
            number (int, optional):

        Returns:
            PostsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v2/posts/{post_id}/reposts",
            params=params,
            return_type=PostsResponse,
        )

    async def get_posts(self, post_ids: List[int]) -> PostsResponse:
        """複数の投稿を取得する

        Args:
            post_ids (List[int]):

        Returns:
            PostsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/posts/multiple",
            params={"post_ids[]": post_ids},
            return_type=PostsResponse,
        )

    async def get_recommended_post_tags(self, **params) -> PostTagsResponse:
        """おすすめのタグ候補を取得する

        Args:
            tag (str, optional):
            save_recent_search (bool, optional):

        Returns:
            PostTagsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/posts/recommended_tag",
            json=params,
            return_type=PostTagsResponse,
        )

    async def get_recommended_posts(self, **params) -> PostsResponse:
        """おすすめの投稿を取得する

        Args:
            experiment_num (int):
            variant_num (int, optional):
            number (int, optional):


        Returns:
            PostsResponse:
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
        """キーワードでタイムラインを検索する

        Args:
            keyword (str, optional):
            from_post_id (int, optional):
            number (int, optional):

        Returns:
            PostsResponse:
        """
        params["keyword"] = keyword
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/posts/search",
            params=params,
            return_type=PostsResponse,
        )

    async def get_timeline(self, **params) -> PostsResponse:
        # - from: str - (optional)
        """タイムラインを取得する

        Args:
            noreply_mode (bool, optional):
            from_post_id (int, optional):
            number (int, optional):
            order_by (str, optional):
            experiment_older_age_rules (bool, optional):
            shared_interest_categories (bool, optional):
            mxn (int, optional):
            en (int, optional):
            vn (int, optional):
            reduce_selfie (bool, optional):
            custom_generation_range (bool, optional):

        Returns:
            PostsResponse:
        """
        endpoint = "/v2/posts/timeline"
        if "noreply_mode" in params and params["noreply_mode"] is True:
            endpoint = "/v2/posts/noreply_timeline"
        return await self.__client.request(
            "GET", config.API_HOST + endpoint, params=params, return_type=PostsResponse
        )

    async def get_url_metadata(self, url: str) -> SharedUrl:
        """URLのメタデータを取得する

        Args:
            url (str):

        Returns:
            SharedUrl:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/posts/url_metadata",
            params={"url": url},
            return_type=SharedUrl,
        )

    async def get_user_timeline(self, user_id: int, **params) -> PostsResponse:
        """ユーザーのタイムラインを取得する

        Args:
            user_id (int):
            from_post_id (int, optional):
            number (int, optional):
            post_type (str, optional):

        Returns:
            PostsResponse:
        """
        params["user_id"] = user_id
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/posts/user_timeline",
            params=params,
            return_type=PostsResponse,
        )

    async def like(self, post_ids: List[int]) -> LikePostsResponse:
        """投稿にいいねする

        Args:
            post_ids (List[int]):

        Returns:
            LikePostsResponse:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v2/posts/like",
            json={"post_ids": post_ids},
            return_type=LikePostsResponse,
        )

    async def delete_bookmark(self, user_id: int, post_id: int) -> Response:
        """ブックマークを削除する

        Args:
            user_id (int):
            post_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v1/users/{user_id}/bookmarks/{post_id}",
            return_type=Response,
        )

    async def delete_group_highlight_post(
        self, group_id: int, post_id: int
    ) -> Response:
        """サークルのまとめから投稿を解除する

        Args:
            group_id (int):
            post_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v1/groups/{group_id}/highlights/{post_id}",
            return_type=Response,
        )

    async def delete_posts(self, post_ids: List[int]) -> Response:
        """投稿を削除する

        Args:
            post_ids (List[int]):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v2/posts/mass_destroy",
            json={"posts_ids": post_ids},
            return_type=Response,
        )

    async def unlike(self, post_id: int) -> Response:
        """いいねを解除する

        Args:
            post_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/posts/{post_id}/unlike",
            return_type=Response,
        )

    async def update_post(
        self,
        post_id: int,
        text: str = None,
        **params,
    ) -> Post:
        """投稿を編集する

        Args:
            post_id (int):
            text (str, optional):
            font_size (int, optional):
            color (int, optional):
            message_tags (List, optional):

        Returns:
            Post:
        """
        result = build_message_tags(text)
        if result is not None:
            text, message_tags = result

        params.update(
            {
                "text": text,
                "message_tags": [] if result is None else message_tags,
                "api_key": config.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__client.device_uuid, int(datetime.now().timestamp()), False
                ),
            }
        )

        return await self.__client.request(
            "PUT",
            config.API_HOST + f"/v3/posts/{post_id}",
            json=params,
        )

    async def view_video(self, video_id: int) -> Response:
        """動画を視聴する

        Args:
            video_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/posts/videos/{video_id}/view",
            return_type=Response,
        )

    async def vote_survey(self, survey_id: int, choice_id: int) -> VoteSurveyResponse:
        """アンケートに投票する

        Args:
            survey_id (int):
            choice_id (int):

        Returns:
            VoteSurveyResponse:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v2/surveys/{survey_id}/vote",
            json={"choice_id": choice_id},
            return_type=VoteSurveyResponse,
        )
