// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct RealmGiftingAbility {
    #[serde(rename = "can_receive", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub can_receive: Option<Option<bool>>,
    #[serde(rename = "can_send", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub can_send: Option<Option<bool>>,
    #[serde(rename = "enabled", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub enabled: Option<Option<bool>>,
    #[serde(rename = "user_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user_id: Option<Option<i64>>,
}

impl RealmGiftingAbility {
    pub fn new() -> RealmGiftingAbility {
        RealmGiftingAbility {
            can_receive: None,
            can_send: None,
            enabled: None,
            user_id: None,
        }
    }
}

