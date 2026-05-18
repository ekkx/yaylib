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
 * @interface ModelCoinAmount
 */
export interface ModelCoinAmount {
    /**
     * 
     * @type {number}
     * @memberof ModelCoinAmount
     */
    free?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelCoinAmount
     */
    paid?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelCoinAmount
     */
    total?: number | null;
}

/**
 * Check if a given object implements the ModelCoinAmount interface.
 */
export function instanceOfModelCoinAmount(value: object): value is ModelCoinAmount {
    return true;
}

export function ModelCoinAmountFromJSON(json: any): ModelCoinAmount {
    return ModelCoinAmountFromJSONTyped(json, false);
}

export function ModelCoinAmountFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelCoinAmount {
    if (json == null) {
        return json;
    }
    return {
        
        'free': json['free'] == null ? undefined : json['free'],
        'paid': json['paid'] == null ? undefined : json['paid'],
        'total': json['total'] == null ? undefined : json['total'],
    };
}

export function ModelCoinAmountToJSON(json: any): ModelCoinAmount {
    return ModelCoinAmountToJSONTyped(json, false);
}

export function ModelCoinAmountToJSONTyped(value?: ModelCoinAmount | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'free': value['free'],
        'paid': value['paid'],
        'total': value['total'],
    };
}

