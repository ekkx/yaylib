// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct LikePostsResponse {
    #[serde(rename = "like_ids", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub like_ids: Option<Option<Vec<i64>>>,
}

impl LikePostsResponse {
    pub fn new() -> LikePostsResponse {
        LikePostsResponse {
            like_ids: None,
        }
    }
}

