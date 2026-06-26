// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ModelInAppPurchaseProduct {
    #[serde(rename = "amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub amount: Option<Option<f64>>,
    #[serde(rename = "bonus", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub bonus: Option<Option<f64>>,
    #[serde(rename = "currency_code", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub currency_code: Option<Option<String>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<String>>,
    #[serde(rename = "price", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub price: Option<Option<f64>>,
    #[serde(rename = "purchasable", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub purchasable: Option<Option<bool>>,
}

impl ModelInAppPurchaseProduct {
    pub fn new() -> ModelInAppPurchaseProduct {
        ModelInAppPurchaseProduct {
            amount: None,
            bonus: None,
            currency_code: None,
            id: None,
            price: None,
            purchasable: None,
        }
    }
}

