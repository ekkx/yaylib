// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ChatRoomsResponse {
    #[serde(rename = "chat_rooms", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub chat_rooms: Option<Option<Vec<models::RealmChatRoom>>>,
    #[serde(rename = "next_page_value", default, skip_serializing_if = "Option::is_none", deserialize_with = "crate::gen::cursor::deserialize_option_string_or_number")]
    pub next_page_value: Option<String>,
    #[serde(rename = "pinned_chat_rooms", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pinned_chat_rooms: Option<Option<Vec<models::RealmChatRoom>>>,
}

impl ChatRoomsResponse {
    pub fn new() -> ChatRoomsResponse {
        ChatRoomsResponse {
            chat_rooms: None,
            next_page_value: None,
            pinned_chat_rooms: None,
        }
    }
}

