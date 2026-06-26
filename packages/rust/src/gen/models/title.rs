// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Title {
    #[serde(rename = "api_value", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub api_value: Option<Option<String>>,
}

impl Title {
    pub fn new() -> Title {
        Title {
            api_value: None,
        }
    }
}

