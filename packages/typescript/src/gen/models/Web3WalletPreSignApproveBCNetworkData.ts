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
import type { Web3WalletGasPercent } from './Web3WalletGasPercent.js';
import {
    Web3WalletGasPercentFromJSON,
    Web3WalletGasPercentFromJSONTyped,
    Web3WalletGasPercentToJSON,
    Web3WalletGasPercentToJSONTyped,
} from './Web3WalletGasPercent.js';

/**
 * 
 * @export
 * @interface Web3WalletPreSignApproveBCNetworkData
 */
export interface Web3WalletPreSignApproveBCNetworkData {
    /**
     * 
     * @type {Web3WalletGasPercent}
     * @memberof Web3WalletPreSignApproveBCNetworkData
     */
    additionGasPercent?: Web3WalletGasPercent | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletPreSignApproveBCNetworkData
     */
    chainId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletPreSignApproveBCNetworkData
     */
    chainUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletPreSignApproveBCNetworkData
     */
    name?: string | null;
}

/**
 * Check if a given object implements the Web3WalletPreSignApproveBCNetworkData interface.
 */
export function instanceOfWeb3WalletPreSignApproveBCNetworkData(value: object): value is Web3WalletPreSignApproveBCNetworkData {
    return true;
}

export function Web3WalletPreSignApproveBCNetworkDataFromJSON(json: any): Web3WalletPreSignApproveBCNetworkData {
    return Web3WalletPreSignApproveBCNetworkDataFromJSONTyped(json, false);
}

export function Web3WalletPreSignApproveBCNetworkDataFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3WalletPreSignApproveBCNetworkData {
    if (json == null) {
        return json;
    }
    return {
        
        'additionGasPercent': json['addition_gas_percent'] == null ? undefined : Web3WalletGasPercentFromJSON(json['addition_gas_percent']),
        'chainId': json['chain_id'] == null ? undefined : json['chain_id'],
        'chainUrl': json['chain_url'] == null ? undefined : json['chain_url'],
        'name': json['name'] == null ? undefined : json['name'],
    };
}

export function Web3WalletPreSignApproveBCNetworkDataToJSON(json: any): Web3WalletPreSignApproveBCNetworkData {
    return Web3WalletPreSignApproveBCNetworkDataToJSONTyped(json, false);
}

export function Web3WalletPreSignApproveBCNetworkDataToJSONTyped(value?: Web3WalletPreSignApproveBCNetworkData | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'addition_gas_percent': Web3WalletGasPercentToJSON(value['additionGasPercent']),
        'chain_id': value['chainId'],
        'chain_url': value['chainUrl'],
        'name': value['name'],
    };
}

