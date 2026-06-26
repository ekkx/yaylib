// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct RecentSearchesResponse {
    #[serde(rename = "recent_searches", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub recent_searches: Option<Option<Vec<models::RecentSearch>>>,
}

impl RecentSearchesResponse {
    pub fn new() -> RecentSearchesResponse {
        RecentSearchesResponse {
            recent_searches: None,
        }
    }
}

