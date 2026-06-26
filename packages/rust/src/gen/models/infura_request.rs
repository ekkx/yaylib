// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct InfuraRequest {
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "jsonrpc", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub jsonrpc: Option<Option<String>>,
    #[serde(rename = "method", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub method: Option<Option<String>>,
    #[serde(rename = "params", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub params: Option<Option<Vec<serde_json::Value>>>,
}

impl InfuraRequest {
    pub fn new() -> InfuraRequest {
        InfuraRequest {
            id: None,
            jsonrpc: None,
            method: None,
            params: None,
        }
    }
}

