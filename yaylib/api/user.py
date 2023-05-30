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
    # @NotNull @Part("app_version") String str,
    # @Part("timestamp") long j2,
    # @NotNull @Part("api_key") String str2,
    # @NotNull @Part("signed_version") String str3,
    # @NotNull @Part("signed_info") String str4,
    # @NotNull @Part("uuid") String str5,
    pass


def delete_contact_friends(self):
    self._check_authorization()
    return self._make_request(
        "DELETE", endpoint=f"{Endpoints.USERS_V1}/contact_friends"
    )


def delete_footprint(self, user_id: int, footprint_id: int):
    self._check_authorization()
    return self._make_request(
        "DELETE", endpoint=f"{Endpoints.USERS_V2}/{user_id}/footprints/{footprint_id}"
    )


def destroy_user(self):
    # @NotNull @Part("uuid") String str,
    # @NotNull @Part("api_key") String str2,
    # @Part("timestamp") long j2,
    # @NotNull @Part("signed_info") String str3
    pass


def follow_user(self, user_id: int):
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V2}/{user_id}/follow"
    )


def follow_users(self, user_ids: List[int]):
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V2}/follow",
        params={"user_ids[]": user_ids}
    )


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
    # @Path("uuid") @NotNull String str
    pass


def get_contact_status(self, mobile_numbers: List[str]) -> ContactStatusResponse:
    pass


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
    response = self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V2}/follow_requests_count",
        data_type=FollowRequestCountResponse
    )
    return response.users_count


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
    response = self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V2}/footprints",
        params=params, data_type=FootprintsResponse
    )
    return response.footprints


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
    response = self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V2}/hima_users",
        params=params, data_type=HimaUsersResponse
    )
    return response.hima_users


def get_initial_recommended_users_to_follow(
        self,
        en: int = None,
        vn: int = None
) -> UsersResponse:
    pass


def get_recommended_users_to_follow_for_profile(
        self,
        user_id: int,
        number: int = None,
        page: int = None
) -> UsersResponse:
    pass


def get_refresh_counter_requests(self) -> RefreshCounterRequestsResponse:
    pass


def get_social_shared_users(
        self,
        sns_name: str,
        number: int = None,
        from_id: int = None
) -> SocialShareUsersResponse:
    pass


def get_timestamp(self) -> UserTimestampResponse:
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V2}/timestamp",
        data_type=UserTimestampResponse
    )


def get_user(self, user_id: int) -> User:
    response = self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V2}/{user_id}",
        data_type=UserResponse
    )
    return response.user


def get_user_custom_definitions(self) -> UserCustomDefinitionsResponse:
    pass


def get_user_email(self, user_id: int) -> str:
    # UserEmailResponse.email
    pass


def get_user_followers(self, user_id: int, **params) -> FollowUsersResponse:
    """

    Parameters:
    ----------

        - user_id: int
        - from_follow_id: int = None
        - followed_by_me: int = None

    """
    self._check_authorization()
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

    """
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V2}/{user_id}/list_followings",
        params=params, data_type=FollowUsersResponse
    )


def get_user_from_qr(self, qr: str) -> UserResponse:
    pass


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
    pass


def post_social_shared(self, sns_name: str):
    pass


def record_app_review_status(self, uuid: str):
    pass


def reduce_kenta_penalty(self, user_id: int):
    # TODO: yay_version_message とは
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"api/v3/users/{user_id}/reduce_penalty",
        payload={
            "uuid": self.uuid,
            "api_key": self.api_key,
            "timestamp": int(datetime.now().timestamp()),
            "signed_info": signed_info_calculating(
                self.api_key, self.device_uuid,
                int(datetime.now().timestamp())
            ),
            "signed_version": signed_version_calculating(
                self.yay_version_message, self.api_version_key
            )
        }, data_type=CreatePostResponse
    )


def refresh_counter(self, counter: str):
    pass


def remove_user_avatar(self):
    pass


def remove_user_cover(self):
    pass


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
    pass


def reset_password(self, email: str, email_grant_token: str, password: str):
    pass


def search_lobi_users(
        self,
        nickname: str = None,
        number: int = None,
        from_str: str = None
) -> UsersResponse:
    pass


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


def set_additional_setting_enabled(
        self,
        mode: str,
        on: int
):
    pass


def set_follow_permission_enabled(self):
    # @NotNull @Part("nickname") String str,
    # @Part("is_private") boolean z2,
    # @NotNull @Part("uuid") String str2,
    # @NotNull @Part("api_key") String str3,
    # @Part("timestamp") long j2,
    # @NotNull @Part("signed_info") String str4
    pass


def set_setting_follow_recommendation_enabled(self, on: bool):
    pass


def take_action_follow_request(self, target_id: int, action: str):
    pass


def turn_on_hima(self):
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USERS_V1}/hima",
    )


def unfollow_user(self, user_id: int):
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V2}/{user_id}/unfollow",
    )


def update_invite_contact_status(self, mobile_number: str):
    pass


def update_language(self, language: str):
    pass


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
    return self._make_request(
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
            "timestamp": int(datetime.now().timestamp()),
            "signed_info": signed_info_calculating(
                self.api_key, self.device_uuid,
                int(datetime.now().timestamp())
            ),
        }
    )


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
    pass


# BlockApi


def block_user(self, user_id: int):
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V1}/{user_id}/block",
    )


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
    return self._make_request(
        "POST", endpoint=f"{Endpoints.USERS_V2}/{user_id}/unblock"
    )


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
    return self._make_request(
        "POST", endpoint=f"{Endpoints.HIDDEN_V1}/users",
        payload={"user_id": user_id}
    )


def unhide_users(self, user_ids: List[int]):
    self._check_authorization()
    return self._make_request(
        "DELETE", endpoint=f"{Endpoints.HIDDEN_V1}/users",
        params={"user_ids[]": user_ids}
    )
