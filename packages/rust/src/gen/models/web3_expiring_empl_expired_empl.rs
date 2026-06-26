// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3ExpiringEmplExpiredEmpl {
    #[serde(rename = "amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub amount: Option<Option<f64>>,
    #[serde(rename = "date_in_millis", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub date_in_millis: Option<Option<i64>>,
    #[serde(rename = "soon", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub soon: Option<Option<bool>>,
}

impl Web3ExpiringEmplExpiredEmpl {
    pub fn new() -> Web3ExpiringEmplExpiredEmpl {
        Web3ExpiringEmplExpiredEmpl {
            amount: None,
            date_in_millis: None,
            soon: None,
        }
    }
}

