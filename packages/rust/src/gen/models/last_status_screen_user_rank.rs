// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct LastStatusScreenUserRank {
    #[serde(rename = "user_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user_id: Option<Option<i64>>,
    #[serde(rename = "user_rank", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user_rank: Option<Option<String>>,
}

impl LastStatusScreenUserRank {
    pub fn new() -> LastStatusScreenUserRank {
        LastStatusScreenUserRank {
            user_id: None,
            user_rank: None,
        }
    }
}

