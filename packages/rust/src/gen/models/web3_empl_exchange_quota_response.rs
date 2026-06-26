// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3EmplExchangeQuotaResponse {
    #[serde(rename = "empl_exchange_limit", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub empl_exchange_limit: Option<Option<i64>>,
    #[serde(rename = "empl_exchange_quota_remaining", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub empl_exchange_quota_remaining: Option<Option<i64>>,
    #[serde(rename = "timestamp", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub timestamp: Option<Option<i64>>,
}

impl Web3EmplExchangeQuotaResponse {
    pub fn new() -> Web3EmplExchangeQuotaResponse {
        Web3EmplExchangeQuotaResponse {
            empl_exchange_limit: None,
            empl_exchange_quota_remaining: None,
            timestamp: None,
        }
    }
}

