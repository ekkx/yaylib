// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CallStatusResponse {
    #[serde(rename = "phone_status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub phone_status: Option<Option<bool>>,
    #[serde(rename = "room_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub room_url: Option<Option<String>>,
    #[serde(rename = "video_status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub video_status: Option<Option<bool>>,
}

impl CallStatusResponse {
    pub fn new() -> CallStatusResponse {
        CallStatusResponse {
            phone_status: None,
            room_url: None,
            video_status: None,
        }
    }
}

