// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GiftReceivedResponse {
    #[serde(rename = "next_page_value", default, skip_serializing_if = "Option::is_none", deserialize_with = "crate::gen::cursor::deserialize_option_string_or_number")]
    pub next_page_value: Option<String>,
    #[serde(rename = "received_gifts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub received_gifts: Option<Option<Vec<models::ReceivedGift>>>,
    #[serde(rename = "total_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total_count: Option<Option<i32>>,
}

impl GiftReceivedResponse {
    pub fn new() -> GiftReceivedResponse {
        GiftReceivedResponse {
            next_page_value: None,
            received_gifts: None,
            total_count: None,
        }
    }
}

