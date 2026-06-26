// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct EmplFee {
    #[serde(rename = "gas_fee", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gas_fee: Option<Option<String>>,
}

impl EmplFee {
    pub fn new() -> EmplFee {
        EmplFee {
            gas_fee: None,
        }
    }
}

