// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PalLevelUpResponse {
    #[serde(rename = "result", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub result: Option<Option<models::LevelUpDetails>>,
}

impl PalLevelUpResponse {
    pub fn new() -> PalLevelUpResponse {
        PalLevelUpResponse {
            result: None,
        }
    }
}

