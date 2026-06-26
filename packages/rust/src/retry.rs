use std::time::Duration;

use rand::Rng;

/// Configuration for automatic request retries with exponential backoff.
#[derive(Debug, Clone)]
pub struct RetryPolicy {
    pub max_attempts: u32,
    pub base_delay: Duration,
    pub max_delay: Duration,
    pub retry_on_post: bool,
}

impl Default for RetryPolicy {
    fn default() -> Self {
        Self {
            max_attempts: 3,
            base_delay: Duration::from_millis(200),
            max_delay: Duration::from_secs(30),
            retry_on_post: false,
        }
    }
}

impl RetryPolicy {
    /// Returns a policy that disables all retries.
    pub fn disabled() -> Self {
        Self {
            max_attempts: 1,
            ..Default::default()
        }
    }
}

pub(crate) fn should_retry(
    status: Option<u16>,
    is_network_error: bool,
    method: &str,
    policy: &RetryPolicy,
) -> bool {
    // 429 means the server rejected at the rate limiter before processing,
    // so there is no duplicate-creation risk -- retry regardless of method.
    if status == Some(429) {
        return true;
    }

    if !can_retry_method(method, policy) {
        return false;
    }

    if is_network_error {
        return true;
    }

    if let Some(code) = status {
        // 401 is owned by the auth-refresh path; surviving 401s don't fix
        // themselves with more retries.
        if code == 401 {
            return false;
        }
        if (500..=599).contains(&code) {
            return true;
        }
    }

    false
}

fn can_retry_method(method: &str, policy: &RetryPolicy) -> bool {
    match method {
        "GET" | "HEAD" | "PUT" | "DELETE" | "OPTIONS" => true,
        "POST" | "PATCH" => policy.retry_on_post,
        _ => false,
    }
}

pub(crate) fn next_delay(
    server_retry_in: Option<Duration>,
    attempt: u32,
    policy: &RetryPolicy,
) -> Duration {
    if let Some(d) = server_retry_in {
        return if policy.max_delay > Duration::ZERO && d > policy.max_delay {
            policy.max_delay
        } else {
            d
        };
    }

    if policy.base_delay.is_zero() {
        return Duration::ZERO;
    }

    // Exponential: base_delay << (attempt - 2).
    // attempt=2 -> base, attempt=3 -> base*2, ...
    let shift = attempt.saturating_sub(2).min(30) as u64;
    let base_ms = policy.base_delay.as_millis() as u64;
    let max_ms = policy.max_delay.as_millis() as u64;

    let exp_ms = base_ms.saturating_mul(1u64 << shift).min(max_ms);
    if exp_ms == 0 {
        return Duration::ZERO;
    }

    // Full jitter: uniform in [0, exp_ms].
    let jittered = rand::thread_rng().gen_range(0..=exp_ms);
    Duration::from_millis(jittered)
}

/// Parses the Yay! server's `retry_in` field (seconds) from a JSON error body.
pub(crate) fn extract_retry_in(body: &[u8]) -> Option<Duration> {
    if body.len() > 16 * 1024 {
        return None;
    }

    #[derive(serde::Deserialize)]
    struct Body {
        retry_in: Option<i64>,
    }

    let parsed: Body = serde_json::from_slice(body).ok()?;
    let secs = parsed.retry_in?;
    if secs <= 0 {
        return None;
    }
    Some(Duration::from_secs(secs as u64))
}
