// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CodeResponse {
    #[serde(rename = "code", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub code: Option<Option<String>>,
    #[serde(rename = "next_editable_datetime", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub next_editable_datetime: Option<Option<i64>>,
}

impl CodeResponse {
    pub fn new() -> CodeResponse {
        CodeResponse {
            code: None,
            next_editable_datetime: None,
        }
    }
}

