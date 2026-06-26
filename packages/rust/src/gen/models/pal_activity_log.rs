// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PalActivityLog {
    #[serde(rename = "created_at_seconds", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at_seconds: Option<Option<i64>>,
    #[serde(rename = "details", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub details: Option<Option<models::PalActivityLogDetails>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub r#type: Option<Option<serde_json::Value>>,
}

impl PalActivityLog {
    pub fn new() -> PalActivityLog {
        PalActivityLog {
            created_at_seconds: None,
            details: None,
            id: None,
            r#type: None,
        }
    }
}

