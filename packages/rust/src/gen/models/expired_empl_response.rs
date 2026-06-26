// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ExpiredEmplResponse {
    #[serde(rename = "expirations", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub expirations: Option<Option<Vec<models::ExpiredEmpl>>>,
    #[serde(rename = "total_amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total_amount: Option<Option<String>>,
}

impl ExpiredEmplResponse {
    pub fn new() -> ExpiredEmplResponse {
        ExpiredEmplResponse {
            expirations: None,
            total_amount: None,
        }
    }
}

