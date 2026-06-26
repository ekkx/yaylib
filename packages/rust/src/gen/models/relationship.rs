// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Relationship {
    #[serde(rename = "target", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub target: Option<Option<models::Target>>,
}

impl Relationship {
    pub fn new() -> Relationship {
        Relationship {
            target: None,
        }
    }
}

