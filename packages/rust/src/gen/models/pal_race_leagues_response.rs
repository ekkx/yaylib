// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PalRaceLeaguesResponse {
    #[serde(rename = "leagues", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub leagues: Option<Option<Vec<models::PalRaceLeagueDto>>>,
}

impl PalRaceLeaguesResponse {
    pub fn new() -> PalRaceLeaguesResponse {
        PalRaceLeaguesResponse {
            leagues: None,
        }
    }
}

