// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GiftReward {
    #[serde(rename = "empl_reward", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub empl_reward: Option<Option<i64>>,
    #[serde(rename = "gifts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gifts: Option<Option<Vec<models::GiftRewardGift>>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i32>>,
}

impl GiftReward {
    pub fn new() -> GiftReward {
        GiftReward {
            empl_reward: None,
            gifts: None,
            id: None,
        }
    }
}

