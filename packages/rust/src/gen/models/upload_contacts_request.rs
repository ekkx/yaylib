// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct UploadContactsRequest {
    #[serde(rename = "contacts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub contacts: Option<Option<Vec<models::Contact>>>,
}

impl UploadContactsRequest {
    pub fn new() -> UploadContactsRequest {
        UploadContactsRequest {
            contacts: None,
        }
    }
}

