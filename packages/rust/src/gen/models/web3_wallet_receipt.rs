// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3WalletReceipt {
    #[serde(rename = "block_hash", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub block_hash: Option<Option<String>>,
    #[serde(rename = "block_number", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub block_number: Option<Option<String>>,
    #[serde(rename = "contract_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub contract_address: Option<Option<String>>,
    #[serde(rename = "cumulative_gas_used", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub cumulative_gas_used: Option<Option<String>>,
    #[serde(rename = "gas_used", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gas_used: Option<Option<String>>,
    #[serde(rename = "logs", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub logs: Option<Option<Vec<models::Log>>>,
    #[serde(rename = "logs_bloom", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub logs_bloom: Option<Option<String>>,
    #[serde(rename = "root", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub root: Option<Option<String>>,
    #[serde(rename = "status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub status: Option<Option<String>>,
    #[serde(rename = "transaction_hash", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub transaction_hash: Option<Option<String>>,
    #[serde(rename = "transaction_index", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub transaction_index: Option<Option<String>>,
}

impl Web3WalletReceipt {
    pub fn new() -> Web3WalletReceipt {
        Web3WalletReceipt {
            block_hash: None,
            block_number: None,
            contract_address: None,
            cumulative_gas_used: None,
            gas_used: None,
            logs: None,
            logs_bloom: None,
            root: None,
            status: None,
            transaction_hash: None,
            transaction_index: None,
        }
    }
}

