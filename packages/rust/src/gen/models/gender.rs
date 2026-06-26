// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Gender {
    #[serde(rename = "api_int_value", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub api_int_value: Option<Option<i32>>,
    #[serde(rename = "api_string_value", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub api_string_value: Option<Option<String>>,
}

impl Gender {
    pub fn new() -> Gender {
        Gender {
            api_int_value: None,
            api_string_value: None,
        }
    }
}

