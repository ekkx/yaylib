// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CampaignsResponse {
    #[serde(rename = "campaigns", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub campaigns: Option<Option<Vec<models::CampaignDto>>>,
}

impl CampaignsResponse {
    pub fn new() -> CampaignsResponse {
        CampaignsResponse {
            campaigns: None,
        }
    }
}

