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

import os
import time
import random
import logging

from datetime import datetime
from typing import Optional, List, Dict, Any

import httpx
from httpx._types import TimeoutTypes

from .api.auth import AuthAPI
from .api.call import CallAPI
from .api.chat import ChatAPI
from .api.group import GroupAPI
from .api.misc import MiscAPI
from .api.notification import NotificationAPI
from .api.post import PostAPI
from .api.review import ReviewAPI
from .api.thread import ThreadAPI
from .api.user import UserAPI

from . import __version__
from . import ws
from .config import Configs
from .cookie import Cookie, CookieManager
from .errors import (
    HTTPError,
    BadRequestError,
    AuthenticationError,
    ForbiddenError,
    NotFoundError,
    RateLimitError,
    YayServerError,
    ErrorCode,
    ErrorMessage,
)
from .models import (
    ApplicationConfig,
    Attachment,
    BanWord,
    Bgm,
    ChatRoom,
    ConferenceCall,
    CreateGroupQuota,
    Footprint,
    GifImageCategory,
    Group,
    Message,
    Post,
    PresignedUrl,
    PopularWord,
    SharedUrl,
    StickerPack,
    Survey,
    ThreadInfo,
    User,
    UserWrapper,
)
from .responses import (
    ActiveFollowingsResponse,
    ActivitiesResponse,
    BlockedUserIdsResponse,
    BlockedUsersResponse,
    BookmarkPostResponse,
    CallStatusResponse,
    ChatRoomsResponse,
    CreateGroupResponse,
    CreateChatRoomResponse,
    FollowRecommendationsResponse,
    FollowUsersResponse,
    GamesResponse,
    GenresResponse,
    GroupCategoriesResponse,
    GroupsRelatedResponse,
    GroupsResponse,
    GroupThreadListResponse,
    GroupUserResponse,
    GroupUsersResponse,
    HiddenResponse,
    LoginUserResponse,
    LoginUpdateResponse,
    MessageResponse,
    PolicyAgreementsResponse,
    PostsResponse,
    PostLikersResponse,
    PostTagsResponse,
    LikePostsResponse,
    RefreshCounterRequestsResponse,
    RegisterDeviceTokenResponse,
    ReviewsResponse,
    SocialShareUsersResponse,
    TokenResponse,
    UnreadStatusResponse,
    UserResponse,
    UsersResponse,
    RankingUsersResponse,
    UsersByTimestampResponse,
    UserTimestampResponse,
)
from .types import PolicyType
from .utils import Colors, generate_jwt

try:
    from json.decoder import JSONDecodeError
except ImportError:
    JSONDecodeError = ValueError


__all__ = ["Client"]


class HeaderInterceptor(object):
    def __init__(self, cookie: CookieManager, locale: str = "ja") -> None:
        self.__locale: str = locale
        self.__host: str = Configs.PRODUCTION_HOST
        self.__user_agent: str = Configs.USER_AGENT
        self.__device_info: str = Configs.DEVICE_INFO
        self.__app_version: str = Configs.API_VERSION_NAME
        self.__cookie: CookieManager = cookie
        self.__client_ip: str = ""
        self.__connection_speed: str = ""
        self.__connection_type: str = "wifi"
        self.__content_type: str = "application/json;charset=UTF-8"

    def intercept(self, jwt_required: bool = False) -> Dict[str, str]:
        cookie: Cookie = self.__cookie.get()
        headers: dict = {
            "Host": self.__host,
            "User-Agent": self.__user_agent,
            "X-Timestamp": str(int(datetime.now().timestamp())),
            "X-App-Version": self.__app_version,
            "X-Device-Info": self.__device_info,
            "X-Device-UUID": cookie.device.device_uuid,
            "X-Client-IP": self.__client_ip,
            "X-Connection-Type": self.__connection_type,
            "X-Connection-Speed": self.__connection_speed,
            "Accept-Language": self.__locale,
            "Content-Type": self.__content_type,
        }

        if jwt_required:
            headers.update({"X-Jwt": generate_jwt()})

        if len(self.__client_ip):
            headers.update({"X-Client-IP": self.__client_ip})

        if len(cookie.authentication.access_token):
            headers.update(
                {"Authorization": "Bearer " + cookie.authentication.access_token}
            )

        return headers

    def get_client_ip(self) -> str:
        return self.__client_ip

    def get_connection_speed(self) -> str:
        return self.__connection_speed

    def set_client_ip(self, client_ip: str) -> None:
        self.__client_ip = client_ip

    def set_connection_speed(self, connection_speed: str) -> None:
        self.__connection_speed = connection_speed


current_path = os.path.abspath(os.getcwd())


class BaseClient(ws.WebSocketInteractor):
    def __init__(
        self,
        *,
        intents: Optional[ws.Intents] = None,
        proxy_url: Optional[str] = None,
        max_retries: int = 3,
        backoff_factor: float = 1.5,
        wait_on_ratelimit: bool = True,
        min_delay: float = 0.3,
        max_delay: float = 1.2,
        timeout: TimeoutTypes = 30,
        err_lang: str = "ja",
        base_path: str = current_path + "/config/",
        save_cookie_file: bool = True,
        cookie_password: Optional[str] = None,
        cookie_filename: str = "cookies",
        loglevel: int = logging.INFO,
    ) -> None:
        super().__init__(self, intents)

        self.__max_retries: int = max_retries
        self.__backoff_factor: float = backoff_factor
        self.__wait_on_ratelimit: bool = wait_on_ratelimit
        self.__last_request_timestamp: int = 0
        self.__min_delay: float = min_delay
        self.__max_delay: float = max_delay
        self.__err_lang: str = err_lang
        self.__save_cookie_file: bool = save_cookie_file

        self.__cookie: CookieManager = CookieManager(
            save_cookie_file,
            base_path + cookie_filename + ".json",
            cookie_password,
        )

        self.__header_interceptor: HeaderInterceptor = HeaderInterceptor(self.__cookie)
        self.__header_interceptor.set_connection_speed("0 kbps")

        self.__session: httpx.Client = httpx.Client(
            headers=self.__header_interceptor.intercept(),
            http2=True,
            proxies=proxy_url,
            timeout=timeout,
        )

        self.AuthAPI: AuthAPI = AuthAPI(self)
        self.CallAPI: CallAPI = CallAPI(self)
        self.ChatAPI: ChatAPI = ChatAPI(self)
        self.GroupAPI: GroupAPI = GroupAPI(self)
        self.MiscAPI: MiscAPI = MiscAPI(self)
        self.NotificationAPI: NotificationAPI = NotificationAPI(self)
        self.PostAPI: PostAPI = PostAPI(self)
        self.ReviewAPI: ReviewAPI = ReviewAPI(self)
        self.ThreadAPI: ThreadAPI = ThreadAPI(self)
        self.UserAPI: UserAPI = UserAPI(self)

        self.logger: logging.Logger = logging.getLogger(
            "yaylib version: " + __version__
        )

        if not os.path.exists(base_path):
            os.makedirs(base_path)

        ch: logging.StreamHandler = logging.StreamHandler()
        ch.setLevel(loglevel)
        ch.setFormatter(logging.Formatter("%(asctime)s - %(levelname)s - %(message)s"))

        self.logger.addHandler(ch)
        self.logger.setLevel(logging.DEBUG)

        self.logger.info("yaylib version: " + __version__ + " started.")

    @property
    def cookie(self) -> Cookie:
        return self.__cookie.get()

    @property
    def user_id(self) -> int:
        return self.cookie.user.user_id

    @property
    def access_token(self) -> str:
        return self.cookie.authentication.access_token

    @property
    def refresh_token(self) -> str:
        return self.cookie.authentication.refresh_token

    @property
    def uuid(self) -> str:
        return self.cookie.user.uuid

    @property
    def device_uuid(self) -> str:
        return self.cookie.device.device_uuid

    def __make_request(
        self,
        method: str,
        base_url: str,
        route: str,
        params: Optional[Dict[str, Any]] = None,
        payload: Optional[Dict[str, Any]] = None,
        headers: Optional[Dict[str, str]] = None,
        jwt_required: bool = False,
        bypass_delay: bool = False,
    ) -> dict | str:
        endpoint = route
        if base_url is not None:
            endpoint: str = "https://" + base_url + route

        params: Optional[Dict[str, Any]] = self.__filter_params(params)
        payload: Optional[Dict[str, Any]] = self.__filter_params(payload)

        # set client ip address to request header if not exists
        if (
            not self.__header_interceptor.get_client_ip()
            and "v2/users/timestamp" not in endpoint
        ):
            response = self.UserAPI.get_timestamp()
            self.__header_interceptor.set_client_ip(response.ip_address)

        response = None
        backoff_duration: int = 0
        # roughly equivalent to 60 mins, plus extra 15 mins
        max_ratelimit_retries: int = 15
        auth_retry_count, max_auth_retries = 0, 2

        # retry the request based on max_retries
        for i in range(self.__max_retries):
            time.sleep(backoff_duration)

            if headers is None:
                headers = {}
            headers.update(self.__header_interceptor.intercept(jwt_required))

            self.__log_request(method, endpoint, params, headers, payload)

            response = self.__session.request(
                method, endpoint, params=params, json=payload, headers=headers
            )

            if (
                self.__last_request_timestamp
                and int(datetime.now().timestamp()) - self.__last_request_timestamp < 1
            ):
                # insert delay if time between requests is less than 1 second
                sleep_time = random.uniform(self.__min_delay, self.__max_delay)
                time.sleep(sleep_time)

            self.__log_response(response)

            if self.__wait_on_ratelimit and self.__is_ratelimit_error(response):
                # continue attempting request until successful
                # or maximum number of retries is reached
                retries_performed: int = 0

                while retries_performed <= max_ratelimit_retries:
                    retry_after: int = 60 * 5

                    self.logger.warn(
                        f"Rate limit reached. Sleeping for: {retry_after} sec..."
                    )
                    time.sleep(retry_after + 1)

                    headers.update(self.__header_interceptor.intercept())

                    self.__log_request(method, endpoint, params, headers, payload)

                    response = self.__session.request(
                        method, endpoint, params=params, json=payload, headers=headers
                    )

                    self.__log_response(response)

                    if not self.__is_ratelimit_error(response):
                        break

                    retries_performed += 1

                if retries_performed >= max_ratelimit_retries:
                    raise RateLimitError("Maximum rate limit reached.")

            if self.__save_cookie_file and self.__is_access_token_expired_error(
                response
            ):
                # remove the cookie file and stop
                # the proccessing if refresh token has expired
                if "/api/v1/oauth/token" in endpoint:
                    self.__cookie.destroy()
                    raise AuthenticationError(
                        "Refresh token expired. Try logging in again."
                    )

                auth_retry_count += 1

                if auth_retry_count < max_auth_retries:
                    if self.user_id == 0:
                        raise AuthenticationError(
                            "Please log in to perform the action."
                        )
                    self.__refresh_tokens()
                    continue
                else:
                    self.__cookie.destroy()
                    raise AuthenticationError(
                        "Maximum authentication retries exceeded. Try logging in again."
                    )

            if response.status_code not in [500, 502, 503, 504]:
                break

            if response is not None:
                self.logger.error(
                    f"Request failed with status code {response.status_code}. Retrying...",
                    exc_info=True,
                )
            else:
                self.logger.error("Request failed. Retrying...")

            backoff_duration = self.__backoff_factor * (2**i)

        self.__last_request_timestamp = 0
        if not bypass_delay:
            self.__last_request_timestamp = int(datetime.now().timestamp())

        try:
            f_response = response.json()
        except JSONDecodeError:
            f_response = response.text

        if isinstance(f_response, dict):
            f_response = self.__translate_error_message(f_response)

        if response.status_code == 400:
            raise BadRequestError(f_response)
        if response.status_code == 401:
            raise AuthenticationError(f_response)
        if response.status_code == 403:
            raise ForbiddenError(f_response)
        if response.status_code == 404:
            raise NotFoundError(f_response)
        if response.status_code == 429:
            raise RateLimitError(f_response)
        if response.status_code == 500:
            raise YayServerError(f_response)
        if response.status_code and not 200 <= response.status_code < 300:
            raise HTTPError(f_response)

        return f_response

    def __log_request(
        self, method: str, endpoint: str, params: dict, headers: dict, payload: dict
    ):
        self.logger.debug(
            f"{Colors.HEADER}Making API request: {method} {endpoint}{Colors.RESET}\n\n"
            f"Parameters: {params}\n\n"
            f"Headers: {headers}\n\n"
            f"Body: {payload}\n"
        )

    def __log_response(self, response: httpx.Response):
        self.logger.debug(
            f"Received API response:\n\n"
            f"Status code: {response.status_code}\n\n"
            f"Headers: {response.headers}\n\n"
            f"Response: {response.text}\n"
        )

    def __filter_params(
        self, params: Optional[Dict[str, Any]] = None
    ) -> Optional[Dict[str, Any]]:
        if params is None:
            return None
        new_params = {}
        for k in params:
            if params[k] is not None:
                new_params[k] = params[k]
        return new_params

    def __is_access_token_expired_error(self, response: httpx.Response) -> bool:
        json_response: dict = response.json()
        return response.status_code == 401 and (
            json_response.get("error_code") == ErrorCode.AccessTokenExpired.value
            or json_response.get("error_code") == ErrorCode.AccessTokenInvalid.value
        )

    def __is_ratelimit_error(self, response: httpx.Response) -> bool:
        if response.status_code == 429:
            return True
        if response.status_code == 400:
            json_response: dict = response.json()
            if json_response.get("error_code") == ErrorCode.QuotaLimitExceeded.value:
                return True
        return False

    def __refresh_tokens(self) -> None:
        response: TokenResponse = self.AuthAPI.get_token(
            grant_type="refresh_token", refresh_token=self.refresh_token
        )
        self.__cookie.set(
            Cookie(
                {
                    **self.cookie.to_dict(),
                    "authentication": {
                        "access_token": response.access_token,
                        "refresh_token": response.refresh_token,
                    },
                }
            )
        )
        self.__cookie.save()

    def __translate_error_message(self, f_response: dict) -> dict:
        if not self.__err_lang == "ja":
            return f_response
        try:
            error_code = f_response.get("error_code")
            if error_code is not None:
                error_type = ErrorCode(error_code)
                if error_type.name in ErrorMessage.__members__:
                    error_message = ErrorMessage[error_type.name].value
                    f_response["message"] = error_message
            return f_response
        except ValueError:
            return f_response

    def __construct_response(
        self, response: dict, data_type: Optional[object] = None
    ) -> dict:
        if data_type is not None:
            if isinstance(response, list):
                response = [data_type(result) for result in response]
            elif response is not None:
                response = data_type(response)
        return response

    def _request(
        self,
        method: str,
        route: str,
        base_url: str = Configs.PRODUCTION_HOST,
        params: Optional[Dict[str, Any]] = None,
        payload: Optional[Dict[str, Any]] = None,
        data_type: Optional[object] = None,
        headers: Optional[Dict[str, Any]] = None,
        jwt_required: bool = False,
        bypass_delay: bool = False,
    ) -> object | dict:
        res: dict | str = self.__make_request(
            method=method,
            base_url=base_url,
            route=route,
            params=params,
            payload=payload,
            headers=headers,
            jwt_required=jwt_required,
            bypass_delay=bypass_delay,
        )
        if data_type:
            return self.__construct_response(res, data_type)
        return res

    def _prepare(self, email: str, password: str) -> LoginUserResponse:
        try:
            self.__cookie.load(email)
            return LoginUserResponse(
                {
                    "access_token": self.access_token,
                    "refresh_token": self.refresh_token,
                    "user_id": self.user_id,
                }
            )
        except:
            response = self.AuthAPI.login_with_email(email, password)

            if response.access_token is None:
                raise ForbiddenError("Invalid email or password.")

            self.__cookie.set(
                Cookie(
                    {
                        "authentication": {
                            "access_token": response.access_token,
                            "refresh_token": response.refresh_token,
                        },
                        "user": {
                            "user_id": response.user_id,
                            "email": email,
                            "uuid": self.uuid,
                        },
                        "device": {"device_uuid": self.device_uuid},
                    }
                )
            )
            self.__cookie.save()

        self.logger.info(f"Successfully logged in as [{response.user_id}]")

        # agree to the policy stuff
        policy_response = self.MiscAPI.get_policy_agreements()

        if not policy_response.latest_privacy_policy_agreed:
            self.MiscAPI.accept_policy_agreement(type=PolicyType.privacy_policy)

        if not policy_response.latest_terms_of_use_agreed:
            self.MiscAPI.accept_policy_agreement(type=PolicyType.terms_of_use)

        return response


class Client(BaseClient):
    """

    Yay!（イェイ）- API クライアント

    #### Useage

        >>> import yaylib
        >>> api = yaylib.Client()
        >>> api.login(email, password)
        >>> api.create_post("Hello with yaylib!")

    #### Parameters

        - intents: Intents - (optional)
        - proxy_url: str - (optional)
        - max_retries: int - (optional)
        - backoff_factor: float - (optional)
        - wait_on_ratelimit: bool - (optional)
        - min_delay: float - (optional)
        - max_delay: float - (optional)
        - timeout: int - (optional)
        - err_lang: str - (optional)
        - base_path: str - (optional)
        - save_cookie_file: bool - (optional)
        - cookie_password: bool - (optional)
        - cookie_filename: str - (optional)
        - loglevel: int - (optional)

    <https://github.com/ekkx/yaylib>

    ---

    ### Yay! (nanameue, Inc.) API Client

    Copyright (c) 2023 ekkx

    """

    # -CALL

    def get_user_active_call(self, user_id: int) -> Post:
        """

        ユーザーが参加中の通話を取得します

        """
        return self.CallAPI.get_user_active_call(user_id).post

    def get_bgms(self) -> list[Bgm]:
        """

        通話のBGMを取得します

        """
        return self.CallAPI.get_bgms().bgm

    def get_call(self, call_id: int) -> ConferenceCall:
        """

        通話の詳細を取得します

        """
        return self.CallAPI.get_call(call_id).conference_call

    def get_call_invitable_users(
        self, call_id: int, from_timestamp: Optional[int] = None
    ) -> UsersByTimestampResponse:
        """

        通話に招待可能なユーザーを取得します

        """
        return self.CallAPI.get_call_invitable_users(call_id, from_timestamp)

    def get_call_status(self, opponent_id: int) -> CallStatusResponse:
        """

        通話の状態を取得します

        """
        return self.CallAPI.get_call_status(opponent_id)

    def get_games(self, **params) -> GamesResponse:
        """

        ゲームを取得します

        Parameters
        ----------
            - number: int - (optional)
            - ids: list[int] - (optional)
            - from_id: int - (optional)

        """
        return self.CallAPI.get_games(**params)

    def get_genres(self, **params) -> GenresResponse:
        """

        通話のジャンルを取得します

        Parameters
        ----------
            - number: int - (optional)
            - from: int - (optional)

        """
        return self.CallAPI.get_genres(**params)

    def get_group_calls(self, **params) -> PostsResponse:
        """

        サークルの通話を取得します

        Parameters
        ----------
            - number: int - (optional)
            - group_category_id: int - (optional)
            - from_timestamp: int - (optional)
            - scope: str - (optional)

        """
        return self.CallAPI.get_group_calls(**params)

    def invite_online_followings_to_call(
        self,
        call_id: int,
        group_id: Optional[int] = None,
    ) -> dict:
        """

        オンラインの友達をまとめて通話に招待します

        Parameters
        ----------
            - call_id: int - (required)
            - group_id: int - (optional)

        """
        return self.CallAPI.invite_to_call_bulk(call_id, group_id)

    def invite_users_to_call(self, call_id: int, user_ids: list[int]) -> dict:
        """

        ユーザーを通話に招待します

        Parameters
        ----------
            - call_id: int - (required)
            - user_ids: list[int] - (required)

        """
        return self.CallAPI.invite_users_to_call(call_id, user_ids)

    def invite_users_to_chat_call(
        self,
        chat_room_id: int,
        room_id: int,
        room_url: str,
    ) -> dict:
        """

        チャット通話にユーザーを招待します

        Parameters
        ----------

            - chat_room_id: int - (required)
            - room_id: int - (required)
            - room_url: int - (required)

        """
        return self.CallAPI.invite_users_to_chat_call(chat_room_id, room_id, room_url)

    def kick_user_from_call(self, call_id: int, user_id: int) -> dict:
        """

        ユーザーを通話からキックします

        Parameters
        ----------

            - call_id: int - (required)
            - user_id: int - (required)

        """
        return self.CallAPI.kick_and_ban_from_call(call_id, user_id)

    def set_call(
        self,
        call_id: int,
        joinable_by: str,
        game_title: Optional[str] = None,
        category_id: Optional[str] = None,
    ) -> dict:
        """

        通話を開始します

        """
        return self.CallAPI.set_call(call_id, joinable_by, game_title, category_id)

    def set_user_role(self, call_id: int, user_id: int, role: str) -> dict:
        """

        通話に参加中のユーザーに役職を与えます

        """
        return self.CallAPI.set_user_role(call_id, user_id, role)

    def join_call(
        self, conference_id: int, call_sid: Optional[str] = None
    ) -> ConferenceCall:
        """

        通話に参加します

        """
        return self.CallAPI.start_call(conference_id, call_sid).conference_call

    def join_call_as_anonymous(
        self, conference_id: int, agora_uid: str
    ) -> ConferenceCall:
        """

        無名くんとして通話に参加します

        """
        return self.CallAPI.start_anonymous_call(
            conference_id, agora_uid
        ).conference_call

    def leave_call(
        self,
        conference_id: int,
        call_sid: Optional[str] = None,
    ) -> dict:
        """

        通話から退出します

        """
        return self.CallAPI.stop_call(conference_id, call_sid)

    def leave_call_as_anonymous(
        self, conference_id: int, agora_uid: Optional[str] = None
    ):
        """

        通話から退出します

        """
        return self.CallAPI.stop_anonymous_call(conference_id, agora_uid)

    # -CASSANDRA

    def get_activities(self, **params) -> ActivitiesResponse:
        """

        通知を取得します

        Parameters
        ----------
            - important: bool - (required)
            - from_timestamp: int - (optional)
            - number: int - (optional)

        """
        return self.NotificationAPI.get_user_activities(**params)

    def get_merged_activities(self, **params) -> ActivitiesResponse:
        """

        全種類の通知を取得します

        Parameters
        ----------

            - from_timestamp: int - (optional)
            - number: int - (optional)

        """
        return self.NotificationAPI.get_user_merged_activities(**params)

    def received_notification(
        self,
        pid: str,
        type: str,
        opened_at: Optional[int] = None,
    ) -> dict:
        return self.NotificationAPI.received_notification(pid, type, opened_at)

    # -CHAT

    def accept_chat_requests(self, chat_room_ids: list[int]) -> dict:
        """

        チャットリクエストを承認します

        Parameters
        ----------

            - chat_room_ids: list[int] - (required)

        """
        return self.ChatAPI.accept_chat_requests(chat_room_ids)

    def check_unread_status(self, from_time: int) -> UnreadStatusResponse:
        """

        チャットの未読ステータスを確認します

        """
        return self.ChatAPI.check_unread_status(from_time)

    def create_group_chat(
        self,
        name: str,
        with_user_ids: list[int],
        icon_filename: Optional[str] = None,
        background_filename: Optional[str] = None,
    ) -> CreateChatRoomResponse:
        """

        グループチャットを作成します

        Parameters
        ----------

            - name: str - (required)
            - with_user_ids: list[int] - (required)
            - icon_filename: str - (optional)
            - background_filename: str - (optional)

        """
        return self.ChatAPI.create_group_chat(
            name, with_user_ids, icon_filename, background_filename
        )

    def create_private_chat(
        self,
        with_user_id: int,
        matching_id: Optional[int] = None,
        hima_chat: Optional[bool] = False,
    ) -> CreateChatRoomResponse:
        """

        個人チャットを作成します

        Parameters
        ----------

            - with_user_id: int - (required)
            - matching_id: str - (optional)
            - hima_chat: bool - (optional)

        """
        return self.ChatAPI.create_private_chat(with_user_id, matching_id, hima_chat)

    def delete_background(self, room_id: int) -> dict:
        """

        チャットの背景画像を削除します

        Parameters
        ----------

            - room_id: int - (required)

        """
        return self.ChatAPI.delete_background(room_id)

    def delete_message(self, room_id: int, message_id: int) -> dict:
        """

        メッセージを削除します

        Parameters
        ----------

            - room_id: int - (required)
            - message_id: int - (required)

        """
        return self.ChatAPI.delete_message(room_id, message_id)

    def edit_chat_room(
        self,
        chat_room_id: int,
        name: str,
        icon_filename: Optional[str] = None,
        background_filename: Optional[str] = None,
    ) -> dict:
        """

        チャットルームを編集します

        """
        return self.ChatAPI.edit_chat_room(
            chat_room_id, name, icon_filename, background_filename
        )

    def get_chatable_users(
        self,
        from_follow_id: Optional[int] = None,
        from_timestamp: Optional[int] = None,
        order_by: Optional[str] = None,
    ) -> FollowUsersResponse:
        """

        チャット可能なユーザーを取得します

        """
        return self.ChatAPI.get_chatable_users(
            from_follow_id,
            from_timestamp,
            order_by,
        )

    def get_gifs_data(self) -> list[GifImageCategory]:
        """

        チャットルームのGIFデータを取得します

        """
        return self.ChatAPI.get_gifs_data().gif_categories

    def get_hidden_chat_rooms(self, **params) -> ChatRoomsResponse:
        """

        非表示のチャットルームを取得します

        Parameters
        ----------

            - from_timestamp: int - (optional)
            - number: int - (optional)

        """
        return self.ChatAPI.get_hidden_chat_rooms(**params)

    def get_main_chat_rooms(
        self, from_timestamp: Optional[int] = None
    ) -> ChatRoomsResponse:
        """

        メインのチャットルームを取得します

        """
        return self.ChatAPI.get_main_chat_rooms(from_timestamp)

    def get_messages(self, chat_room_id: int, **params) -> list[Message]:
        """

        メッセージを取得します

        Parameters
        ---------------
            - from_message_id: int - (optional)
            - to_message_id: int - (optional)

        """
        return self.ChatAPI.get_messages(chat_room_id, **params).messages

    def get_chat_requests(self, **params) -> ChatRoomsResponse:
        """

        チャットリクエストを取得します

        Parameters:
        -----------

            - number: int (optional)
            - from_timestamp: int (optional)
            - : str (optional)

        """
        return self.ChatAPI.get_request_chat_rooms(**params)

    def get_chat_room(self, chat_room_id: int) -> ChatRoom:
        """

        チャットルームの詳細を取得します

        """
        return self.ChatAPI.get_chat_room(chat_room_id).chat

    def get_sticker_packs(self) -> list[StickerPack]:
        """

        スタンプを取得します

        """
        return self.ChatAPI.get_sticker_packs().sticker_packs

    def get_chat_requests_count(self) -> int:
        """

        チャットリクエストの数を取得します

        """
        return self.ChatAPI.get_total_chat_requests().total

    def hide_chat(self, chat_room_id: int) -> dict:
        """

        チャットルームを非表示にします

        """
        return self.ChatAPI.hide_chat(chat_room_id)

    def invite_to_chat(self, chat_room_id: int, user_ids: list[int]) -> dict:
        """

        チャットルームにユーザーを招待します

        """
        return self.ChatAPI.invite_to_chat(chat_room_id, user_ids)

    def kick_users_from_chat(self, chat_room_id: int, user_ids: list[int]) -> dict:
        """

        チャットルームからユーザーを追放します

        """
        return self.ChatAPI.kick_users_from_chat(chat_room_id, user_ids)

    def pin_chat(self, room_id: int) -> dict:
        """

        チャットルームをピン止めします

        """
        return self.ChatAPI.pin_chat(room_id)

    def read_message(self, chat_room_id: int, message_id: int) -> dict:
        """

        メッセージを既読にします

        """
        return self.ChatAPI.read_message(chat_room_id, message_id)

    def refresh_chat_rooms(self, from_time: Optional[int] = None) -> ChatRoomsResponse:
        """

        チャットルームを更新します

        """
        return self.ChatAPI.refresh_chat_rooms(from_time)

    def remove_chat_rooms(self, chat_room_ids: list[int]) -> dict:
        """

        チャットルームを削除します

        """
        return self.ChatAPI.remove_chat_rooms(chat_room_ids)

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
        """

        チャットルームを通報します

        """
        return self.ChatAPI.report_chat_room(
            chat_room_id,
            opponent_id,
            category_id,
            reason,
            screenshot_filename,
            screenshot_2_filename,
            screenshot_3_filename,
            screenshot_4_filename,
        )

    def send_message(
        self,
        chat_room_id: int,
        **params,
    ) -> MessageResponse:
        """

        チャットルームにメッセージを送信します

        Parameters:
        -----------

            - chat_room_id: int (required)
            - message_type: str (required)
            - call_type: str = (optional)
            - text: str = (optional)
            - font_size: int = (optional)
            - gif_image_id: int = (optional)
            - attachment_file_name: str = (optional)
            - sticker_pack_id: int = (optional)
            - sticker_id: int = (optional)
            - video_file_name: str = (optional)
            - parent_id: int = (optional)
            - : str = (optional)

        """
        return self.ChatAPI.send_message(chat_room_id, **params)

    def unhide_chat(self, chat_room_ids: int) -> dict:
        """

        チャットの非表示を解除します

        """
        return self.ChatAPI.unhide_chat(chat_room_ids)

    def unpin_chat(self, chat_room_id: int) -> dict:
        """

        チャットルームのピン止めを解除します

        """
        return self.ChatAPI.unpin_chat(chat_room_id)

    # -GROUP

    def accept_moderator_offer(self, group_id: int) -> dict:
        """

        サークル副管理人の権限オファーを引き受けます

        """
        return self.GroupAPI.accept_moderator_offer(group_id)

    def accept_ownership_offer(self, group_id: int) -> dict:
        """

        サークル管理人の権限オファーを引き受けます

        """
        return self.GroupAPI.accept_ownership_offer(group_id)

    def accept_group_join_request(self, group_id: int, user_id: int) -> dict:
        """

        サークル参加リクエストを承認します

        """
        return self.GroupAPI.accept_group_join_request(group_id, user_id)

    def add_related_groups(self, group_id: int, related_group_id: list[int]) -> dict:
        """

        関連サークルを追加します

        """
        return self.GroupAPI.add_related_groups(group_id, related_group_id)

    def ban_group_user(self, group_id: int, user_id: int) -> dict:
        """

        サークルからユーザーを追放します

        """
        return self.GroupAPI.ban_group_user(group_id, user_id)

    def check_unread_status(
        self, from_time: Optional[int] = None
    ) -> UnreadStatusResponse:
        """

        サークルの未読ステータスを取得します

        """
        return self.GroupAPI.check_unread_status(from_time)

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
        """

        サークルを作成します

        """
        return self.GroupAPI.create_group(
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

    def pin_group(self, group_id: int) -> dict:
        """

        サークルをピンします

        """
        return self.GroupAPI.create_pin_group(group_id)

    def decline_moderator_offer(self, group_id: int) -> dict:
        """

        サークル副管理人の権限オファーを断ります

        """
        return self.GroupAPI.decline_moderator_offer(group_id)

    def decline_ownership_offer(self, group_id: int) -> dict:
        """

        サークル管理人の権限オファーを断ります

        """
        return self.GroupAPI.decline_ownership_offer(group_id)

    def decline_group_join_request(self, group_id: int, user_id: int) -> dict:
        """

        サークル参加リクエストを断ります

        """
        return self.GroupAPI.decline_group_join_request(group_id, user_id)

    def unpin_group(self, group_id: int) -> dict:
        """

        サークルのピン止めを解除します

        """
        return self.GroupAPI.delete_pin_group(group_id)

    def get_banned_group_members(
        self,
        group_id: int,
        page: Optional[int] = None,
    ) -> UsersResponse:
        """

        追放されたサークルメンバーを取得します

        """
        return self.GroupAPI.get_banned_group_members(group_id, page)

    def get_group_categories(self, **params) -> GroupCategoriesResponse:
        """

        サークルのカテゴリーを取得します

        Parameters
        ----------

            - page: int - (optional)
            - number: int - (optional)

        """
        return self.GroupAPI.get_group_categories(**params)

    def get_create_group_quota(self) -> CreateGroupQuota:
        """

        サークル作成可能な割当量を取得します

        """
        return self.GroupAPI.get_create_group_quota()

    def get_group(self, group_id: int) -> Group:
        """

        サークルの詳細を取得します

        """
        return self.GroupAPI.get_group(group_id).group

    def get_groups(self, **params) -> GroupsResponse:
        """

        複数のサークルの詳細を取得します

        Parameters
        ----------

            - group_category_id: int = None
            - keyword: str = None
            - from_timestamp: int = None
            - sub_category_id: int = None

        """
        return self.GroupAPI.get_groups(**params)

    def get_invitable_users(self, group_id: int, **params) -> UsersByTimestampResponse:
        """

        サークルに招待可能なユーザーを取得します

        Parameters
        ----------

            - from_timestamp: int - (optional)
            - user[nickname]: str - (optional)

        """
        return self.GroupAPI.get_invitable_users(group_id, **params)

    def get_joined_statuses(self, ids: list[int]) -> dict:
        """

        サークルの参加ステータスを取得します

        """
        return self.GroupAPI.get_joined_statuses(ids)

    def get_group_member(self, group_id: int, user_id: int) -> GroupUserResponse:
        """

        特定のサークルメンバーの情報を取得します

        """
        return self.GroupAPI.get_group_member(group_id, user_id)

    def get_group_members(self, group_id: int, **params) -> GroupUsersResponse:
        """

        サークルメンバーを取得します

        Parameters
        ----------

            - id: int - (required)
            - mode: str - (optional)
            - keyword: str - (optional)
            - from_id: int - (optional)
            - from_timestamp: int - (optional)
            - order_by: str - (optional)
            - followed_by_me: bool - (optional)

        """
        return self.GroupAPI.get_group_members(group_id, **params)

    def get_my_groups(self, from_timestamp: Optional[int] = None) -> GroupsResponse:
        """

        自分のサークルを取得します

        """
        return self.GroupAPI.get_my_groups(from_timestamp)

    def get_relatable_groups(self, group_id: int, **params) -> GroupsRelatedResponse:
        """

        関連がある可能性があるサークルを取得します

        Parameters
        ----------

            - group_id: int - (required)
            - keyword: str - (optional)
            - from: str - (optional)

        """
        return self.GroupAPI.get_relatable_groups(group_id, **params)

    def get_related_groups(self, group_id: int, **params) -> GroupsRelatedResponse:
        """

        関連があるサークルを取得します

        Parameters
        ----------

            - group_id: int - (required)
            - keyword: str - (optional)
            - from: str - (optional)

        """
        return self.GroupAPI.get_related_groups(group_id, **params)

    def get_user_groups(self, **params) -> GroupsResponse:
        """

        特定のユーザーが参加しているサークルを取得します

        Parameters
        ----------

            - user_id: int - (required)
            - page: int - (optional)

        """
        return self.GroupAPI.get_user_groups(**params)

    def invite_users_to_group(self, group_id: int, user_ids: list[int]) -> dict:
        """

        サークルにユーザーを招待します

        """
        return self.GroupAPI.invite_users_to_group(group_id, user_ids)

    def join_group(self, group_id: int) -> dict:
        """

        サークルに参加します

        """
        return self.GroupAPI.join_group(group_id)

    def leave_group(self, group_id: int) -> dict:
        """

        サークルから脱退します

        """
        return self.GroupAPI.leave_group(group_id)

    def post_gruop_social_shared(self, group_id: int, sns_name: str) -> dict:
        """

        サークルのソーシャルシェアを投稿します

        """
        return self.GroupAPI.post_gruop_social_shared(group_id, sns_name)

    def remove_group_cover(self, group_id: int) -> dict:
        """

        サークルのカバー画像を削除します

        """
        return self.GroupAPI.remove_group_cover(group_id)

    def remove_moderator(self, group_id: int, user_id: int) -> dict:
        """

        サークルの副管理人を削除します

        """
        return self.GroupAPI.remove_moderator(group_id, user_id)

    def remove_related_groups(
        self, group_id: int, related_group_ids: list[int]
    ) -> dict:
        """

        関連のあるサークルを削除します

        """
        return self.GroupAPI.remove_related_groups(group_id, related_group_ids)

    def report_group(
        self,
        group_id: int,
        category_id: int,
        reason: Optional[str] = None,
        opponent_id: Optional[int] = None,
        screenshot_filename: Optional[str] = None,
        screenshot_2_filename: Optional[str] = None,
        screenshot_3_filename: Optional[str] = None,
        screenshot_4_filename: Optional[str] = None,
    ) -> dict:
        """

        サークルを通報します

        """
        return self.GroupAPI.report_group(
            group_id,
            category_id,
            reason,
            opponent_id,
            screenshot_filename,
            screenshot_2_filename,
            screenshot_3_filename,
            screenshot_4_filename,
        )

    def send_moderator_offers(self, group_id: int, user_ids: list[int]) -> dict:
        """

        複数人にサークル副管理人のオファーを送信します

        """
        return self.GroupAPI.send_moderator_offers(group_id, user_ids)

    def send_ownership_offer(self, group_id: int, user_id: int) -> dict:
        """

        サークル管理人権限のオファーを送信します

        """
        return self.GroupAPI.send_ownership_offer(group_id, user_id)

    def set_group_title(self, group_id: int, title: str) -> dict:
        """

        サークルのタイトルを設定します

        """
        return self.GroupAPI.set_group_title(group_id, title)

    def take_over_group_ownership(self, group_id: int) -> dict:
        """

        サークル管理人の権限を引き継ぎます

        """
        return self.GroupAPI.take_over_group_ownership(group_id)

    def unban_group_member(self, group_id: int, user_id: int) -> dict:
        """

        特定のサークルメンバーの追放を解除します

        """
        return self.GroupAPI.unban_group_member(group_id, user_id)

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
    ) -> Group:
        """

        サークルを編集します

        """
        return self.GroupAPI.update_group(
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
        ).group

    def withdraw_moderator_offer(self, group_id: int, user_id: int) -> dict:
        """

        サークル副管理人のオファーを取り消します

        """
        return self.GroupAPI.withdraw_moderator_offer(group_id, user_id)

    def withdraw_ownership_offer(self, group_id: int, user_id: int) -> dict:
        """

        サークル管理人のオファーを取り消します

        """
        return self.GroupAPI.withdraw_ownership_offer(group_id, user_id)

    # -LOGIN

    def change_password(
        self, current_password: str, new_password: str
    ) -> LoginUpdateResponse:
        """

        パスワードを変更します

        """
        return self.AuthAPI.change_password(current_password, new_password)

    def get_token(
        self,
        grant_type: str,
        refresh_token: Optional[str] = None,
        email: Optional[str] = None,
        password: Optional[str] = None,
    ) -> TokenResponse:
        """

        トークンを再発行します

        """
        return self.AuthAPI.get_token(grant_type, refresh_token, email, password)

    def login(self, email: str, password: str) -> LoginUserResponse:
        """

        メールアドレスでログインします

        """
        return self._prepare(email, password)

    def logout(self) -> dict:
        """

        ログアウトします

        """
        return self.AuthAPI.logout()

    def resend_confirm_email(self) -> dict:
        """

        確認メールを再送信します

        """
        return self.AuthAPI.resend_confirm_email()

    def restore_user(self, user_id: int) -> LoginUserResponse:
        """

        ユーザーを復元します

        """
        return self.AuthAPI.restore_user(user_id)

    def register_device_token(
        self,
        device_token: str,
        device_type: str,
        os_version: str,
        app_version: str,
        screen_resolution: str,
        screen_density: str,
        device_model: str,
        appsflyer_id: str,
        advertising_id: Optional[str] = None,
    ) -> RegisterDeviceTokenResponse:
        """

        デバイストークンを登録します

        """
        return self.AuthAPI.register_device_token(
            device_token,
            device_type,
            os_version,
            app_version,
            screen_resolution,
            screen_density,
            device_model,
            appsflyer_id,
            advertising_id,
        )

    def revoke_tokens(self) -> dict:
        """

        トークンを無効化します

        """
        return self.AuthAPI.revoke_tokens()

    def save_account_with_email(
        self,
        email: str,
        password: Optional[str] = None,
        current_password: Optional[str] = None,
        email_grant_token: Optional[str] = None,
    ) -> LoginUpdateResponse:
        """

        メールアドレスでアカウントを保存します

        """
        return self.AuthAPI.save_account_with_email(
            email, password, current_password, email_grant_token
        )

    # -MISC

    def accept_policy_agreement(self, type: str):
        """

        利用規約、ポリシー同意書に同意します

        """
        return self.MiscAPI.accept_policy_agreement(type)

    def send_verification_code(self, email: str, intent: str, locale="ja"):
        """

        認証コードを送信します

        """
        return self.MiscAPI.send_verification_code(email, intent, locale)

    def get_email_grant_token(self, code: int, email: str) -> str:
        """

        email_grant_tokenを取得します

        """
        return self.MiscAPI.get_email_grant_token(code, email).email_grant_token

    def get_email_verification_presigned_url(
        self, email: str, locale: str, intent: Optional[str] = None
    ) -> str:
        """

        メールアドレス確認用の署名付きURLを取得します

        """
        return self.MiscAPI.get_email_verification_presigned_url(email, locale, intent)

    def get_file_upload_presigned_urls(
        self, file_names: list[str]
    ) -> list[PresignedUrl]:
        """

        ファイルアップロード用の署名付きURLを取得します

        """
        return self.MiscAPI.get_file_upload_presigned_urls(file_names).presigned_urls

    # def get_id_checker_presigned_url(
    #         self,
    #         model: str,
    #         action: str,
    #         **params
    # ) -> str:
    #     return get_id_checker_presigned_url(self, model, action, **params)

    def get_old_file_upload_presigned_url(self, video_file_name: str) -> str:
        """

        動画ファイルアップロード用の署名付きURLを取得します

        """
        return self.MiscAPI.get_old_file_upload_presigned_url(
            video_file_name
        ).presigned_url

    def get_policy_agreements(self) -> PolicyAgreementsResponse:
        """

        利用規約、ポリシー同意書に同意しているかどうかを取得します

        """
        return self.MiscAPI.get_policy_agreements()

    def get_web_socket_token(self) -> str:
        """

        Web Socket Tokenを取得します

        """
        return self.MiscAPI.get_web_socket_token().token

    def upload_image(self, image_paths: list[str], image_type: str) -> list[Attachment]:
        """

        画像をアップロードして、サーバー上のファイル名のリストを返します。

        Parameteres
        -----------

            - image_path: list[str] - (required): 画像パスのリスト
            - image_type: str - (required): 画像の種類

        Examples
        --------

        投稿に画像を付与する場合

        >>> # サーバー上にアップロード
        >>> attachments = api.upload_image(
        >>>     image_type=yaylib.IMAGE_TYPE_POST,
        >>>     image_paths=["./test.jpg"],
        >>> )
        >>> # サーバー上のファイル名を指定
        >>> api.create_post(
        >>>     "Hello with yaylib!",
        >>>     attachment_filename=attachments[0].filename
        >>> )

        """
        return self.MiscAPI.upload_image(image_paths, image_type)

    def upload_video(self, video_path: str) -> str:
        """

        動画をアップロードして、サーバー上のファイル名を返します。

        Parameteres
        -----------

            - video_path: str - (required): 動画ファイルの場所

        Examples
        --------

        投稿に動画を付与する場合

        >>> # サーバー上にアップロード
        >>> filename = api.upload_video(
        >>>     video_path="./example.mp4",
        >>> )
        >>> # サーバー上のファイル名を指定
        >>> api.create_post(
        >>>     "Hello with yaylib!",
        >>>     video_file_name=filename
        >>> )

        """
        return self.MiscAPI.upload_video(video_path)

    def get_app_config(self) -> ApplicationConfig:
        """

        アプリケーションの設定情報を取得します

        """
        return self.MiscAPI.get_app_config().app

    def get_banned_words(self, country_code: str = "jp") -> list[BanWord]:
        """

        禁止ワードの一覧を取得します

        """
        return self.MiscAPI.get_banned_words(country_code).ban_words

    def get_popular_words(self, country_code: str = "jp") -> list[PopularWord]:
        """

        人気のワードの一覧を取得します

        """
        return self.MiscAPI.get_popular_words(country_code).popular_words

    # -POST

    def add_bookmark(self, user_id: int, post_id: int) -> BookmarkPostResponse:
        """

        ブックマークに追加します

        """
        return self.PostAPI.add_bookmark(user_id, post_id)

    def add_group_highlight_post(self, group_id: int, post_id: int) -> dict:
        """

        投稿をグループのまとめに追加します

        """
        return self.PostAPI.add_group_highlight_post(group_id, post_id)

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
        message_tags: Optional[list] = [],
        attachment_filename: Optional[str] = None,
        attachment_2_filename: Optional[str] = None,
        attachment_3_filename: Optional[str] = None,
        attachment_4_filename: Optional[str] = None,
        attachment_5_filename: Optional[str] = None,
        attachment_6_filename: Optional[str] = None,
        attachment_7_filename: Optional[str] = None,
        attachment_8_filename: Optional[str] = None,
        attachment_9_filename: Optional[str] = None,
    ) -> ConferenceCall:
        """

        通話の投稿を作成します

        """
        return self.PostAPI.create_call_post(
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
        ).conference_call

    def pin_group_post(self, post_id: int, group_id: int) -> dict:
        """

        グループの投稿をピンします

        """
        return self.PostAPI.create_group_pin_post(post_id, group_id)

    def pin_post(self, post_id: int) -> dict:
        """

        投稿をピンします

        """
        return self.PostAPI.create_pin_post(post_id)

    def mention(self, user_id: int) -> str:
        """

        メンション用文字列を返します

        """
        return self.PostAPI.mention(user_id)

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
        message_tags: Optional[list] = [],
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
        """

        投稿を作成します

        """
        return self.PostAPI.create_post(
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
        message_tags: Optional[list] = [],
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
        """

        投稿を(´∀｀∩)↑age↑します

        """
        return self.PostAPI.create_repost(
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
        ).post

    def create_share_post(
        self,
        shareable_type: str,
        shareable_id: int,
        text: Optional[str] = None,
        font_size: Optional[int] = None,
        color: Optional[int] = None,
        group_id: Optional[int] = None,
    ) -> Post:
        """

        シェア投稿を作成します

        """
        return self.PostAPI.create_share_post(
            shareable_type, shareable_id, text, font_size, color, group_id
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
        message_tags: Optional[list] = [],
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
        """

        スレッドの投稿を作成します

        """
        return self.PostAPI.create_thread_post(
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

    def delete_all_post(self) -> dict:
        """

        すべての自分の投稿を削除します

        """
        return self.PostAPI.delete_all_post()

    def unpin_group_post(self, group_id: int) -> dict:
        """

        グループのピン投稿を解除します

        """
        return self.PostAPI.delete_group_pin_post(group_id)

    def unpin_post(self, post_id: int) -> dict:
        """

        ピン投稿を削除します

        """
        return self.PostAPI.delete_pin_post(post_id)

    def get_bookmark(
        self,
        user_id: int,
        from_str: Optional[str] = None,
    ) -> PostsResponse:
        """

        ブックマークを取得します

        """
        return self.PostAPI.get_bookmark(user_id, from_str)

    def get_timeline_calls(self, **params) -> PostsResponse:
        """

        誰でも通話を取得します

        Parameters
        ----------

            - group_id: int = None
            - from_timestamp: int = None
            - number: int = None
            - category_id: int = None
            - call_type: str = "voice"
            - include_circle_call: bool = None
            - cross_generation: bool = None
            - exclude_recent_gomimushi: bool = None
            - shared_interest_categories: bool = None

        """
        return self.PostAPI.get_timeline_calls(**params)

    def get_conversation(self, conversation_id: int, **params) -> PostsResponse:
        """

        リプライを含める投稿の会話を取得します

        Parameters
        ----------

            - conversation_id: int
            - group_id: int = None
            - thread_id: int = None
            - from_post_id: int = None
            - number: int = 50
            - reverse: bool = True

        """
        return self.PostAPI.get_conversation(conversation_id, **params)

    def get_conversation_root_posts(self, post_ids: list[int]) -> PostsResponse:
        """

        会話の原点の投稿を取得します

        """
        return self.PostAPI.get_conversation_root_posts(post_ids)

    def get_following_call_timeline(self, **params) -> PostsResponse:
        """

        フォロー中の通話を取得します

        Parameters
        ----------

            - from_timestamp: int = None
            - number: int = None
            - category_id: int = None
            - call_type: str = None
            - include_circle_call: bool = None
            - exclude_recent_gomimushi: bool = None

        """
        return self.PostAPI.get_following_call_timeline(**params)

    def get_following_timeline(self, **params) -> PostsResponse:
        """

        フォロー中のタイムラインを取得します

        Parameters
        ----------

            - from_str: str = None
            - only_root: bool = None
            - order_by: str = None
            - number: int = None
            - mxn: int = None
            - reduce_selfie: bool = None
            - custom_generation_range: bool = None

        """
        return self.PostAPI.get_following_timeline(**params)

    def get_group_highlight_posts(self, group_id: int, **params) -> PostsResponse:
        """

        グループのハイライト投稿を取得します

        Parameters
        ----------

            - group_id: int
            - from_post: int = None
            - number: int = None

        """
        return self.PostAPI.get_group_highlight_posts(group_id, **params)

    def get_group_timeline_by_keyword(
        self, group_id: int, keyword: str, **params
    ) -> PostsResponse:
        """

        グループの投稿をキーワードで検索します

        Parameters
        ----------

            - group_id: int
            - keyword: str
            - from_post_id: int = None
            - number: int = None
            - only_thread_posts: bool = False

        """
        return self.PostAPI.get_group_timeline_by_keyword(group_id, keyword, **params)

    def get_group_timeline(self, group_id: int, **params) -> PostsResponse:
        """

        グループのタイムラインを取得します

        Parameters
        ----------

            - group_id: int
            - from_post_id: int
            - reverse: bool
            - post_type: str
            - number: int
            - only_root: bool

        """
        return self.PostAPI.get_group_timeline(group_id, **params)

    def get_timeline_by_hashtag(self, hashtag: str, **params) -> PostsResponse:
        """

        ハッシュタグでタイムラインを検索します

        Parameters
        ----------

            - hashtag: str - (required)
            - from_post_id: int - (optional)
            - number: int - (optional)

        """
        return self.PostAPI.get_timeline_by_hashtag(hashtag, **params)

    def get_my_posts(self, **params) -> PostsResponse:
        """

        自分の投稿を取得します

        Parameters
        ----------

            - from_post_id: int - (optional)
            - number: int - (optional)
            - include_group_post: bool - (optional)

        """
        return self.PostAPI.get_my_posts(**params)

    def get_post(self, post_id: int) -> Post:
        """

        投稿の詳細を取得します

        """
        return self.PostAPI.get_post(post_id).post

    def get_post_likers(self, post_id: int, **params) -> PostLikersResponse:
        """

        投稿にいいねしたユーザーを取得します

        Parameters
        ----------

            - from_id: int - (optional)
            - number: int - (optional)

        """
        return self.PostAPI.get_post_likers(post_id, **params)

    def get_reposts(self, post_id: int, **params: int) -> PostsResponse:
        """

        投稿の(´∀｀∩)↑age↑を取得します

        Parameters
        ----------

            - post_id: int - (required)
            - from_post_id: int - (optional)
            - number: int - (optional)

        """
        return self.PostAPI.get_post_reposts(post_id, **params)

    def get_posts(self, post_ids: list[int]) -> PostsResponse:
        """

        複数の投稿を取得します

        """
        return self.PostAPI.get_posts(post_ids)

    def get_recommended_post_tags(
        self, tag: Optional[str] = None, save_recent_search: Optional[bool] = False
    ) -> PostTagsResponse:
        """

        おすすめのタグ候補を取得します

        """
        return self.PostAPI.get_recommended_post_tags(tag, save_recent_search)

    def get_recommended_posts(self, **params) -> PostsResponse:
        """

        おすすめの投稿を取得します

        Parameters
        ----------

            - experiment_num: int
            - variant_num: int
            - number: int

        """
        return self.PostAPI.get_recommended_posts(**params)

    def get_timeline_by_keyword(
        self,
        keyword: Optional[str] = None,
        **params,
    ) -> PostsResponse:
        """

        キーワードでタイムラインを検索します

        Parameters
        ----------

            - keyword: str
            - from_post_id: int
            - number: int

        """
        return self.PostAPI.get_timeline_by_keyword(keyword, **params)

    def get_timeline(self, **params) -> PostsResponse:
        """

        タイムラインを取得します

        Parameters
        ----------

            - noreply_mode: bool - (optional)
            - from_post_id: int - (optional)
            - number: int - (optional)
            - order_by: str - (optional)
            - experiment_older_age_rules: bool - (optional)
            - shared_interest_categories: bool - (optional)
            - mxn: int - (optional)
            - en: int - (optional)
            - vn: int - (optional)
            - reduce_selfie: bool - (optional)
            - custom_generation_range: bool - (optional)

        """
        return self.PostAPI.get_timeline(**params)

    def get_url_metadata(self, url: str) -> SharedUrl:
        """

        URLのメタデータを取得します

        """
        return self.PostAPI.get_url_metadata(url)

    def get_user_timeline(self, user_id: int, **params) -> PostsResponse:
        """

        ユーザーのタイムラインを取得します

        Parameters
        ----------

            - from_post_id: int - (optional)
            - number: int - (optional)
            - post_type: str - (optional)

        """
        return self.PostAPI.get_user_timeline(user_id, **params)

    def like(self, post_ids: list[int]) -> LikePostsResponse:
        """

        投稿にいいねします

        ※ 一度にいいねできる投稿数は最大25個

        """
        return self.PostAPI.like_posts(post_ids)

    def remove_bookmark(self, user_id: int, post_id: int) -> dict:
        """

        ブックマークを削除します

        """
        return self.PostAPI.remove_bookmark(user_id, post_id)

    def remove_group_highlight_post(self, group_id: int, post_id: int) -> dict:
        """

        サークルのハイライト投稿を解除します

        """
        return self.PostAPI.remove_group_highlight_post(group_id, post_id)

    def remove_posts(self, post_ids: list[int]) -> dict:
        """

        複数の投稿を削除します

        """
        return self.PostAPI.remove_posts(post_ids)

    def report_post(
        self,
        post_id: int,
        opponent_id: int,
        category_id: int,
        reason: Optional[str] = None,
        screenshot_filename: Optional[str] = None,
        screenshot_2_filename: Optional[str] = None,
        screenshot_3_filename: Optional[str] = None,
        screenshot_4_filename: Optional[str] = None,
    ) -> dict:
        """

        投稿を通報します

        """
        return self.PostAPI.report_post(
            post_id,
            opponent_id,
            category_id,
            reason,
            screenshot_filename,
            screenshot_2_filename,
            screenshot_3_filename,
            screenshot_4_filename,
        )

    def unlike(self, post_id: int) -> dict:
        """

        投稿のいいねを解除します

        """
        return self.PostAPI.unlike_post(post_id)

    def update_post(
        self,
        post_id: int,
        text: Optional[str] = None,
        font_size: Optional[int] = None,
        color: Optional[int] = None,
        message_tags: Optional[list] = [],
    ) -> Post:
        """

        投稿を編集します

        """
        return self.PostAPI.update_post(post_id, text, font_size, color, message_tags)

    def view_video(self, video_id: int) -> dict:
        """

        動画を視聴します

        """
        return self.PostAPI.view_video(video_id)

    def vote_survey(self, survey_id: int, choice_id: int) -> Survey:
        """

        アンケートに投票します

        """
        return self.PostAPI.vote_survey(survey_id, choice_id).survey

    # -REVIEW

    def create_review(self, user_id: int, comment: str) -> dict:
        """

        レターを送信します

        """
        return self.ReviewAPI.create_review(user_id, comment)

    def create_reviews(self, user_ids: list[int], comment: str) -> dict:
        """

        複数人にレターを送信します

        """
        return self.ReviewAPI.create_reviews(user_ids, comment)

    def delete_reviews(self, review_ids: list[int]) -> dict:
        """

        レターを削除します

        """
        return self.ReviewAPI.delete_reviews(review_ids)

    def get_my_reviews(self, **params) -> ReviewsResponse:
        """

        送信したレターを取得します

        Parameters
        ----------

            - from_id: int (optional)
            - number: int = (optional)

        """
        return self.ReviewAPI.get_my_reviews(**params)

    def get_reviews(self, user_id: int, **params) -> ReviewsResponse:
        """

        ユーザーが受け取ったレターを取得します

        Parameters
        ----------

            - user_id: int (required)
            - from_id: int = (optional)
            - number: int = (optional)

        """
        return self.ReviewAPI.get_reviews(user_id, **params)

    def pin_review(self, review_id: int) -> dict:
        """

        レターをピンします

        """
        return self.ReviewAPI.pin_review(review_id)

    def unpin_review(self, review_id: int) -> dict:
        """

        レターのピン止めを解除します

        """
        return self.ReviewAPI.unpin_review(review_id)

    # -THREAD

    def add_post_to_thread(self, post_id: int, thread_id: int) -> ThreadInfo:
        """

        投稿をスレッドに追加します

        """
        return self.ThreadAPI.add_post_to_thread(post_id, thread_id)

    def convert_post_to_thread(
        self,
        post_id: int,
        title: Optional[str] = None,
        thread_icon_filename: Optional[str] = None,
    ) -> ThreadInfo:
        """

        投稿をスレッドに変換します

        """
        return self.ThreadAPI.convert_post_to_thread(
            post_id, title, thread_icon_filename
        )

    def create_thread(
        self,
        group_id: int,
        title: str,
        thread_icon_filename: str,
    ) -> ThreadInfo:
        """

        スレッドを作成します

        """
        return self.ThreadAPI.create_thread(group_id, title, thread_icon_filename)

    def get_group_thread_list(
        self,
        group_id: int,
        from_str: Optional[str] = None,
        **params,
    ) -> GroupThreadListResponse:
        """

        グループのスレッド一覧を取得します

        Parameters
        ----------

            - group_id: int
            - from_str: str = None
            - join_status: str = None

        """
        return self.ThreadAPI.get_group_thread_list(group_id, from_str, **params)

    def get_thread_joined_statuses(self, ids: list[int]) -> dict:
        """

        スレッド参加ステータスを取得します

        """
        return self.ThreadAPI.get_thread_joined_statuses(ids)

    def get_thread_posts(
        self,
        thread_id: int,
        from_str: Optional[str] = None,
        **params,
    ) -> PostsResponse:
        """

        スレッドの投稿を取得します

        Parameters
        ----------

            - post_type: str
            - number: int = None
            - from_str: str = None

        """
        return self.ThreadAPI.get_thread_posts(thread_id, from_str, **params)

    def join_thread(self, thread_id: int, user_id: int) -> dict:
        """

        スレッドに参加します

        """
        return self.ThreadAPI.join_thread(thread_id, user_id)

    def leave_thread(self, thread_id: int, user_id: int) -> dict:
        """

        スレッドから脱退します

        """
        return self.ThreadAPI.leave_thread(thread_id, user_id)

    def remove_thread(self, thread_id: int) -> dict:
        """

        スレッドを削除します

        """
        return self.ThreadAPI.remove_thread(thread_id)

    def update_thread(
        self,
        thread_id: int,
        title: str,
        thread_icon_filename: str,
    ) -> dict:
        """

        スレッドを編集します

        """
        return self.ThreadAPI.update_thread(thread_id, title, thread_icon_filename)

    # -USER

    def delete_footprint(self, user_id: int, footprint_id: int) -> dict:
        """

        足跡を削除します

        """
        return self.UserAPI.delete_footprint(user_id, footprint_id)

    def destroy_user(self) -> dict:
        """

        アカウントを削除します

        """
        answer = input("Are you sure you want to delete your account? Y/N")
        if answer.lower() != "y":
            return
        return self.UserAPI.destroy_user()

    def follow_user(self, user_id: int) -> dict:
        """

        ユーザーをフォローします

        """
        return self.UserAPI.follow_user(user_id)

    def follow_users(self, user_ids: list[int]) -> dict:
        """

        複数のユーザーをフォローします

        """
        return self.UserAPI.follow_users(user_ids)

    def get_active_followings(self, **params) -> ActiveFollowingsResponse:
        """

        アクティブなフォロー中のユーザーを取得します

        Parameters
        ----------

            - only_online: bool
            - from_loggedin_at: int = None

        """
        return self.UserAPI.get_active_followings(**params)

    def get_follow_recommendations(self, **params) -> FollowRecommendationsResponse:
        """

        フォローするのにおすすめのユーザーを取得します

        Parameters
        ----------

            - from_timestamp: int = None,
            - number: int = None,
            - sources: list[str] = None

        """
        return self.UserAPI.get_follow_recommendations(**params)

    def get_follow_request(
        self, from_timestamp: Optional[int] = None
    ) -> UsersByTimestampResponse:
        """

        フォローリクエストを取得します

        """
        return self.UserAPI.get_follow_request(from_timestamp)

    def get_follow_request_count(self) -> int:
        """

        フォローリクエストの数を取得します

        """
        return self.UserAPI.get_follow_request_count().users_count

    def get_following_users_born(
        self, birthdate: Optional[int] = None
    ) -> UsersResponse:
        """

        フォロー中のユーザーの誕生日を取得します

        """
        return self.UserAPI.get_following_users_born(birthdate)

    def get_footprints(self, **params) -> list[Footprint]:
        """

        足跡を取得します

        Parameters
        ----------

            - from_id: int = None
            - number: int = None
            - mode: str = None

        """
        return self.UserAPI.get_footprints(**params).footprints

    def get_fresh_user(self, user_id: int) -> UserResponse:
        """

        認証情報などを含んだユーザー情報を取得します

        """
        return self.UserAPI.get_fresh_user(user_id)

    def get_hima_users(self, **params) -> list[UserWrapper]:
        """

        暇なユーザーを取得します

        Parameters
        ----------

            - from_hima_id: int = None
            - number: int = None

        """
        return self.UserAPI.get_hima_users(**params).hima_users

    def get_user_ranking(self, mode: str) -> RankingUsersResponse:
        """

        ユーザーのフォロワーランキングを取得します

        Examples
        --------

        >>> ルーキーを取得する場合:

        >>> api.get_user_ranking(mode="one_month")

        ---

        >>> ミドルを取得する場合:

        >>> api.get_user_ranking(mode="six_months")

        ---

        >>> 殿堂入りを取得する場合:

        >>> api.get_user_ranking(mode="all_time")

        """
        return self.UserAPI.get_user_ranking(mode)

    def get_refresh_counter_requests(self) -> RefreshCounterRequestsResponse:
        """

        カウンター更新のリクエストを取得します

        """
        return self.UserAPI.get_refresh_counter_requests()

    def get_social_shared_users(self, **params) -> SocialShareUsersResponse:
        """

        SNS共有をしたユーザーを取得します

        Parameters
        ----------

            - sns_name: str - (Required)
            - number: int - (Optional)
            - from_id: int - (Optional)

        """
        return self.UserAPI.get_social_shared_users(**params)

    def get_timestamp(self) -> UserTimestampResponse:
        """

        タイムスタンプを取得します

        """
        return self.UserAPI.get_timestamp()

    def get_user(self, user_id: int) -> User:
        """

        ユーザーの情報を取得します

        """
        return self.UserAPI.get_user(user_id).user

    def get_user_email(self, user_id: int) -> str:
        """

        ユーザーのメールアドレスを取得します

        """
        return self.UserAPI.get_user_email(user_id).email

    def get_user_followers(self, user_id: int, **params) -> FollowUsersResponse:
        """

        ユーザーのフォロワーを取得します

        Parameters
        ----------

            - user_id: int
            - from_follow_id: int = None
            - followed_by_me: int = None

        """
        return self.UserAPI.get_user_followers(user_id, **params)

    def get_user_followings(self, user_id: int, **params) -> FollowUsersResponse:
        """

        フォロー中のユーザーを取得します

        Parameters
        ----------

            - user_id: int
            - from_follow_id: int = None
            - from_timestamp: int = None
            - order_by: str = None

        """
        return self.UserAPI.get_user_followings(user_id, **params)

    def get_user_from_qr(self, qr: str) -> UserResponse:
        """

        QRコードからユーザーを取得します

        """
        return self.UserAPI.get_user_from_qr(qr)

    def get_user_without_leaving_footprint(self, user_id: int) -> UserResponse:
        """

        足跡をつけずにユーザーの情報を取得します

        """
        return self.UserAPI.get_user_without_leaving_footprint(user_id)

    def get_users(self, user_ids: list[int]) -> UsersResponse:
        """

        複数のユーザーの情報を取得します

        """
        return self.UserAPI.get_users(user_ids)

    def refresh_counter(self, counter: str) -> dict:
        """

        カウンターを更新します

        """
        return self.UserAPI.refresh_counter(counter)

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
        """

        Register user

        """
        return self.UserAPI.register(
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

    def remove_user_avatar(self) -> dict:
        """

        ユーザーのアイコンを削除します

        """
        return self.UserAPI.remove_user_avatar()

    def remove_user_cover(self) -> dict:
        """

        ユーザーのカバー画像を削除します

        """
        return self.UserAPI.remove_user_cover()

    def report_user(
        self,
        user_id: int,
        category_id: int,
        reason: Optional[str] = None,
        screenshot_filename: Optional[str] = None,
        screenshot_2_filename: Optional[str] = None,
        screenshot_3_filename: Optional[str] = None,
        screenshot_4_filename: Optional[str] = None,
    ) -> dict:
        """

        ユーザーを通報します

        """
        return self.UserAPI.report_user(
            user_id,
            category_id,
            reason,
            screenshot_filename,
            screenshot_2_filename,
            screenshot_3_filename,
            screenshot_4_filename,
        )

    def reset_password(
        self,
        email: str,
        email_grant_token: str,
        password: str,
    ) -> dict:
        """

        パスワードをリセットします

        """
        return self.UserAPI.reset_password(email, email_grant_token, password)

    def search_lobi_users(self, **params) -> UsersResponse:
        """

        Lobiのユーザーを検索します

        Parameters
        ----------

            - nickname: str = None
            - number: int = None
            - from_str: str = None

        """
        return self.UserAPI.search_lobi_users(**params)

    def search_users(self, **params) -> UsersResponse:
        """

        ユーザーを検索します

        Parameters
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
        return self.UserAPI.search_users(**params)

    def set_follow_permission_enabled(
        self,
        nickname: str,
        is_private: Optional[bool] = None,
    ) -> dict:
        """

        フォローを許可制に設定します

        """
        return self.UserAPI.set_follow_permission_enabled(nickname, is_private)

    def take_action_follow_request(self, target_id: int, action: str) -> dict:
        return self.UserAPI.take_action_follow_request(target_id, action)

    def turn_on_hima(self) -> dict:
        """

        ひまなう

        """
        return self.UserAPI.turn_on_hima()

    def unfollow_user(self, user_id: int) -> dict:
        """

        ユーザーをアンフォローします

        """
        return self.UserAPI.unfollow_user(user_id)

    def update_user(self, nickname: str, **params) -> dict:
        """

        プロフィールを更新します

        Parameters
        ----------

            - nickname: str = (required)
            - biography: str = (optional)
            - prefecture: str = (optional)
            - gender: int = (optional)
            - country_code: str = (optional)
            - profile_icon_filename: str = (optional)
            - cover_image_filename: str = (optional)
            - username: str = (optional)
            - : str = (optional)

        """
        return self.UserAPI.update_user(nickname, **params)

    def block_user(self, user_id: int) -> dict:
        """

        ユーザーをブロックします

        """
        return self.UserAPI.block_user(user_id)

    def get_blocked_user_ids(self) -> BlockedUserIdsResponse:
        """

        あなたをブロックしたユーザーを取得します

        """
        return self.UserAPI.get_blocked_user_ids()

    def get_blocked_users(self, from_id: Optional[int] = None) -> BlockedUsersResponse:
        """

        ブロックしたユーザーを取得します

        """
        return self.UserAPI.get_blocked_users(from_id)

    def unblock_user(self, user_id: int) -> dict:
        """

        ユーザーをアンブロックします

        """
        return self.UserAPI.unblock_user(user_id)

    def get_hidden_users_list(self, **params) -> HiddenResponse:
        """

        非表示のユーザー一覧を取得します

        Parameters
        ----------

            - from: str = None
            - number: int = None

        """
        return self.UserAPI.get_hidden_users_list(**params)

    def hide_user(self, user_id: int) -> dict:
        """

        ユーザーを非表示にします

        """
        return self.UserAPI.hide_user(user_id)

    def unhide_users(self, user_ids: list[int]) -> dict:
        """

        ユーザーの非表示を解除します

        """
        return self.UserAPI.unhide_users(user_ids)
