// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct EmplExpiringResponse {
    #[serde(rename = "expiring_amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub expiring_amount: Option<Option<String>>,
    #[serde(rename = "expiring_date", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub expiring_date: Option<Option<i64>>,
}

impl EmplExpiringResponse {
    pub fn new() -> EmplExpiringResponse {
        EmplExpiringResponse {
            expiring_amount: None,
            expiring_date: None,
        }
    }
}

