// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3WalletGasFee {
    #[serde(rename = "gas_limit", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gas_limit: Option<Option<i64>>,
    #[serde(rename = "gas_price_wei", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gas_price_wei: Option<Option<f64>>,
}

impl Web3WalletGasFee {
    pub fn new() -> Web3WalletGasFee {
        Web3WalletGasFee {
            gas_limit: None,
            gas_price_wei: None,
        }
    }
}

