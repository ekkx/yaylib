// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ReviewsResponse {
    #[serde(rename = "pinned_reviews", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pinned_reviews: Option<Option<Vec<models::Review>>>,
    #[serde(rename = "reviews", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub reviews: Option<Option<Vec<models::Review>>>,
}

impl ReviewsResponse {
    pub fn new() -> ReviewsResponse {
        ReviewsResponse {
            pinned_reviews: None,
            reviews: None,
        }
    }
}

