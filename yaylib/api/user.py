from datetime import datetime
from typing import Dict, List

from ..config import *
from ..errors import *
from ..models import *
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


def contact_friends(self):
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USER_V1}/contact_friends",
    )


def get_user(self, user_id: int):
    return self._make_request(
        "GET", endpoint=f"{Endpoints.USER_V2}/{user_id}",
        data_type=User
    )


def follow(self, user_id: int):
    return self._make_request("GET", endpoint=f"{Endpoints.USER_V2}/{str(user_id)}/follow", data_type=User)


# BlockApi, HiddenApi
