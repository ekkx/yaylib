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
 * @interface Web3Bag
 */
export interface Web3Bag {
    /**
     * 
     * @type {object}
     * @memberof Web3Bag
     */
    empl?: object | null;
    /**
     * 
     * @type {number}
     * @memberof Web3Bag
     */
    palsCount?: number | null;
}

/**
 * Check if a given object implements the Web3Bag interface.
 */
export function instanceOfWeb3Bag(value: object): value is Web3Bag {
    return true;
}

export function Web3BagFromJSON(json: any): Web3Bag {
    return Web3BagFromJSONTyped(json, false);
}

export function Web3BagFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3Bag {
    if (json == null) {
        return json;
    }
    return {
        
        'empl': json['empl'] == null ? undefined : json['empl'],
        'palsCount': json['pals_count'] == null ? undefined : json['pals_count'],
    };
}

export function Web3BagToJSON(json: any): Web3Bag {
    return Web3BagToJSONTyped(json, false);
}

export function Web3BagToJSONTyped(value?: Web3Bag | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'empl': value['empl'],
        'pals_count': value['palsCount'],
    };
}

