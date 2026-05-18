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
 * @interface Web3WalletEmplExpiring
 */
export interface Web3WalletEmplExpiring {
    /**
     * 
     * @type {number}
     * @memberof Web3WalletEmplExpiring
     */
    expiringAmount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletEmplExpiring
     */
    expiringDateMillis?: number | null;
}

/**
 * Check if a given object implements the Web3WalletEmplExpiring interface.
 */
export function instanceOfWeb3WalletEmplExpiring(value: object): value is Web3WalletEmplExpiring {
    return true;
}

export function Web3WalletEmplExpiringFromJSON(json: any): Web3WalletEmplExpiring {
    return Web3WalletEmplExpiringFromJSONTyped(json, false);
}

export function Web3WalletEmplExpiringFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3WalletEmplExpiring {
    if (json == null) {
        return json;
    }
    return {
        
        'expiringAmount': json['expiring_amount'] == null ? undefined : json['expiring_amount'],
        'expiringDateMillis': json['expiring_date_millis'] == null ? undefined : json['expiring_date_millis'],
    };
}

export function Web3WalletEmplExpiringToJSON(json: any): Web3WalletEmplExpiring {
    return Web3WalletEmplExpiringToJSONTyped(json, false);
}

export function Web3WalletEmplExpiringToJSONTyped(value?: Web3WalletEmplExpiring | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'expiring_amount': value['expiringAmount'],
        'expiring_date_millis': value['expiringDateMillis'],
    };
}

