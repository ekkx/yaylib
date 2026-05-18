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
 * @interface ModelWeb3WalletPreSign
 */
export interface ModelWeb3WalletPreSign {
    /**
     * 
     * @type {ModelWeb3WalletGasPercent}
     * @memberof ModelWeb3WalletPreSign
     */
    additionGasPercent?: ModelWeb3WalletGasPercent | null;
    /**
     * 
     * @type {number}
     * @memberof ModelWeb3WalletPreSign
     */
    chainId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ModelWeb3WalletPreSign
     */
    chainUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelWeb3WalletPreSign
     */
    contractAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelWeb3WalletPreSign
     */
    inputsHex?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelWeb3WalletPreSign
     */
    name?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelWeb3WalletPreSign
     */
    txUniqueId?: string | null;
}

/**
 * Check if a given object implements the ModelWeb3WalletPreSign interface.
 */
export function instanceOfModelWeb3WalletPreSign(value: object): value is ModelWeb3WalletPreSign {
    return true;
}

export function ModelWeb3WalletPreSignFromJSON(json: any): ModelWeb3WalletPreSign {
    return ModelWeb3WalletPreSignFromJSONTyped(json, false);
}

export function ModelWeb3WalletPreSignFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelWeb3WalletPreSign {
    if (json == null) {
        return json;
    }
    return {
        
        'additionGasPercent': json['addition_gas_percent'] == null ? undefined : ModelWeb3WalletGasPercentFromJSON(json['addition_gas_percent']),
        'chainId': json['chain_id'] == null ? undefined : json['chain_id'],
        'chainUrl': json['chain_url'] == null ? undefined : json['chain_url'],
        'contractAddress': json['contract_address'] == null ? undefined : json['contract_address'],
        'inputsHex': json['inputs_hex'] == null ? undefined : json['inputs_hex'],
        'name': json['name'] == null ? undefined : json['name'],
        'txUniqueId': json['tx_unique_id'] == null ? undefined : json['tx_unique_id'],
    };
}

export function ModelWeb3WalletPreSignToJSON(json: any): ModelWeb3WalletPreSign {
    return ModelWeb3WalletPreSignToJSONTyped(json, false);
}

export function ModelWeb3WalletPreSignToJSONTyped(value?: ModelWeb3WalletPreSign | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'addition_gas_percent': ModelWeb3WalletGasPercentToJSON(value['additionGasPercent']),
        'chain_id': value['chainId'],
        'chain_url': value['chainUrl'],
        'contract_address': value['contractAddress'],
        'inputs_hex': value['inputsHex'],
        'name': value['name'],
        'tx_unique_id': value['txUniqueId'],
    };
}

