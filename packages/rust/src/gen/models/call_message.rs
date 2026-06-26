// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CallMessage {
    #[serde(rename = "created_at_millis", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at_millis: Option<Option<i64>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<String>>,
    #[serde(rename = "is_message_deleted", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_message_deleted: Option<Option<bool>>,
    #[serde(rename = "text", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub text: Option<Option<String>>,
    #[serde(rename = "user_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user_id: Option<Option<String>>,
}

impl CallMessage {
    pub fn new() -> CallMessage {
        CallMessage {
            created_at_millis: None,
            id: None,
            is_message_deleted: None,
            text: None,
            user_id: None,
        }
    }
}

