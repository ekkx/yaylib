// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PalRaceDetailsResponse {
    #[serde(rename = "race", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub race: Option<Option<models::RaceHistory>>,
}

impl PalRaceDetailsResponse {
    pub fn new() -> PalRaceDetailsResponse {
        PalRaceDetailsResponse {
            race: None,
        }
    }
}

