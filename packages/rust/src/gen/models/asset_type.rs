// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct AssetType {
    #[serde(rename = "api_value", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub api_value: Option<Option<String>>,
}

impl AssetType {
    pub fn new() -> AssetType {
        AssetType {
            api_value: None,
        }
    }
}

