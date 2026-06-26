// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PalGachaAvailabilityResponse {
    #[serde(rename = "state", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub state: Option<Option<String>>,
}

impl PalGachaAvailabilityResponse {
    pub fn new() -> PalGachaAvailabilityResponse {
        PalGachaAvailabilityResponse {
            state: None,
        }
    }
}

