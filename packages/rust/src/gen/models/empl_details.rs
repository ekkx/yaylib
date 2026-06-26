// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct EmplDetails {
    #[serde(rename = "regular", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub regular: Option<Option<f64>>,
    #[serde(rename = "rewarded", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub rewarded: Option<Option<f64>>,
}

impl EmplDetails {
    pub fn new() -> EmplDetails {
        EmplDetails {
            regular: None,
            rewarded: None,
        }
    }
}

