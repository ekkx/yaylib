// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GiftTransactionsResponse {
    #[serde(rename = "hide_gifts_received", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hide_gifts_received: Option<Option<bool>>,
    #[serde(rename = "next_page_value", default, skip_serializing_if = "Option::is_none", deserialize_with = "crate::gen::cursor::deserialize_option_string_or_number")]
    pub next_page_value: Option<String>,
    #[serde(rename = "sent_gifts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub sent_gifts: Option<Option<Vec<models::GiftHistory>>>,
}

impl GiftTransactionsResponse {
    pub fn new() -> GiftTransactionsResponse {
        GiftTransactionsResponse {
            hide_gifts_received: None,
            next_page_value: None,
            sent_gifts: None,
        }
    }
}

