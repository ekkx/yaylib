// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GiftsResponse {
    #[serde(rename = "gifts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gifts: Option<Option<Vec<models::RealmGift>>>,
    #[serde(rename = "next_page_value", default, skip_serializing_if = "Option::is_none", deserialize_with = "crate::gen::cursor::deserialize_option_string_or_number")]
    pub next_page_value: Option<String>,
    #[serde(rename = "total_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total_count: Option<Option<i32>>,
}

impl GiftsResponse {
    pub fn new() -> GiftsResponse {
        GiftsResponse {
            gifts: None,
            next_page_value: None,
            total_count: None,
        }
    }
}

