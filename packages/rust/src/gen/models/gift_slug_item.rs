// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GiftSlugItem {
    #[serde(rename = "icon", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub icon: Option<Option<String>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "quantity", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<Option<i32>>,
    #[serde(rename = "slug", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub slug: Option<Option<String>>,
}

impl GiftSlugItem {
    pub fn new() -> GiftSlugItem {
        GiftSlugItem {
            icon: None,
            id: None,
            quantity: None,
            slug: None,
        }
    }
}

