// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct InAppPurchaseProduct {
    #[serde(rename = "amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub amount: Option<Option<f64>>,
    #[serde(rename = "bonus", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub bonus: Option<Option<f64>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<String>>,
    #[serde(rename = "purchasable", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub purchasable: Option<Option<bool>>,
}

impl InAppPurchaseProduct {
    pub fn new() -> InAppPurchaseProduct {
        InAppPurchaseProduct {
            amount: None,
            bonus: None,
            id: None,
            purchasable: None,
        }
    }
}

