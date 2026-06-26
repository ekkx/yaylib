// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct DetailsDto {
    #[serde(rename = "reward_amount_adult", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub reward_amount_adult: Option<Option<i32>>,
    #[serde(rename = "reward_amount_minor", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub reward_amount_minor: Option<Option<i32>>,
}

impl DetailsDto {
    pub fn new() -> DetailsDto {
        DetailsDto {
            reward_amount_adult: None,
            reward_amount_minor: None,
        }
    }
}

