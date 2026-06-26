// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PalGachaListResponse {
    #[serde(rename = "result", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub result: Option<Option<Vec<models::PalGacha>>>,
}

impl PalGachaListResponse {
    pub fn new() -> PalGachaListResponse {
        PalGachaListResponse {
            result: None,
        }
    }
}

