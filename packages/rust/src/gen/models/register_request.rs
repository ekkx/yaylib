// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct RegisterRequest {
    #[serde(rename = "signed_data_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub signed_data_type: Option<Option<String>>,
    #[serde(rename = "wallet_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub wallet_address: Option<Option<String>>,
}

impl RegisterRequest {
    pub fn new() -> RegisterRequest {
        RegisterRequest {
            signed_data_type: None,
            wallet_address: None,
        }
    }
}

