use std::time::Duration;

pub const DEFAULT_BASE_URL: &str = "https://api.yay.space";
pub const DEFAULT_CASSANDRA_BASE_URL: &str = "https://cas.yay.space";

pub const DEFAULT_API_KEY: &str =
    "ccd59ee269c01511ba763467045c115779fcae3050238a252f1bd1a4b65cfec6";
pub const DEFAULT_API_VERSION_KEY: &str = "34c8c1cdf29b46a492e8ec58e6db37ec";
pub const DEFAULT_API_VERSION_NAME: &str = "4.26";
pub const DEFAULT_APP_VERSION: &str = "4.26.1";

pub const DEFAULT_SHARED_KEY: &str = "yayZ1";

pub const DEFAULT_DEVICE_TYPE: &str = "android";
pub const DEFAULT_DEVICE_OS: &str = "11";
pub const DEFAULT_DEVICE_DENSITY: &str = "3.5";
pub const DEFAULT_DEVICE_SCREEN: &str = "1440x2960";
pub const DEFAULT_DEVICE_MODEL: &str = "Galaxy S9";

pub const DEFAULT_CONNECTION_TYPE: &str = "wifi";
pub const DEFAULT_CONNECTION_SPEED: &str = "0 kbps";
pub const DEFAULT_ACCEPT_LANGUAGE: &str = "ja";

pub const DEFAULT_HTTP_TIMEOUT: Duration = Duration::from_secs(90);

pub fn default_user_agent() -> String {
    format!(
        "{} {} ({}x {} {})",
        DEFAULT_DEVICE_TYPE,
        DEFAULT_DEVICE_OS,
        DEFAULT_DEVICE_DENSITY,
        DEFAULT_DEVICE_SCREEN,
        DEFAULT_DEVICE_MODEL,
    )
}

pub fn default_device_info() -> String {
    format!("yay {} {}", DEFAULT_APP_VERSION, default_user_agent())
}
