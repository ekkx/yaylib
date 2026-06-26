// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Request {
    #[serde(rename = "method", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub method: Option<Option<models::RequestMethod>>,
    #[serde(rename = "params", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub params: Option<Option<serde_json::Value>>,
}

impl Request {
    pub fn new() -> Request {
        Request {
            method: None,
            params: None,
        }
    }
}

