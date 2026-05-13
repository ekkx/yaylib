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
import type { MissionType } from './MissionType';
import {
    MissionTypeFromJSON,
    MissionTypeFromJSONTyped,
    MissionTypeToJSON,
    MissionTypeToJSONTyped,
} from './MissionType';

/**
 * 
 * @export
 * @interface Mission
 */
export interface Mission {
    /**
     * 
     * @type {string}
     * @memberof Mission
     */
    action?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Mission
     */
    detail?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof Mission
     */
    isMultiplier?: boolean | null;
    /**
     * 
     * @type {MissionType}
     * @memberof Mission
     */
    missionType?: MissionType | null;
    /**
     * 
     * @type {number}
     * @memberof Mission
     */
    requiredActionsCount?: number | null;
}



/**
 * Check if a given object implements the Mission interface.
 */
export function instanceOfMission(value: object): value is Mission {
    return true;
}

export function MissionFromJSON(json: any): Mission {
    return MissionFromJSONTyped(json, false);
}

export function MissionFromJSONTyped(json: any, ignoreDiscriminator: boolean): Mission {
    if (json == null) {
        return json;
    }
    return {
        
        'action': json['action'] == null ? undefined : json['action'],
        'detail': json['detail'] == null ? undefined : json['detail'],
        'isMultiplier': json['is_multiplier'] == null ? undefined : json['is_multiplier'],
        'missionType': json['mission_type'] == null ? undefined : MissionTypeFromJSON(json['mission_type']),
        'requiredActionsCount': json['required_actions_count'] == null ? undefined : json['required_actions_count'],
    };
}

export function MissionToJSON(json: any): Mission {
    return MissionToJSONTyped(json, false);
}

export function MissionToJSONTyped(value?: Mission | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'action': value['action'],
        'detail': value['detail'],
        'is_multiplier': value['isMultiplier'],
        'mission_type': MissionTypeToJSON(value['missionType']),
        'required_actions_count': value['requiredActionsCount'],
    };
}

