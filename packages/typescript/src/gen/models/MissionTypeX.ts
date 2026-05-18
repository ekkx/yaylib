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
 * @interface MissionTypeX
 */
export interface MissionTypeX {
    /**
     * 
     * @type {number}
     * @memberof MissionTypeX
     */
    sectionInfoId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof MissionTypeX
     */
    sectionTitleId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof MissionTypeX
     */
    titleId?: number | null;
}

/**
 * Check if a given object implements the MissionTypeX interface.
 */
export function instanceOfMissionTypeX(value: object): value is MissionTypeX {
    return true;
}

export function MissionTypeXFromJSON(json: any): MissionTypeX {
    return MissionTypeXFromJSONTyped(json, false);
}

export function MissionTypeXFromJSONTyped(json: any, ignoreDiscriminator: boolean): MissionTypeX {
    if (json == null) {
        return json;
    }
    return {
        
        'sectionInfoId': json['section_info_id'] == null ? undefined : json['section_info_id'],
        'sectionTitleId': json['section_title_id'] == null ? undefined : json['section_title_id'],
        'titleId': json['title_id'] == null ? undefined : json['title_id'],
    };
}

export function MissionTypeXToJSON(json: any): MissionTypeX {
    return MissionTypeXToJSONTyped(json, false);
}

export function MissionTypeXToJSONTyped(value?: MissionTypeX | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'section_info_id': value['sectionInfoId'],
        'section_title_id': value['sectionTitleId'],
        'title_id': value['titleId'],
    };
}

