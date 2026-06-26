// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GiftReceivedTransactionResponse {
    #[serde(rename = "created_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<Option<i64>>,
    #[serde(rename = "gifts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gifts: Option<Option<Vec<models::TransactionGiftReceived>>>,
    #[serde(rename = "receiver", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub receiver: Option<Option<models::RealmUser>>,
}

impl GiftReceivedTransactionResponse {
    pub fn new() -> GiftReceivedTransactionResponse {
        GiftReceivedTransactionResponse {
            created_at: None,
            gifts: None,
            receiver: None,
        }
    }
}

