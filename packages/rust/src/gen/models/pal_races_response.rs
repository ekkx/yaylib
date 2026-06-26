// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PalRacesResponse {
    #[serde(rename = "races", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub races: Option<Option<Vec<models::RaceHistory>>>,
}

impl PalRacesResponse {
    pub fn new() -> PalRacesResponse {
        PalRacesResponse {
            races: None,
        }
    }
}

