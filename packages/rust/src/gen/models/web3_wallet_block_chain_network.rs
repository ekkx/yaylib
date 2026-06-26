// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3WalletBlockChainNetwork {
    #[serde(rename = "l1_amm_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub l1_amm_address: Option<Option<String>>,
    #[serde(rename = "l1_block_chain_network_addition_gas_percent", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub l1_block_chain_network_addition_gas_percent: Option<Option<models::ModelWeb3WalletGasPercent>>,
    #[serde(rename = "l1_block_chain_network_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub l1_block_chain_network_id: Option<Option<i64>>,
    #[serde(rename = "l1_block_chain_network_name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub l1_block_chain_network_name: Option<Option<String>>,
    #[serde(rename = "l1_block_chain_network_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub l1_block_chain_network_url: Option<Option<String>>,
    #[serde(rename = "l1_quoter_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub l1_quoter_address: Option<Option<String>>,
    #[serde(rename = "l1_weth_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub l1_weth_address: Option<Option<String>>,
    #[serde(rename = "l2_amm_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub l2_amm_address: Option<Option<String>>,
    #[serde(rename = "l2_block_chain_network_addition_gas_percent", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub l2_block_chain_network_addition_gas_percent: Option<Option<models::ModelWeb3WalletGasPercent>>,
    #[serde(rename = "l2_block_chain_network_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub l2_block_chain_network_id: Option<Option<i64>>,
    #[serde(rename = "l2_block_chain_network_name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub l2_block_chain_network_name: Option<Option<String>>,
    #[serde(rename = "l2_block_chain_network_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub l2_block_chain_network_url: Option<Option<String>>,
    #[serde(rename = "l2_quoter_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub l2_quoter_address: Option<Option<String>>,
    #[serde(rename = "l2_weth_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub l2_weth_address: Option<Option<String>>,
    #[serde(rename = "misc_pal_base_image_uri", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub misc_pal_base_image_uri: Option<Option<String>>,
    #[serde(rename = "misc_pal_base_json_uri", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub misc_pal_base_json_uri: Option<Option<String>>,
    #[serde(rename = "smart_contract_empl_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub smart_contract_empl_address: Option<Option<String>>,
    #[serde(rename = "smart_contract_empl_deposit_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub smart_contract_empl_deposit_address: Option<Option<String>>,
    #[serde(rename = "smart_contract_empl_withdraw_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub smart_contract_empl_withdraw_address: Option<Option<String>>,
    #[serde(rename = "smart_contract_l1_yay_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub smart_contract_l1_yay_address: Option<Option<String>>,
    #[serde(rename = "smart_contract_l2_yay_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub smart_contract_l2_yay_address: Option<Option<String>>,
    #[serde(rename = "smart_contract_pal_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub smart_contract_pal_address: Option<Option<String>>,
    #[serde(rename = "smart_contract_pal_deposit_withdraw_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub smart_contract_pal_deposit_withdraw_address: Option<Option<String>>,
    #[serde(rename = "smart_contract_x2y2delegateerc721_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub smart_contract_x2y2delegateerc721_address: Option<Option<String>>,
    #[serde(rename = "smart_contract_x2y2market_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub smart_contract_x2y2market_address: Option<Option<String>>,
}

impl Web3WalletBlockChainNetwork {
    pub fn new() -> Web3WalletBlockChainNetwork {
        Web3WalletBlockChainNetwork {
            l1_amm_address: None,
            l1_block_chain_network_addition_gas_percent: None,
            l1_block_chain_network_id: None,
            l1_block_chain_network_name: None,
            l1_block_chain_network_url: None,
            l1_quoter_address: None,
            l1_weth_address: None,
            l2_amm_address: None,
            l2_block_chain_network_addition_gas_percent: None,
            l2_block_chain_network_id: None,
            l2_block_chain_network_name: None,
            l2_block_chain_network_url: None,
            l2_quoter_address: None,
            l2_weth_address: None,
            misc_pal_base_image_uri: None,
            misc_pal_base_json_uri: None,
            smart_contract_empl_address: None,
            smart_contract_empl_deposit_address: None,
            smart_contract_empl_withdraw_address: None,
            smart_contract_l1_yay_address: None,
            smart_contract_l2_yay_address: None,
            smart_contract_pal_address: None,
            smart_contract_pal_deposit_withdraw_address: None,
            smart_contract_x2y2delegateerc721_address: None,
            smart_contract_x2y2market_address: None,
        }
    }
}

