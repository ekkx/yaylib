// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct TwoFaAuthRequiredDto {
    #[serde(rename = "login", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub login: Option<Option<bool>>,
    #[serde(rename = "wallet", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub wallet: Option<Option<bool>>,
}

impl TwoFaAuthRequiredDto {
    pub fn new() -> TwoFaAuthRequiredDto {
        TwoFaAuthRequiredDto {
            login: None,
            wallet: None,
        }
    }
}

