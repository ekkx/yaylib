// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GroupUsersResponse {
    #[serde(rename = "group_users", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group_users: Option<Option<Vec<models::GroupUser>>>,
}

impl GroupUsersResponse {
    pub fn new() -> GroupUsersResponse {
        GroupUsersResponse {
            group_users: None,
        }
    }
}

