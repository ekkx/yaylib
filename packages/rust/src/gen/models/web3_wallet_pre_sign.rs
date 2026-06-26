// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3WalletPreSign {
    #[serde(rename = "blockchain_network_data", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub blockchain_network_data: Option<Option<models::BcNetworkData>>,
    #[serde(rename = "transaction_and_log_data", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub transaction_and_log_data: Option<Option<models::TxnLogData>>,
}

impl Web3WalletPreSign {
    pub fn new() -> Web3WalletPreSign {
        Web3WalletPreSign {
            blockchain_network_data: None,
            transaction_and_log_data: None,
        }
    }
}

