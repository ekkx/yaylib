// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct SectionNavigation {
    #[serde(rename = "icon_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub icon_id: Option<Option<i32>>,
    #[serde(rename = "is_enabled", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_enabled: Option<Option<bool>>,
    #[serde(rename = "section_index", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub section_index: Option<Option<i32>>,
    #[serde(rename = "title_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title_id: Option<Option<i32>>,
}

impl SectionNavigation {
    pub fn new() -> SectionNavigation {
        SectionNavigation {
            icon_id: None,
            is_enabled: None,
            section_index: None,
            title_id: None,
        }
    }
}

