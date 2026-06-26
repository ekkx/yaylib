// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3TokenInfosDto {
    #[serde(rename = "token_infos", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_infos: Option<Option<Vec<models::Web3TokenInfoDto>>>,
}

impl Web3TokenInfosDto {
    pub fn new() -> Web3TokenInfosDto {
        Web3TokenInfosDto {
            token_infos: None,
        }
    }
}

