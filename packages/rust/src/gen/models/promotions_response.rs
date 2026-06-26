// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PromotionsResponse {
    #[serde(rename = "promotions", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub promotions: Option<Option<Vec<models::Promotion>>>,
}

impl PromotionsResponse {
    pub fn new() -> PromotionsResponse {
        PromotionsResponse {
            promotions: None,
        }
    }
}

