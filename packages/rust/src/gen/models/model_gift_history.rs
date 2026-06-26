// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ModelGiftHistory {
    #[serde(rename = "gift_item", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gift_item: Option<Option<models::Gift>>,
    #[serde(rename = "gifts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gifts: Option<Option<Vec<models::ReceivedGift>>>,
    #[serde(rename = "is_sender_anonymous", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_sender_anonymous: Option<Option<bool>>,
    #[serde(rename = "receiver", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub receiver: Option<Option<models::User>>,
    #[serde(rename = "sender", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub sender: Option<Option<models::User>>,
    #[serde(rename = "transaction_at_seconds", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub transaction_at_seconds: Option<Option<i64>>,
}

impl ModelGiftHistory {
    pub fn new() -> ModelGiftHistory {
        ModelGiftHistory {
            gift_item: None,
            gifts: None,
            is_sender_anonymous: None,
            receiver: None,
            sender: None,
            transaction_at_seconds: None,
        }
    }
}

