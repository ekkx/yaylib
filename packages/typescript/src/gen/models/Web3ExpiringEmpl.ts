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
import type { Web3ExpiringEmplExpiredEmpl } from './Web3ExpiringEmplExpiredEmpl.js';
import {
    Web3ExpiringEmplExpiredEmplFromJSON,
    Web3ExpiringEmplExpiredEmplFromJSONTyped,
    Web3ExpiringEmplExpiredEmplToJSON,
    Web3ExpiringEmplExpiredEmplToJSONTyped,
} from './Web3ExpiringEmplExpiredEmpl.js';

/**
 * 
 * @export
 * @interface Web3ExpiringEmpl
 */
export interface Web3ExpiringEmpl {
    /**
     * 
     * @type {Array<Web3ExpiringEmplExpiredEmpl>}
     * @memberof Web3ExpiringEmpl
     */
    items?: Array<Web3ExpiringEmplExpiredEmpl> | null;
    /**
     * 
     * @type {number}
     * @memberof Web3ExpiringEmpl
     */
    totalAmount?: number | null;
}

/**
 * Check if a given object implements the Web3ExpiringEmpl interface.
 */
export function instanceOfWeb3ExpiringEmpl(value: object): value is Web3ExpiringEmpl {
    return true;
}

export function Web3ExpiringEmplFromJSON(json: any): Web3ExpiringEmpl {
    return Web3ExpiringEmplFromJSONTyped(json, false);
}

export function Web3ExpiringEmplFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3ExpiringEmpl {
    if (json == null) {
        return json;
    }
    return {
        
        'items': json['items'] == null ? undefined : ((json['items'] as Array<any>).map(Web3ExpiringEmplExpiredEmplFromJSON)),
        'totalAmount': json['total_amount'] == null ? undefined : json['total_amount'],
    };
}

export function Web3ExpiringEmplToJSON(json: any): Web3ExpiringEmpl {
    return Web3ExpiringEmplToJSONTyped(json, false);
}

export function Web3ExpiringEmplToJSONTyped(value?: Web3ExpiringEmpl | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'items': value['items'] == null ? undefined : ((value['items'] as Array<any>).map(Web3ExpiringEmplExpiredEmplToJSON)),
        'total_amount': value['totalAmount'],
    };
}

