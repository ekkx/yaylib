// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct HiddenRecommendedPost {
    #[serde(rename = "post", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub post: Option<Option<models::ModelPost>>,
}

impl HiddenRecommendedPost {
    pub fn new() -> HiddenRecommendedPost {
        HiddenRecommendedPost {
            post: None,
        }
    }
}

