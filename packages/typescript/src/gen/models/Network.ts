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
 * @interface Network
 */
export interface Network {
    /**
     * 
     * @type {ModelWeb3WalletGasPercent}
     * @memberof Network
     */
    additionGasPercent?: ModelWeb3WalletGasPercent | null;
    /**
     * 
     * @type {string}
     * @memberof Network
     */
    ammAddress?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Network
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Network
     */
    name?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Network
     */
    quoterAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Network
     */
    url?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Network
     */
    wethAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Network
     */
    yayAddress?: string | null;
}

/**
 * Check if a given object implements the Network interface.
 */
export function instanceOfNetwork(value: object): value is Network {
    return true;
}

export function NetworkFromJSON(json: any): Network {
    return NetworkFromJSONTyped(json, false);
}

export function NetworkFromJSONTyped(json: any, ignoreDiscriminator: boolean): Network {
    if (json == null) {
        return json;
    }
    return {
        
        'additionGasPercent': json['addition_gas_percent'] == null ? undefined : ModelWeb3WalletGasPercentFromJSON(json['addition_gas_percent']),
        'ammAddress': json['amm_address'] == null ? undefined : json['amm_address'],
        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'quoterAddress': json['quoter_address'] == null ? undefined : json['quoter_address'],
        'url': json['url'] == null ? undefined : json['url'],
        'wethAddress': json['weth_address'] == null ? undefined : json['weth_address'],
        'yayAddress': json['yay_address'] == null ? undefined : json['yay_address'],
    };
}

export function NetworkToJSON(json: any): Network {
    return NetworkToJSONTyped(json, false);
}

export function NetworkToJSONTyped(value?: Network | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'addition_gas_percent': ModelWeb3WalletGasPercentToJSON(value['additionGasPercent']),
        'amm_address': value['ammAddress'],
        'id': value['id'],
        'name': value['name'],
        'quoter_address': value['quoterAddress'],
        'url': value['url'],
        'weth_address': value['wethAddress'],
        'yay_address': value['yayAddress'],
    };
}

