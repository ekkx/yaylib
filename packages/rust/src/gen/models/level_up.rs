// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct LevelUp {
    #[serde(rename = "level", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub level: Option<Option<i32>>,
    #[serde(rename = "pal_image_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal_image_url: Option<Option<String>>,
    #[serde(rename = "pal_name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal_name: Option<Option<String>>,
    #[serde(rename = "price", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub price: Option<Option<f64>>,
}

impl LevelUp {
    pub fn new() -> LevelUp {
        LevelUp {
            level: None,
            pal_image_url: None,
            pal_name: None,
            price: None,
        }
    }
}

