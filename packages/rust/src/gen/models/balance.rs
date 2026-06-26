// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Balance {
    #[serde(rename = "egg", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub egg: Option<Option<String>>,
    #[serde(rename = "empl", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub empl: Option<Option<String>>,
    #[serde(rename = "l1_native", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub l1_native: Option<Option<String>>,
    #[serde(rename = "l1_yay", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub l1_yay: Option<Option<String>>,
    #[serde(rename = "native", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub native: Option<Option<String>>,
    #[serde(rename = "pal", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal: Option<Option<String>>,
    #[serde(rename = "yay", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub yay: Option<Option<String>>,
}

impl Balance {
    pub fn new() -> Balance {
        Balance {
            egg: None,
            empl: None,
            l1_native: None,
            l1_yay: None,
            native: None,
            pal: None,
            yay: None,
        }
    }
}

