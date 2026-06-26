// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct UserTimestampResponse {
    #[serde(rename = "country", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub country: Option<Option<String>>,
    #[serde(rename = "ip_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub ip_address: Option<Option<String>>,
    #[serde(rename = "time", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub time: Option<Option<i64>>,
}

impl UserTimestampResponse {
    pub fn new() -> UserTimestampResponse {
        UserTimestampResponse {
            country: None,
            ip_address: None,
            time: None,
        }
    }
}

