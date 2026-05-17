# Raw-escape-hatch parity, driven against the shared behavior-contract
# server's body-malformed contract (PORTING.md §15): the endpoint
# answers HTTP 200 with a JSON scalar no typed model can decode. Two
# halves:
#
#   1. the typed call surfaces the parse failure as an APIError that
#      still carries the raw body (so code_of(err) / inspection works);
#   2. the per-call raw escape hatch (client.execute_raw) hands back the
#      unparsed bytes and the 200 HTTP response without raising.
#
# Skips when YAYLIB_MOCK_BASE_URL is unset, like every parity test.

import pytest

from yaylib.errors import APIError

from ._parity import mock_client, require_mock


async def test_typed_call_on_malformed_body_raises_apierror_with_raw_body():
    require_mock()
    c = mock_client("body-malformed")
    try:
        with pytest.raises(APIError) as excinfo:
            await c.buckets_api.get_bucket_presigned_urls(file_names=["x"])
        err = excinfo.value
        # A 2xx body that does not match the typed model surfaces as an
        # APIError that retains the raw body — not a bare validation
        # error with nothing to inspect.
        assert err.body, "typed parse error must retain the raw body"
        assert len(str(err.body)) > 0
    finally:
        await c.close()


async def test_execute_raw_returns_body_and_response_without_raising():
    require_mock()
    c = mock_client("body-malformed")
    try:
        body, resp = await c.execute_raw(
            c.buckets_api.get_bucket_presigned_urls_without_preload_content(
                file_names=["x"]
            )
        )
        # The raw hatch bypasses the typed decode entirely: a 200 whose
        # body the typed model rejects still comes back as bytes + the
        # successful response, no error.
        assert resp.status == 200
        assert isinstance(body, (bytes, bytearray)) and len(body) > 0
    finally:
        await c.close()
