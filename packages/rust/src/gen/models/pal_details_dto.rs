// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PalDetailsDto {
    #[serde(rename = "cost", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub cost: Option<Option<f64>>,
    #[serde(rename = "level", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub level: Option<Option<i32>>,
    #[serde(rename = "nft_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub nft_type: Option<Option<String>>,
    #[serde(rename = "pal_image", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal_image: Option<Option<String>>,
    #[serde(rename = "pal_name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal_name: Option<Option<String>>,
    #[serde(rename = "rank", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub rank: Option<Option<String>>,
    #[serde(rename = "token_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_id: Option<Option<String>>,
}

impl PalDetailsDto {
    pub fn new() -> PalDetailsDto {
        PalDetailsDto {
            cost: None,
            level: None,
            nft_type: None,
            pal_image: None,
            pal_name: None,
            rank: None,
            token_id: None,
        }
    }
}

