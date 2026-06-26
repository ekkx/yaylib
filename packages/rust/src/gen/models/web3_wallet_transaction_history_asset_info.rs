// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3WalletTransactionHistoryAssetInfo {
    #[serde(rename = "image_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub image_url: Option<Option<String>>,
    #[serde(rename = "name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub name: Option<Option<String>>,
    #[serde(rename = "token_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_id: Option<Option<String>>,
    #[serde(rename = "video_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub video_url: Option<Option<String>>,
}

impl Web3WalletTransactionHistoryAssetInfo {
    pub fn new() -> Web3WalletTransactionHistoryAssetInfo {
        Web3WalletTransactionHistoryAssetInfo {
            image_url: None,
            name: None,
            token_id: None,
            video_url: None,
        }
    }
}

