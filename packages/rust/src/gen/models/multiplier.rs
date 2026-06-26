// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Multiplier {
    #[serde(rename = "activity", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub activity: Option<Option<String>>,
    #[serde(rename = "base", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub base: Option<Option<String>>,
    #[serde(rename = "date", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub date: Option<Option<String>>,
    #[serde(rename = "mission", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub mission: Option<Option<String>>,
    #[serde(rename = "multiplier", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub multiplier: Option<Option<String>>,
    #[serde(rename = "vip", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub vip: Option<Option<String>>,
}

impl Multiplier {
    pub fn new() -> Multiplier {
        Multiplier {
            activity: None,
            base: None,
            date: None,
            mission: None,
            multiplier: None,
            vip: None,
        }
    }
}

