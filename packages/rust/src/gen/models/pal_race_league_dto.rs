// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PalRaceLeagueDto {
    #[serde(rename = "accessible", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub accessible: Option<Option<bool>>,
    #[serde(rename = "entry_fee", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub entry_fee: Option<Option<i32>>,
    #[serde(rename = "game_end_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub game_end_at: Option<Option<i64>>,
    #[serde(rename = "name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub name: Option<Option<String>>,
    #[serde(rename = "price_reward", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub price_reward: Option<Option<i32>>,
    #[serde(rename = "registration_close_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub registration_close_at: Option<Option<i64>>,
}

impl PalRaceLeagueDto {
    pub fn new() -> PalRaceLeagueDto {
        PalRaceLeagueDto {
            accessible: None,
            entry_fee: None,
            game_end_at: None,
            name: None,
            price_reward: None,
            registration_close_at: None,
        }
    }
}

