// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct VoteSurveyResponse {
    #[serde(rename = "survey", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub survey: Option<Option<models::Survey>>,
}

impl VoteSurveyResponse {
    pub fn new() -> VoteSurveyResponse {
        VoteSurveyResponse {
            survey: None,
        }
    }
}

