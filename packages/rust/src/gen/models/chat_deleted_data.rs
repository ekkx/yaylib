// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ChatDeletedData {
    #[serde(rename = "room_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub room_id: Option<Option<i64>>,
}

impl ChatDeletedData {
    pub fn new() -> ChatDeletedData {
        ChatDeletedData {
            room_id: None,
        }
    }
}

