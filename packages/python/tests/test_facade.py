# Flat operation facade (PORTING.md §15 S33 / §2).
#
# Every operation is reachable as client.<op> with no service prefix
# (a generated mixin delegating to the per-service attribute), and a
# hand-written wrapper of the same name shadows the generated op via
# MRO. Pure local check — no server is involved.

import pytest

from yaylib.client import Client


# PORTING:S33
async def test_flat_ops_and_wrapper_shadow():
    c = Client()
    try:
        # Operations are reachable directly on the client (no
        # client.posts_api.).
        assert callable(getattr(c, "get_timeline", None))
        assert callable(getattr(c, "get_recommended_timeline", None))

        # The hand-written cached wrapper shadows the generated
        # login_with_email op: empty credentials fail fast with the
        # wrapper's validation error, never an HTTP attempt.
        with pytest.raises(ValueError, match="requires email and password"):
            await c.login_with_email(email="", password="")
    finally:
        await c.close()
