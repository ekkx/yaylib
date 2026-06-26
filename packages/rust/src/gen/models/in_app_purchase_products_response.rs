// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct InAppPurchaseProductsResponse {
    #[serde(rename = "iap_products", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub iap_products: Option<Option<Vec<models::InAppPurchaseProduct>>>,
    #[serde(rename = "quota", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub quota: Option<Option<models::ProductQuota>>,
}

impl InAppPurchaseProductsResponse {
    pub fn new() -> InAppPurchaseProductsResponse {
        InAppPurchaseProductsResponse {
            iap_products: None,
            quota: None,
        }
    }
}

