// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct StickerPacksResponse {
    #[serde(rename = "sticker_packs", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub sticker_packs: Option<Option<Vec<models::StickerPack>>>,
}

impl StickerPacksResponse {
    pub fn new() -> StickerPacksResponse {
        StickerPacksResponse {
            sticker_packs: None,
        }
    }
}

