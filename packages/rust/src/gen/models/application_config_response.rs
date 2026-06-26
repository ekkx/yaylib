// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ApplicationConfigResponse {
    #[serde(rename = "app", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub app: Option<Option<models::Application>>,
}

impl ApplicationConfigResponse {
    pub fn new() -> ApplicationConfigResponse {
        ApplicationConfigResponse {
            app: None,
        }
    }
}

