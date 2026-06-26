// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct LevelUpDetails {
    #[serde(rename = "available", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub available: Option<Option<bool>>,
    #[serde(rename = "cost", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub cost: Option<Option<String>>,
    #[serde(rename = "empl_earn_limit", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub empl_earn_limit: Option<Option<String>>,
    #[serde(rename = "level", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub level: Option<Option<i32>>,
    #[serde(rename = "pal", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal: Option<Option<models::LevelUpDetailsPal>>,
    #[serde(rename = "stat_changes", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub stat_changes: Option<Option<models::StateChanges>>,
}

impl LevelUpDetails {
    pub fn new() -> LevelUpDetails {
        LevelUpDetails {
            available: None,
            cost: None,
            empl_earn_limit: None,
            level: None,
            pal: None,
            stat_changes: None,
        }
    }
}

