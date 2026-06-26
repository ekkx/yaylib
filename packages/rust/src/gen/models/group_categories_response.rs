// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GroupCategoriesResponse {
    #[serde(rename = "group_categories", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group_categories: Option<Option<Vec<models::GroupCategory>>>,
}

impl GroupCategoriesResponse {
    pub fn new() -> GroupCategoriesResponse {
        GroupCategoriesResponse {
            group_categories: None,
        }
    }
}

