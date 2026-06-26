// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GroupOverallLeaderboardResponse {
    #[serde(rename = "group_leaderboard", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group_leaderboard: Option<Option<Vec<models::GroupLeaderboard>>>,
    #[serde(rename = "updated_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub updated_at: Option<Option<i64>>,
}

impl GroupOverallLeaderboardResponse {
    pub fn new() -> GroupOverallLeaderboardResponse {
        GroupOverallLeaderboardResponse {
            group_leaderboard: None,
            updated_at: None,
        }
    }
}

