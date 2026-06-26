// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GiftCardActivityDetails {
    #[serde(rename = "created_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<Option<i64>>,
    #[serde(rename = "email", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub email: Option<Option<String>>,
    #[serde(rename = "empl_amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub empl_amount: Option<Option<i32>>,
    #[serde(rename = "gift_card_amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gift_card_amount: Option<Option<i32>>,
    #[serde(rename = "status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub status: Option<Option<serde_json::Value>>,
}

impl GiftCardActivityDetails {
    pub fn new() -> GiftCardActivityDetails {
        GiftCardActivityDetails {
            created_at: None,
            email: None,
            empl_amount: None,
            gift_card_amount: None,
            status: None,
        }
    }
}

