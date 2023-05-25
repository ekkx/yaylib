from datetime import datetime
from typing import Dict, List

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
    pass


def delete_footprint(self, user_id: int, footprint_id: int):
    pass


def destroy_user(self):
    # @NotNull @Part("uuid") String str,
    # @NotNull @Part("api_key") String str2,
    # @Part("timestamp") long j2,
    # @NotNull @Part("signed_info") String str3
    pass


def follow_user(self, user_id: int):
    pass


def follow_users(self, user_ids: List[int]):
    pass


def get_active_followings(
        self,
        only_online: bool,
        from_loggedin_at: int = None
) -> ActiveFollowingsResponse:
    pass


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


def getFollowRecommendations(
        self,
        from_timestamp: int = None,
        number: int = None,
        sources: List[str] = None
) -> FollowRecommendationsResponse:
    pass


def get_follow_request(self, from_timestamp: int = None) -> UsersByTimestampResponse:
    pass


def get_follow_request_count(self) -> int:
    # FollowRequestCountResponse.users_count
    pass


def get_following_users_born(self, birthdate: int = None) -> UsersResponse:
    pass


def get_footprints(
        self,
        from_id: int = None,
        number: int = None,
        mode: str = None
) -> List[Footprint]:
    # FootprintsResponse.footprints
    pass


def get_fresh_user(self, user_id: int) -> UserResponse:
    pass


def get_hima_users(
        self,
        from_hima_id: int = None,
        number: int = None
) -> List[UserWrapper]:
    pass


def get_initial_recommended_users_to_follow(
        en: int = None,
        vn: int = None
) -> UsersResponse:
    pass


def get_recommended_users_to_follow_for_profile(
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
    pass


def get_user(self, user_id: int) -> User:
    response = self._make_request(
        "GET", endpoint=f"{Endpoints.USER_V2}/{user_id}",
        data_type=UserResponse
    )
    return response.user


def get_user_custom_definitions(self) -> UserCustomDefinitionsResponse:
    pass


def get_user_email(self, user_id: int) -> str:
    # UserEmailResponse.email
    pass


def get_user_followers(
        self,
        user_id: int,
        from_follow_id: int = None,
        followed_by_me: bool = None,
        # @Nullable @Query("user[nickname]") String str
) -> FollowUsersResponse:
    pass


def get_user_followings(
        self,
        user_id: int,
        from_follow_id: int = None,
        from_timestamp: int = None,
        order_by: str = None,
        # @Body @Nullable SearchUsersRequest searchUsersRequest
) -> FollowUsersResponse:
    pass


def get_user_from_qr(self, qr: str) -> UserResponse:
    pass


def getUserInterests(self):
    pass


def get_user_without_leaving_footprint(
        self, user_id: int
) -> UserResponse:
    pass


def get_users(self, user_ids: List[int]) -> UsersResponse:
    # @Header("X-Jwt") @NotNull String str,
    pass


def get_users_from_uuid(self, uuid: str) -> UsersResponse:
    pass


def post_social_shared(self, sns_name: str):
    pass


def record_app_review_status(self, uuid: str):
    pass


def reduce_kenta_penalty(self, user_id: int):
    # @NotNull @Part("app_version") String str,
    # @Part("timestamp") long j3,
    # @NotNull @Part("api_key") String str2,
    # @NotNull @Part("signed_version") String str3,
    # @NotNull @Part("signed_info") String str4,
    # @NotNull @Part("uuid") String str5
    pass


def refresh_counter(self, counter: str):
    pass


def remove_user_avatar(self):
    pass


def remove_user_cover(self):
    pass


def report_user(
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


def search_users(
        self,
        gender: int = None,
        nickname: str = None,
        title: str = None,
        biography: str = None,
        from_timestamp: int = None,
        similar_age: bool = None,
        not_recent_gomimushi: bool = None,
        recently_created: bool = None,
        same_prefecture: bool = None,
        save_recent_search: bool = None
) -> UsersResponse:
    pass


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
    pass


def contact_friends(self):
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USER_V1}/contact_friends",
    )


def follow(self, user_id: int):
    return self._make_request("GET", endpoint=f"{Endpoints.USER_V2}/{str(user_id)}/follow", data_type=User)


# BlockApi, HiddenApi
