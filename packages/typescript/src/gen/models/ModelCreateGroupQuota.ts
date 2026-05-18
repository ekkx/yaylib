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
 * @interface ModelCreateGroupQuota
 */
export interface ModelCreateGroupQuota {
    /**
     * 
     * @type {number}
     * @memberof ModelCreateGroupQuota
     */
    remainingQuota?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelCreateGroupQuota
     */
    usedQuota?: number | null;
}

/**
 * Check if a given object implements the ModelCreateGroupQuota interface.
 */
export function instanceOfModelCreateGroupQuota(value: object): value is ModelCreateGroupQuota {
    return true;
}

export function ModelCreateGroupQuotaFromJSON(json: any): ModelCreateGroupQuota {
    return ModelCreateGroupQuotaFromJSONTyped(json, false);
}

export function ModelCreateGroupQuotaFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelCreateGroupQuota {
    if (json == null) {
        return json;
    }
    return {
        
        'remainingQuota': json['remaining_quota'] == null ? undefined : json['remaining_quota'],
        'usedQuota': json['used_quota'] == null ? undefined : json['used_quota'],
    };
}

export function ModelCreateGroupQuotaToJSON(json: any): ModelCreateGroupQuota {
    return ModelCreateGroupQuotaToJSONTyped(json, false);
}

export function ModelCreateGroupQuotaToJSONTyped(value?: ModelCreateGroupQuota | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'remaining_quota': value['remainingQuota'],
        'used_quota': value['usedQuota'],
    };
}

