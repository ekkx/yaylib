// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct JsEvent {
    #[serde(rename = "args", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub args: Option<Option<String>>,
    #[serde(rename = "name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub name: Option<Option<String>>,
}

impl JsEvent {
    pub fn new() -> JsEvent {
        JsEvent {
            args: None,
            name: None,
        }
    }
}

