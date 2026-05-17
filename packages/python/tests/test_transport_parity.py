# Transport behavior parity for the 401 -> refresh -> retry chain,
# driven against the shared behavior-contract server (PORTING.md §6.1 /
# §15). Translated from the shared-server subset of the Go reference's
# transport_test.go:
#
#   TestTransport_401TriggersRefreshAndRetries  (expired-token)
#   TestTransport_401RefreshFailureSurfaces401   (fail-401-times-2)
#
# The Go suite's header-injection / Basic-vs-Bearer / lazy-X-Client-IP
# checks stay as in-process fixtures (test_headers.py, test_client_ip.py
# — §15 keep-local: header-injection inspection and host routing), and
# the refresh-then-retry-network-error case stays in
# test_refresh_retry_neterror.py (it needs a socket-fault knob the
# shared contract server does not expose). These skip when
# YAYLIB_MOCK_BASE_URL is unset.

from yaylib.retry import RetryPolicy

from ._parity import mock_client, mock_get

# Mirror Go's RetryPolicy{} so the 401 refresh-and-replay path is
# isolated from transient-status retry.
NO_RETRY = RetryPolicy(max_attempts=0, base_delay=0.0, max_delay=0.0)


# PORTING:S7
async def test_401_triggers_refresh_and_retries():
    # expired-token: protected requests 401 (error_code -3) until the
    # token endpoint is hit (a simulated refresh), then the happy path —
    # exactly the 401 -> refresh -> retry chain. Behavioral assertion:
    # final success and the seeded tokens rotated to the server-issued
    # ones (the Bearer-header progression is an internal mechanic; §15
    # says translate the scenario, not the mechanics).
    c = mock_client("expired-token", retry_policy=NO_RETRY)
    c.set_tokens("STALE", "REF")
    try:
        resp = await mock_get(c)
        assert resp.status == 200, "refresh + retry succeeded"
        tok = c.tokens
        assert tok.access not in ("STALE", ""), (
            "access token should be rotated to the server-issued value"
        )
        assert tok.refresh not in ("REF", ""), (
            "refresh token should be rotated to the server-issued value"
        )
    finally:
        await c.close()


# PORTING:S8
async def test_401_refresh_failure_surfaces_401():
    # fail-401-times-2: the data request 401s and the refresh call (same
    # session+scenario counter) also 401s, so the refresh fails and the
    # original 401 surfaces with its body intact.
    c = mock_client("fail-401-times-2", retry_policy=NO_RETRY)
    c.set_tokens("STALE", "REF")
    try:
        resp = await mock_get(c)
        assert resp.status == 401, (
            "refresh failed, original 401 surfaced"
        )
        body = await resp.read()
        assert body, "the original 401 error body must be preserved"
    finally:
        await c.close()
