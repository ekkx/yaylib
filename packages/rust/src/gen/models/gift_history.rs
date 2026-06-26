// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GiftHistory {
    #[serde(rename = "anonymous", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub anonymous: Option<Option<bool>>,
    #[serde(rename = "gift_item", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gift_item: Option<Option<models::GiftSlugItem>>,
    #[serde(rename = "gifts_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gifts_count: Option<Option<Vec<models::GiftCount>>>,
    #[serde(rename = "receiver", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub receiver: Option<Option<models::RealmUser>>,
    #[serde(rename = "sender", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub sender: Option<Option<models::RealmUser>>,
    #[serde(rename = "sent_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub sent_at: Option<Option<i64>>,
}

impl GiftHistory {
    pub fn new() -> GiftHistory {
        GiftHistory {
            anonymous: None,
            gift_item: None,
            gifts_count: None,
            receiver: None,
            sender: None,
            sent_at: None,
        }
    }
}

