// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct BanWordType {
    #[serde(rename = "api_value", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub api_value: Option<Option<String>>,
}

impl BanWordType {
    pub fn new() -> BanWordType {
        BanWordType {
            api_value: None,
        }
    }
}

