// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PalFreePoolStatus {
    #[serde(rename = "cooldown", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub cooldown: Option<Option<models::Cooldown>>,
    #[serde(rename = "pals_available", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pals_available: Option<Option<bool>>,
    #[serde(rename = "qualifications", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub qualifications: Option<Option<Vec<models::Qualification>>>,
    #[serde(rename = "qualified", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub qualified: Option<Option<bool>>,
}

impl PalFreePoolStatus {
    pub fn new() -> PalFreePoolStatus {
        PalFreePoolStatus {
            cooldown: None,
            pals_available: None,
            qualifications: None,
            qualified: None,
        }
    }
}

