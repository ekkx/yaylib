// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct TxnLogData {
    #[serde(rename = "contract_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub contract_address: Option<Option<String>>,
    #[serde(rename = "inputs_hex", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub inputs_hex: Option<Option<String>>,
    #[serde(rename = "tx_unique_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub tx_unique_id: Option<Option<String>>,
}

impl TxnLogData {
    pub fn new() -> TxnLogData {
        TxnLogData {
            contract_address: None,
            inputs_hex: None,
            tx_unique_id: None,
        }
    }
}

