// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ProductQuota {
    #[serde(rename = "bought", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub bought: Option<Option<f64>>,
    #[serde(rename = "limit", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub limit: Option<Option<f64>>,
}

impl ProductQuota {
    pub fn new() -> ProductQuota {
        ProductQuota {
            bought: None,
            limit: None,
        }
    }
}

