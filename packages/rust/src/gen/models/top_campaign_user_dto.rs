// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct TopCampaignUserDto {
    #[serde(rename = "points", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub points: Option<Option<i32>>,
    #[serde(rename = "ranking", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub ranking: Option<Option<i32>>,
    #[serde(rename = "user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user: Option<Option<models::UserDto>>,
}

impl TopCampaignUserDto {
    pub fn new() -> TopCampaignUserDto {
        TopCampaignUserDto {
            points: None,
            ranking: None,
            user: None,
        }
    }
}

