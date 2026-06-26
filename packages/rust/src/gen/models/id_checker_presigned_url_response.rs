// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct IdCheckerPresignedUrlResponse {
    #[serde(rename = "presigned_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub presigned_url: Option<Option<String>>,
}

impl IdCheckerPresignedUrlResponse {
    pub fn new() -> IdCheckerPresignedUrlResponse {
        IdCheckerPresignedUrlResponse {
            presigned_url: None,
        }
    }
}

