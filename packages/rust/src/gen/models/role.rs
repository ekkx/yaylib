// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Role {
    #[serde(rename = "is_moderator", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_moderator: Option<Option<bool>>,
    #[serde(rename = "is_owner", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_owner: Option<Option<bool>>,
}

impl Role {
    pub fn new() -> Role {
        Role {
            is_moderator: None,
            is_owner: None,
        }
    }
}

