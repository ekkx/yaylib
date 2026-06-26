// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct WalletResponse {
    #[serde(rename = "coins", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub coins: Option<Option<models::CoinAmount>>,
}

impl WalletResponse {
    pub fn new() -> WalletResponse {
        WalletResponse {
            coins: None,
        }
    }
}

