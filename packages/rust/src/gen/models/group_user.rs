// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GroupUser {
    #[serde(rename = "is_moderator", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_moderator: Option<Option<bool>>,
    #[serde(rename = "is_mute", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_mute: Option<Option<bool>>,
    #[serde(rename = "pending_deputize", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pending_deputize: Option<Option<bool>>,
    #[serde(rename = "pending_transfer", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pending_transfer: Option<Option<bool>>,
    #[serde(rename = "title", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title: Option<Option<String>>,
    #[serde(rename = "user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user: Option<Option<Box<models::RealmUser>>>,
}

impl GroupUser {
    pub fn new() -> GroupUser {
        GroupUser {
            is_moderator: None,
            is_mute: None,
            pending_deputize: None,
            pending_transfer: None,
            title: None,
            user: None,
        }
    }
}

