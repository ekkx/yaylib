// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct MissionTypeX {
    #[serde(rename = "section_info_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub section_info_id: Option<Option<i32>>,
    #[serde(rename = "section_title_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub section_title_id: Option<Option<i32>>,
    #[serde(rename = "title_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title_id: Option<Option<i32>>,
}

impl MissionTypeX {
    pub fn new() -> MissionTypeX {
        MissionTypeX {
            section_info_id: None,
            section_title_id: None,
            title_id: None,
        }
    }
}

