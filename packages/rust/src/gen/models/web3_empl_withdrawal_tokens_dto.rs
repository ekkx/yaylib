// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3EmplWithdrawalTokensDto {
    #[serde(rename = "tokens", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub tokens: Option<Option<Vec<models::Web3EmplWithdrawalTokenDto>>>,
}

impl Web3EmplWithdrawalTokensDto {
    pub fn new() -> Web3EmplWithdrawalTokensDto {
        Web3EmplWithdrawalTokensDto {
            tokens: None,
        }
    }
}

