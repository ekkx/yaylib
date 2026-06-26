// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GroupGiftHistory {
    #[serde(rename = "gifts_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gifts_count: Option<Option<Vec<models::GiftCount>>>,
    #[serde(rename = "received_date", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub received_date: Option<Option<i64>>,
    #[serde(rename = "user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user: Option<Option<models::RealmUser>>,
}

impl GroupGiftHistory {
    pub fn new() -> GroupGiftHistory {
        GroupGiftHistory {
            gifts_count: None,
            received_date: None,
            user: None,
        }
    }
}

