// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct SignaturePayload {
    #[serde(rename = "action", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub action: Option<Option<String>>,
    #[serde(rename = "call_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub call_id: Option<Option<i64>>,
    #[serde(rename = "receiver_uuid", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub receiver_uuid: Option<Option<String>>,
    #[serde(rename = "sender_uuid", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub sender_uuid: Option<Option<String>>,
    #[serde(rename = "signature", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub signature: Option<Option<String>>,
    #[serde(rename = "timestamp", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub timestamp: Option<Option<i64>>,
}

impl SignaturePayload {
    pub fn new() -> SignaturePayload {
        SignaturePayload {
            action: None,
            call_id: None,
            receiver_uuid: None,
            sender_uuid: None,
            signature: None,
            timestamp: None,
        }
    }
}

