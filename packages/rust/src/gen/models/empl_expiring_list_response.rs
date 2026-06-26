// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct EmplExpiringListResponse {
    #[serde(rename = "regular_expiring", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub regular_expiring: Option<Option<Vec<models::EmplExpiringResponse>>>,
    #[serde(rename = "rewarded_expiring", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub rewarded_expiring: Option<Option<Vec<models::EmplExpiringResponse>>>,
}

impl EmplExpiringListResponse {
    pub fn new() -> EmplExpiringListResponse {
        EmplExpiringListResponse {
            regular_expiring: None,
            rewarded_expiring: None,
        }
    }
}

