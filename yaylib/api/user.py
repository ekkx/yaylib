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

from .. import client
from ..config import Configs
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
    RefreshCounterRequestsResponse,
    SocialShareUsersResponse,
    UserResponse,
    UsersResponse,
    RankingUsersResponse,
    UserEmailResponse,
    HimaUsersResponse,
    UsersByTimestampResponse,
    UserTimestampResponse,
)
from ..utils import md5, sha256


class UserAPI(object):
    def __init__(self, base: client.BaseClient) -> None:
        self.__base = base

    def delete_footprint(self, user_id: int, footprint_id: int):
        return self.__base._request(
            "DELETE",
            route=f"/v2/users/{user_id}/footprints/{footprint_id}",
        )

    def destroy_user(self):
        return self.__base._request(
            "POST",
            route=f"/v2/users/destroy",
            payload={
                "uuid": self.__base.uuid,
                "api_key": Configs.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__base.uuid, int(datetime.now().timestamp()), True
                ),
            },
        )

    def follow_user(self, user_id: int):
        return self.__base._request("POST", route=f"/v2/users/{user_id}/follow")

    def follow_users(self, user_ids: list[int]):
        return self.__base._request(
            "POST",
            route=f"/v2/users/follow",
            params={"user_ids[]": user_ids},
        )

    def get_active_followings(self, **params) -> ActiveFollowingsResponse:
        """

        Parameters:
        -----------

            - only_online: bool
            - from_loggedin_at: int = None

        """
        return self.__base._request(
            "GET",
            route=f"/v1/users/active_followings",
            params=params,
            data_type=ActiveFollowingsResponse,
        )

    def get_follow_recommendations(self, **params) -> FollowRecommendationsResponse:
        """

        Parameters:
        -----------

            - from_timestamp: int = None,
            - number: int = None,
            - sources: list[str] = None

        """
        return self.__base._request(
            "GET",
            route=f"/v1/friends",
            params=params,
            data_type=FollowRecommendationsResponse,
        )

    def get_follow_request(
        self, from_timestamp: int = None
    ) -> UsersByTimestampResponse:
        params = {}
        if from_timestamp:
            params["from_timestamp"] = from_timestamp
        return self.__base._request(
            "GET",
            route=f"/v2/users/follow_requests",
            params=params,
            data_type=UsersByTimestampResponse,
        )

    def get_follow_request_count(self) -> FollowRequestCountResponse:
        return self.__base._request(
            "GET",
            route=f"/v2/users/follow_requests_count",
            data_type=FollowRequestCountResponse,
        )

    def get_following_users_born(self, birthdate: int = None) -> UsersResponse:
        params = {}
        if birthdate:
            params["birthdate"] = birthdate
        return self.__base._request(
            "GET",
            route=f"/v1/users/following_born_today",
            params=params,
            data_type=UsersResponse,
        )

    def get_footprints(self, **params) -> FootprintsResponse:
        """

        Parameters:
        -----------

            - from_id: int = None
            - number: int = None
            - mode: str = None

        """
        return self.__base._request(
            "GET",
            route=f"/v2/users/footprints",
            params=params,
            data_type=FootprintsResponse,
        )

    def get_fresh_user(self, user_id: int) -> UserResponse:
        return self.__base._request(
            "GET",
            route=f"/v2/users/fresh/{user_id}",
            data_type=UserResponse,
        )

    def get_hima_users(self, **params) -> HimaUsersResponse:
        """

        Parameters:
        -----------

            - from_hima_id: int = None
            - number: int = None

        """
        return self.__base._request(
            "GET",
            route=f"/v2/users/hima_users",
            params=params,
            data_type=HimaUsersResponse,
        )

    def get_user_ranking(self, mode: str) -> RankingUsersResponse:
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
        return self.__base._request(
            "GET",
            route=f"/v1/web/users/ranking",
            params={"mode": mode},
            data_type=RankingUsersResponse,
        )

    def get_refresh_counter_requests(self) -> RefreshCounterRequestsResponse:
        return self.__base._request(
            "GET",
            route=f"/v1/users/reset_counters",
            data_type=RefreshCounterRequestsResponse,
        )

    def get_social_shared_users(self, **params) -> SocialShareUsersResponse:
        """

        Parameters:
        -----------

            - sns_name: str - (Required)
            - number: int - (Optional)
            - from_id: int - (Optional)

        """
        return self.__base._request(
            "GET",
            route=f"/v2/users/social_shared_users",
            params=params,
            data_type=SocialShareUsersResponse,
        )

    def get_timestamp(
        self,
    ) -> UserTimestampResponse:
        return self.__base._request(
            "GET",
            route=f"/v2/users/timestamp",
            data_type=UserTimestampResponse,
            bypass_delay=True,
        )

    def get_user(self, user_id: int) -> UserResponse:
        return self.__base._request(
            "GET", route=f"/v2/users/{user_id}", data_type=UserResponse
        )

    def get_user_email(self, user_id: int) -> UserEmailResponse:
        return self.__base._request(
            "GET",
            route=f"/v2/users/fresh/{user_id}",
            data_type=UserEmailResponse,
        )

    def get_user_followers(self, user_id: int, **params) -> FollowUsersResponse:
        """

        Parameters:
        -----------

            - user_id: int - (required)
            - from_follow_id: int - (optional)
            - followed_by_me: int - (optional)
            - number: int - (optional)

        """
        return self.__base._request(
            "GET",
            route=f"/v2/users/{user_id}/followers",
            params=params,
            data_type=FollowUsersResponse,
        )

    def get_user_followings(self, user_id: int, **params) -> FollowUsersResponse:
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
        return self.__base._request(
            "POST",
            route=f"/v2/users/{user_id}/list_followings",
            params=params,
            data_type=FollowUsersResponse,
        )

    def get_user_from_qr(self, qr: str) -> UserResponse:
        return self.__base._request(
            "GET",
            route=f"/v1/users/qr_codes/{qr}",
            data_type=UserResponse,
        )

    def get_user_without_leaving_footprint(self, user_id: int) -> UserResponse:
        return self.__base._request(
            "GET",
            route=f"/v2/users/info/{user_id}",
            data_type=UserResponse,
        )

    def get_users(self, user_ids: list[int]) -> UsersResponse:
        return self.__base._request(
            "GET",
            route=f"/v1/users/list_id",
            params={"user_ids[]": user_ids},
            data_type=UsersResponse,
            jwt_required=True,
        )

    def refresh_counter(self, counter: str):
        return self.__base._request(
            "POST",
            route=f"/v1/users/reset_counters",
            payload={"counter": counter},
        )

    def register(
        self,
        email: str,
        email_grant_token: str,
        password: str,
        nickname: str,
        birth_date: str,
        gender: int = -1,
        country_code: str = "JP",
        biography: str = None,
        prefecture: str = None,
        profile_icon_filename: str = None,
        cover_image_filename: str = None,
        # @Nullable @Part("sns_info") SignUpSnsInfoRequest signUpSnsInfoRequest,
        en: int = None,
        vn: int = None,
    ) -> CreateUserResponse:
        return self.__base._request(
            "POST",
            route=f"/v3/users/register",
            payload={
                "app_version": Configs.API_VERSION_NAME,
                "api_key": Configs.API_KEY,
                "signed_version": sha256(),
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__base.device_uuid, int(datetime.now().timestamp()), False
                ),
                "uuid": self.__base.uuid,
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
            data_type=CreateUserResponse,
        )

    def remove_user_avatar(self):
        return self.__base._request("POST", route=f"/v2/users/remove_profile_photo")

    def remove_user_cover(self):
        return self.__base._request("POST", route=f"/v2/users/remove_cover_image")

    def report_user(
        self,
        user_id: int,
        category_id: int,
        reason: str = None,
        screenshot_filename: str = None,
        screenshot_2_filename: str = None,
        screenshot_3_filename: str = None,
        screenshot_4_filename: str = None,
    ):
        return self.__base._request(
            "POST",
            route=f"/v3/users/{user_id}/report",
            payload={
                "category_id": category_id,
                "reason": reason,
                "screenshot_filename": screenshot_filename,
                "screenshot_2_filename": screenshot_2_filename,
                "screenshot_3_filename": screenshot_3_filename,
                "screenshot_4_filename": screenshot_4_filename,
            },
        )

    def reset_password(self, email: str, email_grant_token: str, password: str):
        return self.__base._request(
            "PUT",
            route=f"/v1/users/reset_password",
            payload={
                "email": email,
                "email_grant_token": email_grant_token,
                "password": password,
            },
        )

    def search_lobi_users(self, **params) -> UsersResponse:
        """

        Parameters:
        -----------

            - nickname: str = None
            - number: int = None
            - from_str: str = None

        """
        return self.__base._request(
            "GET",
            route=f"/v1/lobi_friends",
            params=params,
            data_type=UsersResponse,
        )

    def search_users(self, **params) -> UsersResponse:
        """

        Parameters:
        -----------

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
        return self.__base._request(
            "GET",
            route=f"/v1/web/users/search",
            params=params,
            data_type=UsersResponse,
        )

    def set_follow_permission_enabled(self, nickname: str, is_private: bool = None):
        return self.__base._request(
            "POST",
            route=f"/v2/users/edit",
            payload={
                "nickname": nickname,
                "is_private": is_private,
                "uuid": self.__base.uuid,
                "api_key": Configs.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__base.uuid, int(datetime.now().timestamp()), True
                ),
                "signed_version": sha256(),
            },
        )

    def set_setting_follow_recommendation_enabled(self, on: bool):
        return self.__base._request(
            "POST",
            route=f"/v1/users/visible_on_sns_friend_recommendation_setting",
            params={"on": on},
        )

    def take_action_follow_request(self, target_id: int, action: str):
        return self.__base._request(
            "POST",
            route=f"/v2/users/{target_id}/follow_request",
            payload={"action": action},
        )

    def turn_on_hima(self):
        return self.__base._request(
            "GET",
            route=f"/v1/users/hima",
        )

    def unfollow_user(self, user_id: int):
        return self.__base._request("POST", route=f"/v2/users/{user_id}/unfollow")

    def update_language(self, language: str):
        return self.__base._request(
            "POST",
            route=f"/v1/users/language",
            payload={
                "uuid": self.__base.uuid,
                "api_key": Configs.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__base.uuid, int(datetime.now().timestamp()), True
                ),
                "language": language,
            },
        )

    def update_user(self, nickname: str, **params):
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
                "uuid": self.__base.uuid,
                "api_key": Configs.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__base.uuid, int(datetime.now().timestamp()), True
                ),
            }
        )
        return self.__base._request(
            "POST",
            route=f"/v3/users/edit",
            payload=params,
        )

    # BlockApi

    def block_user(
        self,
        user_id: int,
    ):
        return self.__base._request("POST", route=f"/v1/users/{user_id}/block")

    def get_blocked_user_ids(self) -> BlockedUserIdsResponse:
        return self.__base._request(
            "GET",
            route=f"/v1/users/block_ids",
            data_type=BlockedUserIdsResponse,
        )

    def get_blocked_users(self, from_id: int = None) -> BlockedUsersResponse:
        # @Body @NotNull SearchUsersRequest searchUsersRequest
        params = {}
        if from_id:
            params["from_id"] = from_id
        return self.__base._request(
            "GET",
            route=f"/v2/users/blocked",
            params=params,
            data_type=BlockedUsersResponse,
        )

    def unblock_user(self, user_id: int):
        return self.__base._request("POST", route=f"/v2/users/{user_id}/unblock")

    # HiddenApi

    def get_hidden_users_list(self, **params) -> HiddenResponse:
        """

        Parameters:
        -----------

            - from: str = None
            - number: int = None

        """
        return self.__base._request(
            "GET",
            route=f"/v1/hidden/users",
            params=params,
            data_type=HiddenResponse,
        )

    def hide_user(self, user_id: int):
        return self.__base._request(
            "POST",
            route=f"/v1/hidden/users",
            payload={"user_id": user_id},
        )

    def unhide_users(self, user_ids: list[int]):
        return self.__base._request(
            "DELETE",
            route=f"/v1/hidden/users",
            params={"user_ids[]": user_ids},
        )
