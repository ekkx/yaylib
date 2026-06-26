// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3WalletUserStatusDetails {
    #[serde(rename = "activity_rank", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub activity_rank: Option<Option<models::ModelUserRank>>,
    #[serde(rename = "is_rank_calculated", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_rank_calculated: Option<Option<bool>>,
    #[serde(rename = "is_score_calculated", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_score_calculated: Option<Option<bool>>,
    #[serde(rename = "score", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub score: Option<Option<i32>>,
    #[serde(rename = "score_max_format", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub score_max_format: Option<Option<String>>,
    #[serde(rename = "user_rank", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user_rank: Option<Option<models::ModelUserRank>>,
}

impl Web3WalletUserStatusDetails {
    pub fn new() -> Web3WalletUserStatusDetails {
        Web3WalletUserStatusDetails {
            activity_rank: None,
            is_rank_calculated: None,
            is_score_calculated: None,
            score: None,
            score_max_format: None,
            user_rank: None,
        }
    }
}

