// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3WalletBlockStoreEntry {
    #[serde(rename = "items", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub items: Option<Option<Vec<models::Web3WalletBlockStoreItem>>>,
}

impl Web3WalletBlockStoreEntry {
    pub fn new() -> Web3WalletBlockStoreEntry {
        Web3WalletBlockStoreEntry {
            items: None,
        }
    }
}

