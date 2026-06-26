// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GiftRewardResponse {
    #[serde(rename = "gift_rewards", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gift_rewards: Option<Option<Vec<models::GiftRewards>>>,
}

impl GiftRewardResponse {
    pub fn new() -> GiftRewardResponse {
        GiftRewardResponse {
            gift_rewards: None,
        }
    }
}

