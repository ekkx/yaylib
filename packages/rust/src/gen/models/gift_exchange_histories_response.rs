// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GiftExchangeHistoriesResponse {
    #[serde(rename = "gift_exchange_histories", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gift_exchange_histories: Option<Option<Vec<models::GiftExchangeHistory>>>,
    #[serde(rename = "next_page_value", default, skip_serializing_if = "Option::is_none", deserialize_with = "crate::gen::cursor::deserialize_option_string_or_number")]
    pub next_page_value: Option<String>,
}

impl GiftExchangeHistoriesResponse {
    pub fn new() -> GiftExchangeHistoriesResponse {
        GiftExchangeHistoriesResponse {
            gift_exchange_histories: None,
            next_page_value: None,
        }
    }
}

