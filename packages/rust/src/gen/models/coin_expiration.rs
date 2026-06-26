// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CoinExpiration {
    #[serde(rename = "amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub amount: Option<Option<f64>>,
    #[serde(rename = "expired_at_seconds", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub expired_at_seconds: Option<Option<i64>>,
}

impl CoinExpiration {
    pub fn new() -> CoinExpiration {
        CoinExpiration {
            amount: None,
            expired_at_seconds: None,
        }
    }
}

