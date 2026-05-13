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
 * @interface ExpiredEmpl
 */
export interface ExpiredEmpl {
    /**
     * 
     * @type {string}
     * @memberof ExpiredEmpl
     */
    expiringAmount?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ExpiredEmpl
     */
    expiringDate?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof ExpiredEmpl
     */
    soon?: boolean | null;
}

/**
 * Check if a given object implements the ExpiredEmpl interface.
 */
export function instanceOfExpiredEmpl(value: object): value is ExpiredEmpl {
    return true;
}

export function ExpiredEmplFromJSON(json: any): ExpiredEmpl {
    return ExpiredEmplFromJSONTyped(json, false);
}

export function ExpiredEmplFromJSONTyped(json: any, ignoreDiscriminator: boolean): ExpiredEmpl {
    if (json == null) {
        return json;
    }
    return {
        
        'expiringAmount': json['expiring_amount'] == null ? undefined : json['expiring_amount'],
        'expiringDate': json['expiring_date'] == null ? undefined : json['expiring_date'],
        'soon': json['soon'] == null ? undefined : json['soon'],
    };
}

export function ExpiredEmplToJSON(json: any): ExpiredEmpl {
    return ExpiredEmplToJSONTyped(json, false);
}

export function ExpiredEmplToJSONTyped(value?: ExpiredEmpl | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'expiring_amount': value['expiringAmount'],
        'expiring_date': value['expiringDate'],
        'soon': value['soon'],
    };
}

