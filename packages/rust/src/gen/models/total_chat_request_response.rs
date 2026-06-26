// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct TotalChatRequestResponse {
    #[serde(rename = "total", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total: Option<Option<i32>>,
}

impl TotalChatRequestResponse {
    pub fn new() -> TotalChatRequestResponse {
        TotalChatRequestResponse {
            total: None,
        }
    }
}

