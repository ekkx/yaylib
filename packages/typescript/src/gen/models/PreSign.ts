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

import { mapValues } from '../runtime';
import type { Web3WalletPreSign } from './Web3WalletPreSign';
import {
    Web3WalletPreSignFromJSON,
    Web3WalletPreSignFromJSONTyped,
    Web3WalletPreSignToJSON,
    Web3WalletPreSignToJSONTyped,
} from './Web3WalletPreSign';

/**
 * 
 * @export
 * @interface PreSign
 */
export interface PreSign {
    /**
     * 
     * @type {Web3WalletPreSign}
     * @memberof PreSign
     */
    result?: Web3WalletPreSign | null;
}

/**
 * Check if a given object implements the PreSign interface.
 */
export function instanceOfPreSign(value: object): value is PreSign {
    return true;
}

export function PreSignFromJSON(json: any): PreSign {
    return PreSignFromJSONTyped(json, false);
}

export function PreSignFromJSONTyped(json: any, ignoreDiscriminator: boolean): PreSign {
    if (json == null) {
        return json;
    }
    return {
        
        'result': json['result'] == null ? undefined : Web3WalletPreSignFromJSON(json['result']),
    };
}

export function PreSignToJSON(json: any): PreSign {
    return PreSignToJSONTyped(json, false);
}

export function PreSignToJSONTyped(value?: PreSign | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'result': Web3WalletPreSignToJSON(value['result']),
    };
}

