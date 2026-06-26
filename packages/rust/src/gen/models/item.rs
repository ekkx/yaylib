// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Item {
    #[serde(rename = "gift", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gift: Option<Option<models::RealmGift>>,
    #[serde(rename = "quantity", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<Option<i32>>,
}

impl Item {
    pub fn new() -> Item {
        Item {
            gift: None,
            quantity: None,
        }
    }
}

