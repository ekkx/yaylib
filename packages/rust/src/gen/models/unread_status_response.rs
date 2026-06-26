// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct UnreadStatusResponse {
    #[serde(rename = "is_unread", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_unread: Option<Option<bool>>,
}

impl UnreadStatusResponse {
    pub fn new() -> UnreadStatusResponse {
        UnreadStatusResponse {
            is_unread: None,
        }
    }
}

