// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GroupCommunityCampaignResponse {
    #[serde(rename = "eligible_for_apply", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub eligible_for_apply: Option<Option<bool>>,
    #[serde(rename = "viewable_group_leaderboard", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub viewable_group_leaderboard: Option<Option<bool>>,
}

impl GroupCommunityCampaignResponse {
    pub fn new() -> GroupCommunityCampaignResponse {
        GroupCommunityCampaignResponse {
            eligible_for_apply: None,
            viewable_group_leaderboard: None,
        }
    }
}

