// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct UserReputations {
    #[serde(rename = "activity_score", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub activity_score: Option<Option<models::ActivityScore>>,
    #[serde(rename = "trust_score", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub trust_score: Option<Option<models::TrustScore>>,
}

impl UserReputations {
    pub fn new() -> UserReputations {
        UserReputations {
            activity_score: None,
            trust_score: None,
        }
    }
}

