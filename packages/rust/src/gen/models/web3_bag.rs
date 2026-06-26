// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3Bag {
    #[serde(rename = "empl", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub empl: Option<Option<serde_json::Value>>,
    #[serde(rename = "pals_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pals_count: Option<Option<i32>>,
}

impl Web3Bag {
    pub fn new() -> Web3Bag {
        Web3Bag {
            empl: None,
            pals_count: None,
        }
    }
}

