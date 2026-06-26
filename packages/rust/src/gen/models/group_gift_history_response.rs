// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GroupGiftHistoryResponse {
    #[serde(rename = "gift_history", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gift_history: Option<Option<Vec<models::GroupGiftHistory>>>,
    #[serde(rename = "next_page_value", default, skip_serializing_if = "Option::is_none", deserialize_with = "crate::gen::cursor::deserialize_option_string_or_number")]
    pub next_page_value: Option<String>,
}

impl GroupGiftHistoryResponse {
    pub fn new() -> GroupGiftHistoryResponse {
        GroupGiftHistoryResponse {
            gift_history: None,
            next_page_value: None,
        }
    }
}

