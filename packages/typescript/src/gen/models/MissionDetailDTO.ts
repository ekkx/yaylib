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
 * @interface MissionDetailDTO
 */
export interface MissionDetailDTO {
    /**
     * 
     * @type {string}
     * @memberof MissionDetailDTO
     */
    action?: string | null;
    /**
     * 
     * @type {string}
     * @memberof MissionDetailDTO
     */
    detail?: string | null;
    /**
     * 
     * @type {number}
     * @memberof MissionDetailDTO
     */
    missionId?: number | null;
}

/**
 * Check if a given object implements the MissionDetailDTO interface.
 */
export function instanceOfMissionDetailDTO(value: object): value is MissionDetailDTO {
    return true;
}

export function MissionDetailDTOFromJSON(json: any): MissionDetailDTO {
    return MissionDetailDTOFromJSONTyped(json, false);
}

export function MissionDetailDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): MissionDetailDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'action': json['action'] == null ? undefined : json['action'],
        'detail': json['detail'] == null ? undefined : json['detail'],
        'missionId': json['mission_id'] == null ? undefined : json['mission_id'],
    };
}

export function MissionDetailDTOToJSON(json: any): MissionDetailDTO {
    return MissionDetailDTOToJSONTyped(json, false);
}

export function MissionDetailDTOToJSONTyped(value?: MissionDetailDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'action': value['action'],
        'detail': value['detail'],
        'mission_id': value['missionId'],
    };
}

