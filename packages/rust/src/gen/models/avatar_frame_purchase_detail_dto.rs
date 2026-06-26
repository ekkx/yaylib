// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct AvatarFramePurchaseDetailDto {
    #[serde(rename = "frame", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub frame: Option<Option<String>>,
    #[serde(rename = "frame_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub frame_thumbnail: Option<Option<String>>,
    #[serde(rename = "frame_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub frame_type: Option<Option<String>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "vip_only", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub vip_only: Option<Option<bool>>,
}

impl AvatarFramePurchaseDetailDto {
    pub fn new() -> AvatarFramePurchaseDetailDto {
        AvatarFramePurchaseDetailDto {
            frame: None,
            frame_thumbnail: None,
            frame_type: None,
            id: None,
            vip_only: None,
        }
    }
}

