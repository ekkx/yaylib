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
 * @interface CreateGroupQuota
 */
export interface CreateGroupQuota {
    /**
     * 
     * @type {number}
     * @memberof CreateGroupQuota
     */
    remainingQuota?: number | null;
    /**
     * 
     * @type {number}
     * @memberof CreateGroupQuota
     */
    usedQuota?: number | null;
}

/**
 * Check if a given object implements the CreateGroupQuota interface.
 */
export function instanceOfCreateGroupQuota(value: object): value is CreateGroupQuota {
    return true;
}

export function CreateGroupQuotaFromJSON(json: any): CreateGroupQuota {
    return CreateGroupQuotaFromJSONTyped(json, false);
}

export function CreateGroupQuotaFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateGroupQuota {
    if (json == null) {
        return json;
    }
    return {
        
        'remainingQuota': json['remaining_quota'] == null ? undefined : json['remaining_quota'],
        'usedQuota': json['used_quota'] == null ? undefined : json['used_quota'],
    };
}

export function CreateGroupQuotaToJSON(json: any): CreateGroupQuota {
    return CreateGroupQuotaToJSONTyped(json, false);
}

export function CreateGroupQuotaToJSONTyped(value?: CreateGroupQuota | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'remaining_quota': value['remainingQuota'],
        'used_quota': value['usedQuota'],
    };
}

