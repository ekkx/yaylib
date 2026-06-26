// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct UserCampaign {
    #[serde(rename = "multiplier", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub multiplier: Option<Option<f64>>,
    #[serde(rename = "multiplier_breakdown", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub multiplier_breakdown: Option<Option<models::MultiplierBreakdown>>,
    #[serde(rename = "points", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub points: Option<Option<i32>>,
    #[serde(rename = "ranking", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub ranking: Option<Option<i32>>,
}

impl UserCampaign {
    pub fn new() -> UserCampaign {
        UserCampaign {
            multiplier: None,
            multiplier_breakdown: None,
            points: None,
            ranking: None,
        }
    }
}

