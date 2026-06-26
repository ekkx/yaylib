// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct MissionSection {
    #[serde(rename = "mission_list", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub mission_list: Option<Option<Vec<models::MissionItem>>>,
    #[serde(rename = "mission_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub mission_type: Option<Option<serde_json::Value>>,
    #[serde(rename = "section", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub section: Option<Option<models::MissionSectionHeader>>,
}

impl MissionSection {
    pub fn new() -> MissionSection {
        MissionSection {
            mission_list: None,
            mission_type: None,
            section: None,
        }
    }
}

