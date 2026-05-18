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
import type { MissionItem } from './MissionItem.js';
import {
    MissionItemFromJSON,
    MissionItemFromJSONTyped,
    MissionItemToJSON,
    MissionItemToJSONTyped,
} from './MissionItem.js';
import type { MissionSectionHeader } from './MissionSectionHeader.js';
import {
    MissionSectionHeaderFromJSON,
    MissionSectionHeaderFromJSONTyped,
    MissionSectionHeaderToJSON,
    MissionSectionHeaderToJSONTyped,
} from './MissionSectionHeader.js';

/**
 * 
 * @export
 * @interface MissionSection
 */
export interface MissionSection {
    /**
     * 
     * @type {Array<MissionItem>}
     * @memberof MissionSection
     */
    missionList?: Array<MissionItem> | null;
    /**
     * 
     * @type {object}
     * @memberof MissionSection
     */
    missionType?: object | null;
    /**
     * 
     * @type {MissionSectionHeader}
     * @memberof MissionSection
     */
    section?: MissionSectionHeader | null;
}

/**
 * Check if a given object implements the MissionSection interface.
 */
export function instanceOfMissionSection(value: object): value is MissionSection {
    return true;
}

export function MissionSectionFromJSON(json: any): MissionSection {
    return MissionSectionFromJSONTyped(json, false);
}

export function MissionSectionFromJSONTyped(json: any, ignoreDiscriminator: boolean): MissionSection {
    if (json == null) {
        return json;
    }
    return {
        
        'missionList': json['mission_list'] == null ? undefined : ((json['mission_list'] as Array<any>).map(MissionItemFromJSON)),
        'missionType': json['mission_type'] == null ? undefined : json['mission_type'],
        'section': json['section'] == null ? undefined : MissionSectionHeaderFromJSON(json['section']),
    };
}

export function MissionSectionToJSON(json: any): MissionSection {
    return MissionSectionToJSONTyped(json, false);
}

export function MissionSectionToJSONTyped(value?: MissionSection | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'mission_list': value['missionList'] == null ? undefined : ((value['missionList'] as Array<any>).map(MissionItemToJSON)),
        'mission_type': value['missionType'],
        'section': MissionSectionHeaderToJSON(value['section']),
    };
}

