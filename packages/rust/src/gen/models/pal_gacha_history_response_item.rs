// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PalGachaHistoryResponseItem {
    #[serde(rename = "created_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<Option<i64>>,
    #[serde(rename = "gacha_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gacha_type: Option<Option<String>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<String>>,
    #[serde(rename = "pal_image_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal_image_url: Option<Option<String>>,
    #[serde(rename = "pal_name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal_name: Option<Option<String>>,
    #[serde(rename = "pal_state", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal_state: Option<Option<String>>,
    #[serde(rename = "pal_token_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal_token_id: Option<Option<String>>,
    #[serde(rename = "price", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub price: Option<Option<i32>>,
    #[serde(rename = "rarity", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub rarity: Option<Option<String>>,
}

impl PalGachaHistoryResponseItem {
    pub fn new() -> PalGachaHistoryResponseItem {
        PalGachaHistoryResponseItem {
            created_at: None,
            gacha_type: None,
            id: None,
            pal_image_url: None,
            pal_name: None,
            pal_state: None,
            pal_token_id: None,
            price: None,
            rarity: None,
        }
    }
}

