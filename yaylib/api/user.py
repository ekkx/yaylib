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
from typing import List, Optional

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


class UserApi:
    """ユーザー API"""

    def __init__(self, client) -> None:
        from ..client import Client  # pylint: disable=import-outside-toplevel

        self.__client: Client = client

    async def delete_footprint(self, user_id: int, footprint_id: int) -> Response:
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v2/users/{user_id}/footprints/{footprint_id}",
            return_type=Response,
        )

    async def destroy_user(self) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v2/users/destroy",
            json={
                "uuid": self.__client.device_uuid,
                "api_key": config.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__client.device_uuid, int(datetime.now().timestamp()), True
                ),
            },
            return_type=Response,
        )

    async def follow_user(self, user_id: int) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v2/users/{user_id}/follow",
            return_type=Response,
        )

    async def follow_users(self, user_ids: List[int]) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v2/users/follow",
            params={"user_ids[]": user_ids},
            return_type=Response,
        )

    async def get_active_followings(self, **params) -> ActiveFollowingsResponse:
        """

        Parameters:
        -----------

            - only_online: bool
            - from_loggedin_at: int = None

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
        """

        Parameters:
        -----------

            - from_timestamp: int = None,
            - number: int = None,
            - sources: list[str] = None

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/friends",
            params=params,
            return_type=FollowRecommendationsResponse,
        )

    async def get_follow_request(
        self, from_timestamp: Optional[int] = None
    ) -> UsersByTimestampResponse:
        params = {}
        if from_timestamp:
            params["from_timestamp"] = from_timestamp
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/users/follow_requests",
            params=params,
            return_type=UsersByTimestampResponse,
        )

    async def get_follow_request_count(self) -> FollowRequestCountResponse:
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/users/follow_requests_count",
            return_type=FollowRequestCountResponse,
        )

    async def get_following_users_born(
        self, birthdate: Optional[int] = None
    ) -> UsersResponse:
        params = {}
        if birthdate:
            params["birthdate"] = birthdate
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/users/following_born_today",
            params=params,
            return_type=UsersResponse,
        )

    async def get_footprints(self, **params) -> FootprintsResponse:
        """

        Parameters:
        -----------

            - from_id: int = None
            - number: int = None
            - mode: str = None

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/users/footprints",
            params=params,
            return_type=FootprintsResponse,
        )

    async def get_fresh_user(self, user_id: int) -> UserResponse:
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v2/users/fresh/{user_id}",
            return_type=UserResponse,
        )

    async def get_hima_users(self, **params) -> HimaUsersResponse:
        """

        Parameters:
        -----------

            - from_hima_id: int = None
            - number: int = None

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/users/hima_users",
            params=params,
            return_type=HimaUsersResponse,
        )

    async def get_user_ranking(self, mode: str) -> RankingUsersResponse:
        """

        ユーザーのランキングを取得します

        Examples:
        ---------

        >>> ルーキーを取得する場合:

        >>> api.get_user_ranking(mode="one_month")

        ---

        >>> ミドルを取得する場合:

        >>> api.get_user_ranking(mode="six_months")

        ---

        >>> 殿堂入りを取得する場合:

        >>> api.get_user_ranking(mode="all_time")

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
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/users/reset_counters",
            return_type=RefreshCounterRequestsResponse,
        )

    async def get_social_shared_users(self, **params) -> SocialShareUsersResponse:
        """

        Parameters:
        -----------

            - sns_name: str - (Required)
            - number: int - (Optional)
            - from_id: int - (Optional)

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/users/social_shared_users",
            params=params,
            return_type=SocialShareUsersResponse,
        )

    async def get_timestamp(
        self,
    ) -> UserTimestampResponse:
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/users/timestamp",
            return_type=UserTimestampResponse,
        )

    async def get_user(self, user_id: int) -> UserResponse:
        return await self.__client.request(
            "GET", config.API_HOST + f"/v2/users/{user_id}", return_type=UserResponse
        )

    async def get_user_followers(self, user_id: int, **params) -> FollowUsersResponse:
        """

        Parameters:
        -----------

            - user_id: int - (required)
            - from_follow_id: int - (optional)
            - followed_by_me: int - (optional)
            - number: int - (optional)

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v2/users/{user_id}/followers",
            params=params,
            return_type=FollowUsersResponse,
        )

    async def get_user_followings(self, user_id: int, **params) -> FollowUsersResponse:
        # @Body @Nullable SearchUsersRequest searchUsersRequest
        """

        Parameters:
        -----------

            - user_id: int
            - from_follow_id: int = None
            - from_timestamp: int = None
            - order_by: str = None
            - number: int - (optional)

        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v2/users/{user_id}/list_followings",
            params=params,
            return_type=FollowUsersResponse,
        )

    async def get_user_from_qr(self, qr: str) -> UserResponse:
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v1/users/qr_codes/{qr}",
            return_type=UserResponse,
        )

    async def get_user_without_leaving_footprint(self, user_id: int) -> UserResponse:
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v2/users/info/{user_id}",
            return_type=UserResponse,
        )

    async def get_users(self, user_ids: List[int]) -> UsersResponse:
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/users/list_id",
            params={"user_ids[]": user_ids},
            return_type=UsersResponse,
            jwt_required=True,
        )

    async def refresh_profile_counter(self, counter: str) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/users/reset_counters",
            json={"counter": counter},
            return_type=Response,
        )

    async def register(
        self,
        email: str,
        email_grant_token: str,
        password: str,
        nickname: str,
        birth_date: str,
        gender: int = -1,
        country_code: str = "JP",
        biography: Optional[str] = None,
        prefecture: Optional[str] = None,
        profile_icon_filename: Optional[str] = None,
        cover_image_filename: Optional[str] = None,
        # @Nullable @Part("sns_info") SignUpSnsInfoRequest signUpSnsInfoRequest,
        en: Optional[int] = None,
        vn: Optional[int] = None,
    ) -> CreateUserResponse:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v3/users/register",
            json={
                "app_version": config.API_VERSION_NAME,
                "api_key": config.API_KEY,
                "signed_version": sha256(),
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__client.device_uuid, int(datetime.now().timestamp()), False
                ),
                "uuid": self.__client.device_uuid,
                "nickname": nickname,
                "birth_date": birth_date,
                "gender": gender,
                "country_code": country_code,
                "biography": biography,
                "prefecture": prefecture,
                "profile_icon_filename": profile_icon_filename,
                "cover_image_filename": cover_image_filename,
                "email": email,
                "password": password,
                "email_grant_token": email_grant_token,
                "en": en,
                "vn": vn,
            },
            return_type=CreateUserResponse,
        )

    async def delete_user_avatar(self) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v2/users/remove_profile_photo",
            return_type=Response,
        )

    async def delete_user_cover(self) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v2/users/remove_cover_image",
            return_type=Response,
        )

    async def reset_password(
        self, email: str, email_grant_token: str, password: str
    ) -> Response:
        return await self.__client.request(
            "PUT",
            config.API_HOST + "/v1/users/reset_password",
            json={
                "email": email,
                "email_grant_token": email_grant_token,
                "password": password,
            },
            return_type=Response,
        )

    async def search_lobi_users(self, **params) -> UsersResponse:
        """

        Parameters:
        -----------

            - nickname: str = None
            - number: int = None
            - from_str: str = None

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/lobi_friends",
            params=params,
            return_type=UsersResponse,
        )

    async def search_users(self, **params) -> UsersResponse:
        """

        ユーザーを検索します

        Parameters
        ----------

            - gender: int = None
            - nickname: str = None
            - title: str = None
            - biography: str = None
            - from_timestamp: int = None
            - similar_age: bool = None
            - not_recent_gomimushi: bool = None
            - recently_created: bool = None
            - same_prefecture: bool = None
            - save_recent_search: bool = None

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/web/users/search",
            params=params,
            return_type=UsersResponse,
        )

    async def set_follow_permission_enabled(
        self, nickname: str, is_private: Optional[bool] = None
    ) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v2/users/edit",
            json={
                "nickname": nickname,
                "is_private": is_private,
                "uuid": self.__client.device_uuid,
                "api_key": config.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__client.device_uuid, int(datetime.now().timestamp()), True
                ),
                "signed_version": sha256(),
            },
            return_type=Response,
        )

    async def set_setting_follow_recommendation_enabled(self, on: bool) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/users/visible_on_sns_friend_recommendation_setting",
            params={"on": on},
            return_type=Response,
        )

    async def take_action_follow_request(self, user_id: int, action: str) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v2/users/{user_id}/follow_request",
            json={"action": action},
            return_type=Response,
        )

    async def turn_on_hima(self) -> Response:
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/users/hima",
            return_type=Response,
        )

    async def unfollow_user(self, user_id: int) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v2/users/{user_id}/unfollow",
            return_type=Response,
        )

    async def update_language(self, language: str) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/users/language",
            json={
                "uuid": self.__client.device_uuid,
                "api_key": config.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__client.device_uuid, int(datetime.now().timestamp()), True
                ),
                "language": language,
            },
            return_type=Response,
        )

    async def update_user(self, nickname: str, **params) -> Response:
        """

        プロフィールを更新します

        Parameters
        ----------

            - nickname: str = (required)
            - biography: str = (optional)
            - prefecture: str = (optional)
            - gender: int = (optional)
            - country_code: str = (optional)
            - profile_icon_filename: str = (optional)
            - cover_image_filename: str = (optional)
            - username: str = (optional)

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
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/users/{user_id}/block",
            return_type=Response,
        )

    async def get_blocked_user_ids(self) -> BlockedUserIdsResponse:
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/users/block_ids",
            return_type=BlockedUserIdsResponse,
        )

    async def get_blocked_users(
        self, from_id: Optional[int] = None
    ) -> BlockedUsersResponse:
        # @Body @NotNull SearchUsersRequest searchUsersRequest
        params = {}
        if from_id:
            params["from_id"] = from_id
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/users/blocked",
            params=params,
            return_type=BlockedUsersResponse,
        )

    async def unblock_user(self, user_id: int) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v2/users/{user_id}/unblock",
            return_type=Response,
        )

    # HiddenApi

    async def get_hidden_users_list(self, **params) -> HiddenResponse:
        """

        Parameters:
        -----------

            - from: str = None
            - number: int = None

        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/hidden/users",
            params=params,
            return_type=HiddenResponse,
        )

    async def hide_user(self, user_id: int) -> Response:
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/hidden/users",
            json={"user_id": user_id},
            return_type=Response,
        )

    async def unhide_users(self, user_ids: List[int]) -> Response:
        return await self.__client.request(
            "DELETE",
            config.API_HOST + "/v1/hidden/users",
            params={"user_ids[]": user_ids},
            return_type=Response,
        )
