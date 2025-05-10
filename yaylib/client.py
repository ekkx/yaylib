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

import asyncio
import logging
import os
import random
from datetime import datetime
from typing import Dict, List, Optional

import aiohttp

from . import __version__
from .api.auth import AuthApi
from .api.call import CallApi
from .api.chat import ChatApi
from .api.group import GroupApi
from .api.misc import MiscApi
from .api.notification import NotificationApi
from .api.post import PostApi
from .api.review import ReviewApi
from .api.thread import ThreadApi
from .api.user import UserApi
from .config import API_HOST, API_VERSION_NAME
from .device import Device
from .errors import (
    AccessTokenExpiredError,
    AccessTokenInvalidError,
    HTTPInternalServerError,
    QuotaLimitExceededError,
    TooManyRequestsError,
    UnauthorizedError,
    raise_for_code,
    raise_for_status,
)
from .models import Attachment, CreateGroupQuota, Model, Post, SharedUrl, ThreadInfo
from .responses import (
    ActiveFollowingsResponse,
    ActivitiesResponse,
    ApplicationConfigResponse,
    BanWordsResponse,
    BgmsResponse,
    BlockedUserIdsResponse,
    BlockedUsersResponse,
    BookmarkPostResponse,
    CallActionSignatureResponse,
    CallStatusResponse,
    ChatRoomResponse,
    ChatRoomsResponse,
    ConferenceCallResponse,
    CreateChatRoomResponse,
    CreateGroupResponse,
    CreatePostResponse,
    CreateUserResponse,
    EmailGrantTokenResponse,
    EmailVerificationPresignedUrlResponse,
    FollowRecommendationsResponse,
    FollowRequestCountResponse,
    FollowUsersResponse,
    FootprintsResponse,
    GamesResponse,
    GenresResponse,
    GifsDataResponse,
    GroupCategoriesResponse,
    GroupResponse,
    GroupsRelatedResponse,
    GroupsResponse,
    GroupThreadListResponse,
    GroupUserResponse,
    GroupUsersResponse,
    HiddenResponse,
    HimaUsersResponse,
    IdCheckerPresignedUrlResponse,
    LikePostsResponse,
    LoginUpdateResponse,
    LoginUserResponse,
    MessageResponse,
    MessagesResponse,
    PolicyAgreementsResponse,
    PopularWordsResponse,
    PostLikersResponse,
    PostResponse,
    PostsResponse,
    PostTagsResponse,
    PresignedUrlResponse,
    PresignedUrlsResponse,
    RankingUsersResponse,
    RefreshCounterRequestsResponse,
    Response,
    ReviewsResponse,
    SocialShareUsersResponse,
    StickerPacksResponse,
    TokenResponse,
    TotalChatRequestResponse,
    UnreadStatusResponse,
    UserResponse,
    UsersByTimestampResponse,
    UsersResponse,
    UserTimestampResponse,
    VoteSurveyResponse,
    WebSocketTokenResponse,
)
from .state import LocalUser, State
from .utils import CustomFormatter, filter_dict, generate_jwt
from .ws import Intents, WebSocketInteractor

__all__ = ["Client"]


class RateLimit:
    """レート制限を管理するクラス"""

    def __init__(
        self, wait_on_ratelimit: bool, max_retries: int, retry_after=60 * 5
    ) -> None:
        self.__wait_on_ratelimit = wait_on_ratelimit
        self.__max_retries = max_retries
        self.__retries_performed = 0
        self.__retry_after = retry_after

    @property
    def retries_performed(self) -> int:
        """レート制限によるリトライ回数

        Returns:
            int: リトライ回数
        """
        return self.__retries_performed

    @retries_performed.setter
    def retries_performed(self, value: int) -> None:
        self.__retries_performed = min(value, self.__max_retries)

    @property
    def max_retries(self) -> int:
        """レート制限によるリトライ回数の上限

        Returns:
            int: リトライ回数の上限
        """
        return self.__max_retries

    @property
    def max_retries_reached(self) -> bool:
        """リトライ回数上限に達したか否か"""
        return not self.__wait_on_ratelimit or (
            self.__retries_performed >= self.__max_retries
        )

    def reset(self) -> None:
        """レート制限をリセットする"""
        self.__retries_performed = 0

    async def wait(self, err: Exception) -> None:
        """レート制限が解除されるまで待機する

        Raises:
            Exception: リトライ回数の上限でスロー
        """
        if not self.__wait_on_ratelimit or self.max_retries_reached:
            raise err
        await asyncio.sleep(self.__retry_after)
        self.__retries_performed += 1


class HeaderManager:
    """HTTP ヘッダーのマネージャークラス"""

    def __init__(self, device: Device, state: State, locale="ja") -> None:
        self.__state = state
        self.__locale = locale
        self.__host = API_HOST
        self.__user_agent = device.get_user_agent()
        self.__device_info = device.get_device_info()
        self.__app_version = API_VERSION_NAME
        self.__client_ip = ""
        self.__connection_speed = "0 kbps"
        self.__connection_type = "wifi"
        self.__content_type = "application/json;charset=UTF-8"

    @property
    def locale(self):
        """ロケール"""
        return self.__locale

    @property
    def user_agent(self):
        """ユーザーエージェント"""
        return self.__user_agent

    @property
    def device_info(self) -> str:
        """デバイス情報"""
        return self.__device_info

    @property
    def app_version(self) -> str:
        """アプリバージョン"""
        return self.__app_version

    @property
    def client_ip(self) -> str:
        """クライアント IP アドレス"""
        return self.__client_ip

    @client_ip.setter
    def client_ip(self, value: str) -> None:
        self.__client_ip = value

    @property
    def connection_speed(self) -> str:
        """コネクション速度"""
        return self.__connection_speed

    @property
    def connection_type(self) -> str:
        """コネクション種別"""
        return self.__connection_type

    @property
    def content_type(self) -> str:
        """コンテントタイプ"""
        return self.__content_type

    def generate(self, jwt_required=False) -> Dict[str, str]:
        """HTTPヘッダーを生成する"""
        headers = {
            "Host": self.__host,
            "User-Agent": self.__user_agent,
            "X-Timestamp": str(int(datetime.now().timestamp())),
            "X-App-Version": self.__app_version,
            "X-Device-Info": self.__device_info,
            "X-Device-UUID": self.__state.device_uuid,
            "X-Client-IP": self.__client_ip,
            "X-Connection-Type": self.__connection_type,
            "X-Connection-Speed": self.__connection_speed,
            "Accept-Language": self.__locale,
            "Content-Type": self.__content_type,
        }

        if jwt_required:
            headers.update({"X-Jwt": generate_jwt()})

        if self.__client_ip != "":
            headers.update({"X-Client-IP": self.__client_ip})

        if self.__state.access_token != "":
            headers.update({"Authorization": "Bearer " + self.__state.access_token})

        return headers


current_path = os.path.abspath(os.getcwd())


# pylint: disable=too-many-public-methods
class Client(WebSocketInteractor):
    """yaylib のエントリーポイント"""

    def __init__(
        self,
        *,
        intents: Optional[Intents] = None,
        proxy_url: Optional[str] = None,
        timeout=30,
        max_retries=3,
        backoff_factor=1.5,
        wait_on_ratelimit=True,
        max_ratelimit_retries=15,
        min_delay=0.3,
        max_delay=1.2,
        base_path=current_path + "/.config/",
        state: Optional[State] = None,
        loglevel=logging.INFO,
    ) -> None:
        super().__init__(self, intents)

        self.__proxy_url = proxy_url
        self.__timeout = timeout
        self.__session = None
        self.__min_delay = min_delay
        self.__max_delay = max_delay
        self.__last_request_ts = 0
        self.__max_retries = max_retries
        self.__backoff_factor = backoff_factor

        self.auth = AuthApi(self)
        self.call = CallApi(self)
        self.chat = ChatApi(self)
        self.group = GroupApi(self)
        self.misc = MiscApi(self)
        self.notification = NotificationApi(self)
        self.post = PostApi(self)
        self.review = ReviewApi(self)
        self.thread = ThreadApi(self)
        self.user = UserApi(self)

        if not os.path.exists(base_path):
            os.makedirs(base_path)

        self.__state = state or State(storage_path=base_path + "secret.db")
        self.__header_manager = HeaderManager(Device.create(), self.__state)
        self.__ratelimit = RateLimit(wait_on_ratelimit, max_ratelimit_retries)

        self.logger = logging.getLogger("yaylib version: " + __version__)

        ch = logging.StreamHandler()
        ch.setLevel(loglevel)
        ch.setFormatter(CustomFormatter())

        self.logger.addHandler(ch)
        self.logger.setLevel(logging.DEBUG)

        self.logger.info("yaylib version: %s started.", __version__)

    @property
    def state(self) -> State:
        """状態管理オブジェクト"""
        return self.__state

    @property
    def user_id(self) -> int:
        """ログインしているユーザーの識別子"""
        return self.__state.user_id

    @property
    def access_token(self) -> str:
        """アクセストークン"""
        return self.__state.access_token

    @property
    def refresh_token(self) -> str:
        """リフレッシュトークン"""
        return self.__state.refresh_token

    @property
    def device_uuid(self) -> str:
        """デバイスの識別子"""
        return self.__state.device_uuid

    def __construct_response(
        self, response: Optional[dict], data_type: Optional[Model] = None
    ) -> Optional[dict | Model]:
        """辞書型レスポンスからモデルを生成する"""
        if data_type is not None:
            if isinstance(response, list):
                response = [data_type(result) for result in response]
            elif response is not None:
                response = data_type(response)
        return response

    async def base_request(
        self, method: str, url: str, **kwargs
    ) -> aiohttp.ClientResponse:
        """共通の基底リクエストを行う"""
        self.logger.debug(
            "Making API request: [%s] %s\n\nParameters: %s\n\nHeaders: %s\n\nBody: %s\n",
            method,
            url,
            kwargs.get("params"),
            kwargs.get("headers"),
            kwargs.get("json"),
        )

        session = self.__session or aiohttp.ClientSession()
        async with session.request(
            method, url, proxy=self.__proxy_url, timeout=self.__timeout, **kwargs
        ) as response:
            await response.read()

        if self.__session is None:
            await session.close()

        self.logger.debug(
            "Received API response: [%s] %s\n\nHTTP Status: %s\n\nHeaders: %s\n\nResponse: %s\n",
            method,
            url,
            response.status,
            response.headers,
            await response.text(),
        )

        await raise_for_code(response)
        await raise_for_status(response)

        return response

    async def __make_request(
        self,
        method: str,
        url: str,
        *,
        params: Optional[dict] = None,
        json: Optional[dict] = None,
        headers: Optional[dict] = None,
        return_type: Optional[dict] = None,
    ) -> Optional[dict | Model]:
        """リクエストを行いモデルを生成する"""
        response = await self.base_request(
            method, url, params=params, json=json, headers=headers
        )
        response_json: Optional[dict] = await response.json(content_type=None)
        return self.__construct_response(response_json, return_type)

    async def __refresh_client_tokens(self) -> None:
        """認証トークンのリフレッシュを行う"""
        response = await self.auth.get_token(
            grant_type="refresh_token", refresh_token=self.__state.refresh_token
        )
        self.__state.set_user(
            LocalUser(
                user_id=self.__state.user_id,
                email=self.__state.email,
                device_uuid=self.__state.device_uuid,
                access_token=response.access_token,
                refresh_token=response.refresh_token,
            )
        )
        self.__state.update()

    async def __insert_delay(self) -> None:
        """リクエスト間の時間が1秒未満のときに遅延を挿入する"""
        if int(datetime.now().timestamp()) - self.__last_request_ts < 1:
            await asyncio.sleep(random.uniform(self.__min_delay, self.__max_delay))
        self.__last_request_ts = int(datetime.now().timestamp())

    async def request(
        self,
        method: str,
        url: str,
        *,
        params: Optional[dict] = None,
        json: Optional[dict] = None,
        headers: Optional[dict] = None,
        return_type: Optional[Model] = None,
        jwt_required=False,
    ) -> Optional[dict | Model]:
        """リクエストに必要な処理を行う"""
        if not url.startswith("https://"):
            url = "https://" + url

        if not self.__header_manager.client_ip and "v2/users/timestamp" not in url:
            metadata = await self.user.get_timestamp()
            self.__header_manager.client_ip = metadata.ip_address

        backoff_duration = 0
        headers = headers or {}

        response = None

        for i in range(max(1, self.__max_retries + 1)):
            try:
                await asyncio.sleep(backoff_duration)
                while True:
                    try:
                        await self.__insert_delay()
                        headers.update(self.__header_manager.generate(jwt_required))
                        response = await self.__make_request(
                            method,
                            url,
                            params=filter_dict(params),
                            json=filter_dict(json),
                            headers=headers,
                            return_type=return_type,
                        )
                        break
                    except (QuotaLimitExceededError, TooManyRequestsError) as err:
                        self.logger.warning(
                            "Rate limit exceeded. Waiting... (%s/%s)",
                            self.__ratelimit.retries_performed,
                            self.__ratelimit.max_retries,
                        )
                        await self.__ratelimit.wait(err)
                self.__ratelimit.reset()
                break
            except (AccessTokenExpiredError, AccessTokenInvalidError) as err:
                if self.user_id == 0:
                    self.logger.error("Authentication required to perform the action.")
                    raise err
                await self.__refresh_client_tokens()
            except UnauthorizedError as err:
                if "/api/v1/oauth/token" in url:
                    self.__state.destory(self.user_id)
                    self.logger.error(
                        "Failed to refresh credentials. Please try logging in again."
                    )
                    raise err
            except HTTPInternalServerError as err:
                self.logger.error(
                    "Request failed with status %s! Retrying...", err.response.status
                )
                backoff_duration = self.__backoff_factor * (2**i)

        return response

    # ---------- call api ----------

    def get_user_active_call(self, user_id: int) -> PostResponse:
        """ユーザーが参加中の通話を取得する

        Args:
            user_id (int):

        Returns:
            PostResponse:
        """
        return asyncio.run(self.call.get_user_active_call(user_id))

    def get_bgms(self) -> BgmsResponse:
        """通話のBGMを取得する

        Returns:
            BgmsResponse:
        """
        return asyncio.run(self.call.get_bgms())

    def get_call(self, call_id: int) -> ConferenceCallResponse:
        """通話を取得する

        Args:
            call_id (int):

        Returns:
            ConferenceCallResponse:
        """
        return asyncio.run(self.call.get_call(call_id))

    def get_call_invitable_users(self, **params) -> UsersByTimestampResponse:
        """通話に招待可能なユーザーを取得する

        Args:
            call_id (int):
            from_timestamp (int, optional):

        Returns:
            UsersByTimestampResponse:
        """
        return asyncio.run(self.call.get_call_invitable_users(**params))

    def get_call_status(self, opponent_id: int) -> CallStatusResponse:
        """通話の状態を取得します

        Args:
            opponent_id (int):

        Returns:
            CallStatusResponse:
        """
        return asyncio.run(self.call.get_call_status(opponent_id))

    def get_games(self, **params) -> GamesResponse:
        """通話に設定可能なゲームを取得する

        Args:
            number (int, optional):
            ids (List[int], optional):
            from_id (int, optional):

        Returns:
            GamesResponse:
        """
        return asyncio.run(self.call.get_games(**params))

    def get_genres(self, **params) -> GenresResponse:
        """通話のジャンルを取得する

        Args:
            number (int, optional):
            from (int, optional):

        Returns:
            GenresResponse:
        """
        return asyncio.run(self.call.get_genres(**params))

    def get_group_calls(self, **params) -> PostsResponse:
        """サークル内の通話を取得する

        Args:
            number (int, optional):
            group_category_id (int, optional):
            from_timestamp (int, optional):
            scope (str, optional):

        Returns:
            PostsResponse:
        """
        return asyncio.run(self.call.get_group_calls(**params))

    def invite_online_followings_to_call(self, **params) -> Response:
        """オンラインの友達をまとめて通話に招待します

        Args:
            call_id (int, optional):
            group_id (str, optional):

        Returns:
            Response:
        """
        return asyncio.run(self.call.invite_online_followings_to_call(**params))

    def invite_users_to_call(self, call_id: int, user_ids: List[int]) -> Response:
        """通話に複数のユーザーを招待する

        Args:
            call_id (int):
            user_ids (List[int]):

        Returns:
            Response:
        """
        return asyncio.run(self.call.invite_users_to_call(call_id, user_ids))

    def invite_users_to_chat_call(self, **params) -> Response:
        """ユーザーをチャット通話に招待する

        Args:
            chat_room_id (int):
            room_id (int):
            room_url (str):

        Returns:
            Response:
        """
        return asyncio.run(self.call.invite_users_to_chat_call(**params))

    def kick_user_from_call(
        self, call_id: int, **params
    ) -> CallActionSignatureResponse:
        """ユーザーを通話からキックする

        Args:
            call_id (int):
            uuid (int):
            ban (bool):

        Returns:
            Response:
        """
        asyncio.run(self.call.kick_user_from_call(call_id, **params))

    def start_call(self, call_id: int, **params) -> Response:
        """通話を開始する

        Args:
            call_id (int):
            joinable_by (str):
            game_title (str, optional):
            category_id (str, optional):

        Returns:
            Response:
        """
        return asyncio.run(self.call.start_call(call_id, **params))

    def set_user_role(self, call_id: int, user_id: int, role: str) -> Response:
        """通話の参加者に役職を付与する

        Args:
            call_id (int):
            user_id (int):
            role (str):

        Returns:
            Response:
        """
        return asyncio.run(self.call.set_user_role(call_id, user_id, role))

    def join_call(self, **params) -> ConferenceCallResponse:
        """通話に参加する

        Args:
            conference_id (int):
            call_sid (str, optional):

        Returns:
            ConferenceCallResponse:
        """
        return asyncio.run(self.call.join_call(**params))

    def leave_call(self, **params) -> Response:
        """通話から退出する

        Args:
            conference_id (int):
            call_sid (str, optional):

        Returns:
            Response:
        """
        return asyncio.run(self.call.leave_call(**params))

    def join_call_as_anonymous(self, **params) -> ConferenceCallResponse:
        """匿名で通話に参加する

        Args:
            conference_id (int):
            agora_uid (str):

        Returns:
            ConferenceCallResponse:
        """
        return asyncio.run(self.call.join_call_as_anonymous(**params))

    def leave_call_as_anonymous(self, **params) -> Response:
        """匿名で参加した通話を退出する

        Args:
            conference_id (int):
            agora_uid (str, optional):

        Returns:
            Response: _description_
        """
        return asyncio.run(self.call.leave_call_as_anonymous(**params))

    # ---------- notification api ----------

    def get_activities(self, **params) -> ActivitiesResponse:
        """通知を取得する

        Args:
            important (bool):
            from_timestamp (int, optional):
            number (int, optional):

        Returns:
            ActivitiesResponse:
        """
        return asyncio.run(self.notification.get_activities(**params))

    def get_merged_activities(self, **params) -> ActivitiesResponse:
        """全種類の通知を取得する

        Args:
            from_timestamp (int, optional):
            number (int, optional):

        Returns:
            ActivitiesResponse:
        """
        return asyncio.run(self.notification.get_merged_activities(**params))

    # ---------- chat api ----------

    def accept_chat_requests(self, **params) -> Response:
        """チャットリクエストを承認する

        Args:
            chat_room_ids (List[int]):

        Returns:
            Response:
        """
        return asyncio.run(self.chat.accept_chat_requests(**params))

    def check_unread_status(self, **params) -> UnreadStatusResponse:
        """チャットの未読ステータスを確認する

        Args:
            from_time (int):

        Returns:
            UnreadStatusResponse:
        """
        return asyncio.run(self.chat.check_unread_status(**params))

    def create_group_chat(self, **params) -> CreateChatRoomResponse:
        """グループチャットを作成する

        Args:
            name (str):
            with_user_ids (List[int]):
            icon_filename (str, optional):
            background_filename (str, optional):

        Returns:
            CreateChatRoomResponse:
        """
        return asyncio.run(self.chat.create_group_chat(**params))

    def create_private_chat(self, **params) -> CreateChatRoomResponse:
        """個人チャットを作成する

        Args:
            with_user_id (int):
            matching_id (int, optional):
            hima_chat (bool, optional):

        Returns:
            CreateChatRoomResponse:
        """
        return asyncio.run(self.chat.create_private_chat(**params))

    def delete_chat_background(self, room_id: int) -> Response:
        """チャットの背景を削除する

        Args:
            room_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.chat.delete_chat_background(room_id))

    def delete_message(self, room_id: int, message_id: int) -> Response:
        """チャットメッセージを削除する

        Args:
            room_id (int):
            message_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.chat.delete_message(room_id, message_id))

    def edit_chat_room(self, **params) -> Response:
        """チャットルームを編集する

        Args:
            chat_room_id (int):
            name (str):
            icon_filename (str, optional):
            background_filename (str, optional):

        Returns:
            Response:
        """
        return asyncio.run(self.chat.edit_chat_room(**params))

    def get_chatable_users(self, **params) -> FollowUsersResponse:
        """チャット可能なユーザーを取得する

        Args:
            from_follow_id (int, optional):
            from_timestamp (int, optional):
            order_by (str, optional):

        Returns:
            FollowUsersResponse:
        """
        return asyncio.run(self.chat.get_chatable_users(**params))

    def get_gifs_data(self) -> GifsDataResponse:
        """チャット用 GIF データを取得する

        Returns:
            GifsDataResponse:
        """
        return asyncio.run(self.chat.get_gifs_data())

    def get_hidden_chat_rooms(self, **params) -> ChatRoomsResponse:
        """非表示に設定したチャットルームを取得する

        Args:
            from_timestamp (int, optional):
            number (int, optional)

        Returns:
            ChatRoomsResponse:
        """
        return asyncio.run(self.chat.get_hidden_chat_rooms(**params))

    def get_main_chat_rooms(self, **params) -> ChatRoomsResponse:
        """メインのチャットルームを取得する

        Args:
            from_timestamp (int, optional):

        Returns:
            ChatRoomsResponse:
        """
        return asyncio.run(self.chat.get_main_chat_rooms(**params))

    def get_messages(self, chat_room_id: int, **params) -> MessagesResponse:
        """メッセージを取得する

        Args:
            from_message_id (int, optional):
            to_message_id (int, optional):

        Returns:
            MessagesResponse:
        """
        return asyncio.run(self.chat.get_messages(chat_room_id, **params))

    def get_chat_requests(self, **params) -> ChatRoomsResponse:
        """チャットリクエストを取得する

        Args:
            number (int, optional):
            from_timestamp (int, optional):

        Returns:
            ChatRoomsResponse:
        """
        return asyncio.run(self.chat.get_chat_requests(**params))

    def get_chat_room(self, chat_room_id: int) -> ChatRoomResponse:
        """チャットルームを取得する

        Args:
            chat_room_id (int):

        Returns:
            ChatRoomResponse:
        """
        return asyncio.run(self.chat.get_chat_room(chat_room_id))

    def get_sticker_packs(self) -> StickerPacksResponse:
        """チャット用のスタンプを取得する

        Returns:
            StickerPacksResponse:
        """
        return asyncio.run(self.chat.get_sticker_packs())

    def get_total_chat_requests(self) -> TotalChatRequestResponse:
        """チャットリクエストの総数を取得する

        Returns:
            TotalChatRequestResponse:
        """
        return asyncio.run(self.chat.get_total_chat_requests())

    def hide_chat(self, chat_room_id: int) -> Response:
        """チャットルームを非表示にする

        Args:
            chat_room_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.chat.hide_chat(chat_room_id))

    def invite_to_chat(self, **params) -> Response:
        """チャットルームにユーザーを招待する

        Args:
            chat_room_id (int):
            with_user_ids (List[int]):

        Returns:
            Response:
        """
        return asyncio.run(self.chat.invite_to_chat(**params))

    def kick_users_from_chat(self, **params) -> Response:
        """チャットルームからユーザーを追放する

        Args:
            chat_room_id (int):
            with_user_ids (List[int]):

        Returns:
            Response:
        """
        return asyncio.run(self.chat.kick_users_from_chat(**params))

    def pin_chat(self, room_id: int) -> Response:
        """チャットルームをピン留めする

        Args:
            room_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.chat.pin_chat(room_id))

    def read_message(self, chat_room_id: int, message_id: int) -> Response:
        """メッセージを既読にする

        Args:
            chat_room_id (int):
            message_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.chat.read_message(chat_room_id, message_id))

    def refresh_chat_rooms(self, **params) -> ChatRoomsResponse:
        """チャットルームを更新する

        Args:
            from_time (int, optional):

        Returns:
            ChatRoomsResponse:
        """
        return asyncio.run(self.chat.refresh_chat_rooms(**params))

    def delete_chat_rooms(self, **params) -> Response:
        """チャットルームを削除する

        Args:
            chat_room_ids (List[int]):

        Returns:
            Response:
        """
        return asyncio.run(self.chat.delete_chat_rooms(**params))

    def send_message(
        self,
        chat_room_id: int,
        **params,
    ) -> MessageResponse:
        """チャットを送信する

        Args:
            chat_room_id (int):

        Returns:
            MessageResponse:
        """
        return asyncio.run(self.chat.send_message(chat_room_id, **params))

    def unhide_chat(self, **params) -> Response:
        """非表示に設定したチャットルームを表示する

        Args:
            chat_room_ids (int):

        Returns:
            Response:
        """
        return asyncio.run(self.chat.unhide_chat(**params))

    def unpin_chat(self, chat_room_id: int) -> Response:
        """チャットのピン留めを解除する

        Args:
            chat_room_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.chat.unpin_chat(chat_room_id))

    # ---------- group api ----------

    def accept_moderator_offer(self, group_id: int) -> Response:
        """サークル副管理人の権限オファーを引き受ける

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.group.accept_moderator_offer(group_id))

    def accept_ownership_offer(self, group_id: int) -> Response:
        """サークル管理人の権限オファーを引き受けます

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.group.accept_ownership_offer(group_id))

    def accept_group_join_request(self, group_id: int, user_id: int) -> Response:
        """サークル参加リクエストを承認します

        Args:
            group_id (int):
            user_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.group.accept_group_join_request(group_id, user_id))

    def add_related_groups(
        self, group_id: int, related_group_id: List[int]
    ) -> Response:
        """関連サークルを追加する

        Args:
            group_id (int):
            related_group_id (List[int]):

        Returns:
            Response:
        """
        return asyncio.run(self.group.add_related_groups(group_id, related_group_id))

    def ban_group_user(self, group_id: int, user_id: int) -> Response:
        """サークルからユーザーを追放する

        Args:
            group_id (int):
            user_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.group.ban_group_user(group_id, user_id))

    def check_group_unread_status(self, **params) -> UnreadStatusResponse:
        """サークルの未読ステータスを取得する

        Args:
            from_time (int, optional):

        Returns:
            UnreadStatusResponse:
        """
        return asyncio.run(self.group.check_group_unread_status(**params))

    def create_group(self, **params) -> CreateGroupResponse:
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
        return asyncio.run(self.group.create_group(**params))

    def pin_group(self, group_id: int) -> Response:
        """サークルをピン留めする

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.group.pin_group(group_id))

    def decline_moderator_offer(self, group_id: int) -> Response:
        """サークル副管理人の権限オファーを断る

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.group.decline_moderator_offer(group_id))

    def decline_ownership_offer(self, group_id: int) -> Response:
        """サークル管理人の権限オファーを断る

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.group.decline_ownership_offer(group_id))

    def decline_group_join_request(self, group_id: int, user_id: int) -> Response:
        """サークル参加リクエストを断る

        Args:
            group_id (int):
            user_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.group.decline_group_join_request(group_id, user_id))

    def unpin_group(self, group_id: int) -> Response:
        """サークルのピン留めを解除する

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.group.unpin_group(group_id))

    def get_banned_group_members(self, **params) -> UsersResponse:
        """追放されたサークルメンバーを取得する

        Args:
            group_id (int):
            page (int, optional):

        Returns:
            UsersResponse:
        """
        return asyncio.run(self.group.get_banned_group_members(**params))

    def get_group_categories(self, **params) -> GroupCategoriesResponse:
        """サークルのカテゴリーを取得する

        Args:
            page (int, optional):
            number (int, optional):

        Returns:
            GroupCategoriesResponse:
        """
        return asyncio.run(self.group.get_group_categories(**params))

    def get_create_group_quota(self) -> CreateGroupQuota:
        """残りのサークル作成可能回数を取得する

        Returns:
            CreateGroupQuota:
        """
        return asyncio.run(self.group.get_create_group_quota())

    def get_group(self, group_id: int) -> GroupResponse:
        """サークルの詳細を取得する

        Args:
            group_id (int):

        Returns:
            GroupResponse:
        """
        return asyncio.run(self.group.get_group(group_id))

    def get_groups(self, **params) -> GroupsResponse:
        """複数のサークル情報を取得する

        Args:
            group_category_id (int, optional):
            keyword (str, optional):
            from_timestamp (int, optional):
            sub_category_id (int, optional):

        Returns:
            GroupsResponse:
        """
        return asyncio.run(self.group.get_groups(**params))

    def get_invitable_users(self, group_id: int, **params) -> UsersByTimestampResponse:
        """サークルに招待可能なユーザーを取得する

        Args:
            group_id (int):
            from_timestamp (int, optional):
            user[nickname] (str, optional):

        Returns:
            UsersByTimestampResponse:
        """
        return asyncio.run(self.group.get_invitable_users(group_id, **params))

    def get_joined_statuses(self, ids: List[int]) -> Response:
        """サークルの参加ステータスを取得する

        Args:
            ids (List[int]):

        Returns:
            Response:
        """
        return asyncio.run(self.group.get_joined_statuses(ids))

    def get_group_member(self, group_id: int, user_id: int) -> GroupUserResponse:
        """特定のサークルメンバーの情報を取得する

        Args:
            group_id (int):
            user_id (int):

        Returns:
            GroupUserResponse:
        """
        return asyncio.run(self.group.get_group_member(group_id, user_id))

    def get_group_members(self, group_id: int, **params) -> GroupUsersResponse:
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
        return asyncio.run(self.group.get_group_members(group_id, **params))

    def get_my_groups(self, **params) -> GroupsResponse:
        """自分のサークルを取得する

        Args:
            from_timestamp (_type_, optional):

        Returns:
            GroupsResponse:
        """
        return asyncio.run(self.group.get_my_groups(**params))

    def get_relatable_groups(self, group_id: int, **params) -> GroupsRelatedResponse:
        """関連がある可能性があるサークルを取得する

        Args:
            group_id (int):
            keyword (str, optional):
            from (str, optional):

        Returns:
            GroupsRelatedResponse:
        """
        return asyncio.run(self.group.get_relatable_groups(group_id, **params))

    def get_related_groups(self, group_id: int, **params) -> GroupsRelatedResponse:
        """関連があるサークルを取得する

        Args:
            group_id (int):
            keyword (str, optional):
            from (str, optional):

        Returns:
            GroupsRelatedResponse:
        """
        return asyncio.run(self.group.get_related_groups(group_id, **params))

    def get_user_groups(self, **params) -> GroupsResponse:
        """特定のユーザーが参加しているサークルを取得する

        Args:
            user_id (int):
            page (int, optional):

        Returns:
            GroupsResponse:
        """
        return asyncio.run(self.group.get_user_groups(**params))

    def invite_users_to_group(self, group_id: int, user_ids: List[int]) -> Response:
        """サークルにユーザーを招待する

        Args:
            group_id (int):
            user_ids (List[int]):

        Returns:
            Response:
        """
        return asyncio.run(self.group.invite_users_to_group(group_id, user_ids))

    def join_group(self, group_id: int) -> Response:
        """サークルに参加する

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.group.join_group(group_id))

    def leave_group(self, group_id: int) -> Response:
        """サークルから脱退する

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.group.leave_group(group_id))

    def delete_group_cover(self, group_id: int) -> Response:
        """サークルのカバー画像を削除する

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.group.delete_group_cover(group_id))

    def delete_moderator(self, group_id: int, user_id: int) -> Response:
        """サークルの副管理人を削除する

        Args:
            group_id (int):
            user_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.group.delete_moderator(group_id, user_id))

    def delete_related_groups(
        self, group_id: int, related_group_ids: List[int]
    ) -> Response:
        """関連のあるサークルを削除する

        Args:
            group_id (int):
            related_group_ids (List[int]):

        Returns:
            Response:
        """
        return asyncio.run(
            self.group.delete_related_groups(group_id, related_group_ids)
        )

    def send_moderator_offers(self, group_id: int, user_ids: List[int]) -> Response:
        """複数人にサークル副管理人のオファーを送信する

        Args:
            group_id (int):
            user_ids (List[int]):

        Returns:
            Response:
        """
        return asyncio.run(self.group.send_moderator_offers(group_id, user_ids))

    def send_ownership_offer(self, group_id: int, user_id: int) -> Response:
        """サークル管理人権限のオファーを送信する

        Args:
            group_id (int):
            user_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.group.send_ownership_offer(group_id, user_id))

    def set_group_title(self, group_id: int, title: str) -> Response:
        """サークルのタイトルを設定する

        Args:
            group_id (int):
            title (str):

        Returns:
            Response:
        """
        return asyncio.run(self.group.set_group_title(group_id, title))

    def take_over_group_ownership(self, group_id: int) -> Response:
        """サークル管理人の権限を引き継ぐ

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.group.take_over_group_ownership(group_id))

    def unban_group_member(self, group_id: int, user_id: int) -> Response:
        """特定のサークルメンバーの追放を解除する

        Args:
            group_id (int):
            user_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.group.unban_group_member(group_id, user_id))

    def update_group(self, group_id: int, **params) -> GroupResponse:
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
        return asyncio.run(self.group.update_group(group_id, **params))

    def withdraw_moderator_offer(self, group_id: int, user_id: int) -> Response:
        """サークル副管理人のオファーを取り消す

        Args:
            group_id (int):
            user_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.group.withdraw_moderator_offer(group_id, user_id))

    def withdraw_ownership_offer(self, group_id: int, user_id: int) -> Response:
        """サークル管理人のオファーを取り消す

        Args:
            group_id (int):
            user_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.group.withdraw_ownership_offer(group_id, user_id))

    # ---------- auth api ----------

    def change_email(self, **params) -> LoginUpdateResponse:
        """メールアドレスを変更する

        Args:
            email (str):
            password (str):
            email_grant_token (str, optional):

        Returns:
            LoginUpdateResponse:
        """
        return asyncio.run(self.auth.change_email(**params))

    def change_password(self, **params) -> LoginUpdateResponse:
        """パスワードを変更する

        Args:
            current_password (str):
            new_password (str):

        Returns:
            LoginUpdateResponse:
        """
        return asyncio.run(self.auth.change_password(**params))

    def get_token(self, **params) -> TokenResponse:
        """認証トークンを取得する

        Args:
            grant_type (str):
            refresh_token (str, optional):
            email (str, optional):
            password (str, optional):

        Returns:
            TokenResponse:
        """
        return asyncio.run(self.auth.get_token(**params))

    def login(
        self, email: str, password: str, two_fa_code: Optional[str] = None
    ) -> LoginUserResponse:
        """メールアドレスでログインする

        Args:
            email (str):
            password (str):
            two_fa_code (str, optional):

        Returns:
            LoginUserResponse:
        """
        return asyncio.run(self.auth.login(email, password, two_fa_code))

    def resend_confirm_email(self) -> Response:
        """確認メールを再送信する

        Returns:
            Response:
        """
        return asyncio.run(self.auth.resend_confirm_email())

    def restore_user(self, **params) -> LoginUserResponse:
        """ユーザーを復元する

        Args:
            user_id (int):

        Returns:
            LoginUserResponse:
        """
        return asyncio.run(self.auth.restore_user(**params))

    def save_account_with_email(self, **params) -> LoginUpdateResponse:
        """メールアドレスでアカウントを保存する

        Args:
            email (str):
            password (str, optional):
            current_password (str, optional):
            email_grant_token (str, optional):

        Returns:
            LoginUpdateResponse:
        """
        return asyncio.run(self.auth.save_account_with_email(**params))

    # ---------- misc api ----------

    def accept_policy_agreement(self, agreement_type: str) -> Response:
        """利用規約、ポリシー同意書に同意する

        Args:
            agreement_type (str):

        Returns:
            Response:
        """
        return asyncio.run(self.misc.accept_policy_agreement(agreement_type))

    def send_verification_code(self, email: str, intent: str, locale="ja") -> Response:
        """メールアドレス認証コードを送信する

        Args:
            email (str):
            intent (str):
            locale (str):

        Returns:
            Response:
        """
        return asyncio.run(self.misc.send_verification_code(email, intent, locale))

    def get_email_grant_token(self, **params) -> EmailGrantTokenResponse:
        """メールアドレス認証トークンを取得する

        Args:
            code (int):
            email (str):

        Returns:
            EmailGrantTokenResponse:
        """
        return asyncio.run(self.misc.get_email_grant_token(**params))

    def get_email_verification_presigned_url(
        self, email: str, locale: str = "ja", intent: Optional[str] = None
    ) -> EmailVerificationPresignedUrlResponse:
        """メールアドレス確認用の署名付きURLを取得する

        Args:
            email (str):
            locale (str):
            intent (str, optional):

        Returns:
            EmailVerificationPresignedUrlResponse:
        """
        return asyncio.run(
            self.misc.get_email_verification_presigned_url(email, locale, intent)
        )

    def get_file_upload_presigned_urls(
        self, file_names: List[str]
    ) -> PresignedUrlsResponse:
        """ファイルアップロード用の署名付きURLを取得する

        Args:
            file_names (List[str]):

        Returns:
            PresignedUrlsResponse:
        """
        return asyncio.run(self.misc.get_file_upload_presigned_urls(file_names))

    def get_id_checker_presigned_url(
        self, model: str, action: str, **params
    ) -> IdCheckerPresignedUrlResponse:
        """身分証明用の署名付きURLを取得する

        Args:
            model (str):
            action (str):

        Returns:
            IdCheckerPresignedUrlResponse:
        """
        return asyncio.run(
            self.misc.get_id_checker_presigned_url(model, action, **params)
        )

    def get_old_file_upload_presigned_url(
        self, video_file_name: str
    ) -> PresignedUrlResponse:
        """旧版ファイルアップロード用の署名付きURLを取得する

        Args:
            video_file_name (str):

        Returns:
            PresignedUrlResponse:
        """
        return asyncio.run(self.misc.get_old_file_upload_presigned_url(video_file_name))

    def get_policy_agreed(self) -> PolicyAgreementsResponse:
        """利用規約、ポリシー同意書に同意しているかどうかを取得する

        Returns:
            PolicyAgreementsResponse:
        """
        return asyncio.run(self.misc.get_policy_agreed())

    def get_web_socket_token(self) -> WebSocketTokenResponse:
        """WebSocket トークンを取得する

        Returns:
            WebSocketTokenResponse:
        """
        return asyncio.run(self.misc.get_web_socket_token())

    def upload_image(self, image_paths: List[str], image_type: str) -> List[Attachment]:
        """画像をアップロードして、サーバー上のファイルのリストを取得する

        Examples:
            投稿に画像を付与する場合

            >>> # サーバー上にアップロード
            >>> attachments = client.upload_image(
            >>>     image_type=yaylib.ImageType.POST,
            >>>     image_paths=["./example.jpg"],
            >>> )
            >>> # サーバー上のファイル名を指定
            >>> client.create_post(
            >>>     "Hello with yaylib!",
            >>>     attachment_filename=attachments[0].filename
            >>> )

        Args:
            image_paths (List[str]): 画像ファイルのパスのリスト
            image_type (str): 画像の種類

        Raises:
            ValueError: 画像タイプやフォーマットが不正な場合

        Returns:
            List[Attachment]: サーバー上のファイル情報
        """
        return asyncio.run(self.misc.upload_image(image_paths, image_type))

    def upload_video(self, video_path: str) -> str:
        """動画をアップロードして、サーバー上のファイル名を取得する

        Examples:
            投稿に動画を付与する場合

            >>> # サーバー上にアップロード
            >>> filename = client.upload_video("./example.mp4")
            >>> # サーバー上のファイル名を指定
            >>> client.create_post(
            >>>     "Hello with yaylib!",
            >>>     video_file_name=filename
            >>> )

        Args:
            video_path (str): 動画ファイルのパス

        Raises:
            ValueError: 動画フォーマットが不正な場合

        Returns:
            str: サーバー上のファイル名
        """
        return asyncio.run(self.misc.upload_video(video_path))

    def get_app_config(self) -> ApplicationConfigResponse:
        """アプリケーションのメタデータを取得する

        Returns:
            ApplicationConfigResponse:
        """
        return asyncio.run(self.misc.get_app_config())

    def get_banned_words(self, country_code: str = "jp") -> BanWordsResponse:
        """禁止ワードの一覧を取得する

        Args:
            country_code (str, optional):

        Returns:
            BanWordsResponse:
        """
        return asyncio.run(self.misc.get_banned_words(country_code))

    def get_popular_words(self, country_code: str = "jp") -> PopularWordsResponse:
        """人気ワードの一覧を取得する

        Args:
            country_code (str, optional):

        Returns:
            PopularWordsResponse:
        """
        return asyncio.run(self.misc.get_popular_words(country_code))

    # ---------- post api ----------

    def add_bookmark(self, user_id: int, post_id: int) -> BookmarkPostResponse:
        """ブックマークに追加する

        Args:
            user_id (int):
            post_id (int):

        Returns:
            BookmarkPostResponse:
        """
        return asyncio.run(self.post.add_bookmark(user_id, post_id))

    def add_group_highlight_post(self, group_id: int, post_id: int) -> Response:
        """投稿をグループのまとめに追加する

        Args:
            group_id (int):
            post_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.post.add_group_highlight_post(group_id, post_id))

    def create_call_post(
        self, text: Optional[str] = None, **params
    ) -> CreatePostResponse:
        """通話の投稿を作成する

        Args:
            text (str, optional):
            font_size (int, optional):
            color (int, optional):
            group_id (int, optional):
            call_type (str, optional):
            category_id (int, optional):
            game_title (str, optional):
            joinable_by (str, optional):
            message_tags (List, optional):
            attachment_filename (str, optional):
            attachment_2_filename (str, optional):
            attachment_3_filename (str, optional):
            attachment_4_filename (str, optional):
            attachment_5_filename (str, optional):
            attachment_6_filename (str, optional):
            attachment_7_filename (str, optional):
            attachment_8_filename (str, optional):
            attachment_9_filename (str, optional):

        Returns:
            CreatePostResponse:
        """
        return asyncio.run(self.post.create_call_post(text, **params))

    def pin_group_post(self, post_id: int, group_id: int) -> Response:
        """サークルの投稿をピン留めする

        Args:
            post_id (int):
            group_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.post.pin_group_post(post_id, group_id))

    def pin_post(self, post_id: int) -> Response:
        """プロフィールに投稿をピン留めする

        Args:
            post_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.post.pin_post(post_id))

    def create_post(self, text: Optional[str] = None, **params) -> Post:
        """投稿を作成する

        Args:
            text (str, optional):
            font_size (int, optional): . Defaults to 0.
            color (int, optional): . Defaults to 0.
            in_reply_to (int, optional):
            group_id (int, optional):
            mention_ids (List[int], optional):
            choices (List[str], optional):
            shared_url (str, optional):
            message_tags (List, optional):
            attachment_filename (str, optional):
            attachment_2_filename (str, optional):
            attachment_3_filename (str, optional):
            attachment_4_filename (str, optional):
            attachment_5_filename (str, optional):
            attachment_6_filename (str, optional):
            attachment_7_filename (str, optional):
            attachment_8_filename (str, optional):
            attachment_9_filename (str, optional):
            video_file_name (str, optional):

        Returns:
            Post:
        """
        return asyncio.run(self.post.create_post(text, **params))

    def create_repost(
        self, post_id: int, text: Optional[str] = None, **params
    ) -> CreatePostResponse:
        """投稿を(´∀｀∩)↑age↑する

        Args:
            post_id (int):
            text (str, optional):
            font_size (int, optional):
            color (int, optional):
            in_reply_to (int, optional):
            group_id (int, optional):
            mention_ids (List[int], optional):
            choices (List[str], optional):
            shared_url (Dict[str, str  |  int], optional):
            message_tags (List, optional):
            attachment_filename (str, optional):
            attachment_2_filename (str, optional):
            attachment_3_filename (str, optional):
            attachment_4_filename (str, optional):
            attachment_5_filename (str, optional):
            attachment_6_filename (str, optional):
            attachment_7_filename (str, optional):
            attachment_8_filename (str, optional):
            attachment_9_filename (str, optional):
            video_file_name (str, optional):

        Returns:
            CreatePostResponse:
        """
        return asyncio.run(self.post.create_repost(post_id, text, **params))

    def create_share_post(
        self,
        shareable_type: str,
        shareable_id: int,
        text: Optional[str] = None,
        **params,
    ) -> Post:
        """シェア投稿を作成する

        Args:
            shareable_type (str):
            shareable_id (int):
            text (str, optional):
            font_size (int, optional):
            color (int, optional):
            group_id (int, optional):

        Returns:
            Post:
        """
        return asyncio.run(
            self.post.create_share_post(shareable_type, shareable_id, text, **params)
        )

    def create_thread_post(
        self, thread_id: int, text: Optional[str] = None, **params
    ) -> Post:
        """スレッドの投稿を作成する

        Args:
            thread_id (int):
            text (str, optional):
            font_size (int, optional):
            color (int, optional):
            in_reply_to (int, optional):
            group_id (int, optional):
            mention_ids (List[int], optional):
            choices (list[str], optional):
            shared_url (Dict[str, str  |  int], optional):
            message_tags (List, optional):
            attachment_filename (str, optional):
            attachment_2_filename (str, optional):
            attachment_3_filename (str, optional):
            attachment_4_filename (str, optional):
            attachment_5_filename (str, optional):
            attachment_6_filename (str, optional):
            attachment_7_filename (str, optional):
            attachment_8_filename (str, optional):
            attachment_9_filename (str, optional):
            video_file_name (str, optional):

        Returns:
            Post:
        """
        return asyncio.run(self.post.create_thread_post(thread_id, text, **params))

    def delete_all_posts(self) -> Response:
        """すべての自分の投稿を削除する

        Returns:
            Response:
        """
        return asyncio.run(self.post.delete_all_posts())

    def unpin_group_post(self, group_id: int) -> Response:
        """グループのピン投稿を解除する

        Args:
            group_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.post.unpin_group_post(group_id))

    def unpin_post(self, post_id: int) -> Response:
        """プロフィール投稿のピンを解除する

        Args:
            post_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.post.unpin_post(post_id))

    def get_bookmark(self, user_id: int, **params) -> PostsResponse:
        """ブックマークを取得する

        Args:
            user_id (int):
            from_str (str, optional):

        Returns:
            PostsResponse:
        """
        return asyncio.run(self.post.get_bookmark(user_id, **params))

    def get_timeline_calls(self, **params) -> PostsResponse:
        """誰でも通話を取得する

        Args:
            group_id (int, optional):
            from_timestamp (int, optional):
            number (int, optional):
            category_id (int, optional):
            call_type (str, optional): Defaults to "voice".
            include_circle_call (bool, optional):
            cross_generation (bool, optional):
            exclude_recent_gomimushi (bool, optional):
            shared_interest_categories (bool, optional):

        Returns:
            PostsResponse:
        """
        return asyncio.run(self.post.get_timeline_calls(**params))

    def get_conversation(self, conversation_id: int, **params) -> PostsResponse:
        """リプライを含める投稿の会話を取得する

        Args:
            conversation_id (int):
            group_id (int, optional):
            thread_id (int, optional):
            from_post_id (int, optional):
            number (int, optional):
            reverse (bool, optional):

        Returns:
            PostsResponse:
        """
        return asyncio.run(self.post.get_conversation(conversation_id, **params))

    def get_conversation_root_posts(self, post_ids: List[int]) -> PostsResponse:
        """会話の原点の投稿を取得する

        Args:
            post_ids (List[int]):

        Returns:
            PostsResponse:
        """
        return asyncio.run(self.post.get_conversation_root_posts(post_ids))

    def get_following_call_timeline(self, **params) -> PostsResponse:
        """フォロー中の通話を取得する

        Args:
            from_timestamp (int, optional):
            number (int, optional):
            category_id (int, optional):
            include_circle_call (bool, optional):
            exclude_recent_gomimushi (bool, optional):

        Returns:
            PostsResponse:
        """
        return asyncio.run(self.post.get_following_call_timeline(**params))

    def get_following_timeline(self, **params) -> PostsResponse:
        """フォロー中のタイムラインを取得する

        Args:
            from_str (str, optional):
            only_root (bool, optional):
            order_by (str, optional):
            number (int, optional):
            mxn (int, optional):
            reduce_selfie (bool, optional):
            custom_generation_range (bool, optional):

        Returns:
            PostsResponse:
        """
        return asyncio.run(self.post.get_following_timeline(**params))

    def get_group_highlight_posts(self, group_id: int, **params) -> PostsResponse:
        """グループのまとめ投稿を取得する

        Args:
            group_id (int):
            from_post (int, optional):
            number (int, optional):

        Returns:
            PostsResponse:
        """
        return asyncio.run(self.post.get_group_highlight_posts(group_id, **params))

    def get_group_timeline_by_keyword(
        self, group_id: int, keyword: str, **params
    ) -> PostsResponse:
        """グループの投稿をキーワードで検索する

        Args:
            group_id (int):
            keyword (str):
            from_post_id (int, optional):
            number (int, optional):
            only_thread_posts (bool, optional):

        Returns:
            PostsResponse:
        """
        return asyncio.run(
            self.post.get_group_timeline_by_keyword(group_id, keyword, **params)
        )

    def get_group_timeline(self, group_id: int, **params) -> PostsResponse:
        """グループのタイムラインを取得する

        Args:
            group_id (int):
            from_post_id (int, optional):
            reverse (bool, optional):
            post_type (str, optional):
            number (int, optional):
            only_root (bool, optional):

        Returns:
            PostsResponse:
        """
        return asyncio.run(self.post.get_group_timeline(group_id, **params))

    def get_timeline_by_hashtag(self, hashtag: str, **params) -> PostsResponse:
        """ハッシュタグでタイムラインを検索する

        Args:
            hashtag (str):
            from_post_id (int, optional):
            number (int, optional):

        Returns:
            PostsResponse:
        """
        return asyncio.run(self.post.get_timeline_by_hashtag(hashtag, **params))

    def get_my_posts(self, **params) -> PostsResponse:
        """自分の投稿を取得する

        Args:
            from_post_id (int, optional):
            number (int, optional):
            include_group_post (bool, optional):

        Returns:
            PostsResponse:
        """
        return asyncio.run(self.post.get_my_posts(**params))

    def get_post(self, post_id: int) -> PostResponse:
        """投稿の詳細を取得する

        Args:
            post_id (int):

        Returns:
            PostResponse:
        """
        return asyncio.run(self.post.get_post(post_id))

    def get_post_likers(self, post_id: int, **params) -> PostLikersResponse:
        """投稿にいいねしたユーザーを取得する

        Args:
            post_id (int):
            from_id (int, optional):
            number (int, optional):

        Returns:
            PostLikersResponse:
        """
        return asyncio.run(self.post.get_post_likers(post_id, **params))

    def get_reposts(self, post_id: int, **params: int) -> PostsResponse:
        """投稿の(´∀｀∩)↑age↑を取得する

        Args:
            post_id (int):
            from_post_id (int, optional):
            number (int, optional):

        Returns:
            PostsResponse:
        """
        return asyncio.run(self.post.get_reposts(post_id, **params))

    def get_posts(self, post_ids: List[int]) -> PostsResponse:
        """複数の投稿を取得する

        Args:
            post_ids (List[int]):

        Returns:
            PostsResponse:
        """
        return asyncio.run(self.post.get_posts(post_ids))

    def get_recommended_post_tags(self, **params) -> PostTagsResponse:
        """おすすめのタグ候補を取得する

        Args:
            tag (str, optional):
            save_recent_search (bool, optional):

        Returns:
            PostTagsResponse:
        """
        return asyncio.run(self.post.get_recommended_post_tags(**params))

    def get_recommended_posts(self, **params) -> PostsResponse:
        """おすすめの投稿を取得する

        Args:
            experiment_num (int):
            variant_num (int, optional):
            number (int, optional):


        Returns:
            PostsResponse:
        """
        return asyncio.run(self.post.get_recommended_posts(**params))

    def get_timeline_by_keyword(
        self,
        keyword: Optional[str] = None,
        **params,
    ) -> PostsResponse:
        """キーワードでタイムラインを検索する

        Args:
            keyword (str, optional):
            from_post_id (int, optional):
            number (int, optional):

        Returns:
            PostsResponse:
        """
        return asyncio.run(self.post.get_timeline_by_keyword(keyword, **params))

    def get_timeline(self, **params) -> PostsResponse:
        """タイムラインを取得する

        Args:
            noreply_mode (bool, optional):
            from_post_id (int, optional):
            number (int, optional):
            order_by (str, optional):
            experiment_older_age_rules (bool, optional):
            shared_interest_categories (bool, optional):
            mxn (int, optional):
            en (int, optional):
            vn (int, optional):
            reduce_selfie (bool, optional):
            custom_generation_range (bool, optional):

        Returns:
            PostsResponse:
        """
        return asyncio.run(self.post.get_timeline(**params))

    def get_url_metadata(self, url: str) -> SharedUrl:
        """URLのメタデータを取得する

        Args:
            url (str):

        Returns:
            SharedUrl:
        """
        return asyncio.run(self.post.get_url_metadata(url))

    def get_user_timeline(self, user_id: int, **params) -> PostsResponse:
        """ユーザーのタイムラインを取得する

        Args:
            user_id (int):
            from_post_id (int, optional):
            number (int, optional):
            post_type (str, optional):

        Returns:
            PostsResponse:
        """
        return asyncio.run(self.post.get_user_timeline(user_id, **params))

    def like(self, post_ids: List[int]) -> LikePostsResponse:
        """投稿にいいねする

        Args:
            post_ids (List[int]):

        Returns:
            LikePostsResponse:
        """
        return asyncio.run(self.post.like(post_ids))

    def delete_bookmark(self, user_id: int, post_id: int) -> Response:
        """ブックマークを削除する

        Args:
            user_id (int):
            post_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.post.delete_bookmark(user_id, post_id))

    def delete_group_highlight_post(self, group_id: int, post_id: int) -> Response:
        """サークルのまとめから投稿を解除する

        Args:
            group_id (int):
            post_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.post.delete_group_highlight_post(group_id, post_id))

    def delete_posts(self, post_ids: List[int]) -> Response:
        """投稿を削除する

        Args:
            post_ids (List[int]):

        Returns:
            Response:
        """
        return asyncio.run(self.post.delete_posts(post_ids))

    def unlike(self, post_id: int) -> Response:
        """いいねを解除する

        Args:
            post_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.post.unlike(post_id))

    def update_post(self, **params) -> Post:
        """投稿を編集する

        Args:
            post_id (int):
            text (str, optional):
            font_size (int, optional):
            color (int, optional):
            message_tags (List, optional):

        Returns:
            Post:
        """
        return asyncio.run(self.post.update_post(**params))

    def view_video(self, video_id: int) -> Response:
        """動画を視聴する

        Args:
            video_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.post.view_video(video_id))

    def vote_survey(self, survey_id: int, choice_id: int) -> VoteSurveyResponse:
        """アンケートに投票する

        Args:
            survey_id (int):
            choice_id (int):

        Returns:
            VoteSurveyResponse:
        """
        return asyncio.run(self.post.vote_survey(survey_id, choice_id))

    # ---------- review api ----------

    def create_review(self, user_id: int, comment: str) -> Response:
        """レターを送信する

        Args:
            user_id (int):
            comment (str):

        Returns:
            Response:
        """
        return asyncio.run(self.review.create_review(user_id, comment))

    def delete_reviews(self, review_ids: List[int]) -> Response:
        """レターを削除する

        Args:
            review_ids (List[int]):

        Returns:
            Response:
        """
        return asyncio.run(self.review.delete_reviews(review_ids))

    def get_my_reviews(self, **params) -> ReviewsResponse:
        """送信したレターを取得する

        Args:
            from_id (int, optional):
            number (int, optional):

        Returns:
            ReviewsResponse:
        """
        return asyncio.run(self.review.get_my_reviews(**params))

    def get_reviews(self, user_id: int, **params) -> ReviewsResponse:
        """ユーザーが受け取ったレターを取得する

        Args:
            user_id (int):
            from_id (int, optional):
            number (int, optional):

        Returns:
            ReviewsResponse:
        """
        return asyncio.run(self.review.get_reviews(user_id, **params))

    def pin_review(self, review_id: int) -> Response:
        """レターをピン留めする

        Args:
            review_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.review.pin_review(review_id))

    def unpin_review(self, review_id: int) -> Response:
        """レターのピン留めを解除する

        Args:
            review_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.review.unpin_review(review_id))

    # ---------- thread api ----------

    def add_post_to_thread(self, post_id: int, thread_id: int) -> ThreadInfo:
        """投稿をスレッドに追加する

        Args:
            post_id (int):
            thread_id (int):

        Returns:
            ThreadInfo:
        """
        return asyncio.run(self.thread.add_post_to_thread(post_id, thread_id))

    def convert_post_to_thread(self, post_id: int, **params) -> ThreadInfo:
        """投稿をスレッドに変換する

        Args:
            post_id (int):
            title (str, optional):
            thread_icon_filename (str, optional):

        Returns:
            ThreadInfo:
        """
        return asyncio.run(self.thread.convert_post_to_thread(post_id, **params))

    def create_thread(
        self,
        group_id: int,
        title: str,
        thread_icon_filename: str,
    ) -> ThreadInfo:
        """スレッドを作成する

        Args:
            group_id (int):
            title (str):
            thread_icon_filename (str):

        Returns:
            ThreadInfo:
        """
        return asyncio.run(
            self.thread.create_thread(group_id, title, thread_icon_filename)
        )

    def get_group_thread_list(self, **params) -> GroupThreadListResponse:
        """サークルのスレッド一覧を取得する

        Args:
            group_id (int):
            from_str (str, optional):
            join_status (str, optional):

        Returns:
            GroupThreadListResponse:
        """
        return asyncio.run(self.thread.get_group_thread_list(**params))

    def get_thread_joined_statuses(self, ids: List[int]) -> Response:
        """スレッド参加ステータスを取得する

        Args:
            ids (List[int]):

        Returns:
            Response:
        """
        return asyncio.run(self.thread.get_thread_joined_statuses(ids))

    def get_thread_posts(self, thread_id: int, **params) -> PostsResponse:
        """スレッド内のタイムラインを取得する

        Args:
            thread_id (int):
            from_str (str, optional):
            number (str, optional):

        Returns:
            PostsResponse:
        """
        return asyncio.run(self.thread.get_thread_posts(thread_id, **params))

    def join_thread(self, thread_id: int, user_id: int) -> Response:
        """スレッドに参加する

        Args:
            thread_id (int):
            user_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.thread.join_thread(thread_id, user_id))

    def leave_thread(self, thread_id: int, user_id: int) -> Response:
        """スレッドから脱退する

        Args:
            thread_id (int):
            user_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.thread.leave_thread(thread_id, user_id))

    def delete_thread(self, thread_id: int) -> Response:
        """スレッドを削除する

        Args:
            thread_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.thread.delete_thread(thread_id))

    def update_thread(self, thread_id: int, **params) -> Response:
        """スレッドをアップデートする

        Args:
            thread_id (int):
            title (str):
            thread_icon_filename (str):

        Returns:
            Response:
        """
        return asyncio.run(self.thread.update_thread(thread_id, **params))

    # ---------- user api ----------

    def delete_footprint(self, user_id: int, footprint_id: int) -> Response:
        """足跡を削除する

        Args:
            user_id (int):
            footprint_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.user.delete_footprint(user_id, footprint_id))

    def follow_user(self, user_id: int) -> Response:
        """ユーザーをフォローする

        Args:
            user_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.user.follow_user(user_id))

    def follow_users(self, user_ids: List[int]) -> Response:
        """複数のユーザーをフォローする

        Args:
            user_ids (List[int]):

        Returns:
            Response:
        """
        return asyncio.run(self.user.follow_users(user_ids))

    def get_active_followings(self, **params) -> ActiveFollowingsResponse:
        """アクティブなフォロー中のユーザーを取得する

        Args:
            only_online (bool, optional):
            from_loggedin_at (int, optional):

        Returns:
            ActiveFollowingsResponse:
        """
        return asyncio.run(self.user.get_active_followings(**params))

    def get_follow_recommendations(self, **params) -> FollowRecommendationsResponse:
        """フォローするのにおすすめのユーザーを取得する

        Args:
            from_timestamp (int, optional):
            number (int, optional):
            sources (List[str], optional):

        Returns:
            FollowRecommendationsResponse:
        """
        return asyncio.run(self.user.get_follow_recommendations(**params))

    def get_follow_request(self, **params) -> UsersByTimestampResponse:
        """フォローリクエストを取得する

        Args:
            from_timestamp (int, optional):

        Returns:
            UsersByTimestampResponse:
        """
        return asyncio.run(self.user.get_follow_request(**params))

    def get_follow_request_count(self) -> FollowRequestCountResponse:
        """フォローリクエストの数を取得する

        Returns:
            FollowRequestCountResponse:
        """
        return asyncio.run(self.user.get_follow_request_count())

    def get_following_users_born(self, **params) -> UsersResponse:
        """フォロー中のユーザーの誕生日を取得する

        Args:
            birthdate (int, optional):

        Returns:
            UsersResponse:
        """
        return asyncio.run(self.user.get_following_users_born(**params))

    def get_footprints(self, **params) -> FootprintsResponse:
        """足跡を取得する

        Args:
            from_id (int, optional):
            number (int, optional):
            mode (str, optional):

        Returns:
            FootprintsResponse:
        """
        return asyncio.run(self.user.get_footprints(**params))

    def get_fresh_user(self, user_id: int) -> UserResponse:
        """認証情報などを含んだユーザー情報を取得する

        Args:
            user_id (int):

        Returns:
            UserResponse:
        """
        return asyncio.run(self.user.get_fresh_user(user_id))

    def get_hima_users(self, **params) -> HimaUsersResponse:
        """暇なユーザーを取得する

        Args:
            from_hima_id (int, optional):
            number (int, optional):

        Returns:
            HimaUsersResponse:
        """
        return asyncio.run(self.user.get_hima_users(**params))

    def get_user_ranking(self, mode: str) -> RankingUsersResponse:
        """ユーザーのフォロワーランキングを取得する

        Examples:
            >>> # ルーキーを取得する:
            >>> client.get_user_ranking(mode="one_month")

            >>> # ミドルを取得する:
            >>> client.get_user_ranking(mode="six_months")

            >>> # 殿堂入りを取得する:
            >>> client.get_user_ranking(mode="all_time")

        Args:
            mode (str):

        Returns:
            RankingUsersResponse:
        """
        return asyncio.run(self.user.get_user_ranking(mode))

    def get_profile_refresh_counter_requests(self) -> RefreshCounterRequestsResponse:
        """投稿数やフォロワー数をリフレッシュするための残リクエスト数を取得する

        Returns:
            RefreshCounterRequestsResponse:
        """
        return asyncio.run(self.user.get_profile_refresh_counter_requests())

    def get_social_shared_users(self, **params) -> SocialShareUsersResponse:
        """SNS共有をしたユーザーを取得する

        Args:
            sns_name (str):
            number (int, optional):
            from_id (int, optional):

        Returns:
            SocialShareUsersResponse:
        """
        return asyncio.run(self.user.get_social_shared_users(**params))

    def get_timestamp(self) -> UserTimestampResponse:
        """タイムスタンプを取得する

        Returns:
            UserTimestampResponse:
        """
        return asyncio.run(self.user.get_timestamp())

    def get_user(self, user_id: int) -> UserResponse:
        """ユーザーの情報を取得する

        Args:
            user_id (int):

        Returns:
            UserResponse:
        """
        return asyncio.run(self.user.get_user(user_id))

    def get_user_followers(self, user_id: int, **params) -> FollowUsersResponse:
        """ユーザーのフォロワーを取得する

        Args:
            user_id (int):
            from_follow_id (int, optional):
            followed_by_me: (int, optional):
            number: (int, optional):

        Returns:
            FollowUsersResponse:
        """
        return asyncio.run(self.user.get_user_followers(user_id, **params))

    def get_user_followings(self, user_id: int, **params) -> FollowUsersResponse:
        """フォロー中のユーザーを取得する

        Args:
            user_id (int):
            from_follow_id (int, optional):
            from_timestamp (int, optional):
            order_by (str, optional):
            number (int, optional):

        Returns:
            FollowUsersResponse:
        """
        return asyncio.run(self.user.get_user_followings(user_id, **params))

    def get_user_from_qr(self, qr: str) -> UserResponse:
        """QRコードからユーザーを取得する

        Args:
            qr (str):

        Returns:
            UserResponse:
        """
        return asyncio.run(self.user.get_user_from_qr(qr))

    def get_user_without_leaving_footprint(self, user_id: int) -> UserResponse:
        """足跡をつけずにユーザーの情報を取得する

        Args:
            user_id (int):

        Returns:
            UserResponse:
        """
        return asyncio.run(self.user.get_user_without_leaving_footprint(user_id))

    def get_users(self, user_ids: List[int]) -> UsersResponse:
        """複数のユーザーの情報を取得する

        Args:
            user_ids (List[int]):

        Returns:
            UsersResponse:
        """
        return asyncio.run(self.user.get_users(user_ids))

    def refresh_profile_counter(self, counter: str) -> Response:
        """プロフィールのカウンターを更新する

        Args:
            counter (str):

        Returns:
            Response:
        """
        return asyncio.run(self.user.refresh_profile_counter(counter))

    def register(self, **params) -> CreateUserResponse:
        """
        Args:
            email (str):
            email_grant_token (str):
            password (str):
            nickname (str):
            birth_date (str):
            gender (int, optional):
            country_code (str, optional):
            biography (str, optional):
            prefecture (str, optional):
            profile_icon_filename (str, optional):
            cover_image_filename (str, optional):
            en (int, optional):
            vn (int, optional):

        Returns:
            CreateUserResponse:
        """
        return asyncio.run(self.user.register(**params))

    def delete_user_avatar(self) -> Response:
        """ユーザーのアイコンを削除する

        Returns:
            Response:
        """
        return asyncio.run(self.user.delete_user_avatar())

    def delete_user_cover(self) -> Response:
        """ユーザーのカバー画像を削除する

        Returns:
            Response:
        """
        return asyncio.run(self.user.delete_user_cover())

    def reset_password(self, **params) -> Response:
        """パスワードをリセットする

        Args:
            email (str):
            email_grant_token (str):
            password (str):

        Returns:
            Response:
        """
        return asyncio.run(self.user.reset_password(**params))

    def search_lobi_users(self, **params) -> UsersResponse:
        """Lobiのユーザーを検索する

        Args:
            nickname (str, optional):
            number (int, optional):
            from_str (str, optional):

        Returns:
            UsersResponse:
        """
        return asyncio.run(self.user.search_lobi_users(**params))

    def search_users(self, **params) -> UsersResponse:
        """ユーザーを検索する

        Args:
            gender (int, optional):
            nickname (str, optional):
            title (str, optional):
            biography (str, optional):
            from_timestamp (int, optional):
            similar_age (bool, optional):
            not_recent_gomimushi (bool, optional):
            recently_created (bool, optional):
            same_prefecture (bool, optional):
            save_recent_search (bool, optional):

        Returns:
            UsersResponse:
        """
        return asyncio.run(self.user.search_users(**params))

    def set_follow_permission_enabled(self, **params) -> Response:
        """フォローを許可制にするかを設定する

        Args:
            nickname (str):
            is_private (bool, optional):

        Returns:
            Response:
        """
        return asyncio.run(self.user.set_follow_permission_enabled(**params))

    def take_action_follow_request(self, user_id: int, action: str) -> Response:
        """フォローリクエストを操作する

        Args:
            user_id (int):
            action (str):

        Returns:
            Response:
        """
        return asyncio.run(self.user.take_action_follow_request(user_id, action))

    def turn_on_hima(self) -> Response:
        """ひまなうを有効にする

        Returns:
            Response:
        """
        return asyncio.run(self.user.turn_on_hima())

    def unfollow_user(self, user_id: int) -> Response:
        """ユーザーをアンフォローする

        Args:
            user_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.user.unfollow_user(user_id))

    def update_user(self, nickname: str, **params) -> Response:
        """プロフィールを更新する

        Args:
            nickname (str):
            biography (str, optional):
            prefecture (str, optional):
            gender (int, optional):
            country_code (str, optional):
            profile_icon_filename (str, optional):
            cover_image_filename (str, optional):
            username (str, optional):

        Returns:
            Response:
        """
        return asyncio.run(self.user.update_user(nickname, **params))

    def block_user(self, user_id: int) -> Response:
        """ユーザーをブロックする

        Args:
            user_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.user.block_user(user_id))

    def get_blocked_user_ids(self) -> BlockedUserIdsResponse:
        """あなたをブロックしたユーザーを取得する

        Returns:
            BlockedUserIdsResponse:
        """
        return asyncio.run(self.user.get_blocked_user_ids())

    def get_blocked_users(self, **params) -> BlockedUsersResponse:
        """ブロックしたユーザーを取得する

        Args:
            from_id (int, optional):

        Returns:
            BlockedUsersResponse:
        """
        return asyncio.run(self.user.get_blocked_users(**params))

    def unblock_user(self, user_id: int) -> Response:
        """ユーザーをアンブロックする

        Args:
            user_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.user.unblock_user(user_id))

    def get_hidden_users_list(self, **params) -> HiddenResponse:
        """非表示のユーザー一覧を取得する

        Args:
            from_str (str, optional):
            number (int, optional):

        Returns:
            HiddenResponse:
        """
        return asyncio.run(self.user.get_hidden_users_list(**params))

    def hide_user(self, user_id: int) -> Response:
        """ユーザーを非表示にする

        Args:
            user_id (int):

        Returns:
            Response:
        """
        return asyncio.run(self.user.hide_user(user_id))

    def unhide_users(self, user_ids: List[int]) -> Response:
        """ユーザーの非表示を解除する

        Args:
            user_ids (List[int]):

        Returns:
            Response:
        """
        return asyncio.run(self.user.unhide_users(user_ids))
