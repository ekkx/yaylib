// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct UserSetting {
    #[serde(rename = "notification_chat", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_chat: Option<Option<bool>>,
}

impl UserSetting {
    pub fn new() -> UserSetting {
        UserSetting {
            notification_chat: None,
        }
    }
}

