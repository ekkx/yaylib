// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PostsResponse {
    #[serde(rename = "has_more_hot_posts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub has_more_hot_posts: Option<Option<bool>>,
    #[serde(rename = "next_page_value", default, skip_serializing_if = "Option::is_none", deserialize_with = "crate::gen::cursor::deserialize_option_string_or_number")]
    pub next_page_value: Option<String>,
    #[serde(rename = "pinned_posts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pinned_posts: Option<Option<Vec<models::Post>>>,
    #[serde(rename = "posts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub posts: Option<Option<Vec<models::Post>>>,
}

impl PostsResponse {
    pub fn new() -> PostsResponse {
        PostsResponse {
            has_more_hot_posts: None,
            next_page_value: None,
            pinned_posts: None,
            posts: None,
        }
    }
}

