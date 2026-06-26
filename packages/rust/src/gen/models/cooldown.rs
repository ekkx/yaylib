// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Cooldown {
    #[serde(rename = "active", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub active: Option<Option<bool>>,
    #[serde(rename = "retry_after", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub retry_after: Option<Option<i64>>,
}

impl Cooldown {
    pub fn new() -> Cooldown {
        Cooldown {
            active: None,
            retry_after: None,
        }
    }
}

