use base64::engine::general_purpose::{STANDARD as BASE64_STANDARD, URL_SAFE_NO_PAD};
use base64::Engine;
use hmac::{Hmac, Mac};
use md5::{Digest, Md5};
use sha2::Sha256;

const SIGNED_INFO_SHARED_KEY: &str = "yayZ1";
const SIGNED_VERSION_PLATFORM: &str = "yay_android";
const X_JWT_TTL: i64 = 5;

/// The timestamp and MD5 hash pair required by signed endpoints.
#[derive(Debug, Clone, PartialEq, Eq)]
pub struct SignedInfo {
    pub timestamp: i64,
    pub value: String,
}

/// Computes the signed_info MD5 hash for a given timestamp.
///
/// ```text
/// signed_info = lpad(MD5(api_key + device_uuid + timestamp + "yayZ1"), 32, '0')
/// ```
///
/// MD5 hex output is always 32 chars, so the lpad is a no-op in practice,
/// but the format string guarantees it.
pub fn compute_signed_info_at(api_key: &str, device_uuid: &str, timestamp: i64) -> SignedInfo {
    let payload = format!("{}{}{}{}", api_key, device_uuid, timestamp, SIGNED_INFO_SHARED_KEY);
    let hash = Md5::digest(payload.as_bytes());
    SignedInfo {
        timestamp,
        value: format!("{:032x}", hash),
    }
}

/// Computes the signed_version HMAC token.
///
/// ```text
/// signed_version = base64(HMAC-SHA256(key=api_version_key, msg="yay_android/" + api_version_name))
/// ```
///
/// Uses standard base64 (with `=` padding). The output is always 44 characters
/// for a 32-byte HMAC digest.
pub fn compute_signed_version(api_version_key: &str, api_version_name: &str) -> String {
    let payload = format!("{}/{}", SIGNED_VERSION_PLATFORM, api_version_name);
    let mut mac =
        Hmac::<Sha256>::new_from_slice(api_version_key.as_bytes()).expect("HMAC accepts any key");
    mac.update(payload.as_bytes());
    BASE64_STANDARD.encode(mac.finalize().into_bytes())
}

/// Generates the short-lived HS256 JWT required by some write endpoints.
///
/// The token carries only `{iat, exp=iat+5}` and is signed with the
/// api_version_key. Must be regenerated per request because the 5-second
/// validity window expires quickly.
pub fn generate_x_jwt(api_version_key: &str) -> String {
    let iat = chrono::Utc::now().timestamp();
    generate_x_jwt_at(api_version_key, iat, X_JWT_TTL)
}

/// Generates an HS256 JWT with an explicit `iat` and `ttl`. Exposed for
/// testing; production callers should use [`generate_x_jwt`].
///
/// Header is `{"alg":"HS256"}` (no `typ`). Payload key order is `iat` then
/// `exp`, no spaces. Base64url encoding without padding.
pub fn generate_x_jwt_at(api_version_key: &str, iat: i64, ttl: i64) -> String {
    let header = URL_SAFE_NO_PAD.encode(b"{\"alg\":\"HS256\"}");
    let payload_json = format!("{{\"iat\":{},\"exp\":{}}}", iat, iat + ttl);
    let payload = URL_SAFE_NO_PAD.encode(payload_json.as_bytes());
    let signing_input = format!("{}.{}", header, payload);

    let mut mac =
        Hmac::<Sha256>::new_from_slice(api_version_key.as_bytes()).expect("HMAC accepts any key");
    mac.update(signing_input.as_bytes());
    let sig = URL_SAFE_NO_PAD.encode(mac.finalize().into_bytes());

    format!("{}.{}", signing_input, sig)
}

#[cfg(test)]
mod tests {
    use super::*;

    const TEST_API_KEY: &str =
        "ccd59ee269c01511ba763467045c115779fcae3050238a252f1bd1a4b65cfec6";
    const TEST_API_VERSION_KEY: &str = "34c8c1cdf29b46a492e8ec58e6db37ec";
    const TEST_API_VERSION_NAME: &str = "4.26";
    const TEST_DEVICE_UUID: &str = "00000000-0000-0000-0000-000000000000";

    #[test]
    fn signed_info_golden_vector() {
        let si = compute_signed_info_at(TEST_API_KEY, TEST_DEVICE_UUID, 1700000000);
        assert_eq!(si.value, "253b7ca3e67590204034a73326d2be87");
        assert_eq!(si.timestamp, 1700000000);
        assert_eq!(si.value.len(), 32);
    }

    #[test]
    fn signed_info_deterministic() {
        let a = compute_signed_info_at("test-api-key", "test-uuid", 1700000000);
        let b = compute_signed_info_at("test-api-key", "test-uuid", 1700000000);
        assert_eq!(a, b);
    }

    #[test]
    fn signed_info_depends_on_credentials() {
        let a = compute_signed_info_at("key-A", "uuid-shared", 1);
        let b = compute_signed_info_at("key-B", "uuid-shared", 1);
        assert_ne!(a.value, b.value);
    }

    #[test]
    fn signed_version_golden_vector() {
        let sv = compute_signed_version(TEST_API_VERSION_KEY, TEST_API_VERSION_NAME);
        assert_eq!(sv, "zyKEqdZqaEqJjSRItHvchU9HmgWHNLtn02sWXwzQ4iQ=");
    }

    #[test]
    fn signed_version_length() {
        let sv = compute_signed_version(TEST_API_VERSION_KEY, TEST_API_VERSION_NAME);
        assert_eq!(sv.len(), 44, "32-byte HMAC → 44-char standard base64");
    }

    #[test]
    fn signed_version_no_trailing_whitespace() {
        let sv = compute_signed_version(TEST_API_VERSION_KEY, TEST_API_VERSION_NAME);
        assert!(!sv.ends_with('\n'));
        assert!(!sv.ends_with(' '));
    }

    #[test]
    fn signed_version_depends_on_key() {
        let a = compute_signed_version(TEST_API_VERSION_KEY, TEST_API_VERSION_NAME);
        let b = compute_signed_version("00000000000000000000000000000000", TEST_API_VERSION_NAME);
        assert_ne!(a, b);
    }

    #[test]
    fn signed_version_follows_version_name() {
        let a = compute_signed_version(TEST_API_VERSION_KEY, "4.26");
        let b = compute_signed_version(TEST_API_VERSION_KEY, "9.99");
        assert_ne!(a, b);
    }

    #[test]
    fn x_jwt_golden_vector() {
        let jwt = generate_x_jwt_at(TEST_API_VERSION_KEY, 1700000000, 5);
        assert_eq!(
            jwt,
            "eyJhbGciOiJIUzI1NiJ9.\
             eyJpYXQiOjE3MDAwMDAwMDAsImV4cCI6MTcwMDAwMDAwNX0.\
             -UYjGTM53-Uwe8tNGjBf4ZLx644_zUEK8fpW1w16KfA"
        );
    }

    #[test]
    fn x_jwt_three_parts() {
        let jwt = generate_x_jwt_at(TEST_API_VERSION_KEY, 1700000000, 5);
        assert_eq!(jwt.split('.').count(), 3);
    }

    #[test]
    fn x_jwt_header_has_no_typ() {
        let jwt = generate_x_jwt_at(TEST_API_VERSION_KEY, 1700000000, 5);
        let header_b64 = jwt.split('.').next().unwrap();
        let header_json = URL_SAFE_NO_PAD.decode(header_b64).unwrap();
        let header_str = std::str::from_utf8(&header_json).unwrap();
        assert_eq!(header_str, r#"{"alg":"HS256"}"#);
    }
}
