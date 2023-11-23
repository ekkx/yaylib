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

from datetime import datetime

from ..config import Configs, Endpoints
from ..models import Footprint, User, UserWrapper
from ..responses import (
    ActiveFollowingsResponse,
    BlockedUserIdsResponse,
    BlockedUsersResponse,
    CreatePostResponse,
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


def delete_footprint(self, user_id: int, footprint_id: int, access_token: str = None):
    response = self.request(
        "DELETE",
        endpoint=f"{Endpoints.USERS_V2}/{user_id}/footprints/{footprint_id}",
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Footprint has been deleted.")
    return response


def destroy_user(self, access_token: str = None):
    timestamp = int(datetime.now().timestamp())
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.USERS_V2}/destroy",
        payload={
            "uuid": self.uuid,
            "api_key": self.api_key,
            "timestamp": timestamp,
            "signed_info": self.generate_signed_info(self.device_uuid, timestamp),
        },
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("User has been deleted.")
    return response


def follow_user(self, user_id: int, access_token: str = None):
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.USERS_V2}/{user_id}/follow",
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info(f"Followed the user '{user_id}'.")
    return response


def follow_users(self, user_ids: list[int], access_token: str = None):
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.USERS_V2}/follow",
        params={"user_ids[]": user_ids},
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Followed multiple users.")
    return response


def get_active_followings(
    self, access_token: str = None, **params
) -> ActiveFollowingsResponse:
    """

    Parameters:
    ----------

        - only_online: bool
        - from_loggedin_at: int = None

    """
    return self.request(
        "GET",
        endpoint=f"{Endpoints.USERS_V1}/active_followings",
        params=params,
        data_type=ActiveFollowingsResponse,
        auth_required=True,
        access_token=access_token,
    )


def get_follow_recommendations(
    self, access_token: str = None, **params
) -> FollowRecommendationsResponse:
    """

    Parameters:
    ----------

        - from_timestamp: int = None,
        - number: int = None,
        - sources: list[str] = None

    """
    return self.request(
        "GET",
        endpoint=f"{Endpoints.FRIENDS_V1}",
        params=params,
        data_type=FollowRecommendationsResponse,
        auth_required=True,
        access_token=access_token,
    )


def get_follow_request(
    self, from_timestamp: int = None, access_token: str = None
) -> UsersByTimestampResponse:
    params = {}
    if from_timestamp:
        params["from_timestamp"] = from_timestamp
    return self.request(
        "GET",
        endpoint=f"{Endpoints.USERS_V2}/follow_requests",
        params=params,
        data_type=UsersByTimestampResponse,
        auth_required=True,
        access_token=access_token,
    )


def get_follow_request_count(self, access_token: str = None) -> int:
    return self.request(
        "GET",
        endpoint=f"{Endpoints.USERS_V2}/follow_requests_count",
        data_type=FollowRequestCountResponse,
        auth_required=True,
        access_token=access_token,
    ).users_count


def get_following_users_born(
    self, birthdate: int = None, access_token: str = None
) -> UsersResponse:
    params = {}
    if birthdate:
        params["birthdate"] = birthdate
    return self.request(
        "GET",
        endpoint=f"{Endpoints.USERS_V1}/following_born_today",
        params=params,
        data_type=UsersResponse,
        auth_required=True,
        access_token=access_token,
    )


def get_footprints(self, access_token: str = None, **params) -> list[Footprint]:
    """

    Parameters:
    ----------

        - from_id: int = None
        - number: int = None
        - mode: str = None

    """
    return self.request(
        "GET",
        endpoint=f"{Endpoints.USERS_V2}/footprints",
        params=params,
        data_type=FootprintsResponse,
        auth_required=True,
        access_token=access_token,
    ).footprints


def get_fresh_user(self, user_id: int, access_token: str = None) -> UserResponse:
    return self.request(
        "GET",
        endpoint=f"{Endpoints.USERS_V2}/fresh/{user_id}",
        data_type=UserResponse,
        auth_required=True,
        access_token=access_token,
    )


def get_hima_users(self, access_token: str = None, **params) -> list[UserWrapper]:
    """

    Parameters:
    ----------

        - from_hima_id: int = None
        - number: int = None

    """
    return self.request(
        "GET",
        endpoint=f"{Endpoints.USERS_V2}/hima_users",
        params=params,
        data_type=HimaUsersResponse,
        auth_required=True,
        access_token=access_token,
    ).hima_users


def get_user_ranking(self, mode: str) -> RankingUsersResponse:
    """

    ユーザーのランキングを取得します

    Examples:
    --------

    >>> ルーキーを取得する場合:

    >>> api.get_user_ranking(mode="one_month")

    ---

    >>> ミドルを取得する場合:

    >>> api.get_user_ranking(mode="six_months")

    ---

    >>> 殿堂入りを取得する場合:

    >>> api.get_user_ranking(mode="all_time")

    """
    return self.request(
        "GET",
        endpoint=f"{Endpoints.WEB_V1}/users/ranking",
        params={"mode": mode},
        data_type=RankingUsersResponse,
    )


def get_refresh_counter_requests(
    self, access_token: str = None
) -> RefreshCounterRequestsResponse:
    return self.request(
        "GET",
        endpoint=f"{Endpoints.USERS_V1}/reset_counters",
        data_type=RefreshCounterRequestsResponse,
        auth_required=True,
        access_token=access_token,
    )


def get_social_shared_users(
    self, access_token: str = None, **params
) -> SocialShareUsersResponse:
    """

    Parameters:
    ----------

        - sns_name: str - (Required)
        - number: int - (Optional)
        - from_id: int - (Optional)

    """
    return self.request(
        "GET",
        endpoint=f"{Endpoints.USERS_V2}/social_shared_users",
        params=params,
        data_type=SocialShareUsersResponse,
        auth_required=True,
        access_token=access_token,
    )


def get_timestamp(self, access_token: str = None) -> UserTimestampResponse:
    return self.request(
        "GET",
        endpoint=f"{Endpoints.USERS_V2}/timestamp",
        data_type=UserTimestampResponse,
        bypass_delay=True,
        access_token=access_token,
    )


def get_user(self, user_id: int, access_token: str = None) -> User:
    return self.request(
        "GET",
        endpoint=f"{Endpoints.USERS_V2}/{user_id}",
        data_type=UserResponse,
        access_token=access_token,
    ).user


def get_user_email(self, user_id: int, access_token: str = None) -> str:
    return self.request(
        "GET",
        endpoint=f"{Endpoints.USERS_V2}/fresh/{user_id}",
        data_type=UserEmailResponse,
        auth_required=True,
        access_token=access_token,
    ).email


def get_user_followers(
    self, user_id: int, access_token: str = None, **params
) -> FollowUsersResponse:
    """

    Parameters:
    ----------

        - user_id: int - (required)
        - from_follow_id: int - (optional)
        - followed_by_me: int - (optional)
        - number: int - (optional)

    """
    return self.request(
        "GET",
        endpoint=f"{Endpoints.USERS_V2}/{user_id}/followers",
        params=params,
        data_type=FollowUsersResponse,
        access_token=access_token,
    )


def get_user_followings(
    self, user_id: int, access_token: str = None, **params
) -> FollowUsersResponse:
    # @Body @Nullable SearchUsersRequest searchUsersRequest
    """

    Parameters:
    ----------

        - user_id: int
        - from_follow_id: int = None
        - from_timestamp: int = None
        - order_by: str = None
        - number: int - (optional)

    """
    return self.request(
        "POST",
        endpoint=f"{Endpoints.USERS_V2}/{user_id}/list_followings",
        params=params,
        data_type=FollowUsersResponse,
        auth_required=True,
        access_token=access_token,
    )


def get_user_from_qr(self, qr: str, access_token: str = None) -> UserResponse:
    return self.request(
        "GET",
        endpoint=f"{Endpoints.USERS_V1}/qr_codes/{qr}",
        data_type=UserResponse,
        access_token=access_token,
    )


def get_user_without_leaving_footprint(
    self, user_id: int, access_token: str = None
) -> UserResponse:
    return self.request(
        "GET",
        endpoint=f"{Endpoints.USERS_V2}/info/{user_id}",
        data_type=UserResponse,
        access_token=access_token,
    )


def get_users(self, user_ids: list[int], access_token: str = None) -> UsersResponse:
    timestamp = int(datetime.now().timestamp())
    headers = self.session.headers.copy()
    headers["X-Jwt"] = self.generate_jwt(timestamp)
    return self.request(
        "GET",
        endpoint=f"{Endpoints.USERS_V1}/list_id",
        params={"user_ids[]": user_ids},
        data_type=UsersResponse,
        headers=headers,
        access_token=access_token,
    )


def refresh_counter(self, counter: str, access_token: str = None):
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.USERS_V1}/reset_counters",
        payload={"counter": counter},
        access_token=access_token,
    )
    self.logger.info("Requested counter refresh.")
    return response


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
    timestamp = int(datetime.now().timestamp())
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.USERS_V3}/register",
        payload={
            "app_version": self.api_version,
            "timestamp": timestamp,
            "api_key": self.api_key,
            "signed_version": self.generate_signed_version(),
            "signed_info": self.generate_signed_info(self.device_uuid, timestamp),
            "uuid": self.uuid,
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
    self.logger.info(f"A new user has been registered. (USER ID: {response.id})")
    return response


def remove_user_avatar(self, access_token: str = None):
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.USERS_V2}/remove_profile_photo",
        access_token=access_token,
    )
    self.logger.info("Profile image has been removed.")
    return response


def remove_user_cover(self, access_token: str = None):
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.USERS_V2}/remove_cover_image",
        access_token=access_token,
    )
    self.logger.info("Profile cover image has been removed.")
    return response


def report_user(
    self,
    user_id: int,
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
        endpoint=f"{Endpoints.USERS_V3}/{user_id}/report",
        payload={
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
    self.logger.info(f"Reported the user '{user_id}'.")
    return response


def reset_password(
    self, email: str, email_grant_token: str, password: str, access_token: str = None
):
    response = self.request(
        "PUT",
        endpoint=f"{Endpoints.USERS_V1}/reset_password",
        payload={
            "email": email,
            "email_grant_token": email_grant_token,
            "password": password,
        },
        access_token=access_token,
    )
    self.logger.info("Reset the password.")
    return response


def search_lobi_users(self, access_token: str = None, **params) -> UsersResponse:
    """

    Parameters:
    ----------

        - nickname: str = None
        - number: int = None
        - from_str: str = None

    """
    return self.request(
        "GET",
        endpoint=f"{Endpoints.LOBI_FRIENDS_V1}",
        params=params,
        data_type=UsersResponse,
        access_token=access_token,
    )


def search_users(self, access_token: str = None, **params) -> UsersResponse:
    """

    Parameters:
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
    return self.request(
        "GET",
        endpoint=f"{Endpoints.WEB_V1}/users/search",
        params=params,
        data_type=UsersResponse,
        access_token=access_token,
    )


def set_follow_permission_enabled(
    self, nickname: str, is_private: bool = None, access_token: str = None
):
    timestamp = int(datetime.now().timestamp())
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.USERS_V2}/edit",
        payload={
            "nickname": nickname,
            "is_private": is_private,
            "uuid": self.uuid,
            "api_key": self.api_key,
            "timestamp": timestamp,
            "signed_info": self.generate_signed_info(self.device_uuid, timestamp),
            "signed_version": self.generate_signed_version(),
        },
        access_token=access_token,
    )
    self.logger.info("Follow permission has been enabled.")
    return response


def set_setting_follow_recommendation_enabled(self, on: bool, access_token: str = None):
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.USERS_V1}/visible_on_sns_friend_recommendation_setting",
        params={"on": on},
        access_token=access_token,
    )
    self.logger.info("Follow recommendation has been enabled.")
    return response


def take_action_follow_request(
    self, target_id: int, action: str, access_token: str = None
):
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.USERS_V2}/{target_id}/follow_request",
        payload={"action": action},
        access_token=access_token,
    )
    self.logger.info("Took action follow request.")
    return response


def turn_on_hima(self, access_token: str = None):
    response = self.request(
        "GET",
        endpoint=f"{Endpoints.USERS_V1}/hima",
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Turned on 'hima now'.")
    return response


def unfollow_user(self, user_id: int, access_token: str = None):
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.USERS_V2}/{user_id}/unfollow",
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info(f"Unfollowed the user '{user_id}'.")
    return response


def update_language(self, language: str, access_token: str = None):
    timestamp = int(datetime.now().timestamp())
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.USERS_V1}/language",
        payload={
            "uuid": self.uuid,
            "api_key": self.api_key,
            "timestamp": timestamp,
            "signed_info": self.generate_signed_info(self.device_uuid, timestamp),
            "language": language,
        },
        access_token=access_token,
    )
    self.logger.info("Language has been updated.")
    return response


def update_user(self, nickname: str, access_token: str = None, **params):
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
        - access_token: str = (optional)

    """
    timestamp = int(datetime.now().timestamp())
    params.update(
        {
            "nickname": nickname,
            "uuid": self.uuid,
            "api_key": self.api_key,
            "timestamp": timestamp,
            "signed_info": self.generate_signed_info(
                self.uuid, timestamp, shared_key=True
            ),
        }
    )
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.USERS_V3}/edit",
        payload=params,
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("User profile has been updated.")
    return response


# BlockApi


def block_user(self, user_id: int, access_token: str = None):
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.USERS_V1}/{user_id}/block",
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info(f"Blocked the user '{user_id}'.")
    return response


def get_blocked_user_ids(self, access_token: str = None) -> BlockedUserIdsResponse:
    return self.request(
        "GET",
        endpoint=f"{Endpoints.USERS_V1}/block_ids",
        data_type=BlockedUserIdsResponse,
        auth_required=True,
        access_token=access_token,
    )


def get_blocked_users(
    self, from_id: int = None, access_token: str = None
) -> BlockedUsersResponse:
    # @Body @NotNull SearchUsersRequest searchUsersRequest
    params = {}
    if from_id:
        params["from_id"] = from_id
    return self.request(
        "GET",
        endpoint=f"{Endpoints.USERS_V2}/blocked",
        params=params,
        data_type=BlockedUsersResponse,
        auth_required=True,
        access_token=access_token,
    )


def unblock_user(self, user_id: int, access_token: str = None):
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.USERS_V2}/{user_id}/unblock",
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info(f"Unblocked the user '{user_id}'.")
    return response


# HiddenApi


def get_hidden_users_list(self, access_token: str = None, **params) -> HiddenResponse:
    """

    Parameters:
    ----------

        - from: str = None
        - number: int = None

    """
    return self.request(
        "GET",
        endpoint=f"{Endpoints.HIDDEN_V1}/users",
        params=params,
        data_type=HiddenResponse,
        auth_required=True,
        access_token=access_token,
    )


def hide_user(self, user_id: int, access_token: str = None):
    response = self.request(
        "POST",
        endpoint=f"{Endpoints.HIDDEN_V1}/users",
        payload={"user_id": user_id},
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info(f"User '{user_id}' is hidden.")
    return response


def unhide_users(self, user_ids: list[int], access_token: str = None):
    response = self.request(
        "DELETE",
        endpoint=f"{Endpoints.HIDDEN_V1}/users",
        params={"user_ids[]": user_ids},
        auth_required=True,
        access_token=access_token,
    )
    self.logger.info("Unhid users.")
    return response
