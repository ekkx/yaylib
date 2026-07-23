# PORTING.md §6: login_with_email is the ONLY wrapped auth endpoint —
# it is the cache boundary, so "log in or use the cached session" is one
# call. Other auth-flow body endpoints (create_user, edit_user,
# login_with_sns, oauth token / refresh) stay as plain calls on the
# generated API classes; wrapping each would multiply surface without
# changing behaviour.

from __future__ import annotations

import json
from typing import Optional, Tuple

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
    # Delegate to the client's single activation path so the cache-restore
    # flow and the public load_session / save_session never diverge
    # (PORTING.md §6.2: restore the establishing key in BOTH paths).
    client._apply_session(session)


async def _login_with_email_app(
    client,
    *,
    api_key: str,
    email: str,
    password: str,
    uuid: str,
    two_fa_code: Optional[str],
) -> LoginUserResponse:
    """App login (POST /v3/users/login_with_email) via the wrapped
    generated client — header injection, retry, host routing all apply.
    ``api_key`` carries the key the body authenticates under.
    """
    body = LoginEmailUserRequest(
        api_key=api_key,
        email=email,
        password=password,
        uuid=uuid,
        two_f_a_code=two_fa_code,
    )
    return await client.users_api.login_with_email(login_email_user_request=body)


async def _login_with_email_web(
    client,
    *,
    email: str,
    password: str,
    uuid: str,
    two_fa_code: Optional[str],
) -> LoginUserResponse:
    """Primary login against the browser endpoint (POST
    /v2/users/login_with_email) using the client's ``web_api_key``.

    Sent through the UNWRAPPED transport (``send_unwrapped``): no header
    injection, no Bearer, no 401 refresh — a login must authenticate
    purely by its body credentials (PORTING.md §6.1/§6.2). Only the
    body's api_key and the endpoint version differ from the app login.

    Failures surface as the SAME typed error the generated clients raise
    (``ApiException`` / its status subclasses) with the raw body
    attached, so ``code_of`` / ``error_response_of`` / ``except
    APIError`` work uniformly now that this is the primary path
    (PORTING.md §6.2). A non-2xx response becomes a typed error carrying
    the body; a 2xx response that fails to decode is a genuine login
    result surfaced with its raw body (§15), not masked as a 401.
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

    # A transport-level failure (aiohttp error) surfaces from send_unwrapped
    # as a status-0 ApiException; it is not ErrCodeUnauthorized, so the
    # caller surfaces it directly without falling back to the app login.
    res = await client._transport.send_unwrapped(
        "POST", url, headers, json.dumps(payload)
    )

    raw = res.data.decode("utf-8", "replace")
    if not (200 <= res.status <= 299):
        # Same typed error the generated client raises for a non-2xx,
        # carrying the raw body so code_of / error_response_of work.
        raise ApiException.from_response(http_resp=res, body=raw, data=None)
    try:
        return LoginUserResponse.from_json(raw)
    except ApiException:
        raise
    except Exception as err:  # noqa: BLE001
        # A 2xx we couldn't decode is a genuine login result: surface it
        # with the raw body (typed error), not a bare decode error.
        raise ApiException.from_response(
            http_resp=res, body=raw, data=None
        ) from err


async def _do_fresh_login(
    client,
    *,
    email: str,
    password: str,
    api_key: str,
    uuid: str,
    two_fa_code: Optional[str],
    caller_pinned: bool,
) -> Tuple[LoginUserResponse, str]:
    """Perform a fresh (non-cached) email login and report which API key
    established the session — the Basic credential the later token
    refreshes must use (PORTING.md §6.2).

    Web-first: the browser login (``web_api_key``) is the primary path.
    The app login (``api_key``) is only a fallback for the narrow cases
    the browser endpoint doesn't serve, and the sole path when the
    browser endpoint is switched off (``web_api_key == ""``) or the
    caller pinned a specific api_key.
    """
    # Empty web key (off switch) or a caller-pinned key means "use the app
    # login with exactly this key"; skip the browser path so we never
    # silently authenticate under web_api_key instead.
    if not client.web_api_key or caller_pinned:
        resp = await _login_with_email_app(
            client,
            api_key=api_key,
            email=email,
            password=password,
            uuid=uuid,
            two_fa_code=two_fa_code,
        )
        return resp, api_key

    # Primary: browser login with the web key.
    try:
        resp = await _login_with_email_web(
            client,
            email=email,
            password=password,
            uuid=uuid,
            two_fa_code=two_fa_code,
        )
        return resp, client.web_api_key
    except Exception as err:  # noqa: BLE001
        api_err = as_api_error(err)
        # Fall back to the app login ONLY on a plain unauthorized. Any
        # other typed rejection surfaces unchanged — notably
        # ErrCodeRequired2FA (the caller must supply a code, not switch
        # endpoints) and a 2xx that failed to decode (a genuine login
        # result carrying its raw body, §15), neither of which is
        # ErrCodeUnauthorized. Gate on the parsed code, not the status.
        if code_of(api_err) != ErrCodeUnauthorized:
            raise api_err
        client._logger.debug(
            "app login fallback",
            extra={"event": "login_fallback", "outcome": "attempt"},
        )
        try:
            app_resp = await _login_with_email_app(
                client,
                api_key=api_key,
                email=email,
                password=password,
                uuid=uuid,
                two_fa_code=two_fa_code,
            )
        except Exception:  # noqa: BLE001
            client._logger.debug(
                "app login fallback failed",
                extra={"event": "login_fallback", "outcome": "failed"},
            )
            # Both paths failed: surface the primary (web) error, which
            # carries the parsed error code and raw body.
            raise api_err
        client._logger.debug(
            "app login fallback ok",
            extra={"event": "login_fallback", "outcome": "ok"},
        )
        return app_resp, api_key


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

    resolved_key = api_key or client.api_key
    resolved_uuid = uuid or client.device_uuid
    # A caller-pinned api_key (resolved key differs from the client
    # default) forces the app login under exactly that key.
    caller_pinned = resolved_key != client.api_key

    resp, oauth_key = await _do_fresh_login(
        client,
        email=email,
        password=password,
        api_key=resolved_key,
        uuid=resolved_uuid,
        two_fa_code=two_fa_code,
        caller_pinned=caller_pinned,
    )

    access = resp.access_token or ""
    refresh = resp.refresh_token or ""
    user_id = resp.user_id or 0
    client.set_login_identity(email, user_id)
    # Record which key refreshes this session before activating tokens.
    client.set_oauth_key(oauth_key)
    client.set_tokens(access, refresh)

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
