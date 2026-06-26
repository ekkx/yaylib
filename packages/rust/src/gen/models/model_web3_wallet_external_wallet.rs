// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ModelWeb3WalletExternalWallet {
    #[serde(rename = "eggs_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub eggs_count: Option<Option<i32>>,
    #[serde(rename = "pals_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pals_count: Option<Option<i32>>,
    #[serde(rename = "wallet_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub wallet_address: Option<Option<String>>,
}

impl ModelWeb3WalletExternalWallet {
    pub fn new() -> ModelWeb3WalletExternalWallet {
        ModelWeb3WalletExternalWallet {
            eggs_count: None,
            pals_count: None,
            wallet_address: None,
        }
    }
}

