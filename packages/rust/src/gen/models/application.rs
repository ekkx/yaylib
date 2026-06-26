// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Application {
    #[serde(rename = "settings", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub settings: Option<Option<models::Config>>,
}

impl Application {
    pub fn new() -> Application {
        Application {
            settings: None,
        }
    }
}

