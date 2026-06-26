// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CreateGroupResponse {
    #[serde(rename = "group_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group_id: Option<Option<i64>>,
}

impl CreateGroupResponse {
    pub fn new() -> CreateGroupResponse {
        CreateGroupResponse {
            group_id: None,
        }
    }
}

