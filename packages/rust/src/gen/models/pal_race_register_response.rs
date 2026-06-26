// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PalRaceRegisterResponse {
    #[serde(rename = "race_history", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub race_history: Option<Option<models::RaceHistory>>,
}

impl PalRaceRegisterResponse {
    pub fn new() -> PalRaceRegisterResponse {
        PalRaceRegisterResponse {
            race_history: None,
        }
    }
}

