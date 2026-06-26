// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct UserRank {
    #[serde(rename = "ranking", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub ranking: Option<Option<i32>>,
    #[serde(rename = "top_gifts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub top_gifts: Option<Option<Vec<String>>>,
    #[serde(rename = "total_scores", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total_scores: Option<Option<i32>>,
    #[serde(rename = "updated_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub updated_at: Option<Option<i64>>,
    #[serde(rename = "user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user: Option<Option<models::RealmUser>>,
}

impl UserRank {
    pub fn new() -> UserRank {
        UserRank {
            ranking: None,
            top_gifts: None,
            total_scores: None,
            updated_at: None,
            user: None,
        }
    }
}

