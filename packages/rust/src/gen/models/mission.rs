// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Mission {
    #[serde(rename = "action", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub action: Option<Option<String>>,
    #[serde(rename = "detail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub detail: Option<Option<String>>,
    #[serde(rename = "is_multiplier", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_multiplier: Option<Option<bool>>,
    #[serde(rename = "mission_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub mission_type: Option<Option<models::MissionType>>,
    #[serde(rename = "required_actions_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub required_actions_count: Option<Option<i32>>,
}

impl Mission {
    pub fn new() -> Mission {
        Mission {
            action: None,
            detail: None,
            is_multiplier: None,
            mission_type: None,
            required_actions_count: None,
        }
    }
}

