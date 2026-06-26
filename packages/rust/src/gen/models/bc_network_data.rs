// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct BcNetworkData {
    #[serde(rename = "addition_gas_percent", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub addition_gas_percent: Option<Option<models::Web3WalletGasPercent>>,
    #[serde(rename = "chain_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub chain_id: Option<Option<i64>>,
    #[serde(rename = "chain_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub chain_url: Option<Option<String>>,
    #[serde(rename = "name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub name: Option<Option<String>>,
}

impl BcNetworkData {
    pub fn new() -> BcNetworkData {
        BcNetworkData {
            addition_gas_percent: None,
            chain_id: None,
            chain_url: None,
            name: None,
        }
    }
}

