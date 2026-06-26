// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GiftRewardGift {
    #[serde(rename = "icon", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub icon: Option<Option<String>>,
    #[serde(rename = "received_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub received_count: Option<Option<i32>>,
    #[serde(rename = "required_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub required_count: Option<Option<i32>>,
    #[serde(rename = "title", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title: Option<Option<String>>,
}

impl GiftRewardGift {
    pub fn new() -> GiftRewardGift {
        GiftRewardGift {
            icon: None,
            received_count: None,
            required_count: None,
            title: None,
        }
    }
}

