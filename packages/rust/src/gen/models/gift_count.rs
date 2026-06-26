// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GiftCount {
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "quantity", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<Option<i32>>,
}

impl GiftCount {
    pub fn new() -> GiftCount {
        GiftCount {
            id: None,
            quantity: None,
        }
    }
}

