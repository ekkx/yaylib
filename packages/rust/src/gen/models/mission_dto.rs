// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct MissionDto {
    #[serde(rename = "action", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub action: Option<Option<String>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "mission_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub mission_type: Option<Option<models::MissionType>>,
    #[serde(rename = "progress", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub progress: Option<Option<models::ProgressDto>>,
    #[serde(rename = "required_actions_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub required_actions_count: Option<Option<i32>>,
    #[serde(rename = "reward_points", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub reward_points: Option<Option<i32>>,
}

impl MissionDto {
    pub fn new() -> MissionDto {
        MissionDto {
            action: None,
            id: None,
            mission_type: None,
            progress: None,
            required_actions_count: None,
            reward_points: None,
        }
    }
}

