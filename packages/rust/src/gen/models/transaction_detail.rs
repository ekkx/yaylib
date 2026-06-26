// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct TransactionDetail {
    #[serde(rename = "result", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub result: Option<Option<models::Result>>,
}

impl TransactionDetail {
    pub fn new() -> TransactionDetail {
        TransactionDetail {
            result: None,
        }
    }
}

