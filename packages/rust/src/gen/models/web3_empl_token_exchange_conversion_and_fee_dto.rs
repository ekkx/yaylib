// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3EmplTokenExchangeConversionAndFeeDto {
    #[serde(rename = "empl_amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub empl_amount: Option<Option<String>>,
    #[serde(rename = "empl_tx_fee", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub empl_tx_fee: Option<Option<String>>,
    #[serde(rename = "max_amount_out", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub max_amount_out: Option<Option<String>>,
    #[serde(rename = "min_amount_out", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub min_amount_out: Option<Option<String>>,
    #[serde(rename = "token_out_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_out_address: Option<Option<String>>,
}

impl Web3EmplTokenExchangeConversionAndFeeDto {
    pub fn new() -> Web3EmplTokenExchangeConversionAndFeeDto {
        Web3EmplTokenExchangeConversionAndFeeDto {
            empl_amount: None,
            empl_tx_fee: None,
            max_amount_out: None,
            min_amount_out: None,
            token_out_address: None,
        }
    }
}

