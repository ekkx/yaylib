// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GroupUserResponse {
    #[serde(rename = "group_user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group_user: Option<Option<models::GroupUser>>,
}

impl GroupUserResponse {
    pub fn new() -> GroupUserResponse {
        GroupUserResponse {
            group_user: None,
        }
    }
}

