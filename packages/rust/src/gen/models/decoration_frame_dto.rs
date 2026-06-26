// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct DecorationFrameDto {
    #[serde(rename = "current", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub current: Option<Option<bool>>,
    #[serde(rename = "frame", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub frame: Option<Option<String>>,
    #[serde(rename = "frame_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub frame_thumbnail: Option<Option<String>>,
    #[serde(rename = "frame_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub frame_type: Option<Option<String>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "owned", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub owned: Option<Option<bool>>,
    #[serde(rename = "price", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub price: Option<Option<i64>>,
    #[serde(rename = "vip_only", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub vip_only: Option<Option<bool>>,
}

impl DecorationFrameDto {
    pub fn new() -> DecorationFrameDto {
        DecorationFrameDto {
            current: None,
            frame: None,
            frame_thumbnail: None,
            frame_type: None,
            id: None,
            owned: None,
            price: None,
            vip_only: None,
        }
    }
}

