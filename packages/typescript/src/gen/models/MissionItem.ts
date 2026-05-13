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
 * @interface MissionItem
 */
export interface MissionItem {
    /**
     * 
     * @type {object}
     * @memberof MissionItem
     */
    mission?: object | null;
}

/**
 * Check if a given object implements the MissionItem interface.
 */
export function instanceOfMissionItem(value: object): value is MissionItem {
    return true;
}

export function MissionItemFromJSON(json: any): MissionItem {
    return MissionItemFromJSONTyped(json, false);
}

export function MissionItemFromJSONTyped(json: any, ignoreDiscriminator: boolean): MissionItem {
    if (json == null) {
        return json;
    }
    return {
        
        'mission': json['mission'] == null ? undefined : json['mission'],
    };
}

export function MissionItemToJSON(json: any): MissionItem {
    return MissionItemToJSONTyped(json, false);
}

export function MissionItemToJSONTyped(value?: MissionItem | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'mission': value['mission'],
    };
}

