// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct VideoProcessedData {
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "video_processed", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub video_processed: Option<Option<bool>>,
    #[serde(rename = "video_thumbnail_big_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub video_thumbnail_big_url: Option<Option<String>>,
    #[serde(rename = "video_thumbnail_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub video_thumbnail_url: Option<Option<String>>,
    #[serde(rename = "video_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub video_url: Option<Option<String>>,
}

impl VideoProcessedData {
    pub fn new() -> VideoProcessedData {
        VideoProcessedData {
            id: None,
            video_processed: None,
            video_thumbnail_big_url: None,
            video_thumbnail_url: None,
            video_url: None,
        }
    }
}

