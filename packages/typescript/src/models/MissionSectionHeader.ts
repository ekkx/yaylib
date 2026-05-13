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
 * @interface MissionSectionHeader
 */
export interface MissionSectionHeader {
    /**
     * 
     * @type {number}
     * @memberof MissionSectionHeader
     */
    sectionInfoId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof MissionSectionHeader
     */
    sectionTitleId?: number | null;
}

/**
 * Check if a given object implements the MissionSectionHeader interface.
 */
export function instanceOfMissionSectionHeader(value: object): value is MissionSectionHeader {
    return true;
}

export function MissionSectionHeaderFromJSON(json: any): MissionSectionHeader {
    return MissionSectionHeaderFromJSONTyped(json, false);
}

export function MissionSectionHeaderFromJSONTyped(json: any, ignoreDiscriminator: boolean): MissionSectionHeader {
    if (json == null) {
        return json;
    }
    return {
        
        'sectionInfoId': json['section_info_id'] == null ? undefined : json['section_info_id'],
        'sectionTitleId': json['section_title_id'] == null ? undefined : json['section_title_id'],
    };
}

export function MissionSectionHeaderToJSON(json: any): MissionSectionHeader {
    return MissionSectionHeaderToJSONTyped(json, false);
}

export function MissionSectionHeaderToJSONTyped(value?: MissionSectionHeader | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'section_info_id': value['sectionInfoId'],
        'section_title_id': value['sectionTitleId'],
    };
}

