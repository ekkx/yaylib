// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ApiN {
    #[serde(rename = "d", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub d: Option<Option<String>>,
    #[serde(rename = "f", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub f: Option<Option<String>>,
    #[serde(rename = "h", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub h: Option<Option<i64>>,
    #[serde(rename = "i", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub i: Option<Option<i64>>,
}

impl ApiN {
    pub fn new() -> ApiN {
        ApiN {
            d: None,
            f: None,
            h: None,
            i: None,
        }
    }
}

