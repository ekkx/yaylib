// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GroupsRelatedResponse {
    #[serde(rename = "groups", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub groups: Option<Option<Vec<models::Group>>>,
    #[serde(rename = "next_page_value", default, skip_serializing_if = "Option::is_none", deserialize_with = "crate::gen::cursor::deserialize_option_string_or_number")]
    pub next_page_value: Option<String>,
}

impl GroupsRelatedResponse {
    pub fn new() -> GroupsRelatedResponse {
        GroupsRelatedResponse {
            groups: None,
            next_page_value: None,
        }
    }
}

