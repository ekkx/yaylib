// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PostLikersResponse {
    #[serde(rename = "last_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub last_id: Option<Option<i64>>,
    #[serde(rename = "users", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub users: Option<Option<Vec<models::RealmUser>>>,
}

impl PostLikersResponse {
    pub fn new() -> PostLikersResponse {
        PostLikersResponse {
            last_id: None,
            users: None,
        }
    }
}

