// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct LevelUpDetailsPal {
    #[serde(rename = "image", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub image: Option<Option<String>>,
}

impl LevelUpDetailsPal {
    pub fn new() -> LevelUpDetailsPal {
        LevelUpDetailsPal {
            image: None,
        }
    }
}

