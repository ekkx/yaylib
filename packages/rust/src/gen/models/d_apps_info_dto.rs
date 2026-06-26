// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct DAppsInfoDto {
    #[serde(rename = "dapps", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub dapps: Option<Option<Vec<models::DAppDto>>>,
    #[serde(rename = "educational_links", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub educational_links: Option<Option<Vec<models::DAppDto>>>,
}

impl DAppsInfoDto {
    pub fn new() -> DAppsInfoDto {
        DAppsInfoDto {
            dapps: None,
            educational_links: None,
        }
    }
}

