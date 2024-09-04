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
from ..responses import (
    ActiveFollowingsResponse,
    BlockedUserIdsResponse,
    BlockedUsersResponse,
    CreateUserResponse,
    FollowRecommendationsResponse,
    FollowRequestCountResponse,
    FollowUsersResponse,
    FootprintsResponse,
    HiddenResponse,
    HimaUsersResponse,
    RankingUsersResponse,
    RefreshCounterRequestsResponse,
    Response,
    SocialShareUsersResponse,
    UserResponse,
    UsersByTimestampResponse,
    UsersResponse,
    UserTimestampResponse,
)
from ..utils import md5, sha256


# pylint: disable=too-many-public-methods
class UserApi:
    """ユーザー API"""

    def __init__(self, client) -> None:
        # pylint: disable=import-outside-toplevel
        from ..client import Client

        self.__client: Client = client

    async def delete_footprint(self, user_id: int, footprint_id: int) -> Response:
        """足跡を削除する

        Args:
            user_id (int):
            footprint_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v2/users/{user_id}/footprints/{footprint_id}",
            return_type=Response,
        )

    async def follow_user(self, user_id: int) -> Response:
        """ユーザーをフォローする

        Args:
            user_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v2/users/{user_id}/follow",
            return_type=Response,
        )

    async def follow_users(self, user_ids: List[int]) -> Response:
        """複数のユーザーをフォローする

        Args:
            user_ids (List[int]):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v2/users/follow",
            params={"user_ids[]": user_ids},
            return_type=Response,
        )

    async def get_active_followings(self, **params) -> ActiveFollowingsResponse:
        """アクティブなフォロー中のユーザーを取得する

        Args:
            only_online (bool, optional):
            from_loggedin_at (int, optional):

        Returns:
            ActiveFollowingsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/users/active_followings",
            params=params,
            return_type=ActiveFollowingsResponse,
        )

    async def get_follow_recommendations(
        self, **params
    ) -> FollowRecommendationsResponse:
        """フォローするのにおすすめのユーザーを取得する

        Args:
            from_timestamp (int, optional):
            number (int, optional):
            sources (List[str], optional):

        Returns:
            FollowRecommendationsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/friends",
            params=params,
            return_type=FollowRecommendationsResponse,
        )

    async def get_follow_request(self, **params) -> UsersByTimestampResponse:
        """フォローリクエストを取得する

        Args:
            from_timestamp (int, optional):

        Returns:
            UsersByTimestampResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/users/follow_requests",
            params=params,
            return_type=UsersByTimestampResponse,
        )

    async def get_follow_request_count(self) -> FollowRequestCountResponse:
        """フォローリクエストの数を取得する

        Returns:
            FollowRequestCountResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/users/follow_requests_count",
            return_type=FollowRequestCountResponse,
        )

    async def get_following_users_born(self, **params) -> UsersResponse:
        """フォロー中のユーザーの誕生日を取得する

        Args:
            birthdate (int, optional):

        Returns:
            UsersResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/users/following_born_today",
            params=params,
            return_type=UsersResponse,
        )

    async def get_footprints(self, **params) -> FootprintsResponse:
        """足跡を取得する

        Args:
            from_id (int, optional):
            number (int, optional):
            mode (str, optional):

        Returns:
            FootprintsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/users/footprints",
            params=params,
            return_type=FootprintsResponse,
        )

    async def get_fresh_user(self, user_id: int) -> UserResponse:
        """認証情報などを含んだユーザー情報を取得する

        Args:
            user_id (int):

        Returns:
            UserResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v2/users/fresh/{user_id}",
            return_type=UserResponse,
        )

    async def get_hima_users(self, **params) -> HimaUsersResponse:
        """暇なユーザーを取得する

        Args:
            from_hima_id (int, optional):
            number (int, optional):

        Returns:
            HimaUsersResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/users/hima_users",
            params=params,
            return_type=HimaUsersResponse,
        )

    async def get_user_ranking(self, mode: str) -> RankingUsersResponse:
        """ユーザーのフォロワーランキングを取得する

        Examples:
            >>> # ルーキーを取得する:
            >>> client.get_user_ranking(mode="one_month")

            >>> # ミドルを取得する:
            >>> client.get_user_ranking(mode="six_months")

            >>> # 殿堂入りを取得する:
            >>> client.get_user_ranking(mode="all_time")

        Args:
            mode (str):

        Returns:
            RankingUsersResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/web/users/ranking",
            params={"mode": mode},
            return_type=RankingUsersResponse,
        )

    async def get_profile_refresh_counter_requests(
        self,
    ) -> RefreshCounterRequestsResponse:
        """投稿数やフォロワー数をリフレッシュするための残リクエスト数を取得する

        Returns:
            RefreshCounterRequestsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/users/reset_counters",
            return_type=RefreshCounterRequestsResponse,
        )

    async def get_social_shared_users(self, **params) -> SocialShareUsersResponse:
        """SNS共有をしたユーザーを取得する

        Args:
            sns_name (str):
            number (int, optional):
            from_id (int, optional):

        Returns:
            SocialShareUsersResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/users/social_shared_users",
            params=params,
            return_type=SocialShareUsersResponse,
        )

    async def get_timestamp(self) -> UserTimestampResponse:
        """タイムスタンプを取得する

        Returns:
            UserTimestampResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/users/timestamp",
            return_type=UserTimestampResponse,
        )

    async def get_user(self, user_id: int) -> UserResponse:
        """ユーザーの情報を取得する

        Args:
            user_id (int):

        Returns:
            UserResponse:
        """
        return await self.__client.request(
            "GET", config.API_HOST + f"/v2/users/{user_id}", return_type=UserResponse
        )

    async def get_user_followers(self, user_id: int, **params) -> FollowUsersResponse:
        """ユーザーのフォロワーを取得する

        Args:
            user_id (int):
            from_follow_id (int, optional):
            followed_by_me: (int, optional):
            number: (int, optional):

        Returns:
            FollowUsersResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v2/users/{user_id}/followers",
            params=params,
            return_type=FollowUsersResponse,
        )

    async def get_user_followings(self, user_id: int, **params) -> FollowUsersResponse:
        # @Body @Nullable SearchUsersRequest searchUsersRequest
        """フォロー中のユーザーを取得する

        Args:
            user_id (int):
            from_follow_id (int, optional):
            from_timestamp (int, optional):
            order_by (str, optional):
            number (int, optional):

        Returns:
            FollowUsersResponse:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v2/users/{user_id}/list_followings",
            params=params,
            return_type=FollowUsersResponse,
        )

    async def get_user_from_qr(self, qr: str) -> UserResponse:
        """QRコードからユーザーを取得する

        Args:
            qr (str):

        Returns:
            UserResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v1/users/qr_codes/{qr}",
            return_type=UserResponse,
        )

    async def get_user_without_leaving_footprint(self, user_id: int) -> UserResponse:
        """足跡をつけずにユーザーの情報を取得する

        Args:
            user_id (int):

        Returns:
            UserResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v2/users/info/{user_id}",
            return_type=UserResponse,
        )

    async def get_users(self, user_ids: List[int]) -> UsersResponse:
        """複数のユーザーの情報を取得する

        Args:
            user_ids (List[int]):

        Returns:
            UsersResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/users/list_id",
            params={"user_ids[]": user_ids},
            return_type=UsersResponse,
            jwt_required=True,
        )

    async def refresh_profile_counter(self, counter: str) -> Response:
        """プロフィールのカウンターを更新する

        Args:
            counter (str):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/users/reset_counters",
            json={"counter": counter},
            return_type=Response,
        )

    async def register(self, **params) -> CreateUserResponse:
        # @Nullable @Part("sns_info") SignUpSnsInfoRequest signUpSnsInfoRequest,
        """
        Args:
            email (str):
            email_grant_token (str):
            password (str):
            nickname (str):
            birth_date (str):
            gender (int, optional):
            country_code (str, optional):
            biography (str, optional):
            prefecture (str, optional):
            profile_icon_filename (str, optional):
            cover_image_filename (str, optional):
            en (int, optional):
            vn (int, optional):

        Returns:
            CreateUserResponse:
        """
        params.update(
            {
                "app_version": config.API_VERSION_NAME,
                "api_key": config.API_KEY,
                "signed_version": sha256(),
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__client.device_uuid, int(datetime.now().timestamp()), False
                ),
                "uuid": self.__client.device_uuid,
            }
        )
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v3/users/register",
            json=params,
            return_type=CreateUserResponse,
        )

    async def delete_user_avatar(self) -> Response:
        """ユーザーのアイコンを削除する

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v2/users/remove_profile_photo",
            return_type=Response,
        )

    async def delete_user_cover(self) -> Response:
        """ユーザーのカバー画像を削除する

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v2/users/remove_cover_image",
            return_type=Response,
        )

    async def reset_password(self, **params) -> Response:
        """パスワードをリセットする

        Args:
            email (str):
            email_grant_token (str):
            password (str):

        Returns:
            Response:
        """
        return await self.__client.request(
            "PUT",
            config.API_HOST + "/v1/users/reset_password",
            json=params,
            return_type=Response,
        )

    async def search_lobi_users(self, **params) -> UsersResponse:
        """Lobiのユーザーを検索する

        Args:
            nickname (str, optional):
            number (int, optional):
            from_str (str, optional):

        Returns:
            UsersResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/lobi_friends",
            params=params,
            return_type=UsersResponse,
        )

    async def search_users(self, **params) -> UsersResponse:
        """ユーザーを検索する

        Args:
            gender (int, optional):
            nickname (str, optional):
            title (str, optional):
            biography (str, optional):
            from_timestamp (int, optional):
            similar_age (bool, optional):
            not_recent_gomimushi (bool, optional):
            recently_created (bool, optional):
            same_prefecture (bool, optional):
            save_recent_search (bool, optional):

        Returns:
            UsersResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/web/users/search",
            params=params,
            return_type=UsersResponse,
        )

    async def set_follow_permission_enabled(self, **params) -> Response:
        """フォローを許可制にするかを設定する

        Args:
            nickname (str):
            is_private (bool, optional):

        Returns:
            Response:
        """
        params.update(
            {
                "uuid": self.__client.device_uuid,
                "api_key": config.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__client.device_uuid, int(datetime.now().timestamp()), True
                ),
                "signed_version": sha256(),
            }
        )
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v2/users/edit",
            json=params,
            return_type=Response,
        )

    async def take_action_follow_request(self, user_id: int, action: str) -> Response:
        """フォローリクエストを操作する

        Args:
            user_id (int):
            action (str):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v2/users/{user_id}/follow_request",
            json={"action": action},
            return_type=Response,
        )

    async def turn_on_hima(self) -> Response:
        """ひまなうを有効にする

        Returns:
            Response:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/users/hima",
            return_type=Response,
        )

    async def unfollow_user(self, user_id: int) -> Response:
        """ユーザーをアンフォローする

        Args:
            user_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v2/users/{user_id}/unfollow",
            return_type=Response,
        )

    async def update_user(self, nickname: str, **params) -> Response:
        """プロフィールを更新する

        Args:
            nickname (str):
            biography (str, optional):
            prefecture (str, optional):
            gender (int, optional):
            country_code (str, optional):
            profile_icon_filename (str, optional):
            cover_image_filename (str, optional):
            username (str, optional):

        Returns:
            Response:
        """
        params.update(
            {
                "nickname": nickname,
                "uuid": self.__client.device_uuid,
                "api_key": config.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__client.device_uuid, int(datetime.now().timestamp()), True
                ),
            }
        )
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v3/users/edit",
            json=params,
            return_type=Response,
        )

    # BlockApi

    async def block_user(
        self,
        user_id: int,
    ) -> Response:
        """ユーザーをブロックする

        Args:
            user_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/users/{user_id}/block",
            return_type=Response,
        )

    async def get_blocked_user_ids(self) -> BlockedUserIdsResponse:
        """あなたをブロックしたユーザーを取得する

        Returns:
            BlockedUserIdsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/users/block_ids",
            return_type=BlockedUserIdsResponse,
        )

    async def get_blocked_users(self, **params) -> BlockedUsersResponse:
        # @Body @NotNull SearchUsersRequest searchUsersRequest
        """ブロックしたユーザーを取得する

        Args:
            from_id (int, optional):

        Returns:
            BlockedUsersResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/users/blocked",
            params=params,
            return_type=BlockedUsersResponse,
        )

    async def unblock_user(self, user_id: int) -> Response:
        """ユーザーをアンブロックする

        Args:
            user_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v2/users/{user_id}/unblock",
            return_type=Response,
        )

    # HiddenApi

    async def get_hidden_users_list(self, **params) -> HiddenResponse:
        """非表示のユーザー一覧を取得する

        Args:
            from_str (str, optional):
            number (int, optional):

        Returns:
            HiddenResponse:
        """
        params["from"] = params.get("from_str")
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/hidden/users",
            params=params,
            return_type=HiddenResponse,
        )

    async def hide_user(self, user_id: int) -> Response:
        """ユーザーを非表示にする

        Args:
            user_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/hidden/users",
            json={"user_id": user_id},
            return_type=Response,
        )

    async def unhide_users(self, user_ids: List[int]) -> Response:
        """ユーザーの非表示を解除する

        Args:
            user_ids (List[int]):

        Returns:
            Response:
        """
        return await self.__client.request(
            "DELETE",
            config.API_HOST + "/v1/hidden/users",
            params={"user_ids[]": user_ids},
            return_type=Response,
        )
