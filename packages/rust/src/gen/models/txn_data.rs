// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct TxnData {
    #[serde(rename = "contract_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub contract_address: Option<Option<String>>,
    #[serde(rename = "inputs_hex", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub inputs_hex: Option<Option<String>>,
}

impl TxnData {
    pub fn new() -> TxnData {
        TxnData {
            contract_address: None,
            inputs_hex: None,
        }
    }
}

