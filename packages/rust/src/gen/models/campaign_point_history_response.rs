// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CampaignPointHistoryResponse {
    #[serde(rename = "point_histories", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub point_histories: Option<Option<Vec<models::CampaignPointHistoryDto>>>,
}

impl CampaignPointHistoryResponse {
    pub fn new() -> CampaignPointHistoryResponse {
        CampaignPointHistoryResponse {
            point_histories: None,
        }
    }
}

