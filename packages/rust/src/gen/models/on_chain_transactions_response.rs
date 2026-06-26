// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct OnChainTransactionsResponse {
    #[serde(rename = "cursor", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub cursor: Option<Option<String>>,
    #[serde(rename = "result", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub result: Option<Option<Vec<models::OnChainTransactionDto>>>,
}

impl OnChainTransactionsResponse {
    pub fn new() -> OnChainTransactionsResponse {
        OnChainTransactionsResponse {
            cursor: None,
            result: None,
        }
    }
}

