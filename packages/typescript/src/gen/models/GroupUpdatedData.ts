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
 * @interface GroupUpdatedData
 */
export interface GroupUpdatedData {
    /**
     * 
     * @type {number}
     * @memberof GroupUpdatedData
     */
    groupId?: number | null;
}

/**
 * Check if a given object implements the GroupUpdatedData interface.
 */
export function instanceOfGroupUpdatedData(value: object): value is GroupUpdatedData {
    return true;
}

export function GroupUpdatedDataFromJSON(json: any): GroupUpdatedData {
    return GroupUpdatedDataFromJSONTyped(json, false);
}

export function GroupUpdatedDataFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupUpdatedData {
    if (json == null) {
        return json;
    }
    return {
        
        'groupId': json['group_id'] == null ? undefined : json['group_id'],
    };
}

export function GroupUpdatedDataToJSON(json: any): GroupUpdatedData {
    return GroupUpdatedDataToJSONTyped(json, false);
}

export function GroupUpdatedDataToJSONTyped(value?: GroupUpdatedData | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'group_id': value['groupId'],
    };
}

