// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct TransactionGiftReceived {
    #[serde(rename = "icon", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub icon: Option<Option<String>>,
    #[serde(rename = "item", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub item: Option<Option<models::TransactionGiftReceivedItem>>,
    #[serde(rename = "item_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub item_id: Option<Option<i64>>,
    #[serde(rename = "quantity", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<Option<i32>>,
}

impl TransactionGiftReceived {
    pub fn new() -> TransactionGiftReceived {
        TransactionGiftReceived {
            icon: None,
            item: None,
            item_id: None,
            quantity: None,
        }
    }
}

