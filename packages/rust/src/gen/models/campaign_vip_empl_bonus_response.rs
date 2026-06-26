// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CampaignVipEmplBonusResponse {
    #[serde(rename = "app_campaign", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub app_campaign: Option<Option<models::CampaignVipEmplBonusDto>>,
}

impl CampaignVipEmplBonusResponse {
    pub fn new() -> CampaignVipEmplBonusResponse {
        CampaignVipEmplBonusResponse {
            app_campaign: None,
        }
    }
}

