// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct RefreshCounterRequestsResponse {
    #[serde(rename = "reset_counter_requests", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub reset_counter_requests: Option<Option<Vec<models::RefreshCounterRequest>>>,
}

impl RefreshCounterRequestsResponse {
    pub fn new() -> RefreshCounterRequestsResponse {
        RefreshCounterRequestsResponse {
            reset_counter_requests: None,
        }
    }
}

