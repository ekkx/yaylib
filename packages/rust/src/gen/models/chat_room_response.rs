// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ChatRoomResponse {
    #[serde(rename = "chat", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub chat: Option<Option<models::RealmChatRoom>>,
}

impl ChatRoomResponse {
    pub fn new() -> ChatRoomResponse {
        ChatRoomResponse {
            chat: None,
        }
    }
}

