// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ModelWeb3WalletPreSign {
    #[serde(rename = "addition_gas_percent", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub addition_gas_percent: Option<Option<models::ModelWeb3WalletGasPercent>>,
    #[serde(rename = "chain_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub chain_id: Option<Option<i64>>,
    #[serde(rename = "chain_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub chain_url: Option<Option<String>>,
    #[serde(rename = "contract_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub contract_address: Option<Option<String>>,
    #[serde(rename = "inputs_hex", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub inputs_hex: Option<Option<String>>,
    #[serde(rename = "name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub name: Option<Option<String>>,
    #[serde(rename = "tx_unique_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub tx_unique_id: Option<Option<String>>,
}

impl ModelWeb3WalletPreSign {
    pub fn new() -> ModelWeb3WalletPreSign {
        ModelWeb3WalletPreSign {
            addition_gas_percent: None,
            chain_id: None,
            chain_url: None,
            contract_address: None,
            inputs_hex: None,
            name: None,
            tx_unique_id: None,
        }
    }
}

