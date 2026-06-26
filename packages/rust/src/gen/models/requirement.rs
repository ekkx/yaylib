// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Requirement {
    #[serde(rename = "gift", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gift: Option<Option<models::RealmGift>>,
    #[serde(rename = "received_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub received_count: Option<Option<i32>>,
    #[serde(rename = "requires_amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub requires_amount: Option<Option<i32>>,
}

impl Requirement {
    pub fn new() -> Requirement {
        Requirement {
            gift: None,
            received_count: None,
            requires_amount: None,
        }
    }
}

