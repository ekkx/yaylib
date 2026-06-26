// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Contact {
    #[serde(rename = "email", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub email: Option<Option<String>>,
    #[serde(rename = "mobile_number", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub mobile_number: Option<Option<String>>,
}

impl Contact {
    pub fn new() -> Contact {
        Contact {
            email: None,
            mobile_number: None,
        }
    }
}

