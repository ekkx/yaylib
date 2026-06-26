// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ModelVideo {
    #[serde(rename = "bitrate", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub bitrate: Option<Option<i32>>,
    #[serde(rename = "height", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub height: Option<Option<i32>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "is_completed", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_completed: Option<Option<bool>>,
    #[serde(rename = "thumbnail_big_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub thumbnail_big_url: Option<Option<String>>,
    #[serde(rename = "thumbnail_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub thumbnail_url: Option<Option<String>>,
    #[serde(rename = "video_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub video_url: Option<Option<String>>,
    #[serde(rename = "views_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub views_count: Option<Option<i32>>,
    #[serde(rename = "width", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub width: Option<Option<i32>>,
}

impl ModelVideo {
    pub fn new() -> ModelVideo {
        ModelVideo {
            bitrate: None,
            height: None,
            id: None,
            is_completed: None,
            thumbnail_big_url: None,
            thumbnail_url: None,
            video_url: None,
            views_count: None,
            width: None,
        }
    }
}

