// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct TrustScore {
    #[serde(rename = "calculated", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub calculated: Option<Option<bool>>,
    #[serde(rename = "grade", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub grade: Option<Option<models::ModelUserRank>>,
    #[serde(rename = "score", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub score: Option<Option<f64>>,
}

impl TrustScore {
    pub fn new() -> TrustScore {
        TrustScore {
            calculated: None,
            grade: None,
            score: None,
        }
    }
}

