// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Result {
    #[serde(rename = "amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub amount: Option<Option<String>>,
    #[serde(rename = "currency", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub currency: Option<Option<String>>,
    #[serde(rename = "details", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub details: Option<Option<models::Details>>,
    #[serde(rename = "empl_transaction", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub empl_transaction: Option<Option<models::EmplTransaction>>,
    #[serde(rename = "final_status_updated_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub final_status_updated_at: Option<Option<i64>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub status: Option<Option<String>>,
    #[serde(rename = "transaction_code", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub transaction_code: Option<Option<String>>,
    #[serde(rename = "wallet_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub wallet_address: Option<Option<String>>,
}

impl Result {
    pub fn new() -> Result {
        Result {
            amount: None,
            currency: None,
            details: None,
            empl_transaction: None,
            final_status_updated_at: None,
            id: None,
            status: None,
            transaction_code: None,
            wallet_address: None,
        }
    }
}

