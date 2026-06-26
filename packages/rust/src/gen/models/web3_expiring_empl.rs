// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3ExpiringEmpl {
    #[serde(rename = "items", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub items: Option<Option<Vec<models::Web3ExpiringEmplExpiredEmpl>>>,
    #[serde(rename = "total_amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total_amount: Option<Option<f64>>,
}

impl Web3ExpiringEmpl {
    pub fn new() -> Web3ExpiringEmpl {
        Web3ExpiringEmpl {
            items: None,
            total_amount: None,
        }
    }
}

