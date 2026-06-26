// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Error {
    #[serde(rename = "action", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub action: Option<Option<serde_json::Value>>,
    #[serde(rename = "original_code", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub original_code: Option<Option<i32>>,
    #[serde(rename = "throwable", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub throwable: Option<Option<serde_json::Value>>,
    #[serde(rename = "type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub r#type: Option<Option<models::ErrorType>>,
}

impl Error {
    pub fn new() -> Error {
        Error {
            action: None,
            original_code: None,
            throwable: None,
            r#type: None,
        }
    }
}

