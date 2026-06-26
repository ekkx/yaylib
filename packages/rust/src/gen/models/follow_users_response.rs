// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct FollowUsersResponse {
    #[serde(rename = "next_page_value", default, skip_serializing_if = "Option::is_none", deserialize_with = "crate::gen::cursor::deserialize_option_string_or_number")]
    pub next_page_value: Option<String>,
    #[serde(rename = "users", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub users: Option<Option<Vec<models::RealmUser>>>,
}

impl FollowUsersResponse {
    pub fn new() -> FollowUsersResponse {
        FollowUsersResponse {
            next_page_value: None,
            users: None,
        }
    }
}

