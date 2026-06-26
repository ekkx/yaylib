// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3WalletExternalWallet {
    #[serde(rename = "balance", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub balance: Option<Option<models::Balance>>,
    #[serde(rename = "wallet_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub wallet_address: Option<Option<String>>,
}

impl Web3WalletExternalWallet {
    pub fn new() -> Web3WalletExternalWallet {
        Web3WalletExternalWallet {
            balance: None,
            wallet_address: None,
        }
    }
}

