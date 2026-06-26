// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GroupOverallLeaderboard {
    #[serde(rename = "group_rankings", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group_rankings: Option<Option<Vec<models::GroupRanking>>>,
    #[serde(rename = "updated_at_millis", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub updated_at_millis: Option<Option<i64>>,
}

impl GroupOverallLeaderboard {
    pub fn new() -> GroupOverallLeaderboard {
        GroupOverallLeaderboard {
            group_rankings: None,
            updated_at_millis: None,
        }
    }
}

