// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct MissionItem {
    #[serde(rename = "mission", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub mission: Option<Option<serde_json::Value>>,
}

impl MissionItem {
    pub fn new() -> MissionItem {
        MissionItem {
            mission: None,
        }
    }
}

