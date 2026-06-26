// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CallGiftHistoryResponse {
    #[serde(rename = "next_page_value", default, skip_serializing_if = "Option::is_none", deserialize_with = "crate::gen::cursor::deserialize_option_string_or_number")]
    pub next_page_value: Option<String>,
    #[serde(rename = "sent_gifts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub sent_gifts: Option<Option<Vec<models::GiftHistory>>>,
}

impl CallGiftHistoryResponse {
    pub fn new() -> CallGiftHistoryResponse {
        CallGiftHistoryResponse {
            next_page_value: None,
            sent_gifts: None,
        }
    }
}

