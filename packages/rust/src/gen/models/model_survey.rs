// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ModelSurvey {
    #[serde(rename = "choices", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub choices: Option<Option<Vec<models::SurveyChoice>>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "is_voted", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_voted: Option<Option<bool>>,
    #[serde(rename = "votes_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub votes_count: Option<Option<i32>>,
}

impl ModelSurvey {
    pub fn new() -> ModelSurvey {
        ModelSurvey {
            choices: None,
            id: None,
            is_voted: None,
            votes_count: None,
        }
    }
}

