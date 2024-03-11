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
from ..models import CreateGroupQuota
from ..responses import (
    CreateGroupResponse,
    GroupCategoriesResponse,
    GroupResponse,
    GroupsRelatedResponse,
    GroupsResponse,
    GroupUserResponse,
    GroupUsersResponse,
    UnreadStatusResponse,
    UsersResponse,
    UsersByTimestampResponse,
)
from ..utils import md5


class GroupAPI(object):
    def __init__(self, base: client.BaseClient) -> None:
        self.__base = base

    def accept_moderator_offer(self, group_id: int):
        return self.__base._request("PUT", route=f"/v1/groups/{group_id}/deputize")

    def accept_ownership_offer(
        self,
        group_id: int,
    ):
        return self.__base._request("PUT", route=f"/v1/groups/{group_id}/transfer")

    def accept_group_join_request(
        self,
        group_id: int,
        user_id: int,
    ):
        return self.__base._request(
            "POST",
            route=f"/v1/groups/{group_id}/accept/{user_id}",
        )

    def add_related_groups(self, group_id: int, related_group_id: list[int]):
        """

        関連サークルを追加する

        """
        return self.__base._request(
            "PUT",
            route=f"/v1/groups/{group_id}/related",
            params={"related_group_id": related_group_id},
        )

    def ban_group_user(self, group_id: int, user_id: int):
        return self.__base._request(
            "POST",
            route=f"/v1/groups/{group_id}/ban/{user_id}",
        )

    def check_unread_status(self, from_time: int = None) -> UnreadStatusResponse:
        params = {}
        if from_time:
            params["from_time"] = from_time
        return self.__base._request(
            "GET",
            route=f"/v1/groups/unread_status",
            params=params,
            data_type=UnreadStatusResponse,
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
        return self.__base._request(
            "POST",
            route=f"/v3/groups/new",
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
                "uuid": self.__base.uuid,
                "api_key": Configs.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": self.__signed_info,
                "sub_category_id": sub_category_id,
                "hide_from_game_eight": hide_from_game_eight,
                "allow_members_to_post_image_and_video": allow_members_to_post_media,
                "allow_members_to_post_url": allow_members_to_post_url,
                "guidelines": guidelines,
            },
            data_type=CreateGroupResponse,
        )

    def create_pin_group(self, group_id: int):
        return self.__base._request(
            "POST", route=f"/v1/pinned/groups", payload={"id": group_id}
        )

    def decline_moderator_offer(self, group_id: int):
        return self.__base._request("DELETE", route=f"/v1/groups/{group_id}/deputize")

    def decline_ownership_offer(self, group_id: int):
        return self.__base._request("DELETE", route=f"/v1/groups/{group_id}/transfer")

    def decline_group_join_request(self, group_id: int, user_id: int):
        return self.__base._request(
            "POST", route=f"/v1/groups/{group_id}/decline/{user_id}"
        )

    def delete_pin_group(self, group_id: int):
        return self.__base._request("DELETE", route=f"/v1/pinned/groups/{group_id}")

    def get_banned_group_members(
        self, group_id: int, page: int = None
    ) -> UsersResponse:
        params = {}
        if page:
            params["page"] = page
        return self.__base._request(
            "GET", route=f"/v1/groups/{group_id}/ban_list", params=params
        )

    def get_group_categories(self, **params) -> GroupCategoriesResponse:
        """

        Parameters:
        -----------

            - page: int - (optional)
            - number: int - (optional)

        """
        return self.__base._request(
            "GET",
            route=f"/v1/groups/categories",
            params=params,
            data_type=GroupCategoriesResponse,
        )

    def get_create_group_quota(self) -> CreateGroupQuota:
        return self.__base._request(
            "GET",
            route=f"/v1/groups/created_quota",
            data_type=CreateGroupQuota,
        )

    def get_group(self, group_id: int) -> GroupResponse:
        return self.__base._request(
            "GET",
            route=f"/v1/groups/{group_id}",
            data_type=GroupResponse,
        )

    def get_groups(self, **params) -> GroupsResponse:
        """

        Parameters:
        -----------

            - group_category_id: int = None
            - keyword: str = None
            - from_timestamp: int = None
            - sub_category_id: int = None

        """
        return self.__base._request(
            "GET",
            route=f"/v2/groups",
            params=params,
            data_type=GroupsResponse,
        )

    def get_invitable_users(self, group_id: int, **params) -> UsersByTimestampResponse:
        """

        Parameters:
        -----------

            - from_timestamp: int - (optional)
            - user[nickname]: str - (optional)

        """
        return self.__base._request(
            "GET",
            route=f"/v1/groups/{group_id}/users/invitable",
            params=params,
            data_type=UsersByTimestampResponse,
        )

    def get_joined_statuses(self, ids: list[int]) -> dict:
        return self.__base._request(
            "GET",
            route=f"/v1/groups/joined_statuses",
            params={"ids": ids},
        )

    def get_group_member(self, group_id: int, user_id: int) -> GroupUserResponse:
        return self.__base._request(
            "GET", route=f"/v1/groups/{group_id}/members/{user_id}"
        )

    def get_group_members(self, group_id: int, **params) -> GroupUsersResponse:
        """

        Parameters:
        -----------

            - id: int - (required)
            - mode: str - (optional)
            - keyword: str - (optional)
            - from_id: int - (optional)
            - from_timestamp: int - (optional)
            - order_by: str - (optional)
            - followed_by_me: bool - (optional)

        """
        return self.__base._request(
            "GET",
            route=f"/v2/groups/{group_id}/members",
            params=params,
            data_type=GroupUsersResponse,
        )

    def get_my_groups(self, from_timestamp: None) -> GroupsResponse:
        params = {}
        if from_timestamp:
            params["from_timestamp"] = from_timestamp
        return self.__base._request(
            "GET",
            route=f"/v2/groups/mine",
            params=params,
            data_type=GroupsResponse,
        )

    def get_relatable_groups(self, group_id: int, **params) -> GroupsRelatedResponse:
        """

        Parameters:
        -----------

            - group_id: int - (required)
            - keyword: str - (optional)
            - from: str - (optional)

        """
        return self.__base._request(
            "GET",
            route=f"/v1/groups/{group_id}/relatable",
            params=params,
            data_type=GroupsRelatedResponse,
        )

    def get_related_groups(self, group_id: int, **params) -> GroupsRelatedResponse:
        """

        Parameters:
        -----------

            - group_id: int - (required)
            - keyword: str - (optional)
            - from: str - (optional)

        """
        return self.__base._request(
            "GET",
            route=f"/v1/groups/{group_id}/related",
            params=params,
            data_type=GroupsRelatedResponse,
        )

    def get_user_groups(self, **params) -> GroupsResponse:
        """

        Parameters:
        -----------

            - user_id: int - (required)
            - page: int - (optional)

        """
        return self.__base._request(
            "GET",
            route=f"/v1/groups/user_group_list",
            params=params,
            data_type=GroupsResponse,
        )

    def invite_users_to_group(
        self,
        group_id: int,
        user_ids: list[int],
    ):
        return self.__base._request(
            "POST",
            route=f"/v1/groups/{group_id}/invite",
            payload={"user_ids": user_ids},
        )

    def join_group(
        self,
        group_id: int,
    ):
        return self.__base._request(
            "POST",
            route=f"/v1/groups/{group_id}/join",
        )

    def leave_group(
        self,
        group_id: int,
    ):
        return self.__base._request(
            "DELETE",
            route=f"/v1/groups/{group_id}/leave",
        )

    def post_gruop_social_shared(
        self,
        group_id: int,
        sns_name: str,
    ):
        return self.__base._request(
            "POST",
            route=f"/v2/groups/{group_id}/social_shared",
            params={"sns_name": sns_name},
        )

    def remove_group_cover(
        self,
        group_id: int,
    ):
        return self.__base._request(
            "POST",
            route=f"/v1/groups/{group_id}/remove_cover",
        )

    def remove_moderator(
        self,
        group_id: int,
        user_id: int,
    ):
        return self.__base._request(
            "POST",
            route=f"/v1/groups/{group_id}/fire/{user_id}",
        )

    def remove_related_groups(
        self,
        group_id: int,
        related_group_ids: list[int],
    ):
        return self.__base._request(
            "DELETE",
            route=f"/v1/groups/{group_id}/related",
            params={"related_group_id[]": related_group_ids},
        )

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
        return self.__base._request(
            "POST",
            route=f"/v3/groups/{group_id}/report",
            payload={
                "category_id": category_id,
                "reason": reason,
                "opponent_id": opponent_id,
                "screenshot_filename": screenshot_filename,
                "screenshot_2_filename": screenshot_2_filename,
                "screenshot_3_filename": screenshot_3_filename,
                "screenshot_4_filename": screenshot_4_filename,
            },
        )

    def send_moderator_offers(self, group_id: int, user_ids: list[int]):
        return self.__base._request(
            "POST",
            route=f"/v3/groups/{group_id}/deputize/mass",
            payload={
                "user_ids": user_ids,
                "uuid": self.__base.uuid,
                "api_key": Configs.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": self.__signed_info,
            },
        )

    def send_ownership_offer(self, group_id: int, user_id: int):
        return self.__base._request(
            "POST",
            route=f"/v3/groups/{group_id}/transfer",
            payload={
                "user_id": user_id,
                "uuid": self.__base.uuid,
                "api_key": Configs.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": self.__signed_info,
            },
        )

    def set_group_title(self, group_id: int, title: str):
        return self.__base._request(
            "POST",
            route=f"/v1/groups/{group_id}/set_title",
            payload={"title": title},
        )

    def take_over_group_ownership(self, group_id: int):
        return self.__base._request("POST", route=f"/v1/groups/{group_id}/take_over")

    def unban_group_member(self, group_id: int, user_id: int):
        return self.__base._request(
            "POST", route=f"/v1/groups/{group_id}/unban/{user_id}"
        )

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
        return self.__base._request(
            "POST",
            route=f"/v3/groups/{group_id}/update",
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
                "uuid": self.__base.uuid,
                "api_key": Configs.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": self.__signed_info,
                "hide_from_game_eight": hide_from_game_eight,
                "allow_members_to_post_image_and_video": allow_members_to_post_media,
                "allow_members_to_post_url": allow_members_to_post_url,
                "guidelines": guidelines,
            },
            data_type=GroupResponse,
        )

    def withdraw_moderator_offer(self, group_id: int, user_id: int):
        return self.__base._request(
            "PUT",
            route=f"/v1/groups/{group_id}/deputize/{user_id}/withdraw",
        )

    def withdraw_ownership_offer(self, group_id: int, user_id: int):
        return self.__base._request(
            "PUT",
            route=f"/v1/groups/{group_id}/transfer/withdraw",
            payload={"user_id": user_id},
        )

    @property
    def __signed_info(self) -> str:
        return md5(self.__base.uuid, int(datetime.now().timestamp()), True)
