// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3WalletLocalHistoryTransaction {
    #[serde(rename = "amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub amount: Option<Option<f64>>,
    #[serde(rename = "asset_type_name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub asset_type_name: Option<Option<String>>,
    #[serde(rename = "chain", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub chain: Option<Option<String>>,
    #[serde(rename = "chain_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub chain_id: Option<Option<i64>>,
    #[serde(rename = "chain_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub chain_url: Option<Option<String>>,
    #[serde(rename = "contract_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub contract_address: Option<Option<String>>,
    #[serde(rename = "gas_limit", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gas_limit: Option<Option<i64>>,
    #[serde(rename = "gas_percentage", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gas_percentage: Option<Option<i32>>,
    #[serde(rename = "gas_price", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gas_price: Option<Option<f64>>,
    #[serde(rename = "inputs_hex", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub inputs_hex: Option<Option<String>>,
    #[serde(rename = "nonce", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub nonce: Option<Option<i64>>,
    #[serde(rename = "pal_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal_id: Option<Option<String>>,
    #[serde(rename = "timestamp", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub timestamp: Option<Option<i64>>,
    #[serde(rename = "transaction_hash_hex", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub transaction_hash_hex: Option<Option<String>>,
    #[serde(rename = "transaction_history_status_name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub transaction_history_status_name: Option<Option<String>>,
    #[serde(rename = "transfer_mode_key", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub transfer_mode_key: Option<Option<String>>,
    #[serde(rename = "tx_unique_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub tx_unique_id: Option<Option<String>>,
}

impl Web3WalletLocalHistoryTransaction {
    pub fn new() -> Web3WalletLocalHistoryTransaction {
        Web3WalletLocalHistoryTransaction {
            amount: None,
            asset_type_name: None,
            chain: None,
            chain_id: None,
            chain_url: None,
            contract_address: None,
            gas_limit: None,
            gas_percentage: None,
            gas_price: None,
            inputs_hex: None,
            nonce: None,
            pal_id: None,
            timestamp: None,
            transaction_hash_hex: None,
            transaction_history_status_name: None,
            transfer_mode_key: None,
            tx_unique_id: None,
        }
    }
}

