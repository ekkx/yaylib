// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GifImage {
    #[serde(rename = "height", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub height: Option<Option<i32>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub url: Option<Option<String>>,
    #[serde(rename = "width", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub width: Option<Option<i32>>,
}

impl GifImage {
    pub fn new() -> GifImage {
        GifImage {
            height: None,
            id: None,
            url: None,
            width: None,
        }
    }
}

