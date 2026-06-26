// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Evolution {
    #[serde(rename = "grade", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub grade: Option<Option<serde_json::Value>>,
    #[serde(rename = "level", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub level: Option<Option<i32>>,
    #[serde(rename = "pal_image_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal_image_url: Option<Option<String>>,
    #[serde(rename = "pal_name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal_name: Option<Option<String>>,
}

impl Evolution {
    pub fn new() -> Evolution {
        Evolution {
            grade: None,
            level: None,
            pal_image_url: None,
            pal_name: None,
        }
    }
}

