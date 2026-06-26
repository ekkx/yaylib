// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PendingEmplActivity {
    #[serde(rename = "activity", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub activity: Option<Option<serde_json::Value>>,
    #[serde(rename = "details", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub details: Option<Option<models::PendingEmplActivityDetails>>,
}

impl PendingEmplActivity {
    pub fn new() -> PendingEmplActivity {
        PendingEmplActivity {
            activity: None,
            details: None,
        }
    }
}

