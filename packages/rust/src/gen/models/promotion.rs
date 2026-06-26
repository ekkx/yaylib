// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Promotion {
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "image_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub image_url: Option<Option<String>>,
    #[serde(rename = "order", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub order: Option<Option<i32>>,
    #[serde(rename = "promotion_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub promotion_url: Option<Option<String>>,
    #[serde(rename = "title", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title: Option<Option<String>>,
}

impl Promotion {
    pub fn new() -> Promotion {
        Promotion {
            id: None,
            image_url: None,
            order: None,
            promotion_url: None,
            title: None,
        }
    }
}

