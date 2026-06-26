// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3WalletEmplExpiring {
    #[serde(rename = "expiring_amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub expiring_amount: Option<Option<f64>>,
    #[serde(rename = "expiring_date_millis", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub expiring_date_millis: Option<Option<i64>>,
}

impl Web3WalletEmplExpiring {
    pub fn new() -> Web3WalletEmplExpiring {
        Web3WalletEmplExpiring {
            expiring_amount: None,
            expiring_date_millis: None,
        }
    }
}

