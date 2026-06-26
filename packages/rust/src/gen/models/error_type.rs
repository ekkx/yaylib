// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ErrorType {
    #[serde(rename = "code", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub code: Option<Option<i32>>,
}

impl ErrorType {
    pub fn new() -> ErrorType {
        ErrorType {
            code: None,
        }
    }
}

