// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct FeaturesDto {
    #[serde(rename = "pal", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal: Option<Option<models::FeaturesPalDto>>,
}

impl FeaturesDto {
    pub fn new() -> FeaturesDto {
        FeaturesDto {
            pal: None,
        }
    }
}

