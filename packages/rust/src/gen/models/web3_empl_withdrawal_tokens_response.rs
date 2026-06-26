// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3EmplWithdrawalTokensResponse {
    #[serde(rename = "result", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub result: Option<Option<models::Web3EmplWithdrawalTokensDto>>,
}

impl Web3EmplWithdrawalTokensResponse {
    pub fn new() -> Web3EmplWithdrawalTokensResponse {
        Web3EmplWithdrawalTokensResponse {
            result: None,
        }
    }
}

