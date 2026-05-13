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
 * @interface ProductQuota
 */
export interface ProductQuota {
    /**
     * 
     * @type {number}
     * @memberof ProductQuota
     */
    bought?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ProductQuota
     */
    limit?: number | null;
}

/**
 * Check if a given object implements the ProductQuota interface.
 */
export function instanceOfProductQuota(value: object): value is ProductQuota {
    return true;
}

export function ProductQuotaFromJSON(json: any): ProductQuota {
    return ProductQuotaFromJSONTyped(json, false);
}

export function ProductQuotaFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProductQuota {
    if (json == null) {
        return json;
    }
    return {
        
        'bought': json['bought'] == null ? undefined : json['bought'],
        'limit': json['limit'] == null ? undefined : json['limit'],
    };
}

export function ProductQuotaToJSON(json: any): ProductQuota {
    return ProductQuotaToJSONTyped(json, false);
}

export function ProductQuotaToJSONTyped(value?: ProductQuota | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'bought': value['bought'],
        'limit': value['limit'],
    };
}

