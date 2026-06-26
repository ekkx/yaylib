// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CampaignPointHistoryDto {
    #[serde(rename = "breakdown", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub breakdown: Option<Option<models::Breakdown>>,
    #[serde(rename = "created_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<Option<i64>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "mission", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub mission: Option<Option<models::Mission>>,
    #[serde(rename = "multiplier", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub multiplier: Option<Option<f64>>,
    #[serde(rename = "point_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub point_type: Option<Option<String>>,
    #[serde(rename = "total_points", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total_points: Option<Option<i32>>,
}

impl CampaignPointHistoryDto {
    pub fn new() -> CampaignPointHistoryDto {
        CampaignPointHistoryDto {
            breakdown: None,
            created_at: None,
            id: None,
            mission: None,
            multiplier: None,
            point_type: None,
            total_points: None,
        }
    }
}

