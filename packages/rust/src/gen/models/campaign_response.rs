// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CampaignResponse {
    #[serde(rename = "campaign", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub campaign: Option<Option<models::CampaignDto>>,
}

impl CampaignResponse {
    pub fn new() -> CampaignResponse {
        CampaignResponse {
            campaign: None,
        }
    }
}

