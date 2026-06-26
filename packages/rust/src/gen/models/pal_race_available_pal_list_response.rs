// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PalRaceAvailablePalListResponse {
    #[serde(rename = "pals", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pals: Option<Option<Vec<models::PalRaceItemDto>>>,
}

impl PalRaceAvailablePalListResponse {
    pub fn new() -> PalRaceAvailablePalListResponse {
        PalRaceAvailablePalListResponse {
            pals: None,
        }
    }
}

