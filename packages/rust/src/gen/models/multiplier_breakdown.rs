// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct MultiplierBreakdown {
    #[serde(rename = "activity", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub activity: Option<Option<f64>>,
    #[serde(rename = "special_mission", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub special_mission: Option<Option<f64>>,
    #[serde(rename = "vip_multiplier", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub vip_multiplier: Option<Option<f64>>,
}

impl MultiplierBreakdown {
    pub fn new() -> MultiplierBreakdown {
        MultiplierBreakdown {
            activity: None,
            special_mission: None,
            vip_multiplier: None,
        }
    }
}

