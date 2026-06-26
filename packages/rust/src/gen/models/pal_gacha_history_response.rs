// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PalGachaHistoryResponse {
    #[serde(rename = "cursor", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub cursor: Option<Option<String>>,
    #[serde(rename = "result", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub result: Option<Option<Vec<models::PalGachaHistoryResponseItem>>>,
}

impl PalGachaHistoryResponse {
    pub fn new() -> PalGachaHistoryResponse {
        PalGachaHistoryResponse {
            cursor: None,
            result: None,
        }
    }
}

