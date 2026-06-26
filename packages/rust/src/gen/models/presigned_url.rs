// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PresignedUrl {
    #[serde(rename = "filename", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub filename: Option<Option<String>>,
    #[serde(rename = "url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub url: Option<Option<String>>,
}

impl PresignedUrl {
    pub fn new() -> PresignedUrl {
        PresignedUrl {
            filename: None,
            url: None,
        }
    }
}

