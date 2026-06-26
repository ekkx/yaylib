// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PostTagsResponse {
    #[serde(rename = "tags", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub tags: Option<Option<Vec<models::PostTag>>>,
}

impl PostTagsResponse {
    pub fn new() -> PostTagsResponse {
        PostTagsResponse {
            tags: None,
        }
    }
}

