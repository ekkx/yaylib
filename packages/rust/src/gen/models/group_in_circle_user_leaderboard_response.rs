// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GroupInCircleUserLeaderboardResponse {
    #[serde(rename = "user_leaderboard", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user_leaderboard: Option<Option<Vec<models::UserRank>>>,
}

impl GroupInCircleUserLeaderboardResponse {
    pub fn new() -> GroupInCircleUserLeaderboardResponse {
        GroupInCircleUserLeaderboardResponse {
            user_leaderboard: None,
        }
    }
}

