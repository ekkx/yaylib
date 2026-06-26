// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ModelConferenceCallUserRole {
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "role", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub role: Option<Option<serde_json::Value>>,
}

impl ModelConferenceCallUserRole {
    pub fn new() -> ModelConferenceCallUserRole {
        ModelConferenceCallUserRole {
            id: None,
            role: None,
        }
    }
}

