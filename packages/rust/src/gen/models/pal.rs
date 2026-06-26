// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Pal {
    #[serde(rename = "current_level", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub current_level: Option<Option<i32>>,
    #[serde(rename = "fatigue_level", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub fatigue_level: Option<Option<String>>,
    #[serde(rename = "nickname", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub nickname: Option<Option<String>>,
    #[serde(rename = "overall_strength", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub overall_strength: Option<Option<i32>>,
    #[serde(rename = "profile_image", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub profile_image: Option<Option<models::ProfileImage>>,
    #[serde(rename = "rank", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub rank: Option<Option<String>>,
    #[serde(rename = "token_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_id: Option<Option<String>>,
    #[serde(rename = "tribe", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub tribe: Option<Option<String>>,
    #[serde(rename = "type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub r#type: Option<Option<String>>,
    #[serde(rename = "win_rate", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub win_rate: Option<Option<String>>,
    #[serde(rename = "winning_streak", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub winning_streak: Option<Option<i32>>,
}

impl Pal {
    pub fn new() -> Pal {
        Pal {
            current_level: None,
            fatigue_level: None,
            nickname: None,
            overall_strength: None,
            profile_image: None,
            rank: None,
            token_id: None,
            tribe: None,
            r#type: None,
            win_rate: None,
            winning_streak: None,
        }
    }
}

