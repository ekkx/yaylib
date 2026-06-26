// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Callback {
    #[serde(rename = "callback_method", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub callback_method: Option<Option<serde_json::Value>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i32>>,
    #[serde(rename = "response_string", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub response_string: Option<Option<String>>,
}

impl Callback {
    pub fn new() -> Callback {
        Callback {
            callback_method: None,
            id: None,
            response_string: None,
        }
    }
}

