// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct UserInterestsResponse {
    #[serde(rename = "interests", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub interests: Option<Option<Vec<models::Interest>>>,
}

impl UserInterestsResponse {
    pub fn new() -> UserInterestsResponse {
        UserInterestsResponse {
            interests: None,
        }
    }
}

