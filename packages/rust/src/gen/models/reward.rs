// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Reward {
    #[serde(rename = "rewarded_empl", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub rewarded_empl: Option<Option<i32>>,
}

impl Reward {
    pub fn new() -> Reward {
        Reward {
            rewarded_empl: None,
        }
    }
}

