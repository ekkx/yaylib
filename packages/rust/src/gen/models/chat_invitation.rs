// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ChatInvitation {
    #[serde(rename = "first_user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub first_user: Option<Option<String>>,
    #[serde(rename = "inviter", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub inviter: Option<Option<String>>,
    #[serde(rename = "users_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub users_count: Option<Option<i32>>,
}

impl ChatInvitation {
    pub fn new() -> ChatInvitation {
        ChatInvitation {
            first_user: None,
            inviter: None,
            users_count: None,
        }
    }
}

