// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ProgressDto {
    #[serde(rename = "actions_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub actions_count: Option<Option<i32>>,
    #[serde(rename = "completed_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub completed_at: Option<Option<i64>>,
    #[serde(rename = "total_earned_points", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total_earned_points: Option<Option<i32>>,
}

impl ProgressDto {
    pub fn new() -> ProgressDto {
        ProgressDto {
            actions_count: None,
            completed_at: None,
            total_earned_points: None,
        }
    }
}

