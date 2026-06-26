// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ModelCoinAmount {
    #[serde(rename = "free", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub free: Option<Option<f64>>,
    #[serde(rename = "paid", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub paid: Option<Option<f64>>,
    #[serde(rename = "total", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total: Option<Option<f64>>,
}

impl ModelCoinAmount {
    pub fn new() -> ModelCoinAmount {
        ModelCoinAmount {
            free: None,
            paid: None,
            total: None,
        }
    }
}

