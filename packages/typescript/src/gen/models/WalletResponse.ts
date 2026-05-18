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
import type { CoinAmount } from './CoinAmount.js';
import {
    CoinAmountFromJSON,
    CoinAmountFromJSONTyped,
    CoinAmountToJSON,
    CoinAmountToJSONTyped,
} from './CoinAmount.js';

/**
 * 
 * @export
 * @interface WalletResponse
 */
export interface WalletResponse {
    /**
     * 
     * @type {CoinAmount}
     * @memberof WalletResponse
     */
    coins?: CoinAmount | null;
}

/**
 * Check if a given object implements the WalletResponse interface.
 */
export function instanceOfWalletResponse(value: object): value is WalletResponse {
    return true;
}

export function WalletResponseFromJSON(json: any): WalletResponse {
    return WalletResponseFromJSONTyped(json, false);
}

export function WalletResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): WalletResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'coins': json['coins'] == null ? undefined : CoinAmountFromJSON(json['coins']),
    };
}

export function WalletResponseToJSON(json: any): WalletResponse {
    return WalletResponseToJSONTyped(json, false);
}

export function WalletResponseToJSONTyped(value?: WalletResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'coins': CoinAmountToJSON(value['coins']),
    };
}

