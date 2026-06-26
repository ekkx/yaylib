// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ContactStatusResponse {
    #[serde(rename = "contacts", skip_serializing_if = "Option::is_none")]
    pub contacts: Option<std::collections::HashMap<String, models::ContactStatus>>,
}

impl ContactStatusResponse {
    pub fn new() -> ContactStatusResponse {
        ContactStatusResponse {
            contacts: None,
        }
    }
}

