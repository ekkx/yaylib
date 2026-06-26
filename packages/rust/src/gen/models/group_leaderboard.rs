// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GroupLeaderboard {
    #[serde(rename = "group", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group: Option<Option<models::Group>>,
    #[serde(rename = "ranking", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub ranking: Option<Option<i32>>,
    #[serde(rename = "total_scores", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total_scores: Option<Option<i32>>,
}

impl GroupLeaderboard {
    pub fn new() -> GroupLeaderboard {
        GroupLeaderboard {
            group: None,
            ranking: None,
            total_scores: None,
        }
    }
}

