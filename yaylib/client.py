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
import time
from datetime import datetime
from typing import Awaitable, Callable, Dict, Optional

import aiohttp

from . import __version__, config, utils
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
from .device import Device
from .errors import (
    AccessTokenExpiredError,
    AuthenticationError,
    BadRequestError,
    ErrorCode,
    ForbiddenError,
    HTTPError,
    InternalServerError,
    NotFoundError,
    RateLimitError,
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
    CallStatusResponse,
    ChatRoomResponse,
    ChatRoomsResponse,
    ConferenceCallResponse,
    CreateChatRoomResponse,
    CreateGroupResponse,
    CreatePostResponse,
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

__all__ = ["Client"]


class RateLimit:
    """レート制限を管理するクラス"""

    def __init__(self, wait_on_ratelimit: bool, max_retries: int) -> None:
        self.__wait_on_ratelimit = wait_on_ratelimit
        self.__max_retries = max_retries
        self.__retries_performed = 0
        self.__retry_after = 60 * 5

    @property
    def retries_performed(self) -> int:
        """レート制限によるリトライ回数

        Returns:
            int: リトライ回数
        """
        return self.__retries_performed

    @property
    def max_retries(self) -> int:
        """レート制限によるリトライ回数の上限

        Returns:
            int: リトライ回数の上限
        """
        return self.__max_retries

    def __max_retries_reached(self) -> bool:
        return not self.__wait_on_ratelimit or (
            self.__retries_performed >= self.__max_retries
        )

    def reset(self) -> None:
        """レート制限をリセットする"""
        self.__retries_performed = 0

    async def wait(self, exc: RateLimitError) -> None:
        """レート制限が解除されるまで待機する

        Raises:
            RateLimitError: レート制限エラー
        """
        if not self.__wait_on_ratelimit or self.__max_retries_reached():
            raise exc
        await asyncio.sleep(self.__retry_after)
        self.__retries_performed += 1


class HeaderManager:
    """HTTP ヘッダーのマネージャークラス"""

    def __init__(self, device: Device, state: State, locale="ja") -> None:
        self.__state = state
        self.__locale = locale
        self.__host = config.API_HOST
        self.__user_agent = device.generate_user_agent()
        self.__device_info = device.generate_device_info(config.VERSION_NAME)
        self.__app_version = config.API_VERSION_NAME
        self.__client_ip = ""
        self.__connection_speed = "0 kbps"
        self.__connection_type = "wifi"
        self.__content_type = "application/json;charset=UTF-8"

    @property
    def client_ip(self) -> str:
        return self.__client_ip

    @client_ip.setter
    def client_ip(self, value: str) -> None:
        self.__client_ip = value

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
            headers.update({"X-Jwt": utils.generate_jwt()})

        if self.__client_ip != "":
            headers.update({"X-Client-IP": self.__client_ip})

        if self.__state.access_token != "":
            headers.update({"Authorization": "Bearer " + self.__state.access_token})

        return headers


current_path = os.path.abspath(os.getcwd())


class BaseClient:
    """yaylib クライアントの基底クラス"""

    def __init__(
        self,
        *,
        proxy_url: Optional[str] = None,
        timeout=60,
        base_path=current_path + "/.config/",
        loglevel=logging.INFO,
    ) -> None:
        self.__proxy_url = proxy_url
        self.__timeout = timeout
        self.__session = None

        self.logger = logging.getLogger("yaylib version: " + __version__)

        if not os.path.exists(base_path):
            os.makedirs(base_path)

        ch = logging.StreamHandler()
        ch.setLevel(loglevel)
        ch.setFormatter(utils.CustomFormatter())

        self.logger.addHandler(ch)
        self.logger.setLevel(logging.DEBUG)

        self.logger.info("yaylib version: %s started.", __version__)

    async def __is_ratelimit_error(self, response: aiohttp.ClientResponse) -> bool:
        if response.status == 429:
            return True
        if response.status == 400:
            response_json = await response.json(content_type=None)
            if response_json.get("error_code") == ErrorCode.QuotaLimitExceeded.value:
                return True
        return False

    async def __is_access_token_expired(self, response: aiohttp.ClientResponse) -> bool:
        response_json = await response.json(content_type=None)
        return response.status == 401 and (
            response_json.get("error_code") == ErrorCode.AccessTokenExpired.value
            or response_json.get("error_code") == ErrorCode.AccessTokenInvalid.value
        )

    async def base_request(
        self, method: str, url: str, **kwargs
    ) -> aiohttp.ClientResponse:
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

        if await self.__is_ratelimit_error(response):
            raise RateLimitError(response)
        if await self.__is_access_token_expired(response):
            raise AccessTokenExpiredError(response)

        if response.status == 400:
            raise BadRequestError(response)
        if response.status == 401:
            raise AuthenticationError(response)
        if response.status == 403:
            raise ForbiddenError(response)
        if response.status == 404:
            raise NotFoundError(response)
        if response.status >= 500:
            raise InternalServerError(response)
        if response.status and not 200 <= response.status < 300:
            raise HTTPError(response)

        return response


class Client(
    BaseClient,
    # ws.WebSocketInteractor
):
    """yaylib のエントリーポイント"""

    def __init__(
        self,
        *,
        # intents: Optional[ws.Intents] = None,
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
        super().__init__(
            proxy_url=proxy_url, timeout=timeout, base_path=base_path, loglevel=loglevel
        )

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

        self.__state = state or State(storage_path=base_path + "secret.db")
        self.__header_manager = HeaderManager(Device.instance(), self.__state)

        self.__ratelimit = RateLimit(wait_on_ratelimit, max_ratelimit_retries)

    @property
    def state(self) -> State:
        return self.__state

    @property
    def user_id(self) -> int:
        return self.__state.user_id

    @property
    def access_token(self) -> str:
        return self.__state.access_token

    @property
    def refresh_token(self) -> str:
        return self.__state.refresh_token

    @property
    def device_uuid(self) -> str:
        return self.__state.device_uuid

    def __construct_response(
        self, response: dict, data_type: Optional[Model] = None
    ) -> dict | Model:
        if data_type is not None:
            if isinstance(response, list):
                response = [data_type(result) for result in response]
            elif response is not None:
                response = data_type(response)
        return response

    async def __make_request(
        self,
        method: str,
        url: str,
        *,
        params: Optional[dict] = None,
        json: Optional[dict] = None,
        headers: Optional[dict] = None,
        return_type: Optional[Model] = None,
    ) -> dict | Model:
        response = await self.base_request(
            method, url, params=params, json=json, headers=headers
        )
        response_json = await response.json(content_type=None)
        return self.__construct_response(response_json, return_type)

    async def __refresh_client_tokens(self) -> None:
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
    ) -> dict | Model:
        if not url.startswith("https://"):
            url = "https://" + url

        if not self.__header_manager.client_ip and "v2/users/timestamp" not in url:
            metadata = await self.user.get_timestamp()
            self.__header_manager.client_ip = metadata.ip_address

        backoff_duration = 0
        headers = headers or {}

        for i in range(max(1, self.__max_retries + 1)):
            try:
                await asyncio.sleep(backoff_duration)
                while True:
                    try:
                        headers.update(self.__header_manager.generate(jwt_required))
                        response = await self.__make_request(
                            method,
                            url,
                            params=utils.filter_dict(params),
                            json=utils.filter_dict(json),
                            headers=headers,
                            return_type=return_type,
                        )
                        break
                    except RateLimitError as exc:
                        self.logger.warning(
                            "Rate limit exceeded. Waiting... (%s/%s)",
                            self.__ratelimit.retries_performed,
                            self.__ratelimit.max_retries,
                        )
                        await self.__ratelimit.wait(exc)
                self.__ratelimit.reset()
                break
            except AccessTokenExpiredError as exc:
                if self.user_id == 0:
                    self.logger.error("Authentication required to perform the action.")
                    raise AuthenticationError(exc.response) from exc
                await self.__refresh_client_tokens()
            except AuthenticationError as exc:
                if "/api/v1/oauth/token" in url:
                    self.__state.destory(self.user_id)
                    self.logger.error(
                        "Failed to refresh credentials. Please try logging in again."
                    )
                    raise exc
            except InternalServerError as exc:
                self.logger.error(
                    "Request failed with status %s! Retrying...", exc.response.status
                )
                backoff_duration = self.__backoff_factor * (2**i)

        return response

    def insert_delay(self, callback: Callable):
        """リクエスト間の時間が1秒未満のときに遅延を挿入します"""
        if int(datetime.now().timestamp()) - self.__last_request_ts < 1:
            time.sleep(random.uniform(self.__min_delay, self.__max_delay))
        self.__last_request_ts = int(datetime.now().timestamp())
        return callback()

    def __sync_request(self, callback: Awaitable):
        """非同期リクエストを同期リクエストに変換します"""
        return self.insert_delay(lambda: asyncio.run(callback))

    # ---------- call api ----------

    def get_user_active_call(self, user_id: int) -> PostResponse:
        """ユーザーが参加中の通話を取得します"""
        return self.__sync_request(self.call.get_user_active_call(user_id))

    def get_bgms(self) -> BgmsResponse:
        """通話のBGMを取得します"""
        return self.__sync_request(self.call.get_bgms())

    def get_call(self, call_id: int) -> ConferenceCallResponse:
        """通話の詳細を取得します"""
        return self.__sync_request(self.call.get_call(call_id))

    def get_call_invitable_users(
        self, call_id: int, from_timestamp: Optional[int] = None
    ) -> UsersByTimestampResponse:
        """通話に招待可能なユーザーを取得します"""
        return self.__sync_request(
            self.call.get_call_invitable_users(call_id, from_timestamp)
        )

    def get_call_status(self, opponent_id: int) -> CallStatusResponse:
        """通話の状態を取得します"""
        return self.__sync_request(self.call.get_call_status(opponent_id))

    def get_games(self, **params) -> GamesResponse:
        """ゲームを取得します"""
        return self.__sync_request(self.call.get_games(**params))

    def get_genres(self, **params) -> GenresResponse:
        """通話のジャンルを取得します"""
        return self.__sync_request(self.call.get_genres(**params))

    def get_group_calls(self, **params) -> PostsResponse:
        """サークルの通話を取得します"""
        return self.__sync_request(self.call.get_group_calls(**params))

    def invite_online_followings_to_call(
        self,
        call_id: int,
        group_id: Optional[int] = None,
    ) -> dict:
        """オンラインの友達をまとめて通話に招待します"""
        return self.__sync_request(
            self.call.invite_online_followings_to_call(call_id, group_id)
        )

    def invite_users_to_call(self, call_id: int, user_ids: list[int]) -> dict:
        """ユーザーを通話に招待します"""
        return self.__sync_request(self.call.invite_users_to_call(call_id, user_ids))

    def invite_users_to_chat_call(
        self,
        chat_room_id: int,
        room_id: int,
        room_url: str,
    ) -> dict:
        """チャット通話にユーザーを招待します"""
        return self.__sync_request(
            self.call.invite_users_to_chat_call(chat_room_id, room_id, room_url)
        )

    def kick_user_from_call(self, call_id: int, user_id: int) -> dict:
        """ユーザーを通話からキックします"""
        self.__sync_request(self.call.kick_user_from_call(call_id, user_id))

    def set_call(
        self,
        call_id: int,
        joinable_by: str,
        game_title: Optional[str] = None,
        category_id: Optional[str] = None,
    ) -> dict:
        """通話を開始します"""
        return self.__sync_request(
            self.call.set_call(call_id, joinable_by, game_title, category_id)
        )

    def set_user_role(self, call_id: int, user_id: int, role: str) -> dict:
        """通話に参加中のユーザーに役職を与えます"""
        return self.__sync_request(self.call.set_user_role(call_id, user_id, role))

    def join_call(
        self, conference_id: int, call_sid: Optional[str] = None
    ) -> ConferenceCallResponse:
        """通話に参加します"""
        return self.__sync_request(self.call.join_call(conference_id, call_sid))

    def join_call_as_anonymous(
        self, conference_id: int, agora_uid: str
    ) -> ConferenceCallResponse:
        """無名くんとして通話に参加します"""
        return self.__sync_request(
            self.call.join_call_as_anonymous(conference_id, agora_uid)
        )

    def leave_call(
        self,
        conference_id: int,
        call_sid: Optional[str] = None,
    ) -> dict:
        """通話から退出します"""
        return self.__sync_request(self.call.leave_call(conference_id, call_sid))

    def leave_call_as_anonymous(
        self, conference_id: int, agora_uid: Optional[str] = None
    ):
        """通話から退出します"""
        return self.__sync_request(
            self.call.leave_call_as_anonymous(conference_id, agora_uid)
        )

    # ---------- notification api ----------

    def get_activities(self, **params) -> ActivitiesResponse:
        """通知を取得します"""
        return self.__sync_request(self.notification.get_activities(**params))

    def get_merged_activities(self, **params) -> ActivitiesResponse:
        """全種類の通知を取得します"""
        return self.__sync_request(self.notification.get_merged_activities(**params))

    # ---------- chat api ----------

    def accept_chat_requests(self, chat_room_ids: list[int]) -> dict:
        """チャットリクエストを承認します"""
        return self.__sync_request(self.chat.accept_chat_requests(chat_room_ids))

    def check_unread_status(self, from_time: int) -> UnreadStatusResponse:
        """チャットの未読ステータスを確認します"""
        return self.__sync_request(self.chat.check_unread_status(from_time))

    def create_group_chat(
        self,
        name: str,
        with_user_ids: list[int],
        icon_filename: Optional[str] = None,
        background_filename: Optional[str] = None,
    ) -> CreateChatRoomResponse:
        """グループチャットを作成します"""
        return self.__sync_request(
            self.chat.create_group_chat(
                name, with_user_ids, icon_filename, background_filename
            )
        )

    def create_private_chat(
        self,
        with_user_id: int,
        matching_id: Optional[int] = None,
        hima_chat: Optional[bool] = False,
    ) -> CreateChatRoomResponse:
        """個人チャットを作成します"""
        return self.__sync_request(
            self.chat.create_private_chat(with_user_id, matching_id, hima_chat)
        )

    def delete_background(self, room_id: int) -> dict:
        """チャットの背景画像を削除します"""
        return self.__sync_request(self.chat.delete_background(room_id))

    def delete_message(self, room_id: int, message_id: int) -> dict:
        """メッセージを削除します"""
        return self.__sync_request(self.chat.delete_message(room_id, message_id))

    def edit_chat_room(
        self,
        chat_room_id: int,
        name: str,
        icon_filename: Optional[str] = None,
        background_filename: Optional[str] = None,
    ) -> dict:
        """チャットルームを編集します"""
        return self.__sync_request(
            self.chat.edit_chat_room(
                chat_room_id, name, icon_filename, background_filename
            )
        )

    def get_chatable_users(
        self,
        from_follow_id: Optional[int] = None,
        from_timestamp: Optional[int] = None,
        order_by: Optional[str] = None,
    ) -> FollowUsersResponse:
        """チャット可能なユーザーを取得します"""
        return self.__sync_request(
            self.chat.get_chatable_users(from_follow_id, from_timestamp, order_by)
        )

    def get_gifs_data(self) -> GifsDataResponse:
        """チャットルームのGIFデータを取得します"""
        return self.__sync_request(self.chat.get_gifs_data())

    def get_hidden_chat_rooms(self, **params) -> ChatRoomsResponse:
        """非表示のチャットルームを取得します"""
        return self.__sync_request(self.chat.get_hidden_chat_rooms(**params))

    def get_main_chat_rooms(
        self, from_timestamp: Optional[int] = None
    ) -> ChatRoomsResponse:
        """メインのチャットルームを取得します"""
        return self.__sync_request(self.chat.get_main_chat_rooms(from_timestamp))

    def get_messages(self, chat_room_id: int, **params) -> MessagesResponse:
        """メッセージを取得します"""
        return self.__sync_request(self.chat.get_messages(chat_room_id, **params))

    def get_chat_requests(self, **params) -> ChatRoomsResponse:
        """チャットリクエストを取得します"""
        return self.__sync_request(self.chat.get_request_chat_rooms(**params))

    def get_chat_room(self, chat_room_id: int) -> ChatRoomResponse:
        """チャットルームの詳細を取得します"""
        return self.__sync_request(self.chat.get_chat_room(chat_room_id))

    def get_sticker_packs(self) -> StickerPacksResponse:
        """チャットスタンプを取得します"""
        return self.__sync_request(self.chat.get_sticker_packs())

    def get_total_chat_requests(self) -> TotalChatRequestResponse:
        """チャットリクエストの数を取得します"""
        return self.__sync_request(self.chat.get_total_chat_requests())

    def hide_chat(self, chat_room_id: int) -> dict:
        """チャットルームを非表示にします"""
        return self.__sync_request(self.chat.hide_chat(chat_room_id))

    def invite_to_chat(self, chat_room_id: int, user_ids: list[int]) -> dict:
        """チャットルームにユーザーを招待します"""
        return self.__sync_request(self.chat.invite_to_chat(chat_room_id, user_ids))

    def kick_users_from_chat(self, chat_room_id: int, user_ids: list[int]) -> dict:
        """チャットルームからユーザーを追放します"""
        return self.__sync_request(
            self.chat.kick_users_from_chat(chat_room_id, user_ids)
        )

    def pin_chat(self, room_id: int) -> dict:
        """チャットルームをピン止めします"""
        return self.__sync_request(self.chat.pin_chat(room_id))

    def read_message(self, chat_room_id: int, message_id: int) -> dict:
        """メッセージを既読にします"""
        return self.__sync_request(self.chat.read_message(chat_room_id, message_id))

    def refresh_chat_rooms(self, from_time: Optional[int] = None) -> ChatRoomsResponse:
        """チャットルームを更新します"""
        return self.__sync_request(self.chat.refresh_chat_rooms(from_time))

    def delete_chat_rooms(self, chat_room_ids: list[int]) -> dict:
        """チャットルームを削除します"""
        return self.__sync_request(self.chat.delete_chat_rooms(chat_room_ids))

    def report_chat_room(
        self,
        chat_room_id: int,
        opponent_id: int,
        category_id: int,
        reason: Optional[str] = None,
        screenshot_filename: Optional[str] = None,
        screenshot_2_filename: Optional[str] = None,
        screenshot_3_filename: Optional[str] = None,
        screenshot_4_filename: Optional[str] = None,
    ) -> dict:
        """チャットルームを通報します"""
        return self.__sync_request(
            self.chat.report_chat_room(
                chat_room_id,
                opponent_id,
                category_id,
                reason,
                screenshot_filename,
                screenshot_2_filename,
                screenshot_3_filename,
                screenshot_4_filename,
            )
        )

    def send_message(
        self,
        chat_room_id: int,
        **params,
    ) -> MessageResponse:
        """チャットルームにメッセージを送信します"""
        return self.__sync_request(self.chat.send_message(chat_room_id, **params))

    def unhide_chat(self, chat_room_ids: int) -> dict:
        """非表示に設定したチャットルームを表示する"""
        return self.__sync_request(self.chat.unhide_chat(chat_room_ids))

    def unpin_chat(self, chat_room_id: int) -> dict:
        """チャットルームのピン止めを解除します"""
        return self.__sync_request(self.chat.unpin_chat(chat_room_id))

    # ---------- group api ----------

    def accept_moderator_offer(self, group_id: int) -> dict:
        """サークル副管理人の権限オファーを引き受けます"""
        return self.__sync_request(self.group.accept_moderator_offer(group_id))

    def accept_ownership_offer(self, group_id: int) -> dict:
        """サークル管理人の権限オファーを引き受けます"""
        return self.__sync_request(self.group.accept_ownership_offer(group_id))

    def accept_group_join_request(self, group_id: int, user_id: int) -> dict:
        """サークル参加リクエストを承認します"""
        return self.__sync_request(
            self.group.accept_group_join_request(group_id, user_id)
        )

    def add_related_groups(self, group_id: int, related_group_id: list[int]) -> dict:
        """関連サークルを追加します"""
        return self.__sync_request(
            self.group.add_related_groups(group_id, related_group_id)
        )

    def ban_group_user(self, group_id: int, user_id: int) -> dict:
        """サークルからユーザーを追放します"""
        return self.__sync_request(self.group.ban_group_user(group_id, user_id))

    def check_group_unread_status(
        self, from_time: Optional[int] = None
    ) -> UnreadStatusResponse:
        """サークルの未読ステータスを取得します"""
        return self.__sync_request(self.group.check_group_unread_status(from_time))

    def create_group(
        self,
        topic: str,
        description: Optional[str] = None,
        secret: Optional[bool] = None,
        hide_reported_posts: Optional[bool] = None,
        hide_conference_call: Optional[bool] = None,
        is_private: Optional[bool] = None,
        only_verified_age: Optional[bool] = None,
        only_mobile_verified: Optional[bool] = None,
        call_timeline_display: Optional[bool] = None,
        allow_ownership_transfer: Optional[bool] = None,
        allow_thread_creation_by: Optional[bool] = None,
        gender: Optional[int] = None,
        generation_groups_limit: Optional[int] = None,
        group_category_id: Optional[int] = None,
        cover_image_filename: Optional[str] = None,
        sub_category_id: Optional[str] = None,
        hide_from_game_eight: Optional[bool] = None,
        allow_members_to_post_media: Optional[bool] = None,
        allow_members_to_post_url: Optional[bool] = None,
        guidelines: Optional[str] = None,
    ) -> CreateGroupResponse:
        """サークルを作成します"""
        return self.__sync_request(
            self.group.create_group(
                topic,
                description,
                secret,
                hide_reported_posts,
                hide_conference_call,
                is_private,
                only_verified_age,
                only_mobile_verified,
                call_timeline_display,
                allow_ownership_transfer,
                allow_thread_creation_by,
                gender,
                generation_groups_limit,
                group_category_id,
                cover_image_filename,
                sub_category_id,
                hide_from_game_eight,
                allow_members_to_post_media,
                allow_members_to_post_url,
                guidelines,
            )
        )

    def pin_group(self, group_id: int) -> dict:
        """サークルをピンします"""
        return self.__sync_request(self.group.pin_group(group_id))

    def decline_moderator_offer(self, group_id: int) -> dict:
        """サークル副管理人の権限オファーを断ります"""
        return self.__sync_request(self.group.decline_moderator_offer(group_id))

    def decline_ownership_offer(self, group_id: int) -> dict:
        """サークル管理人の権限オファーを断ります"""
        return self.__sync_request(self.group.decline_ownership_offer(group_id))

    def decline_group_join_request(self, group_id: int, user_id: int) -> dict:
        """サークル参加リクエストを断ります"""
        return self.__sync_request(
            self.group.decline_group_join_request(group_id, user_id)
        )

    def unpin_group(self, group_id: int) -> dict:
        """サークルのピン止めを解除します"""
        return self.__sync_request(self.group.unpin_group(group_id))

    def get_banned_group_members(
        self,
        group_id: int,
        page: Optional[int] = None,
    ) -> UsersResponse:
        """追放されたサークルメンバーを取得します"""
        return self.__sync_request(self.group.get_banned_group_members(group_id, page))

    def get_group_categories(self, **params) -> GroupCategoriesResponse:
        """サークルのカテゴリーを取得します"""
        return self.__sync_request(self.group.get_group_categories(**params))

    def get_create_group_quota(self) -> CreateGroupQuota:
        """残りのサークル作成可能回数を取得します"""
        return self.__sync_request(self.group.get_create_group_quota())

    def get_group(self, group_id: int) -> GroupResponse:
        """サークルの詳細を取得します"""
        return self.__sync_request(self.group.get_group(group_id))

    def get_groups(self, **params) -> GroupsResponse:
        """複数のサークルの詳細を取得します"""
        return self.__sync_request(self.group.get_groups(**params))

    def get_invitable_users(self, group_id: int, **params) -> UsersByTimestampResponse:
        """サークルに招待可能なユーザーを取得します"""
        return self.__sync_request(self.group.get_invitable_users(group_id, **params))

    def get_joined_statuses(self, ids: list[int]) -> dict:
        """サークルの参加ステータスを取得します"""
        return self.__sync_request(self.group.get_joined_statuses(ids))

    def get_group_member(self, group_id: int, user_id: int) -> GroupUserResponse:
        """特定のサークルメンバーの情報を取得します"""
        return self.__sync_request(self.group.get_group_member(group_id, user_id))

    def get_group_members(self, group_id: int, **params) -> GroupUsersResponse:
        """サークルメンバーを取得します"""
        return self.__sync_request(self.group.get_group_members(group_id, **params))

    def get_my_groups(self, from_timestamp: Optional[int] = None) -> GroupsResponse:
        """自分のサークルを取得します"""
        return self.__sync_request(self.group.get_my_groups(from_timestamp))

    def get_relatable_groups(self, group_id: int, **params) -> GroupsRelatedResponse:
        """関連がある可能性があるサークルを取得します"""
        return self.__sync_request(self.group.get_relatable_groups(group_id, **params))

    def get_related_groups(self, group_id: int, **params) -> GroupsRelatedResponse:
        """関連があるサークルを取得します"""
        return self.__sync_request(self.group.get_related_groups(group_id, **params))

    def get_user_groups(self, **params) -> GroupsResponse:
        """特定のユーザーが参加しているサークルを取得します"""
        return self.__sync_request(self.group.get_user_groups(**params))

    def invite_users_to_group(self, group_id: int, user_ids: list[int]) -> dict:
        """サークルにユーザーを招待します"""
        return self.__sync_request(self.group.invite_users_to_group(group_id, user_ids))

    def join_group(self, group_id: int) -> dict:
        """サークルに参加します"""
        return self.__sync_request(self.group.join_group(group_id))

    def leave_group(self, group_id: int) -> dict:
        """サークルから脱退します"""
        return self.__sync_request(self.group.leave_group(group_id))

    def remove_group_cover(self, group_id: int) -> dict:
        """サークルのカバー画像を削除します"""
        return self.__sync_request(self.group.remove_group_cover(group_id))

    def remove_moderator(self, group_id: int, user_id: int) -> dict:
        """サークルの副管理人を削除します"""
        return self.__sync_request(self.group.remove_moderator(group_id, user_id))

    def remove_related_groups(
        self, group_id: int, related_group_ids: list[int]
    ) -> dict:
        """関連のあるサークルを削除します"""
        return self.__sync_request(
            self.group.remove_related_groups(group_id, related_group_ids)
        )

    def send_moderator_offers(self, group_id: int, user_ids: list[int]) -> dict:
        """複数人にサークル副管理人のオファーを送信します"""
        return self.__sync_request(self.group.send_moderator_offers(group_id, user_ids))

    def send_ownership_offer(self, group_id: int, user_id: int) -> dict:
        """サークル管理人権限のオファーを送信します"""
        return self.__sync_request(self.group.send_ownership_offer(group_id, user_id))

    def set_group_title(self, group_id: int, title: str) -> dict:
        """サークルのタイトルを設定します"""
        return self.__sync_request(self.group.set_group_title(group_id, title))

    def take_over_group_ownership(self, group_id: int) -> dict:
        """サークル管理人の権限を引き継ぎます"""
        return self.__sync_request(self.group.take_over_group_ownership(group_id))

    def unban_group_member(self, group_id: int, user_id: int) -> dict:
        """特定のサークルメンバーの追放を解除します"""
        return self.__sync_request(self.group.unban_group_member(group_id, user_id))

    def update_group(
        self,
        group_id: int,
        topic: str,
        description: Optional[str] = None,
        secret: Optional[bool] = None,
        hide_reported_posts: Optional[bool] = None,
        hide_conference_call: Optional[bool] = None,
        is_private: Optional[bool] = None,
        only_verified_age: Optional[bool] = None,
        only_mobile_verified: Optional[bool] = None,
        call_timeline_display: Optional[bool] = None,
        allow_ownership_transfer: Optional[bool] = None,
        allow_thread_creation_by: Optional[str] = None,
        gender: Optional[int] = None,
        generation_groups_limit: Optional[int] = None,
        group_category_id: Optional[int] = None,
        cover_image_filename: Optional[str] = None,
        sub_category_id: Optional[str] = None,
        hide_from_game_eight: Optional[bool] = None,
        allow_members_to_post_media: Optional[bool] = None,
        allow_members_to_post_url: Optional[bool] = None,
        guidelines: Optional[str] = None,
    ) -> GroupResponse:
        """サークルを編集します"""
        return self.__sync_request(
            self.group.update_group(
                group_id,
                topic,
                description,
                secret,
                hide_reported_posts,
                hide_conference_call,
                is_private,
                only_verified_age,
                only_mobile_verified,
                call_timeline_display,
                allow_ownership_transfer,
                allow_thread_creation_by,
                gender,
                generation_groups_limit,
                group_category_id,
                cover_image_filename,
                sub_category_id,
                hide_from_game_eight,
                allow_members_to_post_media,
                allow_members_to_post_url,
                guidelines,
            )
        )

    def withdraw_moderator_offer(self, group_id: int, user_id: int) -> dict:
        """サークル副管理人のオファーを取り消します"""
        return self.__sync_request(
            self.group.withdraw_moderator_offer(group_id, user_id)
        )

    def withdraw_ownership_offer(self, group_id: int, user_id: int) -> dict:
        """サークル管理人のオファーを取り消します"""
        return self.__sync_request(
            self.group.withdraw_ownership_offer(group_id, user_id)
        )

    # ---------- auth api ----------

    def change_password(
        self, current_password: str, new_password: str
    ) -> LoginUpdateResponse:
        """パスワードを変更します"""
        return self.__sync_request(
            self.auth.change_password(current_password, new_password)
        )

    def get_token(
        self,
        grant_type: str,
        refresh_token: Optional[str] = None,
        email: Optional[str] = None,
        password: Optional[str] = None,
    ) -> TokenResponse:
        """トークンを再発行します"""
        return self.__sync_request(
            self.auth.get_token(grant_type, refresh_token, email, password)
        )

    def login(
        self, email: str, password: str, two_fa_code: Optional[str] = None
    ) -> LoginUserResponse:
        """メールアドレスでログインします"""
        return self.__sync_request(self.auth.login(email, password, two_fa_code))

    def resend_confirm_email(self) -> dict:
        """確認メールを再送信します"""
        return self.__sync_request(self.auth.resend_confirm_email())

    def restore_user(self, user_id: int) -> LoginUserResponse:
        """ユーザーを復元します"""
        return self.__sync_request(self.auth.restore_user(user_id))

    def save_account_with_email(
        self,
        email: str,
        password: Optional[str] = None,
        current_password: Optional[str] = None,
        email_grant_token: Optional[str] = None,
    ) -> LoginUpdateResponse:
        """メールアドレスでアカウントを保存します"""
        return self.__sync_request(
            self.auth.save_account_with_email(
                email, password, current_password, email_grant_token
            )
        )

    # ---------- misc api ----------

    def accept_policy_agreement(self, agreement_type: str) -> dict:
        """利用規約、ポリシー同意書に同意します"""
        return self.__sync_request(self.misc.accept_policy_agreement(agreement_type))

    def send_verification_code(self, email: str, intent: str, locale="ja") -> Dict:
        """認証コードを送信します"""
        return self.__sync_request(
            self.misc.send_verification_code(email, intent, locale)
        )

    def get_email_grant_token(self, code: int, email: str) -> EmailGrantTokenResponse:
        """email_grant_tokenを取得します"""
        return self.__sync_request(self.misc.get_email_grant_token(code, email))

    def get_email_verification_presigned_url(
        self, email: str, locale: str, intent: Optional[str] = None
    ) -> EmailVerificationPresignedUrlResponse:
        """メールアドレス確認用の署名付きURLを取得します"""
        return self.__sync_request(
            self.misc.get_email_verification_presigned_url(email, locale, intent)
        )

    def get_file_upload_presigned_urls(
        self, file_names: list[str]
    ) -> PresignedUrlsResponse:
        """ファイルアップロード用の署名付きURLを取得します"""
        return self.__sync_request(self.misc.get_file_upload_presigned_urls(file_names))

    def get_id_checker_presigned_url(
        self, model: str, action: str, **params
    ) -> IdCheckerPresignedUrlResponse:
        return self.__sync_request(
            self.misc.get_id_checker_presigned_url(model, action, **params)
        )

    def get_old_file_upload_presigned_url(
        self, video_file_name: str
    ) -> PresignedUrlResponse:
        """動画ファイルアップロード用の署名付きURLを取得します"""
        return self.__sync_request(
            self.misc.get_old_file_upload_presigned_url(video_file_name)
        )

    def get_policy_agreements(self) -> PolicyAgreementsResponse:
        """利用規約、ポリシー同意書に同意しているかどうかを取得します"""
        return self.__sync_request(self.misc.get_policy_agreements())

    def get_web_socket_token(self) -> WebSocketTokenResponse:
        """Web Socket トークンを取得します"""
        return self.__sync_request(self.misc.get_web_socket_token())

    def upload_image(self, image_paths: list[str], image_type: str) -> list[Attachment]:
        """画像をアップロードして、サーバー上のファイルのリストを返します。"""
        return self.__sync_request(self.misc.upload_image(image_paths, image_type))

    def upload_video(self, video_path: str) -> str:
        """動画をアップロードして、サーバー上のファイル名を返します。"""
        return self.__sync_request(self.misc.upload_video(video_path))

    def get_app_config(self) -> ApplicationConfigResponse:
        """アプリケーションの設定情報を取得します"""
        return self.__sync_request(self.misc.get_app_config())

    def get_banned_words(self, country_code: str = "jp") -> BanWordsResponse:
        """禁止ワードの一覧を取得します"""
        return self.__sync_request(self.misc.get_banned_words(country_code))

    def get_popular_words(self, country_code: str = "jp") -> PopularWordsResponse:
        """人気ワードの一覧を取得します"""
        return self.__sync_request(self.misc.get_popular_words(country_code))

    # ---------- post api ----------

    def add_bookmark(self, user_id: int, post_id: int) -> BookmarkPostResponse:
        """ブックマークに追加します"""
        return self.__sync_request(self.post.add_bookmark(user_id, post_id))

    def add_group_highlight_post(self, group_id: int, post_id: int) -> dict:
        """投稿をグループのまとめに追加します"""
        return self.__sync_request(
            self.post.add_group_highlight_post(group_id, post_id)
        )

    def create_call_post(
        self,
        text: Optional[str] = None,
        font_size: Optional[int] = None,
        color: Optional[int] = None,
        group_id: Optional[int] = None,
        call_type: Optional[str] = None,
        category_id: Optional[int] = None,
        game_title: Optional[str] = None,
        joinable_by: Optional[str] = None,
        message_tags: Optional[list] = None,
        attachment_filename: Optional[str] = None,
        attachment_2_filename: Optional[str] = None,
        attachment_3_filename: Optional[str] = None,
        attachment_4_filename: Optional[str] = None,
        attachment_5_filename: Optional[str] = None,
        attachment_6_filename: Optional[str] = None,
        attachment_7_filename: Optional[str] = None,
        attachment_8_filename: Optional[str] = None,
        attachment_9_filename: Optional[str] = None,
    ) -> CreatePostResponse:
        """通話の投稿を作成します"""
        return self.__sync_request(
            self.post.create_call_post(
                text,
                font_size,
                color,
                group_id,
                call_type,
                category_id,
                game_title,
                joinable_by,
                message_tags,
                attachment_filename,
                attachment_2_filename,
                attachment_3_filename,
                attachment_4_filename,
                attachment_5_filename,
                attachment_6_filename,
                attachment_7_filename,
                attachment_8_filename,
                attachment_9_filename,
            )
        )

    def pin_group_post(self, post_id: int, group_id: int) -> dict:
        """グループの投稿をピンします"""
        return self.__sync_request(self.post.create_group_pin_post(post_id, group_id))

    def pin_post(self, post_id: int) -> dict:
        """投稿をピンします"""
        return self.__sync_request(self.post.create_pin_post(post_id))

    def create_post(
        self,
        text: Optional[str] = None,
        font_size: Optional[int] = 0,
        color: Optional[int] = 0,
        in_reply_to: Optional[int] = None,
        group_id: Optional[int] = None,
        mention_ids: Optional[list[int]] = None,
        choices: Optional[list[str]] = None,
        shared_url: Optional[str] = None,
        message_tags: Optional[list] = None,
        attachment_filename: Optional[str] = None,
        attachment_2_filename: Optional[str] = None,
        attachment_3_filename: Optional[str] = None,
        attachment_4_filename: Optional[str] = None,
        attachment_5_filename: Optional[str] = None,
        attachment_6_filename: Optional[str] = None,
        attachment_7_filename: Optional[str] = None,
        attachment_8_filename: Optional[str] = None,
        attachment_9_filename: Optional[str] = None,
        video_file_name: Optional[str] = None,
    ) -> Post:
        """投稿を作成します"""
        return self.__sync_request(
            self.post.create_post(
                text,
                font_size,
                color,
                in_reply_to,
                group_id,
                mention_ids,
                choices,
                shared_url,
                message_tags,
                attachment_filename,
                attachment_2_filename,
                attachment_3_filename,
                attachment_4_filename,
                attachment_5_filename,
                attachment_6_filename,
                attachment_7_filename,
                attachment_8_filename,
                attachment_9_filename,
                video_file_name,
            )
        )

    def create_repost(
        self,
        post_id: int,
        text: Optional[str] = None,
        font_size: Optional[int] = 0,
        color: Optional[int] = 0,
        in_reply_to: Optional[int] = None,
        group_id: Optional[int] = None,
        mention_ids: Optional[list[int]] = None,
        choices: Optional[list[str]] = None,
        shared_url: Optional[str] = None,
        message_tags: Optional[list] = None,
        attachment_filename: Optional[str] = None,
        attachment_2_filename: Optional[str] = None,
        attachment_3_filename: Optional[str] = None,
        attachment_4_filename: Optional[str] = None,
        attachment_5_filename: Optional[str] = None,
        attachment_6_filename: Optional[str] = None,
        attachment_7_filename: Optional[str] = None,
        attachment_8_filename: Optional[str] = None,
        attachment_9_filename: Optional[str] = None,
        video_file_name: Optional[str] = None,
    ) -> CreatePostResponse:
        """投稿を(´∀｀∩)↑age↑します"""
        return self.__sync_request(
            self.post.create_repost(
                post_id,
                text,
                font_size,
                color,
                in_reply_to,
                group_id,
                mention_ids,
                choices,
                shared_url,
                message_tags,
                attachment_filename,
                attachment_2_filename,
                attachment_3_filename,
                attachment_4_filename,
                attachment_5_filename,
                attachment_6_filename,
                attachment_7_filename,
                attachment_8_filename,
                attachment_9_filename,
                video_file_name,
            )
        )

    def create_share_post(
        self,
        shareable_type: str,
        shareable_id: int,
        text: Optional[str] = None,
        font_size: Optional[int] = None,
        color: Optional[int] = None,
        group_id: Optional[int] = None,
    ) -> Post:
        """シェア投稿を作成します"""
        return self.__sync_request(
            self.post.create_share_post(
                shareable_type, shareable_id, text, font_size, color, group_id
            )
        )

    def create_thread_post(
        self,
        post_id: int,
        text: Optional[str] = None,
        font_size: Optional[int] = 0,
        color: Optional[int] = 0,
        in_reply_to: Optional[int] = None,
        group_id: Optional[int] = None,
        mention_ids: Optional[list[int]] = None,
        choices: Optional[list[str]] = None,
        shared_url: Optional[str] = None,
        message_tags: Optional[list] = None,
        attachment_filename: Optional[str] = None,
        attachment_2_filename: Optional[str] = None,
        attachment_3_filename: Optional[str] = None,
        attachment_4_filename: Optional[str] = None,
        attachment_5_filename: Optional[str] = None,
        attachment_6_filename: Optional[str] = None,
        attachment_7_filename: Optional[str] = None,
        attachment_8_filename: Optional[str] = None,
        attachment_9_filename: Optional[str] = None,
        video_file_name: Optional[str] = None,
    ) -> Post:
        """スレッドの投稿を作成します"""
        return self.__sync_request(
            self.post.create_thread_post(
                post_id,
                text,
                font_size,
                color,
                in_reply_to,
                group_id,
                mention_ids,
                choices,
                shared_url,
                message_tags,
                attachment_filename,
                attachment_2_filename,
                attachment_3_filename,
                attachment_4_filename,
                attachment_5_filename,
                attachment_6_filename,
                attachment_7_filename,
                attachment_8_filename,
                attachment_9_filename,
                video_file_name,
            )
        )

    def delete_all_posts(self) -> dict:
        """すべての自分の投稿を削除します"""
        return self.__sync_request(self.post.delete_all_posts())

    def unpin_group_post(self, group_id: int) -> dict:
        """グループのピン投稿を解除します"""
        return self.__sync_request(self.post.unpin_group_post(group_id))

    def unpin_post(self, post_id: int) -> dict:
        """投稿のピンを解除します"""
        return self.__sync_request(self.post.unpin_post(post_id))

    def get_bookmark(
        self,
        user_id: int,
        from_str: Optional[str] = None,
    ) -> PostsResponse:
        """ブックマークを取得します"""
        return self.__sync_request(self.post.get_bookmark(user_id, from_str))

    def get_timeline_calls(self, **params) -> PostsResponse:
        """誰でも通話を取得します"""
        return self.__sync_request(self.post.get_timeline_calls(**params))

    def get_conversation(self, conversation_id: int, **params) -> PostsResponse:
        """リプライを含める投稿の会話を取得します"""
        return self.__sync_request(
            self.post.get_conversation(conversation_id, **params)
        )

    def get_conversation_root_posts(self, post_ids: list[int]) -> PostsResponse:
        """会話の原点の投稿を取得します"""
        return self.__sync_request(self.post.get_conversation_root_posts(post_ids))

    def get_following_call_timeline(self, **params) -> PostsResponse:
        """フォロー中の通話を取得します"""
        return self.__sync_request(self.post.get_following_call_timeline(**params))

    def get_following_timeline(self, **params) -> PostsResponse:
        """フォロー中のタイムラインを取得します"""
        return self.__sync_request(self.post.get_following_timeline(**params))

    def get_group_highlight_posts(self, group_id: int, **params) -> PostsResponse:
        """グループのハイライト投稿を取得します"""
        return self.__sync_request(
            self.post.get_group_highlight_posts(group_id, **params)
        )

    def get_group_timeline_by_keyword(
        self, group_id: int, keyword: str, **params
    ) -> PostsResponse:
        """グループの投稿をキーワードで検索します"""
        return self.__sync_request(
            self.post.get_group_timeline_by_keyword(group_id, keyword, **params)
        )

    def get_group_timeline(self, group_id: int, **params) -> PostsResponse:
        """グループのタイムラインを取得します"""
        return self.__sync_request(self.post.get_group_timeline(group_id, **params))

    def get_timeline_by_hashtag(self, hashtag: str, **params) -> PostsResponse:
        """ハッシュタグでタイムラインを検索します"""
        return self.__sync_request(self.post.get_timeline_by_hashtag(hashtag, **params))

    def get_my_posts(self, **params) -> PostsResponse:
        """自分の投稿を取得します"""
        return self.__sync_request(self.post.get_my_posts(**params))

    def get_post(self, post_id: int) -> PostResponse:
        """投稿の詳細を取得します"""
        return self.__sync_request(self.post.get_post(post_id))

    def get_post_likers(self, post_id: int, **params) -> PostLikersResponse:
        """投稿にいいねしたユーザーを取得します"""
        return self.__sync_request(self.post.get_post_likers(post_id, **params))

    def get_reposts(self, post_id: int, **params: int) -> PostsResponse:
        """投稿の(´∀｀∩)↑age↑を取得します"""
        return self.__sync_request(self.post.get_post_reposts(post_id, **params))

    def get_posts(self, post_ids: list[int]) -> PostsResponse:
        """複数の投稿を取得します"""
        return self.__sync_request(self.post.get_posts(post_ids))

    def get_recommended_post_tags(
        self, tag: Optional[str] = None, save_recent_search: Optional[bool] = False
    ) -> PostTagsResponse:
        """おすすめのタグ候補を取得します"""
        return self.__sync_request(
            self.post.get_recommended_post_tags(tag, save_recent_search)
        )

    def get_recommended_posts(self, **params) -> PostsResponse:
        """おすすめの投稿を取得します"""
        return self.__sync_request(self.post.get_recommended_posts(**params))

    def get_timeline_by_keyword(
        self,
        keyword: Optional[str] = None,
        **params,
    ) -> PostsResponse:
        """キーワードでタイムラインを検索します"""
        return self.__sync_request(self.post.get_timeline_by_keyword(keyword, **params))

    def get_timeline(self, **params) -> PostsResponse:
        """タイムラインを取得します"""
        return self.__sync_request(self.post.get_timeline(**params))

    def get_url_metadata(self, url: str) -> SharedUrl:
        """URLのメタデータを取得します"""
        return self.__sync_request(self.post.get_url_metadata(url))

    def get_user_timeline(self, user_id: int, **params) -> PostsResponse:
        """ユーザーのタイムラインを取得します"""
        return self.__sync_request(self.post.get_user_timeline(user_id, **params))

    def like(self, post_ids: list[int]) -> LikePostsResponse:
        """投稿にいいねします ※ 一度にいいねできる投稿数は最大25個"""
        return self.__sync_request(self.post.like(post_ids))

    def remove_bookmark(self, user_id: int, post_id: int) -> dict:
        """ブックマークを削除します"""
        return self.__sync_request(self.post.remove_bookmark(user_id, post_id))

    def remove_group_highlight_post(self, group_id: int, post_id: int) -> dict:
        """サークルのハイライト投稿を解除します"""
        return self.__sync_request(
            self.post.remove_group_highlight_post(group_id, post_id)
        )

    def remove_posts(self, post_ids: list[int]) -> dict:
        """複数の投稿を削除します"""
        return self.__sync_request(self.post.remove_posts(post_ids))

    def unlike(self, post_id: int) -> dict:
        """投稿のいいねを解除します"""
        return self.__sync_request(self.post.unlike(post_id))

    def update_post(
        self,
        post_id: int,
        text: Optional[str] = None,
        font_size: Optional[int] = None,
        color: Optional[int] = None,
        message_tags: Optional[list] = None,
    ) -> Post:
        """投稿を編集します"""
        return self.__sync_request(
            self.post.update_post(post_id, text, font_size, color, message_tags)
        )

    def view_video(self, video_id: int) -> dict:
        """動画を視聴します"""
        return self.__sync_request(self.post.view_video(video_id))

    def vote_survey(self, survey_id: int, choice_id: int) -> VoteSurveyResponse:
        """アンケートに投票します"""
        return self.__sync_request(self.post.vote_survey(survey_id, choice_id))

    # ---------- review api ----------

    def create_review(self, user_id: int, comment: str) -> dict:
        """レターを送信します"""
        return self.__sync_request(self.review.create_review(user_id, comment))

    def delete_reviews(self, review_ids: list[int]) -> dict:
        """レターを削除します"""
        return self.__sync_request(self.review.delete_reviews(review_ids))

    def get_my_reviews(self, **params) -> ReviewsResponse:
        """送信したレターを取得します"""
        return self.__sync_request(self.review.get_my_reviews(**params))

    def get_reviews(self, user_id: int, **params) -> ReviewsResponse:
        """ユーザーが受け取ったレターを取得します"""
        return self.__sync_request(self.review.get_reviews(user_id, **params))

    def pin_review(self, review_id: int) -> dict:
        """レターをピンします"""
        return self.__sync_request(self.review.pin_review(review_id))

    def unpin_review(self, review_id: int) -> dict:
        """レターのピン止めを解除します"""
        return self.__sync_request(self.review.unpin_review(review_id))

    # ---------- thread api ----------

    def add_post_to_thread(self, post_id: int, thread_id: int) -> ThreadInfo:
        """投稿をスレッドに追加します"""
        return self.__sync_request(self.thread.add_post_to_thread(post_id, thread_id))

    def convert_post_to_thread(
        self,
        post_id: int,
        title: Optional[str] = None,
        thread_icon_filename: Optional[str] = None,
    ) -> ThreadInfo:
        """投稿をスレッドに変換します"""
        return self.__sync_request(
            self.thread.convert_post_to_thread(post_id, title, thread_icon_filename)
        )

    def create_thread(
        self,
        group_id: int,
        title: str,
        thread_icon_filename: str,
    ) -> ThreadInfo:
        """スレッドを作成します"""
        return self.__sync_request(
            self.thread.create_thread(group_id, title, thread_icon_filename)
        )

    def get_group_thread_list(
        self,
        group_id: int,
        from_str: Optional[str] = None,
        **params,
    ) -> GroupThreadListResponse:
        """グループのスレッド一覧を取得します"""
        return self.__sync_request(
            self.thread.get_group_thread_list(group_id, from_str, **params)
        )

    def get_thread_joined_statuses(self, ids: list[int]) -> dict:
        """スレッド参加ステータスを取得します"""
        return self.__sync_request(self.thread.get_thread_joined_statuses(ids))

    def get_thread_posts(
        self,
        thread_id: int,
        from_str: Optional[str] = None,
        **params,
    ) -> PostsResponse:
        """スレッド内のタイムラインを取得します"""
        return self.__sync_request(
            self.thread.get_thread_posts(thread_id, from_str, **params)
        )

    def join_thread(self, thread_id: int, user_id: int) -> dict:
        """スレッドに参加します"""
        return self.__sync_request(self.thread.join_thread(thread_id, user_id))

    def leave_thread(self, thread_id: int, user_id: int) -> dict:
        """スレッドから脱退します"""
        return self.__sync_request(self.thread.leave_thread(thread_id, user_id))

    def remove_thread(self, thread_id: int) -> dict:
        """スレッドを削除します"""
        return self.__sync_request(self.thread.remove_thread(thread_id))

    def update_thread(
        self,
        thread_id: int,
        title: str,
        thread_icon_filename: str,
    ) -> dict:
        """スレッドを編集します"""
        return self.__sync_request(
            self.thread.update_thread(thread_id, title, thread_icon_filename)
        )

    # ---------- user api ----------

    def delete_footprint(self, user_id: int, footprint_id: int) -> dict:
        """足跡を削除します"""
        return self.__sync_request(self.user.delete_footprint(user_id, footprint_id))

    def follow_user(self, user_id: int) -> dict:
        """ユーザーをフォローします"""
        return self.__sync_request(self.user.follow_user(user_id))

    def follow_users(self, user_ids: list[int]) -> dict:
        """複数のユーザーをフォローします"""
        return self.__sync_request(self.user.follow_users(user_ids))

    def get_active_followings(self, **params) -> ActiveFollowingsResponse:
        """アクティブなフォロー中のユーザーを取得します"""
        return self.__sync_request(self.user.get_active_followings(**params))

    def get_follow_recommendations(self, **params) -> FollowRecommendationsResponse:
        """フォローするのにおすすめのユーザーを取得します"""
        return self.__sync_request(self.user.get_follow_recommendations(**params))

    def get_follow_request(
        self, from_timestamp: Optional[int] = None
    ) -> UsersByTimestampResponse:
        """フォローリクエストを取得します"""
        return self.__sync_request(self.user.get_follow_request(from_timestamp))

    def get_follow_request_count(self) -> FollowRequestCountResponse:
        """フォローリクエストの数を取得します"""
        return self.__sync_request(self.user.get_follow_request_count())

    def get_following_users_born(
        self, birthdate: Optional[int] = None
    ) -> UsersResponse:
        """フォロー中のユーザーの誕生日を取得します"""
        return self.__sync_request(self.user.get_following_users_born(birthdate))

    def get_footprints(self, **params) -> FootprintsResponse:
        """足跡を取得します"""
        return self.__sync_request(self.user.get_footprints(**params))

    def get_fresh_user(self, user_id: int) -> UserResponse:
        """認証情報などを含んだユーザー情報を取得します"""
        return self.__sync_request(self.user.get_fresh_user(user_id))

    def get_hima_users(self, **params) -> HimaUsersResponse:
        """暇なユーザーを取得します"""
        return self.__sync_request(self.user.get_hima_users(**params))

    def get_user_ranking(self, mode: str) -> RankingUsersResponse:
        """ユーザーのフォロワーランキングを取得します"""
        return self.__sync_request(self.user.get_user_ranking(mode))

    def get_profile_refresh_counter_requests(self) -> RefreshCounterRequestsResponse:
        """投稿数やフォロワー数をリフレッシュするための残リクエスト数を取得します"""
        return self.__sync_request(self.user.get_profile_refresh_counter_requests())

    def get_social_shared_users(self, **params) -> SocialShareUsersResponse:
        """SNS共有をしたユーザーを取得します"""
        return self.__sync_request(self.user.get_social_shared_users(**params))

    def get_timestamp(self) -> UserTimestampResponse:
        """タイムスタンプを取得します"""
        return self.__sync_request(self.user.get_timestamp())

    def get_user(self, user_id: int) -> UserResponse:
        """ユーザーの情報を取得します"""
        return self.__sync_request(self.user.get_user(user_id))

    def get_user_followers(self, user_id: int, **params) -> FollowUsersResponse:
        """ユーザーのフォロワーを取得します"""
        return self.__sync_request(self.user.get_user_followers(user_id, **params))

    def get_user_followings(self, user_id: int, **params) -> FollowUsersResponse:
        """フォロー中のユーザーを取得します"""
        return self.__sync_request(self.user.get_user_followings(user_id, **params))

    def get_user_from_qr(self, qr: str) -> UserResponse:
        """QRコードからユーザーを取得します"""
        return self.__sync_request(self.user.get_user_from_qr(qr))

    def get_user_without_leaving_footprint(self, user_id: int) -> UserResponse:
        """足跡をつけずにユーザーの情報を取得します"""
        return self.__sync_request(
            self.user.get_user_without_leaving_footprint(user_id)
        )

    def get_users(self, user_ids: list[int]) -> UsersResponse:
        """複数のユーザーの情報を取得します"""
        return self.__sync_request(self.user.get_users(user_ids))

    def refresh_profile_counter(self, counter: str) -> dict:
        """プロフィールのカウンターを更新します"""
        return self.__sync_request(self.user.refresh_profile_counter(counter))

    def register(
        self,
        email: str,
        email_grant_token: str,
        password: str,
        nickname: str,
        birth_date: str,
        gender: Optional[int] = -1,
        country_code: Optional[str] = "JP",
        biography: Optional[str] = None,
        prefecture: Optional[str] = None,
        profile_icon_filename: Optional[str] = None,
        cover_image_filename: Optional[str] = None,
        en: Optional[int] = None,
        vn: Optional[int] = None,
    ):
        return self.__sync_request(
            self.user.register(
                email,
                email_grant_token,
                password,
                nickname,
                birth_date,
                gender,
                country_code,
                biography,
                prefecture,
                profile_icon_filename,
                cover_image_filename,
                en,
                vn,
            )
        )

    def remove_user_avatar(self) -> dict:
        """ユーザーのアイコンを削除します"""
        return self.__sync_request(self.user.remove_user_avatar())

    def remove_user_cover(self) -> dict:
        """ユーザーのカバー画像を削除します"""
        return self.__sync_request(self.user.remove_user_cover())

    def reset_password(
        self,
        email: str,
        email_grant_token: str,
        password: str,
    ) -> dict:
        """パスワードをリセットします"""
        return self.__sync_request(
            self.user.reset_password(email, email_grant_token, password)
        )

    def search_lobi_users(self, **params) -> UsersResponse:
        """Lobiのユーザーを検索します"""
        return self.__sync_request(self.user.search_lobi_users(**params))

    def search_users(self, **params) -> UsersResponse:
        """ユーザーを検索します"""
        return self.__sync_request(self.user.search_users(**params))

    def set_follow_permission_enabled(
        self,
        nickname: str,
        is_private: Optional[bool] = None,
    ) -> dict:
        """フォローを許可制にするかを設定します"""
        return self.__sync_request(
            self.user.set_follow_permission_enabled(nickname, is_private)
        )

    def take_action_follow_request(self, user_id: int, action: str) -> dict:
        """フォローリクエストを操作する"""
        return self.__sync_request(
            self.user.take_action_follow_request(user_id, action)
        )

    def turn_on_hima(self) -> dict:
        """ひまなうを有効にする"""
        return self.__sync_request(self.user.turn_on_hima())

    def unfollow_user(self, user_id: int) -> dict:
        """ユーザーをアンフォローします"""
        return self.__sync_request(self.user.unfollow_user(user_id))

    def update_user(self, nickname: str, **params) -> dict:
        """プロフィールを更新します"""
        return self.__sync_request(self.user.update_user(nickname, **params))

    def block_user(self, user_id: int) -> dict:
        """ユーザーをブロックします"""
        return self.__sync_request(self.user.block_user(user_id))

    def get_blocked_user_ids(self) -> BlockedUserIdsResponse:
        """あなたをブロックしたユーザーを取得します"""
        return self.__sync_request(self.user.get_blocked_user_ids())

    def get_blocked_users(self, from_id: Optional[int] = None) -> BlockedUsersResponse:
        """ブロックしたユーザーを取得します"""
        return self.__sync_request(self.user.get_blocked_users(from_id))

    def unblock_user(self, user_id: int) -> dict:
        """ユーザーをアンブロックします"""
        return self.__sync_request(self.user.unblock_user(user_id))

    def get_hidden_users_list(self, **params) -> HiddenResponse:
        """非表示のユーザー一覧を取得します"""
        return self.__sync_request(self.user.get_hidden_users_list(**params))

    def hide_user(self, user_id: int) -> dict:
        """ユーザーを非表示にします"""
        return self.__sync_request(self.user.hide_user(user_id))

    def unhide_users(self, user_ids: list[int]) -> dict:
        """ユーザーの非表示を解除します"""
        return self.__sync_request(self.user.unhide_users(user_ids))
