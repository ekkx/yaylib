// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Survey {
    #[serde(rename = "choices", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub choices: Option<Option<Vec<models::Choice>>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "voted", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub voted: Option<Option<bool>>,
    #[serde(rename = "votes_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub votes_count: Option<Option<i32>>,
}

impl Survey {
    pub fn new() -> Survey {
        Survey {
            choices: None,
            id: None,
            voted: None,
            votes_count: None,
        }
    }
}

