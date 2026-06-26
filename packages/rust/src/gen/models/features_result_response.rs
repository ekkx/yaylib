// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct FeaturesResultResponse {
    #[serde(rename = "features", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub features: Option<Option<models::FeaturesDto>>,
}

impl FeaturesResultResponse {
    pub fn new() -> FeaturesResultResponse {
        FeaturesResultResponse {
            features: None,
        }
    }
}

