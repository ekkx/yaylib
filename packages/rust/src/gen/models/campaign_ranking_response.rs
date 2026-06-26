// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CampaignRankingResponse {
    #[serde(rename = "my_ranking", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub my_ranking: Option<Option<i32>>,
    #[serde(rename = "next_cursor", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub next_cursor: Option<Option<String>>,
    #[serde(rename = "top_campaign_users", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub top_campaign_users: Option<Option<Vec<models::TopCampaignUserDto>>>,
    #[serde(rename = "updated_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub updated_at: Option<Option<i64>>,
}

impl CampaignRankingResponse {
    pub fn new() -> CampaignRankingResponse {
        CampaignRankingResponse {
            my_ranking: None,
            next_cursor: None,
            top_campaign_users: None,
            updated_at: None,
        }
    }
}

