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
 * @interface ModelProductQuota
 */
export interface ModelProductQuota {
    /**
     * 
     * @type {number}
     * @memberof ModelProductQuota
     */
    bought?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ModelProductQuota
     */
    currencyCode?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ModelProductQuota
     */
    limit?: number | null;
}

/**
 * Check if a given object implements the ModelProductQuota interface.
 */
export function instanceOfModelProductQuota(value: object): value is ModelProductQuota {
    return true;
}

export function ModelProductQuotaFromJSON(json: any): ModelProductQuota {
    return ModelProductQuotaFromJSONTyped(json, false);
}

export function ModelProductQuotaFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelProductQuota {
    if (json == null) {
        return json;
    }
    return {
        
        'bought': json['bought'] == null ? undefined : json['bought'],
        'currencyCode': json['currency_code'] == null ? undefined : json['currency_code'],
        'limit': json['limit'] == null ? undefined : json['limit'],
    };
}

export function ModelProductQuotaToJSON(json: any): ModelProductQuota {
    return ModelProductQuotaToJSONTyped(json, false);
}

export function ModelProductQuotaToJSONTyped(value?: ModelProductQuota | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'bought': value['bought'],
        'currency_code': value['currencyCode'],
        'limit': value['limit'],
    };
}

