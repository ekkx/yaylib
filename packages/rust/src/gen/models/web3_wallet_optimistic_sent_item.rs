// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3WalletOptimisticSentItem {
    #[serde(rename = "chain_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub chain_id: Option<Option<i64>>,
    #[serde(rename = "contract_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub contract_address: Option<Option<String>>,
    #[serde(rename = "created_at_millis", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at_millis: Option<Option<i64>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<String>>,
    #[serde(rename = "token_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_id: Option<Option<String>>,
}

impl Web3WalletOptimisticSentItem {
    pub fn new() -> Web3WalletOptimisticSentItem {
        Web3WalletOptimisticSentItem {
            chain_id: None,
            contract_address: None,
            created_at_millis: None,
            id: None,
            token_id: None,
        }
    }
}

