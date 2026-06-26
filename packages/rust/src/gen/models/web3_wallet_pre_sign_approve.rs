// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3WalletPreSignApprove {
    #[serde(rename = "blockchain_network_data", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub blockchain_network_data: Option<Option<models::Web3WalletPreSignApproveBcNetworkData>>,
    #[serde(rename = "transaction_data", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub transaction_data: Option<Option<models::TxnData>>,
}

impl Web3WalletPreSignApprove {
    pub fn new() -> Web3WalletPreSignApprove {
        Web3WalletPreSignApprove {
            blockchain_network_data: None,
            transaction_data: None,
        }
    }
}

