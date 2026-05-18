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
 * @interface BCNetworkData
 */
export interface BCNetworkData {
    /**
     * 
     * @type {Web3WalletGasPercent}
     * @memberof BCNetworkData
     */
    additionGasPercent?: Web3WalletGasPercent | null;
    /**
     * 
     * @type {number}
     * @memberof BCNetworkData
     */
    chainId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof BCNetworkData
     */
    chainUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof BCNetworkData
     */
    name?: string | null;
}

/**
 * Check if a given object implements the BCNetworkData interface.
 */
export function instanceOfBCNetworkData(value: object): value is BCNetworkData {
    return true;
}

export function BCNetworkDataFromJSON(json: any): BCNetworkData {
    return BCNetworkDataFromJSONTyped(json, false);
}

export function BCNetworkDataFromJSONTyped(json: any, ignoreDiscriminator: boolean): BCNetworkData {
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

export function BCNetworkDataToJSON(json: any): BCNetworkData {
    return BCNetworkDataToJSONTyped(json, false);
}

export function BCNetworkDataToJSONTyped(value?: BCNetworkData | null, ignoreDiscriminator: boolean = false): any {
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

