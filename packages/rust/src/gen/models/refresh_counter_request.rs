// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct RefreshCounterRequest {
    #[serde(rename = "counter", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub counter: Option<Option<String>>,
    #[serde(rename = "last_requested_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub last_requested_at: Option<Option<i64>>,
    #[serde(rename = "status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub status: Option<Option<String>>,
}

impl RefreshCounterRequest {
    pub fn new() -> RefreshCounterRequest {
        RefreshCounterRequest {
            counter: None,
            last_requested_at: None,
            status: None,
        }
    }
}

