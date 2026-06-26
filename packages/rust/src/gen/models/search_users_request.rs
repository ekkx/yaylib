// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct SearchUsersRequest {
    #[serde(rename = "user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user: Option<Option<models::SearchCriteria>>,
}

impl SearchUsersRequest {
    pub fn new() -> SearchUsersRequest {
        SearchUsersRequest {
            user: None,
        }
    }
}

