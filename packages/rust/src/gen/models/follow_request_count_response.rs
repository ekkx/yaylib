// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct FollowRequestCountResponse {
    #[serde(rename = "users_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub users_count: Option<Option<i32>>,
}

impl FollowRequestCountResponse {
    pub fn new() -> FollowRequestCountResponse {
        FollowRequestCountResponse {
            users_count: None,
        }
    }
}

