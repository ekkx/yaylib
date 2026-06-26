// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PostSearchQuery {
    #[serde(rename = "keyword", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub keyword: Option<Option<String>>,
    #[serde(rename = "only_media", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub only_media: Option<Option<bool>>,
    #[serde(rename = "scope", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub scope: Option<Option<models::PostSearchScope>>,
}

impl PostSearchQuery {
    pub fn new() -> PostSearchQuery {
        PostSearchQuery {
            keyword: None,
            only_media: None,
            scope: None,
        }
    }
}

