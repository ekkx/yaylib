use serde::Deserialize;

use crate::auth::SessionSaveFailed;
use crate::error_codes::ErrorCode;

/// Error type for all API operations.
#[derive(Debug, thiserror::Error)]
pub enum ApiError {
    #[error("API error: status {status}")]
    Response {
        status: u16,
        body: Vec<u8>,
        error_response: Option<ErrorResponse>,
    },

    #[error(transparent)]
    Http(#[from] reqwest::Error),

    #[error(transparent)]
    Json(#[from] serde_json::Error),

    #[error(transparent)]
    SessionSaveFailed(SessionSaveFailed),

    #[error("{0}")]
    Other(String),
}

impl ApiError {
    pub fn from_response(status: u16, body: Vec<u8>) -> Self {
        let error_response = serde_json::from_slice(&body).ok();
        Self::Response {
            status,
            body,
            error_response,
        }
    }
}

impl<T: std::fmt::Debug> From<crate::gen::apis::Error<T>> for ApiError {
    fn from(e: crate::gen::apis::Error<T>) -> Self {
        match e {
            crate::gen::apis::Error::Reqwest(e) => ApiError::Http(e),
            crate::gen::apis::Error::Serde(e) => ApiError::Json(e),
            crate::gen::apis::Error::Io(e) => ApiError::Other(e.to_string()),
            crate::gen::apis::Error::ResponseError(r) => {
                ApiError::from_response(r.status.as_u16(), r.content.into_bytes())
            }
        }
    }
}

/// Parsed error response body from the Yay! API.
#[derive(Debug, Clone, Deserialize)]
pub struct ErrorResponse {
    #[serde(default, rename = "error_code")]
    pub code: i32,
    #[serde(default)]
    pub message: String,
    pub ban_until: Option<i64>,
    pub retry_in: Option<i64>,
    pub remaining_quota: Option<i64>,
}

/// Extracts the [`ErrorCode`] from an API error, or [`ErrorCode::Unknown`].
pub fn code_of(err: &ApiError) -> ErrorCode {
    error_response_of(err)
        .and_then(|r| ErrorCode::from_code(r.code))
        .unwrap_or(ErrorCode::Unknown)
}

/// Extracts the parsed [`ErrorResponse`] from an API error, if available.
pub fn error_response_of(err: &ApiError) -> Option<&ErrorResponse> {
    match err {
        ApiError::Response {
            error_response: Some(ref r),
            ..
        } => Some(r),
        ApiError::Response { body, .. } if !body.is_empty() => None,
        _ => None,
    }
}
