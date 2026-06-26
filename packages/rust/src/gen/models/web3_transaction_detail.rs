// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3TransactionDetail {
    #[serde(rename = "amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub amount: Option<Option<String>>,
    #[serde(rename = "chain_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub chain_id: Option<Option<i64>>,
    #[serde(rename = "currency", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub currency: Option<Option<String>>,
    #[serde(rename = "empl_amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub empl_amount: Option<Option<f64>>,
    #[serde(rename = "error", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub error: Option<Option<serde_json::Value>>,
    #[serde(rename = "final_status_updated_at_millis", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub final_status_updated_at_millis: Option<Option<i64>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "max_token_amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub max_token_amount: Option<Option<String>>,
    #[serde(rename = "min_token_amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub min_token_amount: Option<Option<String>>,
    #[serde(rename = "status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub status: Option<Option<serde_json::Value>>,
    #[serde(rename = "token_out_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_out_address: Option<Option<String>>,
    #[serde(rename = "token_out_amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_out_amount: Option<Option<String>>,
    #[serde(rename = "token_symbol", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_symbol: Option<Option<String>>,
    #[serde(rename = "transaction_fee", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub transaction_fee: Option<Option<f64>>,
    #[serde(rename = "wallet_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub wallet_address: Option<Option<String>>,
}

impl Web3TransactionDetail {
    pub fn new() -> Web3TransactionDetail {
        Web3TransactionDetail {
            amount: None,
            chain_id: None,
            currency: None,
            empl_amount: None,
            error: None,
            final_status_updated_at_millis: None,
            id: None,
            max_token_amount: None,
            min_token_amount: None,
            status: None,
            token_out_address: None,
            token_out_amount: None,
            token_symbol: None,
            transaction_fee: None,
            wallet_address: None,
        }
    }
}

