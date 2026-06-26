// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Network {
    #[serde(rename = "addition_gas_percent", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub addition_gas_percent: Option<Option<models::ModelWeb3WalletGasPercent>>,
    #[serde(rename = "amm_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub amm_address: Option<Option<String>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub name: Option<Option<String>>,
    #[serde(rename = "quoter_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub quoter_address: Option<Option<String>>,
    #[serde(rename = "url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub url: Option<Option<String>>,
    #[serde(rename = "weth_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub weth_address: Option<Option<String>>,
    #[serde(rename = "yay_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub yay_address: Option<Option<String>>,
}

impl Network {
    pub fn new() -> Network {
        Network {
            addition_gas_percent: None,
            amm_address: None,
            id: None,
            name: None,
            quoter_address: None,
            url: None,
            weth_address: None,
            yay_address: None,
        }
    }
}

