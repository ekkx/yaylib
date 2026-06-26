// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Navigation {
    #[serde(rename = "navigation_item_list", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub navigation_item_list: Option<Option<Vec<models::SectionNavigation>>>,
}

impl Navigation {
    pub fn new() -> Navigation {
        Navigation {
            navigation_item_list: None,
        }
    }
}

