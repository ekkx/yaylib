// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PurchaseGiftRequest {
    #[serde(rename = "gifts", skip_serializing_if = "Option::is_none")]
    pub gifts: Option<std::collections::HashMap<String, i32>>,
    #[serde(rename = "is_anonymously", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_anonymously: Option<Option<bool>>,
    #[serde(rename = "receiver_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub receiver_id: Option<Option<i64>>,
}

impl PurchaseGiftRequest {
    pub fn new() -> PurchaseGiftRequest {
        PurchaseGiftRequest {
            gifts: None,
            is_anonymously: None,
            receiver_id: None,
        }
    }
}

