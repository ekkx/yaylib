// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3WalletGasPercent {
    #[serde(rename = "fast", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub fast: Option<Option<i32>>,
    #[serde(rename = "normal", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub normal: Option<Option<i32>>,
    #[serde(rename = "slow", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub slow: Option<Option<i32>>,
}

impl Web3WalletGasPercent {
    pub fn new() -> Web3WalletGasPercent {
        Web3WalletGasPercent {
            fast: None,
            normal: None,
            slow: None,
        }
    }
}

