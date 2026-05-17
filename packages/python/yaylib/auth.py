# PORTING.md §6: login_with_email is the ONLY wrapped auth endpoint —
# it is the cache boundary, so "log in or use the cached session" is one
# call. Other auth-flow body endpoints (create_user, edit_user,
# login_with_sns, oauth token / refresh) stay as plain calls on the
# generated API classes; wrapping each would multiply surface without
# changing behaviour.

from __future__ import annotations

from typing import Optional

from yaylib.errors import as_api_error
from yaylib.models.login_email_user_request import LoginEmailUserRequest
from yaylib.models.login_user_response import LoginUserResponse
from yaylib.session import NoSessionError, Session, SessionSaveFailed


async def login_with_email(
    client,
    *,
    email: str,
    password: str,
    api_key: Optional[str] = None,
    uuid: Optional[str] = None,
    two_fa_code: Optional[str] = None,
    no_cache: bool = False,
) -> LoginUserResponse:
    """Wrapped login with transparent session caching, called with
    keyword arguments — the same call shape as every other generated
    Python operation, not a builder. The generated layer accepts only a
    single ``login_email_user_request`` body; this wrapper takes the
    fields directly and auto-fills api_key / uuid from the client.
    ``no_cache=True`` forces a fresh HTTP login.
    """
    return await _execute_login(
        client,
        email=email,
        password=password,
        api_key=api_key,
        uuid=uuid,
        two_fa_code=two_fa_code,
        bypass_cache=no_cache,
    )


def _synthesize_login_response(session: Session) -> LoginUserResponse:
    # Same shape regardless of cache hit vs. miss. Fields the server
    # returns but the cache doesn't carry default to None.
    return LoginUserResponse(
        access_token=session.access_token,
        refresh_token=session.refresh_token,
        user_id=session.user_id,
    )


def _activate_cached_session(client, session: Session) -> None:
    client.set_tokens(session.access_token, session.refresh_token)
    if session.device_uuid:
        client.set_device_uuid(session.device_uuid)
    client.set_login_identity(session.email, session.user_id)


async def _execute_login(
    client,
    *,
    email: str,
    password: str,
    api_key: Optional[str],
    uuid: Optional[str],
    two_fa_code: Optional[str],
    bypass_cache: bool,
) -> LoginUserResponse:
    if not email or not password:
        raise ValueError(
            "yaylib: login_with_email requires email and password"
        )

    # Cache lookup — only with a session store configured and no
    # no_cache opt-out.
    if client.session_store is not None and not bypass_cache:
        try:
            cached = await client.session_store.load(email)
        except NoSessionError:
            cached = None  # cache miss → fall through to HTTP login
        # Any other error (rate-limit, disk failure) propagates: the
        # caller must not silently drop it on the wire-bound load path.
        if cached is not None:
            _activate_cached_session(client, cached)
            return _synthesize_login_response(cached)

    body = LoginEmailUserRequest(
        api_key=api_key or client.api_key,
        email=email,
        password=password,
        uuid=uuid or client.device_uuid,
        two_f_a_code=two_fa_code,
    )

    try:
        resp = await client.users_api.login_with_email(
            login_email_user_request=body
        )
    except Exception as err:  # noqa: BLE001
        raise as_api_error(err)

    access = resp.access_token or ""
    refresh = resp.refresh_token or ""
    user_id = resp.user_id or 0
    client.set_tokens(access, refresh)
    client.set_login_identity(email, user_id)

    if client.session_store is not None:
        session = Session(
            email=email,
            user_id=user_id,
            access_token=access,
            refresh_token=refresh,
            device_uuid=client.device_uuid,
            updated_at="",
        )
        try:
            await client.session_store.save(session)
        except SessionSaveFailed:
            raise
        except Exception as err:  # noqa: BLE001
            # PORTING.md §5: persist failure here surfaces as
            # SessionSaveFailed, but the tokens are already active so the
            # caller can choose to proceed.
            raise SessionSaveFailed(err)

    return resp
