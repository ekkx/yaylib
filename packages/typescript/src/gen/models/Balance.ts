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
 * @interface Balance
 */
export interface Balance {
    /**
     * 
     * @type {string}
     * @memberof Balance
     */
    egg?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Balance
     */
    empl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Balance
     */
    l1Native?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Balance
     */
    l1Yay?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Balance
     */
    _native?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Balance
     */
    pal?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Balance
     */
    yay?: string | null;
}

/**
 * Check if a given object implements the Balance interface.
 */
export function instanceOfBalance(value: object): value is Balance {
    return true;
}

export function BalanceFromJSON(json: any): Balance {
    return BalanceFromJSONTyped(json, false);
}

export function BalanceFromJSONTyped(json: any, ignoreDiscriminator: boolean): Balance {
    if (json == null) {
        return json;
    }
    return {
        
        'egg': json['egg'] == null ? undefined : json['egg'],
        'empl': json['empl'] == null ? undefined : json['empl'],
        'l1Native': json['l1_native'] == null ? undefined : json['l1_native'],
        'l1Yay': json['l1_yay'] == null ? undefined : json['l1_yay'],
        '_native': json['native'] == null ? undefined : json['native'],
        'pal': json['pal'] == null ? undefined : json['pal'],
        'yay': json['yay'] == null ? undefined : json['yay'],
    };
}

export function BalanceToJSON(json: any): Balance {
    return BalanceToJSONTyped(json, false);
}

export function BalanceToJSONTyped(value?: Balance | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'egg': value['egg'],
        'empl': value['empl'],
        'l1_native': value['l1Native'],
        'l1_yay': value['l1Yay'],
        'native': value['_native'],
        'pal': value['pal'],
        'yay': value['yay'],
    };
}

