// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct UserMuted {
    #[serde(rename = "is_muted", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_muted: Option<Option<bool>>,
    #[serde(rename = "user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user: Option<Option<models::User>>,
}

impl UserMuted {
    pub fn new() -> UserMuted {
        UserMuted {
            is_muted: None,
            user: None,
        }
    }
}

