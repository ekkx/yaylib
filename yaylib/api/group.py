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

from datetime import datetime
from typing import List

from .. import config
from ..models import CreateGroupQuota
from ..responses import (
    CreateGroupResponse,
    GroupCategoriesResponse,
    GroupResponse,
    GroupsRelatedResponse,
    GroupsResponse,
    GroupUserResponse,
    GroupUsersResponse,
    Response,
    UnreadStatusResponse,
    UsersByTimestampResponse,
    UsersResponse,
)
from ..utils import md5


# pylint: disable=too-many-public-methods
class GroupApi:
    """サークル API"""

    def __init__(self, client) -> None:
        # pylint: disable=import-outside-toplevel
        from ..client import Client

        self.__client: Client = client

    async def accept_moderator_offer(self, group_id: int) -> Response:
        """サークル副管理人の権限オファーを引き受ける

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "PUT",
            config.API_HOST + f"/v1/groups/{group_id}/deputize",
            return_type=Response,
        )

    async def accept_ownership_offer(
        self,
        group_id: int,
    ) -> Response:
        """サークル管理人の権限オファーを引き受けます

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "PUT",
            config.API_HOST + f"/v1/groups/{group_id}/transfer",
            return_type=Response,
        )

    async def accept_group_join_request(
        self,
        group_id: int,
        user_id: int,
    ) -> Response:
        """サークル参加リクエストを承認します

        Args:
            group_id (int):
            user_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/groups/{group_id}/accept/{user_id}",
            return_type=Response,
        )

    async def add_related_groups(
        self, group_id: int, related_group_id: List[int]
    ) -> Response:
        """関連サークルを追加する

        Args:
            group_id (int):
            related_group_id (List[int]):

        Returns:
            Response:
        """
        return await self.__client.request(
            "PUT",
            config.API_HOST + f"/v1/groups/{group_id}/related",
            params={"related_group_id": related_group_id},
            return_type=Response,
        )

    async def ban_group_user(self, group_id: int, user_id: int) -> Response:
        """サークルからユーザーを追放する

        Args:
            group_id (int):
            user_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/groups/{group_id}/ban/{user_id}",
            return_type=Response,
        )

    async def check_group_unread_status(self, **params) -> UnreadStatusResponse:
        """サークルの未読ステータスを取得する

        Args:
            from_time (int, optional):

        Returns:
            UnreadStatusResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/groups/unread_status",
            params=params,
            return_type=UnreadStatusResponse,
        )

    async def create_group(self, **params) -> CreateGroupResponse:
        """サークルを作成する

        Args:
            topic (str):
            description (str, optional):
            secret (bool, optional):
            hide_reported_posts (bool, optional):
            hide_conference_call (bool, optional):
            is_private (bool, optional):
            only_verified_age (bool, optional):
            only_mobile_verified (bool, optional):
            call_timeline_display (bool, optional):
            allow_ownership_transfer (bool, optional):
            allow_thread_creation_by (str, optional):
            gender (int, optional):
            generation_groups_limit (int, optional):
            group_category_id (int, optional):
            cover_image_filename (str, optional):
            sub_category_id (str, optional):
            hide_from_game_eight (bool, optional):
            allow_members_to_post_media (bool, optional):
            allow_members_to_post_url (bool, optional):
            guidelines (str, optional):

        Returns:
            CreateGroupResponse:
        """
        params.update(
            {
                "uuid": self.__client.device_uuid,
                "api_key": config.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__client.device_uuid, int(datetime.now().timestamp()), True
                ),
            }
        )
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v3/groups/new",
            json=params,
            return_type=CreateGroupResponse,
        )

    async def pin_group(self, group_id: int) -> Response:
        """サークルをピン留めする

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/pinned/groups",
            json={"id": group_id},
            return_type=Response,
        )

    async def decline_moderator_offer(self, group_id: int) -> Response:
        """サークル副管理人の権限オファーを断る

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v1/groups/{group_id}/deputize",
            return_type=Response,
        )

    async def decline_ownership_offer(self, group_id: int) -> Response:
        """サークル管理人の権限オファーを断る

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v1/groups/{group_id}/transfer",
            return_type=Response,
        )

    async def decline_group_join_request(self, group_id: int, user_id: int) -> Response:
        """サークル参加リクエストを断る

        Args:
            group_id (int):
            user_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/groups/{group_id}/decline/{user_id}",
            return_type=Response,
        )

    async def unpin_group(self, group_id: int) -> Response:
        """サークルのピン留めを解除する

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v1/pinned/groups/{group_id}",
            return_type=Response,
        )

    async def get_banned_group_members(self, group_id: int, **params) -> UsersResponse:
        """追放されたサークルメンバーを取得する

        Args:
            group_id (int):
            page (int, optional):

        Returns:
            UsersResponse:
        """
        return await self.__client.request(
            "GET", config.API_HOST + f"/v1/groups/{group_id}/ban_list", params=params
        )

    async def get_group_categories(self, **params) -> GroupCategoriesResponse:
        """サークルのカテゴリーを取得する

        Args:
            page (int, optional):
            number (int, optional):

        Returns:
            GroupCategoriesResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/groups/categories",
            params=params,
            return_type=GroupCategoriesResponse,
        )

    async def get_create_group_quota(self) -> CreateGroupQuota:
        """残りのサークル作成可能回数を取得する

        Returns:
            CreateGroupQuota:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/groups/created_quota",
            return_type=CreateGroupQuota,
        )

    async def get_group(self, group_id: int) -> GroupResponse:
        """サークルの詳細を取得する

        Args:
            group_id (int):

        Returns:
            GroupResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v1/groups/{group_id}",
            return_type=GroupResponse,
        )

    async def get_groups(self, **params) -> GroupsResponse:
        """複数のサークル情報を取得する

        Args:
            group_category_id (int, optional):
            keyword (str, optional):
            from_timestamp (int, optional):
            sub_category_id (int, optional):

        Returns:
            GroupsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/groups",
            params=params,
            return_type=GroupsResponse,
        )

    async def get_invitable_users(
        self, group_id: int, **params
    ) -> UsersByTimestampResponse:
        """サークルに招待可能なユーザーを取得する

        Args:
            group_id (int):
            from_timestamp (int, optional):
            user[nickname] (str, optional):

        Returns:
            UsersByTimestampResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v1/groups/{group_id}/users/invitable",
            params=params,
            return_type=UsersByTimestampResponse,
        )

    async def get_joined_statuses(self, ids: List[int]) -> Response:
        """サークルの参加ステータスを取得する

        Args:
            ids (List[int]):

        Returns:
            Response:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/groups/joined_statuses",
            params={"ids": ids},
            return_type=Response,
        )

    async def get_group_member(self, group_id: int, user_id: int) -> GroupUserResponse:
        """特定のサークルメンバーの情報を取得する

        Args:
            group_id (int):
            user_id (int):

        Returns:
            GroupUserResponse:
        """
        return await self.__client.request(
            "GET", config.API_HOST + f"/v1/groups/{group_id}/members/{user_id}"
        )

    async def get_group_members(self, group_id: int, **params) -> GroupUsersResponse:
        """サークルメンバーを取得する

        Args:
            group_id (int):
            id (int):
            mode (str, optional):
            keyword (str, optional):
            from_id (int, optional):
            from_timestamp (int, optional):
            order_by (str, optional):
            followed_by_me: (bool, optional)

        Returns:
            GroupUsersResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v2/groups/{group_id}/members",
            params=params,
            return_type=GroupUsersResponse,
        )

    async def get_my_groups(self, **params) -> GroupsResponse:
        """自分のサークルを取得する

        Args:
            from_timestamp (_type_, optional):

        Returns:
            GroupsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v2/groups/mine",
            params=params,
            return_type=GroupsResponse,
        )

    async def get_relatable_groups(
        self, group_id: int, **params
    ) -> GroupsRelatedResponse:
        """関連がある可能性があるサークルを取得する

        Args:
            group_id (int):
            keyword (str, optional):
            from (str, optional):

        Returns:
            GroupsRelatedResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v1/groups/{group_id}/relatable",
            params=params,
            return_type=GroupsRelatedResponse,
        )

    async def get_related_groups(
        self, group_id: int, **params
    ) -> GroupsRelatedResponse:
        """関連があるサークルを取得する

        Args:
            group_id (int):
            keyword (str, optional):
            from (str, optional):

        Returns:
            GroupsRelatedResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v1/groups/{group_id}/related",
            params=params,
            return_type=GroupsRelatedResponse,
        )

    async def get_user_groups(self, **params) -> GroupsResponse:
        """特定のユーザーが参加しているサークルを取得する

        Args:
            user_id (int):
            page (int, optional):

        Returns:
            GroupsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/groups/user_group_list",
            params=params,
            return_type=GroupsResponse,
        )

    async def invite_users_to_group(
        self,
        group_id: int,
        user_ids: List[int],
    ) -> Response:
        """サークルにユーザーを招待する

        Args:
            group_id (int):
            user_ids (List[int]):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/groups/{group_id}/invite",
            json={"user_ids": user_ids},
            return_type=Response,
        )

    async def join_group(
        self,
        group_id: int,
    ) -> Response:
        """サークルに参加する

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/groups/{group_id}/join",
            return_type=Response,
        )

    async def leave_group(
        self,
        group_id: int,
    ) -> Response:
        """サークルから脱退する

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v1/groups/{group_id}/leave",
            return_type=Response,
        )

    async def delete_group_cover(
        self,
        group_id: int,
    ) -> Response:
        """サークルのカバー画像を削除する

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/groups/{group_id}/remove_cover",
            return_type=Response,
        )

    async def delete_moderator(
        self,
        group_id: int,
        user_id: int,
    ) -> Response:
        """サークルの副管理人を削除する

        Args:
            group_id (int):
            user_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/groups/{group_id}/fire/{user_id}",
            return_type=Response,
        )

    async def delete_related_groups(
        self,
        group_id: int,
        related_group_ids: List[int],
    ) -> Response:
        """関連のあるサークルを削除する

        Args:
            group_id (int):
            related_group_ids (List[int]):

        Returns:
            Response:
        """
        return await self.__client.request(
            "DELETE",
            config.API_HOST + f"/v1/groups/{group_id}/related",
            params={"related_group_id[]": related_group_ids},
            return_type=Response,
        )

    async def send_moderator_offers(
        self, group_id: int, user_ids: List[int]
    ) -> Response:
        """複数人にサークル副管理人のオファーを送信する

        Args:
            group_id (int):
            user_ids (List[int]):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v3/groups/{group_id}/deputize/mass",
            json={
                "user_ids": user_ids,
                "uuid": self.__client.device_uuid,
                "api_key": config.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__client.device_uuid, int(datetime.now().timestamp()), True
                ),
            },
            return_type=Response,
        )

    async def send_ownership_offer(self, group_id: int, user_id: int) -> Response:
        """サークル管理人権限のオファーを送信する

        Args:
            group_id (int):
            user_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v3/groups/{group_id}/transfer",
            json={
                "user_id": user_id,
                "uuid": self.__client.device_uuid,
                "api_key": config.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__client.device_uuid, int(datetime.now().timestamp()), True
                ),
            },
            return_type=Response,
        )

    async def set_group_title(self, group_id: int, title: str) -> Response:
        """サークルのタイトルを設定する

        Args:
            group_id (int):
            title (str):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/groups/{group_id}/set_title",
            json={"title": title},
            return_type=Response,
        )

    async def take_over_group_ownership(self, group_id: int) -> Response:
        """サークル管理人の権限を引き継ぐ

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/groups/{group_id}/take_over",
            return_type=Response,
        )

    async def unban_group_member(self, group_id: int, user_id: int) -> Response:
        """特定のサークルメンバーの追放を解除する

        Args:
            group_id (int):
            user_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/groups/{group_id}/unban/{user_id}",
            return_type=Response,
        )

    async def update_group(self, group_id: int, **params) -> GroupResponse:
        """サークルを編集する

        Args:
            group_id (int):
            topic (str):
            description (str, optional):
            secret (bool, optional):
            hide_reported_posts (bool, optional):
            hide_conference_call (bool, optional):
            is_private (bool, optional):
            only_verified_age (bool, optional):
            only_mobile_verified (bool, optional):
            call_timeline_display (bool, optional):
            allow_ownership_transfer (bool, optional):
            allow_thread_creation_by (str, optional):
            gender (int, optional):
            generation_groups_limit (int, optional):
            group_category_id (int, optional):
            cover_image_filename (str, optional):
            sub_category_id (str, optional):
            hide_from_game_eight (bool, optional):
            allow_members_to_post_media (bool, optional):
            allow_members_to_post_url (bool, optional):
            guidelines (str, optional):

        Returns:
            GroupResponse:
        """
        params.update(
            {
                "uuid": self.__client.device_uuid,
                "api_key": config.API_KEY,
                "timestamp": int(datetime.now().timestamp()),
                "signed_info": md5(
                    self.__client.device_uuid, int(datetime.now().timestamp()), True
                ),
            }
        )
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v3/groups/{group_id}/update",
            json=params,
            return_type=GroupResponse,
        )

    async def withdraw_moderator_offer(self, group_id: int, user_id: int) -> Response:
        """サークル副管理人のオファーを取り消す

        Args:
            group_id (int):
            user_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "PUT",
            config.API_HOST + f"/v1/groups/{group_id}/deputize/{user_id}/withdraw",
            return_type=Response,
        )

    async def withdraw_ownership_offer(self, group_id: int, user_id: int) -> Response:
        """サークル管理人のオファーを取り消す

        Args:
            group_id (int):
            user_id (int):

        Returns:
            Response:
        """
        return await self.__client.request(
            "PUT",
            config.API_HOST + f"/v1/groups/{group_id}/transfer/withdraw",
            json={"user_id": user_id},
            return_type=Response,
        )
