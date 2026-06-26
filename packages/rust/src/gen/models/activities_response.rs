// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ActivitiesResponse {
    #[serde(rename = "activities", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub activities: Option<Option<Vec<models::Activity>>>,
    #[serde(rename = "last_timestamp", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub last_timestamp: Option<Option<i64>>,
}

impl ActivitiesResponse {
    pub fn new() -> ActivitiesResponse {
        ActivitiesResponse {
            activities: None,
            last_timestamp: None,
        }
    }
}

