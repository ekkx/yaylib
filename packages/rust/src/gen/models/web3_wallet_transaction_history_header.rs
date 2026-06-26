// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3WalletTransactionHistoryHeader {
    #[serde(rename = "title_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title_id: Option<Option<i32>>,
}

impl Web3WalletTransactionHistoryHeader {
    pub fn new() -> Web3WalletTransactionHistoryHeader {
        Web3WalletTransactionHistoryHeader {
            title_id: None,
        }
    }
}

