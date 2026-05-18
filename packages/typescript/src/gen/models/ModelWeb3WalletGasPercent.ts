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
/**
 * 
 * @export
 * @interface ModelWeb3WalletGasPercent
 */
export interface ModelWeb3WalletGasPercent {
    /**
     * 
     * @type {number}
     * @memberof ModelWeb3WalletGasPercent
     */
    fast?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelWeb3WalletGasPercent
     */
    normal?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelWeb3WalletGasPercent
     */
    slow?: number | null;
}

/**
 * Check if a given object implements the ModelWeb3WalletGasPercent interface.
 */
export function instanceOfModelWeb3WalletGasPercent(value: object): value is ModelWeb3WalletGasPercent {
    return true;
}

export function ModelWeb3WalletGasPercentFromJSON(json: any): ModelWeb3WalletGasPercent {
    return ModelWeb3WalletGasPercentFromJSONTyped(json, false);
}

export function ModelWeb3WalletGasPercentFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelWeb3WalletGasPercent {
    if (json == null) {
        return json;
    }
    return {
        
        'fast': json['fast'] == null ? undefined : json['fast'],
        'normal': json['normal'] == null ? undefined : json['normal'],
        'slow': json['slow'] == null ? undefined : json['slow'],
    };
}

export function ModelWeb3WalletGasPercentToJSON(json: any): ModelWeb3WalletGasPercent {
    return ModelWeb3WalletGasPercentToJSONTyped(json, false);
}

export function ModelWeb3WalletGasPercentToJSONTyped(value?: ModelWeb3WalletGasPercent | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'fast': value['fast'],
        'normal': value['normal'],
        'slow': value['slow'],
    };
}

