// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PalActivityLogs {
    #[serde(rename = "result", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub result: Option<Option<Vec<models::PalActivityLog>>>,
}

impl PalActivityLogs {
    pub fn new() -> PalActivityLogs {
        PalActivityLogs {
            result: None,
        }
    }
}

