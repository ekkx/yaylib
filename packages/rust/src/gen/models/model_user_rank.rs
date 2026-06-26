// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ModelUserRank {
    #[serde(rename = "disabled_icon_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub disabled_icon_id: Option<Option<i32>>,
    #[serde(rename = "enabled_icon_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub enabled_icon_id: Option<Option<i32>>,
    #[serde(rename = "value", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub value: Option<Option<String>>,
}

impl ModelUserRank {
    pub fn new() -> ModelUserRank {
        ModelUserRank {
            disabled_icon_id: None,
            enabled_icon_id: None,
            value: None,
        }
    }
}

