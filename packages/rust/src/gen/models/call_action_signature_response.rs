// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CallActionSignatureResponse {
    #[serde(rename = "signature_payload", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub signature_payload: Option<Option<models::SignaturePayload>>,
}

impl CallActionSignatureResponse {
    pub fn new() -> CallActionSignatureResponse {
        CallActionSignatureResponse {
            signature_payload: None,
        }
    }
}

