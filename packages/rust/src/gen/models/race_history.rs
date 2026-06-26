// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct RaceHistory {
    #[serde(rename = "animation_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub animation_url: Option<Option<String>>,
    #[serde(rename = "league", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub league: Option<Option<models::PalRaceLeagueDto>>,
    #[serde(rename = "league_name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub league_name: Option<Option<String>>,
    #[serde(rename = "pals", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pals: Option<Option<Vec<models::Pal>>>,
    #[serde(rename = "race_date", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub race_date: Option<Option<i64>>,
    #[serde(rename = "race_entry_fee", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub race_entry_fee: Option<Option<i32>>,
    #[serde(rename = "race_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub race_id: Option<Option<String>>,
    #[serde(rename = "race_reward", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub race_reward: Option<Option<models::RaceReward>>,
    #[serde(rename = "race_start_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub race_start_at: Option<Option<i64>>,
    #[serde(rename = "race_status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub race_status: Option<Option<String>>,
    #[serde(rename = "results", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub results: Option<Option<Vec<models::PalRank>>>,
    #[serde(rename = "reward", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub reward: Option<Option<models::Reward>>,
}

impl RaceHistory {
    pub fn new() -> RaceHistory {
        RaceHistory {
            animation_url: None,
            league: None,
            league_name: None,
            pals: None,
            race_date: None,
            race_entry_fee: None,
            race_id: None,
            race_reward: None,
            race_start_at: None,
            race_status: None,
            results: None,
            reward: None,
        }
    }
}

