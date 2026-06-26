// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct BgmsResponse {
    #[serde(rename = "bgm", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub bgm: Option<Option<Vec<models::Bgm>>>,
}

impl BgmsResponse {
    pub fn new() -> BgmsResponse {
        BgmsResponse {
            bgm: None,
        }
    }
}

