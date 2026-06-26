// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GiftCardOption {
    #[serde(rename = "empl", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub empl: Option<Option<f64>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<String>>,
    #[serde(rename = "value", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub value: Option<Option<f64>>,
}

impl GiftCardOption {
    pub fn new() -> GiftCardOption {
        GiftCardOption {
            empl: None,
            id: None,
            value: None,
        }
    }
}

