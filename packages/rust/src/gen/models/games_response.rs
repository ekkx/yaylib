// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GamesResponse {
    #[serde(rename = "from_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub from_id: Option<Option<i64>>,
    #[serde(rename = "games", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub games: Option<Option<Vec<models::RealmGame>>>,
}

impl GamesResponse {
    pub fn new() -> GamesResponse {
        GamesResponse {
            from_id: None,
            games: None,
        }
    }
}

