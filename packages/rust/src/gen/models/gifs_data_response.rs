// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GifsDataResponse {
    #[serde(rename = "gif_categories", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gif_categories: Option<Option<Vec<models::GifImageCategory>>>,
}

impl GifsDataResponse {
    pub fn new() -> GifsDataResponse {
        GifsDataResponse {
            gif_categories: None,
        }
    }
}

