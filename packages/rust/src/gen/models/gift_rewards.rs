// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GiftRewards {
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i32>>,
    #[serde(rename = "requirements", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub requirements: Option<Option<Vec<models::Requirement>>>,
    #[serde(rename = "rewarded_amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub rewarded_amount: Option<Option<i64>>,
}

impl GiftRewards {
    pub fn new() -> GiftRewards {
        GiftRewards {
            id: None,
            requirements: None,
            rewarded_amount: None,
        }
    }
}

