// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct EmplTokenExchangeDetails {
    #[serde(rename = "chain_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub chain_id: Option<Option<i64>>,
    #[serde(rename = "max_token_amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub max_token_amount: Option<Option<String>>,
    #[serde(rename = "min_token_amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub min_token_amount: Option<Option<String>>,
    #[serde(rename = "token_out_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_out_address: Option<Option<String>>,
    #[serde(rename = "token_symbol", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_symbol: Option<Option<String>>,
    #[serde(rename = "transaction_fee", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub transaction_fee: Option<Option<f64>>,
}

impl EmplTokenExchangeDetails {
    pub fn new() -> EmplTokenExchangeDetails {
        EmplTokenExchangeDetails {
            chain_id: None,
            max_token_amount: None,
            min_token_amount: None,
            token_out_address: None,
            token_symbol: None,
            transaction_fee: None,
        }
    }
}

