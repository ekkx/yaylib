// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ApiB {
    #[serde(rename = "d", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub d: Option<Option<serde_json::Value>>,
    #[serde(rename = "f", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub f: Option<Option<serde_json::Value>>,
}

impl ApiB {
    pub fn new() -> ApiB {
        ApiB {
            d: None,
            f: None,
        }
    }
}

