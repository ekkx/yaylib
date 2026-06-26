// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CampaignVipEmplBonusDto {
    #[serde(rename = "claimed", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub claimed: Option<Option<bool>>,
    #[serde(rename = "details", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub details: Option<Option<models::DetailsDto>>,
    #[serde(rename = "end_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub end_at: Option<Option<i64>>,
    #[serde(rename = "key", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub key: Option<Option<String>>,
}

impl CampaignVipEmplBonusDto {
    pub fn new() -> CampaignVipEmplBonusDto {
        CampaignVipEmplBonusDto {
            claimed: None,
            details: None,
            end_at: None,
            key: None,
        }
    }
}

