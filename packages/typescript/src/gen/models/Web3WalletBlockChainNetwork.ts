/* tslint:disable */
/* eslint-disable */
/**
 * Yay! API
 * No description provided
 *
 * Schema version: 4.26.1
 * 
 *
 * NOTE: Code generated; DO NOT EDIT.
 * 
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime.js';
import type { ModelWeb3WalletGasPercent } from './ModelWeb3WalletGasPercent.js';
import {
    ModelWeb3WalletGasPercentFromJSON,
    ModelWeb3WalletGasPercentFromJSONTyped,
    ModelWeb3WalletGasPercentToJSON,
    ModelWeb3WalletGasPercentToJSONTyped,
} from './ModelWeb3WalletGasPercent.js';

/**
 * 
 * @export
 * @interface Web3WalletBlockChainNetwork
 */
export interface Web3WalletBlockChainNetwork {
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockChainNetwork
     */
    l1AmmAddress?: string | null;
    /**
     * 
     * @type {ModelWeb3WalletGasPercent}
     * @memberof Web3WalletBlockChainNetwork
     */
    l1BlockChainNetworkAdditionGasPercent?: ModelWeb3WalletGasPercent | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletBlockChainNetwork
     */
    l1BlockChainNetworkId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockChainNetwork
     */
    l1BlockChainNetworkName?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockChainNetwork
     */
    l1BlockChainNetworkUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockChainNetwork
     */
    l1QuoterAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockChainNetwork
     */
    l1WethAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockChainNetwork
     */
    l2AmmAddress?: string | null;
    /**
     * 
     * @type {ModelWeb3WalletGasPercent}
     * @memberof Web3WalletBlockChainNetwork
     */
    l2BlockChainNetworkAdditionGasPercent?: ModelWeb3WalletGasPercent | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletBlockChainNetwork
     */
    l2BlockChainNetworkId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockChainNetwork
     */
    l2BlockChainNetworkName?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockChainNetwork
     */
    l2BlockChainNetworkUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockChainNetwork
     */
    l2QuoterAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockChainNetwork
     */
    l2WethAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockChainNetwork
     */
    miscPalBaseImageUri?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockChainNetwork
     */
    miscPalBaseJsonUri?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockChainNetwork
     */
    smartContractEmplAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockChainNetwork
     */
    smartContractEmplDepositAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockChainNetwork
     */
    smartContractEmplWithdrawAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockChainNetwork
     */
    smartContractL1YayAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockChainNetwork
     */
    smartContractL2YayAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockChainNetwork
     */
    smartContractPalAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockChainNetwork
     */
    smartContractPalDepositWithdrawAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockChainNetwork
     */
    smartContractX2y2delegateerc721Address?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockChainNetwork
     */
    smartContractX2y2marketAddress?: string | null;
}

/**
 * Check if a given object implements the Web3WalletBlockChainNetwork interface.
 */
export function instanceOfWeb3WalletBlockChainNetwork(value: object): value is Web3WalletBlockChainNetwork {
    return true;
}

export function Web3WalletBlockChainNetworkFromJSON(json: any): Web3WalletBlockChainNetwork {
    return Web3WalletBlockChainNetworkFromJSONTyped(json, false);
}

export function Web3WalletBlockChainNetworkFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3WalletBlockChainNetwork {
    if (json == null) {
        return json;
    }
    return {
        
        'l1AmmAddress': json['l1_amm_address'] == null ? undefined : json['l1_amm_address'],
        'l1BlockChainNetworkAdditionGasPercent': json['l1_block_chain_network_addition_gas_percent'] == null ? undefined : ModelWeb3WalletGasPercentFromJSON(json['l1_block_chain_network_addition_gas_percent']),
        'l1BlockChainNetworkId': json['l1_block_chain_network_id'] == null ? undefined : json['l1_block_chain_network_id'],
        'l1BlockChainNetworkName': json['l1_block_chain_network_name'] == null ? undefined : json['l1_block_chain_network_name'],
        'l1BlockChainNetworkUrl': json['l1_block_chain_network_url'] == null ? undefined : json['l1_block_chain_network_url'],
        'l1QuoterAddress': json['l1_quoter_address'] == null ? undefined : json['l1_quoter_address'],
        'l1WethAddress': json['l1_weth_address'] == null ? undefined : json['l1_weth_address'],
        'l2AmmAddress': json['l2_amm_address'] == null ? undefined : json['l2_amm_address'],
        'l2BlockChainNetworkAdditionGasPercent': json['l2_block_chain_network_addition_gas_percent'] == null ? undefined : ModelWeb3WalletGasPercentFromJSON(json['l2_block_chain_network_addition_gas_percent']),
        'l2BlockChainNetworkId': json['l2_block_chain_network_id'] == null ? undefined : json['l2_block_chain_network_id'],
        'l2BlockChainNetworkName': json['l2_block_chain_network_name'] == null ? undefined : json['l2_block_chain_network_name'],
        'l2BlockChainNetworkUrl': json['l2_block_chain_network_url'] == null ? undefined : json['l2_block_chain_network_url'],
        'l2QuoterAddress': json['l2_quoter_address'] == null ? undefined : json['l2_quoter_address'],
        'l2WethAddress': json['l2_weth_address'] == null ? undefined : json['l2_weth_address'],
        'miscPalBaseImageUri': json['misc_pal_base_image_uri'] == null ? undefined : json['misc_pal_base_image_uri'],
        'miscPalBaseJsonUri': json['misc_pal_base_json_uri'] == null ? undefined : json['misc_pal_base_json_uri'],
        'smartContractEmplAddress': json['smart_contract_empl_address'] == null ? undefined : json['smart_contract_empl_address'],
        'smartContractEmplDepositAddress': json['smart_contract_empl_deposit_address'] == null ? undefined : json['smart_contract_empl_deposit_address'],
        'smartContractEmplWithdrawAddress': json['smart_contract_empl_withdraw_address'] == null ? undefined : json['smart_contract_empl_withdraw_address'],
        'smartContractL1YayAddress': json['smart_contract_l1_yay_address'] == null ? undefined : json['smart_contract_l1_yay_address'],
        'smartContractL2YayAddress': json['smart_contract_l2_yay_address'] == null ? undefined : json['smart_contract_l2_yay_address'],
        'smartContractPalAddress': json['smart_contract_pal_address'] == null ? undefined : json['smart_contract_pal_address'],
        'smartContractPalDepositWithdrawAddress': json['smart_contract_pal_deposit_withdraw_address'] == null ? undefined : json['smart_contract_pal_deposit_withdraw_address'],
        'smartContractX2y2delegateerc721Address': json['smart_contract_x2y2delegateerc721_address'] == null ? undefined : json['smart_contract_x2y2delegateerc721_address'],
        'smartContractX2y2marketAddress': json['smart_contract_x2y2market_address'] == null ? undefined : json['smart_contract_x2y2market_address'],
    };
}

export function Web3WalletBlockChainNetworkToJSON(json: any): Web3WalletBlockChainNetwork {
    return Web3WalletBlockChainNetworkToJSONTyped(json, false);
}

export function Web3WalletBlockChainNetworkToJSONTyped(value?: Web3WalletBlockChainNetwork | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'l1_amm_address': value['l1AmmAddress'],
        'l1_block_chain_network_addition_gas_percent': ModelWeb3WalletGasPercentToJSON(value['l1BlockChainNetworkAdditionGasPercent']),
        'l1_block_chain_network_id': value['l1BlockChainNetworkId'],
        'l1_block_chain_network_name': value['l1BlockChainNetworkName'],
        'l1_block_chain_network_url': value['l1BlockChainNetworkUrl'],
        'l1_quoter_address': value['l1QuoterAddress'],
        'l1_weth_address': value['l1WethAddress'],
        'l2_amm_address': value['l2AmmAddress'],
        'l2_block_chain_network_addition_gas_percent': ModelWeb3WalletGasPercentToJSON(value['l2BlockChainNetworkAdditionGasPercent']),
        'l2_block_chain_network_id': value['l2BlockChainNetworkId'],
        'l2_block_chain_network_name': value['l2BlockChainNetworkName'],
        'l2_block_chain_network_url': value['l2BlockChainNetworkUrl'],
        'l2_quoter_address': value['l2QuoterAddress'],
        'l2_weth_address': value['l2WethAddress'],
        'misc_pal_base_image_uri': value['miscPalBaseImageUri'],
        'misc_pal_base_json_uri': value['miscPalBaseJsonUri'],
        'smart_contract_empl_address': value['smartContractEmplAddress'],
        'smart_contract_empl_deposit_address': value['smartContractEmplDepositAddress'],
        'smart_contract_empl_withdraw_address': value['smartContractEmplWithdrawAddress'],
        'smart_contract_l1_yay_address': value['smartContractL1YayAddress'],
        'smart_contract_l2_yay_address': value['smartContractL2YayAddress'],
        'smart_contract_pal_address': value['smartContractPalAddress'],
        'smart_contract_pal_deposit_withdraw_address': value['smartContractPalDepositWithdrawAddress'],
        'smart_contract_x2y2delegateerc721_address': value['smartContractX2y2delegateerc721Address'],
        'smart_contract_x2y2market_address': value['smartContractX2y2marketAddress'],
    };
}

