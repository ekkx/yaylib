// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CreateChatRoomResponse {
    #[serde(rename = "room_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub room_id: Option<Option<i64>>,
}

impl CreateChatRoomResponse {
    pub fn new() -> CreateChatRoomResponse {
        CreateChatRoomResponse {
            room_id: None,
        }
    }
}

