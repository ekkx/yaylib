// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GiftSendersResponse {
    #[serde(rename = "senders", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub senders: Option<Option<Vec<models::RealmUser>>>,
    #[serde(rename = "total_senders_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total_senders_count: Option<Option<i32>>,
}

impl GiftSendersResponse {
    pub fn new() -> GiftSendersResponse {
        GiftSendersResponse {
            senders: None,
            total_senders_count: None,
        }
    }
}

