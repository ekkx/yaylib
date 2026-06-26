// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ModelPostTag {
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "post_hashtags_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub post_hashtags_count: Option<Option<i32>>,
    #[serde(rename = "tag", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub tag: Option<Option<String>>,
}

impl ModelPostTag {
    pub fn new() -> ModelPostTag {
        ModelPostTag {
            id: None,
            post_hashtags_count: None,
            tag: None,
        }
    }
}

