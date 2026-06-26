// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct AppReviewStatusResponse {
    #[serde(rename = "is_app_reviewed", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_app_reviewed: Option<Option<bool>>,
}

impl AppReviewStatusResponse {
    pub fn new() -> AppReviewStatusResponse {
        AppReviewStatusResponse {
            is_app_reviewed: None,
        }
    }
}

