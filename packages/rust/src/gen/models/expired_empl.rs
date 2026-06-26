// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ExpiredEmpl {
    #[serde(rename = "expiring_amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub expiring_amount: Option<Option<String>>,
    #[serde(rename = "expiring_date", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub expiring_date: Option<Option<i64>>,
    #[serde(rename = "soon", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub soon: Option<Option<bool>>,
}

impl ExpiredEmpl {
    pub fn new() -> ExpiredEmpl {
        ExpiredEmpl {
            expiring_amount: None,
            expiring_date: None,
            soon: None,
        }
    }
}

