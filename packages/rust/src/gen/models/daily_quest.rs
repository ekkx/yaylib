// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct DailyQuest {
    #[serde(rename = "calculated", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub calculated: Option<Option<bool>>,
    #[serde(rename = "spent_duration_seconds", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub spent_duration_seconds: Option<Option<i32>>,
    #[serde(rename = "target_duration_seconds", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub target_duration_seconds: Option<Option<i32>>,
    #[serde(rename = "timestamp", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub timestamp: Option<Option<i64>>,
}

impl DailyQuest {
    pub fn new() -> DailyQuest {
        DailyQuest {
            calculated: None,
            spent_duration_seconds: None,
            target_duration_seconds: None,
            timestamp: None,
        }
    }
}

