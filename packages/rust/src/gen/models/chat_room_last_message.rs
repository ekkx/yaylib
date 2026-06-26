// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ChatRoomLastMessage {
    #[serde(rename = "conference_call", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub conference_call: Option<Option<models::RealmConferenceCall>>,
    #[serde(rename = "created_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<Option<i64>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "message_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub message_type: Option<Option<models::MessageType>>,
    #[serde(rename = "room_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub room_id: Option<Option<i64>>,
    #[serde(rename = "text", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub text: Option<Option<String>>,
    #[serde(rename = "user_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user_id: Option<Option<i64>>,
}

impl ChatRoomLastMessage {
    pub fn new() -> ChatRoomLastMessage {
        ChatRoomLastMessage {
            conference_call: None,
            created_at: None,
            id: None,
            message_type: None,
            room_id: None,
            text: None,
            user_id: None,
        }
    }
}

