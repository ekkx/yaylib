// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Breakdown {
    #[serde(rename = "activity", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub activity: Option<Option<f64>>,
    #[serde(rename = "original_points", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub original_points: Option<Option<i32>>,
    #[serde(rename = "special_mission", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub special_mission: Option<Option<f64>>,
    #[serde(rename = "vip_multiplier", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub vip_multiplier: Option<Option<f64>>,
}

impl Breakdown {
    pub fn new() -> Breakdown {
        Breakdown {
            activity: None,
            original_points: None,
            special_mission: None,
            vip_multiplier: None,
        }
    }
}

