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
/**
 * 
 * @export
 * @interface Web3ExpiringEmplExpiredEmpl
 */
export interface Web3ExpiringEmplExpiredEmpl {
    /**
     * 
     * @type {number}
     * @memberof Web3ExpiringEmplExpiredEmpl
     */
    amount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Web3ExpiringEmplExpiredEmpl
     */
    dateInMillis?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof Web3ExpiringEmplExpiredEmpl
     */
    soon?: boolean | null;
}

/**
 * Check if a given object implements the Web3ExpiringEmplExpiredEmpl interface.
 */
export function instanceOfWeb3ExpiringEmplExpiredEmpl(value: object): value is Web3ExpiringEmplExpiredEmpl {
    return true;
}

export function Web3ExpiringEmplExpiredEmplFromJSON(json: any): Web3ExpiringEmplExpiredEmpl {
    return Web3ExpiringEmplExpiredEmplFromJSONTyped(json, false);
}

export function Web3ExpiringEmplExpiredEmplFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3ExpiringEmplExpiredEmpl {
    if (json == null) {
        return json;
    }
    return {
        
        'amount': json['amount'] == null ? undefined : json['amount'],
        'dateInMillis': json['date_in_millis'] == null ? undefined : json['date_in_millis'],
        'soon': json['soon'] == null ? undefined : json['soon'],
    };
}

export function Web3ExpiringEmplExpiredEmplToJSON(json: any): Web3ExpiringEmplExpiredEmpl {
    return Web3ExpiringEmplExpiredEmplToJSONTyped(json, false);
}

export function Web3ExpiringEmplExpiredEmplToJSONTyped(value?: Web3ExpiringEmplExpiredEmpl | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'amount': value['amount'],
        'date_in_millis': value['dateInMillis'],
        'soon': value['soon'],
    };
}

