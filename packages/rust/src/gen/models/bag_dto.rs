// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct BagDto {
    #[serde(rename = "eggs_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub eggs_count: Option<Option<i32>>,
    #[serde(rename = "empl", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub empl: Option<Option<models::EmplDto>>,
    #[serde(rename = "expiring_empl", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub expiring_empl: Option<Option<String>>,
    #[serde(rename = "pals_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pals_count: Option<Option<i32>>,
    #[serde(rename = "race_pal_available", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub race_pal_available: Option<Option<i32>>,
}

impl BagDto {
    pub fn new() -> BagDto {
        BagDto {
            eggs_count: None,
            empl: None,
            expiring_empl: None,
            pals_count: None,
            race_pal_available: None,
        }
    }
}

