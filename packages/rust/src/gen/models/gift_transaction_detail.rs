// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GiftTransactionDetail {
    #[serde(rename = "gifts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gifts: Option<Option<Vec<models::GiftTransaction>>>,
}

impl GiftTransactionDetail {
    pub fn new() -> GiftTransactionDetail {
        GiftTransactionDetail {
            gifts: None,
        }
    }
}

