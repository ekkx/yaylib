// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GroupInCircleRanking {
    #[serde(rename = "gifts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gifts: Option<Option<Vec<String>>>,
    #[serde(rename = "user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user: Option<Option<models::User>>,
}

impl GroupInCircleRanking {
    pub fn new() -> GroupInCircleRanking {
        GroupInCircleRanking {
            gifts: None,
            user: None,
        }
    }
}

