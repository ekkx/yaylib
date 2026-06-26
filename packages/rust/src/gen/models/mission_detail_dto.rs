// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct MissionDetailDto {
    #[serde(rename = "action", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub action: Option<Option<String>>,
    #[serde(rename = "detail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub detail: Option<Option<String>>,
    #[serde(rename = "mission_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub mission_id: Option<Option<i64>>,
}

impl MissionDetailDto {
    pub fn new() -> MissionDetailDto {
        MissionDetailDto {
            action: None,
            detail: None,
            mission_id: None,
        }
    }
}

