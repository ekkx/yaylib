// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct StickerPack {
    #[serde(rename = "cover", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub cover: Option<Option<String>>,
    #[serde(rename = "description", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub description: Option<Option<String>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub name: Option<Option<String>>,
    #[serde(rename = "order", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub order: Option<Option<i32>>,
    #[serde(rename = "stickers", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub stickers: Option<Option<Vec<models::Sticker>>>,
}

impl StickerPack {
    pub fn new() -> StickerPack {
        StickerPack {
            cover: None,
            description: None,
            id: None,
            name: None,
            order: None,
            stickers: None,
        }
    }
}

