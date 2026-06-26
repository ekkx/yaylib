// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ReceivedGift {
    #[serde(rename = "gift", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gift: Option<Option<models::Gift>>,
    #[serde(rename = "received_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub received_count: Option<Option<i32>>,
    #[serde(rename = "senders", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub senders: Option<Option<Vec<models::User>>>,
    #[serde(rename = "total_senders_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total_senders_count: Option<Option<i32>>,
}

impl ReceivedGift {
    pub fn new() -> ReceivedGift {
        ReceivedGift {
            gift: None,
            received_count: None,
            senders: None,
            total_senders_count: None,
        }
    }
}

