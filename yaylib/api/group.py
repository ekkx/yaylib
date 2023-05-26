from datetime import datetime
from typing import Union, List

from ..config import *
from ..errors import *
from ..models import *
from ..responses import *
from ..utils import *


def accept_moderator_offer(self, group_id: int):
    pass


def accept_ownership_offer(self, group_id: int):
    pass


def accept_group_join_request(self, group_id: int, user_id: int):
    pass


def add_related_groups(self, group_id: int, related_group_id: List[int]):
    pass


def ban_group_user(self, group_id: int, user_id: int):
    pass


def check_unread_status(self, from_time: int = None) -> UnreadStatusResponse:
    pass


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
        allow_members_to_post_image_and_video: bool = None,
        allow_members_to_post_url: bool = None,
        guidelines: str = None,
) -> CreateGroupResponse:
    pass


def create_pin_group(self, group_id: int):
    pass


def decline_moderator_offer(self, group_id: int):
    pass


def decline_ownership_offer(self, group_id: int):
    pass


def decline_user_request(self, group_id: int, user_id: int):
    pass


def delete_pin_group(self, group_id: int):
    pass


def get_banned_group_members(
        self,
        group_id: int,
        page: int = None
) -> UsersResponse:
    pass


def get_categories(self, **kwargs: int) -> GroupCategoriesResponse:
    pass


def get_create_group_quota(self) -> CreateGroupQuota:
    # CreateGroupQuota.create
    pass


def get_group(self, group_id: int) -> GroupResponse:
    pass


def get_group_notification_settings(self, group_id: int) -> GroupNotificationSettingsResponse:
    pass


def get_groups(self, **kwargs: Union[int, str]) -> GroupsResponse:
    pass


def get_invitable_users(
        self,
        group_id: int,
        **kwargs: Union[int, str]
) -> UsersByTimestampResponse:
    pass


def get_joined_statuses(self, ids: List[int]) -> dict:
    pass


def get_group_members(
        self,
        group_id: int,
        **kwargs: Union[str, int, bool]
) -> GroupUsersResponse:
    pass


def get_my_groups(self, from_timestamp: None) -> GroupsResponse:
    pass


def get_relatable_groups(self, group_id: int, **kwargs: str) -> GroupsRelatedResponse:
    pass


def get_related_groups(self, group_id: int, **kwargs: str) -> GroupsRelatedResponse:
    pass


def get_user_groups(self, user_id: int, page: int = None) -> GroupsResponse:
    pass


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
    pass


def send_moderator_offers(self, group_id: int, user_ids: List[int]):
    pass


def send_ownership_offer(self, group_id: int, user_id: int):
    pass


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
    pass


def take_over_group_ownership(self, group_id: int):
    pass


def unban_group_member(self, group_id: int, user_id: int):
    pass


def update_group(
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
        allow_members_to_post_image_and_video: bool = None,
        allow_members_to_post_url: bool = None,
        guidelines: str = None,
) -> GroupResponse:
    pass


def visit_group(self, group_id: int):
    pass


def withdraw_moderator_offer(self, group_id: int, user_id: int):
    pass


def withdraw_ownership_offer(self, group_id: int, user_id: int):
    pass
