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
 * @interface OnlineStatus
 */
export interface OnlineStatus {
    /**
     * 
     * @type {string}
     * @memberof OnlineStatus
     */
    apiValue?: string | null;
}

/**
 * Check if a given object implements the OnlineStatus interface.
 */
export function instanceOfOnlineStatus(value: object): value is OnlineStatus {
    return true;
}

export function OnlineStatusFromJSON(json: any): OnlineStatus {
    return OnlineStatusFromJSONTyped(json, false);
}

export function OnlineStatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): OnlineStatus {
    if (json == null) {
        return json;
    }
    return {
        
        'apiValue': json['api_value'] == null ? undefined : json['api_value'],
    };
}

export function OnlineStatusToJSON(json: any): OnlineStatus {
    return OnlineStatusToJSONTyped(json, false);
}

export function OnlineStatusToJSONTyped(value?: OnlineStatus | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'api_value': value['apiValue'],
    };
}

