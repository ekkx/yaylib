// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3WalletBlockStoreItem {
    #[serde(rename = "created_at_millis", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at_millis: Option<Option<i64>>,
    #[serde(rename = "name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub name: Option<Option<String>>,
    #[serde(rename = "private_key", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub private_key: Option<Option<String>>,
    #[serde(rename = "seed_phrase", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub seed_phrase: Option<Option<String>>,
    #[serde(rename = "wallet_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub wallet_address: Option<Option<String>>,
}

impl Web3WalletBlockStoreItem {
    pub fn new() -> Web3WalletBlockStoreItem {
        Web3WalletBlockStoreItem {
            created_at_millis: None,
            name: None,
            private_key: None,
            seed_phrase: None,
            wallet_address: None,
        }
    }
}

