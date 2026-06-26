// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Qualification {
    #[serde(rename = "qualified", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub qualified: Option<Option<bool>>,
    #[serde(rename = "type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub r#type: Option<Option<serde_json::Value>>,
}

impl Qualification {
    pub fn new() -> Qualification {
        Qualification {
            qualified: None,
            r#type: None,
        }
    }
}

