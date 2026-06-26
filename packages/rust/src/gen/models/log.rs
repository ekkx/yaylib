// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Log {
    #[serde(rename = "address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub address: Option<Option<String>>,
    #[serde(rename = "block_hash", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub block_hash: Option<Option<String>>,
    #[serde(rename = "block_number", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub block_number: Option<Option<String>>,
    #[serde(rename = "data", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub data: Option<Option<String>>,
    #[serde(rename = "log_index", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub log_index: Option<Option<String>>,
    #[serde(rename = "removed", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub removed: Option<Option<bool>>,
    #[serde(rename = "topics", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub topics: Option<Option<Vec<String>>>,
    #[serde(rename = "transaction_hash", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub transaction_hash: Option<Option<String>>,
    #[serde(rename = "transaction_index", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub transaction_index: Option<Option<String>>,
}

impl Log {
    pub fn new() -> Log {
        Log {
            address: None,
            block_hash: None,
            block_number: None,
            data: None,
            log_index: None,
            removed: None,
            topics: None,
            transaction_hash: None,
            transaction_index: None,
        }
    }
}

