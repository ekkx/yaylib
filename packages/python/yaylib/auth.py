# PORTING.md §6: login_with_email is the ONLY wrapped auth endpoint —
# it is the cache boundary, so "log in or use the cached session" is one
# call. Other auth-flow body endpoints (create_user, edit_user,
# login_with_sns, oauth token / refresh) stay as plain calls on the
# generated API classes; wrapping each would multiply surface without
# changing behaviour.

from __future__ import annotations

import json
from typing import Optional

from yaylib._error_codes import ErrCodeUnauthorized
from yaylib.errors import as_api_error, code_of
from yaylib.exceptions import ApiException
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
    # Restore which key refreshes this session (empty → app key) so a
    # cached web session refreshes with the web key after a restart.
    client.set_oauth_key(session.oauth_key)


async def _login_with_email_web(
    client,
    *,
    email: str,
    password: str,
    uuid: str,
    two_fa_code: Optional[str],
) -> Optional[LoginUserResponse]:
    """Fallback login against the browser endpoint (POST
    /v2/users/login_with_email) using the client's ``web_api_key``.

    Sent through the UNWRAPPED transport (``send_unwrapped``): no header
    injection, no Bearer, no 401 refresh — a login must authenticate
    purely by its body credentials (PORTING.md §6.1/§6.2). Only the
    body's api_key and the endpoint version differ from the app login.

    Returns the decoded response on success, or ``None`` when the
    fallback is itself rejected (non-2xx) so the caller surfaces the
    primary error. Raises an ``ApiException`` carrying the raw body when
    a 2xx response fails to decode (§15), rather than masking it behind
    the primary 401.
    """
    payload = {
        "email": email,
        "password": password,
        "api_key": client.web_api_key,
        "uuid": uuid,
    }
    if two_fa_code:
        payload["two_fa_code"] = two_fa_code

    # Present as the browser client. The web version string is not
    # otherwise known to the SDK, so app_version stands in for it.
    web_ua = (
        "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) "
        "AppleWebKit/537.36 (KHTML, like Gecko) "
        "Chrome/120.0.0.0 Safari/537.36"
    )
    headers = {
        "Content-Type": "application/json;charset=UTF-8",
        "User-Agent": web_ua,
        "Agent": f"YayWeb {client.app_version}",
        "X-Device-Info": f"Yay {client.app_version} Web ({web_ua})",
    }
    url = f"{client.base_url}/v2/users/login_with_email"
    try:
        res = await client._transport.send_unwrapped(
            "POST", url, headers, json.dumps(payload)
        )
    except Exception:  # noqa: BLE001
        return None
    if not (200 <= res.status <= 299):
        return None
    raw = res.data.decode("utf-8", "replace")
    try:
        return LoginUserResponse.from_json(raw)
    except Exception as err:  # noqa: BLE001
        # A 2xx we couldn't decode is a genuine login result: surface it
        # with the raw body, not the primary 401.
        raise ApiException(status=res.status, reason=res.reason, body=raw) from err


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

    login_key = api_key or client.api_key
    body = LoginEmailUserRequest(
        api_key=login_key,
        email=email,
        password=password,
        uuid=uuid or client.device_uuid,
        two_f_a_code=two_fa_code,
    )
    # The key that establishes the session also refreshes its tokens;
    # the fallback switches it to web_api_key below.
    oauth_key = login_key

    try:
        resp = await client.users_api.login_with_email(
            login_email_user_request=body
        )
    except Exception as err:  # noqa: BLE001
        api_err = as_api_error(err)
        # Fallback (PORTING.md §6.2): a plain "unauthorized" retries
        # against the browser login endpoint with web_api_key. Skipped
        # for other 401s — notably Required2FA (the caller must supply a
        # code, not switch keys) — and when the caller pinned a specific
        # api_key, or when web_api_key is empty (the off switch).
        caller_pinned = api_key is not None and api_key != client.api_key
        if (
            client.web_api_key
            and not caller_pinned
            and code_of(api_err) == ErrCodeUnauthorized
        ):
            fallback = await _login_with_email_web(
                client,
                email=email,
                password=password,
                uuid=uuid or client.device_uuid,
                two_fa_code=two_fa_code,
            )
            if fallback is None:
                # Fallback itself rejected: surface the primary error.
                raise api_err
            client._logger.debug(
                "web login fallback ok",
                extra={"event": "login_fallback", "outcome": "ok"},
            )
            resp = fallback
            # This session was established with the web key, so its token
            # refresh must use it too.
            oauth_key = client.web_api_key
        else:
            raise api_err

    access = resp.access_token or ""
    refresh = resp.refresh_token or ""
    user_id = resp.user_id or 0
    client.set_tokens(access, refresh)
    client.set_login_identity(email, user_id)
    client.set_oauth_key(oauth_key)

    if client.session_store is not None:
        session = Session(
            email=email,
            user_id=user_id,
            access_token=access,
            refresh_token=refresh,
            device_uuid=client.device_uuid,
            oauth_key=client._session_oauth_key(),
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
