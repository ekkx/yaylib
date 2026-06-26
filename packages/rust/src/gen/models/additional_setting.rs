// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct AdditionalSetting {
    #[serde(rename = "api_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub api_id: Option<Option<String>>,
    #[serde(rename = "is_on", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_on: Option<Option<serde_json::Value>>,
    #[serde(rename = "is_on_from_user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_on_from_user: Option<Option<serde_json::Value>>,
}

impl AdditionalSetting {
    pub fn new() -> AdditionalSetting {
        AdditionalSetting {
            api_id: None,
            is_on: None,
            is_on_from_user: None,
        }
    }
}

