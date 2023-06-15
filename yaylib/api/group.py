from datetime import datetime
from typing import Union, List

from ..config import *
from ..errors import *
from ..models import *
from ..responses import *
from ..utils import *


def accept_moderator_offer(self, group_id: int):
    self._check_authorization()
    response = self._make_request(
        "PUT", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/deputize"
    )
    logger(self, fname="accept_moderator_offer")
    return response


def accept_ownership_offer(self, group_id: int):
    self._check_authorization()
    response = self._make_request(
        "PUT", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/transfer"
    )
    logger(self, fname="accept_ownership_offer")
    return response


def accept_group_join_request(self, group_id: int, user_id: int):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/accept/{user_id}"
    )
    logger(self, fname="accept_group_join_request")
    return response


def add_related_groups(self, group_id: int, related_group_id: List[int]):
    self._check_authorization()
    response = self._make_request(
        "PUT", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/related",
        params={"related_group_id[]": related_group_id}
    )
    logger(self, fname="add_related_groups")
    return response


def ban_group_user(self, group_id: int, user_id: int):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/ban/{user_id}"
    )
    logger(self, fname="ban_group_user")
    return response


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
        self,
        topic: str,
        description: str = None,
        secret: bool = None,
        hide_reported_posts: bool = None,
        hide_conference_call: bool = None,
        is_private: bool = None,
        only_verified_age: bool = None,
        only_mobile_verified: bool = None,
        call_timeline_display: bool = None,
        allow_ownership_transfer: bool = None,
        allow_thread_creation_by: str = None,
        gender: int = None,
        generation_groups_limit: int = None,
        group_category_id: int = None,
        cover_image_filename: str = None,
        sub_category_id: str = None,
        hide_from_game_eight: bool = None,
        allow_members_to_post_media: bool = None,
        allow_members_to_post_url: bool = None,
        guidelines: str = None,
) -> CreateGroupResponse:
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.GROUPS_V3}/new",
        payload={
            "topic": topic,
            "description": description,
            "secret": secret,
            "hide_reported_posts": hide_reported_posts,
            "hide_conference_call": hide_conference_call,
            "is_private": is_private,
            "only_verified_age": only_verified_age,
            "only_mobile_verified": only_mobile_verified,
            "call_timeline_display": call_timeline_display,
            "allow_ownership_transfer": allow_ownership_transfer,
            "allow_thread_creation_by": allow_thread_creation_by,
            "gender": gender,
            "generation_groups_limit": generation_groups_limit,
            "group_category_id": group_category_id,
            "cover_image_filename": cover_image_filename,
            "uuid": self.uuid,
            "api_key": self.api_key,
            "timestamp": int(datetime.now().timestamp()),
            "signed_info": signed_info_calculating(
                self.api_key, self.device_uuid,
                int(datetime.now().timestamp())
            ),
            "sub_category_id": sub_category_id,
            "hide_from_game_eight": hide_from_game_eight,
            "allow_members_to_post_image_and_video": allow_members_to_post_media,
            "allow_members_to_post_url": allow_members_to_post_url,
            "guidelines": guidelines,
        }, data_type=CreateGroupResponse
    )
    logger(self, fname="create_group")
    return response


def create_pin_group(self, group_id: int):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.PINNED_V1}/groups",
        payload={"id": group_id}
    )
    logger(self, fname="create_pin_group")
    return response


def decline_moderator_offer(self, group_id: int):
    self._check_authorization()
    response = self._make_request(
        "DELETE", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/deputize"
    )
    logger(self, fname="decline_moderator_offer")
    return response


def decline_ownership_offer(self, group_id: int):
    self._check_authorization()
    response = self._make_request(
        "DELETE", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/transfer"
    )
    logger(self, fname="decline_ownership_offer")
    return response


def decline_group_join_request(self, group_id: int, user_id: int):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/decline/{user_id}"
    )
    logger(self, fname="decline_group_join_request")
    return response


def delete_pin_group(self, group_id: int):
    self._check_authorization()
    response = self._make_request(
        "DELETE", endpoint=f"{Endpoints.PINNED_V1}/groups/{group_id}"
    )
    logger(self, fname="delete_pin_group")
    return response


def get_banned_group_members(self, group_id: int, page: int = None) -> UsersResponse:
    self._check_authorization()
    params = {}
    if page:
        params["page"] = page
    return self._make_request(
        "GET", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/ban_list",
        params=params
    )


def get_group_categories(self, **params) -> GroupCategoriesResponse:
    """

    Parameters:
    ----------

        - page: int - (optional)
        - number: int - (optional)

    """
    return self._make_request(
        "GET", endpoint=f"{Endpoints.GROUPS_V1}/categories",
        params=params, data_type=GroupCategoriesResponse
    )


def get_create_group_quota(self) -> CreateGroupQuota:
    return self._make_request(
        "GET", endpoint=f"{Endpoints.GROUPS_V1}/created_quota",
        data_type=CreateGroupQuota
    ).create


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
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/invite",
        payload={"user_ids[]": user_ids}
    )
    logger(self, fname="invite_users_to_group")
    return response


def join_group(self, group_id: int):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/join",
    )
    logger(self, fname="join_group")
    return response


def leave_group(self, group_id: int):
    self._check_authorization()
    response = self._make_request(
        "DELETE", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/leave",
    )
    logger(self, fname="leave_group")
    return response


def post_gruop_social_shared(self, group_id: int, sns_name: str):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.GROUPS_V2}/{group_id}/social_shared",
        params={"sns_name": sns_name}
    )
    logger(self, fname="post_gruop_social_shared")
    return response


def remove_group_cover(self, group_id: int):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/remove_cover",
    )
    logger(self, fname="remove_group_cover")
    return response


def remove_moderator(self, group_id: int, user_id: int):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/fire/{user_id}",
    )
    logger(self, fname="remove_moderator")
    return response


def remove_related_groups(self, group_id: int, related_group_ids: List[int]):
    self._check_authorization()
    response = self._make_request(
        "DELETE", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/related",
        params={"related_group_id[]": related_group_ids}
    )
    logger(self, fname="remove_related_groups")
    return response


def report_group(
        self,
        group_id: int,
        category_id: int,
        reason: str = None,
        opponent_id: int = None,
        screenshot_filename: str = None,
        screenshot_2_filename: str = None,
        screenshot_3_filename: str = None,
        screenshot_4_filename: str = None,
):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.GROUPS_V3}/{group_id}/report",
        payload={
            "category_id": category_id,
            "reason": reason,
            "opponent_id": opponent_id,
            "screenshot_filename": screenshot_filename,
            "screenshot_2_filename": screenshot_2_filename,
            "screenshot_3_filename": screenshot_3_filename,
            "screenshot_4_filename": screenshot_4_filename
        }
    )
    logger(self, fname="report_group")
    return response


def send_moderator_offers(self, group_id: int, user_ids: List[int]):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.GROUPS_V3}/{group_id}/deputize/mass",
        payload={
            "user_ids[]": user_ids, "uuid": self.uuid,
            "api_key": self.api_key, "timestamp": int(datetime.now().timestamp()),
            "signed_info": signed_info_calculating(
                self.api_key, self.device_uuid,
                int(datetime.now().timestamp())
            ),
        }
    )
    logger(self, fname="send_moderator_offers")
    return response


def send_ownership_offer(self, group_id: int, user_id: int):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.GROUPS_V3}/{group_id}/transfer",
        payload={
            "user_id": user_id, "uuid": self.uuid,
            "api_key": self.api_key, "timestamp": int(datetime.now().timestamp()),
            "signed_info": signed_info_calculating(
                self.api_key, self.device_uuid,
                int(datetime.now().timestamp())
            ),
        }
    )
    logger(self, fname="send_ownership_offer")
    return response


def set_group_notification_settings(
        self,
        group_id: int,
        notification_group_post: int = None,
        notification_group_join: int = None,
        notification_group_request: int = None,
        notification_group_message_tag_all: int = None,
) -> AdditionalSettingsResponse:
    pass


def set_group_title(self, group_id: int, title: str):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/set_title",
        payload={"title": title}
    )
    logger(self, fname="set_group_title")
    return response


def take_over_group_ownership(self, group_id: int):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/take_over",
    )
    logger(self, fname="take_over_group_ownership", group_id=group_id)
    return response


def unban_group_member(self, group_id: int, user_id: int):
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/unban/{user_id}",
    )
    logger(self, fname="unban_group_member")
    return response


def update_group(
        self,
        group_id: int,
        topic: str,
        description: str = None,
        secret: bool = None,
        hide_reported_posts: bool = None,
        hide_conference_call: bool = None,
        is_private: bool = None,
        only_verified_age: bool = None,
        only_mobile_verified: bool = None,
        call_timeline_display: bool = None,
        allow_ownership_transfer: bool = None,
        allow_thread_creation_by: str = None,
        gender: int = None,
        generation_groups_limit: int = None,
        group_category_id: int = None,
        cover_image_filename: str = None,
        sub_category_id: str = None,
        hide_from_game_eight: bool = None,
        allow_members_to_post_media: bool = None,
        allow_members_to_post_url: bool = None,
        guidelines: str = None,
) -> GroupResponse:
    self._check_authorization()
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.GROUPS_V3}/{group_id}/update",
        payload={
            "topic": topic,
            "description": description,
            "secret": secret,
            "hide_reported_posts": hide_reported_posts,
            "hide_conference_call": hide_conference_call,
            "is_private": is_private,
            "only_verified_age": only_verified_age,
            "only_mobile_verified": only_mobile_verified,
            "call_timeline_display": call_timeline_display,
            "allow_ownership_transfer": allow_ownership_transfer,
            "allow_thread_creation_by": allow_thread_creation_by,
            "gender": gender,
            "generation_groups_limit": generation_groups_limit,
            "group_category_id": group_category_id,
            "cover_image_filename": cover_image_filename,
            "sub_category_id": sub_category_id,
            "uuid": self.uuid,
            "api_key": self.api_key,
            "timestamp": int(datetime.now().timestamp()),
            "signed_info": signed_info_calculating(
                self.api_key, self.device_uuid,
                int(datetime.now().timestamp())
            ),
            "hide_from_game_eight": hide_from_game_eight,
            "allow_members_to_post_image_and_video": allow_members_to_post_media,
            "allow_members_to_post_url": allow_members_to_post_url,
            "guidelines": guidelines,
        }, data_type=GroupResponse
    )
    logger(self, fname="update_group")
    return response


def visit_group(self, group_id: int):
    response = self._make_request(
        "POST", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/visit",
    )
    logger(self, fname="visit_group")
    return response


def withdraw_moderator_offer(self, group_id: int, user_id: int):
    self._check_authorization()
    response = self._make_request(
        "PUT", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/deputize/{user_id}/withdraw",
    )
    logger(self, fname="withdraw_moderator_offer")
    return response


def withdraw_ownership_offer(self, group_id: int, user_id: int):
    self._check_authorization()
    response = self._make_request(
        "PUT", endpoint=f"{Endpoints.GROUPS_V1}/{group_id}/transfer/withdraw",
        payload={"user_id": user_id}
    )
    logger(self, fname="withdraw_ownership_offer")
    return response
