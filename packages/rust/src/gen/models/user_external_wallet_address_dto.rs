// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct UserExternalWalletAddressDto {
    #[serde(rename = "external_wallet_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub external_wallet_address: Option<Option<String>>,
}

impl UserExternalWalletAddressDto {
    pub fn new() -> UserExternalWalletAddressDto {
        UserExternalWalletAddressDto {
            external_wallet_address: None,
        }
    }
}

