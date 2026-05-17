# Golden-vector parity with PORTING.md §7.1 — same inputs as the Go
# reference's signing_test.go / tests/signing.test.ts and the documented
# vectors. A drift in any of the three outputs means the formula changed
# (encoding, key choice, message body, trailing whitespace) and the
# Go / TS / Python ports are no longer byte-compatible.

from yaylib._config import (
    DEFAULT_API_KEY,
    DEFAULT_API_VERSION_KEY,
    DEFAULT_API_VERSION_NAME,
)
from yaylib.signing import (
    generate_signed_info_at,
    generate_signed_version,
    generate_x_jwt,
)


def test_signed_info_golden_vector():
    got = generate_signed_info_at(
        api_key=DEFAULT_API_KEY,
        device_uuid="00000000-0000-0000-0000-000000000000",
        timestamp=1700000000,
    )
    assert got.timestamp == 1700000000
    assert got.value == "253b7ca3e67590204034a73326d2be87"


def test_signed_version_golden_vector():
    got = generate_signed_version(
        api_version_key=DEFAULT_API_VERSION_KEY,
        api_version_name=DEFAULT_API_VERSION_NAME,
    )
    assert got == "zyKEqdZqaEqJjSRItHvchU9HmgWHNLtn02sWXwzQ4iQ="


def test_x_jwt_golden_vector():
    got = generate_x_jwt(
        api_version_key=DEFAULT_API_VERSION_KEY, now=1700000000
    )
    assert got == (
        "eyJhbGciOiJIUzI1NiJ9."
        "eyJpYXQiOjE3MDAwMDAwMDAsImV4cCI6MTcwMDAwMDAwNX0."
        "-UYjGTM53-Uwe8tNGjBf4ZLx644_zUEK8fpW1w16KfA"
    )
