// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3WalletDailyQuest {
    #[serde(rename = "is_calculated", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_calculated: Option<Option<bool>>,
    #[serde(rename = "pal_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal_count: Option<Option<i32>>,
    #[serde(rename = "spent_duration_seconds", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub spent_duration_seconds: Option<Option<i32>>,
    #[serde(rename = "total_duration_needed_seconds", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total_duration_needed_seconds: Option<Option<i32>>,
    #[serde(rename = "updated_at_millis", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub updated_at_millis: Option<Option<i64>>,
}

impl Web3WalletDailyQuest {
    pub fn new() -> Web3WalletDailyQuest {
        Web3WalletDailyQuest {
            is_calculated: None,
            pal_count: None,
            spent_duration_seconds: None,
            total_duration_needed_seconds: None,
            updated_at_millis: None,
        }
    }
}

