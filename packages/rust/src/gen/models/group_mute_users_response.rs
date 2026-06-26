// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GroupMuteUsersResponse {
    #[serde(rename = "group_users", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group_users: Option<Option<Vec<models::GroupUser>>>,
    #[serde(rename = "next_cursor", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub next_cursor: Option<Option<i64>>,
    #[serde(rename = "total", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total: Option<Option<i32>>,
}

impl GroupMuteUsersResponse {
    pub fn new() -> GroupMuteUsersResponse {
        GroupMuteUsersResponse {
            group_users: None,
            next_cursor: None,
            total: None,
        }
    }
}

