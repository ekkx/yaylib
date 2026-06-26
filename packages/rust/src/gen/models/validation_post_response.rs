// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ValidationPostResponse {
    #[serde(rename = "is_allow_to_post", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_allow_to_post: Option<Option<bool>>,
}

impl ValidationPostResponse {
    pub fn new() -> ValidationPostResponse {
        ValidationPostResponse {
            is_allow_to_post: None,
        }
    }
}

