// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CommonErrorResponse {
    #[serde(rename = "ban_until", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub ban_until: Option<Option<i64>>,
    #[serde(rename = "error_code", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub error_code: Option<Option<i32>>,
    #[serde(rename = "message", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub message: Option<Option<String>>,
    #[serde(rename = "remaining_quota", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub remaining_quota: Option<Option<i64>>,
    #[serde(rename = "retry_in", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub retry_in: Option<Option<i64>>,
}

impl CommonErrorResponse {
    pub fn new() -> CommonErrorResponse {
        CommonErrorResponse {
            ban_until: None,
            error_code: None,
            message: None,
            remaining_quota: None,
            retry_in: None,
        }
    }
}

