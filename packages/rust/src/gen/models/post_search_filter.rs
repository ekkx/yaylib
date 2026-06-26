// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PostSearchFilter {
    #[serde(rename = "message_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub message_id: Option<Option<i32>>,
    #[serde(rename = "title_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title_id: Option<Option<i32>>,
}

impl PostSearchFilter {
    pub fn new() -> PostSearchFilter {
        PostSearchFilter {
            message_id: None,
            title_id: None,
        }
    }
}

