// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CampaignMissionsResponse {
    #[serde(rename = "daily_missions", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub daily_missions: Option<Option<Vec<models::MissionDto>>>,
    #[serde(rename = "multiplier_missions", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub multiplier_missions: Option<Option<Vec<models::MissionDto>>>,
    #[serde(rename = "normal_missions", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub normal_missions: Option<Option<Vec<models::MissionDto>>>,
    #[serde(rename = "unlimited_missions", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub unlimited_missions: Option<Option<Vec<models::MissionDto>>>,
}

impl CampaignMissionsResponse {
    pub fn new() -> CampaignMissionsResponse {
        CampaignMissionsResponse {
            daily_missions: None,
            multiplier_missions: None,
            normal_missions: None,
            unlimited_missions: None,
        }
    }
}

