// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PostGiftCardExchange {
    #[serde(rename = "email", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub email: Option<Option<String>>,
    #[serde(rename = "gift_card", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gift_card: Option<Option<String>>,
}

impl PostGiftCardExchange {
    pub fn new() -> PostGiftCardExchange {
        PostGiftCardExchange {
            email: None,
            gift_card: None,
        }
    }
}

