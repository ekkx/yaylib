// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct HimaUsersResponse {
    #[serde(rename = "hima_users", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hima_users: Option<Option<Vec<models::UserWrapper>>>,
}

impl HimaUsersResponse {
    pub fn new() -> HimaUsersResponse {
        HimaUsersResponse {
            hima_users: None,
        }
    }
}

