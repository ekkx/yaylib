// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct RaceReward {
    #[serde(rename = "first_prize", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub first_prize: Option<Option<models::Reward>>,
    #[serde(rename = "second_prize", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub second_prize: Option<Option<models::Reward>>,
    #[serde(rename = "third_prize", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub third_prize: Option<Option<models::Reward>>,
}

impl RaceReward {
    pub fn new() -> RaceReward {
        RaceReward {
            first_prize: None,
            second_prize: None,
            third_prize: None,
        }
    }
}

