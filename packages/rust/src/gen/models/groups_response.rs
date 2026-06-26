// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GroupsResponse {
    #[serde(rename = "groups", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub groups: Option<Option<Vec<models::Group>>>,
    #[serde(rename = "pinned_groups", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pinned_groups: Option<Option<Vec<models::Group>>>,
}

impl GroupsResponse {
    pub fn new() -> GroupsResponse {
        GroupsResponse {
            groups: None,
            pinned_groups: None,
        }
    }
}

