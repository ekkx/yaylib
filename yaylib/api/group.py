from datetime import datetime
from typing import Union, List

from ..config import *
from ..errors import *
from ..models import *
from ..responses import *
from ..utils import *


def accept_moderator_offer(self, group_id: int):
    self._check_authorization()
    return self._make_request(
        "PUT", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/deputize"
    )


def accept_ownership_offer(self, group_id: int):
    self._check_authorization()
    return self._make_request(
        "PUT", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/transfer"
    )


def accept_group_join_request(self, group_id: int, user_id: int):
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/accept/{user_id}"
    )


def add_related_groups(self, group_id: int, related_group_id: List[int]):
    self._check_authorization()
    return self._make_request(
        "PUT", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/related",
        params={"related_group_id[]": related_group_id}
    )


def ban_group_user(self, group_id: int, user_id: int):
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/ban/{user_id}"
    )


def check_unread_status(self, from_time: int = None) -> UnreadStatusResponse:
    self._check_authorization()
    params = {}
    if from_time:
        params["from_time"] = from_time
    return self._make_request(
        "GET", endpoint=f"{Endpoints.GROUPS_V1}/unread_status",
        params=params, data_type=UnreadStatusResponse
    )


def create_group(
        self, topic: str, description: str = None, secret: bool = None,
        hide_reported_posts: bool = None, hide_conference_call: bool = None, is_private: bool = None,
        only_verified_age: bool = None, only_mobile_verified: bool = None,
        call_timeline_display: bool = None, allow_ownership_transfer: bool = None,
        allow_thread_creation_by: str = None, gender: int = None, generation_groups_limit: int = None,
        group_category_id: int = None, cover_image_filename: str = None, sub_category_id: str = None,
        hide_from_game_eight: bool = None, allow_members_to_post_media: bool = None,
        allow_members_to_post_url: bool = None, guidelines: str = None,
) -> CreateGroupResponse:
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.GROUPS_V3}/new",
        payload={
            "topic": topic, "description": description, "secret": secret,
            "hide_reported_posts": hide_reported_posts,
            "hide_conference_call": hide_conference_call, "is_private": is_private,
            "only_verified_age": only_verified_age,
            "only_mobile_verified": only_mobile_verified,
            "call_timeline_display": call_timeline_display,
            "allow_ownership_transfer": allow_ownership_transfer,
            "allow_thread_creation_by": allow_thread_creation_by,
            "gender": gender, "generation_groups_limit": generation_groups_limit,
            "group_category_id": group_category_id,
            "cover_image_filename": cover_image_filename,
            "sub_category_id": sub_category_id,
            "hide_from_game_eight": hide_from_game_eight,
            "allow_members_to_post_image_and_video": allow_members_to_post_media,
            "allow_members_to_post_url": allow_members_to_post_url,
            "guidelines": guidelines,
        }, data_type=CreateGroupResponse
    )


def create_pin_group(self, group_id: int):
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.PINNED_V1}/groups",
        payload={"id": group_id}
    )


def decline_moderator_offer(self, group_id: int):
    self._check_authorization()
    return self._make_request(
        "DELETE", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/deputize"
    )


def decline_ownership_offer(self, group_id: int):
    self._check_authorization()
    return self._make_request(
        "DELETE", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/transfer"
    )


def decline_group_join_request(self, group_id: int, user_id: int):
    self._check_authorization()
    return self._make_request(
        "POST", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/decline/{user_id}"
    )


def delete_pin_group(self, group_id: int):
    self._check_authorization()
    return self._make_request(
        "DELETE", endpoint=f"{Endpoints.PINNED_V1}/groups/{group_id}"
    )


def get_banned_group_members(self, group_id: int, page: int = None) -> UsersResponse:
    self._check_authorization()
    params = {}
    if page:
        params["page"] = page
    return self._make_request(
        "GET", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/ban_list",
        params=params
    )


def get_categories(self, **params) -> GroupCategoriesResponse:
    pass


def get_create_group_quota(self) -> CreateGroupQuota:
    response = self._make_request(
        "GET", endpoint=f"{Endpoints.GROUPS_V1}/created_quota",
        data_type=CreateGroupQuota
    )
    return response.create


def get_group(self, group_id: int) -> GroupResponse:
    return self._make_request(
        "GET", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}",
        data_type=GroupResponse
    )


def get_group_notification_settings(self, group_id: int) -> GroupNotificationSettingsResponse:
    pass


def get_groups(self, **params) -> GroupsResponse:
    """

    Parameters:
    ----------

        - group_category_id: int = None
        - keyword: str = None
        - from_timestamp: int = None
        - sub_category_id: int = None

    """
    return self._make_request(
        "GET", endpoint=f"{Endpoints.GROUPS_V2}",
        params=params, data_type=GroupsResponse
    )


def get_invitable_users(self, group_id: int, **params) -> UsersByTimestampResponse:
    """

    Parameters:
    ----------

        - from_timestamp: int - (optional)
        - user[nickname]: str - (optional)

    """
    return self._make_request(
        "GET", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/users/invitable",
        params=params, data_type=UsersByTimestampResponse
    )


def get_joined_statuses(self, ids: List[int]) -> dict:
    self._check_authorization()
    return self._make_request(
        "GET", endpoint=f"{Endpoints.GROUPS_V1}/joined_statuses",
        params={"ids[]": ids}
    )


def get_group_member(self, group_id: int, user_id: int) -> GroupUserResponse:
    return self._make_request(
        "GET", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/members/{user_id}",
    )


def get_group_members(self, group_id: int, **params) -> GroupUsersResponse:
    """

    Parameters:
    ----------

        - id: int - (required)
        - mode: str - (optional)
        - keyword: str - (optional)
        - from_id: int - (optional)
        - from_timestamp: int - (optional)
        - order_by: str - (optional)
        - followed_by_me: bool - (optional)

    """
    return self._make_request(
        "GET", endpoint=f"{Endpoints.GROUPS_V2}/{group_id}/members",
        params=params, data_type=GroupUsersResponse
    )


def get_my_groups(self, from_timestamp: None) -> GroupsResponse:
    self._check_authorization()
    params = {}
    if from_timestamp:
        params["from_timestamp"] = from_timestamp
    return self._make_request(
        "GET", endpoint=f"{Endpoints.GROUPS_V2}/mine",
        params=params, data_type=GroupsResponse
    )


def get_relatable_groups(self, group_id: int, **params) -> GroupsRelatedResponse:
    """

    Parameters:
    ----------

        - group_id: int - (required)
        - keyword: str - (optional)
        - from: str - (optional)

    """
    return self._make_request(
        "GET", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/relatable",
        params=params, data_type=GroupsRelatedResponse
    )


def get_related_groups(self, group_id: int, **params) -> GroupsRelatedResponse:
    """

    Parameters:
    ----------

        - group_id: int - (required)
        - keyword: str - (optional)
        - from: str - (optional)

    """
    return self._make_request(
        "GET", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/related",
        params=params, data_type=GroupsRelatedResponse
    )


def get_user_groups(self, **params) -> GroupsResponse:
    """

    Parameters:
    ----------

        - user_id: int - (required)
        - page: int - (optional)

    """
    return self._make_request(
        "GET", endpoint=f"{Endpoints.GROUPS_V1}/user_group_list",
        params=params, data_type=GroupsResponse
    )


def invite_users_to_group(self, group_id: int, user_ids: List[int]):
    pass


def join_group(self, group_id: int):
    pass


def leave_group(self, group_id: int):
    pass


def post_gruop_social_shared(self, group_id: int, sns_name: str):
    pass


def remove_group_cover(self, group_id: int):
    pass


def remove_moderator(self, group_id: int, user_id: int):
    pass


def remove_related_groups(self, group_id: int, related_group_ids: List[int]):
    pass


def report_group(
        self, group_id: int, category_id: int, reason: str = None, opponent_id: int = None,
        screenshot_filename: str = None, screenshot_2_filename: str = None,
        screenshot_3_filename: str = None, screenshot_4_filename: str = None,
):
    pass


def send_moderator_offers(self, group_id: int, user_ids: List[int]):
    pass


def send_ownership_offer(self, group_id: int, user_id: int):
    pass


def set_group_notification_settings(
        self, group_id: int, notification_group_post: int = None,
        notification_group_join: int = None, notification_group_request: int = None,
        notification_group_message_tag_all: int = None,
) -> AdditionalSettingsResponse:
    pass


def set_group_title(self, group_id: int, title: str):
    pass


def take_over_group_ownership(self, group_id: int):
    pass


def unban_group_member(self, group_id: int, user_id: int):
    pass


def update_group(
        self, topic: str, description: str = None, secret: bool = None, hide_reported_posts: bool = None,
        hide_conference_call: bool = None, is_private: bool = None, only_verified_age: bool = None,
        only_mobile_verified: bool = None, call_timeline_display: bool = None, allow_ownership_transfer: bool = None,
        allow_thread_creation_by: str = None, gender: int = None, generation_groups_limit: int = None,
        group_category_id: int = None, cover_image_filename: str = None, sub_category_id: str = None,
        hide_from_game_eight: bool = None, allow_members_to_post_image_and_video: bool = None,
        allow_members_to_post_url: bool = None, guidelines: str = None,
) -> GroupResponse:
    pass


def visit_group(self, group_id: int):
    pass


def withdraw_moderator_offer(self, group_id: int, user_id: int):
    pass


def withdraw_ownership_offer(self, group_id: int, user_id: int):
    pass
