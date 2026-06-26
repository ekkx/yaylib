// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PresignedUrlsResponse {
    #[serde(rename = "presigned_urls", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub presigned_urls: Option<Option<Vec<models::PresignedUrl>>>,
}

impl PresignedUrlsResponse {
    pub fn new() -> PresignedUrlsResponse {
        PresignedUrlsResponse {
            presigned_urls: None,
        }
    }
}

