// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PostSearchScope {
    #[serde(rename = "api_value", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub api_value: Option<Option<String>>,
    #[serde(rename = "title_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title_id: Option<Option<i32>>,
}

impl PostSearchScope {
    pub fn new() -> PostSearchScope {
        PostSearchScope {
            api_value: None,
            title_id: None,
        }
    }
}

