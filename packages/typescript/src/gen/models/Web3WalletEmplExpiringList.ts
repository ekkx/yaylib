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
import type { Web3WalletEmplExpiring } from './Web3WalletEmplExpiring.js';
import {
    Web3WalletEmplExpiringFromJSON,
    Web3WalletEmplExpiringFromJSONTyped,
    Web3WalletEmplExpiringToJSON,
    Web3WalletEmplExpiringToJSONTyped,
} from './Web3WalletEmplExpiring.js';

/**
 * 
 * @export
 * @interface Web3WalletEmplExpiringList
 */
export interface Web3WalletEmplExpiringList {
    /**
     * 
     * @type {Array<Web3WalletEmplExpiring>}
     * @memberof Web3WalletEmplExpiringList
     */
    regularExpiring?: Array<Web3WalletEmplExpiring> | null;
    /**
     * 
     * @type {Array<Web3WalletEmplExpiring>}
     * @memberof Web3WalletEmplExpiringList
     */
    rewardedExpiring?: Array<Web3WalletEmplExpiring> | null;
}

/**
 * Check if a given object implements the Web3WalletEmplExpiringList interface.
 */
export function instanceOfWeb3WalletEmplExpiringList(value: object): value is Web3WalletEmplExpiringList {
    return true;
}

export function Web3WalletEmplExpiringListFromJSON(json: any): Web3WalletEmplExpiringList {
    return Web3WalletEmplExpiringListFromJSONTyped(json, false);
}

export function Web3WalletEmplExpiringListFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3WalletEmplExpiringList {
    if (json == null) {
        return json;
    }
    return {
        
        'regularExpiring': json['regular_expiring'] == null ? undefined : ((json['regular_expiring'] as Array<any>).map(Web3WalletEmplExpiringFromJSON)),
        'rewardedExpiring': json['rewarded_expiring'] == null ? undefined : ((json['rewarded_expiring'] as Array<any>).map(Web3WalletEmplExpiringFromJSON)),
    };
}

export function Web3WalletEmplExpiringListToJSON(json: any): Web3WalletEmplExpiringList {
    return Web3WalletEmplExpiringListToJSONTyped(json, false);
}

export function Web3WalletEmplExpiringListToJSONTyped(value?: Web3WalletEmplExpiringList | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'regular_expiring': value['regularExpiring'] == null ? undefined : ((value['regularExpiring'] as Array<any>).map(Web3WalletEmplExpiringToJSON)),
        'rewarded_expiring': value['rewardedExpiring'] == null ? undefined : ((value['rewardedExpiring'] as Array<any>).map(Web3WalletEmplExpiringToJSON)),
    };
}

