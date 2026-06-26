// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct TotalChatRequestData {
    #[serde(rename = "total_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total_count: Option<Option<i32>>,
}

impl TotalChatRequestData {
    pub fn new() -> TotalChatRequestData {
        TotalChatRequestData {
            total_count: None,
        }
    }
}

