// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Sticker {
    #[serde(rename = "extension", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub extension: Option<Option<String>>,
    #[serde(rename = "height", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub height: Option<Option<i32>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "sticker_pack_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub sticker_pack_id: Option<Option<i64>>,
    #[serde(rename = "url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub url: Option<Option<String>>,
    #[serde(rename = "width", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub width: Option<Option<i32>>,
}

impl Sticker {
    pub fn new() -> Sticker {
        Sticker {
            extension: None,
            height: None,
            id: None,
            sticker_pack_id: None,
            url: None,
            width: None,
        }
    }
}

