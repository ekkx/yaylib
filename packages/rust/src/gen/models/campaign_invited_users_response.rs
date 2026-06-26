// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CampaignInvitedUsersResponse {
    #[serde(rename = "invited_users", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub invited_users: Option<Option<Vec<models::InvitedUserDto>>>,
    #[serde(rename = "total_points", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total_points: Option<Option<i32>>,
}

impl CampaignInvitedUsersResponse {
    pub fn new() -> CampaignInvitedUsersResponse {
        CampaignInvitedUsersResponse {
            invited_users: None,
            total_points: None,
        }
    }
}

