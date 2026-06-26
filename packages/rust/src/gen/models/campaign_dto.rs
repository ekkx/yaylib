// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CampaignDto {
    #[serde(rename = "banner_image", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub banner_image: Option<Option<String>>,
    #[serde(rename = "cover_image", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub cover_image: Option<Option<String>>,
    #[serde(rename = "description", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub description: Option<Option<String>>,
    #[serde(rename = "description_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub description_url: Option<Option<String>>,
    #[serde(rename = "end_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub end_at: Option<Option<i64>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "multiplier_detail_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub multiplier_detail_url: Option<Option<String>>,
    #[serde(rename = "name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub name: Option<Option<String>>,
    #[serde(rename = "ranking", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub ranking: Option<Option<i32>>,
    #[serde(rename = "start_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub start_at: Option<Option<i64>>,
    #[serde(rename = "total_points", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total_points: Option<Option<i32>>,
    #[serde(rename = "user_campaign", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user_campaign: Option<Option<models::UserCampaign>>,
}

impl CampaignDto {
    pub fn new() -> CampaignDto {
        CampaignDto {
            banner_image: None,
            cover_image: None,
            description: None,
            description_url: None,
            end_at: None,
            id: None,
            multiplier_detail_url: None,
            name: None,
            ranking: None,
            start_at: None,
            total_points: None,
            user_campaign: None,
        }
    }
}

