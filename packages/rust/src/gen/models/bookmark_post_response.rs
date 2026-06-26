// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct BookmarkPostResponse {
    #[serde(rename = "is_bookmarked", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_bookmarked: Option<Option<bool>>,
}

impl BookmarkPostResponse {
    pub fn new() -> BookmarkPostResponse {
        BookmarkPostResponse {
            is_bookmarked: None,
        }
    }
}

