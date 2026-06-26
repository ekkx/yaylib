// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GiftExchangeHistory {
    #[serde(rename = "created_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<Option<i64>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i32>>,
    #[serde(rename = "items", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub items: Option<Option<Vec<models::Item>>>,
    #[serde(rename = "quantity", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<Option<i32>>,
    #[serde(rename = "rewarded_amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub rewarded_amount: Option<Option<i64>>,
    #[serde(rename = "status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub status: Option<Option<String>>,
    #[serde(rename = "uuid", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub uuid: Option<Option<String>>,
}

impl GiftExchangeHistory {
    pub fn new() -> GiftExchangeHistory {
        GiftExchangeHistory {
            created_at: None,
            id: None,
            items: None,
            quantity: None,
            rewarded_amount: None,
            status: None,
            uuid: None,
        }
    }
}

