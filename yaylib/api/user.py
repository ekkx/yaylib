from datetime import datetime
from typing import Union, Dict, List

from ..config import *
from ..errors import *
from ..models import *
from ..responses import *
from ..utils import *


def create_user(
        self,
        nickname: str,
        birth_date: str,
        gender: int = -1,
        country_code: str = "JP",
        biography: str = None,
        prefecture: str = None,
        profile_icon_filename: str = None,
        cover_image_filename: str = None,
        # @Nullable @Part("sns_info") SignUpSnsInfoRequest signUpSnsInfoRequest,
        email: str = None,
        password: str = None,
        email_grant_token: str = None,
        en: int = None,
        vn: int = None
) -> CreateUserResponse:
    timestamp = int(datetime.now().timestamp())
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V3}/edit",
        payload={
            "app_version": self.api_version,
            "timestamp": timestamp,
            "api_key": self.api_key,
            "signed_version": signed_version_calculating(),
            "signed_info": signed_info_calculating(
                self.device_uuid, timestamp
            ),
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
        }
    )
    logger(self, fname="create_user")
    return response


def delete_contact_friends(self):
    self._check_authorization()
    response = self._make_request(
        "DELETE", endpoint=f"{Endpoints.USERS_V1}/contact_friends"
    )
    logger(self, fname="delete_contact_friends")
    return response


def delete_footprint(self, user_id: int, footprint_id: int):
    self._check_authorization()
    response = self._make_request(
        "DELETE", endpoint=f"{Endpoints.USERS_V2}/{user_id}/footprints/{footprint_id}"
    )
    logger(self, fname="delete_footprint")
    return response


def destroy_user(self):
    self._check_authorization()
    answer = input("Are you sure you want to delete your account? Y/N")
    if answer.lower() != "y":
        return
    timestamp = int(datetime.now().timestamp())
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V2}/destroy",
        payload={
            "uuid": self.uuid,
            "api_key": self.api_key,
            "timestamp": timestamp,
            "signed_info": signed_info_calculating(
                self.device_uuid, timestamp
            ),
        }
    )
    logger(self, fname="destroy_user")
    return response


def follow_user(self, user_id: int):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V2}/{user_id}/follow"
    )
    logger(self, fname="follow_user", user_id=user_id)
    return response


def follow_users(self, user_ids: List[int]):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V2}/follow",
        params={"user_ids[]": user_ids}
    )
    logger(self, fname="follow_users")
    return response


def get_active_followings(self, **params) -> ActiveFollowingsResponse:
    """

    Parameters:
    ----------

        - only_online: bool
        - from_loggedin_at: int = None

    """
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V1}/active_followings",
        params=params, data_type=ActiveFollowingsResponse
    )


def get_additional_settings(self) -> Settings:
    # AdditionalSettingsResponse
    pass


def get_app_review_status(self) -> AppReviewStatusResponse:
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V1}/{self.url_uuid}/app_review_status",
        payload={"uuid": self.uuid}, data_type=AppReviewStatusResponse
    )


def get_contact_status(self, mobile_numbers: List[str]) -> ContactStatusResponse:
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V1}/contact_status",
        payload={"mobile_numbers[]": mobile_numbers}, data_type=ContactStatusResponse
    )


def get_default_settings(self) -> TimelineSettings:
    # DefaultSettingsResponse
    pass


def get_follow_recommendations(self, **params) -> FollowRecommendationsResponse:
    """

    Parameters:
    ----------

        - from_timestamp: int = None,
        - number: int = None,
        - sources: List[str] = None

    """
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.FRIENDS_V1}",
        params=params, data_type=ActiveFollowingsResponse
    )


def get_follow_request(self, from_timestamp: int = None) -> UsersByTimestampResponse:
    params = {}
    if from_timestamp:
        params["from_timestamp"] = from_timestamp
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V2}/follow_requests",
        params=params, data_type=UsersByTimestampResponse
    )


def get_follow_request_count(self) -> int:
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V2}/follow_requests_count",
        data_type=FollowRequestCountResponse
    ).users_count


def get_following_users_born(self, birthdate: int = None) -> UsersResponse:
    params = {}
    if birthdate:
        params["birthdate"] = birthdate
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V1}/following_born_today",
        params=params, data_type=UsersResponse
    )


def get_footprints(self, **params) -> List[Footprint]:
    """

    Parameters:
    ----------

        - from_id: int = None
        - number: int = None
        - mode: str = None

    """
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V2}/footprints",
        params=params, data_type=FootprintsResponse
    ).footprints


def get_fresh_user(self, user_id: int) -> UserResponse:
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V2}/fresh/{user_id}",
        data_type=UserResponse
    )


def get_hima_users(self, **params) -> List[UserWrapper]:
    """

    Parameters:
    ----------

        - from_hima_id: int = None
        - number: int = None

    """
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V2}/hima_users",
        params=params, data_type=HimaUsersResponse
    ).hima_users


def get_initial_recommended_users_to_follow(self, **params) -> UsersResponse:
    """

    Parameters:
    ----------

        - en: int - (Optional)
        - vn: int - (Optional)

    """
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V1}/initial_follow_recommended",
        params=params, data_type=UsersResponse
    )


def get_recommended_users_to_follow_for_profile(self, user_id: int, **params) -> UsersResponse:
    """

    Parameters:
    ----------

        - user_id: int - (Required)
        - number: int - (Optional)
        - page: int - (Optional)

    """
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V1}/{user_id}/follow_recommended",
        params=params, data_type=UsersResponse
    )


def get_refresh_counter_requests(self) -> RefreshCounterRequestsResponse:
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V1}/reset_counters",
        data_type=RefreshCounterRequestsResponse
    )


def get_social_shared_users(self, **params) -> SocialShareUsersResponse:
    """

    Parameters:
    ----------

        - sns_name: str - (Required)
        - number: int - (Optional)
        - from_id: int - (Optional)

    """
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V2}/social_shared_users",
        params=params, data_type=SocialShareUsersResponse
    )


def get_timestamp(self) -> UserTimestampResponse:
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V2}/timestamp",
        data_type=UserTimestampResponse
    )


def get_user(self, user_id: int) -> User:
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V2}/{user_id}",
        data_type=UserResponse
    ).user


def get_user_custom_definitions(self) -> UserCustomDefinitionsResponse:
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V1}/custom_definitions",
        data_type=UserCustomDefinitionsResponse
    )


def get_user_email(self, user_id: int) -> str:
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V2}/fresh/{user_id}",
        data_type=UserEmailResponse
    ).email


def get_user_followers(self, user_id: int, **params) -> FollowUsersResponse:
    """

    Parameters:
    ----------

        - user_id: int - (required)
        - from_follow_id: int - (optional)
        - followed_by_me: int - (optional)
        - number: int - (optional)

    """
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V2}/{user_id}/followers",
        params=params, data_type=FollowUsersResponse
    )


def get_user_followings(self, user_id: int, **params) -> FollowUsersResponse:
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
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V2}/{user_id}/list_followings",
        params=params, data_type=FollowUsersResponse
    )


def get_user_from_qr(self, qr: str) -> UserResponse:
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V1}/qr_codes/{qr}",
        data_type=UserResponse
    )


def get_user_interests(self):
    # @NotNull Continuation<? super UserInterestsResponse> continuation
    pass


def get_user_without_leaving_footprint(self, user_id: int) -> UserResponse:
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V2}/info/{user_id}",
        data_type=UserResponse
    )


def get_users(self, user_ids: List[int]) -> UsersResponse:
    headers = self.session.headers
    headers["X-Jwt"] = self.get_web_socket_token()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V1}/list_id",
        params={"user_ids[]": user_ids}, data_type=UsersResponse,
        headers=headers
    )


def get_users_from_uuid(self, uuid: str) -> UsersResponse:
    # TODO: what it does?
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V2}/list_uuid/",
        params={"uuid": uuid}, data_type=UsersResponse
    )


def post_social_shared(self, sns_name: str):
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V2}/social_shared",
        params={"sns_name": sns_name}
    )
    logger(self, fname="post_social_shared")
    return response


def record_app_review_status(self):
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V1}/{self.url_uuid}/app_review_status",
        params={"uuid": self.uuid}
    )
    logger(self, fname="record_app_review_status")
    return response


def reduce_kenta_penalty(self, user_id: int):
    self._check_authorization()
    timestamp = int(datetime.now().timestamp())
    response = self._make_request(
        "POST", endpoint=f"{self.host}api/v3/users/{user_id}/reduce_penalty",
        payload={
            "uuid": self.uuid,
            "api_key": self.api_key,
            "timestamp": timestamp,
            "signed_info": signed_info_calculating(
                self.device_uuid, timestamp
            ),
            "signed_version": signed_version_calculating()
        }, data_type=CreatePostResponse
    )
    logger(self, fname="reduce_kenta_penalty")
    return response


def refresh_counter(self, counter: str):
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V1}/reset_counters",
        payload={"counter": counter}
    )
    logger(self, fname="refresh_counter")
    return response


def remove_user_avatar(self):
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V2}/remove_profile_photo"
    )
    logger(self, fname="remove_user_avatar")
    return response


def remove_user_cover(self):
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V2}/remove_cover_image"
    )
    logger(self, fname="remove_user_cover")
    return response


def report_user(
        self,
        user_id: int,
        category_id: int,
        reason: str = None,
        screenshot_filename: str = None,
        screenshot_2_filename: str = None,
        screenshot_3_filename: str = None,
        screenshot_4_filename: str = None
):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V3}/{user_id}/report",
        payload={
            "category_id": category_id,
            "reason": reason,
            "screenshot_filename": screenshot_filename,
            "screenshot_2_filename": screenshot_2_filename,
            "screenshot_3_filename": screenshot_3_filename,
            "screenshot_4_filename": screenshot_4_filename,
        }
    )
    logger(self, fname="report_user", user_id=user_id)
    return response


def reset_password(self, email: str, email_grant_token: str, password: str):
    response = self._make_request(
        "PUT", endpoint=f"{Endpoints.USERS_V1}/reset_password",
        payload={
            "email": email,
            "email_grant_token": email_grant_token,
            "password": password,
        }
    )
    logger(self, fname="reset_password")
    return response


def search_lobi_users(self, **params) -> UsersResponse:
    """

    Parameters:
    ----------

        - nickname: str = None
        - number: int = None
        - from_str: str = None

    """
    return self._make_request(
        "GET", endpoint=f"{Endpoints.LOBI_FRIENDS_V1}",
        params=params, data_type=UsersResponse
    )


def search_users(self, **params) -> UsersResponse:
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
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V1}/search",
        params=params, data_type=UsersResponse,
    )


def set_additional_setting_enabled(self, mode: str, on: int = None):
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V1}/additonal_notification_setting",
        payload={"mode": mode, "on": on}
    )
    logger(self, fname="set_additional_setting_enabled")
    return response


def set_follow_permission_enabled(self, nickname: str, is_private: bool = None):
    timestamp = int(datetime.now().timestamp())
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V2}/edit",
        payload={
            "nickname": nickname,
            "is_private": is_private,
            "uuid": self.uuid,
            "api_key": self.api_key,
            "timestamp": timestamp,
            "signed_info": signed_info_calculating(
                self.device_uuid, timestamp
            ),
            "signed_version": signed_version_calculating()
        }
    )
    logger(self, fname="set_follow_permission_enabled")
    return response


def set_setting_follow_recommendation_enabled(self, on: bool):
    response = self._make_request(
        "POST",
        endpoint=f"{Endpoints.USERS_V1}/visible_on_sns_friend_recommendation_setting",
        params={"on": on}
    )
    logger(self, fname="set_setting_follow_recommendation_enabled")
    return response


def take_action_follow_request(self, target_id: int, action: str):
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V2}/{target_id}/follow_request",
        payload={"action": action}
    )
    logger(self, fname="take_action_follow_request")
    return response


def turn_on_hima(self):
    self._check_authorization()
    response = self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V1}/hima",
    )
    logger(self, fname="turn_on_hima")
    return response


def unfollow_user(self, user_id: int):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V2}/{user_id}/unfollow",
    )
    logger(self, fname="unfollow_user", user_id=user_id)
    return response


def update_invite_contact_status(self, mobile_number: str):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V1}/invite_contact",
        params={"mobile_number": mobile_number}
    )
    logger(self, fname="update_invite_contact_status")
    return response


def update_language(self, language: str):
    timestamp = int(datetime.now().timestamp())
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V1}/language",
        payload={
            "uuid": self.uuid,
            "api_key": self.api_key,
            "timestamp": timestamp,
            "signed_info": signed_info_calculating(
                self.device_uuid, timestamp),
            "language": language,
        }
    )
    logger(self, fname="update_language")
    return response


def update_user(
        self,
        nickname: str,
        biography: str = None,
        prefecture: str = None,
        gender: int = None,
        country_code: str = None,
        profile_icon_filename: str = None,
        cover_image_filename: str = None,
        username: str = None,
):
    self._check_authorization()
    timestamp = int(datetime.now().timestamp())
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V3}/edit",
        payload={
            "nickname": nickname,
            "biography": biography,
            "prefecture": prefecture,
            "gender": gender,
            "country_code": country_code,
            "profile_icon_filename": profile_icon_filename,
            "cover_image_filename": cover_image_filename,
            "username": username,
            "uuid": self.uuid,
            "api_key": self.api_key,
            "timestamp": timestamp,
            "signed_info": signed_info_calculating(
                self.device_uuid, timestamp
            ),
        }
    )
    logger(self, fname="update_user")
    return response


def update_user_interests(
        self,
        # @Body @NotNull CommonIdsRequest commonIdsRequest,
        # @NotNull Continuation<? super Unit> continuation
):
    pass


def upload_contacts_friends(
        self,
        # @Body @Nullable UploadContactsRequest uploadContactsRequest
):
    pass


def upload_twitter_friend_ids(self, twitter_friend_ids: List[str]):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V1}/twitter_friends",
        payload={"twitter_friend_ids[]": twitter_friend_ids}
    )
    logger(self, fname="upload_twitter_friend_ids")
    return response


# BlockApi


def block_user(self, user_id: int):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V1}/{user_id}/block",
    )
    logger(self, fname="block_user", user_id=user_id)
    return response


def get_blocked_user_ids(self) -> BlockedUserIdsResponse:
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V1}/block_ids",
        data_type=BlockedUserIdsResponse
    )


def get_blocked_users(self, from_id: int = None) -> BlockedUsersResponse:
    # @Body @NotNull SearchUsersRequest searchUsersRequest
    self._check_authorization()
    params = {}
    if from_id:
        params["from_id"] = from_id
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V2}/blocked",
        params=params, data_type=BlockedUsersResponse
    )


def unblock_user(self, user_id: int):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V2}/{user_id}/unblock"
    )
    logger(self, fname="unblock_user", user_id=user_id)
    return response


# HiddenApi


def get_hidden_users_list(self, **params: Union[str, int]) -> HiddenResponse:
    """

    Parameters:
    ----------

        - from: str = None
        - number: int = None

    """
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.HIDDEN_V1}/users",
        params=params, data_type=HiddenResponse
    )


def hide_user(self, user_id: int):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.HIDDEN_V1}/users",
        payload={"user_id": user_id}
    )
    logger(self, fname="hide_user", user_id=user_id)
    return response


def unhide_users(self, user_ids: List[int]):
    self._check_authorization()
    response = self._make_request(
        "DELETE", endpoint=f"{Endpoints.HIDDEN_V1}/users",
        params={"user_ids[]": user_ids}
    )
    logger(self, fname="unhide_users")
    return response
