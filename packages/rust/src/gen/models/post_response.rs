// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PostResponse {
    #[serde(rename = "post", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub post: Option<Option<models::Post>>,
}

impl PostResponse {
    pub fn new() -> PostResponse {
        PostResponse {
            post: None,
        }
    }
}

