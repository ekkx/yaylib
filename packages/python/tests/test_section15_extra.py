# Two contract scenarios that don't fit the existing parity / error
# files (PORTING.md §15):
#
#   S29  any 2xx is a success — a non-200 2xx (the shared server's
#        ok-status-201) still decodes into the typed success model
#        rather than discarding the body (§11 rule 3). Driven against
#        the shared behavior-contract server; skips when its env is
#        unset, like every parity test.
#   S30  code_of resolves to the typed ErrorCode constant, hermetically
#        (no server): an envelope carrying error_code -31 maps to the
#        exported ErrCodeRequired2FA constant, and a non-API error
#        yields the zero-ish ErrCodeUnknown without raising.

import json

from yaylib.errors import APIError, ErrCodeRequired2FA, code_of

from ._parity import mock_client, require_mock


# PORTING:S29
async def test_any_2xx_is_success():
    require_mock()
    c = mock_client("ok-status-201")
    try:
        # The server answers HTTP 201 with an empty-object body. The SDK
        # must not key strictly on the documented 200 and discard the
        # body — it falls back to the documented 2xx success type and
        # decodes the (all-optional) model.
        out = await c.buckets_api.get_bucket_presigned_urls(file_names=["x"])
        assert out is not None
    finally:
        await c.close()


# PORTING:S30
def test_code_of_resolves_typed_constant():
    # An error envelope carrying error_code -31 must resolve to the
    # exported ErrCodeRequired2FA constant (not just an arbitrary int).
    # Same hermetic construction the error suite uses for raw-body
    # retention: a real APIError with a chosen JSON body.
    err = APIError(status=401, reason="Unauthorized")
    err.body = json.dumps({"error_code": -31, "message": "x"})
    code = code_of(err)
    assert code == ErrCodeRequired2FA
    assert code == -31  # the typed constant's value, pinned

    # A non-API / code-less error yields the zero-ish ErrCodeUnknown and
    # never raises.
    assert code_of(Exception("plain")) == 0
