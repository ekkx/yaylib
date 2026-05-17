# PORTING.md §2 / §3: Client is the single entry point. Every operation
# is reachable via the per-tag API instances attached to the client
# (``client.users_api.get_user(...)`` etc.); they all share one
# ApiClient so the headers / refresh / retry transport is applied
# uniformly. A flat promotion layer (PORTING.md §2) is generated in a
# later step — for now the per-service handles are the stable surface,
# matching the TypeScript port's ``client.usersAPI`` form.

from __future__ import annotations

import asyncio
import contextlib
import json
import logging
import uuid as _uuid
from typing import Optional
from urllib.parse import urlencode

from yaylib.api.activities_api import ActivitiesApi
from yaylib.api.apps_api import AppsApi
from yaylib.api.auth_api import AuthApi
from yaylib.api.buckets_api import BucketsApi
from yaylib.api.calls_api import CallsApi
from yaylib.api.chat_rooms_api import ChatRoomsApi
from yaylib.api.conversations_api import ConversationsApi
from yaylib.api.email_verification_urls_api import EmailVerificationUrlsApi
from yaylib.api.games_api import GamesApi
from yaylib.api.genres_api import GenresApi
from yaylib.api.gifts_api import GiftsApi
from yaylib.api.group_mute_api import GroupMuteApi
from yaylib.api.groups_api import GroupsApi
from yaylib.api.hidden_api import HiddenApi
from yaylib.api.moderation_api import ModerationApi
from yaylib.api.notification_settings_api import NotificationSettingsApi
from yaylib.api.pinned_api import PinnedApi
from yaylib.api.posts_api import PostsApi
from yaylib.api.received_gifts_api import ReceivedGiftsApi
from yaylib.api.sticker_packs_api import StickerPacksApi
from yaylib.api.surveys_api import SurveysApi
from yaylib.api.threads_api import ThreadsApi
from yaylib.api.users_api import UsersApi
from yaylib.api_client import ApiClient
from yaylib.configuration import Configuration

from yaylib._config import (
    DEFAULT_ACCEPT_LANGUAGE,
    DEFAULT_API_KEY,
    DEFAULT_API_VERSION_KEY,
    DEFAULT_API_VERSION_NAME,
    DEFAULT_APP_VERSION,
    DEFAULT_BASE_URL,
    DEFAULT_CASSANDRA_BASE_URL,
    DEFAULT_CONNECTION_SPEED,
    DEFAULT_CONNECTION_TYPE,
    DEFAULT_DEVICE_DENSITY,
    DEFAULT_DEVICE_MODEL,
    DEFAULT_DEVICE_OS,
    DEFAULT_DEVICE_SCREEN,
    DEFAULT_DEVICE_TYPE,
    DEFAULT_EVENT_STREAM_URL,
    build_device_info,
    build_user_agent,
)
from yaylib.retry import DEFAULT_RETRY_POLICY, RetryPolicy
from yaylib.session import Session, SessionStore
from yaylib.tokens import Tokens, empty_tokens
from yaylib.transport import Transport, TransportContext

def _default_logger() -> logging.Logger:
    """The silent default. PORTING.md §12.2: the SDK is a library — with
    no logger configured it produces zero output. A NullHandler discards
    every record and propagation is off so nothing reaches the root
    logger. Callers opt in by passing ``logger=`` to Client.
    """
    lg = logging.getLogger("yaylib")
    if not any(isinstance(h, logging.NullHandler) for h in lg.handlers):
        lg.addHandler(logging.NullHandler())
    lg.propagate = False
    return lg


def new_uuid_v4() -> str:
    """Canonical 8-4-4-4-12 hex UUID v4 from the OS CSPRNG. PORTING.md
    §14: a failure of the randomness primitive MUST raise — never
    degrade to a weak fallback (``uuid.uuid4`` is backed by os.urandom
    and raises on failure).
    """
    return str(_uuid.uuid4())


class _WrappedApiClient(ApiClient):
    """ApiClient whose HTTP layer is our Transport. The generated
    ``__init__`` builds an aiohttp session eagerly (and so requires a
    running loop); we bypass it so Client stays synchronously
    constructible.
    """

    def __init__(self, configuration: Configuration, transport: Transport):
        self.configuration = configuration
        self.rest_client = transport
        # No default User-Agent: the transport composes and injects the
        # Yay!-required one (PORTING.md §3.1). A non-empty default here
        # would win via set-if-absent and pin the wrong value.
        self.default_headers = {}
        self.cookie = None
        self.client_side_validation = configuration.client_side_validation


class Client:
    def __init__(
        self,
        *,
        base_url: Optional[str] = None,
        cassandra_base_url: Optional[str] = None,
        event_stream_url: Optional[str] = None,
        api_key: Optional[str] = None,
        api_version_key: Optional[str] = None,
        api_version_name: Optional[str] = None,
        app_version: Optional[str] = None,
        user_agent: Optional[str] = None,
        device_info: Optional[str] = None,
        device_uuid: Optional[str] = None,
        connection_type: Optional[str] = None,
        connection_speed: Optional[str] = None,
        accept_language: Optional[str] = None,
        session_store: Optional[SessionStore] = None,
        retry_policy: Optional[RetryPolicy] = None,
        logger: Optional[logging.Logger] = None,
    ) -> None:
        self.base_url = base_url or DEFAULT_BASE_URL
        self.cassandra_base_url = cassandra_base_url or DEFAULT_CASSANDRA_BASE_URL
        self.event_stream_url = event_stream_url or DEFAULT_EVENT_STREAM_URL
        self.api_key = api_key or DEFAULT_API_KEY
        self.api_version_key = api_version_key or DEFAULT_API_VERSION_KEY
        self.api_version_name = api_version_name or DEFAULT_API_VERSION_NAME
        self.app_version = app_version or DEFAULT_APP_VERSION
        self.user_agent = user_agent or build_user_agent(
            device_type=DEFAULT_DEVICE_TYPE,
            device_os=DEFAULT_DEVICE_OS,
            device_density=DEFAULT_DEVICE_DENSITY,
            device_screen=DEFAULT_DEVICE_SCREEN,
            device_model=DEFAULT_DEVICE_MODEL,
        )
        self.device_info = device_info or build_device_info(
            self.app_version, self.user_agent
        )
        self._device_uuid = device_uuid or new_uuid_v4()
        self.connection_type = connection_type or DEFAULT_CONNECTION_TYPE
        self.connection_speed = connection_speed or DEFAULT_CONNECTION_SPEED
        self.accept_language = accept_language or DEFAULT_ACCEPT_LANGUAGE
        self.session_store = session_store
        self.retry_policy = retry_policy or DEFAULT_RETRY_POLICY
        self._logger = logger if logger is not None else _default_logger()

        self._tokens: Tokens = empty_tokens()
        self._user_id = 0
        self._current_email = ""
        self._client_ip = ""
        self._client_ip_fetching = False
        self._bg_tasks: "set[asyncio.Task]" = set()
        self._event_streams: set = set()
        self._refresh_inflight: Optional["asyncio.Future[bool]"] = None

        transport = Transport(
            ctx=TransportContext(
                user_agent=self.user_agent,
                app_version=self.api_version_name,
                device_info=self.device_info,
                device_uuid=self._device_uuid,
                connection_type=self.connection_type,
                connection_speed=self.connection_speed,
                accept_language=self.accept_language,
                api_key=self.api_key,
                client_ip=lambda: self._client_ip,
                access_token=lambda: self._tokens.access,
                cassandra_base_url=self.cassandra_base_url,
                logger=self._logger,
            ),
            refresh=self._try_refresh,
            policy=self.retry_policy,
            on_response=self._maybe_fetch_client_ip,
        )
        self._transport = transport
        config = Configuration(host=self.base_url)
        api_client = _WrappedApiClient(config, transport)
        self._api_client = api_client

        self.activities_api = ActivitiesApi(api_client)
        self.apps_api = AppsApi(api_client)
        self.auth_api = AuthApi(api_client)
        self.buckets_api = BucketsApi(api_client)
        self.calls_api = CallsApi(api_client)
        self.chat_rooms_api = ChatRoomsApi(api_client)
        self.conversations_api = ConversationsApi(api_client)
        self.email_verification_urls_api = EmailVerificationUrlsApi(api_client)
        self.games_api = GamesApi(api_client)
        self.genres_api = GenresApi(api_client)
        self.gifts_api = GiftsApi(api_client)
        self.group_mute_api = GroupMuteApi(api_client)
        self.groups_api = GroupsApi(api_client)
        self.hidden_api = HiddenApi(api_client)
        self.moderation_api = ModerationApi(api_client)
        self.notification_settings_api = NotificationSettingsApi(api_client)
        self.pinned_api = PinnedApi(api_client)
        self.posts_api = PostsApi(api_client)
        self.received_gifts_api = ReceivedGiftsApi(api_client)
        self.sticker_packs_api = StickerPacksApi(api_client)
        self.surveys_api = SurveysApi(api_client)
        self.threads_api = ThreadsApi(api_client)
        self.users_api = UsersApi(api_client)

    # ---- tokens / identity (PORTING.md §4) ----

    @property
    def tokens(self) -> Tokens:
        # Frozen dataclass — already a value snapshot; mutating it cannot
        # reach back into the Client.
        return self._tokens

    def set_tokens(self, access: str, refresh: str) -> None:
        self._tokens = Tokens(access=access, refresh=refresh)

    @property
    def device_uuid(self) -> str:
        return self._device_uuid

    @property
    def user_id(self) -> int:
        return self._user_id

    @property
    def current_email(self) -> str:
        return self._current_email

    @property
    def client_ip(self) -> str:
        return self._client_ip

    def set_login_identity(self, email: str, user_id: int) -> None:
        """Internal — used by auth.py when activating a session."""
        self._current_email = email
        self._user_id = user_id

    def set_device_uuid(self, uuid: str) -> None:
        """Internal — used by auth.py when a restored session carries a
        persisted device UUID (PORTING.md §14)."""
        self._device_uuid = uuid
        self._transport._ctx.device_uuid = uuid

    # ---- lifecycle ----

    async def __aenter__(self) -> "Client":
        return self

    async def __aexit__(self, exc_type, exc, tb) -> None:
        await self.close()

    def open_event_stream(self, opts=None):
        """Open a multiplexed real-time event stream (PORTING.md §10).

        Usable both as an async context manager and as a coroutine::

            async with client.open_event_stream() as stream:
                sub = await stream.subscribe(chat_room_channel())

            stream = await client.open_event_stream(opts)
            ...
            await stream.close()
        """
        from yaylib.event_stream import _OpenEventStreamCM

        return _OpenEventStreamCM(self, opts)

    async def close(self) -> None:
        # PORTING.md §3: cancel background work owned by the Client — the
        # lazy X-Client-IP fetch and every open event stream.
        for stream in list(self._event_streams):
            with contextlib.suppress(Exception):
                await stream.close()
        self._event_streams.clear()
        tasks = list(self._bg_tasks)
        for task in tasks:
            task.cancel()
        for task in tasks:
            try:
                await task
            except (asyncio.CancelledError, Exception):  # noqa: BLE001
                pass
        self._bg_tasks.clear()
        await self._transport.close()

    # ---- lazy X-Client-IP (PORTING.md §12) ----

    def _maybe_fetch_client_ip(self) -> None:
        """Single-flight: no-op when the IP is already known or a fetch
        is in flight; otherwise spawn a background fetch. On failure the
        field is left empty and the next request retries.
        """
        if self._client_ip or self._client_ip_fetching:
            return
        self._client_ip_fetching = True
        task = asyncio.ensure_future(self._fetch_client_ip())
        self._bg_tasks.add(task)
        task.add_done_callback(self._bg_tasks.discard)

    async def _fetch_client_ip(self) -> None:
        try:
            resp = await asyncio.wait_for(
                self.users_api.get_user_timestamp(), timeout=30
            )
            ip = getattr(resp, "ip_address", "") or ""
            if ip:
                self._client_ip = ip
        except asyncio.CancelledError:
            raise
        except Exception:  # noqa: BLE001
            # Leave the field empty; the next request retries the fetch.
            pass
        finally:
            self._client_ip_fetching = False

    # ---- auth wrapper (PORTING.md §6) ----

    def login_with_email(self):
        """LoginWithEmail with transparent session caching. A hit returns
        the persisted session synthesized into a LoginUserResponse with
        no HTTP; a miss (or ``.no_cache()``) issues the OAuth login,
        persists the session, and activates tokens / user id / email.
        """
        from yaylib.auth import login_with_email

        return login_with_email(self)

    # ---- 401 auto-refresh (PORTING.md §6.1) ----

    async def _try_refresh(self, stale_access: str) -> bool:
        # Single-flight: concurrent 401s collapse onto one OAuth /token
        # round-trip; the others await this same future.
        if self._refresh_inflight is not None:
            return await self._refresh_inflight
        loop = asyncio.get_event_loop()
        fut: "asyncio.Future[bool]" = loop.create_future()
        self._refresh_inflight = fut
        try:
            result = await self._do_refresh(stale_access)
            fut.set_result(result)
            return result
        except BaseException as exc:  # noqa: BLE001
            fut.set_exception(exc)
            raise
        finally:
            self._refresh_inflight = None

    async def _do_refresh(self, stale_access: str) -> bool:
        # A concurrent caller already rotated the token while we queued;
        # skip the round-trip and let the caller retry with the current
        # Bearer.
        if self._tokens.access != stale_access:
            return True
        if not self._tokens.refresh:
            self._logger.debug(
                "token refresh skipped",
                extra={"event": "token_refresh", "outcome": "no_token"},
            )
            return False

        url = f"{self.base_url}/api/v1/oauth/token"
        import base64

        basic = base64.b64encode(self.api_key.encode("utf-8")).decode("ascii")
        form = urlencode(
            {
                "grant_type": "refresh_token",
                "refresh_token": self._tokens.refresh,
            }
        )
        try:
            res = await self._transport.send_unwrapped(
                "POST",
                url,
                {
                    "Authorization": f"Basic {basic}",
                    "Content-Type": "application/x-www-form-urlencoded",
                },
                form,
            )
        except Exception:  # noqa: BLE001
            self._logger.debug(
                "token refresh failed",
                extra={"event": "token_refresh", "outcome": "failed"},
            )
            return False

        if not (200 <= res.status <= 299):
            self._logger.debug(
                "token refresh failed",
                extra={"event": "token_refresh", "outcome": "failed"},
            )
            return False
        try:
            parsed = json.loads(res.data)
        except (ValueError, TypeError):
            self._logger.debug(
                "token refresh failed",
                extra={"event": "token_refresh", "outcome": "failed"},
            )
            return False
        access = parsed.get("access_token")
        refresh = parsed.get("refresh_token")
        if not access or not refresh:
            self._logger.debug(
                "token refresh failed",
                extra={"event": "token_refresh", "outcome": "failed"},
            )
            return False

        # The server rotates BOTH tokens; persist the new refresh token.
        self._tokens = Tokens(access=access, refresh=refresh)
        self._logger.debug(
            "token refresh ok",
            extra={"event": "token_refresh", "outcome": "ok"},
        )

        if self.session_store is not None and self._current_email:
            session = Session(
                email=self._current_email,
                user_id=self._user_id,
                access_token=self._tokens.access,
                refresh_token=self._tokens.refresh,
                device_uuid=self._device_uuid,
                updated_at="",
            )
            try:
                await self.session_store.save(session)
            except Exception as exc:  # noqa: BLE001
                # PORTING.md §5/§12.2: persist failure on the refresh
                # path is non-fatal — the rotated tokens are still
                # active. Surface it via the injected logger.
                self._logger.warning(
                    "persist refreshed tokens failed",
                    extra={
                        "event": "token_persist_fail",
                        "user_id": self._user_id,
                        "err": str(exc),
                    },
                )

        return True
